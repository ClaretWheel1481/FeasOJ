import { createRouter, createWebHistory } from 'vue-router';

const routes = [
    { path: '/', component: () => import('../pages/MainPage.vue')}, 
    { path: '/about', component: () => import('../pages/AboutPage.vue')},
    { path: '/rank', component: () => import('../pages/RankPage.vue')},
    { path: '/problemset', component: () => import('../pages/ProblemSetPage.vue')},
    { path: '/contest', component: () => import('../pages/ContestPage.vue')},
    { path: '/login', component: () => import('../pages/LoginPage.vue')},
    { path: '/register', component: () => import('../pages/RegisterPage.vue')},
    { path: '/profile', component: () => import('../pages/ProfilePage.vue')},
    { path: '/problem/:id', component: () => import('../pages/ProblemInfoPage.vue')},
    { path: '/admin', component: () => import('../pages/BackendManagementPage.vue')},
    { path: '/reset', component: () => import('../pages/PasswordResetPage.vue')},
];

const router = createRouter({
    history: createWebHistory(),
    routes, 
});

export default router;
