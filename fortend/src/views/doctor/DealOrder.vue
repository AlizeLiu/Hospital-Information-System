<template>
  <div>
    <el-card>
      <el-row>
        <el-col :span="7">
          挂号单号：<el-input disabled v-model="oId" class="orderInput"></el-input>
        </el-col>
        <el-col :span="7">
          患者账号：<el-input disabled v-model="pId" class="orderInput"></el-input></el-col>
        <el-col :span="7">
          患者姓名：<el-input disabled v-model="pName" class="orderInput"></el-input>
        </el-col>
        <el-col :span="3">
          <el-button type="success" @click="submitClick">提交</el-button>
        </el-col>
      </el-row>
      <el-row>
        <el-col :span="7">
          患者性别：<el-input disabled v-model="pGender" class="orderInput"></el-input>
        </el-col>
        <el-col :span="7">
          联系方式：<el-input disabled v-model="pPhone" class="orderInput"></el-input></el-col>
        <el-col :span="7">
          医生姓名：<el-input disabled v-model="dName" class="orderInput"></el-input>
        </el-col>
        <el-col :span="3">
          <el-button type="success" @click="openReason">医嘱编写</el-button>
        </el-col>
      </el-row>
      <!-- 药物表格 -->
      <el-row>
        <el-col :span="12">
          <el-input v-model="queryDrug" placeholder="请输入名称查询" class="drugInput">
            <template #append>
              <el-button @click="requestDrug" :icon="Search"></el-button>
            </template>
          </el-input>
          <el-table :data="drugData" stripe border>
            <el-table-column label="编号" prop="drId"></el-table-column>
            <el-table-column label="名称" prop="drName"></el-table-column>
            <el-table-column label="剩余数量" prop="drNumber"></el-table-column>
            <el-table-column label="单位" prop="drUnit"></el-table-column>
            <el-table-column label="单价/元" prop="drPrice"></el-table-column>
            <el-table-column label="操作" width="120" fixed="right">
              <template #default="scope">
                <el-button type="success" style="font-size: 14px;" @click="addDrug(scope.row.drId)" :icon="Plus">
                  增加</el-button>
              </template>
            </el-table-column>
          </el-table>
          <!-- 分页 -->
          <el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange" background
            layout="total, sizes, prev, pager, next" :total="total" :page-size="size" :page-sizes="[1, 2, 4, 8, 16]">
          </el-pagination>
          <el-row></el-row>
        </el-col>
        <!-- 右边已选择的药物 -->
        <el-col :span="12" class="drugRigth">
          <el-table stripe border :data="drugBuyData" class="rigthTable">
            <el-table-column label="编号" prop="drId"></el-table-column>
            <el-table-column label="名称" prop="drName"></el-table-column>
            <el-table-column label="单价/元" prop="drPrice"></el-table-column>
            <el-table-column label="数量" prop="drNum"></el-table-column>
            <el-table-column label="小记" prop="drSum"></el-table-column>
            <el-table-column label="操作" width="120" fixed="right">
              <template #default="scope">
                <el-button type="danger" style="font-size: 14px;" :icon="Minus" @click="deleteDrug(scope.row.drId)">
                  移除</el-button>
              </template>
            </el-table-column>
          </el-table>
          <el-tag>共计：{{ drugTotalPrice }}元</el-tag>
        </el-col>
      </el-row>
      <!-- 检查项目编写 -->
      <el-row>
        <el-col :span="12">
          <el-input v-model="queryCheck" placeholder="请输入名称查询" class="drugInput">
            <template #append>
              <el-button :icon="Search" @click="requestCheck"></el-button>
            </template>
          </el-input>
          <el-table stripe border :data="checkData">
            <el-table-column label="编号" prop="chId"></el-table-column>
            <el-table-column label="项目" prop="chName"></el-table-column>
            <el-table-column label="价格/元" prop="chPrice"></el-table-column>
            <el-table-column label="操作" width="120" fixed="right">
              <template #default="scope">
                <el-button type="success" style="font-size: 14px;" :icon="Plus" @click="addCheck(scope.row.chId)">
                  增加</el-button>
              </template>
            </el-table-column>
          </el-table>
          <!-- 分页 -->
          <el-pagination @size-change="checkSizeChange" @current-change="checkCurrentChange" background
            layout="total, prev, pager, next" :total="checkTotal" :page-size="checkSize">
          </el-pagination>
        </el-col>
        <!-- 右边已选择的检查 -->
        <el-col :span="12" class="drugRigth">
          <el-table stripe border class="rigthTable" :data="checkBuyData">
            <el-table-column label="编号" prop="chId"></el-table-column>
            <el-table-column label="项目" prop="chName"></el-table-column>
            <el-table-column label="价格/元" prop="chPrice"></el-table-column>
            <el-table-column label="操作" width="120" fixed="right">
              <template #default="scope">
                <el-button type="danger" style="font-size: 14px;" :icon="Minus" @click="deleteCheck(scope.row.chId)">
                  移除</el-button>
              </template>
            </el-table-column>
          </el-table>
          <el-tag>共计：{{ checkTotalPrice }}元</el-tag>
        </el-col>
      </el-row>
    </el-card>

    <!-- 病因编写对话框 -->
    <el-dialog title="病因编写" v-model="reasonFormVisible">
      <el-input type="textarea" :rows="8" placeholder="请输入内容" v-model="reason">
      </el-input>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="reasonFormVisible = false" style="font-size: 18px;"><i class="el-icon-close"
              style="font-size: 20px;"></i> 取 消</el-button>
          <el-button type="primary" @click="holdReason">保存</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>
