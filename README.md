English | [简体中文](README_CN.md)
<p align="center">
    <a href="https://github.com/ClaretWheel1481/FeasOJ">
        <img src="public/logo.png" height="200"/>
    </a>
</p>

# FeasOJ
### Project Description
FeasOJ is an online programming practice platform based on Vue and Golang, aiming to provide users with a convenient and efficient learning and practice environment.
[Demo](https://oj.claret.space)

### Project Structure
- src：Frontend
  - assets：
  - components：
  - pages：
  - routers：
  - utils：
  - plugins：

- backend-src：Backend
  - src：
    - account：
    - codehandler：
    - ginrouter：
    - global：
    - utils：
    - main.go：
    - go.mod：
  - Sandbox: Dockerfile

### Screenshots
![image](/public/Screenshot_1.png)
![image](/public/Screenshot_2.png)

### How to run

1. Clone this repository.
2. Install Docker.
3. Run `npm install -g pnpm` Install Pnpm.
4. Run `pnpm i` Install dependencies.
5. Run `pnpm dev` to start the front-end server.
6. Run `cd backend-src\src` and `go mod tiny` to get the dependencies.
7. Run `go run main.go` to start the backend server.

### Notice

This is the first time I've written a big project with Vue + Golang, so the code is going to be terrible, but I'll keep going to improve it!
If you find any bugs, please open an issue.


### Thanks

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