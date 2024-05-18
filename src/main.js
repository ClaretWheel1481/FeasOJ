import { createApp } from 'vue';
import App from './App.vue';
import router from './routers/index';
import './style.css';
import { createPinia } from 'pinia'

const pinia = createPinia()

router.beforeEach((to, from, next) => {
    /* 路由发生变化修改页面title */
    if (to.meta.title) {
        document.title = to.meta.title
    }
    next()
})

const app = createApp(App);
app.use(router).use(pinia).mount('#app');
