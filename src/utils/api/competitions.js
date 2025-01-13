import axios from 'axios';
import { apiUrl } from '../axios';
import { language } from '../account';

// 获取竞赛列表
export const getAllCompetitions = async (username, token) => {
    return await axios.get(`${apiUrl}/competitions`, {
        headers: {
            Username: encodeURIComponent(username),
            Authorization: token
        }
    })
}

// 获取指定竞赛ID信息
export const getCompetitionById = async (username, token, competitionId) => {
    return await axios.get(`${apiUrl}/competitions/info/${competitionId}`, {
        headers: {
            Username: encodeURIComponent(username),
            Authorization: token,
        }
    })
}

// 加入竞赛
export const joinCompetition = async (username, token, competitionId) => {
    return await axios.post(`${apiUrl}/competitions/join/${competitionId}`,{},{
        headers: {
            Username: encodeURIComponent(username),
            Authorization: token,
            "Accept-Language": language.value
        }
    })
}

// 加入有密码的竞赛
export const joinCompWithPwd = async (username, token, competitionId,competitionPwd) => {
    return await axios.post(`${apiUrl}/competitions/join/pwd/${competitionId}`,{},{
        params:{
            password: competitionPwd
        },
        headers: {
            Username: encodeURIComponent(username),
            Authorization: token,
            "Accept-Language": language.value
        }
    })
}

// 退出竞赛
export const quitCompetition = async (username, token, competitionId) => {
    return await axios.post(`${apiUrl}/competitions/quit/${competitionId}`,{},{
        headers: {
            Username: encodeURIComponent(username),
            Authorization: token,
            "Accept-Language": language.value
        }
    })
}

// 查询用户是否在指定竞赛中
export const isInCompetition = async (username, token, competitionId) => {
    return await axios.get(`${apiUrl}/competitions/${competitionId}/in`, {
        headers: {
            Username: encodeURIComponent(username),
            Authorization: token,
            "Accept-Language": language.value
        }
    })
}

// 获取指定竞赛的所有用户
export const getCompetitionUsers = async (username, token, competitionId) => {
    return await axios.get(`${apiUrl}/competitions/info/${competitionId}/users`,{
        headers: {
            Username: encodeURIComponent(username),
            Authorization: token,
        }
    })
}

// 获取指定竞赛的所有题目
export const getCompetitionProblems = async (username, token, competitionId) => {
    return await axios.get(`${apiUrl}/competitions/info/${competitionId}/problems`,{
        headers: {
            Username: encodeURIComponent(username),
            Authorization: token,
        }
    })
}