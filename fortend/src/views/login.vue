<template>
    <div style="height: 100vh;
    display: flex; align-items: center;
    justify-content: center;
   background-color: #427cb3;;
  " class="login-wrap">
        <div style="display: flex; background-color: white; width: 50%; border-radius: 5px; overflow: hidden">
            <div style="flex: 1">
                <img src="@/assets/17.jpg" alt="" style="width: 400px;height: 400px" />
            </div>
            <div style="flex: 1; display: flex; align-items: center; justify-content: center">
                <el-form :model="loginForm" style="width: 80%" :rules="loginRules" ref="ruleForm">
                    <div style="font-size: 20px; font-weight: bold; text-align: center; margin-bottom: 20px">
                        欢迎登录医院管理系统
                    </div>
                    <el-form-item prop="id">
                        <!--必须绑定v-model输入框才能输入字符---->
                        <el-input v-model="loginForm.id">
                            <i slot="prefix" class="el-input__icon el-icon-user"></i>
                        </el-input>
                    </el-form-item>
                    <el-form-item prop="password">
                        <el-input v-model="loginForm.password" placeholder="请输入密码" clearable show-password>
                            <i slot="prefix" class="el-input__icon el-icon-lock"></i>
                        </el-input>
                    </el-form-item>
                    <el-form-item class="role">
                        <el-radio-group v-model="role" size="small">
                            <el-radio label="patient">患者</el-radio>
                            <el-radio label="doctor">医生</el-radio>
                            <el-radio label="admin">管理员</el-radio>
                        </el-radio-group>
                    </el-form-item>
                    <el-form-item>
                        <el-button type="primary" style="width: 100%" @click="submitLoginForm(ruleForm)">登 录</el-button>
                    </el-form-item>

                    <div style="display: flex">
                        <div style="flex: 1">
                            还没有账号？请
                            <span style="color: #0f9876; cursor: pointer" @click="registerFormVisible = true">注册</span>
                        </div>

                    </div>
                </el-form>
            </div>
        </div>
        <!-- 注册对话框 -->
        <el-dialog title="患者注册" v-model="registerFormVisible">
            <el-form class="findPassword" :model="registerForm" :rules="registerRules" ref="registerFormRule">
                <el-form-item label="账号" label-width="80px" prop="pId">
                    <el-input v-model.number="registerForm.pId"></el-input>
                </el-form-item>
                <el-form-item label="性别" label-width="80px">
                    <el-radio v-model="registerForm.pGender" label="男">男</el-radio>
                    <el-radio v-model="registerForm.pGender" label="女">女</el-radio>
                </el-form-item>
                <el-form-item label="密码" label-width="80px" prop="pPassword">
                    <el-input v-model="registerForm.pPassword"></el-input>
                </el-form-item>
                <el-form-item label="姓名" label-width="80px" prop="pName">
                    <el-input v-model="registerForm.pName"></el-input>
                </el-form-item>
                <el-form-item label="出生日期" label-width="80px" prop="pBirthday">
                    <el-date-picker v-model="registerForm.pBirthday" type="date" placeholder="选择日期"
                        value-format="yyyy-MM-dd">
                    </el-date-picker>
                </el-form-item>
                <el-form-item label="手机号" label-width="80px" prop="pPhone">
                    <el-input v-model="registerForm.pPhone"></el-input>
                </el-form-item>
                <el-form-item label="邮箱号" label-width="80px" prop="pEmail">
                    <el-input v-model="registerForm.pEmail"></el-input>
                </el-form-item>
                <el-form-item label="身份证号" label-width="80px" prop="pCard">
                    <el-input v-model="registerForm.pCard"></el-input>
                </el-form-item>
            </el-form>
            <template #footer>
                <div class="dialog-footer">
                    <el-button @click="registerFormVisible = false">取 消</el-button>
                    <el-button type="primary" @click="registerClick(registerForm)">确 定</el-button>
                </div>
            </template>
        </el-dialog>

    </div>
</template>
  
<script setup>
import { ref, reactive, onMounted,watch } from 'vue'
import { useRouter, useRoute } from 'vue-router';
import { ElMessage, ElMessageBox } from 'element-plus'
import { Check, Delete, Close, Search, Edit, List, ChatRound, Memo } from '@element-plus/icons-vue'
import { getlogin, getRegist, } from '@/api/login.ts'
import { setToken, } from '@/utils/storage.ts'


