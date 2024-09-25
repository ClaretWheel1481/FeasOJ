<script setup>
import { ref, onMounted, computed, reactive } from 'vue'
import { token, userName } from '../../utils/account'
import { useI18n } from 'vue-i18n';
import { getAllCompetitionsInfo, verifyUserInfo, getCompetitionInfoByIDAdmin, deleteCompetition, updateComInfo } from '../../utils/axios';
import { showAlert } from '../../utils/alert';
import moment from 'moment';

const { locale } = useI18n();
const { t } = useI18n();

// 计算属性来判断用户是否已经登录
const userLoggedIn = computed(() => !!token.value)
const loading = ref(true)
const networkloading = ref(false)
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
});

const headers = ref([
    { title: 'ID', value: 'contest_id', align: 'center' },
    { title: t('message.competition'), value: 'title', align: 'center' },
])

// 字段检查
const validateFields = () => {
    for (const key in competitionFields) {
        if (key === "password" && !competitionFields.have_password) {
            continue; // 跳过 password 检查
        }
        if (competitionFields[key] === "" || (Array.isArray(competitionFields[key]) && competitionFields[key].length === 0)) {
            showAlert(t("message.formCheckfailed") + "!", "");
            console.log(key + competitionFields[key]);
            return false;
        }
    }
    return true;
};

// 删除竞赛
const delCompetition = async () => {
    networkloading = true;
    const delResp = await deleteCompetition(competitionFields.contest_id, userName.value, token.value);
    networkloading = false;
    if (delResp.data.status === 200) {
        showAlert(t("message.success") + "!", "reload");
    } else {
        showAlert(t("message.failed") + "!", "");
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
}

// 数据传至后端
const save = async () => {
    competitionFields.start_at = new Date(formatStartDate.value).toISOString();
    competitionFields.end_at = new Date(formatEndDate.value).toISOString();
    if (!validateFields()) return;
    networkloading.value = true;
    const comData = { ...competitionFields };
    try {
        const updateResp = await updateComInfo(userName.value, token.value, comData);
        if (updateResp.data.status === 200) {
            networkloading.value = false;
            showAlert(t("message.success") + "!", "reload");
        }
    } catch (error) {
        showAlert(t("message.failed") + "!", "reload");
    }
    dialog.value = false;
};

// 编辑/创建竞赛
const goToEditCompetition = async (contest_id) => {
    isCreate.value = false;
    dialog.value = true;
    networkloading.value = true;
    const comResp = await getCompetitionInfoByIDAdmin(userName.value, token.value, contest_id);
    networkloading.value = false;

    Object.assign(competitionFields, comResp.data.contest);
    formatStartDate.value = moment(competitionFields.start_at).format('YYYY-MM-DD HH:mm');
    formatEndDate.value = moment(competitionFields.end_at).format('YYYY-MM-DD HH:mm');
}

// 从后端获取数据
const fetchData = async () => {
    loading.value = true
    try {
        const response = await getAllCompetitionsInfo(userName.value, token.value);
        competitions.value = response.data.contests
        totalCompetitions.value = competitions.value.length
        competitionFields.contest_id = totalCompetitions.value + 1
    } catch (error) {
        showAlert(t("message.failed") + "!", "")
    } finally {
        loading.value = false
    }
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
    <v-data-table-server :headers="headers" :items="competitions" :items-length="totalCompetitions" :loading="loading"
        :loading-text="$t('message.loading')" @update="fetchData" :hide-default-footer="true"
        :no-data-text="!userLoggedIn ? $t('message.nologin') : $t('message.nodata')">
        <template v-slot:item="{ item }">
            <tr>
                <td>{{ item.contest_id }}</td>
                <td class="tabletitle">
                    <v-btn @click="goToEditCompetition(item.contest_id)" variant="text" block>{{ item.title }}</v-btn>
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
                    </v-form>
                </v-card-text>
                <v-card-actions>
                    <v-spacer></v-spacer>
                    <!-- TODO: 部分方法待实现 -->
                    <v-btn color="blue darken-1" text @click="dialog = false" rounded="xl">{{ $t('message.cancel')
                        }}</v-btn>
                    <v-btn v-if="!isCreate" color="primary" @click="delCompetition" rounded="xl">{{ $t('message.delete')
                        }}</v-btn>
                    <v-btn color="primary" @click="save" rounded="xl">{{ $t('message.save') }}</v-btn>
                </v-card-actions>
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
</style>