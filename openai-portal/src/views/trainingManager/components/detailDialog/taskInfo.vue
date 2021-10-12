<template>
    <div>
      <el-form ref="ruleForm" :model="ruleForm">
        <el-form-item label="子任务名:" prop="subTaskItem">
          <el-select 
              v-model="ruleForm.subTaskItem"
              value-key="label"
              placeholder="请选择" 
              @change="selectedSubTaskOption" 
          >
              <el-option 
                  v-for="item in subTaskOptions" 
                  :key="item.label" 
                  :label="item.label" 
                  :value="item" 
              />
          </el-select>
        </el-form-item>
      </el-form>
    </div>
</template>

<script>
    import { getTempalteInfo } from '@/api/trainingManager'
    export default {
        name: "TaskInfo",
        props: {
            row: {
                type: Object,
                default: () => { }
            }
        },
        data() {
          return {
            initInfo: "",
            subTaskOptions: [],
            ruleForm: {
              subTaskItem: ""
            }
          }
        },
        created() {
            console.log("123:",this.row)
              for (let i = 0; i < this.row.config.length; i++) {
                  for (let j = 0; j < this.row.config[i].taskNumber; j++) {
                      this.subTaskOptions.push({ label: this.row.config[i].replicaStates[j].key, taskIndex: i + 1, replicaIndex: j + 1})
                  }
              }
              // console.log("subTaskOptions:",this.subTaskOptions)
            // const taskInfoString = this.row.initInfo ? this.row.initInfo.replace(/\n/g, "<br>") : ''
            // const taskInfoData = JSON.parse(taskInfoString)
            // for (const pid in taskInfoData['podEvents']) {
            //     const eventList = taskInfoData['podEvents'][pid]
            //     const roleName = taskInfoData['podRoleName'][pid]
            //     if (roleName == "") {
            //         continue
            //     }
            //     let message = ""
            //     for (const key in eventList) {
            //         const event = eventList[key]
            //         if (event['reason'] == "" && event['message'] == "") {
            //             continue
            //         }
            //         message += "[" + event['reason'] + "]" + "<br>"
            //         message += event['message'] + "<br><br>"
            //     }
            //     for (const key in taskInfoData['extras']) {
            //         const event = taskInfoData['extras'][key]
            //         if (event['reason'] == "" && event['message'] == "") {
            //             continue
            //         }
            //         message += "[" + event['reason'] + "]" + "<br>"
            //         message += event['message'] + "<br><br>"
            //     }
            //     message += "<br>"
            //     tempTaskInfoData[roleName] = message
            // }
            // let obj = {}
            // Object.keys(tempTaskInfoData).sort().forEach(function(key) {
            //     obj[key] = tempTaskInfoData[key];
            // });
            // this.initInfo = obj 
        },
        methods: {
          selectedSubTaskOption() {
              const param = {
                id: this.row.id,
                pageIndex: 1,
                pageSize: 10,
                taskIndex: this.ruleForm.subTaskItem.taskIndex,
                replicaIndex: this.ruleForm.subTaskItem.replicaIndex
              }
              getTempalteInfo(param).then(response => {
                if (response.success) {
                  console.log("payload:",response.payload)
                } else [
                  console.log("failed:",response)
                ]
              })
              console.log("this.subTaskItem:",this.ruleForm.subTaskItem)
          }
        }
    }
</script>

<style lang="scss" scoped>
    .select {
        margin-left: 5px;
    }
</style>