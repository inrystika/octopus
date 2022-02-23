<template>
    <div>
        <div class="searchForm">
            <searchForm :search-form="searchForm" :blur-name="'请输入服务名称'" @searchData="getSearchData" />
        </div>
        <el-button v-if="flag" type="primary" class="create" @click="create">创建</el-button>
        <el-table :data="tableData" style="width: 100%;font-size: 15px;"
            :header-cell-style="{'text-align':'left','color':'black'}" :cell-style="{'text-align':'left'}">
            <el-table-column label="名称" align="center">
                <template slot-scope="scope">
                    <span>{{ scope.row.name }}</span>
                </template>
            </el-table-column>
            <el-table-column label="模型名称" align="center">
                <template slot-scope="scope">
                    <span>{{ scope.row.modelName }}</span>
                </template>
            </el-table-column>
            <el-table-column label="模型版本" align="center">
                <template slot-scope="scope">
                    <span>{{ scope.row.modelVersion }}</span>
                </template>
            </el-table-column>
            <el-table-column label="模型描述" align="center" :show-overflow-tooltip="true">
                <template slot-scope="scope">
                    <span>{{ scope.row.desc }}</span>
                </template>
            </el-table-column>
            <el-table-column label="创建时间" align="center">
                <template slot-scope="scope">
                    <span>{{ scope.row.startedAt | parseTime }}</span>
                </template>
            </el-table-column>
            <el-table-column label="状态" align="center">
                <template slot-scope="scope">
                    <span :class="statusText[scope.row.status][0]" v-if="statusText[scope.row.status][0]"></span>
                    <span v-if="statusText[scope.row.status][1]">{{statusText[scope.row.status][1] }}</span>
                </template>
            </el-table-column>
            <el-table-column label="操作" align="center">
                <template slot-scope="scope">
                    <el-button
                        v-if="scope.row.status==='Available'||scope.row.status==='Creating'"
                        type="text" @click="open2(scope.row.id)">
                        停止
                    </el-button>
                    <el-button type="text" @click="handleDetail(scope.row)">详情</el-button>
                    <el-button
                        v-if="scope.row.status==='Failed'||scope.row.status==='Succeeded'||scope.row.status==='Stopped'"
                        type="text" @click="open(scope.row.id)">删除
                    </el-button>
                </template>
            </el-table-column>
        </el-table>
        <div class="block">
            <el-pagination :current-page="searchData.pageIndex" :page-sizes="[10, 20, 50, 80]"
                :page-size="searchData.pageSize" layout="total, sizes, prev, pager, next, jumper" :total="total"
                @size-change="handleSizeChange" @current-change="handleCurrentChange" />
        </div>
        <!-- 创建在线服务 -->
        <dialogForm v-if="createDialog" :row="row" @cancel="cancel" @confirm="confirm" @close="close" />
        <!-- 详情对话框 -->
        <detailDialog v-if="detailDialog" :row="row" @cancel="cancel" @confirm="confirm" @close="close" />
    </div>
</template>
<script>
    import searchForm from '@/components/search/index.vue'
    import dialogForm from "./components/dialogForm.vue";
    import detailDialog from "./components/index.vue";
    import { getDeployList, deleteDeploy, stopDeploy, deployDetail } from '@/api/deployManager.js'
    import store from '@/store'
    export default {
        name: "PreImage",
        components: {
            detailDialog,
            dialogForm,
            searchForm
        },
        props: {
        },
        data() {
            return {
                tableData: [],
                row: undefined,
                total: undefined,
                createDialog: false,
                detailDialog: false,
                flag: true,
                Logo: true,
                searchForm: [],
                searchData: {
                    pageIndex: 1,
                    pageSize: 10
                },
                statusText: { 'Preparing': ['status-ready', '初始中'], 'Creating': ['status-agent', '创建中'], 'Available': ['status-running', '运行中'], 'Failed': ['status-danger', '失败'],'Stopped': ['status-stopping', '已停止'] },
            }
        },
        created() {
            if (this.$route.params.flag) {
                this.createDialog = true
                this.row = this.$route.params.data
            }
            this.getList(this.searchData)
        },
        mounted() {
            window.addEventListener('beforeunload', e => {
                sessionStorage.clear()
            });

        },
        destroyed() {
            window.removeEventListener('beforeunload', e => {
                sessionStorage.clear()
            })
        },
        methods: {
            getList(data) {
                if (data.time && data.time.length !== 0) {
                    data.createAtGte = data.time[0] / 1000
                    data.createAtLt = data.time[1] / 1000
                    delete data.time
                }
                getDeployList(data).then(response => {
                    if (response.success) {
                        this.tableData = response.data.depInfos
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
                stopDeploy(id).then(response => {
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
                deleteDeploy({jobIds:val}).then(response => {
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
            },
            handleDetail(row) {
                deployDetail(row.id).then(response => {
                    if (response.success) {
                        this.row = response.data.depInfo
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
                this.stop(row);
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
            },
            getSearchData(val) {
                this.searchData = { pageIndex: 1, pageSize: this.searchData.pageSize }
                this.searchData = Object.assign(val, this.searchData)
                this.getList(this.searchData)
            },
            formatDuring(val) {
                return formatDuring(val)
            },
            // 删除确认
            // open(val) {
            //     let message = '此操作将永久删除该部署服务'
            //     this.$confirm(message, '提示', {
            //         confirmButtonText: '确定',
            //         cancelButtonText: '取消',
            //         type: 'warning'
            //     }).then(() => {
            //         this.handleDelete(val)
            //     }).catch(() => {
            //         this.$message({
            //             type: 'info',
            //             message: '已取消删除'
            //         });
            //     });
            // },
            // 停止确认
            open2(val) {
                this.$confirm('此操作将停止运行该部署服务, 是否继续?', '提示', {
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
    .block {
        float: right;
        margin: 20px;
    }

    .create {
        float: right;
    }

    .searchForm {
        display: inline-block;
    }
</style>