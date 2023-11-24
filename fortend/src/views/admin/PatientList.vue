<template>
    <!-- 卡片 -->
    <el-card>
        <!-- 搜索栏 -->
        <el-row type="flex">
            <el-col :span="6">
                <el-input v-model="query" placeholder="请输入姓名查询">
                    <template #append>
                        <el-button :icon="Search" @click="requestPatients"></el-button>
                    </template>
                </el-input>
            </el-col>
        </el-row>
        <!-- 表格 -->
        <el-table :data="patientData" stripe style="width: 100%" border>
            <el-table-column prop="pId" label="账号" width="100">
            </el-table-column>
            <el-table-column prop="pName" label="姓名" width="80">
            </el-table-column>
            <el-table-column prop="pGender" label="性别" width="60">
            </el-table-column>
            <el-table-column prop="pAge" label="年龄/岁" width="180">
            </el-table-column>
            <el-table-column prop="pCard" label="证件号"> </el-table-column>
            <el-table-column prop="pPhone" label="手机号"> </el-table-column>
            <el-table-column prop="pEmail" label="邮箱" width="170">
            </el-table-column>
            <el-table-column prop="pState" label="状态" width="150">
                <template #default="scope">
                    <el-tag type="success" v-if="scope.row.pState === 1">正常</el-tag>
                    <el-tag type="danger" v-else>已禁用</el-tag>
                </template>
            </el-table-column>
            <el-table-column label="操作" width="200" fixed="right">
                <template #default="scope">
                    <el-button :icon="Delete" style="font-size: 14px" type="danger"
                        @click="deleteDialog(scope.row.pId)"></el-button>
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
import { getDeletePatient, getFindAllPatients } from '@/api/admin/patientList.ts'

let pageNumber = ref(1)
let size = ref(8)
let query = ref('')
let total = ref(3)
let patientData = reactive([
    {
        pId: 1,
        pName: '张',
        pGender: '男',//0 男，1女
        pAge:33,
        pCard: 1000001,
        pPhone: 13512341234,
        pEmail: '13456@qq.com',
        pState: 1,
    }
]);
//删除病人操作
function deletePatient(id) {
    getDeletePatient({
                pId: id,
            })
        .then((res) => {
            requestPatients();
            console.log(res);
        });
}
//删除对话框
function deleteDialog(id) {
    ElMessageBox.confirm("此操作将删除该患者信息, 是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
    })
        .then(() => {
            deletePatient(id);
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
    // console.log(typeof size);
    size.value = size1;
    requestPatients();
}
//   页码改变时触发
function handleCurrentChange(num) {
    console.log(num);
    pageNumber.value = num;
    requestPatients();
}
// 加载患者列表
function requestPatients() {
    getFindAllPatients({
                pageNumber: pageNumber.value,
                size: size.value,
                query: query.value,
        })
        .then((res) => {
            patientData = res.data.patients;

            total.value = res.data.total;
            console.log(res);
        });
}
onMounted(() => {
    requestPatients();
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