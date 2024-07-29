<script setup>
import { onMounted, ref, computed } from 'vue';
import { getAllUsersInfo, getUserInfo, promoteUser, demoteUser, banUser, unbanUser } from '../utils/axios';
import { showAlert } from '../utils/alert';
import { token, userName } from '../utils/account'
import { VBtn } from 'vuetify/components';
import moment from 'moment';
import { useI18n } from 'vue-i18n';

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

const handleMenuClick = async (menu, item) => {
    switch (menu.title) {
        case t("message.banUser"):
            networkloading.value = true;
            const resp = await banUser(userName.value, token.value, item.uid)
            if (resp.data.status === 200) {
                networkloading.value = false;
                showAlert(t("message.success") + "!", "reload")
            } else {
                networkloading.value = false;
                showAlert(t("message.failed") + "!", "")
            }
            break;
        case t("message.unbanUser"):
            networkloading.value = true;
            const resp2 = await unbanUser(userName.value, token.value, item.uid)
            if (resp2.data.status === 200) {
                networkloading.value = false;
                showAlert(t("message.success") + "!", "reload")
            } else {
                networkloading.value = false;
                showAlert(t("message.failed") + "!", "")
            }
            break;
        case t("message.demoteUser"):
            networkloading.value = true;
            const resp3 = await demoteUser(userName.value, token.value, item.uid)
            if (resp3.data.status === 200) {
                networkloading.value = false;
                showAlert(t("message.success") + "!", "reload")
            } else {
                networkloading.value = false;
                showAlert(t("message.failed") + "!", "")
            }
            break;
        case t("message.promoteUser"):
            networkloading.value = true;
            const resp4 = await promoteUser(userName.value, token.value, item.uid)
            if (resp4.data.status === 200) {
                networkloading.value = false;
                showAlert(t("message.success") + "!", "reload")
            } else {
                networkloading.value = false;
                showAlert(t("message.failed") + "!", "")
            }
            break;
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
        const usersInfoResp = await getAllUsersInfo(userName.value, token.value)
        users.value = usersInfoResp.data.usersInfo;
        totalUsers.value = usersInfoResp.data.usersInfo.length;
        loading.value = false;
    } catch (error) {
        showAlert(t("message.failed") + "!", "");
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
        const userInfoResponse = await getUserInfo(userName.value, token.value);
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
    <v-dialog v-model="networkloading" max-width="500px">
        <v-card rounded=xl>
            <div class="networkloading">
                <v-progress-circular indeterminate color="primary" :width="12" :size="100"></v-progress-circular>
            </div>
        </v-card>
    </v-dialog>
    <div class="searchbar">
        <v-text-field v-model="searchQuery" variant="solo-filled" :placeholder="$t('message.searchUser')"
            rounded="sm"></v-text-field>
    </div>
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
                        <v-list>
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
    position: sticky;
    top: 0;
    z-index: 100;
}

.networkloading {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100%;
    margin: 100px;
}
</style>