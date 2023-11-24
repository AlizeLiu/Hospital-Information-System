import http from '@/utils/request'

//添加菜单列表
export function getAddMenu(params: any) {
    return http.request({
        url: '/admin/addMenu',
        method: 'post',
        params
    })
}
// 加载菜单列表
export function getFindMenuList(params: any) {
    return http.request({
        url: '/admin/menuList',
        method: 'get',
        params
    })
}
// 删除菜单列表
export function getDeleteMenu(params: any) {
    return http.request({
        url: '/admin/deleteMenu',
        method: 'get',
        params
    })
}