<template>
    <div>
        <searchForm :search-form="searchForm" :blur-name="'任务名称 搜索'" @searchData="getSearchData" />
        <el-table :data="tableData" style="width: 100%;font-size: 15px;"
            :header-cell-style="{'text-align':'left','color':'black'}" :cell-style="{'text-align':'left'}">
            <el-table-column label="用户名">
                <template slot-scope="scope">
                    <span>{{ scope.row.name }}</span>
                </template>
            </el-table-column>
            <el-table-column label="群组">
                <template slot-scope="scope">
                    <span>{{ scope.row.workspaceId }}</span>
                </template>
            </el-table-column>
            <el-table-column label="模型名称">
                <template slot-scope="scope">
                    <span>{{ scope.row.modelName }}</span>
                </template>
            </el-table-column>
            <el-table-column label="版本">
                <template slot-scope="scope">
                    <span>{{ scope.row.modelVersion }}</span>
                </template>
            </el-table-column>
            <el-table-column label="模型描述">
                <template slot-scope="scope">
                    <span>{{ scope.row.desc }}</span>
                </template>
            </el-table-column>
            <el-table-column label="URL">
                <template slot-scope="scope">
                    <span>{{ scope.row.serviceUrl }}</span>
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
                    <span>{{ parseTime(scope.row.startedAt) }}</span>
                </template>
            </el-table-column>
            <el-table-column label="操作">
                <template slot-scope="scope">
                    <el-button
                        v-if="scope.row.status==='pending'||scope.row.status==='running'||scope.row.status==='preparing'"
                        type="text" @click="open(scope.row)">停止
                    </el-button>
                    <el-button type="text" @click="handledetail( scope.row)">详情</el-button>
                </template>
            </el-table-column>
        </el-table>
        <div class="block">
            <el-pagination :current-page="searchData.pageIndex" :page-sizes="[10, 20, 30,50]"
                :page-size="searchData.pageSize" :total="total" layout="total, sizes, prev, pager, next, jumper"
                @size-change="handleSizeChange" @current-change="handleCurrentChange" />
        </div>
        <!-- 详情对话框 -->
        <detailDialog v-if="detailDialog" :data="data" @cancel="cancel" @confirm="confirm" @close="close" />
    </div>
</template>

<script>
    import detailDialog from "./components/index.vue";
    import { getDeployList } from '@/api/deployManager.js'
    import { parseTime, formatDuring } from '@/utils/index'
    import searchForm from '@/components/search/index.vue'
    import { getErrorMsg } from '@/error/index'
    export default {
        name: "TraningTask",
        components: {
            detailDialog,
            searchForm
        },
        data() {
            return {
                total: undefined,
                tableData: [{}],
                detailDialog: false,
                data: {},
                statusText: { 'preparing': ['status-ready', '初始中'], 'available': ['status-agent', '可部署'], 'creating': ['status-running', '创建中'], 'failed': ['status-danger', '失败'], 'stopped': ['status-stopping', '已停止'] },
                searchForm: [
                ],
                searchData: {
                    pageIndex: 1,
                    pageSize: 10
                }

            }
        },
        created() {
            // this.getDeployList(this.searchData)
        },
        beforeDestroy() {

        },
        methods: {
            // 错误码
            getErrorMsg(code) {
                return getErrorMsg(code)
            },
            getDeployList(data) {
                if (data.time && data.time.length !== 0) {
                    data.createAtGte = data.time[0] / 1000
                    data.createAtLt = data.time[1] / 1000
                    delete data.time
                }
                getDeployList(data).then(response => {
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
            handledetail(row) {
                this.detailDialog = true
                // deployDetail(row.id).then(response => {
                //     if (response.success) {
                //         this.data = response.data.trainJob
                //         this.detailDialog = true
                //     } else {
                //         this.$message({
                //             message: this.getErrorMsg(response.error.subcode),
                //             type: 'warning'
                //         });
                //     }
                // })
            },
            open(val) {
                this.$confirm('此操作将停止改模型部署, 是否继续?', '提示', {
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
                stopDeploy(row.id).then(response => {
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
                    this.getDeployList(this.searchData)
                })
            },
            handleSizeChange(val) {
                this.searchData.pageSize = val
                this.getDeployList(this.searchData)
            },
            handleCurrentChange(val) {
                this.searchData.pageIndex = val
                this.getDeployList(this.searchData)
            },
            cancel(val) {
                this.detailDialog = val
                this.getDeployList(this.searchData)
            },
            confirm(val) {
                this.detailDialog = val
                this.getDeployList(this.searchData)
            },
            close(val) {
                this.detailDialog = val
                this.getDeployList(this.searchData)
            },
            getSearchData(val) {
                this.searchData = { pageIndex: 1, pageSize: this.searchData.pageSize }
                this.searchData = Object.assign(val, this.searchData)
                this.getDeployList(this.searchData)
            },
            // 时间戳转换日期
            parseTime(val) {
                return parseTime(val)
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