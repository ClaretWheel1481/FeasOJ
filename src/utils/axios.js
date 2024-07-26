import axios from "axios";

// TODO:每次编译前需要将其设置为你反向代理的地址，例如http://cloud.claret.space:38080/api
export const apiUrl = "http://127.0.0.1:37881/api"
// TODO:打包前记得修改为你的头像服务器地址
export const avatarServer = "http://127.0.0.1:37881/avatar/"

//  GET请求

// 登录
export const loginRequest = async (username, password) => {
    return await axios.get(`${apiUrl}/v1/login`, {
        params: {
            username: username,
            password: password
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
export const getUserInfo = async (username,token) => {
    return await axios.get(`${apiUrl}/v1/getUserInfo`, {
        headers: {
            username: encodeURIComponent(username),
            Authorization: token
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
export const getAllDis = async (page,itemsPerPage) => {
    return await axios.get(`${apiUrl}/v1/getAllDiscussions`,{
        params: {
            page: page,
            itemsPerPage: itemsPerPage
        }
    })
}

// 获取讨论详细信息
export const getDisDetails = async (Did) => {
    return await axios.get(`${apiUrl}/v1/getDiscussionByDid/${Did}`)
}

// 获取指定用户提交记录
export const getUserSubmitRecords = async (userId) => {
    return await axios.get(`${apiUrl}/v1/getSubmitRecordsByUid/${userId}`)
}

// 获取30天内的提交记录
export const getSubmitRecords = async () => {
    return await axios.get(`${apiUrl}/v1/getAllSubmitRecords`)
}

// 获取指定讨论的所有回复
export const getComments = async (did, username, token) => {
    return await axios.get(`${apiUrl}/v1/getComment/${did}`, {
        headers: {
            username: encodeURIComponent(username),
            Authorization: token
        }
    })
}

// 管理员获取指定题目所有信息
export const getProblemAllInfoByAdmin = async (pid, username, token) => {
    return await axios.get(`${apiUrl}/v1/getProblemAllInfo/${pid}`, {
        headers: {
            username: encodeURIComponent(username),
            Authorization: token
        },
    })
}

// 获取所有用户信息
export const getAllUsersInfo = async(username,token) => {
    return await axios.get(`${apiUrl}/v1/getAllUserInfo`, {
        headers: {
            username: encodeURIComponent(username),
            Authorization: token
        }
    })
}

//    POST请求

// 注册
export const registerRequest = async (username, password, email, vcode) => {
    return await axios.post(`${apiUrl}/v2/register`, {
        email: email,
        username: username,
        password: password,
        vcode: vcode
    });
}

// 修改密码
export const updatePassword = async (email, vcode, newPassword) => {
    return await axios.post(`${apiUrl}/v2/updatePassword`, {
        email: email,
        vcode: vcode,
        newpassword: newPassword
    });
}

// 添加讨论
export const addDiscussion = async (Title, Content, Username) => {
    const formData = new FormData();
    formData.append('title', Title);
    formData.append('content', Content);
    return await axios.post(`${apiUrl}/v2/addDiscussion`, formData, {
        headers: {
            username: encodeURIComponent(Username),
        },
    });
}

// 删除讨论
export const deleteDiscussion = async (username, token, Did) => {
    const tokenVerifyResp = await verifyJWT(username, token);
    if (tokenVerifyResp.status === 200) {
        return await axios.post(`${apiUrl}/v2/deleteDiscussion/${Did}`)
    }
}

// 添加讨论评论
export const addComment = async (did, content, username, token) => {
    const formData = new FormData();
    formData.append('content', content);
    return await axios.post(`${apiUrl}/v2/addComment/${did}`, formData, {
        headers: {
            username: encodeURIComponent(username),
            Authorization: token
        },
    })
}

// 删除讨论评论
export const deleteComment = async (username, token, Cid) => {
    return await axios.post(`${apiUrl}/v2/delComment/${Cid}`, {}, {
        headers: {
            username: encodeURIComponent(username),
            Authorization: token
        },
    })
}

// 上传代码文件
export const uploadCode = async (file, pid, username, token) => {
    let formData = new FormData();
    formData.append('code', file);
    return await axios.post(`${apiUrl}/v2/uploadCode`, formData, {
        params: {
            'problem': pid
        },
        headers: {
            'Content-Type': 'multipart/form-data',
            username: encodeURIComponent(username),
            Authorization: token
        },
    });
}

// 上传头像
export const uploadAvatar = async (file, username, token) => {
    let formData = new FormData();
    formData.append('avatar', file);
    return await axios.post(`${apiUrl}/v2/uploadAvatar`, formData, {
        headers: {
            'Content-Type': 'multipart/form-data',
            username: encodeURIComponent(username),
            Authorization: token
        },
    });
}

// 更新题目信息
export const updateProblemInfo = async (username, token, problemInfo) => {
    return await axios.post(`${apiUrl}/v2/updateProblemInfo/`, problemInfo, {
        headers: {
            username: encodeURIComponent(username),
            Authorization: token,
            "Content-Type": "application/json"
        }
    })
}

// 删除题目信息
export const deleteProblemAllInfo = async (pid, username, token) => {
    return await axios.post(`${apiUrl}/v2/delProblemAllInfo/${pid}`, {}, {
        headers: {
            username: encodeURIComponent(username),
            Authorization: token,
        }
    })
}

// 封禁用户
export const banUser = async (username, token, uid) => {
    return await axios.post(`${apiUrl}/v2/banUser/`, {}, {
        params: {
            uid: uid
        },
        headers: {
            username: encodeURIComponent(username),
            Authorization: token
        }
    })
}

// 解封用户
export const unbanUser = async (username, token, uid) => {
    return await axios.post(`${apiUrl}/v2/unbanUser/`, {}, {
        params: {
            uid: uid
        },
        headers: {
            username: encodeURIComponent(username),
            Authorization: token
        }
    })
}

// 晋升用户
export const promoteUser = async (username, token, uid) => {
    return await axios.post(`${apiUrl}/v2/promoteUser/`, {}, {
        params: {
            uid: uid
        },
        headers: {
            username: encodeURIComponent(username),
            Authorization: token
        }
    })
}

// 降级用户
export const demoteUser = async (username, token, uid) => {
    return await axios.post(`${apiUrl}/v2/demoteUser/`, {}, {
        params: {
            uid: uid
        },
        headers: {
            username: encodeURIComponent(username),
            Authorization: token
        }
    })
}