// [
//     {
//         "path": "/index",
//         "name": "Index",
//         "component": "index",
//         "children": [
//             {
//                 "component": "admin/AdminLayout",
//                 "path": "/adminLayout",
//                 "meta": {
//                     "title": "首页大屏",
//                     "requireAuth": true
//                 }
//             },
//             {
//                 "path": "/doctorList",
//                 "component": "admin/DoctorList",
//                 "meta": {
//                     "title": "医生信息管理",
//                     "requireAuth": true,  // 该路由项需要权限校验
//                 }
//             },
//             {
//                 "path": "/patientList",
//                 "component": "admin/PatientList",
//                 "meta": {
//                     "title": "患者信息管理",
//                     "requireAuth": true,  // 该路由项需要权限校验
//                 }
//             },
//             {
//                 "path": "/orderOperate",
//                 "component": "patient/OrderOperate",
//                 "meta": {
//                     "title": "预约挂号",
//                     "requireAuth": true,  // 该路由项需要权限校验
//                 }
//             },
//             {
//                 "path": "/sectionMessage",
//                 "component": "patient/SectionMessage",
//                 "meta": {
//                     "title": "挂号信息",
//                     "requireAuth": true,  // 该路由项需要权限校验
//                 }
//             },
//             {
//                 "path": "/myOrder",
//                 "component": "patient/MyOrder",
//                 "meta": {
//                     "title": "我的挂号",
//                     "requireAuth": true,  // 该路由项需要权限校验
//                 }
//             }

//         ]
//     }
// ]

export const arrDoctor = [
    {
        component: "admin/AdminLayout",
        path: "/adminLayout",
        meta: {
            title: "首页大屏",
            requireAuth: true
        }
    },
    {
        path: "/orderToday",
        component: 'doctor/OrderToday',
        meta: {
            title: '今日挂号',
            requireAuth: true,  // 该路由项需要权限校验
        },
    },
    {
        path: "/dealOrder",
        component: 'doctor/DealOrder',
        meta: {
            title: '处理挂号',
            requireAuth: true,  // 该路由项需要权限校验
            hidden: true
        },
    },
    {
        path: "/doctorOrder",
        component: 'doctor/DoctorOrder',
        meta: {
            title: '历史挂号',
            requireAuth: true,  // 该路由项需要权限校验
        },
    },
]