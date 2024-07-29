<!-- 403 -->
<script setup>
import { VEmptyState, VBtn, VCard } from 'vuetify/components';

const back = () => {
  window.location = '#/'
};

const logout = () => {
  localStorage.clear();
  location.reload();
  window.location = '#/';
};

</script>

<template>
  <div style="margin: 10%"></div>
  <v-card class="mx-auto" max-width="50%" min-width="50%" rounded="xl" elevation="10">
    <v-empty-state headline="403" :title="$t('message.nopermission')"></v-empty-state>
    <v-btn color="primary" variant="text" rounded="xl" style="margin-right: 10px;" @click="back">返回首页</v-btn>
    <v-btn color="primary" variant="text" rounded="xl" style="margin-right: 10px;" @click="logout">{{$t('message.logout')}}</v-btn>
    <div style="margin: 10px"></div>
  </v-card>
</template>