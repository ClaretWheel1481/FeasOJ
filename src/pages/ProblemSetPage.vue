<script setup>
import { ref, onMounted } from 'vue'
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
const handleRowClick = (row) => {
  router.push({ path: `/problem/${row}` })
}
onMounted(fetchData)
</script>

<template>
  <div class="title">
    <h1>题目列表</h1>
  </div>
  <v-data-table-server
    :headers="headers"
    :items="problems"
    :items-length="totalProblems"
    :loading="loading"
    @update:options="fetchData"
  >
  <template v-slot:item="{ item }">
      <tr>
        <td>{{ item.Pid }}</td>
        <td @click="handleRowClick(item.Pid)" class="tableTitle">{{ item.Title }}</td>
      </tr>
  </template>
  </v-data-table-server>
</template>

<style>
.text{
    justify-content: center;
}

.tableTitle{
  /* 超链接颜色 */
  color: #1e65ff;
  /* 字体大小 */
  font-size: 14px;
  /* 字体加粗 */
  font-weight: bold;
  cursor: pointer;
}
</style>