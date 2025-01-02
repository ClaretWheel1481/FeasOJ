import axios from 'axios';
import { apiUrl } from '../axios';
import { language } from '../account';

// 管理员获取竞赛列表
export const getAllCompetitionsInfo = async (username, token) => {
    return await axios.get(`${apiUrl}/admin/competitions`, {
        headers: {
            Username: encodeURIComponent(username),
            Authorization: token,
            "Accept-Language": language.value
        }
    })
}

// 管理员获取指定列表信息
export const getCompetitionInfoByIDAdmin = async (username, token, cid) => {
    return await axios.get(`${apiUrl}/admin/competitions/${cid}`, {
        headers: {
            Username: encodeURIComponent(username),
            Authorization: token
        }
    })
}

// 管理员获取指定题目所有信息
export const getProblemAllInfoByAdmin = async (pid, username, token) => {
    return await axios.get(`${apiUrl}/admin/problems/${pid}`, {
        headers: {
            Username: encodeURIComponent(username),
            Authorization: token
        },
    })
}

// 管理员获取所有用户信息
export const getAllUsersInfo = async (username, token) => {
    return await axios.get(`${apiUrl}/admin/users`, {
        headers: {
            Username: encodeURIComponent(username),
            Authorization: token
        }
    })
}

// 管理员添加/更新题目信息
export const updateProblemInfo = async (username, token, problemInfo) => {
    return await axios.post(`${apiUrl}/admin/problems`, problemInfo, {
        headers: {
            Username: encodeURIComponent(username),
            Authorization: token,
            "Content-Type": "application/json",
            "Accept-Language": language.value
        }
    })
}

// 管理员删除题目信息
export const deleteProblemAllInfo = async (pid, username, token) => {
    return await axios.delete(`${apiUrl}/admin/problems/${pid}`, {
        headers: {
            Username: encodeURIComponent(username),
            Authorization: token,
            "Accept-Language": language.value
        }
    })
}

// 封禁用户
export const banUser = async (username, token, uid) => {
    return await axios.put(`${apiUrl}/admin/users/ban`, {}, {
        params: {
            uid: uid
        },
        headers: {
            Username: encodeURIComponent(username),
            Authorization: token,
            "Accept-Language": language.value
        }
    })
}

// 解封用户
export const unbanUser = async (username, token, uid) => {
    return await axios.put(`${apiUrl}/admin/users/unban`, {}, {
        params: {
            uid: uid
        },
        headers: {
            Username: encodeURIComponent(username),
            Authorization: token,
            "Accept-Language": language.value
        }
    })
}

// 晋升用户
export const promoteUser = async (username, token, uid) => {
    return await axios.put(`${apiUrl}/admin/users/promote`, {}, {
        params: {
            uid: uid
        },
        headers: {
            Username: encodeURIComponent(username),
            Authorization: token,
            "Accept-Language": language.value
        }
    })
}

// 降级用户
export const demoteUser = async (username, token, uid) => {
    return await axios.put(`${apiUrl}/admin/users/demote`, {}, {
        params: {
            uid: uid
        },
        headers: {
            Username: encodeURIComponent(username),
            Authorization: token,
            "Accept-Language": language.value
        }
    })
}

// 管理员获取题目列表
export const getAllProblemsAdmin = async (username, token) => {
    return await axios.get(`${apiUrl}/admin/problems`, {
        headers: {
            Username: encodeURIComponent(username),
            Authorization: token
        }
    });
}

// 管理员删除竞赛
export const deleteCompetition = async (cid, username, token) => {
    return await axios.delete(`${apiUrl}/admin/competitions/${cid}`, {
        headers: {
            Username: encodeURIComponent(username),
            Authorization: token,
            "Accept-Language": language.value
        }
    })
}

// 管理员添加/更新竞赛信息
export const updateComInfo = async (username, token, comInfo) => {
    return await axios.post(`${apiUrl}/admin/competitions`, comInfo, {
        headers: {
            Username: encodeURIComponent(username),
            Authorization: token,
            "Content-Type": "application/json",
            "Accept-Language": language.value
        }
    })
}