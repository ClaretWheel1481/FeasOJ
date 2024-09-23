<script setup>
import { ref, onMounted, computed, reactive } from 'vue'
import { token, userName } from '../../utils/account'
import { useI18n } from 'vue-i18n';
import { getAllCompetitionsInfo, verifyUserInfo, getCompetitionInfoByIDAdmin } from '../../utils/axios';

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
        if (competitionFields[key] === "" || (Array.isArray(competitionFields[key]) && competitionFields[key].length === 0)) {
            showAlert(t("message.formCheckfailed") + "!", "");
            return false;
        }
    }
    for (const testCase of competitionFields.test_cases) {
        if (testCase.input === "" || testCase.output === "") {
            showAlert(t("message.formCheckfailed") + "!", "");
            return false;
        }
    }
    return true;
};

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

// 编辑/创建竞赛
const goToEditCompetition = async (contest_id) => {
    isCreate.value = false;
    dialog.value = true;
    networkloading.value = true;
    const comResp = await getCompetitionInfoByIDAdmin(userName.value, token.value, contest_id);
    networkloading.value = false;

    Object.assign(competitionFields, comResp.data.contest);
};

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
                        <!-- 难易程度 -->
                        <v-select :items="['简单', '中等', '困难']" :label="$t('message.difficulty')"
                            v-model="competitionFields.difficulty" variant="solo-filled"></v-select>
                        <!-- TODO：密码、密码是否可视化、是否可见 -->
                    </v-form>
                </v-card-text>
                <v-card-actions>
                    <v-spacer></v-spacer>
                    <!-- TODO: 部分方法待实现 -->
                    <v-btn color="blue darken-1" text @click="dialog = false" rounded="xl">{{ $t('message.cancel')
                        }}</v-btn>
                    <v-btn v-if="!isCreate" color="primary" @click="" rounded="xl">{{ $t('message.delete')
                        }}</v-btn>
                    <v-btn color="primary" @click="" rounded="xl">{{ $t('message.save') }}</v-btn>
                </v-card-actions>
            </div>
        </v-card>
    </v-dialog>
    <div class="fab">
        <v-fab fixed icon="mdi-plus" size="64" color="primary" elevation="10" v-if="!loading" @click="createCompetition"></v-fab>
    </div>
</template>

<style>
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