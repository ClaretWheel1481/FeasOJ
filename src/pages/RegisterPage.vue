<!-- 注册页 -->
<script setup>
import { ref, reactive } from 'vue';
import { VAppBar, VBtn, VTextField, VForm, VSheet, VIcon } from 'vuetify/components';
import { getCaptchaCode, registerRequest } from '../utils/axios';
import { rules, regex } from '../utils/rules';
import { showAlert } from '../utils/alert';

const formState = reactive({
  username: '',
  userEmail: '',
  password: '',
  confirmPassword: '',
  vcode: ''
});
const networkloading = ref(false);
const isButtonDisabled = ref(false);
const countdown = ref(60);

const register = async () => {
  if (formState.username === "" || formState.userEmail === "" || formState.password === "" || formState.confirmPassword === "" || formState.vcode === "") {
    showAlert('请确认所有内容都已经输入。', "");
    return;
  }
  if (formState.password.length < 8) {
    showAlert('密码长度至少为8个字符。', "");
    return;
  }
  if (formState.password !== formState.confirmPassword) {
    showAlert('两次输入的密码不一致，请重新输入。', "");
    return;
  }
  try {
    networkloading.value = true;
    const response = await registerRequest(formState.username, formState.password, formState.userEmail, formState.vcode);
    if (response.data.status === 200) {
      networkloading.value = false;
      showAlert(response.data.message, "/login");
      return;
    } else {
      networkloading.value = false;
      showAlert(response.data.message, "");
      return;
    }
  } catch (error) {
    networkloading.value = false;
    showAlert(error.response.data.message, "");
    return;
  }
};

const getCaptcha = async () => {
  if (!regex.test(formState.userEmail)) {
    showAlert("请输入有效的邮箱地址。", "");
    return;
  }
  if (!formState.userEmail) {
    showAlert("请输入邮箱地址。", "");
    return;
  }
  try {
    networkloading.value = true;
    const response = await getCaptchaCode(formState.userEmail);
    if (response.data.status === 200) {
      networkloading.value = false;
      showAlert('验证码发送成功，请检查您的邮箱。', "");
      if (isButtonDisabled.value) {
        return;
      }
      isButtonDisabled.value = true;
      let timer = setInterval(() => {
        if (countdown.value > 0) {
          countdown.value--;
        } else {
          clearInterval(timer);
          isButtonDisabled.value = false;
          countdown.value = 60; // 重置倒计时
        }
      }, 1000);
    } else {
      networkloading.value = false;
      showAlert("验证码发送失败，请稍后再试。", "");
    }
  } catch (error) {
    networkloading.value = false;
    showAlert('验证码请求失败，请稍后再试。', "");
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
  <v-app-bar :elevation="0">
    <template v-slot:prepend>
      <v-btn icon="mdi-chevron-left" size="x-large" @click="$router.back"></v-btn>
    </template>
  </v-app-bar>
  <div class="title">
    <h1>注册</h1>
  </div>
  <v-sheet class="constrainsheet" rounded="xl" :elevation="10">
    <v-form fast-fail width="400px" class="mx-auto" @submit.prevent="register">
      <v-text-field v-model="formState.username" :rules="[rules.username.required]" rounded="xl" variant="solo-filled"
        label="用户名" />
      <v-text-field v-model="formState.userEmail" :rules="[rules.userEmail.required, rules.userEmail.email]"
        rounded="xl" variant="solo-filled" label="邮箱" />
      <v-text-field v-model="formState.vcode" :rules="[rules.vcode.required]" rounded="xl" variant="solo-filled"
        label="邮箱验证码">
        <template v-slot:append>
          <v-btn @click="getCaptcha" :disabled="isButtonDisabled" size="25" icon>
            <v-icon v-if="!isButtonDisabled" icon="mdi-email" />
            <span v-if="isButtonDisabled">{{ countdown }}</span>
          </v-btn>
        </template>
      </v-text-field>
      <v-text-field v-model="formState.password" :rules="[rules.password.required, rules.password.minLength]"
        rounded="xl" variant="solo-filled" type="password" label="密码" />
      <v-text-field v-model="formState.confirmPassword"
        :rules="[rules.confirmPassword.required, rules.confirmPassword.minLength]" rounded="xl" variant="solo-filled"
        type="password" label="确认密码" />
      <v-btn type="submit" color="primary" rounded="xl">注册</v-btn>
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