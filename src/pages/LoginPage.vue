<!-- 登录页 -->
<script setup>
import { reactive } from 'vue';
import { VBtn, VTextField, VForm, VSheet, VRow } from 'vuetify/components';
import { loginRequest, getUserInfo } from '../utils/axios.js'
import { showAlert } from '../utils/alert';

const networkloading = ref(false);
const forms = reactive({
  username: '',
  password: '',
});

// 登录逻辑
const login = async () => {
  if (forms.username === '' || forms.password === '') {
    showAlert('用户名或密码不能为空。', "");
    return;
  }
  try {
    const loginResponse = await loginRequest(forms.username, forms.password);
    networkloading.value = true;
    if (loginResponse.data.status === 200) {
      localStorage.setItem('token', loginResponse.data.token)
      localStorage.setItem('username', forms.username)
      const response = await getUserInfo(forms.username, loginResponse.data.token);
      localStorage.setItem('userid', response.data.Info.uid);
      networkloading.value = false;
      showAlert(loginResponse.data.message, "/");
    } else {
      networkloading.value = true;
      showAlert(loginResponse.data.message, "");
      return;
    }
  } catch (error) {
    networkloading.value = true;
    showAlert(error.response.data.message, "");
    return;
  }
}
</script>

<template>
  <v-dialog v-model="networkloading" max-width="600px">
    <v-card>
        <div class="networkloading">
            <v-progress-circular indeterminate color="primary" :width="12" :size="100"></v-progress-circular>
        </div>
    </v-card>
  </v-dialog>
  <div class="title">
    <h1>登录</h1>
  </div>
  <v-sheet class="constrainsheet" rounded="xl" :elevation="10">
    <v-form fast-fail width="400px" class="mx-auto" @submit.prevent="login" style="margin: 20px;">
      <v-text-field v-model="forms.username" rounded="xl" variant="solo-filled" label="用户名" />
      <v-text-field v-model="forms.password" rounded="xl" variant="solo-filled" type="password" label="密码" />
      <v-row justify="end">
        <v-btn type="submit" color="primary" rounded="xl">登录</v-btn>
        <v-btn color="primary" variant="text" rounded="xl" @click="$router.push('/register')">注册</v-btn>
        <v-btn color="primary" variant="text" rounded="xl" @click="$router.push('/reset')">忘记密码</v-btn>
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