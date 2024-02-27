<template>
    <div class="router-history">

        <a-tabs v-model:activeKey="multipleTabStore.currentKey" hide-add type="editable-card" @edit="editTab"
            @tab-click="changeTab" @contextmenu.prevent="openContextMenu($event)">
            <a-tab-pane v-for="(v, _) in multipleTabStore.getTabList" :key="v.KEY" :tab="v.name" :closable="true">
                <!-- {{ v.name }} -->
            </a-tab-pane>
        </a-tabs>

        <!-- <context-menu v-model:context-menu-visible="contextMenuOptions.contextMenuVisible" :left="contextMenuOptions.left"
            :top="contextMenuOptions.top" /> -->
    </div>
</template>
<script setup lang="ts">
import { useMultipleTabStore } from '@/store/modules/multiple-tab'
import contextMenu from './contextmenu.vue'
import { useRoute } from 'vue-router'
import { ref, watch } from 'vue'

defineOptions({
    name: 'multipleTab',
})

const route = useRoute()

const multipleTabStore = useMultipleTabStore()

const isCollapse = ref(false)

const isMobile = ref(false)

const contextMenuOptions = ref({
    contextMenuVisible: false,
    left: 0,
    top: 0,
    currentTab: '',
})

const editTab = (targetKey: string, action: string) => {
    console.log('targetKey', targetKey, 'action', action);
    switch (action) {
        case 'remove':
            removeTab(targetKey)
            break
        default:
            break
    }
}


const changeTab = (tabName: any) => {
    console.log(tabName);

    multipleTabStore.changeTab(tabName)
}

const removeTab = (tab: string) => {
    multipleTabStore.removeTab(tab)
}

const openContextMenu = (e: MouseEvent) => {
    const targetElement = e?.target as HTMLElement;
    if (targetElement.tagName.toLowerCase() !== "span") return
    const tabName = targetElement.getAttribute('tab')
    multipleTabStore.setContextMenuKey(tabName as string)

    e.preventDefault()

    let width: number
    if (isCollapse.value) {
        width = 54
    } else {
        width = 220
    }
    if (isMobile.value) width = 0

    // 获取当前点击的位置
    contextMenuOptions.value.left = e.clientX - width
    contextMenuOptions.value.top = e.clientY + 10
    contextMenuOptions.value.contextMenuVisible = true
}

watch(
    route,
    (to, _) => {
        if (to.name === 'Login' || to.name === 'Reload') return
        multipleTabStore.setTab(to)
    },
    {
        immediate: true,
        deep: true,
    },
)
</script>
  
<style lang="scss">
.contextmenu {
    width: 100px;
    margin: 0;
    border: 1px solid #ccc;
    background: #fff;
    z-index: 3000;
    position: absolute;
    list-style-type: none;
    padding: 5px 0;
    border-radius: 4px;
    font-size: 14px;
    color: #333;
    box-shadow: 2px 2px 3px 0 rgba(0, 0, 0, 0.2);
}

.el-tabs__item .el-icon-close {
    color: initial !important;
}

.el-tabs__item .dot {
    content: '';
    width: 9px;
    height: 9px;
    margin-right: 8px;
    display: inline-block;
    border-radius: 50%;
    transition: background-color 0.2s;
}

.contextmenu li {
    margin: 0;
    padding: 7px 16px;
}

.contextmenu li:hover {
    background: #f2f2f2;
    cursor: pointer;
}
</style>