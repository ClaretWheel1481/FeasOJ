<script setup>
import { ref,onMounted } from 'vue'
import axios from 'axios';
import { useRoute } from 'vue-router';
import { VCard,VCardActions,VCardText,VRow,VProgressCircular,VTextField,VBtn } from 'vuetify/components';
import '@mdi/font/css/materialdesignicons.css';

const route = useRoute();
const username = route.params.Username;
const userInfo = ref({});

const loading = ref(false);

const logout = () => {
  localStorage.clear();
  window.location = '/'
};

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
    <v-row style="justify-content: space-between; margin-inline:5px">
      <v-card class="mx-auto" max-width="50%" min-width="50%" rounded="xl" elevation="10">
        <v-card-text>
          <p class="text-h4 font-weight-black">{{userInfo.Username}}</p>
          <div style="margin: 10px"></div>
          <p class="text-medium-emphasis">{{userInfo.Synopsis}}</p>
          <div style="margin: 20px"></div>
          <v-row style="justify-content: space-between;margin-inline:5px">
            <v-text-field
                label="Email"
                :model-value="userInfo.Email"
                readonly
                rounded="xl"
                variant="solo-filled"
              ></v-text-field>
              <div style="margin: 5px;"></div>
              <v-text-field
                label="Score"
                :model-value="userInfo.Score"
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
    </v-row>
  </div>
</template>

<style scoped>
.loading{
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100%;
}
</style>