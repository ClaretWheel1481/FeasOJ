import { createI18n } from "vue-i18n";

const messages = {
    en: {
      message: {
        mainpage: 'Home',
        problem: 'ProblemSet',
        status: 'Status',
        discussion: 'Discussion',
        about: 'About',
        usermanagement: 'User Management',
        problemmanagement: 'Problem Management',
        lang: 'Language',
        login: 'Login',
        register: 'Register',
        forget: 'Forget Password?',
        username: 'Username',
        password: 'Password',
      },
    },
    zh_CN: {
      message: {
        mainpage: '主页',
        problem: '题目',
        status: '状态',
        discussion: '讨论',
        about: '关于',
        usermanagement: '用户管理',
        problemmanagement: '题目管理',
        lang: '语言',
        login: '登录',
        register: '注册',
        forget: '忘记密码？',
        username: '用户名',
        password: '密码',
      },
    },
    ja: {
      message: {
        mainpage: 'ホームページ',
        problem: '問題セット',
        status: 'ステータス',
        discussion: 'ディスカッション',
        about: 'アバウト',
        usermanagement: 'ユーザー管理',
        problemmanagement: '問題管理',
        lang: '言語',
        login: 'ログイン',
        register: '登録ユーザー',
        forget: '忘れたのか？',
        username: 'ユーザー名',
        password: 'パスワード',
      },
    },
    zh_TW: {
      message: {
        mainpage: '主頁',
        problem: '題目',
        status: '狀態',
        discussion: '討論',
        about: '關於',
        usermanagement: '用戶管理',
        problemmanagement: '題目管理',
        lang: '語言',
        login: '登錄',
        register: '註冊',
        forget: '忘記密碼？',
        username: '用戶名',
        password: '密碼',
      },
    },
};

export const i18n = createI18n({
    legacy: false,
    locale: 'zh_CN', 
    fallbackLocale: 'en',
    messages,
})