<script setup>
import { token, userName } from "../../utils/account";
import { ref, onMounted, computed } from "vue";
import { getAllCompetitions } from "../../utils/axios";
import { useI18n } from "vue-i18n";
import moment from "moment";

const { t } = useI18n();

const competitions = ref([]);
const loading = ref(false);
// 计算属性来判断用户是否已经登录
const userLoggedIn = computed(() => !!token.value);
const networkloading = ref(false);

// 根据题目难度显示不同字体
const difficultyColor = (difficulty) => {
    switch (difficulty) {
        case "简单":
            return "font-weight: bold;color: green;";
        case "中等":
            return "font-weight: bold;color: orange;";
        case "困难":
            return "font-weight: bold;color: red;";
        default:
            return "";
    }
};

onMounted(async () => {
    if (!userLoggedIn.value) {
        loading.value = false;
        setTimeout(() => {
            window.location = "#/login";
            window.location.reload();
        }, 2000);
        return;
    }
    loading.value = true;
    const response = await getAllCompetitions(userName.value, token.value);
    competitions.value = response.data.contests;
    loading.value = false;
});
</script>

<template>
    <template>
        <v-dialog v-model="networkloading" max-width="500px">
            <v-card rounded=xl>
                <div class="networkloading">
                    <v-progress-circular indeterminate color="primary" :width="12" :size="100"></v-progress-circular>
                </div>
            </v-card>
        </v-dialog>
    </template>
    <div v-if="loading" class="loading">
        <v-progress-circular indeterminate color="primary" :width="12" :size="100"></v-progress-circular>
    </div>
    <div v-else>
        <div class="title">
            <h1>{{ t("message.competition") }}</h1>
        </div>
        <div v-if="userLoggedIn">
            <div v-if="competitions.length === 0">
                <p>{{ t("message.nodata") }}</p>
            </div>
            <v-container>
                <v-row>
                    <v-col v-for="contest in competitions" :key="contest.contest_id" cols="12" md="4">
                        <v-card rounded="xl" elevation="8">
                            <v-card-title style="font-weight: bold;font-size:28px;justify-self: left;">{{ contest.title }}</v-card-title>
                            <v-card-subtitle style="justify-self: left;">{{ contest.subtitle }}</v-card-subtitle>
                            <template v-slot:append>
                                <v-chip :style="difficultyColor(contest.difficulty)">
                                    {{ contest.difficulty }}
                                </v-chip>
                                <br>
                            </template>
                            <v-card-text>
                                <p style="justify-self: left;">
                                    {{ moment(contest.start_at).format("MM/DD HH:mm") }} -
                                    {{ moment(contest.end_at).format("MM/DD HH:mm") }}
                                </p>
                            </v-card-text>
                            <template v-slot:actions>
                                <v-btn color="primary" append-icon="mdi-chevron-right" @click="">{{ $t("message.enter")
                                    }}</v-btn>
                            </template>
                        </v-card>
                    </v-col>
                </v-row>
            </v-container>
        </div>
        <div v-else>
            <div class="title" style="margin: 50px">
                <h1>{{ t("message.competition") }}</h1>
            </div>
            <p>{{ t("message.nologin") }}</p>
        </div>
    </div>
</template>

<style scoped>
.loading {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100%;
}
</style>
