<template>
  <div class="dashboard-container">
    <el-card v-loading="loading" :body-style="{padding: '0px'}">
      <div class='top' v-if="show">
        <div class='topTitle'>
          <span>任务概览</span>
        </div >

        <el-row :gutter="20" class='rowPadding' type="flex">
          <el-col :span="8">
            <el-row :gutter="20">
              <el-col :span="12">
                <el-card class='topCard'>
                  <span class='topCardTitle'>运行中的训练任务</span>
                  <br/>
                  <span class='topCardRunningNum'>
                    {{this.count.running}}
                  </span>
                </el-card>
              </el-col>
              <el-col :span="12">
                <el-card class='topCard'>
                  <span class='topCardTitle'>等待的训练任务</span>
                  <br/>
                  <span class='topCardNum'>
                    {{this.count.preparing + this.count.pending}}
                  </span>
                </el-card>
              </el-col>
            </el-row>
            <el-row :gutter="20">
              <el-col :span="12">
                <el-card class='topCard'>
                  <span class='topCardTitle'>已成功的训练任务</span>
                  <br/>
                  <span class='topCardNum'>
                    {{this.count.succeeded}}
                  </span>
                </el-card>
              </el-col>
              <el-col :span="12">
                <el-card class='topCard'>
                  <span class='topCardTitle'>已终止的训练任务</span>
                  <br/>
                  <span class='topCardNum'>
                    {{this.count.failed + this.count.stopped}}
                  </span>
                </el-card>
              </el-col>
            </el-row>
          </el-col>

          <el-col :span="8">
            <div class='topCircle'> 
                  <el-progress type="circle" 
                  :show-text="false"
                  :percentage="100" 
                  :color="customColor" 
                  :width="200"></el-progress>
                  <div class='topCircleContent'>
                    <div class='topCircleContentTitle'>剩余机时</div>
                    <span class='topCircleContentNum'>{{ this.billAmount }}</span>
                    <!-- <div class='topCircleContentText'>机时/h</div> -->
                  </div>
            </div>
          </el-col>
          <el-col :span="8">
            <div class="topHour">
              <!-- <el-button type="primary" size="small">机时充值</el-button> -->
              <el-button class='topHourButton' type="primary" @click="getConsumption" size="small">消费记录</el-button>
              <el-button class='topHourButton' type="primary" @click="getRecharge" size="small">充值记录</el-button>
              <br/>
              <div class='topHourInstrucTitle'>
                充值说明: 
                <span class='topHourInstrucText'>
                  充值请向管理员提交申请
                </span>
              </div>
            </div>
          </el-col>
        </el-row>
      </div>

      <div class='main'>
        <el-row :gutter="20" class='rowPadding'>
          <el-col :span="8">
            <span class='mainTitle'>
              模型开发
            </span>
            <div class='mainBlock'>
              <el-button @click="create('notebook')" class='mainButtonBorder' size="small">
                <span class='mainButtonText'>
                  创建Notebook
                </span>
              </el-button>
              <el-button @click="create('algorithm')" class='mainButtonBorder' size="small">
                <span class='mainButtonText'>
                  创建算法
                </span>
              </el-button>
            </div>
          </el-col>
          <el-col :span="8">
            <span class='mainTitle'>
              模型训练
            </span>
            <div class='mainBlock'>
              <el-button @click="create('trainingTask')" class='mainButtonBorder' size="small">
                <span class='mainButtonText'>
                  创建训练任务
                </span>
              </el-button>
              <el-button @click="create('trainingTemplate')" class='mainButtonBorder' size="small">
                <span class='mainButtonText'>
                  创建训练模板
                </span>
              </el-button>
            </div>
            <div class='mainColText'>
                总训练任务：
              <span class="mainNum">
                {{this.totalTrainingTaskNum}}
              </span>
              个
            </div>
            <div class='mainBlockText'>
              任务模板：
              <span class="mainNum">
                {{this.trainingTemplateNum}}
              </span>
              个
            </div>
          </el-col>
          <el-col :span="8">
            <span class='mainTitle'>
              模型管理
            </span>
            <div class='mainBlock'>
              <div class='mainBlockText'>
                我的模型：
                <span class="mainNum">
                  {{this.myModelNum}}
                </span>
                个
              </div>
              <div class='mainColText'>
                公共模型：
                <span class="mainNum">
                  {{this.pubModelNum}}
                </span>
                个
              </div>
              <div class='mainColText'>
                预置模型：
                <span class="mainNum">
                  {{this.preModelNum}}
                </span>
                个
              </div>
            </div>
          </el-col>
        </el-row>
      </div>

      <el-divider></el-divider>
      
      <div class='mainBlock'>
        <el-row :gutter="20" class='rowPadding'>
          <el-col :span="8">
            <el-row>
              <el-col :span="12">
                <span class='mainTitle'>
                  算法
                </span>
                <div class='mainBlock'>
                  <div class='mainBlockText'>
                    我的算法：
                    <span class="mainNum">
                      {{this.myAlgorithmNum}}
                    </span>
                    个
                  </div>
                  <div class='mainColText'>
                    公共算法：
                    <span class="mainNum">
                      {{this.pubAlgorithmNum}}
                    </span>
                    个
                  </div>
                  <div class='mainColText'>
                    预置算法：
                    <span class="mainNum">
                      {{this.preAlgorithmNum}}
                    </span>
                    个
                  </div>
                </div>
                <el-button @click="create('algorithm')" class='mainButtonBorder' size="small">
                  <span class='mainButtonText'>
                    创建算法
                  </span>
                </el-button>
              </el-col>
              <el-col :span="12">
                <el-divider direction="vertical"></el-divider>
              </el-col>
            </el-row>
          </el-col>

          <el-col :span="8">
            <el-row>
              <el-col :span="12">
                <span class='mainTitle'>
                  数据集
                </span>
                <div class='mainBlock'>
                  <div class='mainBlockText'>
                    我的数据集：
                    <span class="mainNum">
                      {{this.myDatasetNum}}
                    </span>
                    个
                  </div>
                  <div class='mainColText'>
                    公共数据集：
                    <span class="mainNum">
                      {{this.pubDatasetNum}}
                    </span>
                    个
                  </div>
                  <div class='mainColText'>
                    预置数据集：
                    <span class="mainNum">
                      {{this.preDatasetNum}}
                    </span>
                    个
                  </div>
                </div>
                <el-button @click="create('dataset')" class='mainButtonBorder' size="small">
                  <span class='mainButtonText'>
                    创建数据集
                  </span>
                </el-button>
              </el-col>
              <el-col :span="12">
                  <el-divider direction="vertical"></el-divider>
              </el-col>
            </el-row>
          </el-col>

          <el-col :span="8">
            <span class='mainTitle'>
              镜像
            </span>
            <div class='mainBlock'>
              <div class='mainBlockText'>
                我的镜像：
                <span class="mainNum">
                  {{this.myImageNum}}
                </span>
                个
              </div>
              <div class='mainColText'>
                公共镜像：
                <span class="mainNum">
                  {{this.pubImageNum}}
                </span>
                个
              </div>
              <div class='mainColText'>
                预置镜像：
                <span class="mainNum">
                  {{this.preImageNum}}
                </span>
                个
              </div>
            </div>
            <el-button @click="create('image')" class='mainButtonBorder' size="small">
              <span class='mainButtonText'>
                创建镜像
              </span>
            </el-button>
          </el-col>
        </el-row>
      </div>
    </el-card>

    <record v-if="recordVisible" :groupName="groupName" :recordType="recordType" @close="close">           
    </record>
  </div>
