<script setup>
import { ref,onMounted } from 'vue'
import axios from 'axios';
import { useRoute } from 'vue-router';
import { VCard,VCardActions,VCardText,VCardTitle,VRow,VCardItem } from 'vuetify/components';
import '@mdi/font/css/materialdesignicons.css';

const route = useRoute();
const username = route.params.Username;
const userInfo = ref({});

const logout = () => {
  localStorage.clear();
  window.location = '/'
};

onMounted(async () => {
  try {
    const response = await axios.get('http://127.0.0.1:37881/api/v1/getUserInfo',{
      params:{
        username:username
      }
    }
    );
    userInfo.value = response.data.Info;
  } catch (error) {
    console.error('请求题目信息时发生错误:', error);
  }
});
</script>

<template>
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
        <v-btn color="primary" variant="text" rounded="xl" style="margin-right: 10px;" @click="logout">退出登录</v-btn>
      </v-card-actions>
    </v-card>
    <!-- <v-card class="mx-auto" max-width="30%" min-width="30%" rounded="xl" elevation="10">

    </v-card> -->
  </v-row>
</template>

<style scoped>

</style>