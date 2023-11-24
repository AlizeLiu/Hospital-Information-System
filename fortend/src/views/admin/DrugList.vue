<template>
    <div>
        <!-- 卡片 -->
        <el-card>
            <!-- 搜索栏及增加药物 -->
            <el-row type="flex">
                <el-col :span="6">
                    <el-input v-model="query" placeholder="请输入名称查询">
                        <template #append>
                            <el-button :icon="Search" @click="requestDrugs"></el-button>
                        </template>
                    </el-input>
                </el-col>
                <el-col :span="1"></el-col>
                <el-col :span="6">
                    <el-button type="primary" @click="addFormVisible = true" style="font-size: 18px">
                        <i class="el-icon-circle-plus-outline" style="font-size: 22px;"></i>
                        增加药物</el-button>
                </el-col>
            </el-row>
            <!-- 表格 -->
            <el-table :data="drugData" stripe style="width: 100%" border>
                <el-table-column label="编号" prop="drId"></el-table-column>
                <el-table-column label="名称" prop="drName"></el-table-column>
                <el-table-column label="剩余数量" prop="drNumber"></el-table-column>
                <el-table-column label="单位" prop="drUnit"></el-table-column>
                <el-table-column label="单价" prop="drPrice"></el-table-column>
                <el-table-column label="出版商" prop="drPublisher"></el-table-column>
                <el-table-column label="操作" width="200" fixed="right">
                    <template #default="scope">
                        <el-button style="font-size: 14px" type="success" @click="modifyDialog(scope.row.drId)">
                            <el-icon style="font-size: 22px;">
                                <Edit />
                            </el-icon></el-button>
                        <el-button style="font-size: 14px" type="danger" @click="deleteDialog(scope.row.drId)"
                            :icon="Delete"><i class="el-icon-delete" style="font-size: 22px;"></i></el-button>
                    </template>
                </el-table-column>
            </el-table>

            <!-- 分页 -->
            <el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange" background
                layout="total, sizes, prev, pager, next, jumper" :current-page="pageNumber" :page-size="size"
                :page-sizes="[1, 2, 4, 8, 16]" :total="total">
            </el-pagination>
        </el-card>

        <!-- 增加药物对话框 -->
        <el-dialog title="增加药物" v-model="addFormVisible">
            <el-form :model="addForm" :rules="rules" ref="ruleForm">
                <el-form-item label="编号" prop="drId" label-width="80px">
                    <el-input v-model.number="addForm.drId"></el-input>
                </el-form-item>
                <el-form-item label="名称" prop="drName" label-width="80px">
                    <el-input v-model="addForm.drName"></el-input>
                </el-form-item>
                <el-form-item label="数量" prop="drNumber" label-width="80px">
                    <el-input-number v-model="addForm.drNumber" :min="0" :max="1000"></el-input-number>
                </el-form-item>
                <el-form-item label="单位" prop="drUnit" label-width="80px">
                    <el-radio v-model="addForm.drUnit" label="盒">盒</el-radio>
                    <el-radio v-model="addForm.drUnit" label="袋">袋</el-radio>
                    <el-radio v-model="addForm.drUnit" label="剂">剂</el-radio>
                </el-form-item>
                <el-form-item label="单价" prop="drPrice" label-width="80px">
                    <el-input v-model="addForm.drPrice"></el-input>
                </el-form-item>
                <el-form-item label="出版商" prop="drPublisher" label-width="80px">
                    <el-input v-model="addForm.drPublisher"></el-input>
                </el-form-item>
            </el-form>
            <template #footer>
                <div class="dialog-footer">
                    <el-button @click="addFormVisible = false" style="font-size: 18px;"><i class="el-icon-close"
                            style="font-size: 20px;"></i> 取 消</el-button>
                    <el-button type="primary" @click="addDrug('ruleForm')" style="font-size: 18px;"><i class="el-icon-check"
                            style="font-size: 20px;"></i> 确 定</el-button>
                </div>
            </template>
        </el-dialog>

        <!-- 修改药物对话框 -->
        <el-dialog title="修改药物" v-model="modifyFormVisible">
            <el-form :model="modifyForm" :rules="rules" ref="ruleForm">
                <el-form-item label="编号" prop="drId" label-width="80px">
                    <el-input v-model.number="modifyForm.drId" disabled></el-input>
                </el-form-item>
                <el-form-item label="名称" prop="drName" label-width="80px">
                    <el-input v-model="modifyForm.drName"></el-input>
                </el-form-item>
                <el-form-item label="数量" prop="drNumber" label-width="80px">
                    <el-input-number v-model="modifyForm.drNumber" :min="0" :max="1000"></el-input-number>
                </el-form-item>
                <el-form-item label="单位" prop="drUnit" label-width="80px">
                    <el-radio v-model="modifyForm.drUnit" label="盒">盒</el-radio>
                    <el-radio v-model="modifyForm.drUnit" label="袋">袋</el-radio>
                    <el-radio v-model="modifyForm.drUnit" label="剂">剂</el-radio>
                </el-form-item>
                <el-form-item label="单价" prop="drPrice" label-width="80px">
                    <el-input v-model="modifyForm.drPrice"></el-input>
                </el-form-item>
                <el-form-item label="出版商" prop="drPublisher" label-width="80px">
                    <el-input v-model="modifyForm.drPublisher"></el-input>
                </el-form-item>
            </el-form>
            <template #footer>
                <div class="dialog-footer">
                    <el-button @click="modifyFormVisible = false" style="font-size: 18px;"><i class="el-icon-close"
                            style="font-size: 20px;"></i> 取 消</el-button>
                    <el-button type="primary" @click="modifyDrug('ruleForm')" style="font-size: 18px;"><i
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
import { getDeleteDrug, getFindDrug, getFindAllDrugs, getAddDrug, getModifyDrug } from '@/api/admin/drugList.ts'

