<script setup>
import { ref,onMounted,computed } from 'vue';

const title = ref('')
const content = ref('')
const token = ref('')

onMounted(() => {
  token.value = localStorage.getItem('token')
})
// 计算属性来判断用户是否已经登录
const userLoggedIn = computed(() => !!token.value)

</script>

<template>
    
</template>