<template>
    <div>
        <!-- 卡片 -->
        <el-card>
            <!-- 搜索栏及增加检查 -->
            <el-row type="flex">
                <el-col :span="6">
                    <el-input v-model="query" placeholder="请输入名称查询">
                        <template #append>
                            <el-button icon="el-icon-search" @click="requestChecks"></el-button>
                        </template>
                    </el-input>
                </el-col>
                <el-col :span="6"></el-col>
                <el-col :span="6">
                    <el-button type="primary" style="font-size: 18px" @click="addFormVisible = true">
                        <i class="el-icon-circle-plus-outline" style="font-size: 22px;"></i>
                        增加项目</el-button>
                </el-col>
            </el-row>
            <!-- 表格 -->
            <el-table :data="checkData" stripe style="width: 100%" border>
                <el-table-column label="编号" prop="chId"></el-table-column>
                <el-table-column label="项目" prop="chName"></el-table-column>
                <el-table-column label="价格/元" prop="chPrice"></el-table-column>
                <el-table-column label="操作" width="200" fixed="right">
                    <template #default="scope">
                        <el-button style="font-size: 14px" type="success" @click="modifyDialog(scope.row.chId)"><i
                                class="el-icon-edit-outline" style="font-size: 22px;"></i></el-button>
                        <el-button style="font-size: 14px" type="danger" @click="deleteDialog(scope.row.chId)"><i
                                class="el-icon-delete" style="font-size: 22px;"></i></el-button>
                    </template>
                </el-table-column>
            </el-table>

            <!-- 分页 -->
            <el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange" background
                layout="total, sizes, prev, pager, next, jumper" :current-page="pageNumber" :page-size="size"
                :page-sizes="[1, 2, 4, 8, 16]" :total="total">
            </el-pagination>
        </el-card>

        <!-- 增加检查项目对话框 -->
        <el-dialog title="增加检查项目" v-model="addFormVisible">
            <el-form :model="addForm" :rules="rules" ref="ruleForm">
                <el-form-item label="编号" prop="chId" label-width="80px">
                    <el-input v-model.number="addForm.chId"></el-input>
                </el-form-item>
                <el-form-item label="名称" prop="chName" label-width="80px">
                    <el-input v-model="addForm.chName"></el-input>
                </el-form-item>
                <el-form-item label="价格" prop="chPrice" label-width="80px">
                    <el-input v-model="addForm.chPrice"></el-input>
                </el-form-item>
            </el-form>
            <template #footer>
                <div class="dialog-footer">
                    <el-button @click="addFormVisible = false" style="font-size: 18px;"><i class="el-icon-close"
                            style="font-size: 20px;"></i> 取 消</el-button>
                    <el-button type="primary" @click="addCheck('ruleForm')" style="font-size: 18px;"><i
                            class="el-icon-check" style="font-size: 20px;"></i> 确 定</el-button>
                </div>
            </template>

        </el-dialog>

        <!-- 修改检查项目药物对话框 -->
        <el-dialog title="修改检查项目" v-model="modifyFormVisible">
            <el-form :model="modifyForm" :rules="rules" ref="ruleForm">
                <el-form-item label="编号" prop="chId" label-width="80px">
                    <el-input v-model.number="modifyForm.chId" disabled></el-input>
                </el-form-item>
                <el-form-item label="名称" prop="chName" label-width="80px">
                    <el-input v-model="modifyForm.chName"></el-input>
                </el-form-item>
                <el-form-item label="价格" prop="chPrice" label-width="80px">
                    <el-input v-model="modifyForm.chPrice"></el-input>
                </el-form-item>
            </el-form>
            <template #footer>
                <div class="dialog-footer">
                    <el-button @click="modifyFormVisible = false" style="font-size: 18px;"><i class="el-icon-close"
                            style="font-size: 20px;"></i> 取 消</el-button>
                    <el-button type="primary" @click="modifyCheck('ruleForm')" style="font-size: 18px;"><i
                            class="el-icon-check" style="font-size: 20px;"></i> 确 定</el-button>
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
    Delete, Close, Search
} from '@element-plus/icons-vue'
import { getModifyCheck, getFindCheck, getDeleteCheck, getAddCheck, getFindAllChecks } from '@/api/admin/checkList.ts'


