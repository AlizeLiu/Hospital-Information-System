// import * as login from './module/login';
// import * as index from './module/index';

// export default Object.assign({}, login, index);

import http from '../utils/request'
//获取信息
export function getInfo(params: any) {
    return http.request({
        url: '/info',
        method: 'get',
        params
    })
}
// 

