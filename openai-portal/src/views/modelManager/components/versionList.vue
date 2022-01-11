<template>
    <div>
        <el-dialog :title="'版本列表/'+modelName" width="70%" :visible.sync="CreateFormVisible"
            :before-close="handleDialogClose" class="wrapper" :close-on-click-modal="false">
            <el-table :data="tableData" style="width: 100%" height="500">
                <el-table-column prop="version" label="模型版本" align="center" />
                <el-table-column prop="descript" label="模型描述" align="center" />
                <el-table-column label="状态" align="center">
                    <template slot-scope="scope">
                        <span style="margin-left: 10px">{{ fileStatus(scope.row.fileStatus) }}</span>
                    </template>
                </el-table-column>
                <el-table-column label="创建时间" align="center">
                    <template slot-scope="scope">
                        <span style="margin-left: 10px">{{ scope.row.createdAt | parseTime }}</span>
                    </template>
                </el-table-column>
                <el-table-column label="操作" align="center" width="300px">
                    <template slot-scope="scope">
                        <el-button v-if="modelType===3" type="text" @click="deploy(scope.row)">部署</el-button>
                        <el-button type="text" :disabled="scope.row.fileStatus!==2" @click="handlePreview(scope.row)">预览
                        </el-button>
                        <el-button v-if="!scope.row.isShared&&modelType===1" type="text" @click="open(scope.row)">分享
                        </el-button>
                        <el-button v-if="scope.row.isShared&&scope.row.isShared&&modelType===1" type="text"
                            @click="open(scope.row)">取消分享</el-button>
                        <el-button v-if="modelType===1" type="text" @click="open2(scope.row)">删除</el-button>
                        <!-- <el-button type="text" @click="handleDelete(scope.row)" v-if="modelType==1">删除</el-button> -->
                        <el-button type="text" :disabled="scope.row.fileStatus!==2" @click="handledDownload(scope.row)">
                            下载</el-button>
                    </template>
                </el-table-column>
            </el-table>
            <div class="block">
                <el-pagination :current-page="pageIndex" :page-sizes="[10, 20, 50, 80]" :page-size="pageSize"
                    layout="total, sizes, prev, pager, next, jumper" :total="total" @size-change="handleSizeChange"
                    @current-change="handleCurrentChange" />
            </div>
            <div slot="footer">
                <!-- <el-button @click="cancel">取 消</el-button>
                <el-button type="primary" @click="confirm">确 定</el-button> -->
            </div>
        </el-dialog>
        <!-- 预览对话框 -->
        <previewDialog v-if="preVisible" :row="data" @close="closeShareDialog" />
    </div>
</template>

