import http from '@/utils/request'
//get有值
export function getPeople(params: any) {
    return http.request({
        url: '/order/orderPeople',
        method: 'get',
        params
    })
}
export function getBed(params: any) {
    return http.request({
        url: '/bed/bedPeople',
        method: 'get',
        params
    })
}