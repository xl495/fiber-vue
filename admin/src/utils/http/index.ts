import http, { Response } from './request'

const GET = (url: string, params: any) => {
    return http.get<Response, any>(url, {
        params
    })
}

const POST = (url: string, data: any) => {
    return http.post<Response, any>(url, data)
}

const PUT = (url: string, data: any) => {
    return http.put<any, any>(url, data)
}


const DELETE = (url: string, data: any) => {
    return http.delete<any, any>(url, data)
}


export {
    GET,
    POST,
    PUT,
    DELETE
}