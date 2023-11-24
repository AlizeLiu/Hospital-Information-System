<template>
    <!-- 卡片 -->
    <el-card>
        <!-- 搜索栏 -->
        <!-- <el-row type="flex">
            <el-col :span="6">
                <el-button type="primary" @click="addFormVisible = true" style="font-size: 18px;">
                    <el-icon style="font-size: 22px;">
                        <CirclePlusFilled />
                    </el-icon>
                    增加菜单
                </el-button>
            </el-col>
        </el-row> -->
        <!-- 表格 -->
        <el-table :data="menuData" stripe style="width: 100%" border>
            <el-table-column prop="meta.title" label="菜单名称" width="200">
            </el-table-column>
            <el-table-column prop="path" label="路径">
            </el-table-column>
            <el-table-column prop="component" label="组件路径">
            </el-table-column>
            <el-table-column prop="meta.hidden" label="是否显示" width="180">
                <template #default="scope">
                    <el-tag type="success" v-if="!scope.row.meta.hidden">显示</el-tag>
                    <el-tag type="danger" v-else>隐藏</el-tag>
                </template>
            </el-table-column>
            <el-table-column prop="meta.requireAuth" label="是否需要路由校验" width="150">
                <template #default="scope">
                    <el-tag type="success" v-if="scope.row.meta.requireAuth">需要</el-tag>
                    <el-tag type="danger" v-else>不需要</el-tag>
                </template>
            </el-table-column>
            <el-table-column label="操作" width="200" fixed="right">
                <template #default="scope">
                    <el-button :icon="Delete" style="font-size: 14px" type="danger"
                        @click="deleteDialog(scope.row.mId)"></el-button>
                </template>
            </el-table-column>
        </el-table>

        <!-- 分页 -->
        <!-- <el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange" background
            layout="total, sizes, prev, pager, next, jumper" :current-page="pageNumber" :page-size="size"
            :page-sizes="[1, 2, 4, 8, 16]" :total="total">
        </el-pagination> -->
        <el-dialog title="增加菜单" v-model="addFormVisible">
            <el-form :model="addForm" ref="ruleForm">
                <el-form-item label="菜单名称" label-width="150px" prop="meta.title">
                    <el-input v-model="addForm.meta.title" autocomplete="off"></el-input>
                </el-form-item>
                <el-form-item label="菜单路径" label-width="150px">
                    <el-input v-model="addForm.path" autocomplete="off"></el-input>
                </el-form-item>
                <el-form-item label="组件路径" label-width="150px" prop="component">
                    <el-input v-model="addForm.component" autocomplete="off"></el-input>
                </el-form-item>
                <el-form-item label="是否显示" label-width="150px">
                    <el-radio v-model="addForm.meta.hidden" :label="true">隐藏</el-radio>
                    <el-radio v-model="addForm.meta.hidden" :label="false">显示</el-radio>
                </el-form-item>
                <el-form-item label="是否需要路由校验" label-width="150px">
                    <el-radio v-model="addForm.meta.requireAuth" :label="true">需要</el-radio>
                    <el-radio v-model="addForm.meta.requireAuth" :label="false">不需要</el-radio>
                </el-form-item>
            </el-form>
            <template #footer>
                <div class="dialog-footer">
                    <el-button @click="addFormVisible = false" style="font-size: 18px;">
                        <el-icon style="font-size: 20px;">
                            <Close />
                        </el-icon>取 消</el-button>
                    <el-button type="primary" @click="addMenu(ruleForm)" style="font-size: 18px;">
                        <el-icon style="font-size: 20px;">
                            <Check />
                        </el-icon>确 定</el-button>
                </div>
            </template>
        </el-dialog>
    </el-card>
</template>
<script setup>
import { ref, reactive, onMounted, } from 'vue'
import request from '@/utils/request.ts'
import { useRoute } from 'vue-router';
import { ElMessage, ElMessageBox } from 'element-plus'
import { Check, Delete, Close, Search } from '@element-plus/icons-vue'
import { getFindMenuList, getAddMenu, getDeleteMenu } from '@/api/admin/menuList.ts'
import { arrList } from '@/router/list.ts'

let pageNumber = ref(1)
let size = ref(8)
let query = ref('')
let total = ref(3)
let addFormVisible = ref(false)
let ruleForm = ref(null)
let addForm = reactive({
    meta: {
        title: '',
        hidden: true,
        requireAuth: true
    }
});
let menuData = ref([
    {
        path: "/patientList",
        component: 'admin/PatientList',
        meta: {
            title: '患者信息管理',
            requireAuth: true,  // 该路由项需要权限校验
        },
        mId: 1,
    },
]);
//删除病人操作
function deleteMenu(id) {
    getDeleteMenu({
        mId: id,
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
            deleteMenu(id);
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
function addMenu(formName) {
    ruleForm.value.validate((valid) => {
        if (valid) {
            console.log(addForm);
            getAddMenu(addForm)
                .then((res) => {
                    if (res.status !== 200)
                        return ElMessage.error(
                            "请重新添加菜单！"
                        );
                    addFormVisible.value = false;
                    requestDoctors();
                    ElMessage.success("增加菜单成功！");
                    // console.log(res);
                });
        } else {
            console.log("error submit!!");
            return false;
        }
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
    getFindMenuList()
        .then((res) => {
            // menuData.value = res.data
            menuData.value = arrList
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