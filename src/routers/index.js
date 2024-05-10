import { TimelineItem } from '@opentiny/vue';
import { createRouter, createWebHistory } from 'vue-router';

const routes = [
    { 
        path: '/', 
        component: () => import('../pages/MainPage.vue'),
        meta: {
            title: 'FeasOJ - 首页'
        }
    }, 
    { 
        path: '/about', 
        component: () => import('../pages/AboutPage.vue'),
        meta: {
            title: 'FeasOJ - 关于'
        }
    },
    { 
        path: '/rank', 
        component: () => import('../pages/RankPage.vue'),
        meta: {
            title: 'FeasOJ - 排名'
        }
    },
    { 
        path: '/problemset', 
        component: () => import('../pages/ProblemSetPage.vue'),
        meta: {
            title: 'FeasOJ - 题目'
        }
    },
    { 
        path: '/contest', 
        component: () => import('../pages/ContestPage.vue'),
        meta: {
            title: 'FeasOJ - 竞赛'
        }
    },
    { 
        path: '/login', 
        component: () => import('../pages/LoginPage.vue'),
        meta: {
            title: 'FeasOJ - 登录'
        }
    },
    { 
        path: '/register', 
        component: () => import('../pages/RegisterPage.vue'),
        meta: {
            title: 'FeasOJ - 注册'
        }
    },
    { 
        path: '/profile', 
        component: () => import('../pages/ProfilePage.vue'),
        meta: {
            title: 'FeasOJ - 个人资料'
        }
    },
    { 
        path: '/problem/:id', 
        component: () => import('../pages/ProblemInfoPage.vue'),
        meta: {
            title: 'Problem :id'
        }
    },
    { 
        path: '/admin', 
        component: () => import('../pages/BackendManagementPage.vue'),
        meta: {
            title: 'FeasOJ - 后台管理'
        }
    },
    { 
        path: '/reset', component: () => import('../pages/PasswordResetPage.vue'),
        meta: {
            title: 'FeasOJ - 忘记密码'
        }
    },
];

const router = createRouter({
    history: createWebHistory(),
    routes, 
});

export default router;
