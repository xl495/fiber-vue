<template>
    <div class="menu-operation">
        <a-form ref="formRef" :model="formState" :label-col="labelCol" :wrapper-col="wrapperCol" :rules="rules">
            <a-form-item label="name" name="name">
                <a-input v-model:value="formState.name" />
            </a-form-item>
            <a-form-item label="path" name="path">
                <a-input v-model:value="formState.path" />
            </a-form-item>
            <a-form-item label="sort" name="sort">
                <a-input v-model:value="formState.sort" type="number" />
            </a-form-item>
            <a-form-item label="component" name="component">
                <a-input v-model:value="formState.component" />
            </a-form-item>
            <a-form-item label="是否隐藏">
                <a-switch v-model:checked="formState.mate.hidden" />
            </a-form-item>
            <a-form-item label="是否缓存">
                <a-switch v-model:checked="formState.mate.keepAlive" />
            </a-form-item>
            <a-form-item label="图标" :name="['mate', 'icon']"
                :rules="{ required: true, message: 'Please input your icon', trigger: 'blur' }">
                <a-input v-model:value="formState.mate.icon" />
            </a-form-item>
            <a-form-item label="显示名称" :name="['mate', 'title']"
                :rules="{ required: true, message: 'Please input your title', trigger: 'blur' }">
                <a-input v-model:value="formState.mate.title" />
            </a-form-item>
        </a-form>
    </div>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue';
import type { UnwrapRef } from 'vue';

const formRef = ref();

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
const labelCol = { style: { width: '100px' } };
const wrapperCol = { span: 14 };

const rules = reactive({
    name: [
        { required: true, message: 'Please input your name', trigger: 'blur' },
    ],
    path: [
        { required: true, message: 'Please input your path', trigger: 'blur' },
    ],
    sort: [
        { required: true, message: 'Please input your sort', trigger: 'blur' },
    ],
    component: [
        { required: true, message: 'Please input your component', trigger: 'blur' },
    ],

});

defineExpose({
    formRef: formRef,
    formState: formState,
});

</script>

<style lang="scss">
.menu-operation {
    padding-top: 20px;
}
</style>