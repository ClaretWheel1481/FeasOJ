import { createApp } from 'vue';
import { createVuetify } from 'vuetify';
import App from './App.vue';
import 'vuetify/styles';
import router from './routers/index';
import './style.css';
import * as components from 'vuetify/components';
import * as directives from 'vuetify/directives';

const vuetify = createVuetify({
    components,
    directives,
    theme: {
        themes: {
            light: {
                primary: '#42A5F5',
                secondary: '#42A5F5',
            },
        },
    }
})

router.beforeEach((to, from, next) => {
    /* 路由发生变化修改页面title */
    if (to.meta.title) {
        document.title = to.meta.title
    }
    next()
})

const app = createApp(App);
app.use(router).use(vuetify).mount('#app');
