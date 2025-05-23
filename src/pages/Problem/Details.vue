<!-- 题目详细页 -->
<script setup>
import { ref, onMounted, computed, watch } from 'vue';
import { useRoute } from 'vue-router';
import { getPbDetails, uploadCode } from '../../utils/api/problems';
import { VAceEditor } from 'vue3-ace-editor';
import { showAlert } from '../../utils/alert';
import { token } from "../../utils/account";
import { MdPreview } from 'md-editor-v3';
import { useI18n } from 'vue-i18n';
import { difficultyColor, difficultyLang } from '../../utils/dynamic_styles';
import 'md-editor-v3/lib/preview.css';

const { t } = useI18n();

const networkloading = ref(false);
const id = 'preview-only';
const route = useRoute();
const loading = ref(true);
const problemInfo = ref({});
const content = ref('');
const lang = ref('c_cpp'); // 初始选中语言

// Ace Editor字体
const fontSize = ref(14);
const minFontSize = 10;
const maxFontSize = 36;
const stepFontSize = 1;

// Ace Editor主题
const theme = ref("chrome")

// 支持的语言
const langFileExtension = {
    java: 'java',
    c_cpp: 'cpp',
    golang: 'go',
    python: 'py',
    rust: 'rs',
    php: 'php',
    pascal: 'pas'
};

// 计算属性来判断用户是否已经登录
const userLoggedIn = computed(() => !!token.value)

// 代码模板
const templates = {
    java:
        `
// Please don't edit the 'Main' class name. 请不要编辑 'Main' 类名。
public class Main {
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
        ``,
    rust: `

fn main() {
    
}`,
    php: `<?php

?>`,
    pascal: `{ O2 Enabled / O2已启用 }
begin
    
end.`
};

// 代码上传
const uploadContentAsFile = async () => {
    const blob = new Blob([content.value], { type: 'text/plain' });
    const codefile = new File([blob], `main.${langFileExtension[lang.value]}`, { type: 'text/plain' });
    try {
        networkloading.value = true;
        const resp = await uploadCode(codefile, route.params.Pid);
        networkloading.value = false;
        showAlert(resp.data.message, "reload")
    } catch (error) {
        networkloading.value = false;
        showAlert(error.response.data.message, "");
    }
};

onMounted(async () => {
    loading.value = true;
    if (userLoggedIn.value) {
        try {
            const problemId = route.params.Pid;
            const resp = await getPbDetails(problemId);
            problemInfo.value = resp.data.problemInfo;
        } catch (error) {
            showAlert(error.response.data.message, "");
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
    <div v-if="loading" class="loading">
        <v-progress-circular indeterminate color="primary" :width="12" :size="100"></v-progress-circular>
    </div>
    <div v-else>
        <v-app-bar :elevation="2">
            <template v-slot:prepend>
                <v-btn icon="mdi-chevron-left" size="x-large" @click="$router.back"></v-btn>
            </template>
            <v-col class="align-left">
                <v-row style="align-items: center;">
                    <p style="font-size: 24px;">{{ problemInfo.title }}</p>
                    <div style="margin-left: 10px;"></div>
                    <v-chip :style="difficultyColor(problemInfo.difficulty)">
                        {{ $t(difficultyLang(problemInfo.difficulty)) }}
                    </v-chip>
                </v-row>
            </v-col>
        </v-app-bar>
        <v-container fluid>
            <v-row>
                <v-col cols="12" md="6">
                    <div style="margin-bottom: 20px;"></div>
                    <p class="subtitle">{{ $t("message.timeLimit") + ": " + problemInfo.time_limit }} S</p>
                    <p class="subtitle">{{ $t("message.memoryLimit") + ": " + problemInfo.memory_limit }} MB</p>
                    <div style="margin-bottom: 20px;"></div>
                    <v-divider></v-divider>
                    <div style="margin-top: 20px;"></div>
                    <md-preview style="margin-left: 30px;" :modelValue="problemInfo.content" :editorId="id"
                        class="md_preview" />
                    <p class="tags">{{ $t("message.displayInputCase") }}</p>
                    <p class="example">{{ problemInfo.input }}</p>
                    <p class="tags">{{ $t("message.displayOutputCase") }}</p>
                    <p class="example">{{ problemInfo.output }}</p>
                </v-col>
                <v-divider vertical></v-divider>
                <v-col cols="12" md="6">
                    <v-alert 
                        :text="$t('message.a_code')"
                        :title="$t('message.attention')" align="left" type="info" variant="tonal" border="start" closable></v-alert>
                    <v-card width="100%" height="100vh" elevation="0">
                        <v-row style="margin-inline: 0;margin-top: 0;">
                            <v-select :label="$t('message.lang')" v-model="lang"
                                :items="['c_cpp', 'golang', 'java', 'pascal', 'python', 'php', 'rust']"
                                variant="outlined" class="mx-auto mt-4" elevation="0"></v-select>
                            <v-select :label="$t('message.editor_theme')" v-model="theme"
                                :items="['ambiance', 'clouds', 'cobalt', 'chaos', 'crimson_editor', 'dawn', 'dracula', 'dreamweaver', 'chrome', 'github', 'terminal', 'monokai', 'mono_industrial', 'pastel_on_dark', 'sqlserver', 'solarized_light', 'solarized_dark']"
                                variant="outlined" class="mx-auto mt-4" elevation="0"
                                style="padding-left: 5px;"></v-select>
                            <v-slider class="mt-4" v-model="fontSize" :min="minFontSize" :max="maxFontSize"
                                :step="stepFontSize" thumb-label prepend-icon="mdi-format-size"
                                elevation="2"></v-slider>
                        </v-row>

                        <v-ace-editor v-model:value="content" :theme="theme" :lang=lang
                            :style="`height:800px; font-size: ${fontSize}px;`" />
                    </v-card>
                    <v-btn color="primary" rounded="xl" @click="uploadContentAsFile">{{ $t("message.submit") }}</v-btn>
                </v-col>
            </v-row>
        </v-container>
    </div>
</template>

<style scoped>
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
    text-align: left;
    margin-left: 30px;
}

.loading {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100%;
}
</style>