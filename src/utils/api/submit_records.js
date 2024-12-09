import axios from 'axios';
import { apiUrl } from '../axios';

// 获取指定用户提交记录
export const getUserSubmitRecords = async (username) => {
    return await axios.get(`${apiUrl}/users/${username}/submitrecords`)
}

// 获取30天内的提交记录
export const getSubmitRecords = async (username, token) => {
    return await axios.get(`${apiUrl}/submitrecords`, {
        headers: {
            username: encodeURIComponent(username),
            Authorization: token
        }
    })
}