<script setup>
import { useRoute } from 'vue-router';
import Navigation from './components/Navigation.vue';
import Filing from './components/Filing.vue';

const route = useRoute();
</script>

<template>
  <Navigation v-if="route.path !== '/login'"></Navigation>
  <router-view></router-view>
  <Filing></Filing>
</template>

<style scoped>

</style>
