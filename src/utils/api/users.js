import axios from 'axios';
import { apiUrl } from '../axios';

// 更新用户简介
export const updateSynopsis = async (username, token, synopsis) => {
    const formData = new FormData();
    formData.append('synopsis', synopsis);
    return await axios.put(`${apiUrl}/users/synopsis`, formData, {
        headers: {
            username: encodeURIComponent(username),
            Authorization: token,
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
            username: encodeURIComponent(username),
            Authorization: token
        },
    });
}