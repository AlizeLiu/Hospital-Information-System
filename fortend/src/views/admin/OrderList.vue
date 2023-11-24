<template>
    <!-- 卡片 -->
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
        <!-- 表格 -->
        <el-table :data="orderData" stripe style="width: 100%" border>
            <el-table-column prop="oId" label="挂号单号" width="80px"></el-table-column>
            <el-table-column prop="pId" label="患者id" width="80px"></el-table-column>

            <el-table-column prop="dId" label="医生id" width="100px">
            </el-table-column>

            <el-table-column prop="oStart" label="挂号时间" width="180px"></el-table-column>
            <el-table-column prop="oEnd" label="结束时间" width="180px"></el-table-column>
            <el-table-column prop="oRecord" label="病因" width="400px"></el-table-column>
            <el-table-column prop="oDrug" label="药物" width="180px"></el-table-column>
            <el-table-column prop="oCheck" label="检查项目" width="180px"></el-table-column>
            <el-table-column prop="oTotalPrice" label="费用/元" width="80px"></el-table-column>
            <el-table-column prop="oPriceState" label="缴费状态" width="100px">
                <template #default="scope">
                    <el-tag type="success" v-if="scope.row.oPriceState === 1">已缴费</el-tag>
                    <el-tag type="danger" v-if="scope.row.oPriceState === 0">未缴费</el-tag>
                    <!-- <el-button type="danger" size="mini" v-if="scope.row.oPriceState === 0 &&
                        scope.row.oState === 1
                        " @click="priceClick(scope.row.oId)">点击缴费</el-button> -->
                </template>
            </el-table-column>
            <el-table-column prop="oState" label="挂号状态" width="100px">
                <template #default="scope">
                    <el-tag type="success" v-if="scope.row.oState === 1
                        ">已挂号</el-tag>
                    <el-tag type="danger" v-if="scope.row.oState === 0">未挂号</el-tag>
                </template>
            </el-table-column>
            <el-table-column label="操作" width="140" fixed="right">
                <template #default="scope">
                    <el-button :icon="Delete" style="font-size: 14px" type="danger"
                        @click="deleteDialog(scope.row.oId)"></el-button>
                </template>
            </el-table-column>
        </el-table>

        <!-- 分页 -->
        <el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange" background
            layout="total, sizes, prev, pager, next, jumper" :current-page="pageNumber" :page-size="size"
            :page-sizes="[1, 2, 4, 8, 16]" :total="total">
        </el-pagination>
    </el-card>
</template>
<script setup>
import { ref, reactive, onMounted, } from 'vue'
import request from '@/utils/request.ts'
import { useRoute } from 'vue-router';
import { ElMessage, ElMessageBox } from 'element-plus'
import {
    Check,
    Delete, Close, Search
} from '@element-plus/icons-vue'
import { getDeleteOrder, getFindAllOrders } from '@/api/admin/orderList.ts'

const pageNumber = ref(1)
const size = ref(8)
const query = ref('')
const total = ref(3)
const orderData = reactive([
    {
        oId: 1,
        pId: 1,
        dId: 1,
        oStart: '2023-04-04 23:00',
        oEnd: '2023-04-04 23:00',
        oRecord: '肚子不适',//病因
        oDrug: '阿莫西林',//
        oCheck: 'B超',//
        oTotalPrice: '15元',
        oPriceState: 1,
        oState: 1,
    },
    {
        oId: 2,
        pId: 2,
        dId: 2,
        oStart: '2023-04-04 23:00',
        oEnd: '2023-04-04 23:00',
        oRecord: '肚子不适',//病因
        oDrug: '阿莫西林',//
        oCheck: 'B超',//
        oTotalPrice: '15元',
        oPriceState: 0,
        oState: 0,
    },
]);
//删除挂号操作
function deleteOrder(id) {
    getDeleteOrder({
        oId: id,
    })
        .then((res) => {
            requestOrders();
            console.log(res);
        });
}
//删除对话框
function deleteDialog(id) {
    ElMessageBox.confirm("此操作将永久删除该挂号信息, 是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
    })
        .then(() => {
            deleteOrder(id);
            ElMessage({
                type: "success",
                message: "删除成功!",
            });
        })
        .catch(() => {
            ElMessage({
                type: "info",
                message: "已取消删除",
            });
        });
}
//页面大小改变时触发
function handleSizeChange(size1) {
    size.value = size1;
    requestOrders();
}
//   页码改变时触发
function handleCurrentChange(num) {
    console.log(num);
    pageNumber.value = num;
    requestOrders();
}
// 加载订单列表
function requestOrders() {
    getFindAllOrders({
        pageNumber:pageNumber,
        size: size,
        query: query,

    }).then((res) => {
        orderData = res.data.data.orders;
        total.value = res.data.data.total;
        console.log(res.data.data);
    });
}
onMounted(() => {
    requestOrders();
});
</script>
<style scoped lang="scss">
.el-table {
    margin-top: 20px;
    margin-bottom: 20px;
}

.el-form {
    margin-top: 0;
}
</style>