<script setup>
import { VApp,VMain,VToolbar } from 'vuetify/components';
import VNavigation from './components/VNavigation.vue'
import VFooter from './components/Footer.vue'
</script>

<template>
  <v-app>
    <VNavigation></VNavigation>
    <VFooter></VFooter>
    <v-main style="height: 650px">
      <RouterView>
      </RouterView>
    </v-main>
  </v-app>
</template>

<style scoped>

</style>
