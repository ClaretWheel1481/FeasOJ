[English](README.md) | 简体中文
<p align="center">
    <a href="https://github.com/ClaretWheel1481/easyweather">
        <img src="./public/logo.png" height="200"/>
    </a>
</p>

# FeasOJ
### 项目简介
FeasOJ 是一个基于 Vue 和 Golang 的在线编程练习平台，旨在为用户提供一个方便、高效的学习和练习环境。

### 项目结构
- src：前端代码
  - assets：静态资源
  - components：组件
  - pages：页面
  - routers：路由
  - utils：工具函数
  - plugins：插件

- backend-src：后端代码
  - src：主要代码文件夹
    - account：用户模块
    - codehandler：代码处理模块
    - ginrouter：路由模块
    - global：全局变量
    - utils：工具函数
    - main.go：主文件
    - go.mod：依赖管理文件
  - Sandbox: Dockerfile

### 项目截图
![image](/public/Screenshot_1.png)
![image](/public/Screenshot_2.png)

### 如何运行

1. 克隆此库。
2. 安装 Docker。
3. 运行 `npm install -g pnpm` 安装 Pnpm。
4. 运行 `pnpm i` 安装依赖项。
5. 运行 `pnpm dev` 启动前端服务器。
6. 运行 `cd backend-src\src` 和 `go mod tidy` 下载依赖。
7. 运行 `go run main.go` 启动后端服务器。

### 注意

这是我第一次用Vue + Golang写大项目，所以代码会很糟糕，不过我会一直去改进它！
如果你找到任何Bug请发布Issue告诉我。

### 致谢

- [Vue](https://github.com/vuejs/vue)
- [Vuetify](https://github.com/vuetifyjs/vuetify)
- [vue-avatar-cropper](https://github.com/overtrue/vue-avatar-cropper)
- [md-editor-v3](https://github.com/imzbf/md-editor-v3)
- [axios](https://github.com/axios/axios)
- [Go](https://github.com/golang/go)
- [Gin](https://github.com/gin-gonic/gin)
- [gorm](https://github.com/go-gorm/gorm)
- [jwt-go](https://github.com/golang-jwt/jwt)
- [docker](https://github.com/moby/moby)
- [gomail](https://github.com/go-gomail/gomail)
- [go-redis](https://github.com/redis/go-redis)