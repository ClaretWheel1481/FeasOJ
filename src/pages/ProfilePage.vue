<!-- 个人信息页 -->
<script setup>
import { ref, onMounted, computed, reactive } from 'vue'
import { useRoute, useRouter } from 'vue-router';
import { VCard, VCardActions, VCardText, VRow, VProgressCircular, VTextField, VBtn, VAvatar, VImg, VDataTableServer } from 'vuetify/components';
import moment from 'moment';
import { getUserSubmitRecords, uploadAvatar, avatarServer, getUserInfo } from '../utils/axios';
import { showAlert } from '../utils/alert';
import { userId, userName, token } from '../utils/account';
import { useI18n } from 'vue-i18n';

const { t } = useI18n();

const userInfo = ref({});
const showCropper = ref(false);
const route = useRoute();
const router = useRouter();
const currentUsername = route.params.Username;
const loading = ref(false);
const userSubmitRecords = ref([]);
const userSubmitRecordsLength = ref(0);
const sparklineData = ref([]);
const networkloading = ref(false);
const headers = ref([
  { title: t('message.problemId'), value: 'Pid', align: 'center' },
  { title: t('message.result'), value: 'Result', align: 'center' },
  { title: t('message.lang'), value: 'Language', align: 'center' },
  { title: t('message.when'), value: 'Time', align: 'center' },
])

// 计算属性来判断用户是否已经登录
const userLoggedIn = computed(() => !!token.value)

// 登出
const logout = () => {
  localStorage.clear();
  window.location.reload();
  window.location = '#/login';
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
    networkloading.value = true;
    const resp = await uploadAvatar(newFile, userName.value, token.value);
    if (resp.data.status === 200) {
      networkloading.value = false;
      showAlert(t("message.success") + "!", "reload");
    } else {
      networkloading.value = false;
      showAlert(t("message.failed") + "!", "")
    }
  } catch (error) {
    showAlert(t("message.failed") + "!", "")
  }
};

// 获取用户提交信息
const fetchData = async () => {
  try {
    const submitResponse = await getUserSubmitRecords(userId.value);
    userSubmitRecords.value = submitResponse.data.submitrecords;
    userSubmitRecordsLength.value = userSubmitRecords.value.length;
    processSparklineData();
  } catch (error) {
    showAlert(t("message.failed") + "!", "")
  }
}

// 图表数据计算
const processSparklineData = () => {
  const counts = {};
  const now = new Date();
  const sevenDaysAgo = new Date();
  sevenDaysAgo.setDate(now.getDate() - 6);

  // 初始化最近7天的日期，包括当天
  for (let d = new Date(sevenDaysAgo); d <= now; d.setDate(d.getDate() + 1)) {
    const dateString = d.toISOString().split('T')[0];
    counts[dateString] = 0;
  }

  // 统计提交记录
  userSubmitRecords.value.forEach(record => {
    const date = new Date(record.Time);
    if (date >= sevenDaysAgo && date <= now) {
      const dateString = date.toISOString().split('T')[0];
      counts[dateString] = (counts[dateString] || 0) + 1;
    }
  });

  console.log('Counts:', counts); // 调试输出

  sparklineData.value = Object.keys(counts)
    .map(date => ({ date, count: counts[date] }))
    .sort((a, b) => new Date(a.date) - new Date(b.date)); // 按日期排序

  console.log('Sparkline Data:', sparklineData.value); // 调试输出
};

// 根据结果不同显示不同颜色
const getResultStyle = (result) => {
  switch (result) {
    case 'Compile Failed':
      return 'color: red; font-weight: bold;';
    case 'Success':
      return 'color: green; font-weight: bold;';
    case 'Failed':
      return 'color: orange; font-weight: bold;';
    case 'Wrong Answer':
      return 'color: orange; font-weight: bold;';
    default:
      return '';
  }
};

// 检验获取用户信息
onMounted(async () => {
  loading.value = true;
  try {
    if (!userLoggedIn.value) {
      router.push({ path: '/login' });
      return;
    }
    const userInfoResponse = await getUserInfo(userName.value, token.value);
    if (userInfoResponse.data.status !== 200) {
      router.push({ path: '/403' });
      return;
    }
    userInfo.value = userInfoResponse.data.Info;
    if (currentUsername !== userInfo.value.username) {
      router.push({ path: '/403' });
      return;
    }
    await fetchData();
    loading.value = false;
  } catch (error) {
    router.push({ path: '/403' });
  }
});
</script>

