import http from '@/utils/request'
//点击修改医生信息
export function getModifyDoctor(params: any) {
    return http.request({
        url: '/admin/modifyDoctor',
        method: 'post',
        params
    })
}
// 根据id查询医生信息
export function getFindDoctor(params: any) {
    return http.request({
        url: '/admin/findDoctor',
        method: 'get',
        params
    })
}
//删除医生操作
export function getDeleteDoctor(params: any) {
    return http.request({
        url: '/admin/deleteDoctor',
        method: 'get',
        params
    })
}
// 增加医生
export function getAddDoctor(params: any) {
    return http.request({
        url: '/admin/addDoctor',
        method: 'get',
        params
    })
}
// 加载医生列表
export function getFindAllDoctors(params: any) {
    return http.request({
        url: '/admin/findAllDoctors',
        method: 'get',
        params
    })
}