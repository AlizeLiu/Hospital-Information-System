<!--

-->
<template>
    <div>
        <el-card>
            <div>
                <div>
                    <el-icon style="margin: 10px; font-size: 22px"><Calendar />请选择值班日期：</el-icon>
                </div>
                <br />
                <ul class="dateUl">
                    <li v-for="monthDay in monthDays" :key="monthDay">
                        <el-button
                            :icon="Calendar"
                            type="primary"
                            style="margin: 5px"
                            @click="dateClick(monthDay)"
                        >
                            {{ monthDay }}</el-button
                        >
                    </li>
                </ul>
            </div>
            <div class="router-view">
                <router-view></router-view>
            </div>
        </el-card>
    </div>
</template>
<script>
import { getActivePath, setActivePath } from "@/utils/storage.js";
import {
    Calendar,
    Delete, Close, Search
} from '@element-plus/icons-vue'
const ARRANGEDATE = "arrangeDate";
export default {
    name: "ArrangeIndex",
    data() {
        return {
            monthDays: [],
            monthDay: "",
            activePath: "",
        };
    },

    methods: {
        //日历点击
        dateClick(monthDay) {
            console.log(monthDay);

            const nowDate = new Date();
            let year = nowDate.getFullYear();
            let dateTime = year + "-" + monthDay;
            sessionStorage.setItem(ARRANGEDATE, dateTime);

            this.activePath = "sectionIndex";
            setActivePath("sectionIndex");
            if (this.$route.path !== "/sectionIndex")
                this.$router.push("sectionIndex");
        },
        //获取当天及后7天的日期星期
        nowDay(num) {
            var nowDate = new Date();
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
            this.monthDays.push(time);
        },
    },
    created() {
        //获取当天的后7天
        for (var i = 0; i < 25; i++) {
            this.nowDay(i);
            //  获取激活路径
            this.activePath = getActivePath();
        }
    },
};
</script>
<style scoped lang="scss">
.disabled {
    background-color: #ddd;
    border-color: #ddd;
    color: black;
    cursor: not-allowed; // 鼠标变化
    pointer-events: none;
}
.router-view {
    margin-top: 20px;
}
.sectionUl li {
    display: inline;
    padding: 60px;
}
.dateUl li {
    display: inline;
    //margin: 5px;
    padding: 1px;
}
</style>