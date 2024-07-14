<script setup>
import { VApp,VMain } from 'vuetify/components';
import VNavigation from './components/VNavigation.vue'
import VFill from './components/VFill.vue';
import VAlert from './components/VAlert.vue';
import { RouterView } from 'vue-router';
</script>

<template>
  <v-app>
    <VAlert />
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
