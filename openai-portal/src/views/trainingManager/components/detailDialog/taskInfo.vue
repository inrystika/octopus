<template>
    <div>
      <el-row>
            <el-col :span="12">
                <div>任务名称:<span>{{ data.name }}</span></div>
            </el-col>
            <el-col :span="12">
                <div>是否分布式:<span>{{ data.isDistributed?'是':'否' }}</span></div>
            </el-col>
        </el-row>
        <el-row>
            <el-col :span="12">
                <el-form ref="ruleForm" :model="ruleForm">
                    <el-form-item prop="subTaskItem">
                        <div style="font-size: 15px">子任务名:
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
                        </div>
                    </el-form-item>
                </el-form>
            </el-col>
        </el-row>

        <div>
            <el-input
                v-if="showInfo"
                v-model="subTaskInfo"
                type="textarea"
                :readonly="true"
                :autosize="true"
            />
        </div>

        <div class="block">
            <el-pagination
                v-if="showInfo"
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
                data: {},
                initInfo: "",
                subTaskOptions: [],
                ruleForm: {
                  subTaskItem: ""
                },
                subTaskInfo: "",
                pageIndex: 1,
                pageSize: 10,
                total: 0,
                showInfo: false
            }
        },
        created() {
            this.data = JSON.parse(JSON.stringify(this.row))
            for (let i = 0; i < this.row.config.length; i++) {
                for (let j = 0; j < this.row.config[i].taskNumber; j++) {
                    this.subTaskOptions.push({
                        label: this.row.config[i].replicaStates[j].key,
                        taskIndex: i + 1,
                        replicaIndex: j + 1
                    })
                }
            }
        },
        methods: {
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
                        this.showInfo = response.payload.jobEvents.length
                        this.total = response.payload.totalSize
                        let infoMessage = ""
                        response.payload.jobEvents.forEach(function(element) {
                            const title = element.reason
                            const message = element.message
                            infoMessage += "\n" + "[" + title + "]"
                            infoMessage += "\n" + "[" + message + "]" + "\n"
                        })
                        this.subTaskInfo = infoMessage
                    } else {
                        this.$message({
                            message: "暂无相关运行信息",
                            type: 'warning'
                        });
                    }
                }).catch(err => {
                    console.log("err:", err)
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
    .el-col {
        margin: 10px 0 20px 0;
        font-size: 15px;
        font-weight: 800;

        span {
            font-weight: 400;
            margin-left: 20px
        }
    }

    .select {
        margin-left: 5px;
    }

    .block {
        float: right;
        margin: 20px;
    }
</style>