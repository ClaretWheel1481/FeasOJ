<script setup>
import { onMounted, ref, computed } from 'vue';
import { getAllUsersInfo, promoteUser, demoteUser, banUser, unbanUser } from '../../utils/api/admin';
import { verifyUserInfo } from '../../utils/api/auth';
import { showAlert } from '../../utils/alert';
import { token, userName } from '../../utils/account'
import { useI18n } from 'vue-i18n';
import moment from 'moment';

const { t } = useI18n();

// 计算属性来判断用户是否已经登录
const userLoggedIn = computed(() => !!token.value)
const loading = ref(true)
const userPrivilege = ref("")
const users = ref([])
const totalUsers = ref(0)
const searchQuery = ref('')
const networkloading = ref(false);

const getMenus = (item) => {
    let filteredMenus = []
    if (!item.isBan) {
        filteredMenus.push({ title: t("message.banUser"), icon: 'mdi-account-off' })
    } else {
        filteredMenus.push({ title: t("message.unbanUser"), icon: 'mdi-account-check' })
    }
    if (item.role === 1) {
        filteredMenus.push({ title: t("message.demoteUser"), icon: 'mdi-account-minus' })
    } else {
        filteredMenus.push({ title: t("message.promoteUser"), icon: 'mdi-account-plus' })
    }
    return filteredMenus
}

// 简化公共处理逻辑
const handleAction = async (action, item) => {
    networkloading.value = true;
    try {
        const resp = await action(item.uid)
        networkloading.value = false;
        showAlert(resp.data.message, "reload")
    } catch (error) {
        networkloading.value = false;
        showAlert(error.response.data.message, "")
    }
}

const handleMenuClick = (menu, item) => {
    const actions = {
        [t("message.banUser")]: banUser,
        [t("message.unbanUser")]: unbanUser,
        [t("message.demoteUser")]: demoteUser,
        [t("message.promoteUser")]: promoteUser
    };

    const action = actions[menu.title];
    if (action) {
        handleAction(action, item);
    }
}

const headers = ref([
    { title: 'UID', value: 'uid', align: 'center' },
    { title: t("message.username"), value: 'username', align: 'center' },
    { title: t("message.email"), value: 'email', align: 'center' },
    { title: t("message.role"), value: 'role', align: 'center' },
    { title: t("message.timeOfRegister"), value: 'registerTime', align: 'center' },
    { title: t("message.status"), value: 'status', align: 'center' },
    { title: t("message.operation"), value: 'actions', align: 'center' },
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
        const usersInfoResp = await getAllUsersInfo()
        users.value = usersInfoResp.data.usersInfo;
        totalUsers.value = usersInfoResp.data.usersInfo.length;
    } catch (error) {
        showAlert(t("message.failed") + "!", "");
    } finally {
        loading.value = false;
    }
}

onMounted(async () => {
    loading.value = true;
    try {
        if (!userLoggedIn.value) {
            window.location = "#/login";
            window.location.reload();
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
    <v-app-bar :elevation="2" height="55">
        <template v-slot:prepend>
            <v-btn icon="mdi-chevron-left" size="x-large" @click="$router.back"></v-btn>
        </template>
        <div class="searchbar">
            <v-text-field v-model="searchQuery" flat variant="solo-filled" :placeholder="$t('message.searchUser')"></v-text-field>
        </div>
    </v-app-bar>
    <v-dialog v-model="networkloading" max-width="500px">
        <v-card rounded=xl>
            <div class="networkloading">
                <v-progress-circular indeterminate color="primary" :width="12" :size="100"></v-progress-circular>
            </div>
        </v-card>
    </v-dialog>
    <v-data-table-server :headers="headers" :items="filteredUsers" :items-length="totalUsers" :loading="loading"
        :loading-text="$t('message.loading')" @update="fetchData" :hide-default-footer="true"
        :no-data-text="!userLoggedIn ? $t('message.nologin') : $t('message.nodata')">
        <template v-slot:item="{ item }">
            <tr>
                <td>{{ item.uid }}</td>
                <td>{{ item.username }}</td>
                <td>{{ item.email }}</td>
                <td>{{ item.role === 1 ? $t('message.admin') : $t('message.regularUser') }}</td>
                <td>{{ moment(item.create_at).format('YYYY-MM-DD HH:mm') }}</td>
                <td :style="getStatusStyle(item.isBan)">{{ item.isBan ? $t('message.isBan') : $t('message.normal') }}
                </td>
                <td>
                    <v-menu>
                        <template v-slot:activator="{ props }">
                            <v-btn v-bind="props" variant="text" icon="mdi-dots-horizontal"></v-btn>
                        </template>
                        <v-list rounded="xl">
                            <v-list-item v-for="(menu, index) in getMenus(item)" :key="index" :value="index"
                                @click="handleMenuClick(menu, item)">
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
    width: 100%;
    margin-top: 20px;
}
</style>