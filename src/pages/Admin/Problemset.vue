<!-- 题目后台管理页 -->
<script setup>
import { ref, onMounted, computed, reactive } from 'vue'
import { token, userName } from '../../utils/account'
import { verifyUserInfo, getAllProblems, getProblemAllInfoByAdmin, updateProblemInfo, deleteProblemAllInfo, getAllCompetitionsInfo, getAllProblemsAdmin } from '../../utils/axios';
import { VDataTableServer, VFab, VDialog, VCard, VCardTitle, VCardText, VBtn, VTextField, VSelect, VForm, VSpacer, VCardActions, VRow } from 'vuetify/components'
import { showAlert } from '../../utils/alert';
import { MdEditor } from 'md-editor-v3';
import 'md-editor-v3/lib/style.css';
import { useI18n } from 'vue-i18n';

const { locale } = useI18n();
const { t } = useI18n();

// 计算属性来判断用户是否已经登录
const userLoggedIn = computed(() => !!token.value)
const loading = ref(true)
const networkloading = ref(false)
const userPrivilege = ref("")
const problems = ref([])
const totalProblems = ref(0)
const competitionIds = ref([0])
const dialog = ref(false)
const problemFields = reactive({
    pid: null,
    title: "",
    content: "",
    difficulty: "",
    time_limit: "",
    memory_limit: "",
    input: "",
    output: "",
    contest_id: null,
    is_visible: true,
    test_cases: [{ input: '', output: '' }]
});
const headers = ref([
    { title: 'ID', value: 'Pid', align: 'center' },
    { title: t('message.problem'), value: 'Title', align: 'center' },
])
const isCreate = ref(false)

// 添加测试样例
const addTestCase = () => {
    problemFields.test_cases.push({ input: '', output: '' });
};

// 删除测试样例
const removeTestCase = () => {
    if (problemFields.test_cases.length > 0) {
        problemFields.test_cases.pop();
    }
};

// 字段检查
const validateFields = () => {
    for (const key in problemFields) {
        if (problemFields[key] === "" || (Array.isArray(problemFields[key]) && problemFields[key].length === 0)) {
            showAlert(t("message.formCheckfailed") + "!", "");
            return false;
        }
    }
    for (const testCase of problemFields.test_cases) {
        if (testCase.input === "" || testCase.output === "") {
            showAlert(t("message.formCheckfailed") + "!", "");
            return false;
        }
    }
    return true;
};

// 数据传至后端
const save = async () => {
    if (!validateFields()) return;
    networkloading.value = true;
    const problemData = { ...problemFields };
    try {
        const updateResp = await updateProblemInfo(userName.value, token.value, problemData);
        if (updateResp.data.status === 200) {
            networkloading.value = false;
            showAlert(t("message.success") + "!", "reload");
        }
    } catch (error) {
        showAlert(t("message.failed") + "!", "reload");
    }
    dialog.value = false;
};

// 从后端获取数据
const fetchData = async () => {
    loading.value = true
    try {
        const response = await getAllProblemsAdmin(userName.value, token.value);
        problems.value = response.data.problems
        totalProblems.value = problems.value.length
        problemFields.pid = totalProblems.value + 1
    } catch (error) {
        showAlert(t("message.failed") + "!", "")
    } finally {
        loading.value = false
    }
}

// 创建题目
const createProblem = async () => {
    networkloading.value = true;
    isCreate.value = true;
    dialog.value = true;
    // 重置内容
    problemFields.pid = totalProblems.value + 1;
    problemFields.title = "";
    problemFields.content = "";
    problemFields.difficulty = "简单";
    problemFields.time_limit = "";
    problemFields.memory_limit = "";
    problemFields.input = "";
    problemFields.output = "";
    problemFields.contest_id = 0;
    problemFields.is_visible = true;
    problemFields.test_cases = [{ input: '', output: '' }];
    const compResp = await getAllCompetitionsInfo(userName.value, token.value);
    competitionIds.value = [0, ...compResp.data.contests.map(contest => contest.contest_id)];
    networkloading.value = false;
}

// 编辑题目显示
const goToEditProblem = async (pid) => {
    isCreate.value = false;
    dialog.value = true;
    networkloading.value = true;
    const problemResp = await getProblemAllInfoByAdmin(pid, userName.value, token.value);
    Object.assign(problemFields, problemResp.data.problemInfo);
    const compResp = await getAllCompetitionsInfo(userName.value, token.value);
    competitionIds.value = [0, ...compResp.data.contests.map(contest => contest.contest_id)];
    networkloading.value = false;
};

// 删除题目
const delProblem = async () => {
    const delProblemResp = await deleteProblemAllInfo(problemFields.pid, userName.value, token.value);
    if (delProblemResp.data.status === 200) {
        showAlert(t("message.success") + "!", "reload");
    } else {
        showAlert(t("message.failed") + "!", "");
    }
    dialog.value = false;
}

