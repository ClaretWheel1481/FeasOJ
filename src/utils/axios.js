import axios from "axios";

// TODO:每次编译前需要将其设置为你反向代理的地址，例如http://cloud.claret.space:38080/api
export const apiUrl = "http://127.0.0.1:37881/api"
// TODO:打包前记得修改为你的头像服务器地址
export const avatarServer = "http://127.0.0.1:37881/avatar/"

// 登录
export const loginRequest = async (username,password) => {
    return await axios.get(`${apiUrl}/v1/login`, {
        params: {
            username:username,
            password:password
        }
    })
}

// Token校验
export const verifyJWT = async (username,token) => {
    return await axios.get(`${apiUrl}/v1/verifyToken`,{
        headers:{
            username:encodeURIComponent(username),
            Authorization:token
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

// 获取用户信息
export const getUserInfo = async (username) => {
    return await axios.get(`${apiUrl}/v1/getUserInfo`,{
        headers:{
          username:encodeURIComponent(username)
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

// 注册
export const registerRequest = async (username,password,email,vcode) => {
    return await axios.post(`${apiUrl}/v2/register`, {
        email: email,
        username: username,
        password: password,
        vcode: vcode
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

// 添加讨论
export const addDiscussion = async (Title,Content,Username) => {
    const formData = new FormData();
    formData.append('title', Title);
    formData.append('content', Content);
    return await axios.post(`${apiUrl}/v2/addDiscussion`,formData,{
        headers: {
            username: encodeURIComponent(Username),
        },
    });
}

// 删除讨论
export const deleteDiscussion = async(username,token,Tid) => {
    const tokenVerifyResp = await verifyJWT(username,token);
    if(tokenVerifyResp.status === 200){
        return await axios.post(`${apiUrl}/v2/deleteDiscussion/${Tid}`)
    }
}

// 上传代码文件
export const uploadCode = async (file,pid,username,token) => {
    let formData = new FormData();
    formData.append('code', file);
    return await axios.post(`${apiUrl}/v2/uploadCode`,formData,{
        params: {
            'problem' : pid
        },
        headers: {
            'Content-Type':'multipart/form-data',
            username: encodeURIComponent(username),
            Authorization: token
        },
    });
}

// 上传头像
export const uploadAvatar = async (file,username,token) => {
    let formData = new FormData();
    formData.append('avatar', file);
    return await axios.post(`${apiUrl}/v2/uploadAvatar`,formData,{
        headers: {
            'Content-Type':'multipart/form-data',
            username: encodeURIComponent(username),
            Authorization: token
        },
    });
}