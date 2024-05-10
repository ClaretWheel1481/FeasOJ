<!-- TODO:个人资料页，包括用户名、头像、邮箱、简介、Score、答题记录、修改密码、退出登录 -->
<!-- 根据用户UID来显示所有内容 -->
<script setup>
import { UserHead as TinyUserHead } from '@opentiny/vue'
</script>

<template>
  <div>
    <h1>个人资料</h1>
  </div>
  <tiny-user-head type="icon" round></tiny-user-head>
</template>