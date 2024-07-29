<!-- 讨论帖子列表页 -->
<script setup>
import { VCard, VDataTableServer, VBtn, VFab } from 'vuetify/components'
import { useRouter } from 'vue-router'
import moment from 'moment';
import { ref, onMounted, computed, watch } from 'vue'
import { getAllDis } from '../utils/axios';
import { showAlert } from '../utils/alert';
import { token } from '../utils/account';
import { useI18n } from 'vue-i18n';

const { t } = useI18n();

const headers = ref([
  { title: t('message.discussion'), value: 'Title', align: 'center' },
  { title: t('message.discussionOwner'), value: 'Owner', align: 'center' },
  { title: t('message.when'), value: 'Description', align: 'center' }
])

const router = useRouter()
// 用作分页
const itemsPerPage = ref(12)
const page = ref(1)
const discuss = ref([])
const discussCount = ref(0)
const loading = ref(true)

// 计算属性来判断用户是否已经登录
const userLoggedIn = computed(() => !!token.value)

// 从后端获取数据
const fetchData = async () => {
  loading.value = true
  try {
    const response = await getAllDis(page.value, itemsPerPage.value);
    discuss.value = response.data.discussions
    discussCount.value = response.data.total
  } catch (error) {
    showAlert(t("message.failed") + "!", '/discussion')
  } finally {
    loading.value = false
  }
}

// 点击讨论跳转
const handleRowClick = (row) => {
  router.push({ path: `/discussion/${row}` })
}

// 监听分页变化
watch(page, async () => {
  await fetchData()
})

// 初始化数据
onMounted(async () => {
  if (!userLoggedIn.value) {
    loading.value = false;
    setTimeout(() => {
      window.location = '#/login'
      window.location.reload()
    }, 2000);
    return;
  }
  await fetchData()
})
</script>

<template>
  <div class="title">
    <h1>{{ $t('message.discussion') }}</h1>
  </div>
  <v-card class="mx-auto my-8" width="80%" elevation="10" rounded="xl">
    <v-data-table-server :headers="headers" :items="discuss" :items-length="discussCount" :loading="loading"
      :loading-text="$t('message.loading')" @update="fetchData" :hide-default-footer="true"
      :items-per-page="itemsPerPage" v-model:page="page"
      :no-data-text="!userLoggedIn ? $t('message.nologin') : $t('message.nodata')">
      <template v-slot:item="{ item }">
        <tr>
          <td class="disctitle">
            <v-btn @click="handleRowClick(item.did)" variant="text" block>{{ item.title }}</v-btn>
          </td>
          <td>{{ item.username }}</td>
          <td>{{ moment(item.create_at).format('YYYY-MM-DD HH:mm') }}</td>
        </tr>
      </template>
      <template v-slot:bottom>
        <v-pagination v-model="page" :length="Math.ceil(discussCount / itemsPerPage)" rounded="xl" />
      </template>
    </v-data-table-server>
  </v-card>
  <div class="fab">
    <v-fab fixed icon="mdi-plus" size="64" color="primary" elevation="10" v-if="userLoggedIn && !loading"
      @click="$router.push('/discussion/create')"></v-fab>
  </div>
</template>

<style scoped>
.disctitle {
  color: #1e65ff;
  width: 60%;
}

.fab {
  position: fixed;
  bottom: 80px;
  right: 180px;
}
</style>