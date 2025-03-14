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
  <!-- 搜索栏 -->
  <div class="searchbar">
    <v-text-field v-model="searchQuery" variant="solo-filled" :placeholder="$t('message.searchProblem')"
      rounded="sm"></v-text-field>
  </div>
  <v-data-table-server :headers="headers" :items="filteredProblems" :items-length="totalProblems" :loading="loading"
    :loading-text="$t('message.loading')" @update="fetchData" :hide-default-footer="true"
    :no-data-text="!userLoggedIn ? $t('message.nologin') : $t('message.nodata')">
    <template v-slot:item="{ item }">
      <tr>
        <td>{{ item.Pid }}</td>
        <td class="tabletitle">
          <v-btn @click="router.push({ path: `/problem/${item.Pid}` })" variant="text" block>{{ item.Title }}</v-btn>
        </td>
        <td>
          <v-chip :style="difficultyColor(item.Difficulty)">
            {{ t(difficultyLang(item.Difficulty)) }}
          </v-chip>
        </td>
      </tr>
    </template>
  </v-data-table-server>
</template>

<style>
.text {
  justify-content: center;
}

.tabletitle {
  color: #1e65ff;
}

.searchbar {
  position: sticky;
  top: 0;
  z-index: 100;
}
</style>