const pageNumber = ref(1)
const size = ref(8)
const query = ref('')
const total = ref(3)
const checkData = reactive([
    {
        chId: 1,//编号
        chName: 'B超',//项目
        chPrice: '11元',//价格/元
    },
    {
        chId: 1,//编号
        chName: '核磁',//项目
        chPrice: '11元',//价格/元
    },
]);
const addFormVisible = ref(false)
const addForm = reactive({});
const rules = reactive({
    chId: [
        { required: true, message: "请输入编号", trigger: "blur" },
        {
            type: "number",
            message: "账号必须数字类型",
            trigger: "blur",
        },
    ],
    chName: [
        { required: true, message: "请输入名称", trigger: "blur" },
        {
            min: 1,
            max: 50,
            message: "账号必须是1到50个字符",
            trigger: "blur",
        },
    ],
    chPrice: [
        { required: true, message: "请输入单价", trigger: "blur" },
    ],
});
const modifyFormVisible = ref(false)
const modifyForm = reactive({});
//点击修改药物信息
function modifyCheck(formName) {
    this.$refs[formName].validate((valid) => {
        if (valid) {
            getModifyCheck({
                        chId: modifyForm.chId,
                        chName: modifyForm.chName,
                        chPrice: modifyForm.chPrice,
                    
                })
                .then((res) => {
                    if (res.data.status !== 200)
                        return ElMessage.error("修改信息失败！");
                    modifyFormVisible.value = false;
                    requestChecks();
                    ElMessage.success("修改检查项目信息成功！");
                    console.log(res);
                });
        } else {
            console.log("error submit!!");
            return false;
        }
    });
}
//打开修改对话框
function modifyDialog(id) {
    getFindCheck({
                chId: id,
            
        }).then((res) => {
            if (res.data.status !== 200)
                return ElMessage.error("请求数据失败");
            modifyForm = res.data.data;
            modifyFormVisible.value = true;
            console.log(res);
        });
}  //删除检查操作
function deleteCheck(id) {
    getDeleteCheck({
                chId: id,            
        }).then((res) => {
            requestChecks();
            console.log(res);
        });
}
//删除对话框
function deleteDialog(id) {
    ElMessageBox.confirm("此操作将删除该检查项目信息, 是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
    })
        .then(() => {
            deleteCheck(id);
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
function addCheck(formName) {
    this.$refs[formName].validate((valid) => {
        if (valid) {
            getAddCheck({
                        chId: addForm.chId,
                        chName: addForm.chName,
                        chPrice: addForm.chPrice,
                }).then((res) => {
                    if (res.data.status !== 200)
                        return ElMessage.error(
                            "编号不合法或已被占用！"
                        );
                    addFormVisible.value = false;
                    requestChecks();
                    ElMessage.success("增加检查项目成功！");
                    console.log(res);
                });
        } else {
            console.log("error submit!!");
            return false;
        }
    });
}
//页面大小改变时触发
const handleSizeChange = (size1) => {
    size.value = size1;
    requestChecks();
}
//   页码改变时触发
function handleCurrentChange(num) {
    console.log(num);
    pageNumber.value = num;
    requestChecks();
}
// 加载检查列表
function requestChecks() {
    getFindAllChecks({
                pageNumber: pageNumber.value,
                size: size.value,
                query: query.value,
        })
        .then((res) => {
            checkData.value = res.data.data.checks;
            total.value = res.data.data.total;
            console.log(res.data.data);
        });
}
onMounted(() => {
    requestChecks();
})
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