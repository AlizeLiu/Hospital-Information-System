import http from '@/utils/request'
//点击修改药物信息
export function getModifyCheck(params: any) {
    return http.request({
        url: '/check/modifyCheck',
        method: 'get',
        params
    })
}
// 打开修改对话框
export function getFindCheck(params: any) {
    return http.request({
        url: '/check/findCheck',
        method: 'get',
        params
    })
}
//删除检查操作
export function getDeleteCheck(params: any) {
    return http.request({
        url: 'check/deleteCheck',
        method: 'get',
        params
    })
}
// 点击增加确认按钮
export function getAddCheck(params: any) {
    return http.request({
        url: '/check/addCheck',
        method: 'get',
        params
    })
}
// 加载医生列表
export function getFindAllChecks(params: any) {
    return http.request({
        url: '/check/findAllChecks',
        method: 'get',
        params
    })
}