import { createApp } from 'vue';
import App from './App.vue';
import 'vuetify/styles';
import router from './routers/index';
import './assets/style.css';
import { registerPlugins } from './plugins';

// FIXME:Firefox无效
window.addEventListener('storage', function (e) {
    localStorage.setItem(e.key, e.oldValue)
});

router.beforeEach((to, from, next) => {
    /* 路由发生变化修改页面title */
    if (to.meta.title) {
        document.title = to.meta.title
    }
    next()
})

const app = createApp(App);

registerPlugins(app);

app.mount('#app');
