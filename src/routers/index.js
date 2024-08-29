import { createRouter, createWebHashHistory } from 'vue-router';

const routes = [
    {
        path: '/',
        component: () => import('../pages/Main.vue'),
        meta: {
            titleKey: 'message.mainpage'
        }
    },
    {
        path: '/about',
        component: () => import('../pages/About.vue'),
        meta: {
            titleKey: 'message.about'
        }
    },
    {
        path: '/problemset',
        component: () => import('../pages/Problem/Main.vue'),
        meta: {
            titleKey: 'message.problemset'
        }
    },
    {
        path: '/login',
        component: () => import('../pages/Auth/Login.vue'),
        meta: {
            titleKey: 'message.login'
        }
    },
    {
        path: '/register',
        component: () => import('../pages/Auth/Register.vue'),
        meta: {
            titleKey: 'message.register'
        }
    },
    {
        path: '/profile/:Username',
        component: () => import('../pages/Profile.vue'),
        meta: {
            titleKey: 'message.profile'
        },
        beforeEnter: (to, from, next) => {
            to.meta.title = to.params.Username;
            next();
        }
    },
    {
        path: '/problem/:Pid',
        component: () => import('../pages/Problem/Details.vue'),
        meta: {
            title: 'Problem'
        },
        beforeEnter: (to, from, next) => {
            to.meta.title = 'Problem ' + to.params.Pid;
            next();
        }
    },
    {
        path: '/reset', component: () => import('../pages/Auth/Reset.vue'),
        meta: {
            titleKey: 'message.resetpwd'
        }
    },
    {
        path: '/status', component: () => import('../pages/Status.vue'),
        meta: {
            titleKey: 'message.status'
        }
    },
    {
        path: '/discussion', component: () => import('../pages/Discuss/Main.vue'),
        meta: {
            titleKey: 'message.discussion'
        }
    },
    {
        path: '/discussion/create', component: () => import('../pages/Discuss/New.vue'),
        meta: {
            titleKey: 'message.createDiscussion'
        }
    },
    {
        path: '/discussion/:Did',
        component: () => import('../pages/Discuss/Details.vue'),
        meta: {
            titleKey: 'message.discussion'
        },
    },
    {
        path: '/psm', component: () => import('../pages/Admin/Problemset.vue'),
        meta: {
            titleKey: 'message.problemmanagement'
        }
    },
    {
        path: '/am', component: () => import('../pages/Admin/Account.vue'),
        meta: {
            titleKey: 'message.usermanagement'
        }
    },
    {
        path: '/403', component: () => import('../pages/403.vue'),
        meta: {
            titleKey: 'message.nopermission'
        }
    }
];

const router = createRouter({
    history: createWebHashHistory(),
    routes,
});

export default router;
