<script setup>
import { VNavigationDrawer,VList,VListItem,VDivider } from 'vuetify/components';
import { ref,computed,onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { getUserInfo, verifyJWT } from '../utils/axios';
import { token,userName } from '../utils/account';

const privilege = ref("")
// 计算属性来判断用户是否已经登录
const userLoggedIn = computed(() => !!token.value)

// 路由实例
const router = useRouter()

// 根据用户登录状态修改导航目的
const navigate = () => {
  if (userLoggedIn.value) {
    router.push('/profile/' + userName.value)
  } else {
    router.push('/login')
  }
}

// 校验Token后获取用户信息，若privilege为1则表明是管理员
onMounted(async () => {
  if (userLoggedIn.value) {
      const tokenResp = await verifyJWT(userName.value,token.value);
      if(tokenResp.data.status === 200){
        const response = await getUserInfo(userName.value);
        privilege.value = response.data.Info.role;
      }else {
        router.push('/403')
      }
  }
})
</script>

<template>
  <v-navigation-drawer :width="150" permanent>
    <v-list nav style="display: flex; flex-direction: column; height: 100%;">
      <v-list-item rounded="xl" prepend-icon="mdi-home" title="首页" value="HOME" @click="$router.push('/')" color="primary"></v-list-item>
      <v-list-item rounded="xl" prepend-icon="mdi-file" title="题目" value="PROBLEMSET" @click="$router.push('/problemset')" color="primary"></v-list-item>
      <v-list-item rounded="xl" prepend-icon="mdi-checkbox-multiple-marked" title="状态" value="STATUS" @click="$router.push('/status')" color="primary"></v-list-item>
      <v-list-item rounded="xl" prepend-icon="mdi-chat" title="讨论" value="DISCUSS" color="primary" @click="$router.push('/discussion')"></v-list-item>
      <v-list-item rounded="xl" prepend-icon="mdi-help-circle" title="关于" value="ABOUT" @click="$router.push('/about')" color="primary"></v-list-item>
      <v-divider></v-divider>
      <div class="flex-grow-space"></div>
      <v-list-item
        v-if="privilege===1"
        rounded="xl"
        prepend-icon="mdi-file"
        title="题目管理"
        @click="router.push('/psm')"
        value="BACKEND"
        base-color="primary"
      ></v-list-item>
      <v-list-item
        rounded="xl"
        :prepend-icon="userLoggedIn ? 'mdi-account-circle' : 'mdi-account'"
        :title="userLoggedIn ? userName : '登录'"
        @click="navigate"
        value="PROFILE"
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