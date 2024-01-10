<template>
    <!-- 卡片 -->
    <el-card>
        <!-- 搜索栏及增加医生 -->
        <el-row type="flex">
            <el-col :span="8">
                <el-input v-model="query" placeholder="请输入姓名查询">
                    <template #append>
                        <el-button @click="requestDoctors">
                            <el-icon>
                                <Search />
                            </el-icon>
                        </el-button>
                    </template>
                </el-input>
            </el-col>
            <el-col :span="1"></el-col>
            <el-col :span="4">
                <el-button type="primary" @click="addFormVisible = true" style="font-size: 18px;">
                    <el-icon style="font-size: 22px;">
                        <CirclePlusFilled />
                    </el-icon>
                    增加医生
                </el-button>
            </el-col>
            <el-col :span="1"></el-col>

        </el-row>
    
        <!-- 表格 -->
        <el-table :data="doctorData" stripe style="width: 100%" border>
            <el-table-column type="index" width="100" label="序号" />
            <el-table-column prop="dId" label="账号" width="100">
            </el-table-column>
            <el-table-column prop="dName" label="姓名" width="80">
            </el-table-column>
            <el-table-column prop="dGender" label="性别" width="60">
            </el-table-column>
            <el-table-column prop="dPost" label="职位" width="100">
            </el-table-column>
            <el-table-column prop="dSection" label="科室" width="100">
            </el-table-column>
            <el-table-column prop="dCard" label="证件号">
            </el-table-column>
            <el-table-column prop="dPhone" label="手机号">
            </el-table-column>
            <el-table-column prop="dEmail" label="邮箱" width="170">
            </el-table-column>
            <el-table-column prop="dAvgStar" label="评分/5分" width="80">
            </el-table-column>
            <el-table-column prop="dPrice" label="挂号费/元" width="100">
            </el-table-column>
            <el-table-column prop="dState" label="是否在职" width="80">
                <template #default="scope">
                    <el-tag type="success" v-if="scope.row.dState === 1">在职</el-tag>
                    <el-tag type="danger" v-else>离职</el-tag>
                </template>
            </el-table-column>
            <el-table-column label="操作" width="200" fixed="right">
                <template #default="scope">
                    <el-button style="font-size: 14px;" type="success" @click="modifyDialog(scope.row)"
                        :icon="Edit"></el-button>
                    <el-button style="font-size: 14px;" type="danger" @click="deleteDialog(scope.row.dId)"
                        :icon="Delete"></el-button>
                </template>
            </el-table-column>
        </el-table>

        <!-- 分页 -->
        <el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange" background
            layout="total, sizes, prev, pager, next, jumper" :current-page="pageNumber" :page-size="size"
            :page-sizes="[1, 2, 4, 8, 16]" :total="total">
        </el-pagination>
        <!-- 增加医生对话框 -->
        <el-dialog title="增加医生" v-model="addFormVisible">
            <el-form :model="addForm" :rules="rules" ref="ruleForm">
                <el-form-item label="账号" label-width="80px" prop="dId">
                    <el-input v-model.number="addForm.dId" autocomplete="off"></el-input>
                </el-form-item>
                <el-form-item label="密码" label-width="80px">
                    <el-input v-model="addForm.dPassword" autocomplete="off" disabled></el-input>
                </el-form-item>
                <el-form-item label="姓名" label-width="80px" prop="dName">
                    <el-input v-model="addForm.dName" autocomplete="off"></el-input>
                </el-form-item>
                <el-form-item label="性别" label-width="80px">
                    <el-radio v-model="addForm.dGender" label="男">男</el-radio>
                    <el-radio v-model="addForm.dGender" label="女">女</el-radio>
                </el-form-item>
                <el-form-item label="职位" label-width="80px" prop="dPost">
                    <el-select v-model="addForm.dPost" placeholder="请选择职称">
                        <el-option v-for="post in posts" :key="post" :label="post" :value="post">
                        </el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="科室" label-width="80px" prop="dSection">
                    <el-select v-model="addForm.dSection" filterable placeholder="请选择科室">
                        <el-option v-for="section in sections" :key="section" :label="section" :value="section">
                        </el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="身份证号" label-width="80px" prop="dCard">
                    <el-input v-model="addForm.dCard" autocomplete="off"></el-input>
                </el-form-item>
                <el-form-item label="手机号" label-width="80px" prop="dPhone">
                    <el-input v-model="addForm.dPhone" autocomplete="off"></el-input>
                </el-form-item>
                <el-form-item label="邮箱" label-width="80px" prop="dEmail">
                    <el-input v-model="addForm.dEmail" autocomplete="off"></el-input>
                </el-form-item>
                <el-form-item label="挂号费" label-width="80px" prop="dPrice">
                    <el-input v-model="addForm.dPrice" autocomplete="off"></el-input>
                </el-form-item>
                <el-form-item label="简介" label-width="80px" prop="dIntroduction">
                    <el-input type="textarea" :rows="5" placeholder="请输入内容" v-model="addForm.dIntroduction">
                    </el-input>
                </el-form-item>
            </el-form>
            <template #footer>
                <div class="dialog-footer">
                    <el-button @click="addFormVisible = false" style="font-size: 18px;">
                        <el-icon style="font-size: 20px;">
                            <Close />
                        </el-icon>取 消</el-button>
                    <el-button type="primary" @click="addDoctor(ruleForm)" style="font-size: 18px;">
                        <el-icon style="font-size: 20px;">
                            <Check />
                        </el-icon>确 定</el-button>
                </div>
            </template>
        </el-dialog>

        <!-- 修改医生对话框 -->
        <el-dialog title="修改医生信息" v-model="modifyFormVisible">
            <el-form :model="modifyForm" :rules="rules" ref="ruleForm2">
                <el-form-item label="账号" label-width="80px" prop="dId">
                    <el-input v-model.number="modifyForm.dId" autocomplete="off" disabled></el-input>
                </el-form-item>

                <el-form-item label="姓名" label-width="80px" prop="dName">
                    <el-input v-model="modifyForm.dName" autocomplete="off"></el-input>
                </el-form-item>
                <el-form-item label="性别" label-width="80px">
                    <el-radio v-model="modifyForm.dGender" label="男">男</el-radio>
                    <el-radio v-model="modifyForm.dGender" label="女">女</el-radio>
                </el-form-item>
                <el-form-item label="职位" label-width="80px" prop="dPost">
                    <el-select v-model="modifyForm.dPost" placeholder="请选择职称">
                        <el-option v-for="post in posts" :key="post" :label="post" :value="post">
                        </el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="科室" label-width="80px" prop="dSection">
                    <el-select v-model="modifyForm.dSection" filterable placeholder="请选择科室">
                        <el-option v-for="section in sections" :key="section" :label="section" :value="section">
                        </el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="身份证号" label-width="80px" prop="dCard">
                    <el-input v-model="modifyForm.dCard" autocomplete="off"></el-input>
                </el-form-item>
                <el-form-item label="手机号" label-width="80px" prop="dPhone">
                    <el-input v-model="modifyForm.dPhone" autocomplete="off"></el-input>
                </el-form-item>
                <el-form-item label="邮箱" label-width="80px" prop="dEmail">
                    <el-input v-model="modifyForm.dEmail" autocomplete="off"></el-input>
                </el-form-item>
                <el-form-item label="挂号费" label-width="80px" prop="dPrice">
                    <el-input v-model="modifyForm.dPrice" autocomplete="off"></el-input>
                </el-form-item>
                <el-form-item label="简介" label-width="80px" prop="dIntroduction">
                    <el-input type="textarea" :rows="5" placeholder="请输入内容" v-model="modifyForm.dIntroduction">
                    </el-input>
                </el-form-item>
                <el-form-item label="是否在职" label-width="80px" prop="dState">
                    <el-radio-group v-model="modifyForm.dState" class="ml-4">
                        <el-radio :label="1" size="large">在职</el-radio>
                        <el-radio :label="0" size="large">离职</el-radio>
                    </el-radio-group>
                </el-form-item>
            </el-form>
            <template #footer>
                <div class="dialog-footer">
                    <el-button @click="modifyFormVisible = false" style="font-size: 18px;"><i class="el-icon-close"
                            style="font-size: 20px;"></i> 取 消</el-button>
                    <el-button type="primary" @click="modifyDoctor" style="font-size: 18px;"><i
                            class="el-icon-check" style="font-size: 20px;"></i> 确 定</el-button>
                </div>
            </template>
        </el-dialog>
    </el-card>
