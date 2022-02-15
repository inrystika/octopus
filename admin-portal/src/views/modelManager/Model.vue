<template>
    <div>
        <div class="searchForm">
            <searchForm :search-form="searchForm" :blur-name="'模型名称/描述 搜索'" @searchData="getSearchData" />
        </div>
        <el-button v-if="modelTabType===3" type="primary" class="create" @click="createModel">创建</el-button>
        <div class="index">
            <el-table :data="tableData" style="width: 100%;font-size: 15px;"
                :header-cell-style="{'text-align':'left','color':'black'}" :cell-style="{'text-align':'left'}">
                <el-table-column prop="modelName" label="模型名称" />
                <el-table-column prop="latestVersion" label="模型版本" />
                <el-table-column prop="algorithmName" label="算法名称" />
                <el-table-column prop="algorithmVersion" label="算法版本" />
                <el-table-column prop="modelDescript" label="模型描述" />
                <el-table-column label="群组名">
                    <template slot-scope="scope">
                        <span>{{ scope.row.spaceName===''?'默认群组':scope.row.spaceName }}</span>
                    </template>
                </el-table-column>
                <el-table-column v-if="modelTabType===1" label="提供者">
                    <template slot-scope="scope">
                        <el-tooltip trigger="hover" :content="scope.row.userEmail" placement="top">
                            <span>{{ scope.row.userName }}</span>
                        </el-tooltip>
                    </template>
                </el-table-column>
                <el-table-column label="创建时间">
                    <template slot-scope="scope">
                        <span>{{ scope.row.createdAt | parseTime }}</span>
                    </template>
                </el-table-column>
                <el-table-column label="操作">
                    <template slot-scope="scope">
                        <el-button type="text" @click="getVersionList(scope.row)">版本列表</el-button>
                        <el-button v-if="modelTabType===3" type="text" @click="createList(scope.row)">创建新版本</el-button>
                        <el-button v-if="modelTabType===3" type="text" @click="open(scope.row)">删除</el-button>
                    </template>
                </el-table-column>
            </el-table>
        </div>
        <div class="block">
            <el-pagination :current-page="searchData.pageIndex" :page-sizes="[10, 20, 50,80]"
                :page-size="searchData.pageSize" :total="total" layout="total, sizes, prev, pager, next, jumper"
                @size-change="handleSizeChange" @current-change="handleCurrentChange" />
        </div>
        <!-- 版本列表对话框 -->
        <versionList v-if="FormVisible" :model-id="modelId" :model-name="modelName" :model-type="type" @close="close"
            @cancel="cancel" @confirm="confirm" />
        <!-- 创建对话框 -->
        <createDialog v-if="CreateVisible" :row="row" :is-list="isList" @close="close" @cancel="cancel"
            @confirm="confirm" />
    </div>
</template>

<script>
    import versionList from './components/versionList.vue'
    import createDialog from './components/createDialog.vue'
    import { getMyModel, getPreModel, deletePreModel } from '@/api/modelManager.js'
    import searchForm from '@/components/search/index.vue'
    export default {
        name: "MyModel",
        components: {
            versionList,
            searchForm,
            createDialog
        },
        props: {
            modelTabType: { type: Number, default: undefined }
        },
        data() {
            return {
                input: '',
                tableData: [

                ],
                total: undefined,
                FormVisible: false,
                dialogVisible: false,
                modelId: undefined,
                type: undefined,
                // timer: null,
                searchForm: [
                    { type: 'InputSelectUser', label: '用户', prop: 'userId', placeholder: '请输入用户名' },
                    { type: 'InputSelectGroup', label: '群组', prop: 'spaceId', placeholder: '请输入群组名' }

                ],
                CreateVisible: false,
                row: {},
                isList: true,
                searchData: {
                    pageIndex: 1,
                    pageSize: 10
                },
                modelName: ''
            }
        },
        created() {
            this.getModel(this.searchData)
            // this.timer = setInterval(this.getModel, 1000);
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
                this.CreateVisible = val;
                this.getModel(this.searchData)
            },
            cancel(val) {
                this.FormVisible = val;
                this.CreateVisible = val;
                this.getModel(this.searchData)
            },
            confirm(val) {
                this.FormVisible = val;
                this.CreateVisible = val;
                this.getModel(this.searchData)
            },
            getVersionList(val) {
                this.FormVisible = true
                this.modelId = val.modelId
                this.modelName = val.modelName
            },
            handledDelete(row) {
                const data = JSON.parse(JSON.stringify(row));
                data.version = row.modelVersion
                deletePreModel(data).then(response => {
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
                if (!data) { data = { pageIndex: this.pageIndex, pageSize: this.pageSize } }
                this.type = this.modelTabType
                if (this.type === 1) {
                    getMyModel(data).then(response => {
                        if (response.success) {
                            this.total = response.data.totalSize
                            this.tableData = response.data.models
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
                            this.total = response.data.totalSize
                            this.tableData = response.data.models
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
            createModel() {
                this.isList = false
                this.CreateVisible = true
                this.row = {}
            },
            createList(val) {
                this.isList = true
                this.CreateVisible = true
                this.row = val
            },
            // 删除确认
            open(val) {
                this.$confirm('此操作将永久删除该预置模型, 是否继续?', '提示', {
                    confirmButtonText: '确定',
                    cancelButtonText: '取消',
                    type: 'warning'
                }).then(() => {
                    this.handledDelete(val)
                }).catch(() => {
                    this.$message({
                        type: 'info',
                        message: '已取消删除'
                    });
                });
            }

        }
    }
</script>

<style lang="scss" scoped>
    .create {
        float: right;
        margin: 10px;
    }

    .block {
        float: right;
        margin: 20px;
    }

    .searchForm {
        display: inline-block;
    }
</style>