<script setup lang="ts">
import zhCN from 'ant-design-vue/es/locale/zh_CN';
import dayjs from 'dayjs';
import 'dayjs/locale/zh-cn';
dayjs.locale('zh-cn');

const locale = zhCN

</script>

<template>
  <a-config-provider :locale="locale">
    <a-app>
      <router-view />
    </a-app>
  </a-config-provider>
</template>
