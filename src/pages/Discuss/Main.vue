<!-- 讨论帖子列表页 -->
<script setup>
import { useRouter } from 'vue-router';
import { ref, onMounted, computed, watch } from 'vue';
import { getAllDis } from '../../utils/api/discussions';
import { showAlert } from '../../utils/alert';
import { token } from '../../utils/account';
import { useI18n } from 'vue-i18n';
import moment from 'moment';

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
    discuss.value = response.data.discussions ? response.data.discussions : [];
    discussCount.value = response.data.total ? response.data.total : 0;
  } catch (error) {
    showAlert(t("message.failed") + "!", '/discussion')
  } finally {
    loading.value = false
  }
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
  <v-app-bar :elevation="2">
    <p style="font-size: 24px;margin-left: 20px;">{{ t('message.discussion') }}</p>
  </v-app-bar>
  <v-card class="mx-auto my-8" width="85%" elevation="10" rounded="xl">
    <v-data-table-server :headers="headers" :items="discuss" :items-length="discussCount" :loading="loading"
      :loading-text="$t('message.loading')" @update="fetchData" :hide-default-footer="true"
      :items-per-page="itemsPerPage" v-model:page="page"
      :no-data-text="!userLoggedIn ? $t('message.nologin') : $t('message.nodata')">
      <template v-slot:item="{ item }">
        <tr>
          <td class="disctitle">
            <v-btn @click="router.push({ path: `/discussion/${item.did}` })" variant="text" block>{{ item.title
              }}</v-btn>
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
    <v-fab fixed icon="mdi-plus" size="64" color="primary" elevation="8" v-if="userLoggedIn && !loading"
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