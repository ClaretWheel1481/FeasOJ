<!-- 关于页面 -->
<script setup>
import { VImg, VContainer } from 'vuetify/components'

function opengithub() {
    window.open('https://github.com/ClaretWheel1481/FeasOJ')
}
</script>

<template>
    <div class="title">
        <h1>关于</h1>
    </div>
    <div style="margin-top: 50px;">
        <h3>有问题反馈或想参与开发？前往Github！</h3>
    </div>
    <v-container class="d-flex justify-center align-center">
        <v-img src="https://res.hc-cdn.com/tinyui-design-common/1.0.5.20240401143336/assets/github.svg" max-width="50"
            @click="opengithub" style="cursor: pointer;"></v-img>
    </v-container>
</template>