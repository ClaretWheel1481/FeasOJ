<!-- TODO:个人资料页，包括用户名、头像、邮箱、简介、注册时间、提交记录记录、修改密码、退出登录、注销账号 -->
<!-- 根据用户UID来显示所有内容 -->
<script setup>
import { UserHead as TinyUserHead } from '@opentiny/vue'
import { Grid as TinyGrid, GridColumn as TinyGridColumn } from '@opentiny/vue'
import { Collapse as TinyCollapse, CollapseItem as TinyCollapseItem } from '@opentiny/vue'
import { ref } from 'vue'

const activeNames = ref(['1'])
</script>

<template>
  <div>
    <h1>个人资料</h1>
  </div>
  <div class="parent">
      <tiny-user-head type="icon" round></tiny-user-head>
      <h3 style="margin-left: 20px;">Username</h3>
  </div>
  <h3>Score:</h3>
  <div class="parent">
    <div class="restrictmaincomponents">
      <tiny-collapse v-model="activeNames">
          <tiny-collapse-item title="简介" name="1">
              <p>这个人很懒，还没有写简介。</p>
          </tiny-collapse-item>
      </tiny-collapse>
    </div>
  </div>
  <div class="parent" style="margin-top: 30px;">
        <div class="restrictmaincomponents">
          <tiny-grid seq-serial min-height="800" align="center" height="20" :resizable="false">
              <tiny-grid-column field="question" title="已提交的题目"></tiny-grid-column>
              <tiny-grid-column field="status" title="状态" width="15%"></tiny-grid-column>
          </tiny-grid>
        </div>
  </div>
</template>