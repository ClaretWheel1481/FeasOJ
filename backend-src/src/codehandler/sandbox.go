package codehandler

import (
	"context"
	"fmt"
	"io"
	"os"
	"src/global"

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
func StartContainer() bool {
	// 创建context
	ctx := context.Background()

	// 创建Docker客户端
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
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
	}

	// 创建容器
	resp, err := cli.ContainerCreate(ctx, containerConfig, hostConfig, nil, nil, "")
	if err != nil {
		panic(err)
	}

	// 启动容器
	if err := cli.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
		panic(err)
	}

	out, err := cli.ContainerLogs(ctx, resp.ID, container.LogsOptions{ShowStdout: true})
	if err != nil {
		panic(err)
	}
	defer out.Close()
	io.Copy(os.Stdout, out)
	global.ContainerID = resp.ID
	return true
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

	// 删除容器
	if err := cli.ContainerRemove(ctx, containerID, container.RemoveOptions{}); err != nil {
		panic(err)
	}

	return true
}
