<script setup>
import { ref, onMounted, computed, reactive, watch } from 'vue'
import { token, userName } from '../../utils/account'
import { useI18n } from 'vue-i18n';
import { getAllCompetitionsInfo, getCompetitionInfoByIDAdmin, deleteCompetition, updateComInfo, caculateComScore, getScores } from '../../utils/api/admin';
import { verifyUserInfo } from '../../utils/api/auth';
import { showAlert } from '../../utils/alert';
import { MdEditor } from "md-editor-v3";
import moment from 'moment';
import "md-editor-v3/lib/style.css";

const { locale } = useI18n();
const { t } = useI18n();

// 计算属性来判断用户是否已经登录
const userLoggedIn = computed(() => !!token.value)

const loading = ref(true)
const scoreloading = ref(false)
const networkloading = ref(false)
const delDialog = ref(false)
const scoreDialog = ref(false)

const id = "preview-only";

const scores = ref([])
const totalScores = ref(0)
const itemsPerPage = ref(10)
const page = ref(1)

const cid = ref()
const userPrivilege = ref("")
const competitions = ref([])
const totalCompetitions = ref(0)
const dialog = ref(false)
const isCreate = ref(false)
const showPwd = ref(false)
const formatStartDate = ref("")
const formatEndDate = ref("")
const competitionFields = reactive({
    contest_id: null,
    title: "",
    subtitle: "",
    difficulty: "",
    password: "",
    have_password: false,
    is_visible: true,
    start_at: "",
    end_at: "",
    announcement: ""
});

const headers = ref([
    { title: 'ID', value: 'contest_id', align: 'center' },
    { title: t('message.competition'), value: 'title', align: 'center' },
    { title: t("message.operation"), value: 'actions', align: 'center' }
])
const scoreHeaders = ref([
    { title: 'ID', value: 'Uid', align: 'center' },
    { title: t('message.username'), value: 'Username', align: 'center' },
    { title: 'Score', value: 'Score', align: 'center' },
])

// 字段检查
const validateFields = () => {
    for (const key in competitionFields) {
        if (key === "password" && !competitionFields.have_password || key === "announcement") {
            continue;
        }
        if (competitionFields[key] === "" || (Array.isArray(competitionFields[key]) && competitionFields[key].length === 0)) {
            showAlert(t("message.formCheckfailed") + "!", "");
            return false;
        }
    }
    return true;
};

// 删除竞赛
const delCompetition = async () => {
    networkloading.value = true;
    try {
        await deleteCompetition(competitionFields.contest_id);
        showAlert(t("message.success") + "!", "reload");
    } catch (error) {
        showAlert(t("message.failed") + "!", "");
    } finally {
        networkloading.value = false;
    }
    dialog.value = false;
}

// 清除密码
const clearPassword = () => {
    if (!competitionFields.have_password) {
        competitionFields.password = '';
    }
}

// 创建竞赛
const createCompetition = async () => {
    isCreate.value = true;
    dialog.value = true;
    // 重置内容
    competitionFields.contest_id = totalCompetitions.value + 1;
    competitionFields.title = "";
    competitionFields.subtitle = "";
    competitionFields.difficulty = "简单";
    competitionFields.password = "";
    competitionFields.have_password = false;
    competitionFields.is_visible = true;
    competitionFields.start_at = "";
    competitionFields.end_at = "";
    competitionFields.announcement = "";
    formatStartDate.value = "";
    formatEndDate.value = "";
}

// 数据传至后端
const save = async () => {
    try {
        competitionFields.start_at = new Date(formatStartDate.value).toISOString();
        competitionFields.end_at = new Date(formatEndDate.value).toISOString();
    } catch (error) {
        showAlert(t('message.formRuleCheckfailed') + "!", "");
        return;
    }
    if (!validateFields()) return;
    networkloading.value = true;
    const comData = { ...competitionFields };
    try {
        await updateComInfo(comData);
        showAlert(t("message.success") + "!", "reload");
    } catch (error) {
        showAlert(t("message.failed") + "!", "reload");
    } finally{
        networkloading.value = false;
    }
    dialog.value = false;
};

// 编辑/创建竞赛
const goToEditCompetition = async (contest_id) => {
    isCreate.value = false;
    dialog.value = true;
    networkloading.value = true;
    const comResp = await getCompetitionInfoByIDAdmin(contest_id);
    networkloading.value = false;

    Object.assign(competitionFields, comResp.data.contest);
    formatStartDate.value = moment(competitionFields.start_at).format('YYYY-MM-DD HH:mm');
    formatEndDate.value = moment(competitionFields.end_at).format('YYYY-MM-DD HH:mm');
}

