<script setup>
import { ref, computed, onMounted } from "vue";
import { useRouter } from "vue-router";
import { verifyUserInfo } from "../utils/api/auth";
import { token, userName } from "../utils/account";
import { showAlert } from "../utils/alert";
import { useI18n } from 'vue-i18n';

const { t } = useI18n();
const { locale } = useI18n();

const langs = ref([
  { title: "English", value: "en" },
  { title: "Español", value: "es" },
  { title: "Français", value: "fr" },
  { title: "Italiano", value: "it" },
  { title: "日本語", value: "ja" },
  { title: "Русский", value: "ru" },
  { title: "简体中文", value: "zh_CN" },
  { title: "繁體中文", value: "zh_TW" },
]);

const privilege = ref("");
// 计算属性来判断用户是否已经登录
const userLoggedIn = computed(() => !!token.value);

// 路由实例
const router = useRouter();

// 改变语言
const changeLanguage = (lang) => {
  location.reload();
  locale.value = lang;
  localStorage.setItem('language', lang);
};

// 根据用户登录状态修改导航目的
const navigate = async () => {
  if (userLoggedIn.value) {
    router.push("/profile/" + userName.value);
  } else {
    router.push("/login");
  }
};

// 校验Token是否过期
const isTokenExpired = () => {
  const expirationTime = localStorage.getItem('tokenExpiration');
  if (!expirationTime) {
    return true;
  }
  return Date.now() > expirationTime;
}

// 校验Token后获取用户信息，若privilege为1则表明是管理员
onMounted(async () => {
  if (userLoggedIn.value) {
    if (isTokenExpired()) {
      showAlert(t("message.tokenExpired") + "!", "#/login");
      localStorage.clear();
      userLoggedIn.value = false;
      return;
    }
    try {
      const resp = await verifyUserInfo(userName.value, token.value);
      privilege.value = resp.data.info.role;
    } catch (error) {
      showAlert(t("message.tokenCheckfailed") + "!", "reload");
      localStorage.clear();
      userLoggedIn.value = false;
    }
  }
});
</script>

<template>
  <v-navigation-drawer :width="175" permanent location="left" :elevation="2">
    <v-list nav style="display: flex; flex-direction: column; height: 100%">
      <v-list-item rounded="xl" prepend-icon="mdi-home" value="HOME" @click="router.push('/')" color="primary"
        class="list-item">
        <template v-slot:title>
          <span class="multi-line-title">{{ $t('message.mainpage') }}</span>
        </template>
      </v-list-item>
      <v-list-item rounded="xl" prepend-icon="mdi-file" value="PROBLEMSET" @click="router.push('/problemset')"
        color="primary" class="list-item">
        <template v-slot:title>
          <span class="multi-line-title">{{ $t('message.problemset') }}</span>
        </template>
      </v-list-item>
      <v-list-item rounded="xl" prepend-icon="mdi-code-array" value="COMPETITION" color="primary"
        @click="router.push('/competitions')" class="list-item">
        <template v-slot:title>
          <span class="multi-line-title">{{ $t('message.competition') }}</span>
        </template>
      </v-list-item>
      <v-list-item rounded="xl" prepend-icon="mdi-chat" value="DISCUSS" color="primary"
        @click="router.push('/discussion')" class="list-item">
        <template v-slot:title>
          <span class="multi-line-title">{{ $t('message.discussion') }}</span>
        </template>
      </v-list-item>
      <v-list-item rounded="xl" prepend-icon="mdi-playlist-play" value="RANK" @click="router.push('/rank')"
        color="primary" class="list-item">
        <template v-slot:title>
          <span class="multi-line-title">{{ $t('message.rank') }}</span>
        </template>
      </v-list-item>
      <v-list-item rounded="xl" prepend-icon="mdi-checkbox-multiple-marked" value="STATUS"
        @click="router.push('/status')" color="primary" class="list-item">
        <template v-slot:title>
          <span class="multi-line-title">{{ $t('message.status') }}</span>
        </template>
      </v-list-item>
      <v-divider></v-divider>
      <div class="flex-grow-space"></div>
      <v-list-item v-if="privilege === 1" rounded="xl" prepend-icon="mdi-tie" @click="router.push('/admin')"
        value="admin" base-color="primary" class="list-item">
        <template v-slot:title>
          <span class="multi-line-title">{{ $t('message.management') }}</span>
        </template>
      </v-list-item>
      <v-menu location="end">
        <template v-slot:activator="{ props }">
          <v-list-item rounded="xl" prepend-icon="mdi-translate" base-color="primary" v-bind="props" class="list-item">
            <template v-slot:title>
              <span class="multi-line-title">{{ $t('message.lang') }}</span>
            </template>
          </v-list-item>
        </template>
        <v-list rounded="xl">
          <v-list-item v-for="(item, index) in langs" :key="index" :value="index" @click="changeLanguage(item.value)">
            <v-list-item-title style="font-size: 13px;font-weight: bold;">{{ item.title }}</v-list-item-title>
          </v-list-item>
        </v-list>
      </v-menu>
      <v-list-item rounded="xl" :prepend-icon="userLoggedIn ? 'mdi-account-circle' : 'mdi-account'" @click="navigate"
        value="PROFILE" base-color="primary" class="list-item">
        <template v-slot:title>
          <span class="multi-line-title">{{ userLoggedIn ? userName : $t('message.login') }}</span>
        </template>
      </v-list-item>
    </v-list>
  </v-navigation-drawer>
</template>

<style>
.flex-grow-space {
  flex-grow: 1;
}

.list-item {
  height: 40px;
  overflow: hidden;
}

.multi-line-title {
  white-space: normal;
  word-wrap: break-word;
}
</style>
