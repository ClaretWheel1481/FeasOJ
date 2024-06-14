<!-- 登录页 -->
<script setup>
import { reactive,ref } from 'vue';
import { VBtn, VTextField, VForm,VSheet,VDialog,VCard,VCardActions,VCardTitle,VCardText,VRow } from 'vuetify/components';
import { loginRequest } from '../utils/axios.js'

const dialog = ref(false);
const dialogMessage = ref('');

const showAlert = (message) => {
  dialogMessage.value = message;
  dialog.value = true;
};

const forms = reactive({
  username: '',
  password: '',
});

// 登录逻辑
const login = async () => {
  if (forms.username === '' || forms.password === '') {
    showAlert('用户名或密码不能为空。');
    return;
  }
  try {
    const response = await loginRequest(forms.username, forms.password);
    if (response.data.status === 200) {
      showAlert(response.data.message);
      localStorage.setItem('token',response.data.token)
      localStorage.setItem('username',forms.username)
      setTimeout(() => {
        window.location = '/'
      }, 500);
    } else {
      showAlert(response.data.message);
      return;
    }
  } catch (error) {
    showAlert(error.response.data.message);
    return;
  }
}

</script>

<template>
  <div class="title">
    <h1>登录</h1>
  </div>
  <v-dialog v-model="dialog" persistent max-width="290">
    <v-card rounded="xl">
      <v-card-title class="text-h5">提示</v-card-title>
      <v-card-text>{{ dialogMessage }}</v-card-text>
      <v-card-actions>
        <v-btn color="blue darken-1" text @click="dialog = false" rounded="xl">关闭</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
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

<style>

</style>