<template>
    <div>
        <el-card>
            <el-table :data="orderData" stripe style="width: 100%" border>
                <el-table-column prop="oId" label="挂号单号" width="100px"></el-table-column>
                <!-- <el-table-column prop="pId" label="本人id" width="75px"></el-table-column> -->
                <el-table-column prop="pName" label="姓名" width="80px"></el-table-column>
                <!-- <el-table-column prop="dId" label="医生id" width="75px"></el-table-column> -->
                <el-table-column prop="dName" label="医生姓名" width="100px"></el-table-column>

                <el-table-column prop="oStart" label="挂号时间" width="195px"></el-table-column>
                <el-table-column prop="oEnd" label="结束时间" width="185px"></el-table-column>
                <el-table-column prop="oTotalPrice" label="需交费用/元" width="150px"></el-table-column>
                <el-table-column prop="oPrice" label="诊疗费用/元" width="150px"></el-table-column>

                <el-table-column prop="oStatePrice" label="挂号费用/元" width="150px"></el-table-column>
                <el-table-column prop="oState" label="挂号状态" width="150px">
                    <template #default="scope">
                        <el-tag type="success" v-if="scope.row.oState === 1
                            ">已完成</el-tag>
                        <el-tag type="danger" v-if="scope.row.oState === 0
                            ">未完成</el-tag>
                    </template>
                </el-table-column>
                <el-table-column prop="oPriceState" label="缴费状态" width="150">
                    <template #default="scope">
                        <el-tag type="success" v-if="scope.row.oPriceState === 1">已缴费</el-tag>
                        <!-- <el-tag type="danger" v-if="scope.row.oPriceState === 0 && scope.row.oState === 1">未缴费</el-tag> -->
                        <el-button type="warning" :icon="Memo" style="font-size: 14px" v-if="scope.row.oPriceState === 0
                            " @click="priceClick(scope.row.oId, scope.row.dId)">
                            点击缴费</el-button>
                    </template>
                </el-table-column>
                <el-table-column label="报告单" width="150" fixed="right">
                    <template #default="scope">
                        <el-button type="success" style="font-size: 14px" @click="seeReport(scope.row.oId, scope.row.dName)">
                            查看</el-button>
                    </template>
                </el-table-column>
            </el-table>
        </el-card>
        <!-- 缴费 -->
        <el-dialog title="用户缴费" v-model="starVisible">
            <div>
                <div>
                    <div style="display: flex;align-items: center;">
                        <h4>医生姓名：</h4>{{ dName }}
                    </div>
                    <div style="display: flex;align-items: center;">
                        <h4> 缴费金额：</h4>{{ sumPrice }}
                    </div>

                </div>
                <h3 style="text-align: left;">
                    请选择支付方式：
                </h3>
                <div style="display: flex;">
                    <div class="left">微信:
                        <img src="" alt="">
                    </div>
                    <div class="center">支付宝:<img src="" alt=""></div>
                    <div class="right">银行卡:<img src="" alt=""></div>
                </div>
            </div>
            <template #footer>
                <div class="dialog-footer">
                    <el-button @click="starVisible = false" style="font-size: 18px;"><i class="el-icon-close"
                            style="font-size: 20px;"></i> 取 消</el-button>
                    <el-button type="primary" @click="starClick" style="font-size: 18px;"><i class="el-icon-check"
                            style="font-size: 20px;"></i> 确认支付</el-button>
                </div>
            </template>
        </el-dialog>
        <!-- 查看 -->
        <el-dialog title="报告单" v-model="formVisible">
            <el-form :model="resultForm" ref="ruleForm" style="height:60vh;overflow-y: scroll;">
                <el-form-item label="医生:" label-width="100px" prop="dName">
                    <el-input v-model="resultForm.dName" autocomplete="off" disabled></el-input>
                </el-form-item>
                <el-form-item label="药物:" label-width="100px">
                    <el-table stripe border :data="resultForm.oDrugBuyData">
                        <el-table-column type="index" width="80" label="序号" />
                        <!-- <el-table-column label="编号" prop="drId"></el-table-column> -->
                        <el-table-column label="名称" prop="drName"></el-table-column>
                        <el-table-column label="单价/元" prop="drPrice"></el-table-column>
                        <el-table-column label="数量" prop="drNum"></el-table-column>=
                    </el-table>
                </el-form-item>
                <el-form-item label="检查项目:" label-width="100px" prop="oCheckBuyData">
                    <el-table stripe border class="rigthTable" :data="resultForm.oCheckBuyData">
                        <el-table-column type="index" width="80" label="序号" />
                        <!-- <el-table-column label="编号" prop="chId"></el-table-column> -->
                        <el-table-column label="项目" prop="chName"></el-table-column>
                        <el-table-column label="价格/元" prop="chPrice"></el-table-column>
                    </el-table>
                </el-form-item>
                <el-form-item label="医嘱:" label-width="100px" prop="oRecord">
                    <el-input v-model="resultForm.oRecord" autocomplete="off" disabled></el-input>
                </el-form-item>
            </el-form>
            <template #footer>
                <div class="dialog-footer">
                    <el-button type="primary" @click="formVisible = false" style="font-size: 18px;"><i class="el-icon-close"
                            style="font-size: 20px;"></i> 关 闭</el-button>
                    <!-- <el-button type="primary" @click="formVisible = false" style="font-size: 18px;"><i class="el-icon-check"
                            style="font-size: 20px;"></i> 确 定</el-button> -->
                </div>
            </template>
        </el-dialog>
    </div>
