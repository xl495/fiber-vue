import { createApp } from 'vue'
import { message, notification } from 'ant-design-vue';
import './style.css'
import 'ant-design-vue/dist/reset.css';

import Antd from 'ant-design-vue';

import App from './App.vue'

import route from './router';

import store from './store';

const app = createApp(App);

app.config.globalProperties.$message = message;

app.config.globalProperties.$notification = notification;

app.use(Antd).use(store).use(route).mount('#app')
