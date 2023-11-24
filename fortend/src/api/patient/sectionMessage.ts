import http from '@/utils/request'

//选择日期触发时间
export function getFindByTime(params: any) {
    return http.request({
        url: '/arrange/findByTime',
        method: 'get',
        params
    })
}
//选择日期触发时间段
export function getFindNowTime(params: any) {
    return http.request({
        url: '/arrange/findNowTime',
        method: 'get',
        params
    })
}
//请求部门医生信息
export function getFindDoctorBySection(params: any) {
    return http.request({
        url: '/patient/findDoctorBySection',
        method: 'get',
        params
    })
}
//挂号点击确认
export function getAddOrder(params: any) {
    return http.request({
        url: '/patient/addOrder',
        method: 'get',
        params
    })
}
// 获取挂号时间段已剩余票数
export function getFindOrderTime(params: any) {
    return http.request({
        url: '/order/findOrderTime',
        method: 'get',
        params
    })
}