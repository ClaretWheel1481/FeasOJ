<script setup>
import { ref,reactive } from 'vue';
import { VAppBar, VBtn, VTextField, VForm, VDialog, VCard, VCardTitle, VCardText, VCardActions,VSheet } from 'vuetify/components';
import { useVuelidate } from '@vuelidate/core';
import { useRouter } from 'vue-router';
import axios from 'axios';

const router = useRouter();
const dialog = ref(false);
const dialogMessage = ref('');

const showAlert = async (message) => {
  dialogMessage.value = message;
  dialog.value = true;
};

const rules = {
  username: { required: value => !!value || '用户名是必填项。' },
  userEmail: { required: value => !!value || '邮箱是必填项。', email: value => /.+@.+\..+/.test(value) || '请输入有效的邮箱地址。' },
  password: { required: value => !!value || '密码是必填项。', minLength: value => value.length >= 8 || '密码长度至少为8个字符。' },
  confirmPassword: { required: value => !!value || '确认密码是必填项。', minLength: value => value.length >= 8 || '确认密码长度至少为8个字符。' },
  vcode: { required: value => !!value || '验证码是必填项。' }
};
const state = reactive({
  username: '',
  userEmail: '',
  password: '',
  confirmPassword: '',
  vcode: ''
});
const $v = useVuelidate(rules, state);

const isButtonDisabled = ref(false);
const countdown = ref(60);
const username = ref('');
const userEmail = ref('');
const password = ref('');
const confirmPassword = ref('');
const vcode = ref('');



const register = async () => {
    //FIXME: 表单验证不生效
  $v.value.$touch();
  if ($v.value.$invalid) {
    showAlert('表单填写有误，请检查后再提交。');
    return;
  }
  if (password.value !== confirmPassword.value) {
    showAlert('两次输入的密码不一致，请重新输入。');
    return;
  }
  try {
    const response = await axios.post('http://127.0.0.1:37881/api/v2/register', {
        email: userEmail.value,
        username: username.value,
        password: password.value,
        vcode: vcode.value
    });
    console.log(response.data);
    if (response.data.status === 200) {
      await showAlert('注册成功，请登录。');
      router.push('/login');
      return;
    } else {
      showAlert("注册失败，请确认邮箱或用户名的输入以及是否没有重复。");
    }
  } catch (error) {
    showAlert('注册失败，请确认所有内容已经输入。');
  }
};

const getCaptcha = async () => {
    if(!userEmail.value){
        showAlert("请输入邮箱地址。");
        return;
    }
    try {
        const response = await axios.get('http://127.0.0.1:37881/api/v1/getCaptcha', {
            params: { email: userEmail.value }
        });
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
        <v-btn color="blue darken-1" text @click="dialog = false">关闭</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
  <v-sheet class="constrainsheet">
    <v-form fast-fail width="400px" class="mx-auto" @submit.prevent="register">
      <v-text-field v-model="username" :rules="[rules.username.required]" rounded="xl" variant="solo-filled" label="用户名" />
      <v-text-field v-model="userEmail" :rules="[rules.userEmail.required, rules.userEmail.email]" rounded="xl" variant="solo-filled" label="邮箱" />
      <v-text-field v-model="vcode" :rules="[rules.vcode.required]" rounded="xl" variant="solo-filled" label="邮箱验证码">
        <template v-slot:append>
            <v-btn icon @click="getCaptcha" :disabled="isButtonDisabled">
                <v-icon>mdi-mail</v-icon>
                <span v-if="isButtonDisabled">{{ countdown }}</span>
            </v-btn>
        </template>
      </v-text-field>
      <v-text-field v-model="password" :rules="[rules.password.required, rules.password.minLength]" rounded="xl" variant="solo-filled" type="password" label="密码" />
      <v-text-field v-model="confirmPassword" :rules="[rules.confirmPassword.required, rules.confirmPassword.minLength]" rounded="xl" variant="solo-filled" type="password" label="确认密码" />
      <v-btn type="submit" color="primary" rounded="xl">注册</v-btn>
    </v-form>
  </v-sheet>
</template>
