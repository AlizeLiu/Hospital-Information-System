<template>
  <div>
    <!-- <div class="indexImage">
    <img src="@/assets/hospital.jpeg" class="layoutImage"/>
    <span>今天预约挂号总人数：{{orderPeople}}</span>
  </div> -->

    <el-row>
      <div class="indexPeople" style="margin-left: 350px">
        <div class="userImage">
          <el-icon style="font-size: 132px"><User /></el-icon>
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
      <div class="indexPeople">
        <div class="userImage">
          <el-icon style="font-size: 132px"><OfficeBuilding /></el-icon>
        </div>
        <div class="userFont">
          <div class="spanCure">
            <span>住院概况</span>
          </div>
          <div class="spanPeople">
            <span>今天住院总人数：{{ bedPeople }}</span>
          </div>
        </div>
      </div>
    </el-row>
    <el-row>
        <img src="@/assets/16.png" style="width: 550px;margin: 0 auto;">
    </el-row>
  </div>
</template>
<script setup>
import { ref, reactive, onMounted, } from 'vue'
import { getPeople, getBed } from '@/api/admin/adminLayout.ts'
import { ElMessage } from 'element-plus'

const orderPeople = ref(1);
const bedPeople = ref(1);
// 今天预约挂号总人数：
const requestPeople = () => {
  getPeople().then((res) => {
    console.log(res);
    if (res.status !== 200)
      return ElMessage.error('数据请求失败');
    orderPeople.value = res.data.data;
  })
    .catch((err) => {
      console.error(err);
    });
}
const requestBed = () => {
  getBed().then((res) => {
    if (res.status !== 200)
      return ElMessage.error('数据请求失败');
    bedPeople.value = res.data.data;
  })
    .catch((err) => {
      console.error(err);
    });
}
onMounted(() => {
  requestPeople();
  requestBed();
});
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