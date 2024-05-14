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

// function SendEmail(){
//     // TODO:发送验证码并限制60秒内只能点击一次
//     console.log('图标被点击了!')
// }
</script>

<template>
    <h1>注册</h1>
    <div class="parent">
            <tiny-card size="large">
                <div>
                    <h3>昵称</h3>
                </div>
                <div class="input">
                    <!-- TODO:限制不超过16个字符 -->
                    <tiny-input v-model="username" placeholder="Username"></tiny-input>
                </div>
                <div>
                    <h3>邮箱</h3>
                </div>
                <div class="input">
                    <!-- TODO:邮箱格式验证 -->
                    <tiny-input v-model="mails" placeholder="Email"></tiny-input>
                </div>
                <div>
                    <h3>邮箱验证码</h3>
                </div>
                <div>
                    <div class="input" style="display: inline-block; margin-left: 80px;">
                        <tiny-input v-model="verificationcode" placeholder="点击右侧按钮获取验证码"></tiny-input>
                    </div>
                    <div style="display: inline-block;">
                        <tiny-button type="text">发送验证码</tiny-button>
                    </div>
                </div>
                <!-- TODO:右侧“获取验证码”按钮 -->
                <div>
                    <h3>密码</h3>
                </div>
                <div class="input">
                    <!-- TODO:密码安全性验证 -->
                    <tiny-input type="password" v-model="passwords" show-password placeholder="Password"></tiny-input>
                </div>
                <div>
                    <h3>确认密码</h3>
                </div>
                <div class="input">
                    <!-- TODO:检验密码是否相同 -->
                    <tiny-input type="password" v-model="confirmpassword" show-password placeholder="Confirm Password"></tiny-input>
                </div>
                <div style="height: 18px"></div>
                <!-- TODO:若以上校验条件不达成，按下后跳转至不达成项 -->
                <tiny-button type="primary" @click="$router.push('/login')">确认</tiny-button>
            </tiny-card>
        </div>
</template>
