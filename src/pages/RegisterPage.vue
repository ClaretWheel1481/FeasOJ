<!-- 注册页 -->
<script setup>
import { ref,reactive } from 'vue';
import { VAppBar, VBtn, VTextField, VForm, VDialog, VCard, VCardTitle, VCardText, VCardActions,VSheet } from 'vuetify/components';
import { useRouter } from 'vue-router';
import { getCaptchaCode, registerRequest } from '../utils/axios';

const router = useRouter();
const dialog = ref(false);
const dialogMessage = ref('');
const regex = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;

const showAlert = (message) => {
  dialogMessage.value = message;
  dialog.value = true;
};

const rules = {
  username: { required: value => !!value || '用户名是必填项。' },
  userEmail: { required: value => !!value || '邮箱是必填项。', email: value => regex.test(value) || '请输入有效的邮箱地址。' },
  password: { required: value => !!value || '密码是必填项。', minLength: value => value.length >= 8 || '密码长度至少为8个字符。' },
  confirmPassword: { required: value => !!value || '确认密码是必填项。', minLength: value => value.length >= 8 || '确认密码长度至少为8个字符。' },
  vcode: { required: value => !!value || '验证码是必填项。',minLength: value => value.length === 4 || '验证码长度必须为4个字符。'}
};

const formState = reactive({
  username: '',
  userEmail: '',
  password: '',
  confirmPassword: '',
  vcode: ''
});

const isButtonDisabled = ref(false);
const countdown = ref(60);

const register = async () => {
  if (formState.username === "" || formState.userEmail === "" || formState.password === "" || formState.confirmPassword === "" || formState.vcode === ""){
    showAlert('请确认所有内容都已经输入。');
    return;
  }
  if (formState.password.length < 8){
    showAlert('密码长度至少为8个字符。');
    return;
  }
  if (formState.password !== formState.confirmPassword) {
    showAlert('两次输入的密码不一致，请重新输入。');
    return;
  }
  try{
    const response = await registerRequest(formState.username, formState.password, formState.userEmail, formState.vcode);
    if (response.data.status === 200) {
      showAlert(response.data.message);
      // 2秒后跳转
      setTimeout(() => {
        router.push('/login');
      }, 2000);
      return;
    } else {
      showAlert(response.data.message);
      return;
    }
  }catch(error){
    alert(error.response.data.message);
    return;
  }
};

const getCaptcha = async () => {
  if(!regex.test(formState.userEmail)){
    showAlert("请输入有效的邮箱地址。");
      return;
  }
  if(!formState.userEmail){
      showAlert("请输入邮箱地址。");
      return;
  }
  try {
      const response = await getCaptchaCode(formState.userEmail);
      if (response.data.status === 200) {
          showAlert('验证码发送成功，请检查您的邮箱。');
          if(isButtonDisabled.value){
              return;
          }
          isButtonDisabled.value = true;
          let timer = setInterval(() => {
              if(countdown.value > 0){
                  countdown.value--;
              } else {
                  clearInterval(timer);
                  isButtonDisabled.value = false;
                  countdown.value = 60; // 重置倒计时
              }
          }, 1000);
      } else{
          showAlert("验证码发送失败，请稍后再试。");
      }
  } catch (error) {
      showAlert('验证码请求失败，请稍后再试。');
  }
}
</script>

<template>
  <v-app-bar :elevation="0">
    <template v-slot:prepend>
      <v-btn icon="mdi-chevron-left" size="x-large" @click="$router.back"></v-btn>
    </template>
  </v-app-bar>
  <div class="title">
    <h1>注册</h1>
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
  <v-sheet class="constrainsheet"  rounded="xl" :elevation="10">
    <v-form fast-fail width="400px" class="mx-auto" @submit.prevent="register">
      <v-text-field v-model="formState.username" :rules="[rules.username.required]" rounded="xl" variant="solo-filled" label="用户名" />
      <v-text-field v-model="formState.userEmail" :rules="[rules.userEmail.required, rules.userEmail.email]" rounded="xl" variant="solo-filled" label="邮箱" />
      <v-text-field v-model="formState.vcode" :rules="[rules.vcode.required]" rounded="xl" variant="solo-filled" label="邮箱验证码">
        <template v-slot:append>
            <v-btn icon @click="getCaptcha" :disabled="isButtonDisabled">
                <v-icon icon="mdi-email"/>
                <span v-if="isButtonDisabled">{{ countdown }}</span>
            </v-btn>
        </template>
      </v-text-field>
      <v-text-field v-model="formState.password" :rules="[rules.password.required, rules.password.minLength]" rounded="xl" variant="solo-filled" type="password" label="密码" />
      <v-text-field v-model="formState.confirmPassword" :rules="[rules.confirmPassword.required, rules.confirmPassword.minLength]" rounded="xl" variant="solo-filled" type="password" label="确认密码" />
      <v-btn type="submit" color="primary" rounded="xl">注册</v-btn>
    </v-form>
  </v-sheet>
</template>