</template>
<script setup>
import { ref, reactive, onMounted, } from 'vue'
import request from '@/utils/request.ts'
import { useRouter } from 'vue-router';
import { ElMessage, ElMessageBox } from 'element-plus'
import { Check, Delete, Close, Search, Edit } from '@element-plus/icons-vue'
import { getModifyDoctor, getFindDoctor, getDeleteDoctor, getAddDoctor, getFindAllDoctors } from '@/api/admin/doctorList.ts'
var validateMoblie = (rule, value, callback) => {
    if (value === undefined) {
        callback(new Error("请输入手机号"));
    } else {
        let reg =
            /^1(3[0-9]|4[5,7]|5[0,1,2,3,5,6,7,8,9]|6[2,5,6,7]|7[0,1,7,8]|8[0-9]|9[1,8,9])\d{8}$/;
        if (!reg.test(value)) {
            callback(new Error("请输入合法的手机号"));
        }
        callback();
    }
};
var validateCard = (rule, value, callback) => {
    if (value === undefined) {
        callback(new Error("请输入身份证号"));
    } else {
        let reg = /(^\d{15}$)|(^\d{18}$)|(^\d{17}(\d|X|x)$)/;
        if (!reg.test(value)) {
            callback(new Error("请输入合法的身份证号码"));
        }
        callback();
    }
};
let fileList = reactive([]);
let pageNumber = ref(1)
let size = ref(8)
let query = ref('')
let total = ref(3)
let ruleForm = ref()
let ruleForm2 = ref()
let doctorData = ref([]);
let addFormVisible = ref(false)
let addForm = reactive({ dPassword: 123456, dGender: "男", });
let modifyFormVisible = ref(false)
let modifyForm = ref({});
let posts = reactive(["主任医师", "副主任医师", "主治医生"])
let sections = reactive([
    "神经内科",
    "内分泌科",
    "呼吸与危重症医学科",
    "消化内科",
    "心血管内科",
    "发热门诊",
    "手足外科",
    "普通外科",
    "肛肠外科",
    "神经外科",
    "骨科",
    "烧伤整形外科",
    "妇科",
    "产科",
    "儿科",
    "耳鼻咽喉科",
    "眼科",
    "中医科",
    "急诊科",
    "皮肤病科",
    "口腔科",
])
let rules = reactive({
    dId: [
        { required: true, message: "请输入账号", trigger: "blur" },
        // {
        //     type: "number",
        //     message: "账号必须数字类型",
        //     trigger: "blur",
        // },
    ],
    dName: [
        { required: true, message: "请输入姓名", trigger: "blur" },
        {
            min: 2,
            max: 5,
            message: "账号必须是2到5个字符",
            trigger: "blur",
        },
    ],
    dPost: [
        { required: true, message: "请选择职位", trigger: "blur" },
    ],
    dSection: [
        {
            required: true,
            message: "请选择所属科室",
            trigger: "blur",
        },
    ],
    dEmail: [
        { required: true, message: "请输入邮箱", trigger: "blur" },
        {
            type: "email",
            message: "请输入正确的邮箱地址",
            trigger: ["blur", "change"],
        },
    ],
    dPrice: [
        {
            required: true,
            message: "请输入挂号费",
            trigger: "blur",
        },
    ],
    dPhone: [{ validator: validateMoblie }],
    dCard: [{ validator: validateCard }],
    dIntroduction: [
        {
            required: true,
            message: "请输入您的个人简介",
            trigger: "blur",
        },
    ],
});
onMounted(() => {
    requestDoctors();
});
//导出医生信息
function exportDoctors() {
    window.location.href = "http://localhost:9999/doctor/downloadExcel";
}
//文件上传
function handleProgress() {
    ElMessage.warning("文件正在解析中！");
}
// 文件上传成功时的钩子
function handleSuccess() {
    ElMessage.success("数据导入成功！");
    requestDoctors();
}
function handleError() {
    //ElMessage.error("数据导入失败！");
    ElMessage.success("数据导入成功！");
    requestDoctors();
}
function handleExceed() {
    ElMessage.warning("当前限制选择 1 个文件");
}
//点击修改医生信息
function modifyDoctor() {
    console.log('modifyDoctor: ', ruleForm)
    ruleForm2.value.validate((valid) => {
        if (valid) {
            console.log('valid success!')
            getModifyDoctor(modifyForm.value)
                .then((res) => {
                    console.log(res);
                    if (res.status !== 200)
                        return ElMessage.error("修改信息失败！");
                    modifyFormVisible.value = false;
                    requestDoctors();
                    ElMessage.success("修改医生信息成功！");
                    console.log(res);
                });
        } else {
            console.log("error submit!!");
            return false;
        }
    });
}
//打开修改对话框
function modifyDialog(row) {
    modifyForm.value = row;
    modifyFormVisible.value = true;
    // getFindDoctor({
    //     dId: id,
    // })
    //     .then((res) => {
    //         if (res.status !== 200)
    //             ElMessage.error("请求数据失败");
    //         modifyForm.value = res.data.data;
    //         modifyFormVisible.value = true;
    //         console.log(res);
    //     });
}
//删除医生操作
function deleteDoctor(id) {
    getDeleteDoctor({
        dId: id,
    })
        .then((res) => {
            requestDoctors();
            // console.log('删除医生操作',res);
        });
}
//删除对话框
function deleteDialog(id) {
    ElMessageBox.confirm("此操作将删除该医生信息, 是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
    })
        .then(() => {
            deleteDoctor(id);
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
// 增加医生
function addDoctor() {
    ruleForm.value.validate((valid) => {
        if (valid) {
            getAddDoctor({
                dId: addForm.dId,// 
                dGender: addForm.dGender,
                dPassword: addForm.dPassword,
                dName: addForm.dName,
                dPost: addForm.dPost,
                dSection: addForm.dSection,
                dPhone: addForm.dPhone,
                dEmail: addForm.dEmail,
                dCard: addForm.dCard,
                dPrice: addForm.dPrice,
                dIntroduction: addForm.dIntroduction,
                // doctor: addForm,
                role:'doctor'
            })
                .then((res) => {
                    if (res.status !== 200)
                        return ElMessage.error(
                            "账号不合法或已被占用！"
                        );
                    addFormVisible.value = false;
                    requestDoctors();
                    ElMessage.success("增加医生成功！");
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
    size.value = size1;
    requestDoctors();
}
//   页码改变时触发
function handleCurrentChange(num) {
    console.log(num);
    pageNumber.value = num;
    requestDoctors();
}
// 加载医生列表
function requestDoctors() {
    getFindAllDoctors({
        pageNumber: pageNumber.value,
        size: size.value,
        query: query.value,
    })
        .then((res) => {
            console.log(res.data.doctors);
            doctorData.value = res.data.doctors;
            total.value = res.data.total;
            // console.log(res.data.data.map((item) => item.doctor));
            // console.log(res.data.data.doctors.map((item) => item.dId));
        });
}

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