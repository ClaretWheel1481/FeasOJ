<script setup>
import { Card as TinyCard } from '@opentiny/vue'
import { ref } from 'vue'
import { Input as TinyInput } from '@opentiny/vue'
import { Button as TinyButton } from '@opentiny/vue'

const username = ref('')
const mails = ref('')
const passwords = ref('')
const confirmpassword = ref('')
const verificationcode = ref('')
</script>

<template>
    <div class="parent">
        <div class="card-wrap">
            <tiny-card size="large">
                <div>
                    <h3>昵称</h3>
                </div>
                <div class="input">
                    <tiny-input v-model="username" placeholder="Username"></tiny-input>
                </div>
                <div>
                    <h3>邮箱</h3>
                </div>
                <div class="input">
                    <tiny-input v-model="mails" placeholder="Email"></tiny-input>
                </div>
                <div>
                    <h3>邮箱验证码</h3>
                </div>
                <div class="input">
                    <tiny-input v-model="verificationcode" placeholder="Code"></tiny-input>
                </div>
                <!-- TODO:右侧“获取验证码”按钮 -->
                <div>
                    <h3>密码</h3>
                </div>
                <div class="input">
                    <tiny-input type="password" v-model="passwords" show-password placeholder="Password"></tiny-input>
                </div>
                <div>
                    <h3>确认密码</h3>
                </div>
                <div class="input">
                    <tiny-input type="password" v-model="confirmpassword" show-password placeholder="Confirm Password"></tiny-input>
                </div>
                <div style="height: 18px"></div>
                <tiny-button type="primary" @click="$router.push('/login')">确认</tiny-button>
            </tiny-card>
        </div>
    </div>
</template>

<style>
.input .tiny-input {
  width: 250px;
}
</style>