<script setup>
import { ref, reactive, onMounted, computed } from 'vue'
import request from '@/utils/request.ts'
import { useRouter, useRoute } from 'vue-router';
import { ElMessage, ElMessageBox } from 'element-plus'
import { Check, Delete, Close, Search, Edit, List, ChatRound, Memo, Plus, Minus } from '@element-plus/icons-vue'
import { getReduceDrugNumber, getUpdateOrder, getFindCheck, getFindAllChecks, getAddDrug, getFindAllDrugs, getFindPatientById } from '@/api/doctor/dealOrder.ts'
let oId = ref(1)
let pId = ref(1)
let dId = ref(0)
let pName = ref("aa")
let pGender = ref("as")
let pPhone = ref("asd")
let dName = ref("")
// 药物表格
let drugData = ref([
  {
    drId: 1,
    drName: '青霉素',
    drNumber: 99,
    drUnit: '盒',
    drPrice: 40,
  },
  {
    drId: 2,
    drName: '阿莫西林',
    drNumber: 46,
    drUnit: '盒',
    drPrice: 24,
  },
  {
    drId: 3,
    drName: '头孢',
    drNumber: 36,
    drUnit: '盒',
    drPrice: 16,
  },
  {
    drId: 4,
    drName: '999感冒灵',
    drNumber: 54,
    drUnit: '盒',
    drPrice: 16,
  },
])
let size = ref(4)
let pageNumber = ref(1)
let total = ref(3)
let queryDrug = ref("")
// 右边已选择的药物
let drugBuyData = ref([]
  // [
  //   {
  //     drId: 1,
  //     drName: '青霉素',
  //     drNum: 1,
  //     drUnit: '盒',
  //     drPrice: 40,
  //     drSum: 40,
  //   },
  //   {
  //     drId: 2,
  //     drName: '阿莫西林',
  //     drNum: 2,
  //     drUnit: '盒',
  //     drPrice: 24,
  //     drSum: 48,
  //   },
  //   {
  //     drId: 3,
  //     drName: '头孢',
  //     drNum: 1,
  //     drUnit: '盒',
  //     drPrice: 16,
  //     drSum: 16,
  //   },
  // ]
)
let drugTotalPrice = ref(0)
// let drugTotalPrice = computed(() => {
//   let sum = null;
//   drugBuyData.map((item) => {
//     sum = sum + item.drSum
//     console.log(sum);
//   })
//   return sum
// })  //
let reason = ref("")   
// 检查项目编写
let checkData = reactive([
  {
    chId: 1,
    chName: 'B超',
    chPrice: 180,
  },
  {
    chId: 2,
    chName: '核磁共振',
    chPrice: 200,
  },
  {
    chId: 3,
    chName: 'CT',
    chPrice: 425,
  },
  {
    chId: 4,
    chName: '心电图',
    chPrice: 100,
  },
  {
    chId: 5,
    chName: '验血',
    chPrice: 80,
  },
])
let queryCheck = ref("")
let checkTotal = ref(3)
let checkSize = ref(4)
let checkPageNumber = ref(1)
let checkTotalPrice = ref(0)
let checkBuyData = reactive([
  // {
  //   chId: 1,
  //   chName: 'B超',
  //   chPrice: 180,
  // },
  // {
  //   chId: 2,
  //   chName: '核磁共振',
  //   chPrice: 200,
  // },
  // {
  //   chId: 3,
  //   chName: 'CT',
  //   chPrice: 425,
  // },
])    // 病因编写对话框显示参数
let reasonFormVisible = ref(false)

