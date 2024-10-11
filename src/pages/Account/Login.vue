<!-- 登录页 -->
<script setup>
import { reactive, ref } from 'vue';
import { loginRequest, verifyUserInfo } from '../../utils/axios.js'
import { showAlert } from '../../utils/alert.js';
import { useI18n } from 'vue-i18n';
import { jwtDecode } from "jwt-decode";

const { t } = useI18n();

const networkloading = ref(false);
const forms = reactive({
  username: '',
  password: '',
});

// 登录逻辑
const login = async () => {
  if (forms.username === '' || forms.password === '') {
    showAlert(t("message.formCheckfailed") + "!", "");
    return;
  }
  try {
    networkloading.value = true;
    const loginResponse = await loginRequest(forms.username, forms.password);
    const token = loginResponse.data.token;
    const decodedToken = jwtDecode(token);
    const expirationTime = decodedToken.exp * 1000; // 将秒转换为毫秒
    localStorage.setItem('token', token);
    localStorage.setItem('tokenExpiration', expirationTime);
    const response = await verifyUserInfo(forms.username, token);
    localStorage.setItem('username', response.data.Info.username);
    networkloading.value = false;
    showAlert(loginResponse.data.message, "/");
    return;
  } catch (error) {
    networkloading.value = false;
    showAlert(error.response.data.message, "");
    return;
  }
}

</script>

<template>
  <v-dialog v-model="networkloading" max-width="500px">
    <v-card rounded=xl>
      <div class="networkloading">
        <v-progress-circular indeterminate color="primary" :width="12" :size="100"></v-progress-circular>
      </div>
    </v-card>
  </v-dialog>
  <div class="title">
    <h1>{{ $t('message.login') }}</h1>
  </div>
  <v-sheet class="constrainsheet" rounded="xl" :elevation="10">
    <v-form fast-fail width="400px" class="mx-auto" @submit.prevent="login" style="margin: 20px;">
      <v-text-field v-model="forms.username" rounded="xl" variant="solo-filled"
        :label="$t('message.usernameoremail')" />
      <v-text-field v-model="forms.password" rounded="xl" variant="solo-filled" type="password"
        :label="$t('message.password')" />
      <v-row justify="end">
        <v-btn type="submit" color="primary" rounded="xl">{{ $t('message.login') }}</v-btn>
        <v-btn color="primary" variant="text" rounded="xl" @click="$router.push('/register')">{{ $t('message.register')
          }}</v-btn>
        <v-btn color="primary" variant="text" rounded="xl" @click="$router.push('/reset')">{{ $t('message.forget')
          }}</v-btn>
      </v-row>
    </v-form>
  </v-sheet>
</template>

<style scoped>
.networkloading {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100%;
  margin: 100px;
}
</style>