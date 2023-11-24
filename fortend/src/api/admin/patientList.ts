import http from '@/utils/request'

//删除病人操作
export function getDeletePatient(params: any) {
    return http.request({
        url: '/admin/deletePatient',
        method: 'get',
        params
    })
}
// 加载患者列表
export function getFindAllPatients(params: any) {
    return http.request({
        url: '/admin/findAllPatients',
        method: 'get',
        params
    })
}