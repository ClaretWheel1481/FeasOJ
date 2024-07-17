<!-- 题目后台管理页 -->
<script setup>
import { ref, onMounted, computed } from 'vue'
import { token, userName } from '../utils/account'
import { getUserInfo, verifyJWT, getAllProblems, getProblemAllInfoByAdmin, updateProblemInfo, deleteProblemAllInfo } from '../utils/axios';
import { VDataTableServer, VFab, VDialog,VCard,VCardTitle,VCardText,VBtn,VTextField,VSelect,VForm,VSpacer,VCardActions,VRow } from 'vuetify/components'
import { showAlert } from '../utils/alert';
import { MdEditor } from 'md-editor-v3';
import 'md-editor-v3/lib/style.css';

// 计算属性来判断用户是否已经登录
const userLoggedIn = computed(() => !!token.value)
const loading = ref(true)
const networkloading = ref(false)
const userPrivilege = ref("")
const problems = ref([])
const totalProblems = ref(0)
const dialog = ref(false)
const problemFields = {
    pids: ref(),
    title: ref(""),
    content: ref(""),
    difficulty: ref(""),
    timeLimit: ref(""),
    memoryLimit: ref(""),
    input: ref(""),
    output: ref(""),
    testCases: ref([{ input: '', output: '' }])
};
const headers = ref([
    { title: 'ID', value: 'Pid', align: 'center' },
    { title: '题目', value: 'Title', align: 'center' },
])
const isCreate = ref(false)

// 添加测试样例
const addTestCase = () => {
    problemFields.testCases.value.push({ input: '', output: '' });
};

const removeTestCase = () => {
  if (problemFields.testCases.value.length > 0) {
    problemFields.testCases.value.pop();
  }
};

// 数据传至后端
const save = async() => {
    // TODO: 快来个大神帮我优化掉大芬循环
    for (const key in problemFields) {
        if (problemFields[key].value === "" || (Array.isArray(problemFields[key].value) && problemFields[key].value.length === 0)) {
            showAlert("请填写所有数据。", "");
            return;
        }
    }

    for (const testCase of problemFields.testCases.value) {
        if (testCase.input === "" || testCase.output === "") {
            showAlert("请填写所有测试样例的输入和输出。", "");
            return;
        }
    }
    networkloading.value = true;
    const problemData = {
        pid: problemFields.pids.value,
        title: problemFields.title.value,
        content: problemFields.content.value,
        difficulty: problemFields.difficulty.value,
        time_limit: problemFields.timeLimit.value,
        memory_limit: problemFields.memoryLimit.value,
        input: problemFields.input.value,
        output: problemFields.output.value,
        test_cases: problemFields.testCases.value
    };
    try{
        const updateResp = await updateProblemInfo(userName.value,token.value,problemData);
        if(updateResp.data.status === 200){
            networkloading.value = false;
            showAlert("更新成功！","reload");
        }
    }catch(error){
        showAlert("更新失败，请重试。","reload");
    }
    dialog.value = false;
};

// 从后端获取数据
const fetchData = async () => {
    loading.value = true
    try {
        const response = await getAllProblems();
        problems.value = response.data.problems
        totalProblems.value = problems.value.length
        problemFields.pids.value = totalProblems.value+1
    } catch (error) {
        showAlert('无法获取数据，请重试。', "")
    } finally {
        loading.value = false
    }
}

// 创建题目
const createProblem = async() => {
    isCreate.value = true;
    dialog.value = true;
    problemFields.pids.value = totalProblems.value+1;
    problemFields.title.value = "";
    problemFields.content.value = "";
    problemFields.difficulty.value = "";
    problemFields.timeLimit.value = "";
    problemFields.memoryLimit.value = "";
    problemFields.input.value = "";
    problemFields.output.value = "";
    problemFields.testCases.value = [{ input: '', output: '' }];
}

// 编辑题目显示
const goToEditProblem = async(pid) => {
    isCreate.value = false;
    dialog.value = true;
    networkloading.value = true;
    const problemResp = await getProblemAllInfoByAdmin(pid,userName.value,token.value);
    networkloading.value = false;
    problemFields.pids.value = problemResp.data.problemInfo.pid;
    problemFields.title.value = problemResp.data.problemInfo.title;
    problemFields.content.value = problemResp.data.problemInfo.content;
    problemFields.difficulty.value = problemResp.data.problemInfo.difficulty;
    problemFields.timeLimit.value = problemResp.data.problemInfo.time_limit;
    problemFields.memoryLimit.value = problemResp.data.problemInfo.memory_limit;
    problemFields.input.value = problemResp.data.problemInfo.input;
    problemFields.output.value = problemResp.data.problemInfo.output;
    problemFields.testCases.value = problemResp.data.problemInfo.test_cases;
}

