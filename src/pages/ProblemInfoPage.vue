<!-- 题目详细页 -->
<script setup>
import { VAppBar,VBtn,VDivider,VCard,VCardText,VProgressCircular } from 'vuetify/components'
import { ref,onMounted,computed } from 'vue'
import { useRoute } from 'vue-router';
import { getPbDetails } from '../utils/axios';

const route = useRoute();
const loading = ref(true)
const problemInfo = ref({});
const token = ref(localStorage.getItem('token'))

// 计算属性来判断用户是否已经登录
const userLoggedIn = computed(() => !!token.value)

onMounted(async () => {
    loading.value = true;
    if(userLoggedIn.value){
        try {
            const problemId = route.params.Pid;
            const response = await getPbDetails(problemId);
            if (response.status === 200) {
                problemInfo.value = response.data.problemInfo;
            }
        } catch (error) {
            console.error('请求题目信息时发生错误:', error);
        } finally{
            loading.value = false;
        }
    }else{
        window.location='/login'
    }
});
</script>

<template>
    <v-app-bar :elevation="0">
        <template v-slot:prepend>
            <v-btn icon="mdi-chevron-left" size="x-large" @click="$router.back"></v-btn>
        </template>
    </v-app-bar>
    <div v-if="loading" class="loading">
        <v-progress-circular indeterminate color="primary" :width="12" :size="100"></v-progress-circular>
    </div>
    <div v-else>
        <h1>{{problemInfo.title}}</h1>
        <div style="margin: 10px;"></div>
        <p class="subtitle">时间限制: {{problemInfo.time_limit}} S</p>
        <p class="subtitle">内存限制: {{problemInfo.memory_limit}} MB</p>
        <div style="margin: 10px;"></div>
        <v-divider></v-divider>
        <div style="margin: 10px"></div>
        <v-card class="mx-auto my-8" width="80%" elevation="10" rounded="xl">
            <template v-slot:title>
                <span class="font-weight-black">题目详细</span>
            </template>
            <v-card-text>
                {{ problemInfo.content }}
            </v-card-text>
        </v-card>
        <div style="margin: 10px"></div>
        <v-card class="mx-auto my-8" width="80%" elevation="10" rounded="xl">
            <template v-slot:title>
                <span class="font-weight-black">输入样例</span>
            </template>
            <v-card-text>
                {{ problemInfo.input }}
            </v-card-text>
        </v-card>
        <div style="margin: 10px"></div>
        <v-card class="mx-auto my-8" width="80%" elevation="10" rounded="xl">
            <template v-slot:title>
                <span class="font-weight-black">输出样例</span>
            </template>
            <v-card-text>
                {{ problemInfo.output }}
            </v-card-text>
        </v-card>
        <v-card class="mx-auto my-8 editor-container" width="80%" elevation="10" rounded="xl">
            <!-- TODO:代码编辑器 -->
        </v-card>
    </div>
</template>

<style scoped>
.subtitle {
    font-size: 0.9rem;
}
.editor-container {
  height: 600px;
}
.loading{
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100%;
}
</style>