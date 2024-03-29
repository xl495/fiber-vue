import { useUserStore } from '@/store/modules/user'
import { createRouter, createWebHistory } from 'vue-router'


const routes = [
    {
        path: '/',
        name: 'Layout',
        component: () => import('@/layout/index.vue'),
        children: [
            {
                path: 'system',
                name: 'System',
                children: [
                    {
                        path: 'menu',
                        name: 'Menu',
                        component: () => import('@/views/system/menu/index.vue'),
                    },
                ]
            }
        ]
    },
    {
        path: '/login',
        name: 'Login',
        component: () => import('@/views/login/index.vue'),
    },

]

const route = createRouter({
    history: createWebHistory(),
    routes
})


route.beforeEach((to, from, next) => {
    const token = localStorage.getItem('token')
    const userStore = useUserStore()
    // 获取 token 不存在直接跳转登录页
    if (to.path !== '/login' && !token) {
        return next('/login')
    }

    if (token && to.path !== '/login') {
        userStore.fetchUser()
    }
    next()
})

export default route