// 删除题目
const delProblem = async() => {
    const delProblemResp = await deleteProblemAllInfo(problemFields.pids.value,userName.value,token.value);
    if(delProblemResp.data.status === 200){
        showAlert("删除成功！","reload");
    }else{
        showAlert("删除失败，请重试。","");
    }
    dialog.value = false;
}

onMounted(async () => {
    loading.value = true;
    try {
        if (!userLoggedIn.value) {
            window.location = "/login";
            return;
        }
        const userInfoResponse = await getUserInfo(userName.value);
        userPrivilege.value = userInfoResponse.data.Info.role;
        if (userPrivilege.value !== 1) {
            window.location = '/403';
            return;
        }
        const tokenVerificationResponse = await verifyJWT(userName.value, token.value);
        if (tokenVerificationResponse.data.status !== 200) {
            window.location = '/403';
            return;
        }
        await fetchData();
        loading.value = false;
    } catch (error) {
        window.location = '/403';
    }
});
</script>

<template>
    <v-data-table-server :headers="headers" :items="problems" :items-length="totalProblems" :loading="loading"
        loading-text="加载中..." @update="fetchData" :hide-default-footer="true"
        :no-data-text="!userLoggedIn ? '' : '没有题目数据。'">
        <template v-slot:item="{ item }">
            <tr>
                <td>{{ item.Pid }}</td>
                <td class="tabletitle">
                    <v-btn @click="goToEditProblem(item.Pid)" variant="text" block>{{ item.Title }}</v-btn>
                </td>
            </tr>
        </template>
    </v-data-table-server>
    <v-dialog v-model="dialog" max-width="1200px">
        <v-card>
            <div v-if="networkloading" class="loading">
                <v-progress-circular indeterminate color="primary" :width="12" :size="100"></v-progress-circular>
            </div>
            <div v-else>
                <v-card-title>
                    <span class="headline">编辑题目</span>
                </v-card-title>
                <v-card-text>
                <v-form>
                    <v-text-field label="题目ID" v-model="problemFields.pids.value" variant="solo-filled" readonly></v-text-field>
                    <!-- 题目名称 -->
                    <v-text-field label="题目名称" v-model="problemFields.title.value" variant="solo-filled"></v-text-field>
                    <!-- 题目描述 -->
                    <md-editor v-model="problemFields.content.value" :noUploadImg="true" :tabWidth="4"/>
                    <div style="margin-top: 20px;"></div>
                    <!-- 难易程度 -->
                    <v-select
                    :items="['简单', '中等', '困难']"
                    label="难易程度"
                    v-model="problemFields.difficulty.value"
                    variant="solo-filled"
                    ></v-select>
                    <v-row class="limitRow">
                        <v-text-field label="时间限制s" v-model="problemFields.timeLimit.value" variant="solo-filled"></v-text-field>
                        <div style="margin-inline: 30px;"></div>
                        <v-text-field label="内存限制MB" v-model="problemFields.memoryLimit.value" variant="solo-filled"></v-text-field>
                    </v-row>
                    <v-row class="limitRow">
                        <v-text-field label="显示输入案例" v-model="problemFields.input.value" variant="solo-filled"></v-text-field>
                        <div style="margin-inline: 30px;"></div>
                        <v-text-field label="显示输出案例" v-model="problemFields.output.value" variant="solo-filled"></v-text-field>
                    </v-row>
                    <!-- 输入输出测试样例 -->
                    <div v-for="(testCase, index) in problemFields.testCases.value" :key="index">
                        <span>测试样例 {{index+1}}</span>
                        <v-text-field
                            label="输入"
                            v-model="testCase.input"
                            variant="solo-filled"
                        ></v-text-field>
                        <v-text-field
                            label="输出"
                            v-model="testCase.output"
                            variant="solo-filled"
                        ></v-text-field>
                    </div>
                    <!-- 添加新的输入输出测试样例 -->
                    <v-btn @click="addTestCase" color="primary" rounded="xl">添加测试样例</v-btn>
                    <v-btn @click="removeTestCase" text rounded="xl">移除测试样例</v-btn>
                </v-form>
                </v-card-text>
                <v-card-actions>
                    <v-spacer></v-spacer>
                    <v-btn color="blue darken-1" text @click="dialog = false" rounded="xl">取消</v-btn>
                    <v-btn v-if="!isCreate" color="primary" @click="delProblem" rounded="xl">删除</v-btn>
                    <v-btn color="primary" @click="save" rounded="xl">保存</v-btn>
                </v-card-actions>
            </div>
        </v-card>
    </v-dialog>
    <div class="fab">
        <v-fab fixed icon="mdi-plus" size="64" color="primary" elevation="10" v-if="!loading" @click="createProblem"></v-fab>
    </div>
</template>

<style scoped>
.loading {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100%;
    margin: 100px;
}

.tabletitle{
  color: #1e65ff;
}

.fab {
    position: fixed;
    bottom: 80px;
    right: 180px;
}

.limitRow {
    display: flex;
    justify-content: space-between;
    margin: 5px 0;
}
</style>