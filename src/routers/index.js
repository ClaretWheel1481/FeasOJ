import { createRouter, createWebHistory } from 'vue-router';
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
            transition: 'slide-left'
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
        path: '/admin', 
        component: () => import('../pages/BackendManagementPage.vue'),
        meta: {
            title: 'FeasOJ - 后台管理',
            hideComponent: true
        },
        children: [
            {
                path: 'accounts',
                component: () => import('../pages/AccountManagementPage.vue'),
                meta: {
                    title: 'FeasOJ - 用户管理',
                }
            }
        ]
    },
    { 
        path: '/reset', component: () => import('../pages/PasswordResetPage.vue'),
        meta: {
            title: 'FeasOJ - 忘记密码'
        }
    },
    {
        path: '/status', component: () => import('../pages/StatusPage.vue'),
        meta: {
            title: 'FeasOJ - 状态'
        }
    },
];

const router = createRouter({
    history: createWebHistory(),
    routes, 
});

router.afterEach((to) => {
    nextTick(() => {
      document.title = to.meta.title || 'FeasOJ';
    });
  });
  
export default router;
