import axios from "axios";

// TODO:每次编译前需要将其设置为你反向代理的地址，例如http://cloud.claret.space:38080/api
export const apiUrl = "http://127.0.0.1:37881/api/v1/"
export const avatarServer = "http://127.0.0.1:37881/avatar/"

// 登录
export const loginRequest = async (username, password) => {
    return await axios.get(`${apiUrl}login`, {
        params: {
            username: username,
            password: password
        }
    })
}

// 获取验证码
export const getCaptchaCode = async (email,iscreate) => {
    return await axios.get(`${apiUrl}captcha`, {
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
    return await axios.get(`${apiUrl}verify`, {
        headers: {
            username: encodeURIComponent(username),
            Authorization: token
        }
    });
}

// 获取用户信息
export const getUserInfo = async (username) => {
    return await axios.get(`${apiUrl}users/${username}`);
}

// 获取题目列表
export const getAllProblems = async (username, token) => {
    return await axios.get(`${apiUrl}problems`, {
        headers: {
            username: encodeURIComponent(username),
            Authorization: token
        }
    });
}

// 获取题目详细信息
export const getPbDetails = async (Pid, username, token) => {
    return await axios.get(`${apiUrl}problems/${Pid}`, {
        headers: {
            username: encodeURIComponent(username),
            Authorization: token
        }
    });
}

// 获取讨论列表
export const getAllDis = async (page, itemsPerPage, username, token) => {
    return await axios.get(`${apiUrl}discussions`, {
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
    return await axios.get(`${apiUrl}discussions/${Did}`, {
        headers: {
            username: encodeURIComponent(username),
            Authorization: token
        }
    })
}

// 获取竞赛列表
export const getAllCompetitions = async (username, token) => {
    return await axios.get(`${apiUrl}competitions`, {
        headers: {
            username: encodeURIComponent(username),
            Authorization: token
        }
    })
}

// 管理员获取竞赛列表
export const getAllCompetitionsInfo = async (username, token) => {
    return await axios.get(`${apiUrl}admin/competitions`, {
        headers: {
            username: encodeURIComponent(username),
            Authorization: token
        }
    })
}

// 管理员获取指定列表信息
export const getCompetitionInfoByIDAdmin = async (username, token, cid) => {
    return await axios.get(`${apiUrl}admin/competitions/${cid}`, {
        headers: {
            username: encodeURIComponent(username),
            Authorization: token
        }
    })
}

// 获取指定用户提交记录
export const getUserSubmitRecords = async (username) => {
    return await axios.get(`${apiUrl}users/${username}/submitrecords`)
}

// 获取30天内的提交记录
export const getSubmitRecords = async (username, token) => {
    return await axios.get(`${apiUrl}submitrecords`, {
        headers: {
            username: encodeURIComponent(username),
            Authorization: token
        }
    })
}

// 获取指定讨论的所有回复
export const getComments = async (Did, username, token) => {
    return await axios.get(`${apiUrl}discussions/${Did}/comments`, {
        headers: {
            username: encodeURIComponent(username),
            Authorization: token
        }
    })
}

// 管理员获取指定题目所有信息
export const getProblemAllInfoByAdmin = async (pid, username, token) => {
    return await axios.get(`${apiUrl}admin/problems/${pid}`, {
        headers: {
            username: encodeURIComponent(username),
            Authorization: token
        },
    })
}

// 获取所有用户信息
export const getAllUsersInfo = async (username, token) => {
    return await axios.get(`${apiUrl}admin/users`, {
        headers: {
            username: encodeURIComponent(username),
            Authorization: token
        }
    })
}

// 注册
export const registerRequest = async (username, password, email, vcode) => {
    return await axios.post(`${apiUrl}register`, {
        email: email,
        username: username,
        password: password,
        vcode: vcode
    });
}

// 添加讨论
export const addDiscussion = async (Title, Content, Username, token) => {
    const formData = new FormData();
    formData.append('title', Title);
    formData.append('content', Content);
    return await axios.post(`${apiUrl}discussions`, formData, {
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
    return await axios.post(`${apiUrl}discussions/${Did}/comments`, formData, {
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
    return await axios.post(`${apiUrl}users/code`, formData, {
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

// 添加/更新题目信息
export const updateProblemInfo = async (username, token, problemInfo) => {
    return await axios.post(`${apiUrl}admin/problems`, problemInfo, {
        headers: {
            username: encodeURIComponent(username),
            Authorization: token,
            "Content-Type": "application/json"
        }
    })
}

// 删除讨论
export const deleteDiscussion = async (username, token, Did) => {
    return await axios.delete(`${apiUrl}discussions/${Did}`, {
        headers: {
            username: encodeURIComponent(username),
            Authorization: token
        }
    })
}

// 删除讨论评论
export const deleteComment = async (username, token, Cid) => {
    return await axios.delete(`${apiUrl}comments/${Cid}`, {
        headers: {
            username: encodeURIComponent(username),
            Authorization: token
        },
    })
}

// 删除题目信息
export const deleteProblemAllInfo = async (pid, username, token) => {
    return await axios.delete(`${apiUrl}admin/problems/${pid}`, {
        headers: {
            username: encodeURIComponent(username),
            Authorization: token,
        }
    })
}

// 更新用户简介
export const updateSynopsis = async (username, token, synopsis) => {
    const formData = new FormData();
    formData.append('synopsis', synopsis);
    return await axios.put(`${apiUrl}users/synopsis`, formData, {
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
    return await axios.put(`${apiUrl}users/avatar`, formData, {
        headers: {
            'Content-Type': 'multipart/form-data',
            username: encodeURIComponent(username),
            Authorization: token
        },
    });
}

// 修改密码
export const updatePassword = async (email, vcode, newPassword) => {
    return await axios.put(`${apiUrl}users/password`, {
        email: email,
        vcode: vcode,
        newpassword: newPassword
    });
}

// 封禁用户
export const banUser = async (username, token, uid) => {
    return await axios.put(`${apiUrl}admin/users/ban`, {}, {
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
    return await axios.put(`${apiUrl}admin/users/unban`, {}, {
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
    return await axios.put(`${apiUrl}admin/users/promote`, {}, {
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
    return await axios.put(`${apiUrl}admin/users/demote`, {}, {
        params: {
            uid: uid
        },
        headers: {
            username: encodeURIComponent(username),
            Authorization: token
        }
    })
}

// 管理眼获取题目列表
export const getAllProblemsAdmin = async (username, token) => {
    return await axios.get(`${apiUrl}admin/problems`, {
        headers: {
            username: encodeURIComponent(username),
            Authorization: token
        }
    });
}

// 删除竞赛
export const deleteCompetition = async (cid, username, token) => {
    return await axios.delete(`${apiUrl}admin/competitions/${cid}`, {
        headers: {
            username: encodeURIComponent(username),
            Authorization: token
        }
    })
}

// 添加/更新竞赛信息
export const updateComInfo = async (username, token, comInfo) => {
    return await axios.post(`${apiUrl}admin/competitions`, comInfo, {
        headers: {
            username: encodeURIComponent(username),
            Authorization: token,
            "Content-Type": "application/json"
        }
    })
}