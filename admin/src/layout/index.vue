<template>
    <a-layout :style="{ height: '100vh' }">
        <a-layout-sider width="230" :style="siderStyle" v-model:collapsed="userStore.isCollapseMenu" collapsible>
            <Menu />
        </a-layout-sider>
        <a-layout>
            <a-layout-header :style="headerStyle">
            </a-layout-header>
            <a-layout-content :style="contentStyle">
            
                <div v-for="item in 1000" :key="item" >{{  item  }}</div>

            </a-layout-content>
            <a-layout-footer :style="footerStyle">Footer</a-layout-footer>
        </a-layout>
    </a-layout>
</template>

<script setup lang="ts">
import type { CSSProperties } from 'vue';
import Menu from './menu/index.vue'
import { useUserStore } from '@/store/modules/user';
import { MenuUnfoldOutlined, MenuFoldOutlined } from '@ant-design/icons-vue';

const userStore = useUserStore();



const headerStyle: CSSProperties = {
    textAlign: 'center',
    color: '#fff',
    height: 64,
    paddingInline: 50,
    lineHeight: '64px',
    backgroundColor: '#7dbcea',
};

const siderStyle: CSSProperties = {
    textAlign: 'center',
    lineHeight: '120px',
    color: '#fff',
    backgroundColor: '#3ba0e9',
};

const contentStyle: CSSProperties = {
    textAlign: 'center',
    minHeight: 120,
    lineHeight: '120px',
    color: '#fff',
    backgroundColor: '#108ee9',
    overflow: 'auto',
};


const footerStyle: CSSProperties = {
    textAlign: 'center',
    color: '#fff',
    backgroundColor: '#7dbcea',
};
</script>

<style scoped></style>