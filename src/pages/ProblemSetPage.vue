<!-- 题库页 -->
<script setup>
import { ref, onMounted,computed } from 'vue';
import { useRouter } from 'vue-router';
import { VTextField,VDataTableServer,VBtn } from 'vuetify/lib/components/index.mjs';
import { getAllProblems } from '../utils/axios';

const router = useRouter()

const headers = ref([
  { title: 'ID', value: 'Pid', align:'center'},
  { title: '题目', value: 'Title', align:'center'}
])
const problems = ref([])
const totalProblems = ref(0)
const loading = ref(true)
const searchQuery = ref('')
const token = ref(localStorage.getItem('token'))

// 计算属性来判断用户是否已经登录
const userLoggedIn = computed(() => !!token.value)

// 搜索功能
const filteredProblems = computed(() => {
  return problems.value.filter((problem) =>
    problem.Title.toLowerCase().includes(searchQuery.value.toLowerCase())
  )
})

// 从后端获取数据
const fetchData = async () => {
  loading.value = true
  try {
    const response = await getAllProblems();
    problems.value = response.data.problems
    totalProblems.value = problems.value.length
  } catch (error) {
    alert('错误: ', error)
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
  <!-- 搜索栏 -->
  <div class="searchbar">
    <v-text-field v-model="searchQuery" variant="solo-filled" placeholder="搜索题目" rounded="sm"></v-text-field>
  </div>
  <v-data-table-server
    :headers="headers"
    :items="filteredProblems"
    :items-length="totalProblems"
    :loading="loading"
    loading-text="加载中..."
    @update="fetchData"
    :hide-default-footer="true"
    :no-data-text="!userLoggedIn ? '你没有登录，将在2秒后跳转到登录界面。' : '没有题目数据。'"
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
</template>

<style>
.text{
    justify-content: center;
}

.tabletitle{
  color: #1e65ff;
}

.searchbar {
  position: sticky;
  top: 0;
  z-index: 100;
}
</style>