import { createRouter, createWebHistory } from 'vue-router';

const routes = [
    { path: '/', component: () => import('../components/MainPage.vue')}, 
    { path: '/about', component: () => import('../components/AboutPage.vue')},
    { path: '/rank', component: () => import('../components/Rank.vue')},
    { path: '/problemset', component: () => import('../components/ProblemSetPage.vue')},
    { path: '/contest', component: () => import('../components/Contest.vue')},
    { path: '/login', component: () => import('../components/LoginPage.vue')},
];

const router = createRouter({
    history: createWebHistory(),
    routes, 
});

export default router;
