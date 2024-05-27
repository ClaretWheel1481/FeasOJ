<script setup>
import { VNavigationDrawer, VToolbar,VList,VListItem,VDivider } from 'vuetify/components';
import '@mdi/font/css/materialdesignicons.css';
</script>

<template>
  <v-navigation-drawer :width="150" permanent>
    <v-list nav>
      <v-list-item prepend-icon="mdi-home" title="首页" value="HOME" @click="$router.push('/')"></v-list-item>
      <v-list-item prepend-icon="mdi-signal" title="排名" value="TOP"></v-list-item>
      <v-list-item prepend-icon="mdi-format-align-justify" title="题目" value="PROBLEMSET"></v-list-item>
      <v-list-item prepend-icon="mdi-file-outline" title="竞赛" value="CONTEST"></v-list-item>
      <v-list-item prepend-icon="mdi-help-circle" title="关于" value="ABOUT" @click="$router.push('/about')"></v-list-item>
      <v-divider></v-divider>
    </v-list>
  </v-navigation-drawer>
</template>