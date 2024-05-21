<script setup>
import { Card as TinyCard } from '@opentiny/vue'
import { ref, reactive } from 'vue'
import { Form as TinyForm, FormItem as TinyFormItem, Input as TinyInput, Button as TinyButton, Row as TinyRow } from '@opentiny/vue'

const ruleFormRef = ref()
const createData = reactive({
  email: '',
  newpassword: '',
})
const emailcode = ref('')

let validatePass = (rule, value, callback) => {
  if (!/^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)[a-zA-Z\d]{8,}$/.test(value)) {
    callback(new Error('最少八个字符，至少包含一个大写字母，一个小写字母和一个数字'))
  } else {
    callback()
  }
}

let confirmPasswordValidate = (rule, value, callback) => {
  if (value !== createData.password) {
    callback(new Error('两次输入密码不一致'))
  } else {
    callback()
  }
}
const isValidate = ref(true)
const rules = ref({
  email: [
    { required: true, message: '必填', trigger: 'blur' },
    { type: 'email', message: '请输入正确的邮箱地址', trigger: ['change', 'blur'] }
  ],
  emailcode: [
    { required: true, message: '必填', trigger: 'blur' },
    { min: 4, max: 4, message: '长度必须为4', trigger: ['change', 'blur'] }
  ],
  newpassword: [
    { required: true, message: '必填', trigger: 'blur' },
    { validator: validatePass, trigger: 'blur' }
  ],
  confirmPassword: [
    { required: true, message: '必填', trigger: 'blur' },
    { validator: confirmPasswordValidate, trigger: 'blur' }
  ]
})

function handleSubmit() {
  ruleFormRef.value.validate(() => {
    // TODO:验证输入后调用转移至新的组建，旧组件隐藏
  })
}
</script>

<template>
    <div>
      <h1>修改密码</h1>
    </div>
    <div class="parent">
        <tiny-card size="large">
            <tiny-form ref="ruleFormRef" :label-align="true" label-position="right" label-width=80px :inline="true" :model="createData" :rules="rules" :validate-on-rule-change="isValidate" validate-position="right" class="input">
                <tiny-form-item label="邮箱" prop="email">
                    <tiny-input v-model="createData.email"></tiny-input>
                </tiny-form-item>
                <tiny-row>
                    <tiny-form-item label="验证码" prop="emailcode" style="margin-left: 80px;">
                        <tiny-input v-model="emailcode"></tiny-input>
                    </tiny-form-item>
                    <!-- TODO:发送验证码待处理 -->
                    <tiny-button type="text">发送验证码</tiny-button>
                </tiny-row>
                <tiny-form-item label="新密码" prop="newpassword">
                    <tiny-input v-model="createData.newpassword" type="password" show-password></tiny-input>
                </tiny-form-item>
                <tiny-form-item label="确认密码" prop="confirmPassword">
                    <tiny-input v-model="createData.confirmPassword" type="password" show-password></tiny-input>
                </tiny-form-item>
                <br>
                <tiny-button type="primary" @click="handleSubmit()">下一步</tiny-button>
                <tiny-button @click="$router.push('/login')">返回</tiny-button>
            </tiny-form>
        </tiny-card>
    </div>
</template>
