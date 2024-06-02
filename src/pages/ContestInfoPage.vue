<!-- 竞赛详细信息页 -->
<script setup>
import { ref, onMounted,computed } from 'vue'
import axios from 'axios'
import { useRouter,useRoute } from 'vue-router'
import { VDataTableServer,VBtn,VCard,VAppBar } from 'vuetify/lib/components/index.mjs';

const headers = ref([
  { title: '题目ID', value: 'Pid', align:'center'},
  { title: '题目名称', value: 'Title', align:'center'},
])

const router = useRouter()
const route = useRoute()

const problems = ref([])
const allProblems = ref(0)
const loading = ref(true)
const token = ref(localStorage.getItem('token'))

// 计算属性来判断用户是否已经登录
const userLoggedIn = computed(() => !!token.value)

// 从后端获取数据
const fetchData = async () => {
  loading.value = true
  try {
    const contestId = route.params.Cid;
    const response = await axios.get(`http://127.0.0.1:37881/api/v1/getContestInfo/${contestId}`)
    problems.value = response.data.contestInfo
    allProblems.value = problems.value.length
  } catch (error) {
    console.error('Error fetching data: ', error)
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
    <v-app-bar :elevation="0">
        <template v-slot:prepend>
            <v-btn icon="mdi-chevron-left" size="x-large" @click="$router.back"></v-btn>
        </template>
    </v-app-bar>
    <div class="title">
        <h1>竞赛 - {{ route.params.Cid }}</h1>
    </div>
    <v-card style="margin: 100px;" rounded="xl" elevation="10">
        <v-data-table-server
        :headers="headers"
        :items="problems"
        :items-length="allProblems"
        :loading="loading"
        loading-text="加载中..."
        @update="fetchData"
        :hide-default-footer="true"
        :no-data-text="!userLoggedIn ? '你没有登录，将在2秒后跳转到登录界面。' : '当前无竞赛数据'"
        >
        <template v-slot:item="{ item }">
            <tr>
                <td>{{ item.Pid }}</td>
                <td class="tabletitle">
                <v-btn @click="handleRowClick(item.Pid)" variant="text" block>{{ item.Title }}</v-btn>
                </td>
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