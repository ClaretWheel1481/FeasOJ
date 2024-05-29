<script setup>
import { VNavigationDrawer,VList,VListItem,VDivider } from 'vuetify/components';
import { ref,computed } from 'vue';
import { useRouter } from 'vue-router';
import '@mdi/font/css/materialdesignicons.css';

const userName = ref(localStorage.getItem('username'))
const token = ref(localStorage.getItem('token'))

// 计算属性来判断用户是否已经登录
const userLoggedIn = computed(() => !!token.value)

// 路由实例
const router = useRouter()

// 根据用户登录状态进行导航
const navigate = () => {
  if (userLoggedIn.value) {
    // router.push('/user-profile')
  } else {
    router.push('/login')
  }
}
</script>

<template>
  <v-navigation-drawer :width="150" permanent>
    <v-list nav style="display: flex; flex-direction: column; height: 100%;">
      <v-list-item rounded="xl" prepend-icon="mdi-home" title="首页" value="HOME" @click="$router.push('/')" color="primary"></v-list-item>
      <v-list-item rounded="xl" prepend-icon="mdi-file" title="题目" value="PROBLEMSET" @click="$router.push('/problemset')" color="primary"></v-list-item>
      <v-list-item rounded="xl" prepend-icon="mdi-flag" title="竞赛" value="CONTEST" @click="$router.push('/contest')" color="primary"></v-list-item>
      <v-list-item rounded="xl" prepend-icon="mdi-comment-processing" title="状态" value="STATUS" @click="$router.push('/status')" color="primary"></v-list-item>
      <v-list-item rounded="xl" prepend-icon="mdi-help-circle" title="关于" value="ABOUT" @click="$router.push('/about')" color="primary"></v-list-item>
      <v-divider></v-divider>
      <div class="flex-grow-space"></div>
      <v-list-item
        rounded="xl"
        :prepend-icon="userLoggedIn ? 'mdi-account-circle' : 'mdi-account'"
        :title="userLoggedIn ? userName : '登录'"
        @click="navigate"
        base-color="primary"
      ></v-list-item>
    </v-list>
  </v-navigation-drawer>
</template>

<style>
.flex-grow-space {
  flex-grow: 1;
}
</style>