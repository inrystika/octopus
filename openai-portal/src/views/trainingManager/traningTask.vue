<template>
    <div>
        <div class="searchForm">
            <searchForm :searchForm="searchForm" :blurName="'任务名称 搜索'" @searchData="getSearchData" />
        </div>
        <el-button type="primary" class="create" @click="open()">批量删除</el-button>
        <el-button type="primary" class="create" @click="create">创建任务</el-button>
        <el-table
            ref="multipleTable"
            :data="tableData"
            style="width: 100%;font-size: 15px;"
            :header-cell-style="{'text-align':'left','color':'black'}"
            :cell-style="{'text-align':'left'}"
            @selection-change="handleSelectionChange"
        >
            <el-table-column type="selection" width="55" :selectable="checkSelectable" />
            <el-table-column label="任务名称" align="center">
                <template slot-scope="scope">
                    <span>{{ scope.row.name }}</span>
                </template>
            </el-table-column>
            <el-table-column label="算法名称" align="center">
                <template slot-scope="scope">
                    <span>{{ scope.row.algorithmName }}</span>
                </template>
            </el-table-column>
            <el-table-column label="数据集名称" align="center">
                <template slot-scope="scope">
                    <span>{{ scope.row.dataSetName }}</span>
                </template>
            </el-table-column>
            <el-table-column label="描述" align="center" :show-overflow-tooltip="true">
                <template slot-scope="scope">
                    <span>{{ scope.row.desc }}</span>
                </template>
            </el-table-column>
            <el-table-column label="分布式任务" align="center">
                <template slot-scope="scope">
                    <span>{{ scope.row.isDistributed?'是':'否' }}</span>
                </template>
            </el-table-column>
            <el-table-column label="状态" align="center">
                <template slot-scope="scope">
                    <span :class="statusText[scope.row.status][0]"></span>
                    <span>{{ statusText[scope.row.status][1] }}</span>
                </template>
            </el-table-column>
            <el-table-column label="创建时间" align="center">
                <template slot-scope="scope">
                    <span>{{ parseTime(scope.row.createdAt) }}</span>
                </template>
            </el-table-column>
            <el-table-column label="运行时长" align="center">
                <template slot-scope="scope">
                    <span>{{ formatDuring(scope.row.runSec) }}</span>
                </template>
            </el-table-column>
            <el-table-column label="操作" align="center">
                <template slot-scope="scope">
                    <el-button
                        v-if="scope.row.status==='pending'||scope.row.status==='running'||scope.row.status==='preparing'"
                        type="text"
                        @click="open2(scope.row)"
                    >
                        停止
                    </el-button>
                    <el-button type="text" @click="handleDetail(scope.row)">详情</el-button>
                    <el-button
                        v-if="scope.row.status==='failed'||scope.row.status==='succeeded'||scope.row.status==='stopped'"
                        type="text"
                        @click="open(scope.row)"
                    >删除
                    </el-button>
                </template>
            </el-table-column>
        </el-table>
        <div class="block">
            <el-pagination
                :current-page="searchData.pageIndex"
                :page-sizes="[10, 20, 50, 80]"
                :page-size="searchData.pageSize"
                layout="total, sizes, prev, pager, next, jumper"
                :total="total"
                @size-change="handleSizeChange"
                @current-change="handleCurrentChange"
            />
        </div>
        <!-- 创建对话框 -->
        <createDialog v-if="createDialog" :row="row" :flag="flag" @cancel="cancel" @confirm="confirm" @close="close" />
        <!-- 详情对话框 -->
        <detailDialog v-if="detailDialog" :data="data" @cancel="cancel" @confirm="confirm" @close="close" />

    </div>
