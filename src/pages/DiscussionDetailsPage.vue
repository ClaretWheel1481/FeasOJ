<script setup>
import { VRow,VAppBar,VBtn,VAvatar,VProgressCircular,VImg,VCard,VCardText } from 'vuetify/components';
import { ref,onMounted,computed } from 'vue'
import axios from 'axios';
import { useRoute } from 'vue-router';

const route = useRoute();
const loading = ref(true)
const discussionInfos = ref({});
const token = ref(localStorage.getItem('token'))
// 计算属性来判断用户是否已经登录
const userLoggedIn = computed(() => !!token.value)

onMounted(async () => {
    loading.value = true;
    if(userLoggedIn.value){
        try {
            const Tid = route.params.Tid;
            const response = await axios.get(`http://127.0.0.1:37881/api/v1/getDiscussionByTid/${Tid}`);
            if (response.status === 200) {
                discussionInfos.value = response.data.discussionInfo;
            }
        } catch (error) {
            console.error('请求帖子信息时发生错误:', error);
        } finally{
            loading.value = false;
        }
    }else{
        window.location='/login'
    }
});
</script>

<template>
    <div v-if="loading" class="loading">
        <v-progress-circular indeterminate color="primary" :width="12" :size="100"></v-progress-circular>
    </div>
    <div v-else>
        <v-app-bar :elevation="0">
            <template v-slot:prepend>
                <v-btn icon="mdi-chevron-left" size="x-large" @click="$router.back"></v-btn>
            </template>
            <v-row style="align-items: center;">
                <div style="margin-left: 80px;"></div>
                <v-avatar size="50" color="surface-variant">
                    <v-img :src="discussionInfos.avatar" cover></v-img>
                </v-avatar>
                <div style="margin-left: 10px;"></div>
                <p class="font-weight-black">{{ discussionInfos.username }}</p>
            </v-row>
        </v-app-bar>
        <div style="margin-top: 30px;"></div>
        <v-card class="mx-auto" width="85%" rounded="xl" elevation="10">
            <template v-slot:title>
                <span class="font-weight-black">{{ discussionInfos.title }}</span>
            </template>
            <v-card-text class="bg-surface-light pt-4">
                {{ discussionInfos.content }}
            </v-card-text>
        </v-card>
    </div>
</template>

<style scoped>
.loading{
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100%;
}
</style>