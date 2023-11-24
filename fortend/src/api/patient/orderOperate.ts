import http from '@/utils/request'

//请求科室信息
export function getSection(params: any) {
    return http.request({
        url: '/admin/sectionInfo',
        method: 'get',
        params
    })
}