<template>
    <div>
        <div class="searchForm">
            <searchForm :search-form="searchForm" :blur-name="'模板名称 搜索'" @searchData="getSearchData" />

        </div>
        <el-button type="primary" class="create" @click="open()">批量删除</el-button>
        <el-button type="primary" class="create" @click="create">创建模板</el-button>
        <div class="index">
            <el-table
                ref="multipleTable"
                :data="tableData"
                style="width: 100%;font-size: 15px;"
                :header-cell-style="{'text-align':'left','color':'black'}"
                :cell-style="{'text-align':'left'}"
                :selectable="checkSelectable"
                @selection-change="handleSelectionChange"
            >
                <el-table-column type="selection" width="55" />
                <el-table-column prop="name" label="任务模板名称" align="center" />
                <el-table-column prop="algorithmName" label="算法名称" align="center" />>
                <el-table-column prop="dataSetName" label="数据集名称" align="center" />
                <el-table-column prop="desc" label="描述" align="center" :show-overflow-tooltip="true" />
                <el-table-column label="是否为分布式">
                    <template slot-scope="scope">
                        <span>{{ scope.row.isDistributed?'是':'否' }}</span>
                    </template>
                </el-table-column>
                <el-table-column label="创建时间" align="center">
                    <template slot-scope="scope">
                        <span>{{ scope.row.createdAt | parseTime }}</span>
                    </template>
                </el-table-column>
                <el-table-column label="操作" width="250">
                    <template slot-scope="scope">
                        <el-button type="text" @click="handleEdit(scope.row,'editTemplate')">编辑</el-button>
                        <el-button type="text" @click="open(scope.row)">删除</el-button>
                        <el-button type="text" @click="handleEdit(scope.row,'createTask')">创建训练任务</el-button>
                        <el-button type="text" @click="handleCopy(scope.row)">复制</el-button>
                    </template>
                </el-table-column>
            </el-table>

        </div>
        <!-- 创建任务模板 -->
        <createDialog v-if="createDialog" :flag="flag" :row="row" @cancel="cancel" @confirm="confirm" @close="close" />
        <!-- 编辑对话框 -->
        <editDialog
            v-if="editDialog"
            :flag="flag"
            :row="row"
            @cancel="cancel"
            @confirm="confirm"
            @close="close"
            @createTraning="createTraning"
        />
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
    </div>
</template>
<script>
    import { getTemplate, getTempalteDetail, deleteTemplate, getResourceList, copyTemplate } from '@/api/trainingManager'
    import createDialog from "./components/createDialog/index.vue";
    import editDialog from "./components/editDialog/index.vue";
    import searchForm from '@/components/search/index.vue'
    export default {
        name: "PreImage",
        components: {
            createDialog,
            editDialog,
            searchForm

        },
        props: {
            trainingTemplate: {
                type: Boolean,
                default: false
            }
        },
        data() {
            return {
                tableData: [],
                createDialog: false,
                editDialog: false,
                row: {},
                flag: undefined,
                pageIndex: 1,
                pageSize: 10,
                total: undefined,
                multipleSelection: [],
                resourceOptions: [],
                searchForm: [
                    { type: 'Time', label: '创建时间', prop: 'time', placeholder: '请选择时间段' }

                ],
                timer: null,
                searchData: {
                    pageSize: 10,
                    pageIndex: 1
                }
            }
        },
        created() {
            this.getTemplate(this.searchData)
            if (this.trainingTemplate) {
                this.createDialog = true;
                this.flag = 2
            }
        },
        methods: {
            getTemplate(data) {
                if (data.time && data.time.length !== 0) {
                    data.createAtGte = data.time[0] / 1000
                    data.createAtLt = data.time[1] / 1000
                    delete data.time
                }
                getTemplate(data).then(response => {
                    if (response.success) {
                        this.tableData = response.data.jobTemplates
                        this.total = response.data.totalSize
                    } else {
                        this.$message({
                            message: this.getErrorMsg(response.error.subcode),
                            type: 'warning'
                        });
                    }
                })
            },
            handleEdit(val, name) {
                getTempalteDetail(val.id).then(response => {
                            if (response.success) {
                                this.editDialog = true
                                this.row = response.data.jobTemplate
                                if (name === 'editTemplate') { this.flag = 1 } else { this.flag = 2 }
                            } else {
                                this.$message({
                                    message: this.getErrorMsg(response.error.subcode),
                                    type: 'warning'
                                });
                            }
                        })
            },
            handleDelete(val) {
                let templateIds = []
                if (!val) {
                    if (this.multipleSelection.length !== 0) {
                        this.multipleSelection.forEach(
                            item => { templateIds.push(item.id) }
                        )
                        deleteTemplate({ templateIds }).then(response => {
                            if (response.success) {
                                this.$message({
                                    message: '删除成功',
                                    type: 'success'
                                });
                                this.getTemplate(this.searchData)
                            } else {
                                this.$message({
                                    message: this.getErrorMsg(response.error.subcode),
                                    type: 'warning'
                                });
                            }
                        })
                    } else {
                        this.$message({
                            message: '请勾选需要删除的任务模板',
                            type: 'warning'
                        });
                    }
                } else {
                    templateIds = [val.id]
                    deleteTemplate({ templateIds }).then(response => {
                        if (response.success) {
                            this.$message({
                                message: '删除成功',
                                type: 'success'
                            });
                            this.getTemplate(this.searchData)
                        } else {
                            this.$message({
                                message: this.getErrorMsg(response.error.subcode),
                                type: 'warning'
                            });
                        }
                    })
                }
            },
            handleCopy(row) {
                copyTemplate(row).then(response => {
                    if (response.success) {
                        this.$message({
                            message: '复制成功',
                            type: 'success'
                        });
                        this.getTemplate(this.searchData)
                    } else {
                        this.$message({
                            message: this.getErrorMsg(response.error.subcode),
                            type: 'warning'
                        });
                    }
                })
            },
            cancel(val) {
                this.getTemplate(this.searchData)
                this.createDialog = val
                this.editDialog = val
            },
            confirm(val) {
                this.getTemplate(this.searchData)
                this.createDialog = val
                this.editDialog = val
            },
            close(val) {
                this.getTemplate(this.searchData)
                this.createDialog = val
                this.editDialog = val
            },
            // 新增创建任务模板
            create() {
                this.row = {}
                this.createDialog = true
                this.flag = 2
            },
            handleSizeChange(val) {
                this.searchData.pageSize = val
                this.getTemplate(this.searchData)
            },
            handleCurrentChange(val) {
                this.searchData.pageIndex = val
                this.getTemplate(this.searchData)
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
                this.getTemplate(this.searchData)
            },
            // 删除确认
            open(val) {
                let message = ''
                if (val) { message = '此操作将永久删除该任务模板' } else { message = '此操作将永久批量删除该任务模板' }
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
            // 创建任务跳转
            createTraning() {
                // 这里写你要跳转的标签页的name
                this.$emit('createTraning')
                // this.resetZuJian()
            }

        }
    }
</script>

<style lang="scss" scoped>
    .searchForm {
        display: inline-block;
    }

    .create {
        float: right;
        margin-bottom: 15px;
        margin-left: 10px;
    }

    .block {
        float: right;
        margin: 20px;
    }
</style>