<script setup>
import { onMounted, ref, computed } from 'vue';
import { getAllUsersInfo, getUserInfo,promoteUser,demoteUser,banUser,unbanUser } from '../utils/axios';
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

const getMenus = (item) => {
  let filteredMenus = []
  if (!item.isBan) {
    filteredMenus.push({ title: '封禁用户', icon: 'mdi-account-off' })
  } else {
    filteredMenus.push({ title: '解封用户', icon: 'mdi-account-check' })
  }
  if (item.role === 1) {
    filteredMenus.push({ title: '降级为普通用户', icon: 'mdi-account-minus' })
  } else {
    filteredMenus.push({ title: '提升为管理员', icon: 'mdi-account-plus' })
  }
  return filteredMenus
}

const handleMenuClick = async(menu,item)=> {
    switch(menu.title){
        case "封禁用户":
            const resp = await banUser(userName.value,token.value,item.uid)
            if(resp.data.status === 200){
                showAlert("封禁用户成功！","reload")
            }else{
                showAlert("封禁用户失败！","")
            }
            break;
        case "解封用户":
            const resp2 = await unbanUser(userName.value,token.value,item.uid)
            if(resp2.data.status === 200){
                showAlert("解封用户成功！","reload")
            }else{
                showAlert("解封用户失败！","")
            }
            break;
        case "降级为普通用户":
            const resp3 = await demoteUser(userName.value,token.value,item.uid)
            if(resp3.data.status === 200){
                showAlert("降级为普通用户成功！","reload")
            }else{
                showAlert("降级为普通用户失败！","")
            }
            break;
        case "提升为管理员":
            const resp4 = await promoteUser(userName.value,token.value,item.uid)
            if(resp4.data.status === 200){
                showAlert("提升为管理员成功！","reload")
            }else{
                showAlert("提升为管理员失败！","")
            }
            break;
    }
}

const headers = ref([
    { title: 'UID', value: 'uid', align: 'center' },
    { title: '用户名', value: 'username', align: 'center' },
    { title: '邮箱', value: 'email', align: 'center' },
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

// 根据结果不同显示不同颜色
const getStatusStyle = (status) => {
    return status ? 'color: red; font-weight: bold;' : 'color: green; font-weight: bold;';
};

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
        const userInfoResponse = await getUserInfo(userName.value, token.value);
        if(userInfoResponse.data.status !== 200){
            window.location = '/403';
            return;
        }
        userPrivilege.value = userInfoResponse.data.Info.role;
        if (userPrivilege.value !== 1) {
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
                <td :style="getStatusStyle(item.isBan)">{{ item.isBan ? '封禁' : '正常' }}</td>
                <td>
                    <v-menu>
                        <template v-slot:activator="{ props }">
                            <v-btn v-bind="props" variant="text" @click="" icon="mdi-dots-horizontal"></v-btn>
                        </template>
                        <v-list>
                            <v-list-item v-for="(menu, index) in getMenus(item)" :key="index" :value="index" @click="handleMenuClick(menu,item)">
                                <template v-slot:default="{ active, toggle }">
                                    <div class="d-flex align-center">
                                        <v-icon :icon="menu.icon" class="me-2"></v-icon>
                                        <v-list-item-title>{{ menu.title }}</v-list-item-title>
                                    </div>
                                </template>
                            </v-list-item>
                        </v-list>
                    </v-menu>
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