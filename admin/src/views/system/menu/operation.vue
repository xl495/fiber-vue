<template>
    <div class="menu-operation">
        <a-form :model="formState" :label-col="labelCol" :wrapper-col="wrapperCol">
            <a-form-item label="name">
                <a-input v-model:value="formState.name" />
            </a-form-item>
            <a-form-item label="path">
                <a-input v-model:value="formState.path" />
            </a-form-item>
            <a-form-item label="sort">
                <a-input v-model:value="formState.sort" />
            </a-form-item>
            <a-form-item label="component">
                <a-input v-model:value="formState.component" />
            </a-form-item>
            <a-form-item label="是否隐藏">
                <a-switch v-model:checked="formState.mate.hidden" />
            </a-form-item>
            <a-form-item label="是否缓存">
                <a-switch v-model:checked="formState.mate.keepAlive" />
            </a-form-item>
            <a-form-item label="图标">
                <a-input v-model:value="formState.mate.icon" />
            </a-form-item>
            <a-form-item label="显示名称">
                <a-input v-model:value="formState.mate.title" />
            </a-form-item>
        </a-form>
    </div>
</template>

<script setup lang="ts">
import { reactive } from 'vue';
import type { UnwrapRef } from 'vue';

const formState: UnwrapRef<MenuItem> = reactive({
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
const labelCol = { style: { width: '80px' } };
const wrapperCol = { span: 14 };

</script>

<style lang="scss">
.menu-operation {
    padding-top: 20px;
}
</style>

