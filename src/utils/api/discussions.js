import axios from 'axios';
import { apiUrl } from '../axios';

// 获取讨论列表
export const getAllDis = async (page, itemsPerPage, username, token) => {
    return await axios.get(`${apiUrl}/discussions`, {
        params: {
            page: page,
            itemsPerPage: itemsPerPage
        },
        headers: {
            username: encodeURIComponent(username),
            Authorization: token
        }
    })
}

// 获取讨论详细信息
export const getDisDetails = async (did, username, token) => {
    return await axios.get(`${apiUrl}/discussions/${did}`, {
        headers: {
            username: encodeURIComponent(username),
            Authorization: token
        }
    })
}

// 获取指定讨论的所有回复
export const getComments = async (did, username, token) => {
    return await axios.get(`${apiUrl}/discussions/comments/${did}`, {
        headers: {
            username: encodeURIComponent(username),
            Authorization: token
        }
    })
}

// 添加讨论
export const addDiscussion = async (Title, Content, Username, token) => {
    const formData = new FormData();
    formData.append('title', Title);
    formData.append('content', Content);
    return await axios.post(`${apiUrl}/discussions`, formData, {
        headers: {
            username: encodeURIComponent(Username),
            Authorization: token
        },
    });
}

// 添加评论
export const addComment = async (did, content, username, token) => {
    const formData = new FormData();
    formData.append('content', content);
    return await axios.post(`${apiUrl}/discussions/comments/${did}`, formData, {
        headers: {
            username: encodeURIComponent(username),
            Authorization: token
        },
    })
}

// 删除讨论
export const deleteDiscussion = async (username, token, did) => {
    return await axios.delete(`${apiUrl}/discussions/${did}`, {
        headers: {
            username: encodeURIComponent(username),
            Authorization: token
        }
    })
}

// 删除讨论评论
export const deleteComment = async (username, token, cid) => {
    return await axios.delete(`${apiUrl}/discussions/comments/${cid}`, {
        headers: {
            username: encodeURIComponent(username),
            Authorization: token
        },
    })
}