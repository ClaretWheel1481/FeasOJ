<!-- 状态页 -->
<script setup>
import { VDataTableServer,VCard,VBtn } from 'vuetify/lib/components/index.mjs';
import { getSubmitRecords } from '../utils/axios.js';
import { onMounted,computed, ref } from 'vue';
import moment from 'moment';
import { useRouter } from 'vue-router';
import { showAlert } from '../utils/alert';
import { token } from "../utils/account";

const router = useRouter();

const headers = ref([
  { title: '题目ID', value: 'Pid', align:'center'},
  { title: '用户UID', value: 'Uid', align:'center'},
  { title: '结果', value: 'Result', align:'center'},
  { title: '语言', value: 'Language', align:'center'},
  { title: '时间', value: 'Time', align:'center'},
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
    showAlert(error,"")
  } finally {
    loading.value = false
  }
}

// 点击题目跳转
const handleRowClick = (row) => {
  router.push({ path: `/problem/${row}` })
}

// 初始化数据
onMounted(async () => {
  if(!userLoggedIn.value){
    loading.value = false;
    setTimeout(() => {
        window.location = '/login'
    }, 2000);
    return;
  }
  await fetchData()
})
</script>

<template>
  <div class="title">
    <h1>状态</h1>
  </div>
  <v-card style="margin: 50px;" rounded="xl" elevation="10">
    <v-data-table-server
    :headers="headers"
    :items="submitrecords"
    :items-length="submitRecordsLength"
    :loading="loading"
    loading-text="加载中..."
    @update="fetchData"
    :hide-default-footer="true"
    :no-data-text="!userLoggedIn ? '你没有登录，将在2秒后跳转到登录界面。' : '没有状态数据。'"
    >
    <template v-slot:item="{ item }">
      <tr>
        <td class="tabletitle">
          <v-btn @click="handleRowClick(item.Pid)" variant="text" block>{{ item.Pid }}</v-btn>
        </td>
        <td>{{ item.Uid }}</td>
        <td>{{ item.Result }}</td>
        <td>{{ item.Language }}</td>
        <td>{{ moment(item.Time).format('YYYY-MM-DD HH:mm') }}</td>
      </tr>
    </template>
    </v-data-table-server>
  </v-card>
</template>

<style scoped>
.tabletitle{
  color: #1e65ff;
}
</style>