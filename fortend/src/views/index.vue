<template>
  <div class="main common-layout">
    <el-container>
      <el-header>
        <div class="head-bar">
          <div class="head-left">
            <div class="header-ico">
              <!--      <i class="el-icon-s-home"></i>-->
              <img src="@/assets/img/1.png">
            </div>
            <div class="logo">医院管理系统</div>
          </div>
          <div class="head-right">
            <div class="head-user-con">
              <div class="user-avatar">
                <img src="../assets/11.jpg" />
              </div>
              <el-dropdown @command="handleCommand" class="user-name" trigger="click">
                <span class="el-dropdown-link">
                  <span>欢迎您，<b>{{ userName }}</b>&nbsp;{{ roleValue }}&nbsp;</span>
                  <el-icon>
                    <ArrowDown />
                  </el-icon>
                </span>
                <template #dropdown>
                  <el-dropdown-menu>
                    <el-dropdown-item command="logout">退出登录</el-dropdown-item>
                  </el-dropdown-menu>
                </template>
              </el-dropdown>
            </div>
          </div>
        </div>
      </el-header>
      <el-container>
        <el-aside width="200px"><!-- 导航菜单 -->
          <el-menu background-color="#353744" text-color="#fff" active-text-color="#ffd04b" :default-active="activePath">
            <template v-for="(item, index) in menuList" :key="index">
              <template v-if="!item.meta.hidden">
                <div>
                  <el-menu-item :index="item.path" @click="menuClick(item.path)">
                    <i class="el-icon-s-home" style="font-size: 18px;"> {{ item.meta.title }}</i>
                  </el-menu-item>
                </div>
              </template>

            </template>
            <!-- <el-menu-item index="adminLayout" @click="menuClick('adminLayout')">
              <i class="el-icon-s-home" style="font-size: 18px;"> 首页</i>
            </el-menu-item> -->
            <!-- <el-menu-item index="doctorList" @click="menuClick('doctorList')">
              <i class="el-icon-user" style="font-size: 18px;"> 医生信息管理</i>
            </el-menu-item>
            <el-menu-item index="patientList" @click="menuClick('patientList')">
              <i class="el-icon-user-solid" style="font-size: 18px;"> 患者信息管理</i>
            </el-menu-item>
            <el-menu-item index="orderList" @click="menuClick('orderList')">
              <i class="el-icon-postcard" style="font-size: 18px;"> 菜单管理</i>
            </el-menu-item>
            <el-menu-item index="oderOperate" @click="menuClick('orderOperate')">
              <i class="el-icon-monitor" style="font-size: 18px">
                预约挂号</i>
            </el-menu-item>
            <el-menu-item index="myOrder" @click="menuClick('myOrder')">
              <i class="el-icon-postcard" style="font-size: 18px">
                我的挂号</i>
            </el-menu-item>
            <el-menu-item index="orderToday" @click="menuClick('orderToday')">
              <i class="el-icon-news" style="font-size: 18px;"> 今日挂号列表</i>
            </el-menu-item>
            <el-menu-item index="doctorOrder" @click="menuClick('doctorOrder')">
              <i class="el-icon-monitor" style="font-size: 18px;"> 历史挂号列表</i>
            </el-menu-item>
            <el-menu-item index="inBed" @click="menuClick('inBed')">
              <i class="el-icon-postcard" style="font-size: 18px;"> 住院申请管理</i>
            </el-menu-item> -->
          </el-menu>
        </el-aside>
        <el-main>
          <router-view></router-view>
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>
<script setup>
import { ref, reactive, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router';
import { getInfo, } from '@/api/index.ts'
import { getToken, clearToken, getActivePath, setActivePath, } from "@/utils/storage.ts";
import { ElMessage, ElMessageBox } from 'element-plus'
// import { arrAdmin } from '@/router/list'
import { arrAdmin } from '@/router/admin.ts'
import { arrPatient } from '@/router/patient.ts'
import { arrDoctor } from '@/router/doctor.ts'

let route = useRoute()
let router = useRouter()

let userName = ref('');
let activePath = ref('');
let roleValue = ref('');
let role = localStorage.getItem('role')

let menuList = reactive([]);
if (role == 'admin') {
  menuList =arrAdmin
} else if (role == 'doctor') {
  menuList =arrDoctor
} else if (role == 'patient') {
  menuList =arrPatient
}
onMounted(() => {
  requestInfo()
  if (role == 'admin') {
    roleValue = '管理员'
  } else if (role == 'doctor') {
    roleValue = '医生'
  } else if (role == 'patient') {
    roleValue = '患者'
  }
})

const menuClick = (url) => {
  router.push(url)
}
function requestInfo() {
  getInfo({
    userId: route.query.userId,
    role: localStorage.getItem('role')
  }).then(res => {
    // console.log('res==', res);
    if (res.status != 200)
      return ElMessage.error("请重新登录！");
    userName.value = res.data.username;
    // role.value = route.query.role;
  }).catch(err => {
    ElMessage.error("用户名或密码错误");
    console.error(err);
  })
}
function handleCommand(command) {
  if (command === "logout") {
    ElMessageBox.confirm("此操作将退出登录, 是否继续?", "提示", {
      confirmButtonText: "确定",
      cancelButtonText: "取消",
      type: "warning",
    })
      .then(() => {
        clearToken();
        localStorage.clear();
        ElMessage({
          type: "success",
          message: "退出登录成功!",
        });
        router.push({
          path:"login",
          query:{
            date:new Date().getTime()
          }
        });
      
      })
      .catch(() => {
        ElMessage({
          type: "info",
          message: "已取消",
        });
      });

  }
}

</script>
<style lang="less">
.title {
  cursor: pointer;
}

.el-header {
  background-color: #427cb3;
  display: flex;
  justify-content: space-between;
  align-items: center;

  .words {
    text-align: center;

    span {
      color: black;
    }
  }

  //border-bottom: 1px solid lightgrey;
}

.el-container {
  height: 100vh;
}

.el-aside {
  background-color: #353744;
  border-right: 1px solid lightgrey;
  width: 150px;
}

.el-menu {
  border: 0;
}

.head-bar {
  position: relative;
  box-sizing: border-box;
  width: 100%;
  height: 70px;
  font-size: 22px;
  color: #fff;

}

.header-ico {
  float: left;
  padding: 0 21px;
  line-height: 70px;
  width: 71px;
  height: 55px;
  margin-left: -25px;
  margin-top: 5px;

  img {
    width: 100%;
    height: 100%;
  }

}

.head-bar .logo {
  float: left;
  width: 250px;
  line-height: 70px;
  margin-left: -25px;
}

.head-right {
  float: right;
  padding-right: 50px;
}

.head-user-con {
  display: flex;
  height: 70px;
  align-items: center;
}

.btn-fullscreen {
  transform: rotate(45deg);
  margin-right: 5px;
  font-size: 24px;
}

.btn-fullscreen {
  position: relative;
  width: 30px;
  height: 30px;
  text-align: center;
  border-radius: 15px;
  cursor: pointer;
}

.btn-bell .el-icon-bell {
  color: #fff;
}

.user-name {
  margin-left: 10px;
}

.user-avatar {
  margin-left: 20px;
}

.user-avatar img {
  display: block;
  width: 40px;
  height: 40px;
  border-radius: 50%;
}

.el-dropdown-link {
  color: #fff;
  cursor: pointer;
}

.el-dropdown-menu__item {
  text-align: center;
}
</style>