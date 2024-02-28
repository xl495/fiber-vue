import { GET, POST } from '@/utils/http/index'

export const login = (data: any) => POST('/login', data)

export const getUsers = (params: any = {}) => GET('/users', { params })