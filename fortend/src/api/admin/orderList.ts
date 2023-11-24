import http from '@/utils/request'

//删除挂号操作
export function getDeleteOrder(params: any) {
    return http.request({
        url: '/admin/deleteOrder',
        method: 'get',
        params
    })
}
// 加载医生列表
export function getFindAllOrders(params: any) {
    return http.request({
        url: '/admin/findAllOrders',
        method: 'get',
        params
    })
}