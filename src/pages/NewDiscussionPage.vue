<!-- 创建讨论页 -->
<script setup>
import { VProgressCircular,VAppBar,VRow,VBtn,VTextarea,VForm,VTextField } from 'vuetify/lib/components/index.mjs';
import { ref,computed,onMounted } from 'vue';
import { addDiscussion } from '../utils/axios.js';

const title = ref('');
const content = ref('');
const loading = ref(false);
const token = ref(localStorage.getItem('token'))
const userName = ref(localStorage.getItem('username'))

// 计算属性来判断用户是否已经登录
const userLoggedIn = computed(() => !!token.value)

onMounted(async () => {
    loading.value = true;
    if(!userLoggedIn.value){
        window.location = "/login";
    }else{
        loading.value = false;
    }
});

const submit = async () => {
    loading.value = true;
    const response = await addDiscussion(title.value,content.value,userName.value)
    if(response.status === 200){
        loading.value = false;
        window.location = "/discussion";
    }else{
        loading.value = false;
        alert("创建讨论失败。");
    }
};
</script>

<template>
    <!-- TODO:创建讨论页 -->
    <div v-if="loading" class="loading">
        <v-progress-circular indeterminate color="primary" :width="12" :size="100"></v-progress-circular>
    </div>
    <div v-else>
        <v-app-bar :elevation="0">
            <template v-slot:prepend>
                <v-btn icon="mdi-chevron-left" size="x-large" @click="$router.back"></v-btn>
            </template>
            <v-row style="align-items: center;">
                <div style="margin-left: 30px;"></div>
                <p class="font-weight-black">创建讨论</p>
            </v-row>
        </v-app-bar>
        <div style="margin-top: 60px"></div>
        <div class="form-align">
            <v-form style="min-width: 50%" @submit.prevent="submit">
                <v-text-field v-model="title" label="标题" rounded="xl" variant="solo-filled"></v-text-field>
                <v-textarea v-model="content" label="内容" rounded="xl" variant="solo-filled"></v-textarea>
                <v-btn rounded="xl" type="submit" color="primary">提交</v-btn>
            </v-form>
        </div>
    </div>
</template>

<style scoped>
.loading{
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100%;
}

.form-align{
    display: flex;
    justify-content: center;
    align-items: center;
}
</style>