package codehandler

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"src/global"
	"src/utils"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
)

// 构建Sandbox
func BuildImage() bool {
	// 创建context
	ctx := context.Background()

	// 创建Docker客户端
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	// 将Dockerfile目录打包成tar格式
	tar, err := archive.TarWithOptions(global.ParentDir, &archive.TarOptions{})
	if err != nil {
		panic(err)
	}

	// 设置镜像构建选项
	buildOptions := types.ImageBuildOptions{
		Context:    tar,                          // 构建上下文
		Dockerfile: "Sandbox",                    // Dockerfile文件名
		Tags:       []string{"judgecore:latest"}, // 镜像标签
	}

	fmt.Println("[FeasOJ]正在构建SandBox...")
	// 构建Docker镜像
	buildResponse, err := cli.ImageBuild(ctx, tar, buildOptions)
	if err != nil {
		panic(err)
	}
	defer buildResponse.Body.Close()

	// 打印构建响应
	_, err = io.Copy(os.Stdout, buildResponse.Body)
	if err != nil {
		panic(err)
	}
	return true
}

// 启动Docker容器
func StartContainer() (string, error) {
	// 创建context
	ctx := context.Background()

	// 创建Docker客户端
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return "", err
	}
	fmt.Println("[FeasOJ]正在启动SandBox...")

	// 配置容器配置
	containerConfig := &container.Config{
		Image: "judgecore:latest",
		Cmd:   []string{"bash"},
		Tty:   true,
	}

	// 配置主机配置
	hostConfig := &container.HostConfig{
		Resources: container.Resources{
			Memory:   1024 * 1024 * 1024, // 内存限制为1GB
			NanoCPUs: 2 * 1e9,            // CPU限制为2核
		},
		Binds: []string{
			global.CodeDir + ":/workspace", // 绑定卷
		},
		AutoRemove: true, // 容器退出后自动删除
	}

	// 创建容器
	resp, err := cli.ContainerCreate(ctx, containerConfig, hostConfig, nil, nil, "")
	if err != nil {
		return "", err
	}

	// 启动容器
	if err := cli.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
		return "", err
	}

	global.ContainerID = resp.ID
	return resp.ID, nil
}

// TODO: 待测试 - 启动容器并编译运行文件、放入输入、捕获输出、对照输出
func CompileAndRun(filename string) (bool, string) {
	ext := filepath.Ext(filename)
	var compileCmd *exec.Cmd

	// FIXME: Python找不到路径
	switch ext {
	case ".cpp":
		compileCmd = exec.Command("docker", "exec", global.ContainerID, "g++", fmt.Sprintf("/workspace/%s", filename), "-o", "/workspace/a.out")
		if err := compileCmd.Run(); err != nil {
			return false, "Compile Failed"
		}
		// Java果然是我最讨厌的语言，没有之一!!!!
	case ".java":
		// 临时重命名为Main.java
		originalName := filename
		tempName := "Main.java"
		renameCmd := exec.Command("docker", "exec", global.ContainerID, "mv", fmt.Sprintf("/workspace/%s", originalName), fmt.Sprintf("/workspace/%s", tempName))
		if err := renameCmd.Run(); err != nil {
			return false, "Compile Failed"
		}

		compileCmd = exec.Command("docker", "exec", global.ContainerID, "javac", fmt.Sprintf("/workspace/%s", tempName))
		output, err := compileCmd.CombinedOutput()
		if err != nil {
			fmt.Println("Error compiling Java code:", err)
			fmt.Println("Compile output:", string(output))
			return false, "Compile Failed"
		}

		// 编译完成后改回原名
		renameBackCmd := exec.Command("docker", "exec", global.ContainerID, "mv", fmt.Sprintf("/workspace/%s", tempName), fmt.Sprintf("/workspace/%s", originalName))
		if err := renameBackCmd.Run(); err != nil {
			return false, "Compile Failed"
		}
	}
	// 获取输入输出样例
	testCases := utils.SelectTestCasesByPid(strings.Split(filename, "_")[1])
	for _, testCase := range testCases {
		var runCmd *exec.Cmd
		switch ext {
		case ".cpp":
			runCmd = exec.Command("docker", "exec", "-i", global.ContainerID, "/workspace/a.out")
		case ".py":
			runCmd = exec.Command("docker", "exec", "-i", global.ContainerID, "python", fmt.Sprintf("/workspace/%s", filename))
		case ".go":
			runCmd = exec.Command("docker", "exec", "-i", global.ContainerID, "go", "run", fmt.Sprintf("/workspace/%s", filename))
		case ".java":
			runCmd = exec.Command("docker", "exec", "-i", global.ContainerID, "java", "Main")
		default:
			return false, "Failed"
		}

		// TODO: 测试完成后记得把输出删除！！！
		fmt.Println("Input Data:", testCase.InputData)
		runCmd.Stdin = strings.NewReader(testCase.InputData)
		output, err := runCmd.CombinedOutput()
		fmt.Println(string(output))
		if err != nil {
			return false, "Failed"
		}
		outputStr := string(output)
		fmt.Println("Output:", outputStr)
		if strings.TrimSpace(outputStr) != strings.TrimSpace(testCase.OutputData) {
			fmt.Println("Test case failed. Expected:", testCase.OutputData, "Got:", outputStr)
			return false, "Failed"
		}
	}

	return true, "Success"
}

// 终止并删除Docker容器
func TerminateContainer(containerID string) bool {
	// 创建context
	ctx := context.Background()

	// 创建Docker客户端
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	// 终止容器
	if err := cli.ContainerStop(ctx, containerID, container.StopOptions{}); err != nil {
		panic(err)
	}

	return true
}
