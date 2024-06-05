<script setup>
import { ref,onMounted } from 'vue'
import axios from 'axios';
import { useRoute,useRouter } from 'vue-router';
import { VCard,VCardActions,VCardText,VRow,VProgressCircular,VTextField,VBtn,VAvatar,VImg,VDataTableServer } from 'vuetify/components';
import '@mdi/font/css/materialdesignicons.css';
import moment from 'moment';

const route = useRoute();
const router = useRouter();
const username = route.params.Username;
const userInfo = ref({});
const userId = ref('');
const loading = ref(false);
const userSubmitRecords = ref([])
const userSubmitRecordsLength = ref(0)
const headers = ref([
  { title: '题目ID', value: 'Pid', align:'center'},
  { title: '结果', value: 'Result', align:'center'},
  { title: '语言', value: 'Language', align:'center'},
  { title: '时间', value: 'Time', align:'center'},
])

const logout = () => {
  localStorage.clear();
  window.location = '/'
};

// 点击题目跳转
const handleRowClick = (row) => {
  router.push({ path: `/problem/${row}` })
}

const fetchData = async () => {
  try {
    const submitResponse = await axios.get(`http://127.0.0.1:37881/api/v1/getSubmitRecordsByUid/${userId.value}`)
    userSubmitRecords.value = submitResponse.data.submitrecords
    userSubmitRecordsLength.value = userSubmitRecords.value.length
  } catch (error) {
    console.error('Error fetching data: ', error)
  }
}

onMounted(async () => {
  loading.value = true;
  try{
    const tokenResponse = await axios.get('http://127.0.0.1:37881/api/v1/verifyToken',{
      params:{
        username:username
      },
      headers:{
        token:localStorage.getItem('token')
      }
    })
    if(tokenResponse.data.status === 200){
      const response = await axios.get('http://127.0.0.1:37881/api/v1/getUserInfo',{
        params:{
          username:username
        }
      }
      );
      userInfo.value = response.data.Info;
      userId.value = response.data.Info.uid;
      await fetchData();
    }else {
      window.location='/403'
    }
  }catch(error){
    window.location='/403'
  }finally{
    loading.value=false;
  }
});
</script>

<template>
  <div v-if="loading" class="loading">
    <v-progress-circular indeterminate color="primary" :width="12" :size="100"></v-progress-circular>
  </div>
  <div v-else>
    <div style="margin: 10%"></div>
      <v-card class="mx-auto" max-width="50%" min-width="50%" rounded="xl" elevation="10">
        <div style="margin: 10px"></div>
        <!-- TODO:头像点击后弹出修改头像组件 -->
        <v-btn icon size="120">
          <v-avatar size="120" color="surface-variant">
            <v-img :src="userInfo.avatar" cover>
            </v-img>
          </v-avatar>
        </v-btn>
        <v-card-text>
          <p class="text-h4 font-weight-black">{{userInfo.username}}</p>
          <div style="margin: 10px"></div>
          <p class="text-medium-emphasis">{{userInfo.synopsis}}</p>
          <div style="margin: 20px"></div>
          <v-row style="justify-content: space-between;margin-inline:5px">
            <v-text-field
                label="Email"
                :model-value="userInfo.email"
                readonly
                rounded="xl"
                variant="solo-filled"
              ></v-text-field>
              <div style="margin: 5px;"></div>
              <v-text-field
                label="Score"
                :model-value="userInfo.score"
                readonly
                rounded="xl"
                variant="solo-filled"
              ></v-text-field>
          </v-row>
        </v-card-text>
        <v-card-actions style="justify-content: end;">
          <v-btn color="primary" variant="text" rounded="xl" style="margin-right: 10px;" @click="$router.push('/reset')">修改密码</v-btn>
          <v-btn color="primary" variant="text" rounded="xl" style="margin-right: 10px;" @click="logout">退出登录</v-btn>
        </v-card-actions>
      </v-card>
      <div style="margin: 30px"></div>
      <v-card class="mx-auto" max-width="50%" min-width="50%" rounded="xl" elevation="10">
        <v-data-table-server
        :headers="headers"
        :items="userSubmitRecords"
        :items-length="userSubmitRecordsLength"
        :loading="loading"
        loading-text="加载中..."
        @update="fetchData"
        :hide-default-footer="true"
        no-data-text="没有状态数据。"
        >
        <template v-slot:item="{ item }">
          <tr>
            <td class="tabletitle">
              <v-btn @click="handleRowClick(item.Pid)" variant="text" block>{{ item.Pid }}</v-btn>
            </td>
            <td>{{ item.Result }}</td>
            <td>{{ item.Language }}</td>
            <td>{{ moment(item.Time).format('YYYY-MM-DD HH:mm') }}</td>
          </tr>
        </template>
        </v-data-table-server>
      </v-card>
  </div>
</template>

<style scoped>
.loading{
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100%;
}
.tabletitle{
  color: #1e65ff;
}
</style>