</template>
<script setup>
import { ref, reactive, onMounted, } from 'vue'
import request from '@/utils/request.ts'
import { useRouter, useRoute } from 'vue-router';
import { ElMessage, ElMessageBox } from 'element-plus'
import { Check, Delete, Close, Search, Edit, List, ChatRound, Memo } from '@element-plus/icons-vue'
import { getUpdatePrice, getFindDoctor, getFindOrderByPid, getFindOrderInfo } from '@/api/patient/myOrder.ts'
let route = useRoute()
let userId = ref(1)
let orderData = ref([
    {
        oId: 214345,//挂号单号
        // pId:1,//本人id
        pName: '张三',//姓名
        dName: '李医生',//姓名
        oStart: '2023-10-10 09:00',//挂号时间
        oEnd: '2023-10-10 10:00',//结束时间
        oTotalPrice: 200,//需交费用/元
        oPrice: 180,//诊疗费用/元
        oPriceState: 1,//缴费状态
        oStatePrice: 20,//挂号费用
        oState: 1,//挂号状态
    },
])
let star = ref(5)
let starVisible = ref(false)
let formVisible = ref(false)
let dId = ref(1)
let dName = ref('');
let pName = ref('');
let sumPrice = ref('')
let resultForm = ref({
    dName: '李医生',//姓名
    oCheckBuyData: [
        { chId: 1049, chPrice: 62, chName: 'B超' }
    ],
    oDrugBuyData: [
        { drId: 86, drPrice: 36, drName: '安深五', drNum: 1, drSum: 36 }
    ],
    oRecord: '有有有有有有有有有有有有有有有有有有有'
})
let formData = reactive([
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
//评价点击确认
// function starClick() {
//     console.log(star.value);
//     console.log(dId.value);
//     request
//         .get("doctor/updateStar", {
//             params: {
//                 dId: dId.value,
//                 dStar: star.value,
//             },
//         })
//         .then((res) => {
//             if (res.data.status !== 200)
//                 return this.$message.error("评价失败");
//             this.$message.success("谢谢您的评价");
//             starVisible.value = false;
//         });
// }
//查看报告单
function seeReport(id, name) {
    dName.value = name
    formVisible.value = true
    getFindOrderInfo({
        oId: id,
    }).then((res) => {
        if (res.status !== 200) {
            ElMessage.error("请求数据失败");
            return;
        }
        resultForm.value = res.data;
    });
}
//点击缴费按钮
function priceClick(oId, dId) {
    starVisible.value = true;
    getUpdatePrice({
        oId: oId,
    })
        .then((res) => {
            if (res.status !== 200) {
                ElMessage.error("请求数据失败");
                return;
            }
            // ElMessage.success("单号 " + oId + " 缴费成功！");
            dName.value = res.data.dName;
            sumPrice.value = res.data.sumPrice
            // starVisible.value = true;
            // requestOrder();
        });
}
//请求挂号信息
function requestOrder() {
    getFindOrderByPid({
        pId: userId.value,
    })
        .then((res) => {
            if (res.status !== 200)
                ElMessage.error("请求数据失败");
            orderData.value = res.data;
            //this.orderData.dSection = res.data.data.map(item => item.doctor.dSection);
            //console.log(res.data.data.map(item => item.doctor.dSection));
        });
}
//token解码
// tokenDecode(token) {
//     if (token !== null) return jwtDecode(token);
// },
onMounted(() => {
    // 解码token
    //this.orderData.pName = this.tokenDecode(getToken()).pName;
    //this.orderData.pCard = this.tokenDecode(getToken()).pCard;
    // this.userId = this.tokenDecode(getToken()).pId;
    //this.orderData.pName = "dasda"
    requestOrder();
})
</script>
<style lang="scss" scoped>
.el-dialog div {
    text-align: center;
    margin-bottom: 8px;
}
</style>