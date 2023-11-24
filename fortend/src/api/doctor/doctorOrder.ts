import http from '@/utils/request'

//请求挂号信息
export function getFindOrderByDid(params: any) {
    return http.request({
        url: '/order/findOrderByDid',
        method: 'get',
        params
    })
}
//请求患者流程
export function getFindStarByDid(params: any) {
    return http.request({
        url: '/order/findStarByDid',
        method: 'get',
        params
    })
}