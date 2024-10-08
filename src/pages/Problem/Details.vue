<!-- 题目详细页 -->
<script setup>
import { VAppBar, VBtn, VDivider, VCard, VCardText, VProgressCircular, VSelect, VContainer, VRow, VCol } from 'vuetify/components'
import { ref, onMounted, computed, watch } from 'vue'
import { useRoute } from 'vue-router';
import { getPbDetails, uploadCode } from '../../utils/axios';
import { VAceEditor } from 'vue3-ace-editor';
import { showAlert } from '../../utils/alert';
import { token, userName } from "../../utils/account";
import { MdPreview } from 'md-editor-v3';
import 'md-editor-v3/lib/preview.css';
import { useI18n } from 'vue-i18n';

const { t } = useI18n();

const networkloading = ref(false);
const id = 'preview-only';
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

// 代码上传
const uploadContentAsFile = async () => {
    const blob = new Blob([content.value], { type: 'text/plain' });
    const codefile = new File([blob], `main.${langFileExtension[lang.value]}`, { type: 'text/plain' });
    try {
        networkloading.value = true;
        const uploadResp = await uploadCode(codefile, route.params.Pid, userName.value, token.value);
        if (uploadResp.status === 200) {
            networkloading.value = false;
            showAlert(t("message.success") + "!", "reload")
        }
    } catch (error) {
        showAlert(t("message.failed") + "!", "");
    }
};

onMounted(async () => {
    loading.value = true;
    if (userLoggedIn.value) {
        try {
            const problemId = route.params.Pid;
            const resp = await getPbDetails(problemId, userName.value, token.value);
            if (resp.status === 200 && resp.data.problemInfo.pid !== 0) {
                problemInfo.value = resp.data.problemInfo;
            }else {
                window.location = '#/403'
            }
        } catch (error) {
            showAlert(t("message.failed") + "!", "");
        } finally {
            loading.value = false;
        }
    } else {
        window.location = '#/login'
    }

    watch(lang, (newLang) => {
        content.value = templates[newLang];
    });
});
</script>

<template>
    <v-dialog v-model="networkloading" max-width="600px">
        <v-card rounded="xl">
            <div class="networkloading">
                <v-progress-circular indeterminate color="primary" :width="12" :size="100"></v-progress-circular>
            </div>
        </v-card>
    </v-dialog>
    <v-app-bar :elevation="0">
        <template v-slot:prepend>
            <v-btn icon="mdi-chevron-left" size="x-large" @click="$router.back"></v-btn>
        </template>
    </v-app-bar>
    <div v-if="loading" class="loading">
        <v-progress-circular indeterminate color="primary" :width="12" :size="100"></v-progress-circular>
    </div>
    <div v-else>
        <v-container fluid>
            <v-row>
                <v-col cols="12" md="6">
                    <h1>{{ problemInfo.title }}</h1>
                    <div style="margin: 10px;"></div>
                    <p class="subtitle">{{ $t("message.difficulty") + ": " + problemInfo.difficulty }}</p>
                    <p class="subtitle">{{ $t("message.timeLimit") + ": " + problemInfo.time_limit }} S</p>
                    <p class="subtitle">{{ $t("message.memoryLimit") + ": " + problemInfo.memory_limit }} MB</p>
                    <div style="margin: 10px;"></div>
                    <v-divider></v-divider>
                    <div style="margin-top: 20px;"></div>
                    <md-preview :modelValue="problemInfo.content" :editorId="id" class="md_preview" />
                    <p class="tags">{{ $t("message.displayInputCase") }}</p>
                    <p class="example">{{ problemInfo.input }}</p>
                    <p class="tags">{{ $t("message.displayOutputCase") }}</p>
                    <p class="example">{{ problemInfo.output }}</p>
                </v-col>
                <v-divider vertical></v-divider>
                <v-col cols="12" md="6">
                    <v-card class="mx-auto my-4" width="100%" height="800" elevation="5">
                        <v-select :label="$t('message.lang')" v-model="lang"
                            :items="['python', 'c_cpp', 'golang', 'java']" variant="solo-filled"></v-select>
                        <v-ace-editor v-model:value="content" theme="chrome" :lang=lang
                            style="height: 800px;font-size: 14px;" />
                    </v-card>
                    <v-btn color="primary" rounded="xl" @click="uploadContentAsFile">{{ $t("message.submit") }}</v-btn>
                </v-col>
            </v-row>
        </v-container>
    </div>
</template>

<style scoped>
.networkloading {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100%;
    margin: 100px;
}

.tags {
    font-size: 18px;
    font-weight: bold;
    text-align: left;
    margin-top: 30px;
    margin-left: 30px;
}

.example {
    text-align: left;
    margin-left: 30px;
}

.md_preview {
    text-align: left;
    margin: 10px;
}

.subtitle {
    font-size: 0.9rem;
}

.loading {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100%;
}
</style>