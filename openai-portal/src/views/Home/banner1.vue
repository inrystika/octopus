<template>
  <el-carousel :interval="3000" arrow="always" height="400px">
    <el-carousel-item v-for="(item,index) in banner" :key="index">
      <img :src="item.src" class="image">
    </el-carousel-item>
  </el-carousel>
</template>
<script>
  export default {
    data() {
      return {
        banner: [
          { src: require('@/assets/banner1.jpg'), desc: 'banner1' },
          { src: require('@/assets/banner2.jpg'), desc: 'banner2' },
          { src: require('@/assets/banner3.jpg'), desc: 'banner3' }
        ]
      }
    }
  }
</script>
<style>
  .el-carousel__item h3 {
    color: #475669;
    font-size: 14px;
    opacity: 0.75;
    line-height: 200px;
    margin: 0;
  }

  .el-carousel__item:nth-child(2n) {
    background-color: #99a9bf;
  }

  .el-carousel__item:nth-child(2n+1) {
    background-color: #d3dce6;
  }
  .image{width: 100%;height:100%;}
</style>