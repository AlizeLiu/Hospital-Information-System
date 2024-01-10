import http from '@/utils/request'
//根据id减少药物数量
export function getReduceDrugNumber(params: any) {
    return http.request({
        url: '/drug/reduceDrugNumber',
        method: 'get',
        params
    })
}
//点击提交按钮
export function getUpdateOrder(params: any) {
    return http.request({
        url: '/order/updateOrder',
        method: 'post',
        data: params
    })
}
//检查列表点击增加按钮
export function getFindCheck(params: any) {
    return http.request({
        url: '/check/findCheck',
        method: 'get',
        params
    })
}
//请求检查项目
export function getFindAllChecks(params: any) {
    return http.request({
        url: '/check/findAllChecks',
        method: 'get',
        params
    })
}
// 药物列表点击增加按钮
export function getAddDrug(params: any) {
    return http.request({
        url: '/drug/addDrug',
        method: 'get',
        params
    })
}
// 获取药物列表
export function getFindAllDrugs(params: any) {
    return http.request({
        url: '/drug/findAllDrugs',
        method: 'get',
        params
    })
}
// 获取患者信息
export function getFindPatientById(params: any) {
    return http.request({
        url: '/doctor/findPatientById',
        method: 'get',
        params
    })
}