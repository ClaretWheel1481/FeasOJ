<script setup>
import { VAppBar,VBtn,VDivider,VCard,VCardText,VProgressCircular } from 'vuetify/components'
import { ref,onMounted,computed } from 'vue'
import axios from 'axios';
import { useRoute } from 'vue-router';

const route = useRoute();
const loading = ref(true)
const problemInfo = ref({});
const userName = ref(localStorage.getItem('username'))
const token = ref(localStorage.getItem('token'))

// 计算属性来判断用户是否已经登录
const userLoggedIn = computed(() => !!token.value)

onMounted(async () => {
    if(userLoggedIn.value){
        loading.value = true;
        try {
            const problemId = route.params.Pid;
            const response = await axios.get(`http://127.0.0.1:37881/api/v1/getProblemInfo/${problemId}`);
            if (response.status === 200) {
            problemInfo.value = response.data.problemInfo;
            } else {
            console.error('获取题目信息失败:', response);
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
        <h1>{{problemInfo.Title}}</h1>
        <div style="margin: 10px;"></div>
        <p class="subtitle">时间限制: {{problemInfo.Timelimit}} S</p>
        <p class="subtitle">内存限制: {{problemInfo.Memorylimit}} MB</p>
        <div style="margin: 10px;"></div>
        <v-divider></v-divider>
        <div style="margin: 10px"></div>
        <v-card class="mx-auto my-8" width="80%" elevation="5" rounded="xl">
            <template v-slot:title>
                <span class="font-weight-black">题目详细</span>
            </template>
            <v-card-text>
                {{ problemInfo.Content }}
            </v-card-text>
        </v-card>
        <div style="margin: 10px"></div>
        <v-card class="mx-auto my-8" width="80%" elevation="5" rounded="xl">
            <template v-slot:title>
                <span class="font-weight-black">输入样例</span>
            </template>
            <v-card-text>
                {{ problemInfo.Input }}
            </v-card-text>
        </v-card>
        <div style="margin: 10px"></div>
        <v-card class="mx-auto my-8" width="80%" elevation="5" rounded="xl">
            <template v-slot:title>
                <span class="font-weight-black">输出样例</span>
            </template>
            <v-card-text>
                {{ problemInfo.Output }}
            </v-card-text>
        </v-card>
        <v-card class="mx-auto my-8 editor-container" width="80%" elevation="5" rounded="xl">
            <!-- <v-select v-model="selectedLanguage" :items="languages" label="选择语言" rounded="xl"></v-select> -->
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