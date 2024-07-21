<!-- 讨论帖子详情页 -->
<script setup>
import {
    VRow,
    VAppBar,
    VBtn,
    VAvatar,
    VProgressCircular,
    VImg,
    VCard,
    VListItemTitle,
    VListItemSubtitle,
} from "vuetify/components";
import { ref, onMounted, computed } from "vue";
import { getDisDetails, deleteDiscussion, avatarServer, getComments,deleteComment } from "../utils/axios";
import { useRoute } from "vue-router";
import { showAlert } from "../utils/alert";
import { token, userName } from "../utils/account";
import { MdPreview,MdEditor } from "md-editor-v3";
import moment from "moment";
import "md-editor-v3/lib/style.css";

const route = useRoute();
const loading = ref(true);
const discussionInfos = ref({});
const Did = route.params.Did;
const id = "preview-only";
const page = ref(1);
const comments = ref([]);
const commentContent = ref("写一条评论吧！");
const editorToolbar = [
  'bold',
  'underline',
  'italic',
  '-',
  'title',
  'strikeThrough',
  'sub',
  'sup',
  'quote',
  'unorderedList',
  'orderedList',
  '-',
  'codeRow',
  'code',
  'link',
  'image',
  'katex',
];
// 用作分页
const itemsPerPage = 5; 
const totalPages = computed(() => {
  return Math.ceil(comments.value.length / itemsPerPage);
});

const paginatedComments = computed(() => {
  const start = (page.value - 1) * itemsPerPage;
  const end = start + itemsPerPage;
  return comments.value.slice(start, end);
});

const onPageChange = (newPage) => {
  page.value = newPage;
};
// 计算属性来判断用户是否已经登录
const userLoggedIn = computed(() => !!token.value);

onMounted(async () => {
    loading.value = true;
    if (userLoggedIn.value) {
        try {
            const response = await getDisDetails(Did);
            // 获取评论
            const commentsResp = await getComments(Did,userName.value,token.value);
            comments.value = commentsResp.data.comments;
            if (response.status === 200) {
                discussionInfos.value = response.data.discussionInfo;
            }
        } catch (error) {
            showAlert(`错误，未找到对象。`, "/discussion");
        } finally {
            loading.value = false;
        }
    } else {
        window.location = "/login";
    }
});

// 删除讨论
const deleteCommentByID = async(commentID) => {
    loading.value = true;
    try {
        const delCommentResp = await deleteComment(userName.value,token.value, commentID);
        if (delCommentResp.status === 200) {
            showAlert("删除成功！", "reload");
        }
    } catch (error) {
        window.location = "/403";
    } finally {
        loading.value = false;
    }
}

// 删除讨论
const deleteDis = async () => {
    loading.value = true;
    try {
        const Did = route.params.Did;
        const resp = await deleteDiscussion(
            localStorage.getItem("username"),
            localStorage.getItem("token"),
            Did
        );
        if (resp.status === 200) {
            showAlert("删除成功！", "/discussion");
        }
    } catch (error) {
        window.location = "/403";
    } finally {
        loading.value = false;
    }
};
</script>

<template>
    <div v-if="loading" class="loading">
        <v-progress-circular indeterminate color="primary" :width="12" :size="100"></v-progress-circular>
    </div>
    <div v-else>
        <v-app-bar :elevation="0">
            <template v-slot:prepend>
                <v-btn icon="mdi-chevron-left" size="x-large" @click="$router.back"></v-btn>
            </template>
            <v-row style="align-items: center">
                <div style="margin-left: 80px"></div>
                <v-avatar size="50" color="surface-variant">
                    <v-img :src="avatarServer + discussionInfos.avatar" cover></v-img>
                </v-avatar>
                <div style="margin-left: 10px"></div>
                <p class="font-weight-black">{{ discussionInfos.username }}</p>
            </v-row>
            <template v-if="discussionInfos.username === userName" v-slot:append>
                <v-btn icon="mdi-delete" size="x-large" @click="deleteDis"></v-btn>
            </template>
        </v-app-bar>
        <div style="margin-top: 30px"></div>
        <v-card class="mx-auto" width="50%" rounded="xl" elevation="10">
            <template v-slot:title>
                <span class="font-weight-black">{{ discussionInfos.title }}</span>
            </template>
            <MdPreview :editorId="id" :modelValue="discussionInfos.content" />
        </v-card>
        <div style="margin-top: 50px"></div>
        <v-card class="mx-auto" width="50%" rounded="xl" elevation="10">
            <template v-slot:title>
                <span class="font-weight-black">评论</span>
            </template>
            <div style="max-height: 300px">
                <md-editor v-model="commentContent" :editorId="id" :toolbars="editorToolbar" :noUploadImg="true" :preview="false" :footers="[]"/>
            </div>
            <div style="margin: 10px;">
                <!-- TODO: 发布事件待处理 -->
                <v-btn color="primary" rounded="xl" @click="">发布</v-btn>
            </div>
            <v-divider></v-divider>
            <v-list>
                <!-- TODO: 删除评论待处理 -->
                <v-list-item v-for="comment in paginatedComments" :key="comment.cid">
                    <v-list-item>
                        <v-list-item-title class="username-avatar">
                            <v-avatar size="44" color="surface-variant">
                                <v-img :src="avatarServer + comment.avatar" alt="Avatar" cover></v-img>
                            </v-avatar>
                            <div style="margin: 5px;">
                                <div class="username">
                                    {{ comment.username }}
                                </div>
                                <div class="timeline">
                                    {{ moment(comment.create_at).format('MM-DD HH:mm') }}
                                </div>
                            </div>
                        </v-list-item-title>
                        <v-list-item-title class="comment-content">
                            <md-preview :modelValue="comment.content" />
                        </v-list-item-title>
                    </v-list-item>
                    <div class="buttons">
                        <!-- TODO: 回复待处理 -->
                        <!-- <v-btn rounded="xl" variant="text" color="primary" @click="">回复</v-btn> -->
                        <v-btn v-if="comment.username === userName" rounded="xl" variant="text" color="primary" @click="deleteCommentByID(comment.cid)">删除</v-btn>
                    </div>
                    <v-divider></v-divider>
                </v-list-item>
            </v-list>

            <v-pagination
            v-model="page"
            :length="totalPages"
            @input="onPageChange"
            rounded="xl"
            ></v-pagination>
        </v-card>
    </div>
</template>

<style scoped>
.buttons {
    display: flex;
    justify-content: flex-end;
    margin-top: 10px;
}

.loading {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100%;
}

.username {
    font-weight: bold;
    font-size: 1.0em
}

.comment-content {
    text-align: left;
    margin-left: -20px;
}

.timeline {
    text-align: left;
    font-size: 0.8em;
    color: #999;
}

.username-avatar {
    text-align: left;
    display: flex;
    align-items: center;
}
</style>
