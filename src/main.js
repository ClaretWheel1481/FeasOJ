import { createApp } from 'vue';
import App from './App.vue';
import router from './router/index';
import './assets/style.css';
import { registerPlugins } from './plugins';
import { i18n } from './plugins/vue-i18n';
import { nextTick } from 'vue';

const settingLanguage = localStorage.getItem('language');
if (settingLanguage) {
  i18n.global.locale.value = settingLanguage;
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
