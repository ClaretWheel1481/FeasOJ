<!-- 题库页 -->
<script setup>
import { ref, onMounted, computed } from 'vue';
import { getAllProblems } from '../../utils/api/problems';
import { showAlert } from '../../utils/alert';
import { token } from "../../utils/account";
import { useI18n } from 'vue-i18n';
import { useRouter } from 'vue-router';
import { difficultyColor, difficultyLang } from '../../utils/dynamic_styles';

const router = useRouter();
const { t } = useI18n();

const headers = ref([
  { title: t('message.problemId'), value: 'Pid', align: 'center' },
  { title: t('message.problem'), value: 'Title', align: 'center' },
  { title: t('message.difficulty'), value: 'Difficulty', align: 'center' }
])
const problems = ref([])
const totalProblems = ref(0)
const loading = ref(true)
const searchQuery = ref('')
const showSearch = ref(false) // 控制搜索栏显示状态

// 计算属性来判断用户是否已经登录
const userLoggedIn = computed(() => !!token.value)

// 搜索功能
const filteredProblems = computed(() => {
  return problems.value.filter((problem) =>
    problem.Title.toLowerCase().includes(searchQuery.value.toLowerCase())
  )
})

// 切换搜索栏显示状态
const toggleSearch = () => {
  showSearch.value = !showSearch.value
  if (!showSearch.value) {
    searchQuery.value = '' // 关闭搜索时清空搜索内容
  }
}

// 从后端获取数据
const fetchData = async () => {
  loading.value = true
  try {
    const response = await getAllProblems();
    problems.value = response.data.problems
    totalProblems.value = problems.value.length
  } catch (error) {
    showAlert(t("message.failed") + "!", "")
  } finally {
    loading.value = false
  }
}

// 初始化数据
onMounted(async () => {
  if (!userLoggedIn.value) {
    loading.value = false;
    setTimeout(() => {
      window.location = '#/login'
      window.location.reload()
    }, 2000);
    return;
  }
  await fetchData()
})
</script>

<template>
  <v-container fluid class="pa-0">
    <v-app-bar :elevation="2">
      <p style="font-size: 24px;margin-left: 20px;">{{ t('message.problemset') }}</p>
      <v-spacer></v-spacer>
      <v-btn icon @click="toggleSearch" variant="text">
        <v-icon>mdi-magnify</v-icon>
      </v-btn>
    </v-app-bar>
    <!-- 搜索栏 -->
    <v-expand-transition>
      <v-row v-show="showSearch" class="mb-6">
        <v-col cols="12">
          <v-card elevation="0" rounded="sm">
            <v-card-text class="pa-0">
              <v-text-field 
                v-model="searchQuery" 
                :placeholder="$t('message.searchProblem')" 
                variant="solo-filled"
                density="comfortable" 
                prepend-inner-icon="mdi-magnify" 
                rounded="sm"
                hide-details 
                @keyup.enter="showSearch = false"
              ></v-text-field>
            </v-card-text>
          </v-card>
        </v-col>
      </v-row>
    </v-expand-transition>

    <!-- 数据表格 -->
    <v-row>
      <v-col cols="12">
        <v-card class="problem-table-card" elevation="0" rounded="lg">
          <v-data-table :headers="headers" :items="filteredProblems" :items-length="totalProblems" :loading="loading"
            :loading-text="$t('message.loading')"
            :no-data-text="!userLoggedIn ? $t('message.nologin') : $t('message.nodata')" class="problem-table"
            density="comfortable" hover rounded="lg" elevation="0">
            <template v-slot:item.Pid="{ item }">
              <v-chip variant="tonal" size="small" color="primary">
                {{ item.Pid }}
              </v-chip>
            </template>
            <template v-slot:item.Title="{ item }">
              <v-btn @click="router.push({ path: `/problemset/${item.Pid}` })" variant="text" color="primary">
                {{ item.Title }}
              </v-btn>
            </template>
            <template v-slot:item.Difficulty="{ item }">
              <v-chip :style="difficultyColor(item.Difficulty)">
                {{ t(difficultyLang(item.Difficulty)) }}
              </v-chip>
            </template>
          </v-data-table>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<style scoped>
.problem-table-card {
  background: rgba(var(--v-theme-surface), 0.8);
  backdrop-filter: blur(10px);
  border: 1px solid rgba(var(--v-theme-outline), 0.08);
}

.problem-table {
  --v-data-table-header-background: rgba(var(--v-theme-surface), 0.95);
  --v-data-table-row-hover-background: rgba(var(--v-theme-primary), 0.04);
}

.problem-table :deep(.v-btn) {
  transition: all 0.2s cubic-bezier(0.4, 0, 0.2, 1);
  text-align: left;
  justify-content: flex-start;
  padding: 8px 12px;
  border-radius: 8px;
}

.problem-table :deep(.v-btn:hover) {
  background: rgba(var(--v-theme-primary), 0.08);
  transform: translateX(4px);
}
</style>