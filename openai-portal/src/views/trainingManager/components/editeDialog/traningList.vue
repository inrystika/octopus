<template>
    <div>
        <div class="index">
            <el-button type="primary" @click="add" class="add">添加</el-button>
            <el-table :data="tableData" style="width: 100%" :header-cell-style="{'text-align':'left','color':'black'}"
                :cell-style="{'text-align':'left'}">
                <el-table-column prop="name" label="任务名称" align="center">
                </el-table-column>
                <el-table-column label="是否是主任务" align="center">
                    <template slot-scope="scope">
                        <span style="margin-left: 10px">{{ scope.row.isMainRole?'是':'否' }}</span>
                    </template>
                </el-table-column>
                <el-table-column prop="taskNumber" label="副本个数" align="center">
                </el-table-column>
                <el-table-column prop="minSucceededTaskCount" label="最小副本成功个数" align="center">
                </el-table-column>
                <el-table-column prop="minFailedTaskCount" label="最小副本失败数" align="center">
                </el-table-column>
                <el-table-column prop="resourceSpecId" label="资源规格" align="center">
                    <template slot-scope="scope">
                        <span>{{showResource(scope.row)}}</span>
                    </template>
                </el-table-column>
                <el-table-column label="运行命令" align="center">
                    <template slot-scope="scope">
                        <span>{{command(scope.row)}}</span>
                    </template>
                </el-table-column>
                <el-table-column label="操作">
                    <template slot-scope="scope">
                        <el-button type="text" @click="handleEdit(scope.row)">编辑</el-button>
                        <el-button type="text" @click.native.prevent="handleDelete(scope.$index, tableData)">删除
                        </el-button>
                    </template>
                </el-table-column>
            </el-table>
        </div>
        <!-- 分布式任务对话框 -->
        <distributedTask v-if="FormVisible" @cancel="cancel" @confirm="confirm" @close="close" :row="row" :flag="flag"
            @subTasks="getsubTasksList">
        </distributedTask>
    </div>
</template>

<script>
    import distributedTask from './distributedTask.vue'
    import { getResourceList } from "@/api/trainingManager"
    import { getErrorMsg } from '@/error/index'
    export default {
        name: "traningList",
        components: {
            distributedTask
        },
        props: {
            Table: {
                type: Array,
                default: () => []
            },
            resource: {
                type: Array,
                default: () => []
            }
        },
        data() {
            return {
                tableData: [],
                FormVisible: false,
                row: {},
                flag: true,
                resourceOptions: []

            }

        },
        watch: {
            tableData() {
                this.$emit('tableData', this.tableData)
            }
        },
        created() {
            this.tableData = this.Table
            this.getResourceList()


        },
        methods: {
            // 错误码
            getErrorMsg(code) {
                return getErrorMsg(code)
            },
            add() {
                this.FormVisible = true,
                    this.flag = true,
                    this.row = { parameters: [] }
            },
            handleEdit(row) {
                this.FormVisible = true
                this.row = row
                this.flag = false


            },
            handleDelete(index, rows) {
                rows.splice(index, 1);
            },
            cancel(val) {
                this.FormVisible = val

            },
            confirm(val) {
                this.FormVisible = val

            },
            close(val) {
                this.FormVisible = val

            },
            // 监听子组件表格变化
            getsubTasksList(val) {
                val.taskNumber = parseInt(val.taskNumber)
                val.minFailedTaskCount = parseInt(val.minFailedTaskCount)
                val.minSucceededTaskCount = parseInt(val.minSucceededTaskCount)
                // flag为true新增
                // flag为false编辑

                if (this.flag) { this.tableData.push(val); }

            },
            showResource(row) {
                let name = ''
                this.resourceOptions.forEach(item => {
                    if (item.id === row.resourceSpecId) {
                        name = item.name
                    }
                })
                return name
            },
            // 获取资源规格      
            getResourceList() {
                getResourceList().then(response => {
                    if (response.success) {
                        response.data.mapResourceSpecIdList.train.resourceSpecs.forEach(
                            item => {
                                this.resourceOptions.push({ name: item.name + ' ' + item.price + '机时/h', id: item.id })
                            }
                        )

                    }
                    else {
                        this.$message({
                            message: this.getErrorMsg(response.error.subcode),
                            type: 'warning'
                        });
                    }
                })
            },
            command: function (data) {
                let command = data.command
                if (data.parameters != null && data.parameters.length != 0) {
                    data.parameters.forEach(
                        item => {
                            if (item.key != '' || item.value != '') {
                                command += " " + '--' + item.key + '=' + item.value
                            }
                        }
                    )
                }
                return command
            },
        }
    }
</script>

<style lang="scss" scoped>
    .block {
        float: right;
        margin: 20px;
    }

    .add {
        float: right;
        margin: 20px;
    }
</style>