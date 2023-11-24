import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router';
import { getToken } from '@/utils/storage.ts'
import { getMenu } from '@/api/login.ts'
import { arrAdmin } from './admin.ts'
import { arrPatient } from './patient.ts'
import { arrDoctor } from './doctor.ts'
import { arrList } from './list.ts'

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
    }
    
]

const router = createRouter({
    history: createWebHistory(),
    routes
});
let menu: any = localStorage.getItem('_menu') || null;
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
        //懒加载组件
        function loadComponent(url: string) {
            let path = Module[`/src/views/${url}.vue`];
            console.log('url', url);

            return path;
        }
        //如果没有子集 证明为路由层
        if (!item.children) {
            //实现组件懒加载
            let paths = loadComponent(item.component);
            rts.component = paths;
            // console.log(paths);
            // console.log(rts);
            //添加动态路由
            router.addRoute('Index', rts);
        }
    });

}
router.beforeEach((to, from, next) => {
    if (to.meta.requireAuth) {
        const token = getToken();
        if (token !== null) {
            if (!menu) {
                let role=localStorage.getItem(('role'))                
                getMenu({ role }).then((res: any) => {
                    console.log(res);
                    CompilerMenu(arrList)
                    console.log('---------------------------');
                    menu = router.getRoutes()
                    console.log(router.getRoutes(), router.hasRoute("Housepay"));
                    // localStorage.setItem()
                    // console.log(menu);
                    // let arrList = JSON.stringify(arrAdmin)
                    // localStorage.setItem('list', arrList)
                })
                // if(role=='admin'){
                //     CompilerMenu(arrAdmin)
                //     menu = router.getRoutes()
                // }else if(role=='doctor'){
                //     CompilerMenu(arrDoctor)
                //     menu = router.getRoutes()
                // }else if(role=='patient'){
                //     CompilerMenu(arrPatient)
                //     menu = router.getRoutes()
                // }
            }
            //直接放行
            next();
        } else {
            next("/login");
        }
    }
    else {
        // console.log(router.getRoutes());
        next();
    }
})

export default router;
export function SetRouter() {
    CompilerMenu(arrList)
    menu = router.getRoutes()
    // let arr = JSON.parse(localStorage.getItem('list') || '')
    // let role=localStorage.getItem(('role'))       
    // if(role=='admin'){
    //     CompilerMenu(arrAdmin)
    //     menu = router.getRoutes()
    // }else if(role=='doctor'){
    //     CompilerMenu(arrDoctor)
    //     menu = router.getRoutes()
    // }else if(role=='patient'){
    //     CompilerMenu(arrPatient)
    //     menu = router.getRoutes()
    // }
    // console.log('---------------------------');
    console.log(router.getRoutes(), router.hasRoute("Housepay"));
}