<script>
    import previewDialog from './previewDialog.vue'
    import { getPublicList, getNoPublicList, downloadModel, deleteModelVersion, shareModel, cancelShareModel } from '@/api/modelManager.js'
    export default {
        name: "VersionList",
        components: {
            previewDialog
        },
        props: {
            modelId: {
                type: String,
                default: ""
            },
            modelType: { type: Number, default: undefined },
            modelName: { type: String, default: "" }
        },
        data() {
            return {
                CreateFormVisible: true,
                dialogVisible: false,
                preVisible: false,
                pageIndex: 1,
                total: 1,
                pageSize: 10,
                tableData: [],
                row: { flag: undefined, data: undefined },
                data: {}

            }
        },
        created() {
            this.getList()
        },
        beforeDestroy() {

        },
        methods: {
            handlePreview(row) {
                this.preVisible = true
                this.data = row
            },
            handleShare(row) {
                this.row.flag = true
                this.row.data = row
                this.dialogVisible = true;
            },
            handleCancelShare(row) {
                this.row.data = row
                this.row.flag = false
                this.dialogVisible = true;
            },
            handleDelete(row) {
                const data = JSON.parse(JSON.stringify(row));
                data.version = row.version
                data.modelId = row.modelId
                deleteModelVersion(data).then(response => {
                    if (response.success) {
                        this.$message({
                            message: '删除成功',
                            type: 'success'
                        });
                        this.getList()
                    } else {
                        this.$message({
                            message: this.getErrorMsg(response.error.subcode),
                            type: 'warning'
                        });
                    }
                })
            },
            handledDownload(row) {
                const data = {}
                data.version = row.version
                data.modelId = row.modelId
                data.domian = this.GLOBAL.DOMAIN
                downloadModel(data).then(response => {
                    if (response.success) { this.URLdownload(data, response.data.downloadUrl) } else {
                        this.$message({
                            message: this.getErrorMsg(response.error.subcode),
                            type: 'warning'
                        });
                    }
                })
            },
            // 接受到url下载
            URLdownload(data, url) {
                const link = document.createElement('a')
                link.style = 'display: none'; // 创建一个隐藏的a标签
                // link.setAttribute('download', '下载文件')
                link.setAttribute('href', url)
                link.setAttribute('target', "_blank")
                document.body.appendChild(link);
                link.click(); // 触发a标签的click事件
                document.body.removeChild(link);
                this.$message({
                    message: '下载成功',
                    type: 'success'
                });
            },
            cancel() {
                this.$emit('cancel', false)
            },
            confirm() {
                this.$emit('confirm', false)
            },
            handleDialogClose() {
                this.$emit('close', false)
            },
            handleSizeChange(val) {
                this.pageSize = val
                this.getList()
            },
            handleCurrentChange(val) {
                this.pageIndex = val
                this.getList()
            },
            closeShareDialog(val) {
                this.dialogVisible = val
                this.preVisible = val
                this.getList()
            },
            cancelShareDialog(val) {
                this.dialogVisible = val
                this.getList()
            },
            confirmShareDialog(val) {
                this.dialogVisible = val
                this.getList()
            },
            getList() {
                if (this.modelType !== 2) {
                    getNoPublicList({ pageIndex: this.pageIndex, pageSize: this.pageSize, modelId: this.modelId }).then(response => {
                        if (response.success) {
                            if (response.data.modelVersions !== null) {
                                this.total = response.data.totalSize
                                const data = response.data.modelVersions
                                this.tableData = []
                                data.forEach(item => {
                                    this.tableData.push({ ...item.versionDetail, isShared: item.isShared })
                                })
                            }
                        } else {
                            this.$message({
                                message: this.getErrorMsg(response.error.subcode),
                                type: 'warning'
                            });
                        }
                    })
                } else {
                    getPublicList({ pageIndex: this.pageIndex, pageSize: this.pageSize, modelId: this.modelId }).then(response => {
                        if (response.success) {
                            if (response.data.modelVersions !== null) {
                                this.total = response.data.totalSize
                                this.tableData = response.data.modelVersions
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
            // 版本列表状态
            fileStatus(val) {
                switch (val) {
                    case 0:
                        return '未上传'
                        break;
                    case 1:
                        return '上传中'
                        break;
                    case 2:
                        return '已上传'
                        break;
                    default:
                        return '上传失败'
                }
            },
            // 分享取消分享
            open(val) {
                let message = ''
                let flag = true
                if (!val.isShared) {
                    message = '分享至群组，群组内所有成员可见'
                    flag = true
                } else {
                    message = '取消分享'
                    flag = false
                }
                var data = { modelId: val.modelId, version: val.version }
                this.$confirm(message, '提示', {
                    confirmButtonText: '确定',
                    cancelButtonText: '取消',
                    type: 'warning'
                }).then(() => {
                    this.share(data, flag)
                }).catch(() => {
                    this.$message({
                        type: 'info',
                        message: '已取消操作'
                    });
                });
            },
            // 删除确认
            open2(val) {
                this.$confirm('此操作将永久删除该模型版本(如该模型版本已分享，则分享模型版本也会被删除)， 是否继续?', '提示', {
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

            share(data, flag) {
                if (flag) {
                    shareModel(data).then(response => {
                        if (response.success) {
                            this.$message({
                                message: '分享成功',
                                type: 'success'
                            });
                            this.getList()
                        } else {
                            this.$message({
                                message: this.getErrorMsg(response.error.subcode),
                                type: 'warning'
                            });
                        }
                    })
                } else {
                    cancelShareModel(data).then(response => {
                        if (response.success) {
                            this.$message({
                                message: '取消分享成功',
                                type: 'success'
                            });
                            this.getList()
                        } else {
                            this.$message({
                                message: this.getErrorMsg(response.error.subcode),
                                type: 'warning'
                            });
                        }
                    })
                }
            },
            // 预制模型部署
            deploy(val) {
                val = Object.assign(val, { modelName: this.modelName })
                this.$router.push({ name: 'modelDeploy', params: { data: val, flag: true } })
            }
        }

    }
</script>
<style lang="scss" scoped>
    .block {
        float: right;
        margin: 20px;
    }
</style>