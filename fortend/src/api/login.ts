import http from '../utils/request'
//提交表单
export function getlogin(params: any) {
    return http.request({
        url: '/admin/login',
        method: 'get',
        params
    })
}
// 
//点击注册确认按钮
export function getRegist(params: any) {
    return http.request({
        url: '/patient/addPatient',
        method: 'get',
        params
    })
}
//获取菜单
export function getMenu(params: any) {
    return http.request({
        url: '/login/getMenu',
        method: 'get',
        params
    })
}
