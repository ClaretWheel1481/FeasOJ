<!-- 首页 -->
<script setup>
import { VCard, VExpansionPanel } from 'vuetify/lib/components/index.mjs';
</script>

<template>
    <div class="title">
        <v-img src="logo.png" width="100px" height="100px" style="margin: 20px;"></v-img>
        <h1>FeasOJ</h1>
    </div>
    <v-card rounded="xl" style="margin: 50px;" elevation="5">
        <v-expansion-panels>
            <v-expansion-panel
                title="说明"
                text="1.FeasOJ支持Python、C++(暂不支持C)、Golang、Java
                2.若遇到Bug，欢迎来Github提交Issue或者提交PR帮我修复(
                3.第一次写这种项目，写得不好的地方欢迎指正
                4.还有小部分功能未实现（如竞赛）"
                style="white-space: pre-line;"
            >
            </v-expansion-panel>
        </v-expansion-panels>
    </v-card>
</template>

<style scoped>
.title {
    align-items: center;
    display: flex;
    flex-direction: column;
    justify-content: center;
}
</style>