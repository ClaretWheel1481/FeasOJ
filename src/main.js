import { createApp } from 'vue';
import { createVuetify } from 'vuetify';
import App from './App.vue';
import 'vuetify/styles';
import router from './routers/index';
import './style.css';
import * as components from 'vuetify/components';
import * as directives from 'vuetify/directives';
import hljs from 'highlight.js';
import CodeEditor from 'simple-code-editor';

const vuetify = createVuetify({
    components,
    directives,
    theme: {
      themes: {
        'light': { // 浅色主题
            variables:{},
            primary: '#1976D2',
            secondary: '#424242',
            accent: '#82B1FF',
            error: '#FF5252',
            info: '#2196F3',
            success: '#4CAF50',
            warning: '#FFC107'
        },
      }
    }
})

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
app.use(router).use(vuetify).mount('#app');
