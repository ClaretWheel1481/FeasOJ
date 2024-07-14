<!-- 讨论帖子列表页 -->
<script setup>
import { VCard,VDataTableServer,VBtn,VFab } from 'vuetify/components'
import { useRouter } from 'vue-router'
import moment from 'moment';
import { ref, onMounted,computed } from 'vue'
import { getAllDis } from '../utils/axios';

const headers = ref([
  { title: '帖子', value: 'Title', align:'center'},
  { title: '发帖人', value: 'Owner', align:'center'},
  { title: '发布时间', value: 'Description', align:'center'}
])

const router = useRouter()

const discuss = ref([])
const allDiscuss = ref(0)
const loading = ref(true)
const token = ref(localStorage.getItem('token'))

// 计算属性来判断用户是否已经登录
const userLoggedIn = computed(() => !!token.value)

// 从后端获取数据
const fetchData = async () => {
  loading.value = true
  try {
    const response = await getAllDis();
    discuss.value = response.data.discussions
    allDiscuss.value = discuss.value.length
  } catch (error) {
    alert('错误: ', error)
  } finally {
    loading.value = false
  }
}

// 点击竞赛跳转
const handleRowClick = (row) => {
  router.push({ path: `/discussion/${row}` })
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
        <h1>讨论</h1>
    </div>
    <v-card class="mx-auto my-8" width="80%" elevation="10" rounded="xl">
        <v-data-table-server
        :headers="headers"
        :items="discuss"
        :items-length="allDiscuss"
        :loading="loading"
        loading-text="加载中..."
        @update="fetchData"
        :hide-default-footer="true"
        :no-data-text="!userLoggedIn ? '你没有登录，将在2秒后跳转到登录界面。' : '当前无帖子数据'"
        >
        <template v-slot:item="{ item }">
        <tr>
            <td class="disctitle">
                <v-btn @click="handleRowClick(item.tid)" variant="text" block>{{ item.title }}</v-btn>
            </td>
            <td>{{ item.username }}</td>
            <td>{{ moment(item.create_at).format('YYYY-MM-DD HH:mm') }}</td>
        </tr>
        </template>
        </v-data-table-server>
    </v-card>
    <div class="fab">
      <v-fab fixed icon="mdi-plus" size="64" color="primary" elevation="10" v-if="!loading" @click="$router.push('/discussion/create')"></v-fab>
    </div>
</template>

<style scoped>
.disctitle{
  color: #1e65ff;
  width: 60%;
}

.fab {
  position: fixed;
  bottom: 80px;
  right: 180px;
}
</style>