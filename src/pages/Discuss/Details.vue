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
} from "vuetify/components";
import { ref, onMounted, computed } from "vue";
import {
    getDisDetails,
    deleteDiscussion,
    avatarServer,
    getComments,
    deleteComment,
    addComment,
} from "../../utils/axios";
import { useRoute, useRouter } from "vue-router";
import { showAlert } from "../../utils/alert";
import { token, userName } from "../../utils/account";
import { MdPreview, MdEditor } from "md-editor-v3";
import moment from "moment";
import "md-editor-v3/lib/style.css";
import { useI18n } from "vue-i18n";

const { locale } = useI18n();
const { t } = useI18n();
const route = useRoute();
const router = useRouter();
const loading = ref(true);
const discussionInfos = ref({});
const Did = route.params.Did;
const id = "preview-only";
const page = ref(1);
const comments = ref([]);
const commentContent = ref("");
const editorToolbar = [
    "bold",
    "underline",
    "italic",
    "-",
    "title",
    "strikeThrough",
    "quote",
    "-",
    "codeRow",
    "code",
    "link",
    "katex",
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

// 添加评论
const addComments = async (content) => {
    loading.value = true;
    try {
        const resp = await addComment(Did, content, userName.value, token.value);
        if (resp.status === 200) {
            showAlert(t("message.success") + "!", "reload");
        }
    } catch (error) {
        window.location = "#/403";
    } finally {
        loading.value = false;
    }
};

// 删除评论
const deleteCommentByID = async (commentID) => {
    loading.value = true;
    try {
        const delCommentResp = await deleteComment(
            userName.value,
            token.value,
            commentID
        );
        if (delCommentResp.status === 200) {
            showAlert(t("message.success") + "!", "reload");
        }
    } catch (error) {
        window.location = "#/403";
    } finally {
        loading.value = false;
    }
};

// 删除讨论
const deleteDis = async () => {
    loading.value = true;
    try {
        const resp = await deleteDiscussion(userName.value, token.value, Did);
        if (resp.status === 200) {
            showAlert(t("message.success") + "!", "/discussion");
        }
    } catch (error) {
        window.location = "#/403";
    } finally {
        loading.value = false;
    }
};

onMounted(async () => {
    loading.value = true;
    if (userLoggedIn.value) {
        try {
            const response = await getDisDetails(Did, userName.value, token.value);
            // 获取评论
            const commentsResp = await getComments(Did, userName.value, token.value);
            comments.value = commentsResp.data.comments;
            if (response.status === 200) {
                discussionInfos.value = response.data.discussionInfo;
            }
        } catch (error) {
            showAlert(t("message.failed") + "!", "/discussion");
        } finally {
            loading.value = false;
        }
    } else {
        window.location = "#/login";
        window.location.reload();
    }
});
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
                <div style="margin-left: 50px"></div>
                <v-avatar size="50" color="surface-variant">
                    <v-img :src="avatarServer + discussionInfos.avatar" cover></v-img>
                </v-avatar>
                <v-btn variant="text" class="font-weight-black"
                    @click="router.push({ path: `/profile/${discussionInfos.username}` })">
                    {{ discussionInfos.username }}
                </v-btn>
            </v-row>
            <template v-if="discussionInfos.username === userName" v-slot:append>
                <v-btn icon="mdi-delete" size="x-large" @click="deleteDis"></v-btn>
            </template>
        </v-app-bar>
        <div style="margin-top: 30px"></div>
        <v-card class="mx-auto" width="75%" rounded="xl" elevation="10">
            <template v-slot:title>
                <span class="font-weight-black">{{ discussionInfos.title }}</span>
            </template>
            <MdPreview :editorId="id" :modelValue="discussionInfos.content" />
            <v-card-subtitle style="justify-self: right;">
                <p style="font-size: 12px;">{{ moment(discussionInfos.create_at).format('YYYY-MM-DD HH:mm') }}</p>
            </v-card-subtitle>
        </v-card>
        <div style="margin-top: 50px"></div>
        <v-card class="mx-auto" width="75%" rounded="xl" elevation="10">
            <template v-slot:title>
                <span class="font-weight-black">{{ $t("message.comments") }}</span>
            </template>
            <div style="max-height: 300px">
                <md-editor v-model="commentContent" :editorId="id" :toolbars="editorToolbar" :noUploadImg="true"
                    :preview="false" :footers="[]" :language="locale === 'zh_CN' ? 'zh-CN' : 'en-US'" />
            </div>
            <div style="margin: 10px">
                <v-btn color="primary" rounded="xl" @click="addComments(commentContent)">{{ $t("message.submit")
                    }}</v-btn>
            </div>
            <v-divider></v-divider>
            <v-list>
                <v-list-item v-for="comment in paginatedComments" :key="comment.cid">
                    <v-list-item>
                        <v-list-item-title class="username-avatar">
                            <v-avatar size="44" color="surface-variant">
                                <v-img :src="avatarServer + comment.avatar" alt="Avatar" cover></v-img>
                            </v-avatar>
                            <div style="margin: 5px">
                                <div class="username" @click="router.push({ path: `/profile/${comment.username}` })">
                                    {{ comment.username }}
                                </div>
                                <div class="timeline">
                                    {{ moment(comment.create_at).format("MM-DD HH:mm") }}
                                </div>
                            </div>
                        </v-list-item-title>
                        <v-list-item class="comment-content">
                            <md-preview :modelValue="comment.content" />
                        </v-list-item>
                    </v-list-item>
                    <div class="buttons">
                        <v-btn v-if="comment.username === userName" rounded="xl" variant="text" color="primary"
                            @click="deleteCommentByID(comment.cid)">{{ $t("message.delete") }}</v-btn>
                    </div>
                    <v-divider></v-divider>
                </v-list-item>
            </v-list>
            <v-pagination v-model="page" :length="totalPages" @input="onPageChange" rounded="xl" />
        </v-card>
    </div>
</template>

<style scoped>
.buttons {
    display: flex;
    justify-content: flex-end;
    margin-top: 5px;
    margin-bottom: 10px;
}

.loading {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100%;
}

.username {
    font-weight: bold;
    font-size: 1em;
    cursor: pointer;
}

.comment-content {
    max-width: fit-content;
    max-height: 350px;
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
