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

        <div>
            <el-row>
                <div v-html="subTaskInfo"></div>
            </el-row>
        </div>

        <div class="block">
            <el-pagination
              :current-page="pageIndex"
              :page-sizes="[10, 20, 50, 80]"
              :page-size="pageSize"
              layout="total, sizes, prev, pager, next, jumper"
              :total="total"
              @size-change="handleSizeChange"
              @current-change="handleCurrentChange"
            />
        </div>
    </div>
</template>

<script>
    import { getTempalteInfo } from '@/api/trainingManager'
    import { getErrorMsg } from '@/error/index'
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
                },
                subTaskInfo: "",
                pageIndex: 1,
                pageSize: 10,
                total: 0
            }
        },
        created() {
              for (let i = 0; i < this.row.config.length; i++) {
                  for (let j = 0; j < this.row.config[i].taskNumber; j++) {
                      this.subTaskOptions.push({ label: this.row.config[i].replicaStates[j].key, taskIndex: i + 1, replicaIndex: j + 1})
                  }
              }                    
        },
        methods: {
            // 错误码
            getErrorMsg(code) {
                return getErrorMsg(code)
            },
            selectedSubTaskOption() {
                const param = {
                    id: this.row.id,
                    pageIndex: this.pageIndex,
                    pageSize: this.pageSize,
                    taskIndex: this.ruleForm.subTaskItem.taskIndex,
                    replicaIndex: this.ruleForm.subTaskItem.replicaIndex
                }
                getTempalteInfo(param).then(response => {
                    if (response.success) {
                        this.total = response.payload.totalSize
                        let infoMessage = ""
                        response.payload.jobEvents.forEach(function (element) {
                            const title = element.reason
                            const message = element.message
                            infoMessage += "[" + title + "]" + "<br>"
                            infoMessage += "[" + message + "]" + "<br><br>"
                        })
                        this.subTaskInfo = infoMessage
                    } else {
                        this.$message({
                            message: this.getErrorMsg(response.error.subcode),
                            type: 'warning'
                        });
                    }
                }).catch(err => {
                    console.log("err:",err)
                    this.$message({
                        message: "未知错误",
                        type: 'warning'
                    });
                });
            },
            handleSizeChange(val) {
                this.pageSize = val
                this.selectedSubTaskOption()
            },
            handleCurrentChange(val) {
                this.pageIndex = val
                this.selectedSubTaskOption()
            }
        }
    }
</script>

<style lang="scss" scoped>
    .select {
        margin-left: 5px;
    }

    .block {
        float: right;
        margin: 20px;
    }
</style>