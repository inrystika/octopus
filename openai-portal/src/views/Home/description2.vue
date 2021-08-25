<template>
  <div class="wrapper">
    <div>
      <h1 class="description">鹏城云脑支持的AI引擎</h1>
      <div class="title">D910是由于计算速度快，能迅速反复迭代，研发出更强的AI模型。所以被广泛应用于ModelArts的数据标注，涵盖了自动驾驶、图像识别、声音识别、文本识别等AI智能领域。</div>
      <div class="bottomImage">
        <div class="svgWrapper">
          <div class="svg">
          </div>
        </div>
        <div class="svgWrapper">
          <div class="svg svg2">
          </div>
        </div>
        <div class="svgWrapper">
          <div class="svg svg3">
          </div>
        </div>
      </div>
      <div class="bottomImage">
        <div class="svgWrapper">
          <div class="svg svg4">
          </div>
        </div>
        <div class="svgWrapper">
          <div class="svg svg5">
          </div>
        </div>
        <div class="svgWrapper">
          <div class="svg svg6">
          </div>
        </div>
      </div>
      <div class="bottomImage">
        <div class="svgWrapper">
          <div class="svg svg7">
          </div>
        </div>
        <div class="svgWrapper">
          <div class="svg svg8">
          </div>
        </div>
        <div class="svgWrapper">
          <div class="svg svg9">
          </div>
        </div>

      </div>
    </div>
    <!-- <div class="copyRight">© 启智章鱼</div> -->
  </div>

</template>

<script>
  export default {
    data() {
      return {
      };
    }
  }
</script>
<style lang="scss" scoped>
  .wrapper{margin: 0 auto;width: 60%;}
  .description {
    font-family: PingFangSC-Medium;
    font-size: 24px;
    color: #333333;
    letter-spacing: 0px;
    text-align: center;
  }

  .bottomImage {
    width: 90%;
    display: flex;
    align-items: center;
    margin-left: 5%;
  }



  /* .copyRight {
    text-align: center;
    color: rgba(0, 0, 0, 0.65);
    font-size: 14px;
    margin: 10px 0 20px 0;
  } */

  .copyRight:after {
    display: block;
    content: "";
    width: 100px;
    height: 40px;
  }

  .svgWrapper {
    overflow: hidden;
    border: 1px solid #e8e8e8;
    width: 555px;
    height: 70px;
    margin: 0px 80px 25px 0px;
    padding-top: 10px;
  }

  .svg {
    width: 140px;
    height: 50px;
    background-repeat: no-repeat;
    margin: 0 auto;
    background-image: url("../../assets/ai.svg");
  }

  .svg2 {
    width: 120px;
    background-position: -140px 8px;
  }

  .svg3 {
    width: 160px;
    background-position: -255px 0px;
  }

  .svg4 {
    width: 115px;
    background-position: -410px 0px;
  }

  .svg5 {
    width: 115px;
    background-position: -520px 0px;
  }

  .svg6 {
    width: 115px;
    background-position: -635px 0px;
  }

  .svg7 {
    width: 80px;
    background-position: -750px 0px;
  }

  .svg8 {
    width: 80px;
    background-position: -825px 0px;
  }

  .svg9 {
    width: 60px;
    background-position: -900px 0px;
  }

  .title {
    text-align: center;
    font-family: PingFangSC-Regular;
    font-size: 14px;
    color: #666666;
    letter-spacing: 0px;
    margin-bottom: 18px;
  }
</style>