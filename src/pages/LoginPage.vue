<script setup>
import { Card as TinyCard } from '@opentiny/vue'
import { ref } from 'vue'
import { Input as TinyInput } from '@opentiny/vue'
import { Button as TinyButton } from '@opentiny/vue'

const username = ref('')
const passwords = ref('')
</script>

<template>
    <!-- TODO:新的Navigator -->
    <div class="parent">
        <div class="card-wrap">
            <tiny-card size="large">
                <div>
                    <h3>邮箱</h3>
                </div>
                <div class="input">
                    <tiny-input v-model="username" placeholder="Email"></tiny-input>
                </div>
                <div>
                    <h3>密码</h3>
                </div>
                <div class="input">
                    <tiny-input type="password" v-model="passwords" show-password placeholder="Password"></tiny-input>
                </div>
                <div style="height: 12px"></div>
                <!-- TODO:跳转待处理 -->
                <tiny-button type="text" text="找回密码" @click="$router.push('/')"> </tiny-button>
                <div style="height: 12px"></div>
                <div>
                    <!-- TODO:跳转待处理 -->
                    <tiny-button type="primary" @click="$router.push('/')">登录</tiny-button>
                    <tiny-button @click="$router.push('/register')">注册</tiny-button>
                </div>
            </tiny-card>
        </div>
    </div>
</template>

<style scoped>
.card-wrap {
  background: #f5f5f5;
  padding: 16px;
}

.input .tiny-input {
  width: 250px;
}
</style>