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
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
)

// 构建Sandbox
func BuildImage() bool {
	ctx := context.Background()

	// 创建Docker客户端
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		fmt.Println("[FeasOJ]请确认Docker是否在本机安装，并处于启动状态！")
		return false
	}
	// TODO: 每次编译前需要修改为CurrentDir，debug时用ParentDir
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
		fmt.Println("[FeasOJ]请确认Docker是否在本机安装，并处于启动状态！")
		return false
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
	ctx := context.Background()

	// 创建Docker客户端
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return "", err
	}

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
			NanoCPUs: 1 * 1e9,            // CPU限制为1核
		},
		Binds: []string{
			global.CodeDir + ":/workspace", // 挂载文件夹
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

// 启动容器并编译运行文件、放入输入、捕获输出、对照输出
func CompileAndRun(filename string) string {
	ext := filepath.Ext(filename)
	var compileCmd *exec.Cmd

	switch ext {
	case ".cpp":
		compileCmd = exec.Command("docker", "exec", global.ContainerID, "g++", fmt.Sprintf("/workspace/%s", filename), "-o", "/workspace/a.out")
		if err := compileCmd.Run(); err != nil {
			TerminateContainer(global.ContainerID)
			return "Compile Failed"
		}
	case ".java":
		// 临时重命名为Main.java
		originalName := filename
		tempName := "Main.java"
		renameCmd := exec.Command("docker", "exec", global.ContainerID, "mv", fmt.Sprintf("/workspace/%s", originalName), fmt.Sprintf("/workspace/%s", tempName))
		if err := renameCmd.Run(); err != nil {
			TerminateContainer(global.ContainerID)
			return "Compile Failed"
		}

		compileCmd = exec.Command("docker", "exec", global.ContainerID, "javac", fmt.Sprintf("/workspace/%s", tempName))
		if err := compileCmd.Run(); err != nil {
			TerminateContainer(global.ContainerID)
			return "Compile Failed"
		}

		// 编译完成后改回原名
		renameBackCmd := exec.Command("docker", "exec", global.ContainerID, "mv", fmt.Sprintf("/workspace/%s", tempName), fmt.Sprintf("/workspace/%s", originalName))
		if err := renameBackCmd.Run(); err != nil {
			TerminateContainer(global.ContainerID)
			return "Compile Failed"
		}
	}

	// 设置超时上下文
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 获取输入输出样例
	testCases := utils.SelectTestCasesByPid(strings.Split(filename, "_")[1])
	for _, testCase := range testCases {
		var runCmd *exec.Cmd
		switch ext {
		case ".cpp":
			runCmd = exec.CommandContext(ctx, "docker", "exec", "-i", global.ContainerID, "/workspace/a.out")
		case ".py":
			runCmd = exec.CommandContext(ctx, "docker", "exec", "-i", global.ContainerID, "python", fmt.Sprintf("/workspace/%s", filename))
		case ".go":
			runCmd = exec.CommandContext(ctx, "docker", "exec", "-i", global.ContainerID, "go", "run", fmt.Sprintf("/workspace/%s", filename))
		case ".java":
			runCmd = exec.CommandContext(ctx, "docker", "exec", "-i", global.ContainerID, "java", "Main")
		default:
			TerminateContainer(global.ContainerID)
			return "Failed"
		}

		runCmd.Stdin = strings.NewReader(testCase.InputData)
		output, err := runCmd.CombinedOutput()
		if ctx.Err() == context.DeadlineExceeded {
			// 超时10s后强制停止容器
			TerminateContainer(global.ContainerID)
			return "Time Limit Exceeded"
		}
		if err != nil {
			TerminateContainer(global.ContainerID)
			return "Failed"
		}
		outputStr := string(output)
		if strings.TrimSpace(outputStr) != strings.TrimSpace(testCase.OutputData) {
			TerminateContainer(global.ContainerID)
			return "Wrong Answer"
		}
	}
	TerminateContainer(global.ContainerID)
	return "Success"
}

// 终止并删除Docker容器
func TerminateContainer(containerID string) bool {
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
