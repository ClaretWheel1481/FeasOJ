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
export const getDisDetails = async (Did, username, token) => {
    return await axios.get(`${apiUrl}/discussions/${Did}`, {
        headers: {
            username: encodeURIComponent(username),
            Authorization: token
        }
    })
}

// 获取指定讨论的所有回复
export const getComments = async (Did, username, token) => {
    return await axios.get(`${apiUrl}/discussions/${Did}/comments`, {
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
export const addComment = async (Did, content, username, token) => {
    const formData = new FormData();
    formData.append('content', content);
    return await axios.post(`${apiUrl}/discussions/${Did}/comments`, formData, {
        headers: {
            username: encodeURIComponent(username),
            Authorization: token
        },
    })
}

// 删除讨论
export const deleteDiscussion = async (username, token, Did) => {
    return await axios.delete(`${apiUrl}/discussions/${Did}`, {
        headers: {
            username: encodeURIComponent(username),
            Authorization: token
        }
    })
}

// 删除讨论评论
export const deleteComment = async (username, token, Cid) => {
    return await axios.delete(`${apiUrl}/comments/${Cid}`, {
        headers: {
            username: encodeURIComponent(username),
            Authorization: token
        },
    })
}