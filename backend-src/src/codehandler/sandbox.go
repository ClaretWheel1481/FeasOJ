package codehandler

import (
	"context"
	"fmt"
	"io"
	"log"
	"os/exec"
	"path/filepath"
	"src/global"
	"src/utils/sql"
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
		return false
	}
	// TODO: 每次编译前需要修改为CurrentDir，debug时用ParentDir
	// 将Dockerfile目录打包成tar格式
	tar, err := archive.TarWithOptions(global.ParentDir, &archive.TarOptions{})
	if err != nil {
		log.Panic(err)
	}

	// 设置镜像构建选项
	buildOptions := types.ImageBuildOptions{
		Context:    tar,                          // 构建上下文
		Dockerfile: "Sandbox",                    // Dockerfile文件名
		Tags:       []string{"judgecore:latest"}, // 镜像标签
	}

	log.Println("[FeasOJ]SandBox is being built...")
	// 构建Docker镜像
	buildResponse, err := cli.ImageBuild(ctx, tar, buildOptions)
	if err != nil {
		return false
	}
	defer buildResponse.Body.Close()

	// 打印构建响应
	_, err = io.Copy(log.Writer(), buildResponse.Body)
	if err != nil {
		log.Printf("Error copying build response: %v", err)
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
		Cmd:   []string{"sh"},
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
		compileCmd = exec.Command("docker", "exec", global.ContainerID, "sh", "-c", fmt.Sprintf("g++ /workspace/%s -o /workspace/a.out", filename))
		if err := compileCmd.Run(); err != nil {
			TerminateContainer(global.ContainerID)
			return "Compile Failed"
		}
	case ".java":
		// 临时重命名为Main.java
		originalName := filename
		tempName := "Main.java"
		renameCmd := exec.Command("docker", "exec", global.ContainerID, "sh", "-c", fmt.Sprintf("mv /workspace/%s /workspace/%s", originalName, tempName))
		if err := renameCmd.Run(); err != nil {
			TerminateContainer(global.ContainerID)
			return "Compile Failed"
		}

		compileCmd = exec.Command("docker", "exec", global.ContainerID, "sh", "-c", fmt.Sprintf("javac /workspace/%s", tempName))
		if err := compileCmd.Run(); err != nil {
			TerminateContainer(global.ContainerID)
			return "Compile Failed"
		}

		// 编译完成后改回原名
		renameBackCmd := exec.Command("docker", "exec", global.ContainerID, "sh", "-c", fmt.Sprintf("mv /workspace/%s /workspace/%s", tempName, originalName))
		if err := renameBackCmd.Run(); err != nil {
			TerminateContainer(global.ContainerID)
			return "Compile Failed"
		}
	}

	// 设置超时上下文
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 获取输入输出样例
	testCases := sql.SelectTestCasesByPid(strings.Split(filename, "_")[1])
	for _, testCase := range testCases {
		var runCmd *exec.Cmd
		switch ext {
		case ".cpp":
			runCmd = exec.CommandContext(ctx, "docker", "exec", "-i", global.ContainerID, "sh", "-c", "/workspace/a.out")
		case ".py":
			runCmd = exec.CommandContext(ctx, "docker", "exec", "-i", global.ContainerID, "sh", "-c", fmt.Sprintf("python /workspace/%s", filename))
		case ".go":
			runCmd = exec.CommandContext(ctx, "docker", "exec", "-i", global.ContainerID, "sh", "-c", fmt.Sprintf("go run /workspace/%s", filename))
		case ".java":
			runCmd = exec.CommandContext(ctx, "docker", "exec", "-i", global.ContainerID, "sh", "-c", "java Main")
		default:
			TerminateContainer(global.ContainerID)
			return "Failed"
		}

		runCmd.Stdin = strings.NewReader(testCase.InputData)
		output, err := runCmd.CombinedOutput()
		if ctx.Err() == context.DeadlineExceeded {
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