onMounted(async () => {
    loading.value = true;
    try {
        if (!userLoggedIn.value) {
            window.location = "#/login";
            return;
        }
        const userInfoResponse = await verifyUserInfo(userName.value, token.value);
        if (userInfoResponse.data.status !== 200) {
            window.location = '#/403';
            return;
        }
        userPrivilege.value = userInfoResponse.data.Info.role;
        if (userPrivilege.value !== 1) {
            window.location = '#/403';
            return;
        }
        await fetchData();
        loading.value = false;
    } catch (error) {
        window.location = '#/403';
    }
});
</script>

<template>
    <v-app-bar :elevation="0">
        <template v-slot:prepend>
            <v-btn icon="mdi-chevron-left" size="x-large" @click="$router.back"></v-btn>
        </template>
    </v-app-bar>
    <v-data-table-server :headers="headers" :items="problems" :items-length="totalProblems" :loading="loading"
        :loading-text="$t('message.loading')" @update="fetchData" :hide-default-footer="true"
        :no-data-text="!userLoggedIn ? $t('message.nologin') : $t('message.nodata')">
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
                    <span class="headline">{{ $t('message.edit') }}</span>
                </v-card-title>
                <v-card-text>
                    <v-form>
                        <v-text-field label="ID" v-model="problemFields.pid" variant="solo-filled"
                            readonly></v-text-field>
                        <!-- 题目名称 -->
                        <v-text-field :label="$t('message.title')" v-model="problemFields.title"
                            variant="solo-filled"></v-text-field>
                        <!-- 题目描述 -->
                        <md-editor v-model="problemFields.content" :footers="[]" :noUploadImg="true" :tabWidth="4"
                            :language="locale === 'zh_CN' ? 'zh-CN' : 'en-US'" />
                        <div style="margin-top: 20px;"></div>
                        <!-- 难易程度 -->
                        <v-select :items="['简单', '中等', '困难']" :label="$t('message.difficulty')"
                            v-model="problemFields.difficulty" variant="solo-filled"></v-select>
                        <!-- 所属竞赛ID及是否可见 -->
                        <v-row class="limitRow">
                            <v-select :items="competitionIds" :label="$t('message.contestid')" v-model.number="problemFields.contest_id" variant="solo-filled"></v-select>
                            <div style="margin-inline: 30px;"></div>
                            <v-switch v-model="problemFields.is_visible" :label="$t('message.isvisible')"
                                color="primary" inset></v-switch>
                        </v-row>
                        <!-- 时间、内存限制 -->
                        <v-row class="limitRow">
                            <v-text-field :label="$t('message.timeLimit') + '(S)'" v-model="problemFields.time_limit"
                                variant="solo-filled"></v-text-field>
                            <div style="margin-inline: 30px;"></div>
                            <v-text-field :label="$t('message.memoryLimit') + '(MB)'"
                                v-model="problemFields.memory_limit" variant="solo-filled"></v-text-field>
                        </v-row>
                        <!-- 显示输入输出样例 -->
                        <v-row class="limitRow">
                            <v-text-field :label="$t('message.displayInputCase')" v-model="problemFields.input"
                                variant="solo-filled"></v-text-field>
                            <div style="margin-inline: 30px;"></div>
                            <v-text-field :label="$t('message.displayOutputCase')" v-model="problemFields.output"
                                variant="solo-filled"></v-text-field>
                        </v-row>
                        <!-- 输入输出测试样例 -->
                        <div v-for="(testCase, index) in problemFields.test_cases" :key="index">
                            <span>{{ $t('message.testcases') + (index + 1) }}</span>
                            <v-text-field :label="$t('message.input')" v-model="testCase.input"
                                variant="solo-filled"></v-text-field>
                            <v-text-field :label="$t('message.output')" v-model="testCase.output"
                                variant="solo-filled"></v-text-field>
                        </div>
                        <!-- 添加新的输入输出测试样例 -->
                        <v-btn @click="addTestCase" color="primary" rounded="xl">{{ $t('message.addTestCase') }}</v-btn>
                        <v-btn @click="removeTestCase" text rounded="xl">{{ $t('message.deleteTestCase') }}</v-btn>
                    </v-form>
                </v-card-text>
                <v-card-actions>
                    <v-spacer></v-spacer>
                    <v-btn color="blue darken-1" text @click="dialog = false" rounded="xl">{{ $t('message.cancel')
                        }}</v-btn>
                    <v-btn v-if="!isCreate" color="primary" @click="delProblem" rounded="xl">{{ $t('message.delete')
                        }}</v-btn>
                    <v-btn color="primary" @click="save" rounded="xl">{{ $t('message.save') }}</v-btn>
                </v-card-actions>
            </div>
        </v-card>
    </v-dialog>
    <div class="fab">
        <v-fab fixed icon="mdi-plus" size="64" color="primary" elevation="10" v-if="!loading"
            @click="createProblem"></v-fab>
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

.tabletitle {
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