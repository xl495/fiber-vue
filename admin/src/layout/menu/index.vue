<template>
    <a-menu v-model:openKeys="state.openKeys" v-model:selectedKeys="state.selectedKeys" mode="inline"
        style="min-height: 100vh;" theme="dark" :inline-collapsed="false" :items="items"></a-menu>
</template>

<script setup lang="ts">
import { useUserStore } from '@/store/modules/user';
import { UserOutlined, VideoCameraOutlined, UploadOutlined } from '@ant-design/icons-vue';
import { reactive, ref, h } from 'vue';
import { useRoute } from 'vue-router';
// import { RouteRecordNormalized } from 'vue-router';

const userStore = useUserStore();

const route = useRoute()

interface RouteItem {
    path: string;
    name: string;
    icon?: any;
    component?: string;
    children?: RouteItem[];
    meta?: Record<string, any>;
}


const items = reactive([
    {
        key: '/',
        icon: h(UserOutlined),
        label: '控制台',
        title: '控制台',
    },
    {
        key: 'sub1',
        label: 'Navigation One',
        title: 'Navigation One',
        icon: h(UserOutlined),
        children: [
            {
                key: '5',
                label: 'Option 5',
                title: 'Option 5',
                children: [
                    {
                        key: '8',
                        label: 'Option 5',
                        title: 'Option 5',
                    },
                    {
                        key: '10',
                        label: 'Option 6',
                        title: 'Option 6',
                    },
                ],
            },
            {
                key: '6',
                label: 'Option 6',
                title: 'Option 6',
            },
        ],
    }
])

const state = reactive({
    collapsed: false,
    selectedKeys: [route.fullPath],
    openKeys: [],
    preOpenKeys: [],
});


const routes: RouteItem[] = [
    {
        path: '/login',
        name: 'login',
        meta: {
            icon: 'UserOutlined',
        },
        component: '/views/nav1/index.vue',
    },
    {
        path: '/',
        name: 'home',
        children: [
            {
                path: '/nav1',
                name: 'nav1',
                icon: UserOutlined,
                component: '/views/nav1/index.vue',
            },
            {
                path: '/nav2',
                name: 'nav2',
                icon: VideoCameraOutlined,
                component: '/views/nav2/index.vue',
            },
            {
                path: '/nav3',
                name: 'nav3',
                icon: UploadOutlined,
                component: '/views/nav3/index.vue',
            },
            {
                path: '/nav4',
                name: 'nav4',
                icon: UserOutlined,
                component: '/views/nav4/index.vue',
            },
        ],
    }
]

const selectedKeys = ref<string[]>(['4']);
</script>

<style scoped></style>