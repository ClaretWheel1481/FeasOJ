<script setup>
import { ref, onMounted, computed } from 'vue';
import { initSSE } from '../utils/notification.js';
import { userId, token } from '../utils/account.js';
import { useI18n } from 'vue-i18n';

const { t } = useI18n();

const events = ref([]);
const snackbar = ref(false);
const snackbarMessage = ref('');

// 计算属性来判断用户是否已经登录
const userLoggedIn = computed(() => !!token.value)

function handleSSEMessage(message) {
    events.value.push(message);
    snackbarMessage.value = message;
    snackbar.value = true;
}

// 连接到服务器发送事件
onMounted(() => {
    if (userLoggedIn.value) {
        initSSE(userId.value, handleSSEMessage);
    }
});
</script>

<template>
    <v-snackbar v-model="snackbar" :timeout="4000" bottom>
        {{ snackbarMessage }}
        <template v-slot:actions>
            <v-btn color="blue" text @click="snackbar = false">
                {{ t('message.ok') }}
            </v-btn>
        </template>
    </v-snackbar>
</template>

<style></style>
