<!-- 状态页 -->
<script setup>
import { getSubmitRecords } from '../utils/api/submit_records.js';
import { getResultStyle } from '../utils/dynamic_styles.js';
import { onMounted, computed, ref } from 'vue';
import { useRouter } from 'vue-router';
import { showAlert } from '../utils/alert.js';
import { token } from "../utils/account.js";
import { useI18n } from 'vue-i18n';
import moment from 'moment';

const { t } = useI18n();

const router = useRouter();

const headers = ref([
  { title: t('message.problemId'), value: 'Pid', align: 'center' },
  { title: t('message.username'), value: 'Username', align: 'center' },
  { title: t('message.result'), value: 'Result', align: 'center' },
  { title: t('message.lang'), value: 'Language', align: 'center' },
  { title: t('message.when'), value: 'Time', align: 'center' },
])
const submitrecords = ref([])
const submitRecordsLength = ref(0)
const loading = ref(true)

// 计算属性来判断用户是否已经登录
const userLoggedIn = computed(() => !!token.value)

// 从后端获取数据
const fetchData = async () => {
  loading.value = true
  try {
    const response = await getSubmitRecords()
    submitrecords.value = response.data.submitrecords
    submitRecordsLength.value = submitrecords.value.length
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
    <p style="font-size: 24px;margin-left: 20px;">{{ t("message.status") }}</p>
  </v-app-bar>
  <v-card style="margin: 50px;" rounded="xl" elevation="10">
    <v-data-table-server :headers="headers" :items="submitrecords" :items-length="submitRecordsLength"
      :loading="loading" :loading-text="$t('message.loading')" @update="fetchData" :hide-default-footer="true"
      :no-data-text="!userLoggedIn ? $t('message.nologin') : $t('message.nodata')">
      <template v-slot:item="{ item }">
        <tr>
          <td class="tabletitle">
            <v-btn @click="router.push({ path: `/problem/${item.Pid}` })" variant="text" block>{{ item.Pid }}</v-btn>
          </td>
          <td class="tabletitle">
            <v-btn @click="router.push({ path: `/profile/${item.Username}` })" variant="text" block>{{ item.Username
            }}</v-btn>
          </td>
          <td v-if="item.Result === 'Running...'" colspan="1">
            <v-progress-circular indeterminate color="primary"></v-progress-circular>
          </td>
          <td v-else :style="getResultStyle(item.Result)" @click="" style="cursor: pointer;">
            <!-- TODO: 查看代码，但是当该题目属于某个竞赛且在进行时间时不可用 -->
            {{ item.Result }}
          </td>
          <td>{{ item.Language }}</td>
          <td>{{ moment(item.Time).format('YYYY-MM-DD HH:mm') }}</td>
        </tr>
      </template>
    </v-data-table-server>
  </v-card>
</template>

<style scoped>
.tabletitle {
  color: #1e65ff;
}
</style>