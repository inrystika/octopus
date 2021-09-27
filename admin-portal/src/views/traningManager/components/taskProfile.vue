<template>
    <div>
        <div>
            <el-row>
                <el-col :span="12">
                    <div>任务名称:<span>{{ data.name }}</span></div>
                </el-col>
                <el-col :span="12">
                    <div>描述: <span>{{ data.desc }}</span></div>
                </el-col>
            </el-row>
            <el-row>
                <el-col :span="12">
                    <div>选用算法:<span>{{ data.algorithmName + ":" + data.algorithmVersion }}</span></div>
                </el-col>
                <el-col :span="12">
                    <div>镜像选择:<span>{{ data.imageName + ":" + data.imageVersion }}</span></div>
                </el-col>
            </el-row>
            <el-row>
                <el-col :span="12">
                    <div>选用数据集:<span>{{ data.dataSetName + ":" + data.dataSetVersion }}</span></div>
                </el-col>
                <el-col :span="12">
                    <div>是否分布式:<span>{{ data.isDistributed?'是':'否' }}</span></div>
                </el-col>
            </el-row>
            <el-row>
                <el-col v-if="!show" :span="12">
                    <div>资源规格:<span>{{ data.config[0].resourceSpecName + '' + data.config[0].resourceSpecPrice + '机时/h' }}</span>
                    </div>
                </el-col>
                <el-col v-if="!show" :span="12">
                    <div>运行命令:<span>{{ command(data.config[0]) }}</span></div>
                </el-col>
            </el-row>
            <el-row>
                <el-col :span="12">
                    <div>任务状态:<span style="margin-left: 10px">{{ statusText[data.status][1] }}</span></div>
                </el-col>
                <el-col :span="12">
                    <div></div>
                </el-col>
            </el-row>
        </div>
        <div v-if="show">
            <el-divider />
            <div class="taskList">分布式任务列表</div>
            <div>
                <el-table
                    :data="tableData"
                    style="width: 100%"
                    row-key="name"
                    :tree-props="{children: 'replicaStates', hasChildren: 'hasChildren'}"
                    default-expand-all
                    height="700"
                >
                    <el-table-column prop="name" label="任务名称" width="200px" :show-overflow-tooltip="true" />
                    <el-table-column label="是否主任务">
                        <template slot-scope="scope">
                            <span v-if="!scope.row.isChildren">{{ scope.row.isMainRole?'是':'否' }}</span>
                        </template>
                    </el-table-column>
                    <el-table-column prop="taskNumber" label="副本个数" />
                    <el-table-column prop="minSucceededTaskCount" label="最小副本成功数" />
                    <el-table-column prop="minFailedTaskCount" label="最小副本失败数" />
                    <el-table-column label="资源规格">
                        <template slot-scope="scope">
                            <span v-if="!scope.row.isChildren">
                              {{ scope.row.resourceSpecName + '' + scope.row.resourceSpecPrice + '机时/h' }}
                            </span>
                        </template>
                    </el-table-column>
                    <el-table-column label="运行命令" width="350px" align="center">
                        <template slot-scope="scope">
                            <span>{{ command(scope.row) }}</span>
                        </template>
                    </el-table-column>
                    <el-table-column prop="status" label="状态" align="center">
                        <template slot-scope="scope">
                            <span :class="scope.row.status?statusText[scope.row.status][0]:''"></span>
                            <span>{{ scope.row.status?statusText[scope.row.status][1]:'' }}</span>
                        </template>
                    </el-table-column>
                </el-table>
            </div>
        </div>
    </div>
</template>

<script>
    export default {
        name: "taskProfile",
        props: {
            row: {
                type: Object,
                default: () => { }
            }
        },
        data() {
            return {
                data: {},
                tableData: [],
                statusText: { 'preparing': ['status-ready', '初始中'], 'pending': ['status-agent', '等待中'], 'running': ['status-running', '运行中'], 'failed': ['status-danger', '失败'], 'succeeded': ['status-success', '成功'], 'stopped': ['status-stopping', '已停止'] }
            }
        },
        computed: {
            show: function() {
                if (this.data.isDistributed === true) {
                    return true
                } else {
                    return false
                }
            }
        },
        created() {
            this.data = JSON.parse(JSON.stringify(this.row))
            if (this.data.isDistributed === true) {
                this.data.config.forEach(
                    (item, index) => {
                        item.status = item.subTaskState
                        if (item.subTaskState === 'unknown') {
                            item.status = undefined
                        }
                        if (item.replicaStates.length > 1) {
                            item.replicaStates = item.replicaStates.map((item, index) => { return { name: item.key, status: item.state, isChildren: true, id: index + item.key } })
                        } else {
                            item.replicaStates = []
                        }
                    }
                )
                this.tableData = this.data.config
            }
        },
        methods: {
            command: function(data) {
                let command = data.command
                if (data.parameters != null && data.parameters.length != 0) {
                    data.parameters.forEach(
                        item => {
                            if (item.key != '' || item.value != '') {
                                command += " " + '--' + item.key + '=' + item.value
                            }
                        }
                    )
                } else {
                    command = data.command
                }
                return command
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

    .taskList {
        font-weight: 800;
    }

    .block {
        float: right;
        margin: 20px;
    }
</style>