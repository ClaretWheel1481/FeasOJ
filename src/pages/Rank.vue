<script setup>
import { onMounted, computed, ref } from 'vue';
import { getRanking } from '../utils/api/users';
import { token } from '../utils/account';
import { showAlert } from '../utils/alert';
import { useI18n } from 'vue-i18n';
import { useRouter } from 'vue-router';

const { t } = useI18n();

const router = useRouter();
const loading = ref(true);
const rankingData = ref([])
const rankingDataLength = ref(0)

// 计算属性来判断用户是否已经登录
const userLoggedIn = computed(() => !!token.value)

const headers = ref([
  { title: t('message.username'), value: 'username', align: 'center' },
  { title: t('message.synopsis'), value: 'synopsis', align: 'center' },
  { title: 'Score', value: 'score', align: 'center' },
])

// 从后端获取数据
const fetchData = async () => {
  loading.value = true
  try {
    const response = await getRanking()
    rankingData.value = response.data.ranking;
    rankingDataLength.value = response.data.ranking.length;
  } catch (error) {
    showAlert(t("message.failed") + "!", "")
  } finally {
    loading.value = false
  }
}

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
      <v-col class="align-left-solo">
        <p style="font-size: 24px;">{{ $t('message.rank') }}</p>
        <p style="font-size: 12px;">{{ $t('message.refresh5m') }}</p>
      </v-col>
  </v-app-bar>
    <v-card style="margin: 50px;" rounded="xl" elevation="10">
    <v-data-table-server :headers="headers" :items="rankingData" :items-length="rankingDataLength"
      :loading="loading" :loading-text="$t('message.loading')" @update="fetchData" :hide-default-footer="true"
      :no-data-text="!userLoggedIn ? $t('message.nologin') : $t('message.nodata')">
      <template v-slot:item="{ item }">
        <tr>
          <td class="tabletitle">
            <v-btn @click="router.push({ path: `/profile/${item.username}` })" variant="text" block>{{ item.username
              }}</v-btn>
          </td>
          <td>{{ item.synopsis }}</td>
          <td>{{ item.score }}</td>
        </tr>
      </template>
    </v-data-table-server>
  </v-card>
</template>

<style scoped>
.tabletitle {
  color: #1e65ff;
}

.align-left-solo {
  text-align: left;
  margin-left: 10px;
}
</style>