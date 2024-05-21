<!-- TODO:后台管理页面，限制仅管理员显示 -->
<script setup>
import { TreeMenu as TinyTreeMenu } from '@opentiny/vue'
import { ref } from 'vue' 

const adminData = ref([
  {
    id: 100,
    label: '用户管理',
  },
  {
    id: 200,
    label: '题目管理',
  },
  {
    id: 300,
    label: '竞赛管理',
  },
  {
    id: 400,
    label: '更新日志'
  },
  {
    id: 500,
    label: '返回主页',
    url: '/'
  }
//   TODO:更多后台管理功能
]);
</script>

<template>
    <tiny-tree-menu :data="adminData"></tiny-tree-menu>
</template>