let route = useRoute()
let router = useRouter()
//根据id减少药物数量
function reduceDrugNumber(drId, usedNumber) {
  getReduceDrugNumber({
    drId: drId,
    usedNumber: usedNumber,
  })
    .then((res) => {
      if (res.status !== 200) {
        ElMessage.error("药物数量不足！！");
      }
      ElMessage.success("减少药物成功！！");
    });
}
//点击提交按钮
function submitClick() {
  for (let i = 0; i < drugBuyData.value.length; i++) {
    reduceDrugNumber(
      drugBuyData.value[i].drId,
      drugBuyData.value[i].drNum
    );
  }
  let data = {
    oId: oId.value,
    pId: pId.value,
    dId: dId.value,
    oRecord: reason.value,
    oDrug: drugTotalPrice.value,
    oCheck: checkTotalPrice.value,
    oTotalPrice: dataPackage().oTotalPrice,
    oCheckBuyData:checkBuyData,
    oDrugBuyData:drugBuyData.value
  };
  console.log('data--------',data);
  getUpdateOrder(data)
    .then((res) => {
      if (res.status !== 200) {
        ElMessage.error("请求信息错误");
        return;
      }
      ElMessage.success("提交成功！请通知患者登录系统自助缴费！");
      router.push("/orderToday");
      console.log(res.data);
    })
    .catch((err) => {
      console.error(err);
    });
}
//封装数据
function dataPackage() {
  let oDrug = "";
  let oCheck = "";
  let oTotalPrice = 0;
  // let oId = oId.value;
  let oRecord = reason.value;
  // for (let i = 0; i < drugBuyData.length; i++) {
  //   oDrug +=
  //     drugBuyData.value[i].drPrice +
  //     "(元)*" ;
  // }
  // for (let i = 0; i < checkBuyData.length; i++) {
  //   oCheck +=checkBuyData[i].chPrice;
  // }
  oCheck = " 项目总价" + checkTotalPrice.value + "元 ";
  oDrug = " 药物总价" + drugTotalPrice.value + "元 ";
  oTotalPrice = checkTotalPrice.value + drugTotalPrice.value;
  return {  oRecord, oDrug, oCheck, oTotalPrice };
}
//点击病因保存按钮
function holdReason() {
  reasonFormVisible.value = false;
  ElMessage.success("信息保存成功！");
}
//打开病因编写对话框
function openReason() {
  reasonFormVisible.value = true;
}
//检查列表点击移除按钮
function deleteCheck(chId) {
  for (let i = 0; i < checkBuyData.length; i++) {
    if (checkBuyData[i].chId === chId) {
      checkTotalPrice.value -= checkBuyData[i].chPrice; //药物价格总计
      checkBuyData.splice(i, 1); //！！！！！！删除数组中下标为i的元素
    }
  }
}
//检查列表点击增加按钮
function addCheck(chId) {
  getFindCheck({
    chId: chId,
  })
    .then((res) => {
      if (res.status != 200) return ElMessage.error("信息请求失败");
      //后端传过来的是对象，表格绑定的数组
      checkBuyData.push({
        chId: res.data.chId,
        chPrice: res.data.chPrice,
        chName: res.data.chName,
      });
      checkTotalPrice.value += res.data.chPrice; //药物价格总计
      console.log(res);
    });
}
//药物页面大小切换时触发
function checkSizeChange(size1) {
  checkSize.value = size1;
  requestCheck();
}
//检查页数切换时触发
function checkCurrentChange(num) {
  checkPageNumber.value = num;
  requestCheck();
}
//请求检查项目
function requestCheck() {
  getFindAllChecks({
    size: checkSize.value,
    pageNumber: checkPageNumber.value,
    query: queryCheck.value,
  })
    .then((res) => {
      if (res.status != 200) ElMessage.error("获取信息失败");
      checkData = res.data.checks;
      checkTotal.value = res.data.total;
      console.log(res);
    });
}
/**
 * 此处逻辑较复杂，容易出现错误
 */