const pageNumber = ref(1)
const size = ref(8)
const query = ref('')
const total = ref(3)
const drugData = reactive([
    {
        drId: 1,//编号
        drName: '张',//名称
        drNumber: 33,//剩余数量
        drUnit: 3,//单位
        drPrice: '11元',//单价
        drPublisher: '石家庄',//出版商
    },
    {
        drId: 1,//编号
        drName: '张',//名称
        drNumber: 33,//剩余数量
        drUnit: 3,//单位
        drPrice: '11元',//单价
        drPublisher: '石家庄',//出版商
    },
]);
const addFormVisible = ref(false)
const addForm = reactive({});
const modifyFormVisible = ref(false)
const modifyForm = reactive({});
const rules = reactive({
    drId: [
        { required: true, message: "请输入编号", trigger: "blur" },
        {
            type: "number",
            message: "账号必须数字类型",
            trigger: "blur",
        },
    ],
    drName: [
        { required: true, message: "请输入名称", trigger: "blur" },
        {
            min: 1,
            max: 50,
            message: "账号必须是1到50个字符",
            trigger: "blur",
        },
    ],
    drUnit: [
        { required: true, message: "请选择单位", trigger: "blur" },
    ],
    drPrice: [
        { required: true, message: "请输入单价", trigger: "blur" },
    ],
    drPublisher: [
        {
            required: true,
            message: "请输入出版商",
            trigger: "blur",
        },
        {
            min: 1,
            max: 50,
            message: "账号必须是1到50个字符",
            trigger: "blur",
        },
    ],
});
//点击修改药物信息
function modifyDrug(formName) {
    this.$refs[formName].validate((valid) => {
        if (valid) {
            getModifyDrug({
                drId: modifyForm.drId,
                drName: modifyForm.drName,
                drNumber: modifyForm.drNumber,
                drPrice: modifyForm.drPrice,
                drUnit: modifyForm.drUnit,
                drPublisher: modifyForm.drPublisher,

            }).then((res) => {
                if (res.data.status !== 200)
                    return ElMessage.error("修改信息失败！");
                modifyFormVisible.value = false;
                requestDrugs();
                ElMessage.success("修改药物信息成功！");
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
    getFindDrug({
        drId: id,
    }).then((res) => {
        if (res.data.status !== 200)
            return ElMessage.error("请求数据失败");
        modifyForm = res.data.data;
        modifyFormVisible.value = true;
        console.log(res);
    });
}
//删除药物操作
function deleteDrug(id) {
    getDeleteDrug({
        drId: id,
    }).then((res) => {
        requestDrugs();
        console.log(res);
    });
}
//删除对话框
function deleteDialog(id) {
    ElMessageBox.confirm("此操作将删除该药物信息, 是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
    })
        .then(() => {
            deleteDrug(id);
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
function addDrug(formName) {
    this.$refs[formName].validate((valid) => {
        if (valid) {
            getAddDrug({
                drId: addForm.drId,
                drName: addForm.drName,
                drNumber: addForm.drNumber,
                drPrice: addForm.drPrice,
                drUnit: addForm.drUnit,
                drPublisher: addForm.drPublisher,

            }).then((res) => {
                if (res.data.status !== 200)
                    return ElMessage.error(
                        "编号不合法或已被占用！"
                    );
                addFormVisible.value = false;
                requestDrugs();
                ElMessage.success("增加医生成功！");
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
    size.value = size1;
    requestDrugs();
}
//   页码改变时触发
function handleCurrentChange(num) {
    console.log(num);
    pageNumber.value = num;
    requestDrugs();
}
// 加载医生列表
function requestDrugs() {
    console.log(pageNumber, size, query);
    getFindAllDrugs({
        pageNumber: pageNumber,
        size: size,
        query: query,

    }).then((res) => {
        drugData = res.data.data.drugs;
        total.value = res.data.data.total;
        console.log(res.data.data);
    });
}
onMounted(() => {
    requestDrugs();
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