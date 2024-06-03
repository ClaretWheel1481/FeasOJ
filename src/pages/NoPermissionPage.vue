<script setup>
import { VEmptyState,VBtn,VCard,VCardActions } from 'vuetify/components';
import { useRouter } from 'vue-router';

const router = useRouter();

const back = () => {
  window.location = '/'
};

const logout = () => {
  localStorage.clear();
  window.location = '/'
};

</script>

<template>
    <div style="margin: 10%"></div>
    <v-card class="mx-auto" max-width="50%" min-width="50%" rounded="xl" elevation="10">
        <v-empty-state
        headline="403"
        title="看起来你没有权限访问"
        text="请确保你的请求正确。"
        ></v-empty-state>
        <v-btn color="primary" variant="text" rounded="xl" style="margin-right: 10px;" @click="back">返回首页</v-btn>
        <v-btn color="primary" variant="text" rounded="xl" style="margin-right: 10px;" @click="logout">退出登录</v-btn>
        <div style="margin: 10px"></div>
    </v-card>
    
</template>