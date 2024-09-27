[English](README.md) | 简体中文
<p align="center">
    <a href="https://github.com/ClaretWheel1481/FeasOJ">
        <img src="public/logo.png" height="200"/>
    </a>
</p>

# FeasOJ
### 项目简介
FeasOJ 是一个基于 Vue 和 Golang 的在线编程练习平台，旨在为用户提供一个方便、高效的学习和练习环境。

### 项目结构
```
FeasOJ
│ 
├─src
│  ├─public
│  ├─components
│  ├─pages
│  ├─routers
│  ├─utils
│  └─plugins
│ 
├─backend-src
│  ├─src
│  │  ├─codehandler
│  │  ├─config
│  │  ├─gin
│  │  ├─global
│  │  ├─middleware
│  │  ├─utils
│  │  │  └─sql
│  │  ├─main.go
│  │  └─go.mod
│  └─Sandbox
```

### 软件环境
- Vue 3
- Pnpm
- Golang 1.23.1
- Redis
- MySQL 8.0+
- Docker

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

### 项目截图
![image](/assets/Screenshot1.png)
![image](/assets/Screenshot2.png)

### 致谢

- [Vue](https://github.com/vuejs/vue)
- [Vuetify](https://github.com/vuetifyjs/vuetify)
- [vue-avatar-cropper](https://github.com/overtrue/vue-avatar-cropper)
- [md-editor-v3](https://github.com/imzbf/md-editor-v3)
- [vue-i18n](https://github.com/intlify/vue-i18n)
- [axios](https://github.com/axios/axios)
- [Go](https://github.com/golang/go)
- [Gin](https://github.com/gin-gonic/gin)
- [gorm](https://github.com/go-gorm/gorm)
- [jwt-go](https://github.com/golang-jwt/jwt)
- [docker](https://github.com/moby/moby)
- [gomail](https://github.com/go-gomail/gomail)
- [go-redis](https://github.com/redis/go-redis)