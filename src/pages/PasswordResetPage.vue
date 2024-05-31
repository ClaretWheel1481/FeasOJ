<script setup>
import { VSheet,VForm,VTextField,VRow,VBtn,VDialog,VCard,VCardActions,VCardText,VCardTitle,VAppBar } from 'vuetify/lib/components/index.mjs';
import { ref,reactive } from 'vue';
import { useVuelidate } from '@vuelidate/core';
import { useRouter } from 'vue-router';
import axios from 'axios';

const dialog = ref(false);
const dialogMessage = ref('');

const showAlert = (message) => {
  dialogMessage.value = message;
  dialog.value = true;
};

const forms = reactive({
  email: '',
  vcode: '',
  password: '',
  confirmPassword: ''
});

const rules = {
  userEmail: { required: value => !!value || '邮箱是必填项。', email: value => /.+@.+\..+/.test(value) || '请输入有效的邮箱地址。' },
  vcode: { required: value => !!value || '验证码是必填项。',minLength: value => value.length === 4 || '验证码长度必须为4个字符。'},
  password: { required: value => !!value || '密码是必填项。', minLength: value => value.length >= 8 || '密码长度至少为8个字符。' },
  confirmPassword: { required: value => !!value || '确认密码是必填项。', minLength: value => value.length >= 8 || '确认密码长度至少为8个字符。' },
};

const v$ = useVuelidate(rules, forms);

const isButtonDisabled = ref(false);
const countdown = ref(60);

const nextStep = async () => {
    if (forms.email === "" || forms.password === "" || forms.confirmPassword === "" || forms.vcode === ""){
        showAlert('请确认所有内容都已经输入。');
        return;
    }
    if (forms.password.length < 8){
        showAlert('密码长度至少为8个字符。');
        return;
    }
    if (forms.password !== forms.confirmPassword) {
        showAlert('两次输入的密码不一致，请重新输入。');
        return;
    }
    try{
        const response = await axios.post('http://127.0.0.1:37881/api/v2/updatePassword', {
            email: forms.email,
            vcode: forms.vcode,
            newpassword: forms.password
        });
        if (response.data.status === 200) {
            showAlert(response.data.message);
            setTimeout(() => {
                window.location = '/login';
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
}

const getCaptcha = async () => {
  //TODO: 验证邮箱格式
    if(!forms.email){
        showAlert("请输入邮箱地址。");
        return;
    }
    try {
        const response = await axios.get('http://127.0.0.1:37881/api/v1/getCaptcha', {
            params: { email: forms.email }
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
        <h1>重置密码</h1>
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
        <v-form fast-fail width="400px" class="mx-auto" @submit.prevent="nextStep" style="margin: 20px;">
            <v-text-field v-model="forms.email" :rules="[rules.userEmail.required, rules.userEmail.email]" rounded="xl" variant="solo-filled" label="邮箱" />
            <v-text-field v-model="forms.vcode" :rules="[rules.vcode.required]" rounded="xl" variant="solo-filled" label="邮箱验证码">
                <template v-slot:append>
                    <v-btn icon @click="getCaptcha()" :disabled="isButtonDisabled">
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