<!-- 题目后台管理页 -->
<script setup>
import { ref, onMounted, computed } from 'vue'
import { token, userName } from '../utils/account'
import { getUserInfo, verifyJWT, getAllProblems, getProblemAllInfoByAdmin, updateProblemInfo } from '../utils/axios';
import { VDataTableServer, VFab, VDialog,VCard,VCardTitle,VCardText,VBtn,VTextField,VSelect,VForm,VTextarea,VSpacer,VCardActions,VRow } from 'vuetify/components'
import { showAlert } from '../utils/alert';

// 计算属性来判断用户是否已经登录
const userLoggedIn = computed(() => !!token.value)
const loading = ref(true)
const userPrivilege = ref("")
const problems = ref([])
const totalProblems = ref(0)
const dialog = ref(false)
const difficulty = ref("")
const testCases = ref([{ input: '', output: '' }]);
const title = ref("")
const content = ref("")
const timeLimit = ref("")
const memoryLimit = ref("")
const pids = ref()
const input = ref("")
const output = ref("")
const headers = ref([
    { title: 'ID', value: 'Pid', align: 'center' },
    { title: '题目', value: 'Title', align: 'center' },
])

// 添加测试样例
const addTestCase = () => {
  testCases.value.push({ input: '', output: '' });
};

// 数据传至后端
const save = async() => {
    const problemData = {
        pid: pids.value,
        title: title.value,
        content: content.value,
        difficulty: difficulty.value,
        time_limit: timeLimit.value,
        memory_limit: memoryLimit.value,
        input: input.value,
        output: output.value,
        test_cases: testCases.value
    };
    try{
        const updateResp = await updateProblemInfo(userName.value,token.value,problemData);
        if(updateResp.data.status === 200){
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
        pids.value = totalProblems.value+1
    } catch (error) {
        showAlert('无法获取数据，请重试。', "")
    } finally {
        loading.value = false
    }
}

// 编辑题目显示
const goToEditProblem = async(pid) => {
    dialog.value = true;
    const problemResp = await getProblemAllInfoByAdmin(pid,userName.value,token.value);
    pids.value = problemResp.data.problemInfo.pid;
    title.value = problemResp.data.problemInfo.title;
    content.value = problemResp.data.problemInfo.content;
    difficulty.value = problemResp.data.problemInfo.difficulty;
    timeLimit.value = problemResp.data.problemInfo.time_limit;
    memoryLimit.value = problemResp.data.problemInfo.memory_limit;
    input.value = problemResp.data.problemInfo.input;
    output.value = problemResp.data.problemInfo.output;
    testCases.value = problemResp.data.problemInfo.test_cases;
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
    <v-dialog v-model="dialog" max-width="800px">
        <v-card>
            <v-card-title>
            <span class="headline">编辑题目</span>
            </v-card-title>
            <v-card-text>
            <v-form>
                <v-text-field label="题目ID" v-model="pids" variant="solo-filled" readonly></v-text-field>
                <!-- 题目名称 -->
                <v-text-field label="题目名称" v-model="title" variant="solo-filled"></v-text-field>
                <!-- TODO: 题目描述Markdown编辑器待替换 -->
                <!-- 题目描述 -->
                <v-textarea v-model="content" label="题目描述" variant="solo-filled"></v-textarea>
                <!-- 难易程度 -->
                <v-select
                :items="['简单', '中等', '困难']"
                label="难易程度"
                v-model="difficulty"
                variant="solo-filled"
                ></v-select>
                <v-row class="limitRow">
                    <v-text-field label="时间限制s" v-model="timeLimit" variant="solo-filled"></v-text-field>
                    <div style="margin-inline: 30px;"></div>
                    <v-text-field label="内存限制MB" v-model="memoryLimit" variant="solo-filled"></v-text-field>
                </v-row>
                <v-row class="limitRow">
                    <v-text-field label="用户输入案例" v-model="input" variant="solo-filled"></v-text-field>
                    <div style="margin-inline: 30px;"></div>
                    <v-text-field label="用户输出案例" v-model="output" variant="solo-filled"></v-text-field>
                </v-row>
                <!-- 输入输出测试样例 -->
                <div v-for="(testCase, index) in testCases" :key="index">
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
            </v-form>
            </v-card-text>
            <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn color="blue darken-1" text @click="dialog = false" rounded="xl">取消</v-btn>
            <v-btn color="primary" @click="save" rounded="xl">保存</v-btn>
            </v-card-actions>
        </v-card>
    </v-dialog>
    <!-- TODO:添加题目待处理 -->
    <div class="fab">
        <v-fab fixed icon="mdi-plus" size="64" color="primary" elevation="10" v-if="!loading" @click=""></v-fab>
    </div>
</template>

<style scoped>
.loading {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100%;
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