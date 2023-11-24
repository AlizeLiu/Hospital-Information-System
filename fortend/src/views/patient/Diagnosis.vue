<!-- 诊断 -->
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
                <el-table-column prop="oPriceState" label="缴费状态" width="150">
                    <template #default="scope">
                        <el-tag type="success" v-if="scope.row.oPriceState === 1">已缴费</el-tag>
                        <!-- <el-tag type="danger" v-if="scope.row.oPriceState === 0 && scope.row.oState === 1">未缴费</el-tag> -->
                        <el-button type="warning" :icon="Memo" style="font-size: 14px" v-if="scope.row.oPriceState === 0
                            " @click="priceClick(scope.row.oId, scope.row.dId)">
                            点击缴费</el-button>
                    </template>
                </el-table-column>
                <el-table-column prop="oState" label="挂号状态" width="150px">
                    <template #default="scope">
                        <el-tag type="success" v-if="scope.row.oState === 1
                            ">已完成</el-tag>
                        <el-tag type="danger" v-if="scope.row.oState === 0
                            ">未完成</el-tag>
                    </template>
                </el-table-column>
                <el-table-column label="报告单">
                    <template #default="scope">
                        <el-button type="success" :icon="ChatRound" style="font-size: 14px"
                            @click="seeReport(scope.row.oId)" v-if="scope.row.oState === 1 &&
                                scope.row.oPriceState === 1
                                "> 查看</el-button>
                    </template>
                </el-table-column>
            </el-table>
        </el-card>
        <!-- 缴费 -->
        <el-dialog title="用户缴费" v-model="starVisible">
            <div>
                <h3>
                    请选择支付方式：
                </h3>
                <div>
                    <div class="left">微信:

                    </div>
                    <div class="center">支付宝:</div>
                    <div class="right">银行卡:</div>
                </div>
            </div>
            <div slot="footer" class="dialog-footer">
                <el-button @click="starVisible = false" style="font-size: 18px;"><i class="el-icon-close"
                        style="font-size: 20px;"></i> 取 消</el-button>
                <el-button type="primary" @click="starClick" style="font-size: 18px;"><i class="el-icon-check"
                        style="font-size: 20px;"></i> 确 定</el-button>
            </div>
        </el-dialog>
    </div>
</template>
<script setup>
import { ref, reactive, onMounted, } from 'vue'
import request from '@/utils/request.ts'
import { useRouter, useRoute } from 'vue-router';
import { ElMessage, ElMessageBox } from 'element-plus'
import { Check, Delete, Close, Search, Edit, List, ChatRound, Memo } from '@element-plus/icons-vue'
import { getUpdatePrice, getFindDoctor, getFindOrderByPid } from '@/api/patient/myOrder.ts'
let route = useRoute()
let userId = ref(1)
let orderData = reactive([
    {
        oId: 214345,//挂号单号
        // pId:1,//本人id
        pName: '张三',//姓名
        dName: '李医生',//姓名
        oStart: '2023-10-10 09:00',//挂号时间
        oEnd: '2023-10-10 10:00',//结束时间
        oTotalPrice: 20,//需交费用/元
        oPriceState: 1,//缴费状态
        oState: 1,//挂号状态
    },
    {
        oId: 214345,//挂号单号
        // pId:1,//本人id
        pName: '里三',//姓名
        dName: '李医生',//姓名
        oStart: '2023-10-10 09:00',//挂号时间
        oEnd: '2023-10-10 10:00',//结束时间
        oTotalPrice: 20,//需交费用/元
        oPriceState: 0,//缴费状态
        oState: 1,//挂号状态
    },
    {
        oId: 214345,//挂号单号
        // pId:1,//本人id
        pName: '里三',//姓名
        dName: '李医生',//姓名
        oStart: '2023-10-10 09:00',//挂号时间
        oEnd: '2023-10-10 10:00',//结束时间
        oTotalPrice: 20,//需交费用/元
        oPriceState: 0,//缴费状态
        oState: 0,//挂号状态
    },
])
let star = ref(5)
let starVisible = ref(false)
let dId = ref(1)
let dName = ref('');

onMounted(() => {
})
</script>
<style lang="scss" scoped>
.el-dialog div {
    text-align: center;
    margin-bottom: 8px;
}
</style>