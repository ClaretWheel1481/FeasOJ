import { createRouter, createWebHashHistory } from 'vue-router';
import { nextTick } from 'vue';

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
            title: 'FeasOJ - 关于',
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
        path: '/profile/:Username',
        component: () => import('../pages/ProfilePage.vue'),
        meta: {
            title: 'FeasOJ - 个人资料'
        },
        beforeEnter: (to, from, next) => {
            to.meta.title = to.params.Username + '的资料';
            next();
        }
    },
    {
        path: '/problem/:Pid',
        component: () => import('../pages/ProblemInfoPage.vue'),
        meta: {
            title: 'Problem'
        },
        beforeEnter: (to, from, next) => {
            to.meta.title = 'Problem ' + to.params.Pid;
            next();
        }
    },
    {
        path: '/reset', component: () => import('../pages/PasswordResetPage.vue'),
        meta: {
            title: 'FeasOJ - 重置密码'
        }
    },
    {
        path: '/status', component: () => import('../pages/StatusPage.vue'),
        meta: {
            title: 'FeasOJ - 状态'
        }
    },
    {
        path: '/discussion', component: () => import('../pages/DiscussPage.vue'),
        meta: {
            title: 'FeasOJ - 讨论'
        }
    },
    {
        path: '/discussion/create', component: () => import('../pages/NewDiscussionPage.vue'),
        meta: {
            title: 'FeasOJ - 创建讨论'
        }
    },
    {
        path: '/discussion/:Did',
        component: () => import('../pages/DiscussionDetailsPage.vue'),
        meta: {
            title: 'FeasOJ - 讨论'
        },
    },
    {
        path: '/psm', component: () => import('../pages/ProblemsetManagementPage.vue'),
        meta: {
            title: 'FeasOJ - 题目管理'
        }
    },
    {
        path: '/am', component: () => import('../pages/AccountManagementPage.vue'),
        meta: {
            title: 'FeasOJ - 用户管理'
        }
    },
    {
        path: '/403', component: () => import('../pages/NoPermissionPage.vue'),
        meta: {
            title: 'FeasOJ - 没有权限',
        }
    }
];

const router = createRouter({
    history: createWebHashHistory(),
    routes,
});

router.afterEach((to) => {
    nextTick(() => {
        document.title = to.meta.title || 'FeasOJ';
    });
});

export default router;
