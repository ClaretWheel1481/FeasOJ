<script setup>
import { onMounted, ref, computed } from 'vue';
import { getAllUsersInfo, getUserInfo, verifyJWT } from '../utils/axios';
import { showAlert } from '../utils/alert';
import { token, userName } from '../utils/account'
import { VBtn } from 'vuetify/components';
import moment from 'moment';

// 计算属性来判断用户是否已经登录
const userLoggedIn = computed(() => !!token.value)
const loading = ref(true)
const userPrivilege = ref("")
const users = ref([])
const totalUsers = ref(0)
const searchQuery = ref('')
const headers = ref([
    { title: 'UID', value: 'uid', align: 'center' },
    { title: '用户名', value: 'username', align: 'center' },
    { title: '邮箱',value: 'email', align: 'center' },
    { title: '角色', value: 'role', align: 'center' },
    { title: '注册时间', value: 'registerTime', align: 'center' },
    { title: '状态', value: 'status', align: 'center' },
    { title: '操作', value: 'actions', align: 'center' },
])

// 通过用户名搜索
const filteredUsers = computed(() => {
    return users.value.filter((user) =>
        user.username.toLowerCase().includes(searchQuery.value.toLowerCase())
    )
})

// 获取数据
const fetchData = async () => {
    loading.value = true;
    try {
        const usersInfoResp = await getAllUsersInfo(userName.value, token.value)
        users.value = usersInfoResp.data.usersInfo;
        totalUsers.value = usersInfoResp.data.usersInfo.length;
        loading.value = false;
    } catch (error) {
        showAlert("获取用户信息失败，请稍后重试。");
        loading.value = false;
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
    <div class="searchbar">
        <v-text-field v-model="searchQuery" variant="solo-filled" placeholder="搜索用户" rounded="sm"></v-text-field>
    </div>
    <v-data-table-server :headers="headers" :items="filteredUsers" :items-length="totalUsers" :loading="loading"
        loading-text="加载中..." @update="fetchData" :hide-default-footer="true"
        :no-data-text="!userLoggedIn ? '' : '没有用户数据。'">
        <template v-slot:item="{ item }">
            <tr>
                <td>{{ item.uid }}</td>
                <td>{{ item.username }}</td>
                <td>{{ item.email }}</td>
                <td>{{ item.role === 1 ? '管理员' : '普通用户' }}</td>
                <td>{{ moment(item.create_at).format('YYYY-MM-DD HH:mm') }}</td>
                <td>{{ item.is_ban ? '封禁' : '正常' }}</td>
                <td>
                    <v-btn v-if="!item.is_ban" variant="text" @click="" icon="mdi-account-off"></v-btn>
                    <v-btn variant="text" @click="" icon="mdi-dots-horizontal"></v-btn>
                </td>
            </tr>
        </template>
    </v-data-table-server>
</template>

<style scoped>
.searchbar {
  position: sticky;
  top: 0;
  z-index: 100;
}
</style>