<script setup>
import { Grid as TinyGrid, GridColumn as TinyGridColumn } from '@opentiny/vue'
</script>

<template>
  <div>
    <h1>状态</h1>
  </div>
  <div class="parent" style="margin-top: 30px;">
        <div class="restrictmaincomponents">
          <tiny-grid seq-serial min-height="800" align="center" height="20" :resizable="false">
            <tiny-grid-column field="user" title="用户" width="20%"></tiny-grid-column>
              <tiny-grid-column field="question" title="已提交的题目"></tiny-grid-column>
              <tiny-grid-column field="status" title="状态" width="15%"></tiny-grid-column>
          </tiny-grid>
        </div>
  </div>
</template>