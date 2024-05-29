<script setup>
import { ref, onMounted,computed } from 'vue'
import axios from 'axios'
import { useRouter } from 'vue-router'

const router = useRouter()

const headers = ref([
  { title: 'ID', value: 'Pid', align:'center'},
  { title: '题目', value: 'Title', align:'center'}
])
const problems = ref([])
const totalProblems = ref(0)
const loading = ref(true)
const searchQuery = ref('')

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
    const response = await axios.get('http://127.0.0.1:37881/api/v1/getAllProblems')
    problems.value = response.data.problems
    totalProblems.value = problems.value.length
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
onMounted(fetchData)
</script>

<template>
  <div class="searchbar">
    <v-text-field v-model="searchQuery" variant="solo-filled" placeholder="搜索题目" rounded="sm"></v-text-field>
  </div>
  <v-data-table-server
    :headers="headers"
    :items="filteredProblems"
    :items-length="totalProblems"
    :loading="loading"
    @update="fetchData"
    :hide-default-footer="true"
    no-data-text="当前题库为空"
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