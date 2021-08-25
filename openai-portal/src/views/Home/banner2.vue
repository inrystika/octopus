<template>
    <el-carousel :interval="4000" type="card" height="400px">
        <el-carousel-item v-for="(item,index) in model" :key="index">
            <img :src="item.src" class="image">
          </el-carousel-item>
    </el-carousel>
  </template>
  <script>
    export default {
      data() {
        return {
          model: [
            { src: require('@/assets/model1.jpeg'), desc: 'model1' },
            { src: require('@/assets/model2.jpeg'), desc: 'model2' },
            { src: require('@/assets/model3.jpeg'), desc: 'model3' },
            { src: require('@/assets/model4.jpeg'), desc: 'model3' },
            { src: require('@/assets/model5.jpeg'), desc: 'model3' },
            { src: require('@/assets/model6.jpeg'), desc: 'model3' },
            { src: require('@/assets/model7.jpeg'), desc: 'model3' },
            { src: require('@/assets/model8.png'), desc: 'model3' },
            { src: require('@/assets/model9.jpeg'), desc: 'model3' }
          
          ]
        }
      }
    }
  </script>
  <style>
  .el-carousel__item h3 {
    color: #475669;
    font-size: 18px;
    opacity: 0.75;
    line-height: 300px;
    margin: 0;
  }
  
  .el-carousel__item:nth-child(2n) {
    background-color: #99a9bf;
  }
  
  .el-carousel__item:nth-child(2n+1) {
    background-color: #d3dce6;
  }
  </style>