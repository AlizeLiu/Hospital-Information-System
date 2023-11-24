import http from '@/utils/request'
//点击缴费按钮
export function getUpdatePrice(params: any) {
    return http.request({
        url: '/order/updatePrice',
        method: 'get',
        params
    })
}

export function getFindDoctor(params: any) {
    return http.request({
        url: '/admin/findDoctor',
        method: 'get',
        params
    })
}
//请求挂号信息
export function getFindOrderByPid(params: any) {
    return http.request({
        url: '/patient/findOrderByPid',
        method: 'get',
        params
    })
}
// 请求挂号信息
export function getFindOrderInfo(params: any) {
    return http.request({
        url: '/patient/findOrderInfo',
        method: 'get',
        params
    })
}