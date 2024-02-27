<template>
    <ul v-if="props.contextMenuVisible" :style="{ left: props.left + 'px', top: props.top + 'px' }" class="contextmenu">
        <li @click="closeAll">关闭所有</li>
        <li @click="closeLeft">关闭左侧</li>
        <li @click="closeRight">关闭右侧</li>
        <li @click="closeOther">关闭其他</li>
    </ul>
</template>
  
<script setup lang="ts">
import { useMultipleTabStore } from '@/store/modules/multiple-tab'
import { onUnmounted } from 'vue';

const props = withDefaults(
    defineProps<{
        contextMenuVisible: boolean
        left: number
        top: number
    }>(),
    {
        contextMenuVisible: false,
        left: 0,
        top: 0,
    },
)

const multipleTabStore = useMultipleTabStore()

const emits = defineEmits(['update:context-menu-visible'])

const closeAll = () => {
    multipleTabStore.removeAllTabs()
}
const closeLeft = () => {
    multipleTabStore.removeLeftTabs()
}
const closeRight = () => {
    multipleTabStore.removeRightTabs()
}
const closeOther = () => {
    multipleTabStore.removeOtherTabs()
}

document.body.addEventListener('click', () => {
    if (!props.contextMenuVisible) return
    emits('update:context-menu-visible', false)
})

onUnmounted(() => {
    document.body.removeEventListener('click', () => {
        emits('update:context-menu-visible', false)
    })
})
</script>
  
<style scoped></style>