</template>
<script>
    import createDialog from "./components/createDialog/index.vue";
    import detailDialog from "./components/detailDialog/index.vue";
    import { getList, stop, Delete, getTraningDetail } from '@/api/trainingManager'
    import searchForm from '@/components/search/index.vue'
    import { parseTime, formatDuring } from '@/utils/index'
    import { getErrorMsg } from '@/error/index'
    export default {
        name: "TraningTask",
        components: {
            createDialog,
            detailDialog,
            searchForm
        },
        props: {
          trainingTask: {
            type: Boolean,
            default: false
          }
        },
        data() {
            return {
                tableData: [],
                row: {
                },
                data: {},
                createDialog: false,
                detailDialog: false,
                total: undefined,
                statusText: { 'preparing': ['status-ready', '初始中'], 'pending': ['status-agent', '等待中'], 'running': ['status-running', '运行中'], 'failed': ['status-danger', '失败'], 'succeeded': ['status-success', '成功'], 'stopped': ['status-stopping', '已停止'] },
                flag: undefined,

                multipleSelection: [],
                searchForm: [
                    { type: 'Time', label: '创建时间', prop: 'time', placeholder: '请选择时间段' },
                    {
                        type: 'Select', label: '状态', prop: 'status', placeholder: '请选择状态',
                        options: [{ label: '成功', value: 'succeeded' }, { label: '失败', value: 'failed' }, { label: '运行中', value: 'running' }, { label: '等待中', value: 'pending' }, { label: '已停止', value: 'stopped' }, { label: '初始中', value: 'preparing' }]
                    }

                ],
                timer: null,
                searchData: {
                    pageSize: 10,
                    pageIndex: 1
                }

            }
        },
        created() {
            this.getList(this.searchData)
            if (this.trainingTask) {
                this.createDialog = true; this.row = {
                    algorithmSource: this.$route.params.data.type,
                    algorithmVersion: this.$route.params.data.algorithmVersion,
                    algorithmId: this.$route.params.data.algorithmId,
                    algorithmName: this.$route.params.data.algorithmName
                }
                // 算法页面创建训练任务跳转
                this.flag = 3
            }
        },

        methods: {
            // 错误码
            getErrorMsg(code) {
                return getErrorMsg(code)
            },
            getList(data) {
                if (data.time && data.time.length !== 0) {
                    data.createAtGte = data.time[0] / 1000
                    data.createAtLt = data.time[1] / 1000
                    delete data.time
                }
                getList(data).then(response => {
                    if (response.success) {
                        this.tableData = response.data.trainJobs
                        this.total = response.data.totalSize
                    } else {
                        this.$message({
                            message: this.getErrorMsg(response.error.subcode),
                            type: 'warning'
                        });
                    }
                })
            },
            stop(id) {
                stop(id).then(response => {
                    if (response.success) {
                        this.$message({
                            message: '停止成功',
                            type: 'success'
                        });
                    } else {
                        this.$message({
                            message: response.error.message,
                            type: 'warning'
                        });
                    }
                    this.getList(this.searchData)
                })
            },
            Delete(val) {
                var jobIds = []
                if (!val) {
                    if (this.multipleSelection.length !== 0) {
                        this.multipleSelection.forEach(
                            item => {
                                jobIds.push(item.id)
                            }
                        )
                        Delete({ jobIds }).then(response => {
                            if (response.success) {
                                this.$message({
                                    message: '删除成功',
                                    type: 'success'
                                });
                                this.getList(this.searchData)
                            } else {
                                this.$message({
                                    message: response.error.message,
                                    type: 'warning'
                                });
                            }
                        })
                    } else {
                        this.$message({
                            message: '请勾选需要删除的训练任务',
                            type: 'warning'
                        });
                    }
                } else {
                    jobIds = [val.id]
                    Delete({ jobIds }).then(response => {
                        if (response.success) {
                            this.$message({
                                message: '删除成功',
                                type: 'success'
                            });
                            this.getList(this.searchData)
                        } else {
                            this.$message({
                                message: response.error.message,
                                type: 'warning'
                            });
                        }
                    })
                }
            },
            handleDetail(row) {
                getTraningDetail(row.id).then(response => {
                    if (response.success) {
                        this.data = response.data.trainJob
                        this.detailDialog = true
                    } else {
                        this.$message({
                            message: this.getErrorMsg(response.error.subcode),
                            type: 'warning'
                        });
                    }
                })
            },
            handleStop(row) {
                this.stop(row.id);
            },
            handleDelete(val) {
                this.Delete(val);
            },
            handleSizeChange(val) {
                this.searchData.pageSize = val
                this.getList(this.searchData)
            },
            handleCurrentChange(val) {
                this.searchData.pageIndex = val
                this.getList(this.searchData)
            },
            cancel(val) {
                this.getList(this.searchData)
                this.createDialog = val;
                this.detailDialog = val
            },
            confirm(val) {
                this.getList(this.searchData)
                this.createDialog = val;
                this.detailDialog = val
            },
            close(val) {
                this.getList(this.searchData)
                this.createDialog = val;
                this.detailDialog = val
            },
            create() {
                this.createDialog = true; this.row = {}
                this.flag = 1
            },
            handleSelectionChange(val) {
                this.multipleSelection = val;
            },
            checkSelectable(row) {
                return row.status === 'failed' || row.status === 'succeeded' || row.status === 'stopped'
            },
            getSearchData(val) {
                this.searchData = { pageIndex: 1, pageSize: this.searchData.pageSize }
                this.searchData = Object.assign(val, this.searchData)
                this.getList(this.searchData)
            },
            // 时间戳转换日期
            parseTime(val) {
                return parseTime(val)
            },
            formatDuring(val) {
                return formatDuring(val)
            },
            // 删除确认
            open(val) {
                let message = ''
                if (val) { message = '此操作将永久删除该训练任务' } else { message = '此操作将永久批量删除该训练任务' }
                this.$confirm(message, '提示', {
                    confirmButtonText: '确定',
                    cancelButtonText: '取消',
                    type: 'warning'
                }).then(() => {
                    this.handleDelete(val)
                }).catch(() => {
                    this.$message({
                        type: 'info',
                        message: '已取消删除'
                    });
                });
            },
            // 停止确认
            open2(val) {
                this.$confirm('此操作将停止运行该训练任务, 是否继续?', '提示', {
                    confirmButtonText: '确定',
                    cancelButtonText: '取消',
                    type: 'warning'
                }).then(() => {
                    this.handleStop(val)
                }).catch(() => {
                    this.$message({
                        type: 'info',
                        message: '已取消操作'
                    });
                });
            }

        }
    }
</script>
<style lang="scss" scoped>
    .searchForm {
        display: inline-block;
    }

    .block {
        float: right;
        margin: 20px;
    }

    .create {
        float: right;
        margin-bottom: 15px;
        margin-left: 10px;
    }
</style>