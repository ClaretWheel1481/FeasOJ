<script setup>
import { Input as TinyInput } from '@opentiny/vue'
import { Card as TinyCard } from '@opentiny/vue'
import { iconMailContent } from '@opentiny/vue-icon'
import { ref } from 'vue'
import { Button as TinyButton } from '@opentiny/vue'

const mails = ref('')
const verificationcode = ref('')

const TinyIconMailContent = iconMailContent()
// function SendEmail(){
//     // TODO:发送验证码并限制60秒内只能点击一次
//     console.log('图标被点击了!')
// }
</script>

<template>
    <div>
      <h1>忘记密码</h1>
    </div>
    <div class="parent">
            <tiny-card size="large">
                <div>
                    <h3>邮箱</h3>
                </div>
                <div class="input">
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
                <div style="height: 18px"></div>
                <tiny-button type="primary" @click="$router.push('/login')">下一步</tiny-button>
                <tiny-button  @click="$router.push('/login')">返回</tiny-button>
            </tiny-card>
    </div>
    <!-- TODO:按下下一步后弹出新的组件 -->
</template>
