<template>
    <div>
        <!-- 卡片 -->
        <el-card>
            <!-- 面包屑 -->
            <el-breadcrumb separator-class="el-icon-arrow-right">
                <el-breadcrumb-item :to="{ path: '/orderOperate' }">科室选择</el-breadcrumb-item>
                <!-- <el-breadcrumb-item>日期选择</el-breadcrumb-item> -->
                <el-breadcrumb-item>挂号</el-breadcrumb-item>
            </el-breadcrumb>

            <!-- 两边布局 -->
            <div class="head">
                <div>
                    <i class="iconfont icon-r-user1" style="margin: 5px; font-size: 26px">
                        {{ sectionOpt }}医生列表</i>
                        <h4 style="margin: 10px 0;" v-show="orderDate">挂号日期：{{ orderDate }}</h4>
                </div>

                <!-- 选择挂号时间 -->
                <div>
                    <div style="display: flex;align-items: center;">
                        <span>请选择你要挂号的日期：</span>
                        <ul class="dateUl">
                            <li v-for="monthDay in monthDays" :key="monthDay">
                                <el-button :icon="List" @click="dateClick(monthDay)">
                                    {{ monthDay }}</el-button>
                            </li>
                        </ul>
                    </div>
                    <div style="display: flex;align-items: center;">
                        <span>时间段：</span>
                        <ul class="dateUl">
                            <li v-for="timeDay in timeDays" :key="timeDay">
                                <el-button :icon="List" @click="timeClick(timeDay.times)" :disabled="timeDay.isDisable">
                                    {{ timeDay.times }}</el-button>
                            </li>
                        </ul>
                    </div>
                </div>
            </div>

            <!-- 表格 -->
            <el-table :data="sectionData" stripe style="width: 100%" border>
                <el-table-column type="index" label="序号" width="60"></el-table-column>
                <el-table-column prop="dId" label="工号" width="80">
                </el-table-column>
                <el-table-column prop="dName" label="姓名" width="80">
                </el-table-column>
                <el-table-column prop="dGender" label="性别" width="60">
                </el-table-column>
                <el-table-column prop="dPost" label="职位">
                </el-table-column>
                <el-table-column prop="dSection" labelx="科室"></el-table-column>
                <!-- <el-table-column prop="dIntroduction" label="简介">
                </el-table-column> -->
                <el-table-column prop="dPrice" label="挂号费用/元" width="100">
                </el-table-column>
                <el-table-column prop="dAvgStar" label="评价/5分" width="100">
                </el-table-column>
                <el-table-column prop="dNum" label="剩余挂号数量" width="100">
                </el-table-column>
                <el-table-column label="操作" width="140">
                    <template #default="scope">
                        <el-button class="iconfont icon-r-paper" style="font-size: 14px" type="warning"
                            @click="openClick(scope.row.dId, scope.row.dName)">
                            挂号</el-button>
                    </template>
                </el-table-column>
            </el-table>
        </el-card>
        <!-- 挂号对话框 -->
        <el-dialog title="填写挂号信息" v-model="orderFormVisible">
            <el-form :model="orderForm" ref="orderDorm" :rules="orderRules">
                <!-- <el-form-item label="挂号时间段" label-width="100px" prop="oTime">
                    <el-select v-model="orderForm.oTime" placeholder="请选择" no-data-text="请尝试预约明日">
                        <el-option v-for="time in times" :key="time.value" :label="time.label" :value="time.value">
                        </el-option>
                    </el-select>
                </el-form-item> -->
                <el-form-item label="挂号日期" label-width="100px">
                    <el-input v-model="orderForm.orderDate" autocomplete="off" disabled></el-input>
                </el-form-item>
                <el-form-item label="医生工号" label-width="100px">
                    <el-input v-model="orderForm.dId" autocomplete="off" disabled></el-input>
                </el-form-item>
                <el-form-item label="医生姓名" label-width="100px">
                    <el-input v-model="orderForm.dName" autocomplete="off" disabled></el-input>
                </el-form-item>
                <el-form-item label="患者姓名" label-width="100px">
                    <el-input v-model="orderForm.pName" autocomplete="off"></el-input>
                </el-form-item>
                <el-form-item label="患者身份证号" label-width="100px">
                    <el-input v-model="orderForm.pCard" autocomplete="off"></el-input>
                </el-form-item>
            </el-form>
            <template #footer>
                <div class="dialog-footer">
                    <el-button @click="orderFormVisible = false" style="font-size: 18px;"><i class="el-icon-close"
                            style="font-size: 20px;"></i> 取 消</el-button>
                    <el-button type="primary" @click="orderSuccess()" style="font-size: 18px;"><i class="el-icon-check"
                            style="font-size: 20px;"></i> 确 定</el-button>
                </div>
            </template>
        </el-dialog>
    </div>
