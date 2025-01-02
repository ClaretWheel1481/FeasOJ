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

// 加入竞赛
export const joinCompetition = async (username, token, competitionId) => {
    return await axios.post(`${apiUrl}/competitions/join/${competitionId}`,{
        headers: {
            Username: encodeURIComponent(username),
            Authorization: token,
            "Accept-Language": language.value
        }
    })
}

// 退出竞赛
export const quitCompetition = async (username, token, competitionId) => {
    return await axios.post(`${apiUrl}/competitions/quit/${competitionId}`,{
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
export const getCompetitionUsers = async (competitionId) => {
    return await axios.get(`${apiUrl}/competitions/${competitionId}/users`)
}