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
            <operation ref="operationRef" :model="model" />
            <template #footer>
                <a-button key="back" @click="resetModel()" >返回</a-button>
                <a-button key="submit" type="primary" :loading="isSubmitLoading" @click="handleOk">提交</a-button>
            </template>
        </a-modal>
    </div>
</template>

<script setup lang="ts">
import { Ref, reactive, ref } from 'vue';
import { notification, type TreeProps } from 'ant-design-vue';
import operation from './operation.vue';
import { addMenu as addMenuApi, getMenus } from '@/api/system/menu';

const isOpen = ref(false);

const modelType: Ref<'add' | 'edit'> = ref('add');

const operationRef = ref<any>();

const isSubmitLoading = ref(false);

const model = reactive<MenuItem>({
    ID: null,
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
    operationRef.value.formRef && operationRef.value.formRef.validate().then(() => {
        isSubmitLoading.value = true;
        addMenuApi(operationRef.value.formState).then(() => {
            notification.success({
                message: `${modelType.value === 'add' ? '新增' : '编辑'}菜单成功`,
                description: `${modelType.value === 'add' ? '新增' : '编辑'}菜单成功`,
            });
            fetchMenu();
            resetModel()
            isOpen.value = false;
        }).finally(() => {
            isSubmitLoading.value = false;
        });
    })
        .catch((error: any) => {
            console.log('error', error);
        });

};

const resetModel = () => {
    modelType.value = 'add';
    model.ID = null;
    model.name = '';
    model.path = '';
    model.parentId = 0;
    model.sort = 0;
    model.component = '';
    model.mate.hidden = false;
    model.mate.keepAlive = false;
    model.mate.icon = '';
    model.mate.title = '';
};

const fetchMenu = () => {
    getMenus().then((res: any) => {
        console.log('res', res);
    });
};

fetchMenu()

</script>

<style lang="scss">
.menu-wrap {
    padding: 20px;

    .menu-header {
        margin-bottom: 20px;
    }
}
</style>