</template>

<script>
import { mapGetters } from "vuex";
import jobHourStat from "./components/jobHourStat";
import { getList, getTemplate } from '@/api/trainingManager'
import { getMyModel, getPreModel, getPublicModel } from '@/api/modelManager'
import { getMyAlgorithmList, getPublicAlgorithmList, getPresetAlgorithmList } from '@/api/modelDev'
import { getMyDatasetList, getPublicDatasetList, getPresetDatasetList } from '@/api/datasetManager'
import { getMyImage, getPublicImage, getPreImage } from '@/api/imageManager'
import { getUserHour, getGroupHour } from "@/api/generalView";
import record from './components/record.vue'
import { getErrorMsg } from '@/error/index'

export default {
  name: "Dashboard",
  components: {
    jobHourStat,
    record
  },
  data() {
    return {
      customColor:[{color: '#666699', percentage: 100}],
      count: {},
      show: false,
      recordVisible: false,
      billAmount: undefined,
      groupName: undefined,
      recordType: undefined,
      loading:true,
      totalTrainingTaskNum: undefined,
      trainingTemplateNum: undefined,
      myModelNum: undefined,
      pubModelNum: undefined,
      preModelNum: undefined,
      myAlgorithmNum: undefined,
      pubAlgorithmNum: undefined,
      preAlgorithmNum: undefined,
      myDatasetNum: undefined,
      pubDatasetNum: undefined,
      preDatasetNum: undefined,
      myImageNum: undefined,
      pubImageNum: undefined,
      preImageNum: undefined
    };
  },
  created() {
    this.getTrainingTask();
    this.getHour()
    this.getAllLit()
  },
  methods: {
    getErrorMsg(code) {
      return getErrorMsg(code)
    },
    async getTrainingTask() {
      const statusList = {
        preparing:'preparing',
        pending:'pending',
        running:'running',
        failed:'failed',
        succeeded:'succeeded',
        stopped:'stopped'
      }
      for (let status in statusList) {
        const param = {
          pageIndex: 1,
          pageSize: 20,
          status: status
        }
        await getList(param).then(response => {
          if (response.success) {
            this.count[status] = response.data.totalSize;
          } else {
              this.$message({
                message: this.getErrorMsg(response.error.subcode),
                type: 'warning'
              });          }
        }).catch(err => {
          this.$message({
            message: this.getErrorMsg(response.error.subcode),
            type: 'warning'
          }); 
        });
      }
      this.show = true
      this.loading = false
    },
    getHour(){
      this.groupName = this.workspaceId
      if (this.workspaceId === "default-workspace") {
        getUserHour().then(response => {
          if(response.success) {
            this.billAmount = response.data.billingUser.amount
          } else {
            this.$message({
              message: this.getErrorMsg(response.error.subcode),
              type: 'warning'
            }); 
          }
        })
      } else {
        getGroupHour().then(response => {
          if(response.success) {
            this.billAmount = response.data.billingSpace.amount
          } else {
            this.$message({
              message: this.getErrorMsg(response.error.subcode),
              type: 'warning'
            }); 
          }
        })
      }
    },
    getAllLit(){
      const param = {
          pageIndex: 1,
          pageSize: 20,
        }
      getList(param).then(response => {
        if(response.success) {
            this.totalTrainingTaskNum = response.data.totalSize
          } else {
            this.$message({
              message: this.getErrorMsg(response.error.subcode),
              type: 'warning'
            }); 
          }
      })
      getTemplate(param).then(response => {
        if(response.success) {
            this.trainingTemplateNum = response.data.totalSize
          } else {
            this.$message({
              message: this.getErrorMsg(response.error.subcode),
              type: 'warning'
            }); 
          }
      })
      getMyModel(param).then(response => {
        if(response.success) {
            this.myModelNum = response.data.totalSize
          } else {
            this.$message({
              message: this.getErrorMsg(response.error.subcode),
              type: 'warning'
            }); 
          }
      })
      getPreModel(param).then(response => {
        if(response.success) {
            this.preModelNum = response.data.totalSize
          } else {
            this.$message({
              message: this.getErrorMsg(response.error.subcode),
              type: 'warning'
            }); 
          }
      })
      getPublicModel(param).then(response => {
        if(response.success) {
            this.pubModelNum = response.data.totalSize
          } else {
            this.$message({
              message: this.getErrorMsg(response.error.subcode),
              type: 'warning'
            }); 
          }
      })
      getMyAlgorithmList(param).then(response => {
        if(response.success) {
            this.myAlgorithmNum = response.data.totalSize
          } else {
            this.$message({
              message: this.getErrorMsg(response.error.subcode),
              type: 'warning'
            }); 
          }
      })
      getPublicAlgorithmList(param).then(response => {
        if(response.success) {
            this.pubAlgorithmNum = response.data.totalSize
          } else {
            this.$message({
              message: this.getErrorMsg(response.error.subcode),
              type: 'warning'
            }); 
          }
      })
      getPresetAlgorithmList(param).then(response => {
        if(response.success) {
            this.preAlgorithmNum = response.data.totalSize
          } else {
            this.$message({
              message: this.getErrorMsg(response.error.subcode),
              type: 'warning'
            }); 
          }
      })
      getMyDatasetList(param).then(response => {
        if(response.success) {
            this.myDatasetNum = response.data.totalSize
          } else {
            this.$message({
              message: this.getErrorMsg(response.error.subcode),
              type: 'warning'
            }); 
          }
      })
      getPublicDatasetList(param).then(response => {
        if(response.success) {
            this.pubDatasetNum = response.data.totalSize
          } else {
            this.$message({
              message: this.getErrorMsg(response.error.subcode),
              type: 'warning'
            }); 
          }
      })
      getPresetDatasetList(param).then(response => {
        if(response.success) {
            this.preDatasetNum = response.data.totalSize
          } else {
            this.$message({
              message: this.getErrorMsg(response.error.subcode),
              type: 'warning'
            }); 
          }
      })
      getMyImage(param).then(response => {
        if(response.success) {
            this.myImageNum = response.data.totalSize
          } else {
            this.$message({
              message: this.getErrorMsg(response.error.subcode),
              type: 'warning'
            }); 
          }
      })
      getPublicImage(param).then(response => {
        if(response.success) {
            this.pubImageNum = response.data.totalSize
          } else {
            this.$message({
              message: this.getErrorMsg(response.error.subcode),
              type: 'warning'
            }); 
          }
      })
      getPreImage(param).then(response => {
        if(response.success) {
            this.preImageNum = response.data.totalSize
          } else {
            this.$message({
              message: this.getErrorMsg(response.error.subcode),
              type: 'warning'
            }); 
          }
      })
    },
    getConsumption(){
      this.recordVisible = true;
      this.recordType = 1
    },
    getRecharge(){
      this.recordVisible = true;
      this.recordType = 2
    },
    view() {
      this.recordVisible = true;
    },
    close(val) {
      this.recordVisible = val;
    },
    create(param){
      let data = {}
      data[param] = true
      switch(param) {
        case 'notebook':
          this.$router.push({
            path: '/modelDev/notebook',
            query:{ data: data }
          })
          break
        case 'algorithm':
          this.$router.push({
            path: '/modelDev/algorithmManager',
            query:{ data: data }
          })
          break
        case 'trainingTask':
          this.$router.push({
            name: 'trainingManager',
            params: { data: data }
          })
          break
        case 'trainingTemplate':
          this.$router.push({
            name: 'trainingManager',
            params: { data: data }
          })
          break
        case 'dataset':
          this.$router.push({
            name: 'dataManager',
            params: { data: data }
          })
          break
        case 'image':
          this.$router.push({
            name: 'imageManager',
            params: { data: data }
          })
          break
      }
    },
    
  },
  computed: {
    ...mapGetters(["name","workspaceId"])
  }
};
</script>

