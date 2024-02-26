import { createRouter, createWebHistory } from 'vue-router'


const routes = [
    {
        path: '/',
        name: 'Layout',
        component: () => import('@/layout/index.vue'),
    },
    {
        path: '/login',
        name: 'Login',
        component: () => import('@/views/login/index.vue'),
    }
]

const route = createRouter({
    history: createWebHistory(),
    routes
})


route.beforeEach((to, from, next) => {
    const token = localStorage.getItem('token')
    // 获取 token 不存在直接跳转登录页
    if (to.path !== '/login' && !token) {
        return next('/login')
    }
    next()
})

export default route