</template>
<script setup>
import { ref, reactive, onMounted, } from 'vue'
import request from '@/utils/request.ts'
import { useRouter, useRoute } from 'vue-router';
import { ElMessage, ElMessageBox } from 'element-plus'
import { Check, Delete, Close, Search, Edit, List } from '@element-plus/icons-vue'
import { getFindOrderTime, getAddOrder, getFindByTime, getFindDoctorBySection, getFindNowTime } from '@/api/patient/sectionMessage.ts'

let route = useRoute()
let sectionOpt = ref('');
let sectionData = ref([
    {
        dId: 1,
        dName: '张',
        dGender: '男',
        dPost: '普通医生',
        dSection: '神经内科',
        dPrice: 13,
        dAvgStar: 5
    },
    {
        dId: 2,
        dName: '李医生',
        dGender: '女',
        dPost: '主任医生',
        dSection: '神经内科',
        dPrice: 13,
        dAvgStar: 5
    },
]);
let monthDays = reactive([]);
let timeDays = reactive([
    {
        times: '08:30-09:30',
        isDisable: false
    },
    {
        times: '09:30-10:30',
        isDisable: false
    },
    {
        times: '10:30-11:30',
        isDisable: false
    },
    {
        times: '14:30-15:30',
        isDisable: false
    },
    {
        times: '15:30-16:30',
        isDisable: false
    },
    {
        times: '16:30-17:30',
        isDisable: false
    },
]);
// let timeDays = reactive([
//     '08:30-09:30', '09:30-10:30', '10:30-11:30', '14:30-15:30', '15:30-16:30', '16:30-17:30'
// ]);
// 当前时间
// let dateYear=ref('')
let nowTime = ref('')
let clickTag = ref(false)
let orderFormVisible = ref(false)
let orderForm = reactive({});
let orderDorm = ref(null);
let times = ref([]);
let orderRules = reactive({
    oTime: [
        { required: true, message: "选择时间段", trigger: "blur" },
    ],
});
//挂号日期
let orderDate = ref('')
//拼接时间和日期成为oId
let idTime = ref('')

