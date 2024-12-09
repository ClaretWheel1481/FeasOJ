<!-- 注册页 -->
<script setup>
import { ref, reactive } from 'vue';
import { getCaptchaCode, registerRequest } from '../../utils/api/auth';
import { rules, regex } from '../../utils/rules';
import { showAlert } from '../../utils/alert';
import { useI18n } from 'vue-i18n';

const { t } = useI18n();

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

// 校验表单内容并注册
const register = async () => {
  if (formState.username === "" || formState.userEmail === "" || formState.password === "" || formState.confirmPassword === "" || formState.vcode === "") {
    showAlert(t("message.formCheckfailed") + "!", "");
    return;
  }
  if (formState.password.length < 8) {
    showAlert(t("message.formRuleCheckfailed") + "!", "");
    return;
  }
  if (formState.password !== formState.confirmPassword) {
    showAlert(t("message.formRuleCheckfailed") + "!", "");
    return;
  }
  try {
    networkloading.value = true;
    const response = await registerRequest(formState.username, formState.password, formState.userEmail, formState.vcode);
    if (response.status === 200) {
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

// 获取验证码
const getCaptcha = async () => {
  if (!regex.test(formState.userEmail)) {
    showAlert(t("message.checkEmail") + "!", "");
    return;
  }
  if (!formState.userEmail) {
    showAlert(t("message.checkEmail") + "!", "");
    return;
  }
  try {
    networkloading.value = true;
    const response = await getCaptchaCode(formState.userEmail, "true");
    networkloading.value = false;
    showAlert(response.data.message, "");
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
  } catch (error) {
    networkloading.value = false;
    showAlert(error.response.data.message, "");
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
    <h1>{{ $t('message.register') }}</h1>
  </div>
  <v-sheet class="constrainsheet" rounded="xl" :elevation="10">
    <v-form fast-fail width="400px" class="mx-auto" @submit.prevent="register">
      <v-text-field v-model="formState.username" :rules="[rules.username.required]" rounded="xl" variant="solo-filled"
        :label="$t('message.username')" />
      <v-text-field v-model="formState.userEmail" :rules="[rules.userEmail.required, rules.userEmail.email]"
        rounded="xl" variant="solo-filled" :label="$t('message.email')" />
      <v-text-field v-model="formState.vcode" :rules="[rules.vcode.required]" rounded="xl" variant="solo-filled"
        :label="$t('message.vCode')">
        <template v-slot:append>
          <v-btn @click="getCaptcha" :disabled="isButtonDisabled" size="25" icon>
            <v-icon v-if="!isButtonDisabled" icon="mdi-email" />
            <span v-if="isButtonDisabled">{{ countdown }}</span>
          </v-btn>
        </template>
      </v-text-field>
      <v-text-field v-model="formState.password" :rules="[rules.password.required, rules.password.minLength]"
        rounded="xl" variant="solo-filled" type="password" :label="$t('message.password')" />
      <v-text-field v-model="formState.confirmPassword"
        :rules="[rules.confirmPassword.required, rules.confirmPassword.minLength]" rounded="xl" variant="solo-filled"
        type="password" :label="$t('message.confirmPwd')" />
      <v-btn type="submit" color="primary" rounded="xl">{{ $t('message.register') }}</v-btn>
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