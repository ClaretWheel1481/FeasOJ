import { createRouter, createWebHistory } from 'vue-router';

const MainPage = () => import('../components/MainPage.vue');
const AboutPage = () => import('../components/AboutPage.vue');
const TopPage = () => import('../components/Top.vue');
const PSPage = () => import('../components/ProblemSet.vue');
const Contest = () => import('../components/Contest.vue');

const routes = [
    { path: '/', component: MainPage }, 
    { path: '/about', component: AboutPage},
    { path: '/top', component: TopPage},
    { path: '/problemset', component: PSPage},
    { path: '/contest', component: Contest}
];

const router = createRouter({
    history: createWebHistory(),
    routes, 
});

export default router;
