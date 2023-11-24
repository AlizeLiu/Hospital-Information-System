import http from '@/utils/request'
//点击修改药物信息
export function getModifyDrug(params: any) {
    return http.request({
        url: '/drug/modifyDrug',
        method: 'get',
        params
    })
}
// 打开修改对话框
export function getFindDrug(params: any) {
    return http.request({
        url: '/drug/findDrug',
        method: 'get',
        params
    })
}
//删除药物操作
export function getDeleteDrug(params: any) {
    return http.request({
        url: 'drug/deleteDrug',
        method: 'get',
        params
    })
}
// 增加医生
export function getAddDrug(params: any) {
    return http.request({
        url: '/drug/addDrug',
        method: 'get',
        params
    })
}
// 加载医生列表
export function getFindAllDrugs(params: any) {
    return http.request({
        url: '/drug/findAllDrugs',
        method: 'get',
        params
    })
}