var validateMoblie = (rule, value, callback) => {
    if (value === undefined) {
        callback(new Error("请输入手机号"));
    } else {
        let reg = /^1(3[0-9]|4[5,7]|5[0,1,2,3,5,6,7,8,9]|6[2,5,6,7]|7[0,1,7,8]|8[0-9]|9[1,8,9])\d{8}$/;
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
var validatePass = (rule, value, callback) => {
    if (value === "") {
        callback(new Error("请输入密码"));
    } else {
        if (this.findForm.checkPassword !== "") {
            this.$refs.findForm.validateField("checkPassword");
        }
        callback();
    }
};
var validatePass2 = (rule, value, callback) => {
    if (value === "") {
        callback(new Error("请再次输入密码"));
    } else if (value !== this.findForm.newPassword) {
        callback(new Error("两次输入密码不一致!"));
    } else {
        callback();
    }
};
//背景图片
let backgroundDiv = reactive({
    // backgroundImage: "url(import.meta("../assets/doctor.jpg") + ')",
    backgroundImage: new URL("@/assets/doctor.jpg", import.meta.url),
    backgroundRepeat: "no-repeat",
    backgroundSize: "100% 100%"
})
let loginForm = reactive({
    id: "202301",
    password: "123456"
})
let loginRules = reactive({
    id: [
        { required: true, message: "请输入账号编号", trigger: "blur" },
        { min: 3, max: 50, message: "长度在 3到 50 个字符", trigger: "blur" }
    ],
    password: [{ required: true, message: "请输入密码", trigger: "blur" }]
})
let role = ref("patient")
let findRole = ref("patient")
let ruleForm = ref()
let registerFormRule = ref()
//找回密码
let findFormVisible = ref(false)
let findForm = reactive({
    code: "",
    newPassword: "",
    checkPassword: "",
    pEmail: ""
})

let findRules = reactive({
    pEmail: [
        { required: true, message: "请输入邮箱地址", trigger: "blur" },
        {
            type: "email",
            message: "请输入正确的邮箱地址",
            trigger: ["blur", "change"]
        }
    ],
    code: [{ required: true, message: "请输入验证码", trigger: "blur" }],
    newPassword: [{ validator: validatePass, trigger: "blur" }],
    checkPassword: [{ validator: validatePass2, trigger: "blur" }]
})
let totalTime = ref(60)
let content = ref("发送验证码")
let canClick = ref(true)
//注册
let registerFormVisible = ref(false)
let registerForm = reactive({
    pGender: "男"
})
let registerRules = reactive({
    pId: [
        { required: true, message: "请输入账号", trigger: "blur" },
        { type: "number", message: "账号必须数字类型", trigger: "blur" }
    ],
    pPassword: [
        { required: true, message: "请输入密码", trigger: "blur" },
        { min: 4, max: 50, message: "长度在 4到 50 个字符", trigger: "blur" }
    ],
    pName: [
        { required: true, message: "请输入姓名", trigger: "blur" },
        { min: 2, max: 8, message: "长度在 2到 8 个字符", trigger: "blur" }
    ],
    pEmail: [
        { required: true, message: "请输入邮箱", trigger: "blur" },
        {
            type: "email",
            message: "请输入正确的邮箱地址",
            trigger: ["blur", "change"]
        }
    ],
    pPhone: [{ validator: validateMoblie }],
    pCard: [{ validator: validateCard }],
    pBirthday: [
        { required: true, message: "选择出生日期", trigger: "blur" }
    ]
})

let router = useRouter()
let route = useRoute()
//点击注册确认按钮
function registerClick(formName) {
    registerFormRule.value.validate(valid => {
        console.log(registerForm);
        if (valid) {
            getRegist({
                pId: registerForm.pId,
                pName: registerForm.pName,
                pPassword: registerForm.pPassword,
                pGender: registerForm.pGender,
                pEmail: registerForm.pEmail,
                pPhone: registerForm.pPhone,
                pCard: registerForm.pCard,
                pBirthday: registerForm.pBirthday,
                role:'patient'
            })
                .then(res => {
                    if (res.status !== 200)
                        return ElMessage.error("账号或邮箱已被占用！");
                    registerFormVisible.value = false;
                    ElMessage.success("注册成功！");
                    console.log(res);
                });
        } else {
            console.log("error submit!!");
            return false;
        }
    });
}
//提交表单
function submitLoginForm(formName) {
    ruleForm.value.validate(valid => {
        if (valid) {
            console.log(role.value);
            getlogin({
                username: loginForm.id,
                password: loginForm.password,
                role: role.value
            }).then(res => {
                console.log(res);
                if (res.status != 200)
                    return ElMessage.error("用户名或密码错误");
                setToken(res.data.token);
                localStorage.setItem('role',role.value)
                router.push({
                    path:'/',
                    query:{
                        role:role.value,
                        userId:loginForm.id
                    }
                });
            }).catch(err => {
                    ElMessage.error("用户名或密码错误");
                    console.error(err);
                })
        } else {
            console.log("error submit!!");
            return false;
        }
    });
}
watch(route.query.date, (to, from) => {
  router.go(0)
  console.log(111);
})

</script>
  
<style lang="scss">
.login-wrap {
    position: relative;
    width: 100%;
    height: 100%;
    background: #0f9876;
    background-image: url("../assets/img/login-bg.svg");
    background-size: 100% 100%;
}
</style>
  