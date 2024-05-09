<template>
    <h1>FeasOJ</h1>
    <div class="parent">
        <div class="restrictbulletin">
            <tiny-collapse v-model="activeNames">
                <tiny-collapse-item title="公告" name="1">
                    <h2>现在处于测试阶段</h2>
                </tiny-collapse-item>
                <tiny-collapse-item title="更新日志" name="3">
                    <h2>2023.05.09</h2>
                    <p class="pfont">初步外观框架构建完毕</p>
                </tiny-collapse-item>
            </tiny-collapse>
        </div>
        <!-- TODO:组件待扩展 -->
    </div>
</template>

<script setup>
import { ref } from 'vue'
import { Collapse as TinyCollapse, CollapseItem as TinyCollapseItem } from '@opentiny/vue'

const activeNames = ref(['1', '3'])
</script>

<style scoped>
/* 限制公告栏长度为屏幕70% */
.restrictbulletin {
  min-width: 70%;
  max-width: 70%;
  display: inline-block;
  
}

.pfont{ 
    font-size: 14px;
}
</style>