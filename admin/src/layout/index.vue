<template>
    <a-layout :style="{ height: '100vh' }">
        <a-layout-sider width="230" :style="siderStyle" v-model:collapsed="userStore.isCollapseMenu" collapsible>
            <Menu />
        </a-layout-sider>
        <a-layout>
            <a-layout-header :style="headerStyle">
                <multiple-tab />
            </a-layout-header>
            <a-layout-content :style="contentStyle">
                <router-view />
            </a-layout-content>
            <a-layout-footer :style="footerStyle">Footer</a-layout-footer>
        </a-layout>
    </a-layout>
</template>

<script setup lang="ts">
import type { CSSProperties } from 'vue';
import Menu from './menu/index.vue'
import { useUserStore } from '@/store/modules/user';
import multipleTab from './multiple-tab/tabs.vue'

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
    minHeight: 120,
    overflow: 'auto',
};


const footerStyle: CSSProperties = {
    textAlign: 'center',
    color: '#fff',
    backgroundColor: '#7dbcea',
};
</script>

<style scoped></style>