// 从后端获取数据
const fetchData = async () => {
    loading.value = true
    try {
        const response = await getAllCompetitionsInfo();
        competitions.value = response.data.contests
        totalCompetitions.value = competitions.value.length
        competitionFields.contest_id = totalCompetitions.value + 1
    } catch (error) {
        loading.value = false
        showAlert(t("message.failed") + "!", "")
    } finally {
        loading.value = false
    }
}

// 竞赛计分
const calcCompetition = async (competitionId) => {
    try {
        const response = await caculateComScore(competitionId);
        showAlert(t("message.calculating"), "");
    } catch (error) {
        showAlert(t("message.failed") + "!", "");
    }
}

// 查看竞赛得分情况
const getScoreBoard = async (competitionId) => {
    scoreloading.value = true;
    scoreDialog.value = true;
    cid.value = competitionId;
    try {
        const response = await getScores(competitionId, page.value, itemsPerPage.value);
        scores.value = response.data.users;
        totalScores.value = response.data.total;
    } catch (error) {
        showAlert(t("message.failed") + "!", "");
    } finally {
        scoreloading.value = false;
    }
}

// 监听分页变化
watch(page, async () => {
    await getScoreBoard(cid.value)
})

onMounted(async () => {
    loading.value = true;
    try {
        if (!userLoggedIn.value) {
            window.location = "#/login";
            return;
        }
        const userInfoResponse = await verifyUserInfo(userName.value, token.value);
        userPrivilege.value = userInfoResponse.data.info.role;
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
    <template>
        <v-dialog v-model="scoreDialog" persistent max-width="1200px">
            <v-card rounded="xl">
                <v-data-table-server :headers="scoreHeaders" :items="scores" :items-length="totalScores"
                    :loading="scoreloading" :loading-text="$t('message.loading')" :hide-default-footer="true"
                    :items-per-page="itemsPerPage" v-model:page="page" @update="getScoreBoard(cid.value)">
                    <template v-slot:bottom>
                        <v-pagination v-model="page" :length="Math.ceil(totalScores / itemsPerPage)" rounded="xl" />
                    </template>
                </v-data-table-server>
                <v-card-actions>
                    <v-btn color="primary" @click="scoreDialog = false" rounded="xl">{{ $t('message.ok') }}</v-btn>
                </v-card-actions>
            </v-card>
        </v-dialog>
    </template>
    <template>
        <v-dialog v-model="delDialog" persistent max-width="290">
            <v-card rounded="xl">
                <v-card-title class="text-h5">{{ $t('message.notify') }}</v-card-title>
                <v-card-text>{{ t('message.suredel') }}</v-card-text>
                <v-card-actions>
                    <v-btn variant="elevated" color="primary" @click="delCompetition" rounded="xl">{{ $t('message.yes')
                        }}</v-btn>
                    <v-btn color="primary" @click="delDialog = false" rounded="xl">{{ $t('message.cancel') }}</v-btn>
                </v-card-actions>
            </v-card>
        </v-dialog>
    </template>
    <v-app-bar :elevation="2">
        <template v-slot:prepend>
            <v-btn icon="mdi-chevron-left" size="x-large" @click="$router.back"></v-btn>
        </template>
        <v-col class="align-left">
            <p style="font-size: 24px;">{{ t('message.competitionmanagement') }}</p>
        </v-col>
    </v-app-bar>
    <v-data-table-server :headers="headers" :items="competitions" :items-length="totalCompetitions" :loading="loading"
        :loading-text="$t('message.loading')" @update="fetchData" :hide-default-footer="true"
        :no-data-text="!userLoggedIn ? $t('message.nologin') : $t('message.nodata')">
        <template v-slot:item="{ item }">
            <tr>
                <td>{{ item.contest_id }}</td>
                <td class="tabletitle">
                    <v-btn @click="goToEditCompetition(item.contest_id)" variant="text" block>{{ item.title }}</v-btn>
                </td>
                <td>
                    <v-menu>
                        <template v-slot:activator="{ props }">
                            <v-btn v-bind="props" variant="text" icon="mdi-dots-horizontal"></v-btn>
                        </template>
                        <v-list rounded="xl">
                            <v-list-item :disabled="(item.status != 2) || (item.scored != false)"
                                @click="calcCompetition(item.contest_id)">
                                <template v-slot:default="{ active, toggle }">
                                    <div class="d-flex align-center">
                                        <v-icon icon="mdi-calculator" class="me-2"></v-icon>
                                        <v-list-item-title>{{ t('message.scoring') }}</v-list-item-title>
                                    </div>
                                </template>
                            </v-list-item>
                            <v-list-item @click="getScoreBoard(item.contest_id)" :disabled="item.scored != true">
                                <template v-slot:default="{ active, toggle }">
                                    <div class="d-flex align-center">
                                        <v-icon icon="mdi-note-text" class="me-2"></v-icon>
                                        <v-list-item-title>{{ t("message.viewScores") }}</v-list-item-title>
                                    </div>
                                </template>
                            </v-list-item>
                        </v-list>
                    </v-menu>
                </td>
            </tr>
        </template>
    </v-data-table-server>
    <v-dialog v-model="dialog" max-width="1200px">
        <v-card>
            <div v-if="networkloading" class="networkloading">
                <v-progress-circular indeterminate color="primary" :width="12" :size="100"></v-progress-circular>
            </div>
            <div v-else>
                <v-card-title>
                    <span class="headline">{{ $t('message.edit') }}</span>
                </v-card-title>
                <v-card-text>
                    <v-form>
                        <v-text-field label="ID" v-model="competitionFields.contest_id" variant="solo-filled"
                            readonly></v-text-field>
                        <!-- 竞赛名称 -->
                        <v-text-field :label="$t('message.title')" v-model="competitionFields.title"
                            variant="solo-filled"></v-text-field>
                        <!-- 竞赛副标题 -->
                        <v-text-field :label="$t('message.subtitle')" v-model="competitionFields.subtitle"
                            variant="solo-filled"></v-text-field>
                        <v-row class="limitRow">
                            <!-- 难易程度 -->
                            <v-select :items="['简单', '中等', '困难']" :label="$t('message.difficulty')"
                                v-model="competitionFields.difficulty" variant="solo-filled"></v-select>
                            <div style="margin-inline: 30px;"></div>
                            <!-- 是否可见 -->
                            <v-switch v-model="competitionFields.is_visible" :label="$t('message.isvisible')"
                                color="primary" inset></v-switch>
                            <div style="margin-inline: 30px;"></div>
                            <!-- 启用密码 -->
                            <v-switch v-model="competitionFields.have_password" :label="$t('message.withapwd')"
                                color="primary" inset @change="clearPassword"></v-switch>
                        </v-row>
                        <v-text-field v-if="competitionFields.have_password" v-model="competitionFields.password"
                            :append-icon="showPwd ? 'mdi-eye' : 'mdi-eye-off'" variant="solo-filled"
                            :type="showPwd ? 'text' : 'password'" :label="$t('message.password')" counter
                            @click:append="showPwd = !showPwd"></v-text-field>
                        <!-- 竞赛开始时间日期 -->
                        <v-row class="limitRow">
                            <v-text-field :label="$t('message.start_date') + '(YYYY-MM-DD HH:mm)'"
                                v-model="formatStartDate" variant="solo-filled"></v-text-field>
                            <div style="margin-inline: 30px;"></div>
                            <v-text-field :label="$t('message.end_date') + '(YYYY-MM-DD HH:mm)'" v-model="formatEndDate"
                                variant="solo-filled"></v-text-field>
                        </v-row>
                        <md-editor v-model="competitionFields.announcement" :editorId="id" :noUploadImg="true"
                            :footers="[]" :language="locale === 'zh_CN' ? 'zh-CN' : 'en-US'" />
                    </v-form>
                    <div style="position: fixed; bottom: 16px; right: 16px; z-index: 1000;">
                        <v-btn @click="dialog = false" rounded="xl" style="margin-right: 10px;">{{
                            $t('message.cancel') }}</v-btn>
                        <v-btn v-if="!isCreate" color="red-lighten-1" @click="delDialog = true" rounded="xl"
                            style="margin-right: 10px;">{{ $t('message.delete')
                            }}</v-btn>
                        <v-btn color="primary" @click="save" rounded="xl" style="margin-right: 10px;">{{
                            $t('message.save')
                        }}</v-btn>
                    </div>
                </v-card-text>
            </div>
        </v-card>
    </v-dialog>
    <div class="fab">
        <v-fab fixed icon="mdi-plus" size="64" color="primary" elevation="10" v-if="!loading"
            @click="createCompetition"></v-fab>
    </div>
</template>

<style scoped>
.limitRow {
    display: flex;
    justify-content: space-between;
    margin: 5px 0;
}

.tabletitle {
    color: #1e65ff;
}

.fab {
    position: fixed;
    bottom: 80px;
    right: 180px;
}
</style>