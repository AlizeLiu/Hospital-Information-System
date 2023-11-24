// import { defineConfig } from 'vite'
// import vue from '@vitejs/plugin-vue'

// // https://vitejs.dev/config/
// export default defineConfig({
//   plugins: [vue()],
// })
import { defineConfig } from 'vite';
import vue from '@vitejs/plugin-vue';
import * as path from 'path';

// https://vitejs.dev/config/
export default defineConfig({
    resolve: {
        //设置别名
        alias: {
            '@': path.resolve(__dirname, 'src')
        }
    },
    plugins: [vue()],
    server: {
        //  服务端口
        port: 8081, //启动端口
        // hmr: {
        //     host: '127.0.0.1',
        //     port: 8080
        // },
        // 设置 https 代理
        proxy: {
            '/api': {
                // target: 'http://xxxxx.com',
                changeOrigin: true,
                rewrite: (path: string) => path.replace(/^\/api/, '')
            }
        },
        css: {
            preprocessorOptions: {
                scss: {
                    additionalData: '@import "@/styles/constant.scss";'
                }
            }
        }
    }
});

