English | [简体中文](README_CN.md)
<p align="center">
    <a href="https://github.com/ClaretWheel1481/FeasOJ">
        <img src="public/logo.png" height="200"/>
    </a>
</p>

# FeasOJ
### Project Description
FeasOJ is an online programming practice platform based on Vue and Golang, aiming to provide users with a convenient and efficient learning and practice environment.

### Project Structure
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

### Environment
- Vue 3
- Pnpm
- Golang 1.23.1
- Redis
- MySQL 8.0+
- Docker

### How to run

1. Clone repository.
2. Install Docker.
3. Run `npm install -g pnpm` Install Pnpm.
4. Run `pnpm i` Install dependencies.
5. Run `pnpm dev` to start the front-end server.
6. Run `cd backend-src\src` and `go mod tiny` to get the dependencies.
7. Run `go run main.go` to start the backend server.

### Notice

This is the first time I've written a big project with Vue + Golang, so the code is going to be terrible, but I'll keep going to improve it!
If you find any bugs, please open an issue.

### Screenshots
![image](/assets/Screenshot1.png)
![image](/assets/Screenshot2.png)

### Thanks

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