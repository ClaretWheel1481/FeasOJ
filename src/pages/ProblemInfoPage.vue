<!-- 题目详细页 -->
<script setup>
import { VAppBar,VBtn,VDivider,VCard,VCardText,VProgressCircular,VSelect } from 'vuetify/components'
import { ref,onMounted,computed,watch } from 'vue'
import { useRoute } from 'vue-router';
import { getPbDetails,uploadCode } from '../utils/axios';
import { VAceEditor } from 'vue3-ace-editor';
import { showAlert } from '../utils/alert';
import { token,userName } from "../utils/account";

const route = useRoute();
const loading = ref(true)
const problemInfo = ref({});
const content = ref('');
// 代码模板
const templates = {
java: 
`public class Main {
    public static void main(String[] args) {

    }
}`,
c_cpp: 
`#include <iostream>
using namespace std;

int main() {

    return 0;
}`,
golang: 
`package main

import "fmt"

func main() {

}`,
python:
``
};
const lang = ref('python');

const langFileExtension = {
    java: 'java',
    c_cpp: 'cpp',
    golang: 'go',
    python: 'py'
};
// 计算属性来判断用户是否已经登录
const userLoggedIn = computed(() => !!token.value)

onMounted(async () => {
    loading.value = true;
    if(userLoggedIn.value){
        try {
            const problemId = route.params.Pid;
            const resp = await getPbDetails(problemId);
            if (resp.status === 200) {
                problemInfo.value = resp.data.problemInfo;
            }
        } catch (error) {
            showAlert('请求题目信息时发生错误。',"");
        } finally{
            loading.value = false;
        }
    }else{
        window.location='/login'
    }

    watch(lang, (newLang) => {
        content.value = templates[newLang];
    });
});

// 代码上传
const uploadContentAsFile = async () => {
    const blob = new Blob([content.value], { type: 'text/plain' });
    const codefile = new File([blob], `main.${langFileExtension[lang.value]}`, { type: 'text/plain' });
    try {
        const uploadResp = await uploadCode(codefile,route.params.Pid,userName.value,token.value);
        if(uploadResp.status === 200){
            showAlert('提交成功',"reload")
        }
    } catch (error) {
        showAlert('提交代码时发生错误。',"");
    }
};
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
        <v-card class="mx-auto my-8" width="80%" elevation="5" rounded="xl">
            <template v-slot:title>
                <span class="font-weight-black">题目详细</span>
            </template>
            <v-card-text>
                {{ problemInfo.content }}
            </v-card-text>
        </v-card>
        <div style="margin: 10px"></div>
        <v-card class="mx-auto my-8" width="80%" elevation="5" rounded="xl">
            <template v-slot:title>
                <span class="font-weight-black">输入样例</span>
            </template>
            <v-card-text>
                {{ problemInfo.input }}
            </v-card-text>
        </v-card>
        <div style="margin: 10px"></div>
        <v-card class="mx-auto my-8" width="80%" elevation="5" rounded="xl">
            <template v-slot:title>
                <span class="font-weight-black">输出样例</span>
            </template>
            <v-card-text>
                {{ problemInfo.output }}
            </v-card-text>
        </v-card>
        <v-card class="mx-auto my-8" width="80%" height="1000" elevation="5" rounded="xl">
            <v-select
                label="选择语言"
                v-model="lang"
                :items="['python','c_cpp', 'golang', 'java']"
                variant="solo-filled"
            ></v-select>
            <v-ace-editor
                v-model:value="content"
                theme="chrome"
                :lang=lang
                style="height: 1000px;font-size: 16px;"
            />
        </v-card>
        <v-btn color="primary" rounded="xl" @click="uploadContentAsFile">提交</v-btn>
    </div>
</template>

<style scoped>
.subtitle {
    font-size: 0.9rem;
}

.loading{
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100%;
}
</style>