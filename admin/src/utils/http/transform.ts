import { AxiosResponse } from "axios"
import router from '@/router'
import { notification } from 'ant-design-vue';

const handleError = (error: string, msg: string) => {
    notification.error({
        message: error,
        description: msg,
        duration: 3000,
    })
}

export const handleResponse = (res: AxiosResponse<{ code: number, data: any, msg: string }>): any => {
    if (res.status !== 200) {
        console.error('error = ', res.statusText)
        return Promise.reject(res.statusText)
    }
    const { code, data, msg } = res.data

    if (code === 0) {
        return data
    } else if (code === 401) {
        handleError('error', '登录失效，请重新登录')
        localStorage.removeItem('token')
        router.push('/login')
        return Promise.reject('登录失效，请重新登录')
    } else {
        handleError('error', msg)
        return Promise.reject(msg)
    }
}