//药物列表点击移除按钮
function deleteDrug(drId) {
  for (let i = 0; i < drugBuyData.value.length; i++) {
    if (drugBuyData.value[i].drId === drId) {
      for (let j = 0; j < drugData.value.length; j++) {
        if (drugData.value[j].drId === drId) drugData.value[j].drNumber += 1;
      }
      drugBuyData.value[i].drNum -= 1;
      drugBuyData.value[i].drSum =
        drugBuyData.value[i].drPrice * drugBuyData.value[i].drNum;
      drugTotalPrice.value -= drugBuyData.value[i].drPrice; //药物价格总计
      if (drugBuyData.value[i].drNum === 0) drugBuyData.value.splice(i, 1); //！！！！！！删除数组中下标为i的元素
    }
  }
}
//药物列表点击增加按钮
function addDrug(drId) {
  getAddDrug({
    drId: drId,
  })
    .then((res) => {
      if (res.status != 200) return ElMessage.error("信息请求失败");
      //后端传过来的是对象，表格绑定的数组
      for (let i = 0; i < drugBuyData.value.length; i++) {
        if (drugBuyData.value[i].drId === res.data.drId) {
          for (let j = 0; j < drugData.value.length; j++) {
            if (
              drugData.value[j].drId === res.data.drId &&
              drugData.value[j].drNumber > 0
            ) {
              drugData.value[j].drNumber -= 1;
              drugBuyData.value[i].drNum += 1;
              drugBuyData.value[i].drSum =
                drugBuyData.value[i].drPrice * drugBuyData.value[i].drNum;
              drugTotalPrice.value += drugBuyData.value[i].drPrice; //药物价格总计
              //return;
            }
          }
          return;
        }
      }
      for (let j = 0; j < drugData.value.length; j++) {
        if (
          drugData.value[j].drId === res.data.drId &&
          drugData.value[j].drNumber <= 0
        )
          return;
      }
      drugBuyData.value.push({
        drId: res.data.drId,
        drPrice: res.data.drPrice,
        drName: res.data.drName,
        drNum: 1,
        drSum: res.data.drPrice,
      });
      for (let j = 0; j < drugData.value.length; j++) {
        if (drugData.value[j].drId === res.data.drId)
          drugData.value[j].drNumber -= 1;
      }
      drugTotalPrice.value += res.data.drPrice; //药物价格总计

      console.log(drugBuyData);
    });
}
//药物页面大小切换时触发
function handleSizeChange(size1) {
  size.value = size1;
  requestDrug();
}
//药物页数切换时触发
function handleCurrentChange(num) {
  pageNumber.value = num;
  requestDrug();
}

//获取药物列表
function requestDrug() {
  getFindAllDrugs({
    size: size.value,
    pageNumber: pageNumber.value,
    query: queryDrug.value,
  })
    .then((res) => {
      if (res.status != 200)
        ElMessage.error("获取信息失败");
      // console.log(res);
      drugData.value = res.data.drugs;
      total.value = res.data.total;

      console.log(res);
    });
}
//获取患者信息
function requestPatient() {
  getFindPatientById({
    pId: pId.value,
  })
    .then((res) => {
      console.log(res);
      if (res.status != 200) ElMessage.error("获取信息失败");
      pName.value = res.data.pName;
      pGender.value = res.data.pGender;
      pPhone.value = res.data.pPhone;
    });
}
//token解码
// function tokenDecode(token) {
//   if (token !== null) return jwtDecode(token);
// }
onMounted(() => {
  dName.value = route.query.dName
  // dId.value = this.tokenDecode(getToken()).dId;
  oId.value = route.query.oId;
  pId.value = route.query.pId;
  requestPatient();
  requestDrug();
  requestCheck();
})
</script>
<style lang="scss">
// 让文字居中
.drugRigth {
  text-align: center;
}

.el-tag {
  margin: 8px;
}

.rigthTable {
  margin-top: 56px;
  margin-left: 8px;
}

.drugInput {
  margin-top: 8px;
  margin-bottom: 8px;
}

.el-row {
  margin: 5px;
}

.orderInput {
  width: 240px;
}

.el-pagination {
  margin: 8px;
}
</style>