<!-- 题目后台管理页 -->
<script setup>
import { ref, onMounted, computed } from 'vue'
import { token, userName } from '../utils/account'
import { getUserInfo, verifyJWT, getAllProblems } from '../utils/axios';
import { VDataTableServer, VFab }

// 计算属性来判断用户是否已经登录
const userLoggedIn = computed(() => !!token.value)
const loading = ref(true)
const userPrivilege = ref("")
const problems = ref([])
const totalProblems = ref(0)
const headers = ref([
    { title: 'ID', value: 'Pid', align: 'center' },
    { title: '题目', value: 'Title', align: 'center' },
])

// 从后端获取数据
const fetchData = async () => {
    loading.value = true
    try {
        const response = await getAllProblems();
        problems.value = response.data.problems
        totalProblems.value = problems.value.length
    } catch (error) {
        showAlert('无法获取数据，请重试。', "")
    } finally {
        loading.value = false
    }
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
                    <v-btn @click="goTOEditProblem(item.Pid)" variant="text" block>{{ item.Title }}</v-btn>
                </td>
            </tr>
        </template>
    </v-data-table-server>
    <!-- TODO:添加题目待处理 -->
    <div class="fab">
        <v-fab fixed icon="mdi-plus" size="64" color="primary" elevation="10" v-if="!loading"
            @click="$router.push('')"></v-fab>
    </div>
</template>

<style scoped>
.loading {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100%;
}

.fab {
    position: fixed;
    bottom: 80px;
    right: 180px;
}
</style>