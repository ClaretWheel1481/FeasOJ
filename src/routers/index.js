import { createRouter, createWebHashHistory } from 'vue-router';

const routes = [
    {
        path: '/',
        component: () => import('../pages/MainPage.vue'),
        meta: {
            titleKey: 'message.mainpage'
        }
    },
    {
        path: '/about',
        component: () => import('../pages/AboutPage.vue'),
        meta: {
            titleKey: 'message.about'
        }
    },
    {
        path: '/problemset',
        component: () => import('../pages/ProblemSetPage.vue'),
        meta: {
            titleKey: 'message.problemset'
        }
    },
    {
        path: '/login',
        component: () => import('../pages/LoginPage.vue'),
        meta: {
            titleKey: 'message.login'
        }
    },
    {
        path: '/register',
        component: () => import('../pages/RegisterPage.vue'),
        meta: {
            titleKey: 'message.register'
        }
    },
    {
        path: '/profile/:Username',
        component: () => import('../pages/ProfilePage.vue'),
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
            titleKey: 'message.resetpwd'
        }
    },
    {
        path: '/status', component: () => import('../pages/StatusPage.vue'),
        meta: {
            titleKey: 'message.status'
        }
    },
    {
        path: '/discussion', component: () => import('../pages/DiscussPage.vue'),
        meta: {
            titleKey: 'message.discussion'
        }
    },
    {
        path: '/discussion/create', component: () => import('../pages/NewDiscussionPage.vue'),
        meta: {
            titleKey: 'message.createDiscussion'
        }
    },
    {
        path: '/discussion/:Did',
        component: () => import('../pages/DiscussionDetailsPage.vue'),
        meta: {
            titleKey: 'message.discussion'
        },
    },
    {
        path: '/psm', component: () => import('../pages/ProblemsetManagementPage.vue'),
        meta: {
            titleKey: 'message.problemmanagement'
        }
    },
    {
        path: '/am', component: () => import('../pages/AccountManagementPage.vue'),
        meta: {
            titleKey: 'message.usermanagement'
        }
    },
    {
        path: '/403', component: () => import('../pages/NoPermissionPage.vue'),
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
