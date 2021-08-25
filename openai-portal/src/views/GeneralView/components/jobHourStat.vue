<template>
  <div>
    <el-row :gutter="20">
      <div>
        <el-col :span="16">
          <el-row :gutter="20">
            <el-col :span="12">
              <el-card>
                <div style='float:left;margin-right:5%'>
                  <el-progress type="circle" :percentage="0"></el-progress>
                </div>
                <div style='float:left;margin-top:10%'>
                  <span class="upperNumber">12</span>
                  <br />
                  <span>正在等待的任务</span>
                </div>
              </el-card>
            </el-col>
            <el-col :span="12">
              <el-card>
                <div style='float:left;margin-right:5%'>
                  <el-progress type="circle" :percentage="50"></el-progress>
                </div>
                <div style='float:left;margin-top:10%'>
                  <span class="upperNumber">12</span>
                  <br />
                  <span>正在运行的任务</span>
                </div>
              </el-card>
            </el-col>
          </el-row>
          <el-row :gutter="20">
            <el-col :span="12">
              <el-card>
                <div style='float:left;margin-right:5%'>
                  <el-progress type="circle" :percentage="50" status="warning"></el-progress>
                </div>
                <div style='float:left;margin-top:10%'>
                  <span class="upperNumber">12</span>
                  <br />
                  <span>已终止的任务</span>
                </div>
              </el-card>
            </el-col>
            <el-col :span="12">
              <el-card>
                <div style='float:left;margin-right:5%'>
                  <el-progress type="circle" :percentage="50" status="success"></el-progress>
                </div>
                <div style='float:left;margin-top:10%'>
                  <span class="upperNumber">12</span>
                  <br />
                  <span>已成功的任务</span>
                </div>
              </el-card>
            </el-col>
          </el-row>
        </el-col>
      </div>
      <el-col :span="8">
        <el-card style="heitht:100%">
          机时
        </el-card>
      </el-col>
    </el-row>


    <el-row type="flex" justify="center">
      <el-col :span="24">
        <div style="font-size:18px; color:#333">机时</div>
      </el-col>
    </el-row>
    <el-row justify="space-between" type="flex">
      <el-col :span="6">
        <div class="grid-content-hour">
          <el-row class="subTitle">剩余(小时)</el-row>
          <el-row>
            <span class="upperNumber">{{ this.billAmount }}</span>
          </el-row>
        </div>
      </el-col>
      <el-col :span="6">
        <div class="grid-content-hour">
          <el-row>
            <el-button type="primary" @click="getConsumption">消费记录</el-button>
          </el-row>
        </div>
      </el-col>
      <el-col :span="6">
        <div class="grid-content-hour">
          <el-row>
            <el-button type="primary" @click="getRecharge">充值记录</el-button>
          </el-row>
        </div>
      </el-col>
    </el-row>
    <!-- 消费记录对话框 -->
    <record v-if="recordVisible" :groupName="groupName" :recordType="recordType" @close="close">
    </record>
  </div>
</template>

<script>
  import { mapGetters } from 'vuex'
  import record from "./record.vue";
  import { getUserHour, getGroupHour } from "@/api/generalView";
  export default {
    name: "jobHourStat",
    components: {
      record
    },
    data() {
      return {
        recordVisible: false,
        billAmount: undefined,
        groupName: undefined,
        recordType: undefined
      };
    },
    computed: {
      ...mapGetters([
        'workspaceId'
      ])
    },
    created() {
      this.getHour()
    },
    methods: {
      getHour() {
        this.groupName = this.workspaceId
        if (this.workspaceId === "default-workspace") {
          getUserHour().then(response => {
            if (response.success) {
              this.billAmount = response.data.billingUser.amount
            } else {
              this.$message.error("暂时无法获取剩余任务机时");
            }
          })
        } else {
          getGroupHour().then(response => {
            if (response.success) {
              this.billAmount = response.data.billingSpace.amount
            } else {
              this.$message.error("暂时无法获取剩余任务机时");
            }
          })
        }
      },
      getConsumption() {
        this.recordVisible = true;
        this.recordType = 1
      },
      getRecharge() {
        this.recordVisible = true;
        this.recordType = 2
      },
      view() {
        this.recordVisible = true;
      },
      close(val) {
        this.recordVisible = val;
      },
    }
  };
</script>

<style lang="scss" scoped>
  .grid-content-hour {
    border-radius: 4px;
    // min-height: 36px;
    height: 120px;
    background-color: #DDDAE5 !important;
    border-radius: 8px;
    color: #4a4a4a;
    font-size: 18px;
  }

  .subTitle {
    margin-top: 15px;
    color: rgba(0, 0, 0, 0.45);
    font-size: 14px;
    height: 30px;
  }

  .upperNumber {
    font-size: 24px;
    color: rgba(0, 0, 0, 0.85);
    font-family: -apple-system, BlinkMacSystemFont, Segoe UI, PingFang SC,
      Hiragino Sans GB, Microsoft YaHei, Helvetica Neue, Helvetica, Arial,
      sans-serif, Apple Color Emoji, Segoe UI Emoji, Segoe UI Symbol;
  }

  .lowerNumber {
    font-size: 16px;
    color: rgba(0, 0, 0, 0.85);
    font-family: -apple-system, BlinkMacSystemFont, Segoe UI, PingFang SC,
      Hiragino Sans GB, Microsoft YaHei, Helvetica Neue, Helvetica, Arial,
      sans-serif, Apple Color Emoji, Segoe UI Emoji, Segoe UI Symbol;
  }
</style>