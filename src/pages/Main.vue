<script setup>
import { ref, onMounted } from 'vue';
import { getAnnouncement, getNotification } from '../utils/api/users';

const panel = ref([0, 1]);
const announcement = ref('');
const notice = ref('');

onMounted(async () => {
  announcement.value = await getAnnouncement();
  notice.value = await getNotification();
});
</script>

<template>
  <div class="title">
    <v-img src="logo.png" width="100px" height="100px" style="margin: 20px;"></v-img>
    <h1>FeasOJ</h1>
  </div>
  <v-row>
    <!-- TODO: 固定长度分布 -->
    <v-card rounded="xl" style="margin: 50px;" elevation="5">
      <v-expansion-panels v-model="panel">
        <v-expansion-panel :title="$t('message.announcement')" :text="announcement" style="white-space: pre-line;">
        </v-expansion-panel>
        <v-expansion-panel :title="$t('message.notice')" :text="notice" style="white-space: pre-line;">
        </v-expansion-panel>
      </v-expansion-panels>
    </v-card>
    <v-card rounded="xl" style="margin: 50px;" elevation="5">
      
    </v-card>
  </v-row>
</template>

<style scoped>
.title {
  align-items: center;
  display: flex;
  flex-direction: column;
  justify-content: center;
}
</style>
