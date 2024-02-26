
import axios, { AxiosResponse } from 'axios'
import { handleResponse } from './transform'

export interface Response {
    code: number
    data: any
    msg: string
}
// 1. create an axios instance
const instance = axios.create({
    // 获取vite环境变量
    baseURL: import.meta.env.VITE_APP_BASE_API as string,
    timeout: 150000,
    withCredentials: true
})

// 2. request interceptor
instance.interceptors.request.use(config => {
    return config
}, error => {
    console.error('error = ', error)
    Promise.reject(error)
})

// 3. response interceptor
instance.interceptors.response.use((response): AxiosResponse<Response> => {
    const data = handleResponse(response)
    return data
}, error => {
    return Promise.reject(error)
})

export default instance
