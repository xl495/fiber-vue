import { POST } from '@/utils/http/index'

export const login = (data: any) => POST('/login', data)