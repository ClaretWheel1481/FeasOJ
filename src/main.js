import { createApp } from 'vue';
import App from './App.vue';
import router from './router/index';
import './assets/style.css';
import { registerPlugins } from './plugins';
import { i18n } from './plugins/vue-i18n';
import { nextTick } from 'vue';

// 先检查 localStorage 是否已保存语言设置
const settingLanguage = localStorage.getItem('language');
if (settingLanguage) {
    i18n.global.locale.value = settingLanguage;
} else {
    // 获取浏览器语言，如 "en-US"、"zh-CN"
    let browserLang = navigator.language || navigator.userLanguage;

    // 格式统一
    browserLang = browserLang.replace('-', '_');

    // 定义你支持的语言列表
    const supportedLanguages = ['en', 'es', 'fr', 'it', 'ja', 'zh_CN', 'zh_TW'];

    // 如果获取的语言不在支持列表中，则尝试匹配其前两位
    if (!supportedLanguages.includes(browserLang)) {
        const langPrefix = browserLang.split('_')[0];
        const match = supportedLanguages.find(lang => lang.startsWith(langPrefix));
        browserLang = match || 'en';
    }

    // 设置国际化语言并保存到 localStorage
    i18n.global.locale.value = browserLang;
    localStorage.setItem('language', browserLang);
}

// 避免用户修改 localStorage 中的数据
window.addEventListener('storage', function (e) {
    localStorage.setItem(e.key, e.oldValue)
});

router.beforeEach((to, from, next) => {
    // 路由发生变化修改页面title
    nextTick(() => {
        const titleKey = to.meta.titleKey;
        const title = titleKey ? i18n.global.t(titleKey) : to.meta.titleKey;
        document.title = title || 'FeasOJ';
    });
    next();
});

const app = createApp(App);

registerPlugins(app);

app.mount('#app');
