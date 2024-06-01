<script setup>
import { ref, onMounted,computed } from 'vue'
import axios from 'axios'
import { useRouter } from 'vue-router'
import { VTextField,VDataTableServer,VBtn } from 'vuetify/lib/components/index.mjs';

const headers = ref([
  { title: 'ID', value: 'Cid', align:'center'},
  { title: '竞赛名称', value: 'Title', align:'center'},
  { title: '开始时间', value: 'Start_time', align:'center'},
  { title: '结束时间', value: 'End_time', align:'center'},
  { title: '描述', value: 'Description', align:'center'}
])

const router = useRouter()

const contests = ref([])
const allContests = ref(0)
const loading = ref(true)
const searchQuery = ref('')
const token = ref(localStorage.getItem('token'))

// 计算属性来判断用户是否已经登录
const userLoggedIn = computed(() => !!token.value)

// 搜索功能
const filteredContests = computed(() => {
  return contests.value.filter((contest) =>
    contest.Title.toLowerCase().includes(searchQuery.value.toLowerCase())
  )
})

// 从后端获取数据
const fetchData = async () => {
  loading.value = true
  try {
    const response = await axios.get('http://127.0.0.1:37881/api/v1/getAllContests')
    contests.value = response.data.contests
    allContests.value = contests.value.length
  } catch (error) {
    console.error('Error fetching data: ', error)
  } finally {
    loading.value = false
  }
}

// 点击竞赛跳转
const handleRowClick = (row) => {
  router.push({ path: `/contest/${row}` })
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
    <div class="searchbar">
      <v-text-field v-model="searchQuery" variant="solo-filled" placeholder="搜索竞赛" rounded="sm"></v-text-field>
    </div>
    <v-data-table-server
      :headers="headers"
      :items="filteredContests"
      :items-length="allContests"
      :loading="loading"
      loading-text="加载中..."
      @update="fetchData"
      :hide-default-footer="true"
      :no-data-text="!userLoggedIn ? '你没有登录，将在2秒后跳转到登录界面。' : '没有竞赛数据。'"
    >
    <template v-slot:item="{ item }">
      <tr>
        <td>{{ item.Cid }}</td>
        <td class="tabletitle">
          <v-btn @click="handleRowClick(item.Cid)" variant="text" block>{{ item.Title }}</v-btn>
        </td>
        <td>{{ item.Start_time }}</td>
        <td>{{ item.End_time }}</td>
        <td>{{ item.Description }}</td>
      </tr>
    </template>
    </v-data-table-server>
</template>