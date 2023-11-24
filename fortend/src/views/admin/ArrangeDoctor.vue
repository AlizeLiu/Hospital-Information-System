<template>
    <div>
        <!-- 面包屑 -->
        <el-breadcrumb separator-class="el-icon-arrow-right">
            <el-breadcrumb-item :to="{ path: '/sectionIndex' }">返回科室</el-breadcrumb-item>
            <el-breadcrumb-item>{{ section }}</el-breadcrumb-item>
        </el-breadcrumb>
        <el-input v-model="query" placeholder="请输入姓名查询" class="doctorInput">
            <template #append>
                <el-button icon="el-icon-search" @click="requestDoctors"></el-button>
            </template>
        </el-input>
        <el-table :data="doctorData" border>
            <el-table-column label="账号" prop="dId" v-model="doctorData.dId"></el-table-column>
            <el-table-column label="姓名" prop="dName" v-model="doctorData.dName"></el-table-column>
            <el-table-column label="性别" prop="dGender" v-model="doctorData.dGender"></el-table-column>
            <el-table-column label="职位" prop="dPost" v-model="doctorData.dPost"></el-table-column>
            <el-table-column label="部门" prop="dSection" v-model="doctorData.dSection"></el-table-column>
            <el-table-column label="操作" prop="dSection">
                <template #default="scope">
                    <el-button v-if="scope.row.arrangeId == null" type="success" icon="iconfont icon-r-paper"
                        style="font-size: 14px;" @click="arrangeClick(scope.row.dId)"> 排班</el-button>
                    <el-button v-if="scope.row.arrangeId != null" type="danger" icon="el-icon-delete"
                        style="font-size: 14px;" @click="deleteArrange(scope.row.arrangeId)"> 取消排班</el-button>
                </template>
            </el-table-column>
        </el-table>
        <!-- 分页 -->
        <el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange" background
            layout="total, sizes, prev, pager, next, jumper" :current-page="pageNumber" :page-size="size"
            :page-sizes="[1, 2, 4, 8, 16]" :total="total">
        </el-pagination>
    </div>
</template>
<script setup>
import { ref, reactive, onMounted, } from 'vue'
import request from '@/utils/request.ts'
import { useRoute } from 'vue-router';
import { ElMessage } from 'element-plus'

const doctorData = reactive([]);
const total = ref(3)
const pageNumber = ref(1)
const size = ref(8)
const query = ref('')
const section = useRoute().query.section
//排班点击
const arrangeClick = (dId) => {
    request
        .get("arrange/addArrange", {
            params: {
                arId: dId + sessionStorage.getItem("arrangeDate"),
                arTime: sessionStorage.getItem("arrangeDate"),
                dId: dId,
            },
        })
        .then((res) => {
            if (res.data.status !== 200)
                return ElMessage.error('已排班');;
            ElMessage.success("排班成功！");
            requestDoctors();
        });
}
const deleteArrange = (arrangeId) => {
    request
        .get("arrange/deleteArrange", {
            params: {
                arId: arrangeId,
            },
        })
        .then((res) => {
            if (res.data.status !== 200)
                return ElMessage.error("排班信息不存在");
            ElMessage.success("删除排班信息成功！");
            requestDoctors();
        });
}
//页面大小改变时触发
const handleSizeChange = (size1) => {
    size.value = size1;
    requestDoctors();
}
//   页码改变时触发
const handleCurrentChange = () => {
    console.log(num);
    pageNumber.value = num;
    requestDoctors();
}
//根据部门请求医生信息
const requestDoctors = () => {
    request
        .get("doctor/findDoctorBySectionPage", {
            params: {
                pageNumber: pageNumber.value,
                size: size.value,
                query: query.value,
                dSection: section,
                arrangeDate: sessionStorage.getItem("arrangeDate"),
            },
        })
        .then((res) => {
            console.log(res.data);
            if (res.data.status !== 200)
                return ElMessage.error("数据请求失败");
            doctorData.value = res.data.data.doctors;
            total.value = res.data.data.total;
        });
}

onMounted(() => {
    requestDoctors();
});
</script>
<style scope lang="scss">
.el-breadcrumb {
    margin-bottom: 10px;
}

.doctorInput {
    width: 30%;
    margin-bottom: 10px;
}
</style>