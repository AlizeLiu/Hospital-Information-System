<template>
    <!-- 卡片 -->
    <el-card>
        <!-- 搜索栏 -->
        <el-row type="flex">
            <el-col :span="6">
                <!-- <el-input v-model="query" placeholder="请输入姓名查询">
                    <template #append>
                        <el-button :icon="Search" @click="requestRoles"></el-button>
                    </template>
                </el-input> -->
                <el-button type="primary" @click="addFormVisible = true" style="font-size: 18px;">
                    <el-icon style="font-size: 22px;">
                        <CirclePlusFilled />
                    </el-icon>
                    增加角色
                </el-button>
            </el-col>
        </el-row>
        <!-- 表格 -->
        <el-table :data="roleData" stripe style="width: 100%" border>
            <el-table-column label="序号" type="index" width="100"></el-table-column>
            <el-table-column prop="roleName" label="角色名称" width="200">
            </el-table-column>
            <el-table-column prop="role" label="权限字符">
            </el-table-column>
            <el-table-column prop="startTime" label="创建时间">
            </el-table-column>
            <el-table-column prop=state label="是否使用" width="180">
                <template #default="scope">
                    <el-tag type="success" v-if="scope.row.state">是</el-tag>
                    <el-tag type="danger" v-else>否</el-tag>
                </template>
            </el-table-column>
            <el-table-column label="操作" width="200" fixed="right">
                <template #default="scope">
                    <el-button :icon="Delete" style="font-size: 14px" type="danger"
                        @click="deleteDialog(scope.row.rId)"></el-button>
                </template>
            </el-table-column>
        </el-table>

        <!-- 分页 -->
        <!-- <el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange" background
            layout="total, sizes, prev, pager, next, jumper" :current-page="pageNumber" :page-size="size"
            :page-sizes="[1, 2, 4, 8, 16]" :total="total">
        </el-pagination> -->
        <el-dialog title="增加角色" v-model="addFormVisible">
            <el-form :model="addForm" ref="ruleForm">
                <el-form-item label="角色名称" label-width="100px" prop="roleName">
                    <el-input v-model="addForm.roleName" autocomplete="off"></el-input>
                </el-form-item>
                <el-form-item label="权限字符" label-width="100px">
                    <el-input v-model="addForm.role" autocomplete="off"></el-input>
                </el-form-item>
                <el-form-item label="使用状态" label-width="100px">
                    <el-radio v-model="addForm.state" :label="1">使用</el-radio>
                    <el-radio v-model="addForm.state" :label="0">不使用</el-radio>
                </el-form-item>
                <el-form-item label="菜单权限" label-width="100px">
                    <el-checkbox-group v-model="addForm.menu">
                        <el-checkbox v-for="(item, index) in menuList" :key="index" :label="item.title" />
                    </el-checkbox-group>
                </el-form-item>
            </el-form>
            <template #footer>
                <div class="dialog-footer">
                    <el-button @click="addFormVisible = false" style="font-size: 18px;">
                        <el-icon style="font-size: 20px;">
                            <Close />
                        </el-icon>取 消</el-button>
                    <el-button type="primary" @click="addRole(ruleForm)" style="font-size: 18px;">
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
import {
    Check,
    Delete, Close, Search
} from '@element-plus/icons-vue'
import { getFindMenuList } from '@/api/admin/menuList.ts'
import { getAddRole, getFindRoleList, getDeleteRole } from '@/api/admin/roleList.ts'
import { arrList } from '@/router/list.ts'

let pageNumber = ref(1)
let size = ref(8)
let query = ref('')
let total = ref(3)
let addFormVisible = ref(false)
let ruleForm = ref(null)
let addForm = reactive({
    menu: []
});
let menuList = reactive([])
let roleData = ref([
    {
        roleName: "医生",
        role: 'doctor',
        startTime: '2023-01-20 09:00',
        state: 1,
        rId:1
    },
    {
        roleName: "患者",
        role: 'patient',
        startTime: '2023-01-20 09:00',
        state: 1,
        rId:2
    },
    {
        roleName: "管理员",
        role: 'admin',
        startTime: '2023-01-20 09:00',
        state: 1,
        rId:3
    },
]);
//删除病人操作
function deleteMenu(id) {
    getDeleteRole({
        rId: id,
    })
        .then((res) => {
            requestRoles();
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
function addRole(formName) {
    ruleForm.value.validate((valid) => {
        if (valid) {
            console.log(addForm);
            getAddRole({
                roleName: addForm.roleName,
                role: addForm.role,
                menu:addForm.menu,
                state: addForm.state
            })
                .then((res) => {
                    if (res.status !== 200)
                        return ElMessage.error(
                            "请重新添加角色！"
                        );
                    addFormVisible.value = false;
                    requestRoles();
                    ElMessage.success("增加角色成功！");
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
    requestRoles();
}
//   页码改变时触发
function handleCurrentChange(num) {
    console.log(num);
    pageNumber.value = num;
    requestRoles();
}
// 加载角色列表
function requestRoles() {
    getFindRoleList()
        .then((res) => {
            roleData.value = res.data
            console.log(res);
        });
}
function requestMenu() {
    let arr = []
    getFindMenuList()
        .then((res) => {
            console.log(res);
            arrList.map(item => {
                arr.push({
                    title: item.meta.title,
                    path: item.path
                })
            })
        });
    menuList = arr;
}
onMounted(() => {
    requestRoles();
    requestMenu()
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