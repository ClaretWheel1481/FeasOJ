import axios from "axios";

export const apiUrl = "http://127.0.0.1:37881/api"

// 登录
export const loginRequest = async (username,password) => {
    return await axios.get(`${apiUrl}/v1/login`, {
        params: {
            username:username,
            password:password
        }
    })
}

// 注册
export const registerRequest = async (username,password,email,vcode) => {
    return await axios.post('http://127.0.0.1:37881/api/v2/register', {
        email: email,
        username: username,
        password: password,
        vcode: vcode
      });
}

// Token校验
export const verifyJWT = async (username,token) => {
    return await axios.get('http://127.0.0.1:37881/api/v1/verifyToken',{
        params:{
          username:username
        },
        headers:{
          token:token
        }
    })
}

// 获取验证码
export const getCaptchaCode = async (email) => {
    return await axios.get(`${apiUrl}/v1/getCaptcha`, {
        params: { 
            email: email 
        }
    });
}

// 修改密码
export const updatePassword = async (email,vcode,newPassword) => {
    return await axios.post(`${apiUrl}/v2/updatePassword`, {
        email: email,
        vcode: vcode,
        newpassword: newPassword
    });
}

// 获取用户信息
export const getUserInfo = async (username) => {
    return await axios.get(`${apiUrl}/v1/getUserInfo`,{
        params:{
          username:username
        }
    });
}

// 获取题目列表
export const getAllProblems = async () => {
    return await axios.get(`${apiUrl}/v1/getAllProblems`);
}

// 获取题目详细信息
export const getPbDetails = async (Pid) => {
    return await axios.get(`${apiUrl}/v1/getProblemInfo/${Pid}`);
}

// 获取讨论列表
export const getAllDis = async () => {
    return await axios.get(`${apiUrl}/v1/getAllDiscussions`)
}

// 获取讨论详细信息
export const getDisDetails = async (Tid) => {
    return await axios.get(`${apiUrl}/v1/getDiscussionByTid/${Tid}`)
}

// 获取指定用户提交记录
export const getUserSubmitRecords = async (userId) => {
    return await axios.get(`${apiUrl}/v1/getSubmitRecordsByUid/${userId}`)
}

// 获取30天内的提交记录
export const getSubmitRecords = async () => {
    return await axios.get(`${apiUrl}/v1/getAllSubmitRecords`)
}