<template>
  <v-dialog v-model="networkloading" max-width="500px">
    <v-card rounded=xl>
      <div class="networkloading">
        <v-progress-circular indeterminate color="primary" :width="12" :size="100"></v-progress-circular>
      </div>
    </v-card>
  </v-dialog>
  <div v-if="loading" class="loading">
    <v-progress-circular indeterminate color="primary" :width="12" :size="100"></v-progress-circular>
  </div>
  <div v-else>
    <div style="margin: 10%"></div>
    <v-card class="mx-auto" max-width="50%" min-width="50%" rounded="xl" elevation="10">
      <div style="margin: 10px"></div>
      <div class="avatar-container">
        <v-avatar size="120" color="surface-variant">
          <v-img :src="avatarServer + userInfo.avatar" cover></v-img>
        </v-avatar>
        <v-btn icon="mdi-pencil" size="30" @click="showCropper = true" class="edit-btn"></v-btn>
      </div>
      <v-card-text>
        <p class="text-h4 font-weight-black">{{ userName }}</p>
        <div style="margin: 10px"></div>
        <p class="text-medium-emphasis">{{ userInfo.synopsis }}</p>
        <div style="margin: 20px"></div>
        <v-row style="justify-content: space-between;margin-inline:5px">
          <v-text-field label="Email" :model-value="userInfo.email" readonly rounded="xl"
            variant="solo-filled"></v-text-field>
          <div style="margin: 5px;"></div>
          <v-text-field label="Score" :model-value="userInfo.score" readonly rounded="xl"
            variant="solo-filled"></v-text-field>
        </v-row>
      </v-card-text>
      <v-card-actions style="justify-content: end;">
        <v-btn color="primary" variant="text" rounded="xl" style="margin-right: 10px;"
          @click="$router.push('/reset')">{{ $t('message.resetpwd') }}</v-btn>
        <v-btn color="primary" variant="text" rounded="xl" style="margin-right: 10px;" @click="logout">{{
          $t('message.logout') }}</v-btn>
      </v-card-actions>
    </v-card>
    <v-card class="mx-auto" max-width="50%" min-width="50%" rounded="xl" elevation="10" style="margin-top: 30px;">
      <v-card-text>
        <v-sheet>
          <v-sparkline :model-value="sparklineData.map(item => item.count * 10)" color="cyan" height="100"
            padding="34" line-width="1.5" smooth>
            <template v-slot:label="{ index }">
              {{ sparklineData[index].date.split('-').slice(1).join('-') }}
            </template>
          </v-sparkline>
        </v-sheet>
      </v-card-text>
    </v-card>
    <div style="margin: 30px"></div>
    <v-card class="mx-auto" max-width="50%" min-width="50%" rounded="xl" elevation="10">
      <v-data-table-server :headers="headers" :items="userSubmitRecords" :items-length="userSubmitRecordsLength"
        :loading="loading" :loading-text="$t('message.loading')" @update="fetchData" :hide-default-footer="true"
        :no-data-text="$t('message.nodata')">
        <template v-slot:item="{ item }">
          <tr>
            <td class="tabletitle">
              <v-btn @click="handleRowClick(item.Pid)" variant="text" block>{{ item.Pid }}</v-btn>
            </td>
            <td v-if="item.Result === 'Running...'" colspan="1">
              <v-progress-circular indeterminate color="primary"></v-progress-circular>
            </td>
            <td v-else :style="getResultStyle(item.Result)">{{ item.Result }}</td>
            <td>{{ item.Language }}</td>
            <td>{{ moment(item.Time).format('YYYY-MM-DD HH:mm') }}</td>
          </tr>
        </template>
      </v-data-table-server>
    </v-card>
  </div>
  <avatar-cropper v-model="showCropper" :labels="{ submit: '上传头像', cancel: $t('message.cancel') }"
    :upload-handler="uploadAvat" />
</template>

<style scoped>
.networkloading {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100%;
  margin: 100px;
}

.loading {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100%;
}

.tabletitle {
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