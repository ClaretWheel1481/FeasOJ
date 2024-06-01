<script setup>
import { VApp,VMain } from 'vuetify/components';
import VNavigation from './components/VNavigation.vue'
import VFill from './components/VFill.vue';
import { RouterView } from 'vue-router';
</script>

<template>
  <v-app>
    <VNavigation></VNavigation>
    <v-main>
        <RouterView></RouterView>
        <div style="margin-top: 50px"></div>
        <!-- 备案组件 -->
        <VFill class="fill"></VFill>
    </v-main>
  </v-app>
</template>

<style>
.fill{
  text-align: center;
  display:inline-block;
}
</style>
