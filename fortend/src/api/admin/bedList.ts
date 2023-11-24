import http from '@/utils/request'
//get有值
export function getEmptyBed(params: any) {
    return http.request({
        url: '/bed/emptyBed',
        method: 'get',
        params
    })
}
export function getDeleteBed(params: any) {
    return http.request({
        url: '/bed/deleteBed',
        method: 'get',
        params
    })
}
export function getAddBed(params: any) {
    return http.request({
        url: '/bed/addBed',
        method: 'get',
        params
    })
}
export function getFindAllBeds(params: any) {
    return http.request({
        url: '/bed/findAllBeds',
        method: 'get',
        params
    })
}