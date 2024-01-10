<template>
    <div>
        <el-card>
            <!-- 搜索栏 -->
            <el-row type="flex">
                <el-col :span="6">
                    <el-input v-model="query" placeholder="请输入患者id查询">
                        <template #append>
                            <el-button :icon="Search" @click="requestOrders"></el-button>
                        </template>
                    </el-input>
                </el-col>
            </el-row>
            <el-table :data="orderData" stripe style="width: 100%" border>
                <el-table-column prop="oId" label="挂号单号" width="100px"></el-table-column>
                <el-table-column prop="dId" label="本人id" width="80px"></el-table-column>
                <el-table-column prop="pId" label="患者id" width="100px">
                </el-table-column>
                <el-table-column prop="pName" label="患者姓名" width="100px">
                </el-table-column>
                <el-table-column prop="oStart" label="挂号时间" width="190px"></el-table-column>
                <el-table-column prop="oEnd" label="结束时间" width="180px"></el-table-column>
                <el-table-column prop="oRecord" label="病因" width="400px"></el-table-column>
                <el-table-column prop="drugs" label="药物" width="180px"></el-table-column>
                <el-table-column prop="checks" label="检查项目" width="180px"></el-table-column>
                <el-table-column prop="oTotalPrice" label="需交费用/元" width="80px"></el-table-column>
                <el-table-column prop="oPriceState" label="缴费状态" width="100px">
                    <template #default="scope">
                        <el-tag type="success" v-if="scope.row.oPriceState === 1">已缴费</el-tag>
                        <!-- <el-tag type="danger" v-if="scope.row.oPriceState === 0 && scope.row.oState === 1">未缴费</el-tag> -->
                        <el-tag type="danger" v-if="scope.row.oPriceState === 0 &&
                            scope.row.oState === 1
                            ">未缴费</el-tag>
                    </template>
                </el-table-column>
                <el-table-column prop="oState" label="挂号状态" width="100px">
                    <template #default="scope">
                        <el-tag type="success" v-if="scope.row.oState === 1 &&
                            scope.row.oPriceState === 1
                            ">已完成</el-tag>
                        <el-tag type="danger" v-if="scope.row.oState === 0 && scope.row.oState === 0
                            ">未完成</el-tag>
                    </template>
                </el-table-column>
                <el-table-column label="操作" width="140" fixed="right">
                    <template #default="scope">
                        <el-button type="warning" :icon="List" style="font-size: 14px"
                            @click="dealClick(scope.row.oId, scope.row.pId, scope.row.pName)">查看</el-button>
                    </template>
                </el-table-column>
            </el-table>

            <!-- 分页 -->
            <el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange" background
                layout="total, sizes, prev, pager, next, jumper" :current-page="pageNumber" :page-size="size"
                :page-sizes="[1, 2, 4, 8, 16]" :total="total">
            </el-pagination>
        </el-card>

        <!-- 流程 -->
        <el-dialog title="流程" v-model="starVisible">
            <div style="height:60vh;overflow-y: scroll;">
                <el-timeline>
                    <el-timeline-item v-for="(item, index) in starData" :key="index" :timestamp="item.timestamp"
                        placement="top">
                        <el-card>
                            <h4>{{ starName }}患者{{ item.sDesc }}</h4>
                            <!-- <p>Tom committed 2018/4/12 20:46</p> -->
                        </el-card>
                    </el-timeline-item>
                </el-timeline>
            </div>
            <template #footer>
                <div class="dialog-footer" style="    text-align: center;">
                    <el-button @click="starClick" style="font-size: 18px;" type="primary"><i class="el-icon-close"
                            style="font-size: 20px;"></i> 关闭</el-button>
                    <!-- <el-button type="primary" @click="starClick" style="font-size: 18px;"><i class="el-icon-check"
                            style="font-size: 20px;"></i> 确 定</el-button> -->
                </div>
            </template>
        </el-dialog>
    </div>
