import axios from 'axios';
import { apiUrl } from '../axios';
import { language } from '../account';

// 获取公告
export const getAnnouncement = async () => {
    return await fetch(`${apiUrl}/docs/announcement.md`)
}

// 获取通知
export const getNotification = async () => {
    return await fetch(`${apiUrl}/docs/notification.md`)
}

// 更新用户简介
export const updateSynopsis = async (username, token, synopsis) => {
    const formData = new FormData();
    formData.append('synopsis', synopsis);
    return await axios.put(`${apiUrl}/users/synopsis`, formData, {
        headers: {
            Username: encodeURIComponent(username),
            Authorization: token,
            "Accept-Language": language.value
        }
    })
}

// 修改头像
export const uploadAvatar = async (file, username, token) => {
    let formData = new FormData();
    formData.append('avatar', file);
    return await axios.put(`${apiUrl}/users/avatar`, formData, {
        headers: {
            'Content-Type': 'multipart/form-data',
            Username: encodeURIComponent(username),
            Authorization: token,
            "Accept-Language": language.value
        },
    });
}