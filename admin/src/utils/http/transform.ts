import { AxiosResponse } from "axios"

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
    } else {
        handleError('error', msg)
        return Promise.reject(msg)
    }
}