</template>
<script setup>
import { ref, reactive, onMounted, } from "vue"
import request from "@/utils/request.ts"
import { useRouter, useRoute } from "vue-router";
import { ElMessage, ElMessageBox } from "element-plus"
import { Check, Delete, Close, Search, Edit, List, ChatRound, Memo } from "@element-plus/icons-vue"
import { getFindOrderByDid, getFindStarByDid } from "@/api/doctor/doctorOrder.ts"
let userId = ref(localStorage.getItem("userid"))
let orderData = ref([
    {
        oId: 214345,//挂号单号
        dId: 1,//本人id
        pId: 1243546,//本人id
        pName: "张三",//姓名
        oStart: "2023-10-10 09:00",//挂号时间
        oEnd: "2023-10-10 10:00",//结束时间
        oRecord: "扁桃体发炎",//需交费用/元
        oDrugBuyData: ["阿莫西林", "头孢"],//诊疗费用/元
        oCheckBuyData: ["B超", "心电图"],//缴费状态
        oTotalPrice: 200,//需交费用
        oPriceState: 1,//缴费状态
        oState: 1,//挂号状态
    },
   
])
let pageNumber = ref(1)
let size = ref(8)
let query = ref("")
let total = ref(3)
let starVisible = ref(false)
let starData = ref([
    {
        timestamp: "2023-10-11 09:10",
        sDesc: "已挂号"
    },
    {
        timestamp: "2023-10-11 09:10",
        sDesc: "已就诊"
    },
    {
        timestamp: "2023-10-11 09:10",
        sDesc: "已缴费"
    },
    {
        timestamp: "2023-10-11 09:10",
        sDesc: "已出院"
    },
])
let starName = ref("")

let route = useRoute()
let router = useRouter()

//点击追诊按钮
function dealClick(oId, pId, pName) {
    // 请求患者治疗流程
    getFindStarByDid({
        oId: oId,
        pId: pId,
    }).then((res) => {
            if (res.status !== 200)
                ElMessage.error("请求数据失败");
            starData.value = res.data;
        });
    starName.value = pName
    starVisible.value = true;

}
//页面大小改变时触发
function handleSizeChange(size1) {
    // console.log(size);
    size.value = size1;
    requestOrders();
}
//   页码改变时触发
function handleCurrentChange(num) {
    console.log(num);
    pageNumber.value = num;
    requestOrders();
}

//请求挂号信息
function requestOrders() {
    getFindOrderByDid({
        dId: userId.value,
        pageNumber: pageNumber.value,
        size: size.value,
        query: query.value,
    })
        .then((res) => {
            let orders = res.data.orders || []
            orders.forEach(o => {
                if(o.oDrugBuyData) {
                    let drugs = JSON.parse(o.oDrugBuyData) || []
                    let drugStr = ''
                    drugs.forEach(drug => drugStr += `${drug.DrName}, `) 
                    o.drugs = drugStr.slice(0, -1)
                }
                if(o.oCheckBuyData) {
                    let checks = JSON.parse(o.oCheckBuyData) || []
                    let checkStr = ''
                    checks.forEach(check => checkStr += `${check.ChName}, `)
                    o.checks = checkStr.slice(0, -1)
                }
            });
            if (res.status !== 200)
                ElMessage.error("请求数据失败");
            orderData.value = orders;
            // total.value = res.data.total;
        });
}
function starClick(){
    starVisible.value = false
    starData.value=null
}
//token解码
// tokenDecode(token) {
//     if (token !== null) return jwtDecode(token);
// },
onMounted(() => {
    // 解码token
    //this.orderData.pName = this.tokenDecode(getToken()).pName;
    //this.orderData.pCard = this.tokenDecode(getToken()).pCard;
    // this.userId = this.tokenDecode(getToken()).dId;
    // console.log(this.orderData.pName);
    //this.orderData.pName = "dasda"
    requestOrders();
})
</script>
<style lang="scss" scoped>
.el-table {
    margin-top: 20px;
    margin-bottom: 20px;
}
</style>