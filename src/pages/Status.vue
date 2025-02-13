<!-- 状态页 -->
<script setup>
import { getSubmitRecords } from '../utils/api/submit_records.js';
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

// 根据结果不同显示不同颜色
const getResultStyle = (result) => {
  switch (result) {
    case 'Compile Failed':
      return 'color: red; font-weight: bold;';
    case 'Time Limit Exceeded':
      return 'color: red; font-weight: bold;';
    case 'Success':
      return 'color: green; font-weight: bold;';
    case 'Failed':
      return 'color: orange; font-weight: bold;';
    case 'Wrong Answer':
      return 'color: orange; font-weight: bold;';
    default:
      return '';
  }
};

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
    <h1>{{ $t('message.status') }}</h1>
  </div>
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
          <td v-else :style="getResultStyle(item.Result)">{{ item.Result }}</td>
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