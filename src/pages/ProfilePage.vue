<!-- 个人信息页 -->
<script setup>
import { ref,onMounted } from 'vue'
import { useRoute,useRouter } from 'vue-router';
import { VCard,VCardActions,VCardText,VRow,VProgressCircular,VTextField,VBtn,VAvatar,VImg,VDataTableServer } from 'vuetify/components';
import moment from 'moment';
import { verifyJWT,getUserSubmitRecords,getUserInfo,uploadAvatar } from '../utils/axios';

const showCropper = ref(false);
const route = useRoute();
const router = useRouter();
const username = route.params.Username;
const userInfo = ref({});
const userId = ref('');
const loading = ref(false);
const userSubmitRecords = ref([]);
const userSubmitRecordsLength = ref(0);
const token = ref(localStorage.getItem('token'));
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

// 上传头像至服务器
const uploadAvat = async (cropper) => {
  try {
    const canvas = cropper.getCroppedCanvas();
    const file = await new Promise((resolve) => canvas.toBlob(resolve));

    // 获取原始文件的名称和类型
    const imageData = cropper.getImageData();
    const originalFile = imageData.src ? imageData.src.split('/').pop() : 'avatar.png';
    const fileName = originalFile.split('?')[0]; // 去掉可能的查询参数
    const fileType = file.type || 'image/png';

    // 创建一个新的文件对象，保留原始文件名和类型
    const newFile = new File([file], fileName, { type: fileType });
    await uploadAvatar(newFile,username,token.value);
    alert("上传成功！");
    window.location.reload()
  } catch (error) {
    console.error(error);
  }
};

// 获取用户提交记录
const fetchData = async () => {
  try {
    const submitResponse = await getUserSubmitRecords(userId.value);
    userSubmitRecords.value = submitResponse.data.submitrecords
    userSubmitRecordsLength.value = userSubmitRecords.value.length
  } catch (error) {
    alert('错误: ', error)
  }
}

// 校验token后获取用户信息
onMounted(async () => {
  loading.value = true;
  try{
    const tokenResponse = await verifyJWT(username,token.value);
    if(tokenResponse.data.status === 200){
      const response = await getUserInfo(username);
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
        <div class="avatar-container">
          <v-avatar size="120" color="surface-variant">
            <v-img :src="userInfo.avatar" cover></v-img>
          </v-avatar>
          <v-btn icon="mdi-pencil" size="30" @click="showCropper = true" class="edit-btn"></v-btn>
        </div>
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
  <!-- FIXME:头像文件上传失败 -->
  <avatar-cropper
    v-model="showCropper"
    :labels="{ submit: '上传头像', cancel: '取消' }"
    :upload-handler="uploadAvat"
  />
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
.avatar-container {
  position: relative;
  display: inline-block;
}

.edit-btn {
  position: absolute;
  bottom: 0;
  right: 0;
}
</style>