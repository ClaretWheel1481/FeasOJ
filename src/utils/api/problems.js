import axios from 'axios';
import { apiUrl } from '../axios';
import { language } from '../account';

// 获取题目列表
export const getAllProblems = async (username, token) => {
    return await axios.get(`${apiUrl}/problems`, {
        headers: {
            Username: encodeURIComponent(username),
            Authorization: token
        }
    });
}

// 获取题目详细信息
export const getPbDetails = async (pid, username, token) => {
    return await axios.get(`${apiUrl}/problems/${pid}`, {
        headers: {
            Username: encodeURIComponent(username),
            Authorization: token
        }
    });
}

// 提交代码文件
export const uploadCode = async (file, pid, username, token) => {
    let formData = new FormData();
    formData.append('code', file);
    return await axios.post(`${apiUrl}/problems/${pid}/code`, formData, {
        headers: {
            'Content-Type': 'multipart/form-data',
            Username: encodeURIComponent(username),
            Authorization: token,
            "Accept-Language": language.value
        },
    });
}