//打开挂号对话框触发,获取挂号时间段已剩余票数
async function requestTime(id) {
    idTime.value = id;
    console.log(idTime.value, orderDate.value);
    await getFindOrderTime({
        arId: idTime.value,
        data: orderDate.value
    }).then((res) => {
        const date = new Date(orderDate.value);
        const today = new Date();
        const isToday =
            date.getFullYear() === today.getFullYear() &&
            date.getMonth() === today.getMonth() &&
            date.getDate() === today.getDate();
        var array = [];
        if (!isTimeAfterTarget("09:30") || !isToday) {
            array.push({
                label: "08:30-09:30  " + "   余号 " + res.data.eTOn,
                value: "08:30-09:30  " + "   余号 " + res.data.eTOn
            });
        }
        if (!isTimeAfterTarget("10:30") || !isToday) {
            array.push(
                {
                    label: "09:30-10:30  " + "   余号 " + res.data.nTOt,
                    value: "09:30-10:30  " + "   余号 " + res.data.nTOt
                }
            );
        }
        if (!isTimeAfterTarget("11:30") || !isToday) {
            array.push(
                {
                    label: "10:30-11:30  " + "   余号 " + res.data.tTOe,
                    value: "10:30-11:30  " + "   余号 " + res.data.tTOe
                }
            );
        }
        if (!isTimeAfterTarget("15:30") || !isToday) {
            array.push(
                {
                    label: "14:30-15:30  " + "   余号 " + res.data.fTOf,
                    value: "14:30-15:30  " + "   余号 " + res.data.fTOf
                }
            );
        }
        if (!isTimeAfterTarget("16:30") || !isToday) {
            array.push(
                {
                    label: "15:30-16:30  " + "   余号 " + res.data.fTOs,
                    value: "15:30-16:30  " + "   余号 " + res.data.fTOs
                }
            );
        }
        if (!isTimeAfterTarget("17:30") || !isToday) {
            array.push(
                {
                    label: "16:30-17:30  " + "   余号 " + res.data.sTOs,
                    value: "16:30-17:30  " + "   余号 " + res.data.sTOs
                }
            );
        }
        times.value = array;
        console.log(times);
    });
}
function isTimeAfterTarget(timeString) {
    // 判断时间是否超过timeString(入参格式如：09:30)
    const currentTime = new Date();

    // 解析传入的目标时间字符串，获取小时和分钟
    const [targetHour, targetMinute] = timeString.split(":");

    // 设置要比较的时间
    const targetTime = new Date();
    targetTime.setHours(targetHour);
    targetTime.setMinutes(targetMinute);
    targetTime.setSeconds(0);
    console.log(targetTime);
    // 比较当前时间是否超过了目标时间
    return currentTime > targetTime;
}
//挂号点击确认
function orderSuccess() {
    orderDorm.value.validate((valid) => {
        if (valid) {
            getAddOrder({
                pId: localStorage.getItem("userid"),
                dId: orderForm.dId,
                oStart:
                    orderForm.orderDate +
                    " " +
                    orderForm.oTime,
                data: orderDate.value
            })
                .then((res) => {
                    if (res.status != 200)
                        return ElMessage.error(
                            "该时间段无剩余号源！请重新选择！"
                        );
                    orderFormVisible.value = false;
                    ElMessage.success("挂号成功！");
                    orderForm.oTime = '';
                });
        } else {
            console.log("error submit!!");
            return false;
        }
    });
}
//token解码
// tokenDecode(token) {
//     if (token !== null) return jwtDecode(token);
// },
//打开挂号对话框
function openClick(id, name) {
    orderForm.dId = id;
    orderForm.dName = name;
    //请求挂号时间段
    requestTime(id);
    orderFormVisible.value = true;

}
//选择日期触发时间
function dateClick(date) {
    //获取挂号年月日
    const nowDate = new Date();
    let year = nowDate.getFullYear();
    orderForm.orderDate = year + "-" + date;
    let dateYear = year + "-" + date;
    orderDate.value = dateYear;
    console.log('-----', dateYear, sectionOpt.value);
}
// 选择时间段出发
function timeClick(times){
    console.log(orderDate.value,'----------------------------');
    getFindByTime({
        arTime: orderDate.value,
        dSection: sectionOpt.value,
        times:times
    }).then((res) => {
        //this.sectionData.dId = res.data.data.doctors.dId;
        /**
         * 重点！！！把数组中的对象取出来用map
         */
        sectionData.value = res.data.map((item) => item.doctor);
        clickTag.value = true;
        // console.log(res.data.map((item) => item.doctor));
        //console.log(res.data.data[0].doctor);
    });
}
//获取当天及后7天的日期星期
function nowDay(num) {
    var nowDate = new Date();
    var currentHour = nowDate.getHours();
    var currentMinute = nowDate.getMinutes();
    // 判断当前时间是否已经过了17:30
    if (
        currentHour > 17 ||
        (currentHour === 17 && currentMinute > 30)
    ) {
        num++; // 次日
    }

    nowDate.setDate(nowDate.getDate() + num);
    var month = nowDate.getMonth() + 1;
    var date = nowDate.getDate();
    if (date < 10) {
        date = "0" + date;
    }
    if (month < 10) {
        month = "0" + month;
    }
    var time = month + "-" + date;
    monthDays.push(time);
    console.log(monthDays);
    orderDate.value = nowDate.getFullYear()+'-'+monthDays[0];
}
// 比较当天时间段
function isTimeTarget(timeO, timeS) {
    // 判断时间是否超过timeString(入参格式如：09:30)
    const currentTime = new Date();

    // 解析传入的第一个时间字符串，获取小时和分钟
    const [targetHourO, targetMinuteO] = timeO.split(":");
    const [targetHourS, targetMinuteS] = timeS.split(":");

    // 设置要比较的时间
    const targetTimeO = new Date();
    const targetTimeS = new Date();
    targetTimeO.setHours(targetHourO);
    targetTimeO.setMinutes(targetMinuteO);
    targetTimeO.setSeconds(0);
    targetTimeS.setHours(targetHourS);
    targetTimeS.setMinutes(targetMinuteS);
    targetTimeS.setSeconds(0);
    console.log(targetTimeO, targetTimeS);
    console.log(currentTime);
    // 比较当前时间是否超过了目标时间
    return  (currentTime > targetTimeO)&&(currentTime < targetTimeS);
}
//获取当天时间段
function getNowTime() {
    const date = new Date(orderDate.value);
    const today = new Date();
    var array = timeDays;
    // console.log(isTimeTarget("21:30", "22:30"));
    // console.log(isTimeTarget("22:30", "23:30"));
    if (isTimeTarget("08:30", "09:30")) {
        for (let i = 0; i < 0; i++) {
            array[i].isDisable = true
        }
    }
    if (isTimeTarget("09:30", "10:30")) {
        for (let i = 0; i < 1; i++) {
            array[i].isDisable = true
        }
    }
    if (isTimeTarget("10:30", "11:30")) {
        for (let i = 0; i < 2; i++) {
            array[i].isDisable = true
        }
    }
    // if(isTimeAfterTarget("11:30")){
    //     for (let i = 0; i < 3; i++) {
    //         array[i].isDisable = true
    //     }
    // }
    if (isTimeTarget("11:30", "15:30")) {
        for (let i = 0; i < 3; i++) {
            array[i].isDisable = true
        }
    }
    if (isTimeTarget("15:30", "16:30")) {
        for (let i = 0; i < 4; i++) {
            array[i].isDisable = true
        }
    }
    if (isTimeTarget("16:30", "17:30")) {
        for (let i = 0; i < 5; i++) {
            array[i].isDisable = true
        }
    }
    timeDays=array
}
//请求部门医生信息
function requestSection() {
    getFindDoctorBySection({
        dSection: route.query.sectionOpt
    })
        .then((res) => {
            if (res.status !== 200)
                return ElMessage.error("请求数据失败");
            sectionData.value = res.data.doctors;
            console.log(res);
        });
}

onMounted(() => {
    sectionOpt.value = route.query.sectionOpt;
    getNowTime()
    //获取当天的后7天
    for (var i = 0; i < 7; i++) {
        nowDay(i);
    }
    //按科室请求医生信息
    requestSection();
    // 解码token
    //     orderForm.pName = this.tokenDecode(getToken()).pName;
    //     orderForm.pCard = this.tokenDecode(getToken()).pCard;
    //     orderForm.pId = this.tokenDecode(getToken()).pId;
    //console.log(this.orderForm.pId)
})
</script>
<style scoped lang="scss">
.dateUl li {
    display: inline;
    //margin: 5px;
    padding: 1px;
}

.dateUl {
    margin: 10px;
    padding: 0;
}

.el-breadcrumb {
    margin: 8px;
}

.head {
    display: flex;
    justify-content: space-between;
    align-items: center;
}

.el-form {
    margin-top: 0;
}
</style>