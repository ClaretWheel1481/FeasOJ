<!-- 重置密码页 -->
<script setup>
import { VSheet,VForm,VTextField,VBtn,VAppBar } from 'vuetify/lib/components/index.mjs';
import { ref,reactive } from 'vue';
import { updatePassword, getCaptchaCode } from '../utils/axios';
import { regex,rules } from '../utils/rules'
import { showAlert } from '../utils/alert';

const forms = reactive({
  email: '',
  vcode: '',
  password: '',
  confirmPassword: ''
});

const isButtonDisabled = ref(false);
const countdown = ref(60);

const nextStep = async () => {
    localStorage.clear();
    window.location = '/';
    if (forms.email === "" || forms.password === "" || forms.confirmPassword === "" || forms.vcode === ""){
        showAlert('请确认所有内容都已经输入。',"");
        return;
    }
    if (forms.password.length < 8){
        showAlert('密码长度至少为8个字符。',"");
        return;
    }
    if (forms.password !== forms.confirmPassword) {
        showAlert('两次输入的密码不一致，请重新输入。',"");
        return;
    }
    try{
        const response = await updatePassword(forms.email, forms.vcode, forms.password);
        if (response.data.status === 200) {
            showAlert(response.data.message,"/login");
            return;
        } else {
            showAlert(response.data.message,"");
            return;
        }
    }catch(error){
        showAlert(error.response.data.message,"");
        return;
    }
}

const getCaptcha = async () => {
    if(!regex.test(forms.email)){
        showAlert("请输入有效的邮箱地址。","");
        return;
    }
    if(!forms.email){
        showAlert("请输入邮箱地址。","");
        return;
    }
    try {
        const response = await getCaptchaCode(forms.email);
        if (response.data.status === 200) {
            showAlert('验证码发送成功，请检查您的邮箱。',"");
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
            showAlert("验证码发送失败，请稍后再试。","");
        }
    } catch (error) {
        showAlert('验证码请求失败，请稍后再试。',"");
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
        <h1>重置密码</h1>
    </div>
    <v-sheet class="constrainsheet" rounded="xl" :elevation="10">
        <v-form fast-fail width="400px" class="mx-auto" @submit.prevent="nextStep" style="margin: 20px;">
            <v-text-field v-model="forms.email" :rules="[rules.userEmail.required, rules.userEmail.email]" rounded="xl" variant="solo-filled" label="邮箱" />
            <v-text-field v-model="forms.vcode" :rules="[rules.vcode.required]" rounded="xl" variant="solo-filled" label="邮箱验证码">
                <template v-slot:append>
                    <v-btn icon @click="getCaptcha" :disabled="isButtonDisabled">
                        <v-icon>mdi-email</v-icon>
                        <span v-if="isButtonDisabled">{{ countdown }}</span>
                    </v-btn>
                </template>
            </v-text-field>
            <v-text-field v-model="forms.password" :rules="[rules.password.required, rules.password.minLength]" rounded="xl" variant="solo-filled" type="password" label="密码" />
        <v-text-field v-model="forms.confirmPassword" :rules="[rules.confirmPassword.required, rules.confirmPassword.minLength]" rounded="xl" variant="solo-filled" type="password" label="确认密码" />
        <v-btn type="submit" color="primary" rounded="xl">提交</v-btn>
        </v-form>
    </v-sheet>
</template>