import axios from 'axios';
import { apiUrl } from '../axios';

// 注册
export const registerRequest = async (username, password, email, vcode) => {
    return await axios.post(`${apiUrl}/register`, {
        email: email,
        username: username,
        password: password,
        vcode: vcode
    });
}

// 登录
export const loginRequest = async (username, password) => {
    return await axios.get(`${apiUrl}/login`, {
        params: {
            username: username,
            password: password
        }
    })
}

// 获取验证码
export const getCaptchaCode = async (email,iscreate) => {
    return await axios.get(`${apiUrl}/captcha`, {
        params: {
            email: email
        },
        headers: {
            iscreate: iscreate
        }
    });
}

// 验证个人用户信息
export const verifyUserInfo = async (username, token) => {
    return await axios.get(`${apiUrl}/verify`, {
        headers: {
            username: encodeURIComponent(username),
            Authorization: token
        }
    });
}

// 获取用户信息
export const getUserInfo = async (username) => {
    return await axios.get(`${apiUrl}/users/${username}`);
}

// 修改密码
export const updatePassword = async (email, vcode, newPassword) => {
    return await axios.put(`${apiUrl}/users/password`, {
        email: email,
        vcode: vcode,
        newpassword: newPassword
    });
}