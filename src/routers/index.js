import { createRouter, createWebHistory } from 'vue-router';

const MainPage = () => import('../components/MainPage.vue');
const AboutPage = () => import('../components/AboutPage.vue');
const RankPage = () => import('../components/Rank.vue');
const PSPage = () => import('../components/ProblemSet.vue');
const Contest = () => import('../components/Contest.vue');

const routes = [
    { path: '/', component: MainPage }, 
    { path: '/about', component: AboutPage},
    { path: '/rank', component: RankPage},
    { path: '/problemset', component: PSPage},
    { path: '/contest', component: Contest}
];

const router = createRouter({
    history: createWebHistory(),
    routes, 
});

export default router;
