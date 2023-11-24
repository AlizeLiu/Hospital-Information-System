<template>
    <el-card class="box-card">
        <div class="left"></div>
        <div class="right">
            <h3>欢迎登录医院管理系统</h3>
            <el-form :model="form" label-width="120px">
                <el-form-item label="用户名：">
                    <el-input v-model="form.username" class="w-50 m-2" placeholder="请输入账号">
                        <template #prefix>
                            <el-icon class="el-input__icon">
                                <User />
                            </el-icon>
                        </template>
                    </el-input>
                </el-form-item>
                <el-form-item label="密码：">
                    <el-input v-model="form.password" class="w-50 m-2" placeholder="请输入密码" show-password>
                        <template #prefix>
                            <el-icon class="el-input__icon">
                                <Lock />
                            </el-icon>
                        </template>
                    </el-input>
                </el-form-item>
                <el-form-item label="角色：">
                    <el-radio-group v-model="form.role" change="">
                        <el-radio label="patient">患者</el-radio>
                        <el-radio label="doctor">医生</el-radio>
                        <el-radio label="admin">管理员</el-radio>
                    </el-radio-group>
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" @click="onSubmit">登录</el-button>
                </el-form-item>
            </el-form>
            <div class="desc">
                还没有账号？请<a href="">注册</a>
            </div>
        </div>


    </el-card>
</template>

<script setup lang="ts">
import { ref, onMounted, reactive } from 'vue'
// import { Calendar, Search } from '@element-plus/icons-vue'
import { getlogin } from '@/api/login'
import router from '@/router';

const form = reactive({
    username: '',
    password: '',
    role: 'admin'
})

// onMounted(() => {
    
// });

const onSubmit = () => {
    console.log(form)
    getlogin(form).then((res) => {
        console.log(res);
        sessionStorage.setItem('TOKEN_KEY',res.token)
        router.push('/index')
    });
    // router.push('/index')
}



</script>

<style scoped lang="less">
.box-card {
    width: 480px;
    margin: 100px auto;
}
</style>