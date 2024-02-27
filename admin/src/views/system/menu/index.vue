<template>
    <div class="menu-wrap">
        <div class="menu-header">
            <a-button type="primary" @click="addMenu">新增</a-button>
        </div>

        <a-tree :tree-data="treeData">
            <template #title="{ title, key }">
                <span v-if="key === '0-0-1-0'" style="color: #1890ff">{{ title }}</span>
                <template v-else>{{ title }}</template>
            </template>
        </a-tree>


        <a-modal v-model:open="isOpen" :title="modelType ? '新增菜单' : '编辑菜单'" @ok="handleOk">
            <operation :model="model" />
        </a-modal>
    </div>
</template>

<script setup lang="ts">
import { Ref, reactive, ref } from 'vue';
import { type TreeProps } from 'ant-design-vue';
import operation from './operation.vue';

const isOpen = ref(false);

const modelType: Ref<'add' | 'edit'> = ref('add');

const model = reactive<MenuItem>({
    name: '',
    path: '',
    parentId: 0,
    sort: 0,
    component: '',
    mate: {
        hidden: false,
        keepAlive: false,
        icon: '',
        title: ''
    }
});

const treeData: TreeProps['treeData'] = [
    {
        title: 'parent 1',
        key: '0-0',
        children: [
            {
                title: 'parent 1-0',
                key: '0-0-0',
                disabled: true,
                children: [
                    { title: 'leaf', key: '0-0-0-0', disableCheckbox: true },
                    { title: 'leaf', key: '0-0-0-1' },
                ],
            },
            {
                title: 'parent 1-1',
                key: '0-0-1',
                children: [{ key: '0-0-1-0', title: 'sss' }],
            },
        ],
    },
];

const addMenu = () => {
    isOpen.value = true;
};

const handleOk = () => {
    console.log('handleOk');
};

</script>

<style lang="scss">
.menu-wrap {
    padding: 20px;

    .menu-header {
        margin-bottom: 20px;
    }
}
</style>