<script setup>
import { VApp,VMain } from 'vuetify/components';
import VNavigation from './components/VNavigation.vue'
</script>

<template>
  <v-app>
    <VNavigation></VNavigation>
    <v-main>
        <RouterView></RouterView>
    </v-main>
  </v-app>
</template>

<style>

</style>
