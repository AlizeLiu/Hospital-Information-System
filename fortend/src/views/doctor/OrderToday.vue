<template>
    <el-card>
        <el-table :data="orderData" stripe border>
            <el-table-column label="序号" type="index" width="100">
            </el-table-column>
            <el-table-column label="挂号单号" prop="oId"></el-table-column>
            <!-- <el-table-column label="患者id" prop="pId"></el-table-column> -->
            <el-table-column label="患者姓名" prop="pName"></el-table-column>
            <el-table-column label="医生姓名" prop="dName"></el-table-column>
            <el-table-column label="挂号时间" prop="oStart" width="200"></el-table-column>
            <el-table-column label="操作">
                <template #default="scope">
                    <el-button type="warning" style="font-size: 18px" @click="dealClick(scope.row.oId, scope.row.pId,scope.row.dName)">
                        <i class="el-icon-monitor" style="font-size: 18px"></i>
                        处理
                    </el-button>
                </template>
            </el-table-column>

        </el-table>
    </el-card>
</template>
<script setup>
import { ref, reactive, onMounted, } from 'vue'
import request from '@/utils/request.ts'
import { useRouter, useRoute } from 'vue-router';
import { ElMessage, ElMessageBox } from 'element-plus'
import { Check, Delete, Close, Search, Edit, List, ChatRound, Memo } from '@element-plus/icons-vue'
import { getFindOrderByNull } from '@/api/doctor/orderToday.ts'
let userId = ref(localStorage.getItem("userid"))
let userName = ref('dada')
let today = ref('')
let orderData = ref([
    {
        oId: 312434,// 挂号单号
        pId: 123423,// 挂号单号
        pName: '张三',//患者姓名
        dName: '李医生',//患者姓名
        oStart: '2023-10-10 09:00',//挂号时间
    },
    {
        oId: 312434,// 挂号单号
        pId: 123423,// 挂号单号
        pName: '张三',//患者姓名
        dName: '李医生',//患者姓名
        oStart: '2023-10-10 09:00',//挂号时间
    },
    {
        oId: 312434,// 挂号单号
        pId: 123423,// 挂号单号
        pName: '张三',//患者姓名
        dName: '李医生',//患者姓名
        oStart: '2023-10-10 09:00',//挂号时间
    },
])
let router = useRouter();
//挂号处理//页面跳转传值
function dealClick(oId, pId,dName) {
    console.log(oId, pId);
    router.push(
        {
            path: "/dealOrder",
            query: {
                oId: oId,
                pId: pId,
                dName:dName
            }
        }
    );
}
//获取挂号信息
function requestOrder() {
    getFindOrderByNull({
        dId: userId.value,
        oStart: today.value
    })
        .then(res => {
            if (res.status != 200)
                return ElMessage.error("获取数据失败");
            orderData.value = res.data;
            //this.orderData.dName = this.userName;
            console.log(res);

        })
}
//token解码
// tokenDecode(token){
//   if (token !== null)
//   return jwtDecode(token);
// },
//获取当天日期
function nowDay() {
    const nowDate = new Date();
    let date = {
        year: nowDate.getFullYear(),
        month: nowDate.getMonth() + 1,
        date: nowDate.getDate(),
    };
    if (date.date < 10) {
        date.date = "0" + date.date
    }
    if (date.month < 10) {
        date.month = "0" + date.month
    }
    today.value = date.year + "-" + date.month + "-" + date.date;

}
onMounted(() => {
    //解码token信息
    // this.userId = this.tokenDecode(getToken()).dId;
    // this.userName = this.tokenDecode(getToken()).dName;
    // console.log(this.userId);
    // console.log(this.userName);
    //获取当天日期
    nowDay();
    //获取订单信息
    requestOrder();

})

</script>