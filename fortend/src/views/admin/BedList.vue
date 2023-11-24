<template>
    <div>
        <!-- 卡片 -->
        <el-card>
            <!-- 搜索栏及增加检查 -->
            <el-row type="flex">
                <el-col :span="6">
                    <el-input v-model="query" placeholder="请输入患者id查询">
                        <template #append>
                            <el-button @click="requestBeds" :icon="Search"></el-button>
                        </template>
                    </el-input>
                </el-col>
                <el-col :span="1"></el-col>
                <el-col :span="6">
                    <el-button type="primary" style="font-size: 18px" @click="addFormVisible = true">
                        <el-icon style="font-size: 22px;"><CirclePlusFilled /></el-icon>增加床位
                        </el-button>
                </el-col>
            </el-row>
            <!-- 表格 -->
            <el-table :data="bedData" stripe style="width: 100%" border>
                <el-table-column label="床号" prop="bId"></el-table-column>
                <el-table-column label="患者id" prop="pId"></el-table-column>
                <el-table-column label="医生id" prop="dId"></el-table-column>
                <el-table-column label="开始时间" prop="bStart"></el-table-column>
                <el-table-column label="申请理由" prop="bReason"></el-table-column>
                <el-table-column label="状态" prop="bState">
                    <template #default="scope">
                        <el-tag v-if="scope.row.bState === 1" type="danger">已占用</el-tag>
                        <el-tag v-if="scope.row.bState === 0" type="success">空</el-tag>
                    </template>
                </el-table-column>

                <el-table-column label="操作" width="200" fixed="right">
                    <template #default="scope">
                        <el-button style="font-size: 14px" type="success" @click="deleteDialog(scope.row.bId)" :icon="Delete"></el-button>
                        <el-button style="font-size: 14px" type="danger" @click="emptyDialog(scope.row.bId)" :icon="Check"></el-button>
                    </template>
                </el-table-column>
            </el-table>

            <!-- 分页 -->
            <el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange" background
                layout="total, sizes, prev, pager, next, jumper" :current-page="pageNumber" :page-size="size"
                :page-sizes="[1, 2, 4, 8, 16]" :total="total">
            </el-pagination>
        </el-card>

        <!-- 增加床位对话框 -->
        <el-dialog title="增加床位" v-model="addFormVisible">
            <el-form :model="addForm" :rules="rules" ref="ruleForm">
                <el-form-item label="床号" prop="bId" label-width="80px">
                    <el-input v-model.number="addForm.bId"></el-input>
                </el-form-item>
            </el-form>
            <template #footer>
                <div class="dialog-footer">
                    <el-button @click="addFormVisible = false" style="font-size: 18px;" :icon="Close">取 消</el-button>
                    <el-button type="primary" @click="addBed('ruleForm')" style="font-size: 18px;" :icon="Check"> 确 定</el-button>
                </div>
            </template>

        </el-dialog>
    </div>
</template>
<script setup>
import { ref, reactive, onMounted, } from 'vue'
import request from '@/utils/request.ts'
import { useRoute } from 'vue-router';
import { ElMessage, ElMessageBox } from 'element-plus'
import {
  Check,
  Delete,Close,Search
} from '@element-plus/icons-vue'
import { getEmptyBed, getDeleteBed, getAddBed, getFindAllBeds } from '@/api/admin/bedList.ts'

const pageNumber = ref(1)
const size = ref(8)
const query = ref('')
const total = ref(3)
const bedData = reactive([
    {
        bId: 1,
        pId: 1,
        dId: 1,
        bStart: '2023-01-01',
        bReason: '理由一',
        bState: 1
    },
    {
        bId: 2,
        pId: 2,
        dId: 2,
        bStart: '2023-02-01',
        bReason: '理由一',
        bState: 0
    },
]);
const addFormVisible = ref(false)
const addForm = reactive({});
const rules = reactive({
    bId: [
        { required: true, message: "请输入床号", trigger: "blur" },
        {
            type: "number",
            message: "床号必须数字类型",
            trigger: "blur",
        },
    ],
});
//清空床位操作

function emptyBed(id) {
    getEmptyBed({ bId: id }).then((res) => {
        requestBeds();
        console.log(res);
    });
}
//清空对话框

function emptyDialog(id) {
    ElMessageBox.confirm("此操作将清空该床位, 是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
    })
        .then(() => {
            emptyBed(id);
            ElMessage({
                type: "success",
                message: "清空成功!",
            });
        })
        .catch(() => {
            ElMessage({
                type: "info",
                message: "已取消清空",
            });
        });
}
//删除床位操作
function deleteBed(id) {
    getDeleteBed({
        bId: id,
    })
        .then((res) => {
            requestBeds();
            console.log(res);
        });
}
//删除对话框
function deleteDialog(id) {
    ElMessageBox.confirm("此操作将删除该床位, 是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
    })
        .then(() => {
            deleteBed(id);
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
//点击增加确认按钮
function addBed(formName) {
    this.$refs[formName].validate((valid) => {
        if (valid) {
            getAddBed({
                bId: addForm.bId,
                pId: -1,
                dId: -1,
            })
                .then((res) => {
                    if (res.data.status !== 200)
                        return ElMessage.error("床号或已被占用！");
                    addFormVisible.value = false;
                    requestBeds();
                    ElMessage.success("增加床位成功！");
                    console.log(res);
                });
        } else {
            console.log("error submit!!");
            return false;
        }
    });
}
//页面大小改变时触发
function handleSizeChange(size1) {
    console.log(size1);
    size.value = size;
    requestBeds();
}
//   页码改变时触发
function handleCurrentChange(num) {
    console.log(num);
    pageNumber.value = num;
    requestBeds();
}

// 加载检查列表
function requestBeds() {
    getFindAllBeds({
        pageNumber: pageNumber,
        size: size,
        query: query,
    }).then((res) => {
            bedData.value = res.data.data.beds;
            total.value = res.data.data.total;
            console.log(res.data.data);
        });
}
onMounted(() => {
    requestBeds();
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