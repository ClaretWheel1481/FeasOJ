import { createApp } from 'vue';
import App from './App.vue';
import { createRouter, createWebHashHistory } from 'vue-router';
import './style.css';

const MainPage = () => import('/src/components/MainPage.vue');
const AboutPage = () => import('/src/components/AboutPage.vue');
const TopPage = () => import('/src/components/Top.vue')
const PSPage = () => import('/src/components/ProblemSet.vue');
const Contest = () => import('/src/components/Contest.vue');

// 定义路由
const routes = [
    { path: '/', component: MainPage }, 
    { path: '/about', component: AboutPage},
    { path: '/top', component: TopPage},
    { path: '/problemset', component: PSPage},
    { path: '/contest', component: Contest}
];

// 创建路由实例
const router = createRouter({
    history: createWebHashHistory(),
    routes, 
});

// 创建并挂载Vue应用实例
const app = createApp(App);
app.use(router);
app.mount('#app');