<style lang="scss" scoped>
  .dashboard {
    &-container {
      margin: 30px;
    }
    &-text {
      font-size: 30px; 
      line-height: 46px;
    }
  }
  .el-row {
    margin-bottom: 20px;
    &:last-child {
      margin-bottom: 0;
    }
  }
  .el-col {
    border-radius: 4px;
  }

  .top {
    width: auto;
    height: 338.35px;
    background: #2b2b3a;
  }
  .topTitle {
    // margin-top:2%;
    padding:2% 2%;
    font-family: "MicrosoftYaHeiLight ";
    font-weight: normal;
    font-size: 24px;
    line-height: 28.8px;
    color: #669;
  }
  .rowPadding {
    padding:0 2%
  }
  .topCard {
    background: #2b2b3a;
    border-color:#669
  }
  .topCardTitle {
    font-family: "MicrosoftYaHei ";
    font-weight: normal;
    font-size: 14px;
    line-height: 16.8px;
    text-align: center;
    color: #999
  }
  .topCardRunningNum {
    font-family: "MicrosoftYaHei-Bold ";
    font-weight: normal;
    font-size: 28px;
    line-height: 33.6px;
    float: center;
    color: #0cc;
  }
  .topCardNum {
    font-family: "MicrosoftYaHei-Bold ";
    font-weight: normal;
    font-size: 28px;
    line-height: 33.6px;
    float: center;
    color: #ccc;
  }
  .topCircle {
    position: relative;
    margin: 0 auto;
    width: 200px;
  }
  .topCircleContent {
    position: absolute;
    left: 50%;    //起始是在body中，横向距左50%的位置
		top:50%;      //起始是在body中，纵向距上50%的位置，这个点相当于body的中心点，div的左上角的定位
		transform:translate(-50%,-50%);
  }
  .topCircleContentTitle {
    font-family: "MicrosoftYaHei ";
    font-weight: normal;
    font-size: 14px;
    line-height: 16.8px;
    text-align: center;
    color: #999;
    // padding-bottom: 10%;
    // margin-bottom: 10%;
  }
  .topCircleContentNum {
    font-family: "MicrosoftYaHei-Bold ";
    font-weight: normal;
    font-size: 28px;
    line-height: 33.6px;
    text-align: center;
    color: #fff;
  }
  .topCircleContentText {
    font-family: "MicrosoftYaHeiLight ";
    font-weight: normal;
    font-size: 14px;
    line-height: 16.8px;
    text-align: center;
    color: #999;
    padding-top: 15%;
    margin-top: 20%;
  }
  .topHour {
    position: relative;
    margin: 10% auto
  }
  .topHourInstrucTitle {
    font-family: "MicrosoftYaHei ";
    font-size: 14px;
    line-height: 16.8px;
    color: #999999;
    margin-top: 5%;
  }
  .topHourButton {
    background: #2B2B3A;
    border-color: #666699;
  }
  .topHourInstrucText {
    font-family: "MicrosoftYaHei ";
    font-size: 14px;
    line-height: 16.8px;
     color: #fff;
  }

  .el-divider--vertical{
    width:1px;
    height:180px;		//更改竖向分割线长度
    vertical-align:middle;
    float: right;
  }

  .main {
    margin:3% 0;
  }
  .mainTitle {
    font-family: "MicrosoftYaHeiLight ";
    font-weight: normal;
    font-size: 24px;
    line-height: 28.8px;
    text-align: left;
    color: #669
  }
  .mainBlock {
    margin:15px 0
  }
  .mainBlockText {
    color:#606266;
    font-size: 14px;
  }
  .mainButtonText {
    font-family: "MicrosoftYaHei ";
    font-weight: normal;
    color: #03c;
  }
  .mainButtonBorder {
    border-color:#409EFF;
  }
  .mainColText {
    color:#606266;
    font-size: 14px;
    margin:10px 0;
  }
  .mainNum {
    color:#0033CC;
    font-weight:bold;
  }

</style>

