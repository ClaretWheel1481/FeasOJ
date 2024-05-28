<script setup>
import { VNavigationDrawer,VList,VListItem,VDivider } from 'vuetify/components';
import '@mdi/font/css/materialdesignicons.css';
</script>

<template>
  <v-navigation-drawer :width="120" permanent>
    <v-list nav style="display: flex; flex-direction: column; height: 100%;">
      <v-list-item prepend-icon="mdi-home" title="首页" value="HOME" @click="$router.push('/')"></v-list-item>
      <v-list-item prepend-icon="mdi-file" title="题目" value="PROBLEMSET" @click="$router.push('/problemset')"></v-list-item>
      <v-list-item prepend-icon="mdi-flag" title="竞赛" value="CONTEST"></v-list-item>
      <v-list-item prepend-icon="mdi-comment-processing" title="状态" value="STATUS" @click="$router.push('/status')"></v-list-item>
      <v-list-item prepend-icon="mdi-help-circle" title="关于" value="ABOUT" @click="$router.push('/about')"></v-list-item>
      <v-divider></v-divider>
      <div class="flex-grow-space"></div>
      <v-list-item prepend-icon="mdi-account" title="登录" value="LOGIN" @click="$router.push('/login')"></v-list-item>
    </v-list>
  </v-navigation-drawer>
</template>

<style>
.flex-grow-space {
  flex-grow: 1;
}
</style>