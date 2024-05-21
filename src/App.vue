<script setup>
import { useRoute } from 'vue-router';
import Navigation from './components/Navigation.vue';
import Filing from './components/Filing.vue';

const route = useRoute();
</script>

<template>
  <Navigation v-if="!$route.meta.hideComponent"></Navigation>
  <router-view></router-view>
  <Filing v-if="!$route.meta.hideComponent"></Filing>
</template>

<style scoped>

</style>
