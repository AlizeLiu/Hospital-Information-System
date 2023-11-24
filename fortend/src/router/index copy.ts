import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router';
import { getToken } from '@/utils/storage.ts'
import { getMenu } from '@/api/login.ts'

import NotFound from "@/views/404.vue";
import AdminLayout from "@/views/admin/AdminLayout.vue";
import DoctorList from "@/views/admin/DoctorList.vue";
import PatientList from "@/views/admin/PatientList.vue";
import OrderList from "@/views/admin/OrderList.vue";
import DrugList from "@/views/admin/DrugList.vue";
import CheckList from "@/views/admin/CheckList.vue";
import BedList from "@/views/admin/BedList.vue";

import Patient from "@/views/patient/Patient.vue";
import OrderOperate from "@/views/patient/OrderOperate.vue";
import SectionMessage from "@/views/patient/SectionMessage.vue";
import MyOrder from "@/views/patient/MyOrder.vue";

import Doctor from "@/views/doctor/Doctor.vue";
import OrderToday from "@/views/doctor/OrderToday.vue";
import DealOrder from "@/views/doctor/DealOrder.vue";
import DoctorOrder from "@/views/doctor/DoctorOrder.vue";

const routes: Array<RouteRecordRaw> = [
    { path: '/', redirect: '/index' },
    {
        path: '/login',
        name: 'Login',
        component: () => import('@/views/login.vue')
    },
    {
        path: '/index',
        name: 'Index',
        meta: {
            title: '首页',
            keepAlive: true,
            requireAuth: true
        },
        component: () => import('@/views/index.vue'),
            children: [
                {
                    path: "/adminLayout",
                    component: AdminLayout,
                    meta: {
                        title: '首页大屏',
                        requireAuth: true,  // 该路由项需要权限校验
                    },
                },
                {
                    path: "/doctorList",
                    component: DoctorList,
                    meta: {
                        title: '医生信息管理',
                        requireAuth: true,  // 该路由项需要权限校验
                    },
                },
        {
            path: "/patientList",
            component: PatientList,
            meta: {
                title: '患者信息管理',
                requireAuth: true,  // 该路由项需要权限校验
            },
        },
        //         // {
        //         //     path: "/orderList",
        //         //     component: OrderList,
        //         //     meta: {
        //         //         title: '挂号信息管理',
        //         //         requireAuth: true,  // 该路由项需要权限校验
        //         //     },
        //         // },
        //         // {
        //         //     path: "/drugList",
        //         //     component: DrugList,
        //         //     meta: {
        //         //         title: '药物信息管理',
        //         //         requireAuth: true,  // 该路由项需要权限校验
        //         //     },
        //         // },
        //         // {
        //         //     path: "/checkList",
        //         //     component: CheckList,
        //         //     meta: {
        //         //         title: '检查项目管理',
        //         //         requireAuth: true,  // 该路由项需要权限校验
        //         //     },

        //         // },
        //         // {
        //         //     path: "/bedList",
        //         //     component: BedList,
        //         //     meta: {
        //         //         title: '病床信息管理',
        //         //         requireAuth: true,  // 该路由项需要权限校验
        //         //     },

        //         // },
        //         // {
        //         //     path: "/orderOperate",
        //         //     component: OrderOperate,
        //         //     meta: {
        //         //         title: '预约挂号',
        //         //         requireAuth: true,  // 该路由项需要权限校验
        //         //     },
        //         // },
        //         // {
        //         //     path: "/sectionMessage",
        //         //     component: SectionMessage,
        //         //     meta: {
        //         //         title: '挂号信息',
        //         //         requireAuth: true,  // 该路由项需要权限校验
        //         //     },
        //         // },
        //         // {
        //         //     path: "/myOrder",
        //         //     component: MyOrder,
        //         //     meta: {
        //         //                 title: '我的挂号',
        //         //                 requireAuth: true,  // 该路由项需要权限校验
        //         //             },
        //         // },
        //         // {
        //         //     path: "/orderToday",
        //         //     component: OrderToday,
        //         //     meta: {
        //         //         title: '今日挂号列表',
        //         //         requireAuth: true,  // 该路由项需要权限校验
        //         //     },
        //         // },
        //         // {
        //         //     path: "/dealOrder",
        //         //     component: DealOrder,
        //         //     meta: {
        //         //         title: '处理挂号',
        //         //         requireAuth: true,  // 该路由项需要权限校验
        //         //     },
        //         // },
        //         // {
        //         //     path: "/doctorOrder",
        //         //     component: DoctorOrder,
        //         //     meta: {
        //         //         requireAuth: true,  // 该路由项需要权限校验
        //         //     },
        //         // },
        //     ]
    }
    // {
    //     path: "/patient",
    //     component: Patient,
    //     meta: {
    //         requireAuth: true,  // 该路由项需要权限校验
    //     },
    //     children: [
            {
                path: "/orderOperate",
                component: OrderOperate,
                meta: {
                    title: '预约挂号',
                    requireAuth: true,  // 该路由项需要权限校验
                },
            },
            {
                path: "/sectionMessage",
                component: SectionMessage,
                meta: {
                    title: '挂号信息',
                    requireAuth: true,  // 该路由项需要权限校验
                },
            },
            {
                path: "/myOrder",
                component: MyOrder,
                meta: {
                    title: '我的挂号',
                    requireAuth: true,  // 该路由项需要权限校验
                },
            },
    //     ]
    // },
    // {
    //     path: "/doctor",
    //     component: Doctor,
    //     meta: {
    //         requireAuth: true,  // 该路由项需要权限校验
    //     },
    //     children: [
            {
                path: "/orderToday",
                component: OrderToday,
                meta: {
                    title: '今日挂号列表',
                    requireAuth: true,  // 该路由项需要权限校验
                },
            },
            {
                path: "/dealOrder",
                component: DealOrder,
                meta: {
                    title: '处理挂号',
                    requireAuth: true,  // 该路由项需要权限校验
                },
            },
            {
                path: "/doctorOrder",
                component: DoctorOrder,
                meta: {
                    requireAuth: true,  // 该路由项需要权限校验
                },
            },
    //         // {
    //         //     path: "/inBed",
    //         //     component: InBed,
    //         //     meta: {
    //         //         requireAuth: true,  // 该路由项需要权限校验
    //         //     },
    //         // },
    //         // {
    //         //     path: "/doctorCard",
    //         //     component: DoctorCard,
    //         //     meta: {
    //         //         requireAuth: true,  // 该路由项需要权限校验
    //         //     },
    //         // }
    //     ]
    // }
]

const arrAdmin = [
    {
        children: [
            {
                component: "admin/AdminLayout",
                path: "/adminLayout",
                meta: {
                    title: "首页大屏",
                    requireAuth: true
                }
            },
            {
                path: "/doctorList",
                component: 'admin/DoctorList',
                meta: {
                    title: '医生信息管理',
                    requireAuth: true,  // 该路由项需要权限校验
                },
            },
            {
                path: "/patientList",
                component: 'admin/PatientList',
                meta: {
                    title: '患者信息管理',
                    requireAuth: true,  // 该路由项需要权限校验
                },
            },
        ]
    },
]

const router = createRouter({
    history: createWebHistory(),
    routes
});
let menu: any = null;
//使用vite懒加载
let Module = import.meta.glob("@/views/**/*.vue");
//解析菜单
function CompilerMenu(arr: Array<any>) {
    if (!arr.length) {
        return;
    }
    arr.forEach((item) => {
        //定义对象
        let rts = {
            name: item.name,
            path: item.path,
            meta: item.meta,
            component: item.component,
        };
        //如果存在子集进行递归解析
        if (item.children && item.children.length) {
            CompilerMenu(item.children);
        }
        console.log(rts);
        //如果没有子集 证明为路由层
        if (!item.children) {
            //实现组件懒加载
            let paths = loadComponent(item.component);
            console.log(paths);
        }

        //懒加载组件
        function loadComponent(url: string) {
            let path = Module[`/src/views/${url}.vue`];
            return path;
        }
        //如果没有子集 证明为路由层
        if (!item.children) {
            //实现组件懒加载
            let paths = loadComponent(item.component);
            rts.component = paths;
            console.log(rts);
            //添加动态路由
            router.addRoute('/index', rts);
        }
    });
}
router.beforeEach((to, from, next) => {
    if (to.meta.requireAuth) {
        const token = getToken();
        if (token !== null) {
            if (!menu) {
                getMenu({ role: 'admin' }).then((res: any) => {
                    console.log(res);
                    CompilerMenu(arrAdmin)
                    console.log('---------------------------');
                    console.log(router.getRoutes(), router.hasRoute("Housepay"));
                })
            }
            //直接放行
            next();
        } else {
            next("/login");
        }
    }
    else {
        next();
    }
})
console.log(router.getRoutes(), router.hasRoute("Housepay"));
export default router;
