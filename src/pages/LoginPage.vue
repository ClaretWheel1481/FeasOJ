<script setup>
import { Card as TinyCard } from '@opentiny/vue'
import { ref } from 'vue'
import { Input as TinyInput } from '@opentiny/vue'
import { Button as TinyButton } from '@opentiny/vue'

const username = ref('')
const passwords = ref('')
</script>

<template>
    <h1>登录</h1>
    <div class="parent">
            <tiny-card size="large">
                <div>
                    <h3>用户名</h3>
                </div>
                <div class="input">
                    <tiny-input v-model="username" placeholder="Username"></tiny-input>
                </div>
                <div>
                    <h3>密码</h3>
                </div>
                <div class="input">
                    <tiny-input type="password" v-model="passwords" show-password placeholder="Password"></tiny-input>
                </div>
                <div style="height: 12px"></div>
                <tiny-button type="text" text="忘记密码?" @click="$router.push('/reset')"> </tiny-button>
                <div style="height: 12px"></div>
                <!-- TODO:若密码不正确，显示“用户名或密码错误，请检查并重试。” -->
                <div>
                    <!-- TODO:跳转待处理 -->
                    <tiny-button type="primary" @click="$router.push('/')">登录</tiny-button>
                    <tiny-button @click="$router.push('/register')">注册</tiny-button>
                </div>
            </tiny-card>
    </div>
</template>