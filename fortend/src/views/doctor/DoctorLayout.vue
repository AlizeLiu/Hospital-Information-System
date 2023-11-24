<template>
  <div>
    <!-- <div class="indexImage">
    <img src="@/assets/hospital.jpeg" class="layoutImage"/>
    <span>今天预约挂号总人数：{{orderPeople}}</span>
  </div> -->
    <div class="indexPeople" style="margin-left: 550px">
      <div class="userImage">
        <i class="el-icon-user" style="font-size: 132px"></i>
      </div>

      <div class="userFont">
        <div class="spanCure">
          <span>就诊概况</span>
        </div>
        <div class="spanPeople">
          <span>今天预约挂号总人数：{{ orderPeople }}</span>
        </div>
      </div>
    </div>
<!--    <div class="indexPeople" style="margin-left: 10px">-->
<!--      <div class="userImage">-->
<!--        <i class="el-icon-office-building" style="font-size: 132px"></i>-->
<!--      </div>-->

<!--      <div class="userFont">-->
<!--        <div class="spanCure">-->
<!--          <span>住院概况</span>-->
<!--        </div>-->
<!--        <div class="spanPeople">-->
<!--          <span>今天住院总人数：{{ bedPeople }}</span>-->
<!--        </div>-->
<!--      </div>-->
<!--    </div>-->

    <el-row>
      <el-col :span="24">
        <img src="@/assets/16.png" style="width: 641px;margin-left: 490px;">
      </el-col>
    </el-row>
  </div>
</template>
<script>
import request from "@/utils/request.js";
import jwtDecode from "jwt-decode";
import { getToken } from "@/utils/storage.js";
export default {
    name: "DoctorLayout",
    data() {
        return {
            userId: 1,
            orderPeople: 1,
        };
    },
    methods: {
        //token解码
        tokenDecode(token) {
            if (token !== null) return jwtDecode(token);
        },
        requestPeople() {
            request
                .get("order/orderPeopleByDid", {
                    params: {
                        dId: this.userId,
                    },
                })
                .then((res) => {
                    if (res.data.status !== 200)
                        return ElMessage.error("数据请求失败");
                    this.orderPeople = res.data.data;
                });
        },
    },
    created() {
        this.userId = this.tokenDecode(getToken()).dId;
        this.requestPeople();
    },
};
</script>
<style lang="scss" scoped>
.userFont {
  height: 150px;
  width: 250px;
  float: right;
  color: white;
  .spanCure {
    font-size: 15px;
    margin-top: 60px;
    margin-bottom: 15px;
  }
  .spanPeople {
    font-size: 18px;
  }
}

.userImage {
  height: 150px;
  width: 150px;
  font-size: 130px;
  color: white;
  position: relative;
  left: 40px;
  top: 10px;
  float: left;
}
.indexPeople {
  height: 200px;
  width: 440px;
  background: #58b9ae;
  float: left;
  margin: 30px;
}
</style>