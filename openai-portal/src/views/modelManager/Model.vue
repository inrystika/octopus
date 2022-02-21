<template>
    <div>
        <div>
            <searchForm :search-form="searchForm" :blur-name="'名称/描述 搜索'" @searchData="getSearchData" />
        </div>
        <div class="index">
            <el-table :data="tableData" style="width: 100%;font-size: 15px;"
                :header-cell-style="{'text-align':'left','color':'black'}" :cell-style="{'text-align':'left'}">
                <el-table-column prop="modelName" label="模型名称" align="center" />
                <el-table-column v-if="modelTabType===2" prop="userName" label="提供者" align="center" />
                <el-table-column prop="algorithmName" label="算法名称" align="center" />
                <el-table-column prop="algorithmVersion" label="算法版本" align="center" />
                <el-table-column prop="modelDescript" label="模型描述" align="center" />
                <el-table-column label="创建时间" align="center">
                    <template slot-scope="scope">
                        <span>{{ scope.row.createdAt | parseTime }}</span>
                    </template>
                </el-table-column>
                <el-table-column label="操作" align="center">
                    <template slot-scope="scope">
                        <el-button type="text" @click="getVersionList(scope.row)">版本列表</el-button>
                        <el-button v-if="type===1" type="text" @click="open(scope.row)">删除</el-button>
                    </template>
                </el-table-column>
            </el-table>
        </div>
        <div class="block">
            <el-pagination :current-page="searchData.pageIndex" :page-sizes="[10, 20, 50,80]"
                :page-size="searchData.pageSize" layout="total, sizes, prev, pager, next, jumper" :total="total"
                @size-change="handleSizeChange" @current-change="handleCurrentChange" />
        </div>
        <!-- 版本列表对话框 -->
        <versionList v-if="FormVisible" :model-id="modelId" :model-type="type" :model-name="modelName" @close="close"
            @cancel="cancel" @confirm="confirm" :showDeploy="showDeploy" />
    </div>
</template>

<script>
    import versionList from './components/versionList.vue'
    import { getMyModel, getPreModel, getPublicModel, deleteMyModel } from '@/api/modelManager.js'
    import searchForm from '@/components/search/index.vue'
    import { getPresetAlgorithmList, getPublicAlgorithmList, getMyAlgorithmList } from '@/api/modelDev'
    export default {
        name: "MyModel",
        components: {
            versionList,
            searchForm
        },
        props: {
            modelTabType: { type: Number, default: undefined }
        },
        data() {
            return {
                input: '',
                tableData: [],
                total: undefined,
                FormVisible: false,
                dialogVisible: false,
                modelId: undefined,
                type: undefined,
                searchForm: [],
                searchData: {
                    pageIndex: 1,
                    pageSize: 10
                },
                modelName: '',
                showDeploy: false
            }
        },
        created() {
            this.getModel(this.searchData)
            if (this.modelTabType !== 1) {
                this.flag = false
            }
        },
        methods: {
            handleSizeChange(val) {
                this.searchData.pageSize = val
                this.getModel(this.searchData)
            },
            handleCurrentChange(val) {
                this.searchData.pageIndex = val
                this.getModel(this.searchData)
            },
            close(val) {
                this.FormVisible = val;
                this.dialogVisible = val
                this.getModel(this.searchData)
            },
            cancel(val) {
                this.FormVisible = val;
                this.dialogVisible = val
                this.getModel(this.searchData)
            },
            confirm(val) {
                this.FormVisible = val;
                this.dialogVisible = val
                this.getModel(this.searchData)
            },
            getVersionList(val) {
                let frameworkName = ""
                if (this.modelTabType === 1) {
                    getMyAlgorithmList({ pageIndex: 1, pageSize: 10, searchKey: val.algorithmName }).then(response => {
                        if (response.success) {
                            response.data.algorithms[0] ? frameworkName = response.data.algorithms[0].frameworkName : frameworkName = ""
                            if (frameworkName === "TensorFlow" || frameworkName === "Pytorch") {
                                this.showDeploy = true
                            }
                            else { this.showDeploy = false }
                            this.FormVisible = true;
                            this.modelId = val.modelId
                            this.modelName = val.modelName
                        }

                    })
                }
                if (this.modelTabType === 3) {
                    getPresetAlgorithmList({ pageIndex: 1, pageSize: 10, searchKey: val.algorithmName }).then(response => {
                        if (response.success) {
                            response.data.algorithms[0] ? frameworkName = response.data.algorithms[0].frameworkName : frameworkName = ""
                            if (frameworkName === "TensorFlow" || frameworkName === "Pytorch") {
                                this.showDeploy = true
                            }
                            else { this.showDeploy = false }
                            this.FormVisible = true;
                            this.modelId = val.modelId
                            this.modelName = val.modelName
                        }
                    })
                }
                if (this.modelTabType === 2) {
                    getPublicAlgorithmList({ pageIndex: 1, pageSize: 10, searchKey: val.algorithmName }).then(response => {
                        if (response.success) {
                            response.data.algorithms[0] ? frameworkName = response.data.algorithms[0].frameworkName : frameworkName = ""
                            if (frameworkName === "TensorFlow" || frameworkName === "Pytorch") {
                                this.showDeploy = true
                            }
                            else { this.showDeploy = false }
                            this.FormVisible = true;
                            this.modelId = val.modelId
                            this.modelName = val.modelName
                        }
                    })

                }

            },
            handleDelete(row) {
                const data = JSON.parse(JSON.stringify(row));
                data.version = row.modelVersion
                deleteMyModel(data).then(response => {
                    if (response.success) {
                        this.$message({
                            message: '删除成功',
                            type: 'success'
                        });
                        this.getModel(this.searchData)
                    } else {
                        this.$message({
                            message: this.getErrorMsg(response.error.subcode),
                            type: 'warning'
                        });
                    }
                })
            },
            getModel(data) {
                this.type = this.modelTabType
                if (this.type === 1) {
                    getMyModel(data).then(response => {
                        if (response.success) {
                            if (response.data.models !== null) {
                                this.total = response.data.totalSize
                                this.tableData = response.data.models
                            }
                        } else {
                            this.$message({
                                message: this.getErrorMsg(response.error.subcode),
                                type: 'warning'
                            });
                        }
                    })
                }
                if (this.type === 2) {
                    getPublicModel(data).then(response => {
                        if (response.success) {
                            if (response.data.models !== null) {
                                this.total = response.data.totalSize
                                this.tableData = response.data.models
                            }
                        } else {
                            this.$message({
                                message: this.getErrorMsg(response.error.subcode),
                                type: 'warning'
                            });
                        }
                    })
                }
                if (this.type === 3) {
                    getPreModel(data).then(response => {
                        if (response.success) {
                            if (response.data.models !== null) {
                                this.total = response.data.totalSize
                                this.tableData = response.data.models
                            }
                        } else {
                            this.$message({
                                message: this.getErrorMsg(response.error.subcode),
                                type: 'warning'
                            });
                        }
                    })
                }
            },
            getSearchData(val) {
                this.searchData = { pageIndex: 1, pageSize: this.searchData.pageSize }
                this.searchData = Object.assign(val, this.searchData)
                this.getModel(this.searchData)
            },
            // 删除确认
            open(val) {
                this.$confirm('此操作将永久删除该模型(如该模型已分享，则分享模型也会被删除)，是否继续?', '提示', {
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