<template>
  <a-form :model="formState" ref="formRef" :rules="rules" layout="vertical" :hide-required-mark="true">
    <a-form-item label="用户名" name="username">
      <a-input v-model:value="formState.username" placeholder="Username" />
    </a-form-item>
    <a-form-item label="密码" name="password">
      <a-input v-model:value="formState.password" type="password" placeholder="Password" />
    </a-form-item>
    <a-form-item>
      <a-button type="primary" @click="handleSubmit" :loading="formState.isSubmitting">Login</a-button>
    </a-form-item>
  </a-form>
</template>
  
<script setup lang="ts">
import { ref, reactive } from 'vue';
import type { Rule } from 'ant-design-vue/es/form';
import { useUserStore } from '@/store/modules/user';
import route from '@/router';
import { notification } from 'ant-design-vue';

const userStore = useUserStore();

const formRef = ref();

const formState = reactive({
  username: '',
  password: '',
  isSubmitting: false,
});

const rules: Record<string, Rule[]> = {
  username: [{ required: true, message: 'Please input your username!' }],
  password: [{ required: true, message: 'Please input your password!' }],
};

const handleSubmit = () => {
  formRef.value
    .validate()
    .then(() => {
      formState.isSubmitting = true;
      
      userStore.login(formState)
        .then((res) => {
          console.log('res',res);
          
          if (res.user) {
            console.log('login success');
            route.push('/');
            notification.success({
              message: '登录成功',
              description: '欢迎回来',
            });
          }
          
          formState.isSubmitting = false;
        })
        .catch((err) => {
          console.log(err);
          formState.isSubmitting = false;
        });
    })
    .catch((error : any) => {
      console.log('error', error);
    });
};

</script>
  
<style scoped>
/* add your custom styles here */
</style>
  