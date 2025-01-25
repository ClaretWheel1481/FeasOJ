[English](README.md) | 简体中文
<p align="center">
    <a href="https://github.com/ClaretWheel1481/FeasOJ">
        <img src="public/logo.png" height="200"/>
    </a>
</p>

# FeasOJ
### 项目简介
FeasOJ 是一个基于 Vue 和 Golang 的在线编程练习平台，支持多国语言、讨论区、竞赛等功能，旨在为用户提供一个方便、高效的学习和练习环境。
<br>
[FeasOJ-Backend](https://github.com/ClaretWheel1481/FeasOJ-Backend)
[ImageGuard](https://github.com/ClaretWheel1481/ImageGuard)

### 项目结构
```
FeasOJ
│ 
├─public
├─src
│  ├─assets
│  ├─components
│  ├─pages
│  ├─plugins
│  ├─router
│  └─utils
```

### 环境
- Vue 3
- pnpm
- 最新版本的Chromium系浏览器或Firefox

### 如何运行
1. 克隆此库。
2. 运行 `pnpm i` 安装依赖项。
3. 配置 `src/utils/axios.js` 中的 `apiUrl` 为你的[FeasOJ-Bakcend](https://github.com/ClaretWheel1481/FeasOJ-Backend)服务器地址。
4. 运行 `pnpm dev` 启动前端服务器。

### 注意
这是我第一次用Vue + Golang写大项目，所以代码会很糟糕，不过我会一直去改进它！
如果你找到任何Bug请发布Issue告诉我。

### 本土化
- English
- Español
- Français
- Italiano
- 日本語
- 简体中文
- 繁體中文

如果您想要增加语言翻译或优化当前的翻译，请按照以下步骤：
- [Fork](https://github.com/ClaretWheel1481/FeasOJ/fork) 该项目仓库
- 复制 `src/plugins/locales/en.js` 文件并以您想要的语言代码作为文件名称，粘贴到 `src/plugins/locales`，或者直接修改您想优化的文件
- 翻译文件中的内容
- 创建一个 [pull request](https://github.com/ClaretWheel1481/FeasOJ/pulls) 即可

### 项目截图
![Main](/assets/Main.png)
![Login](/assets/Login.png)
![Problem](/assets/Problem.png)
![Profile](/assets/Profile.png)
更多图片可在 [assets](/assets) 中查看。

### 致谢
- [Vue](https://github.com/vuejs/vue)
- [Vuetify](https://github.com/vuetifyjs/vuetify)
- [vue-avatar-cropper](https://github.com/overtrue/vue-avatar-cropper)
- [md-editor-v3](https://github.com/imzbf/md-editor-v3)
- [vue-i18n](https://github.com/intlify/vue-i18n)
- [axios](https://github.com/axios/axios)