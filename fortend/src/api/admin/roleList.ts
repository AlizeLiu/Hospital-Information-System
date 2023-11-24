import http from '@/utils/request'

//添加角色列表
export function getAddRole(params: any) {
    return http.request({
        url: '/admin/addRole',
        method: 'post',
        params
    })
}
// 加载菜单列表
export function getFindRoleList(params: any) {
    return http.request({
        url: '/admin/roleList',
        method: 'get',
        params
    })
}
// 删除菜单列表
export function getDeleteRole(params: any) {
    return http.request({
        url: '/admin/deleteRole',
        method: 'get',
        params
    })
}