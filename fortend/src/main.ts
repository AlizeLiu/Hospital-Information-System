import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
// 图标
import * as ElementPlusIconsVue from '@element-plus/icons-vue'

// import store from './store'
import router from '@/router';
import {SetRouter} from '@/router/index.ts'
// 创建vue实例
const app = createApp(App);

for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
  app.component(key, component)
}
SetRouter()

app.use(router);

app.use(ElementPlus);
// 挂载实例
app.mount('#app');

