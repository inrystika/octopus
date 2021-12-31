<template>
    <div>
        <searchForm :search-form="searchForm" :blur-name="'任务名称 搜索'" @searchData="getSearchData" />
        <el-table
            :data="tableData"
            style="width: 100%;font-size: 15px;"
            :header-cell-style="{'text-align':'left','color':'black'}"
            :cell-style="{'text-align':'left'}"
        >
            <el-table-column label="任务名称">
                <template slot-scope="scope">
                    <span>{{ scope.row.name }}</span>
                </template>
            </el-table-column>
            <el-table-column label="算法名称">
                <template slot-scope="scope">
                    <span>{{ scope.row.algorithmName }}</span>
                </template>
            </el-table-column>
            <el-table-column label="数据集名称">
                <template slot-scope="scope">
                    <span>{{ scope.row.dataSetName }}</span>
                </template>
            </el-table-column>
            <el-table-column label="群组">
                <template slot-scope="scope">
                    <span>{{ scope.row.workspaceName==""?'默认群组':scope.row.workspaceName }}</span>
                </template>
            </el-table-column>
            <el-table-column label="用户名">
                <template slot-scope="scope">
                    <span>{{ scope.row.userName }}</span>
                </template>
            </el-table-column>
            <el-table-column label="状态">
                <template slot-scope="scope">
                    <span :class="scope.row.status?statusText[scope.row.status][0]:''"></span>
                    <span>{{ scope.row.status?statusText[scope.row.status][1]:'' }}</span>
                </template>
            </el-table-column>
            <el-table-column label="创建时间">
                <template slot-scope="scope">
                    <span>{{ scope.row.createdAt | parseTime }}</span>
                </template>
            </el-table-column>
            <el-table-column label="运行时长" align="center">
                <template slot-scope="scope">
                    <span>{{ formatDuring(scope.row.runSec) }}</span>
                </template>
            </el-table-column>
            <el-table-column label="操作">
                <template slot-scope="scope">
                    <el-button
                        v-if="scope.row.status==='pending'||scope.row.status==='running'||scope.row.status==='preparing'"
                        type="text"
                        @click="open(scope.row)"
                    >
                        停止
                    </el-button>
                    <el-button type="text" @click="handledetail( scope.row)">详情</el-button>
                </template>
            </el-table-column>
        </el-table>
        <div class="block">
            <el-pagination
                :current-page="searchData.pageIndex"
                :page-sizes="[10, 20, 30,50]"
                :page-size="searchData.pageSize"
                :total="total"
                layout="total, sizes, prev, pager, next, jumper"
                @size-change="handleSizeChange"
                @current-change="handleCurrentChange"
            />
        </div>
        <!-- 详情对话框 -->
        <detailDialog v-if="detailDialog" :data="data" @cancel="cancel" @confirm="confirm" @close="close" />
    </div>
</template>

<script>
    import detailDialog from "./components/index.vue";
    import { getTraining, stopTraining, trainingDetail } from '@/api/trainingManager.js'
    import { formatDuring } from '@/utils/index'
    import searchForm from '@/components/search/index.vue'
    export default {
        name: "TraningTask",
        components: {
            detailDialog,
            searchForm
        },
        data() {
            return {
                total: undefined,
                tableData: [],
                state: undefined,
                detailDialog: false,
                // 训练详情
                data: {},
                statusText: { 'preparing': ['status-ready', '初始中'], 'pending': ['status-agent', '等待中'], 'running': ['status-running', '运行中'], 'failed': ['status-danger', '失败'], 'succeeded': ['status-success', '成功'], 'stopped': ['status-stopping', '已停止'] },
                searchForm: [
                    { type: 'Time', label: '创建时间', prop: 'time', placeholder: '请选择时间段' },
                    {
                        type: 'Select', label: '状态', prop: 'status', placeholder: '请选择状态',
                        options: [{ label: '成功', value: 'succeeded' }, { label: '失败', value: 'failed' }, { label: '运行中', value: 'running' }, { label: '等待中', value: 'pending' }, { label: '已停止', value: 'stopped' }, { label: '初始中', value: 'preparing' }]
                    },
                    { type: 'Input', label: '用户名', prop: 'userNameLike', placeholder: '请输入用户名' }

                ],
                searchData: {
                    pageIndex: 1,
                    pageSize: 10
                }

            }
        },
        created() {
            this.getTraining(this.searchData)
        },
        beforeDestroy() {

        },
        methods: {
            handledetail(row) {
                trainingDetail(row.id).then(response => {
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
            open(val) {
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
            },
            handleStop(row) {
                stopTraining(row.id).then(response => {
                    if (response.success) {
                        this.$message({
                            message: '停止成功',
                            type: 'success'
                        });
                    } else {
                        this.$message({
                            message: this.getErrorMsg(response.error.subcode),
                            type: 'warning'
                        });
                    }
                    this.getTraining(this.searchData)
                })
            },
            handleSizeChange(val) {
                this.searchData.pageSize = val
                this.getTraining(this.searchData)
            },
            handleCurrentChange(val) {
                this.searchData.pageIndex = val
                this.getTraining(this.searchData)
            },
            cancel(val) {
                this.detailDialog = val
                this.getTraining(this.searchData)
            },
            confirm(val) {
                this.detailDialog = val
                this.getTraining(this.searchData)
            },
            close(val) {
                this.detailDialog = val
                this.getTraining(this.searchData)
            },
            getTraining(data) {
                if (data.time && data.time.length !== 0) {
                    data.createAtGte = data.time[0] / 1000
                    data.createAtLt = data.time[1] / 1000
                    delete data.time
                }
                getTraining(data).then(response => {
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
            getSearchData(val) {
                this.searchData = { pageIndex: 1, pageSize: this.searchData.pageSize }
                this.searchData = Object.assign(val, this.searchData)
                this.getTraining(this.searchData)
            },
            // 时间戳转换日期
            formatDuring(val) {
                return formatDuring(val)
            }
        }
    }
</script>
<style lang="scss" scoped>
    .function {
        float: right;
        margin: 10px;
    }

    .block {
        float: right;
        margin: 20px;
    }
</style>