import http from '@/utils/request'

//请求挂号信息
export function getFindOrderByNull(params: any) {
    return http.request({
        url: '/doctor/findOrderByNull',
        method: 'get',
        params
    })
}