<template>
  <div>
    <el-card>
      <!-- 面包屑 -->
      <el-breadcrumb separator-class="el-icon-arrow-right">
        <el-breadcrumb-item :to="{ path: '/orderOperate' }">科室选择</el-breadcrumb-item>
        <el-breadcrumb-item>日期选择</el-breadcrumb-item>
        <el-breadcrumb-item>挂号</el-breadcrumb-item>
      </el-breadcrumb>
      <!-- <i class="el-icon-monitor" style="font-size: 22px;"><b>科室门诊</b></i> -->
      <el-divider></el-divider>
      <i class="el-icon-monitor" style="font-size: 22px;"><b> 内科</b></i>
      <ul>
        <li v-for="inter in inters" :key="inter">
          <el-button type="primary" style="margin: 5px;" size="mini" @click="sectionClick(inter)">{{ inter }}</el-button>
        </li>
      </ul>
      <el-divider></el-divider>
      <i class="el-icon-monitor" style="font-size: 22px;"><b> 外科</b></i>
      <ul>
        <li v-for="out in outs" :key="out">
          <el-button type="primary" style="margin: 5px;" size="mini" @click="sectionClick(out)">{{ out }}</el-button>
        </li>
      </ul>
      <el-divider></el-divider>
      <i class="el-icon-monitor" style="font-size: 22px;"><b> 妇产科</b></i>
      <ul>
        <li v-for="woman in women" :key="woman">
          <el-button type="primary" style="margin: 5px;" size="mini" @click="sectionClick(woman)">{{ woman }}</el-button>
        </li>
      </ul>
      <el-divider></el-divider>
      <i class="el-icon-monitor" style="font-size: 22px;"><b> 儿科</b></i>
      <ul>
        <li v-for="kid in kids" :key="kid">
          <el-button type="primary" style="margin: 5px;" size="mini" @click="sectionClick(kid)">{{ kid }}</el-button>
        </li>
      </ul>
      <el-divider></el-divider>
      <i class="el-icon-monitor" style="font-size: 22px;"><b> 五官科</b></i>
      <ul>
        <li v-for="five in fives" :key="five">
          <el-button type="primary" style="margin: 5px;" size="mini" @click="sectionClick(five)">{{ five }}</el-button>
        </li>
      </ul>
      <el-divider></el-divider>
      <i class="el-icon-monitor" style="font-size: 22px;"><b> 中医科</b></i>
      <ul>
        <li v-for="chinese in chineses" :key="chinese">
          <el-button type="primary" style="margin: 5px;" size="mini" @click="sectionClick(chinese)">{{ chinese
          }}</el-button>
        </li>
      </ul>
      <el-divider></el-divider>
      <i class="el-icon-monitor" style="font-size: 22px;"><b> 其他</b></i>
      <ul>
        <li v-for="orther in orthers" :key="orther">
          <el-button type="primary" style="margin: 5px;" size="mini" @click="sectionClick(orther)">{{ orther
          }}</el-button>
        </li>
      </ul>
      <el-divider></el-divider>
    </el-card>
  </div>
</template>
<script setup>
import { ref, reactive, onMounted, } from 'vue'
import { useRouter } from 'vue-router';
import { ElMessage, ElMessageBox } from 'element-plus'
import {Check,Delete, Close, Search, Edit} from '@element-plus/icons-vue'
const inters = reactive([
  "神经内科",
  "呼吸与危重症医学科",
  "内分泌科",
  "消化内科",
  "心血管内科",
  "肾内科",
  "发热门诊",
])
const outs = reactive([
  "手足外科",
  "普通外科",
  "肛肠外科",
  "神经外科",
  "泌尿外科",
  "骨科",
  "烧伤整形外科",
])
const women = reactive(["妇科", "产科"])
const kids = reactive(["儿科", "儿童保健科"])
const fives = reactive(["耳鼻咽喉科", "眼科", "口腔科"])
const chineses = reactive(["中医科"])
const orthers = reactive(["康复医学科", "急诊科", "皮肤性病科", "功能科"])

let router = useRouter()
//点击科室
function sectionClick(sectionOpt) {
  console.log(sectionOpt);
  router.push({ path: '/sectionMessage',query:{sectionOpt} });
}

</script>
<style scoped lang="scss">
ul li {
  display: inline;
}

;

.el-breadcrumb {
  margin: 8px;
}
</style>