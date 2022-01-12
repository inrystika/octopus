<template>
    <div>
        <el-dialog :title="'版本列表/' + modelName" width="70%" :visible.sync="CreateFormVisible"
            :before-close="handleDialogClose" :close-on-click-modal="false">
            <el-table :data="tableData" style="width: 100%" height="350">
                <el-table-column label="算法版本">
                    <template slot-scope="scope">
                        <span>{{ scope.row.version }}</span>
                    </template>
                </el-table-column>
                <el-table-column label="模型描述" :show-overflow-tooltip="true">
                    <template slot-scope="scope">
                        <span>{{ scope.row.descript }}</span>
                    </template>
                </el-table-column>
                <el-table-column label="创建时间">
                    <template slot-scope="scope">
                        {{ scope.row.createdAt | parseTime }}
                    </template>
                </el-table-column>

                <el-table-column label="状态" align="center">
                    <template slot-scope="scope">
                        <span v-if="!(scope.row.progress&&scope.row.progress!=0)"> {{
                            fileStatus(scope.row.fileStatus)}}</span>
                        <span v-if="scope.row.progress&&scope.row.progress!=0">{{ "上传中" }}</span>
                        <el-progress :percentage="parseInt(scope.row.progress-1)"
                            v-if="scope.row.progress&&scope.row.progress!=0"></el-progress>
                    </template>
                </el-table-column>
                <el-table-column label="操作">
                    <template slot-scope="scope">
                        <el-button v-if="scope.row.fileStatus==0||scope.row.fileStatus==3&&modelType==3" type="text"
                            @click="handleEdit(scope.row)" :disabled="scope.row.progress&&scope.row.progress!=0">重新上传
                        </el-button>
                        <el-button type="text" :disabled="scope.row.fileStatus!==2" @click="handlePreview(scope.row)">
                            预览
                        </el-button>
                        <el-button v-if="modelType===3" type="text" @click="open(scope.row)"
                            :disabled="scope.row.progress&&scope.row.progress!=0">删除</el-button>
                        <el-button type="text" :disabled="scope.row.fileStatus!==2" @click="handledDownload(scope.row)">
                            下载
                        </el-button>
                    </template>
                </el-table-column>
            </el-table>
            <div class="block">
                <el-pagination :current-page="pageIndex" :page-sizes="[10, 20, 50, 80]" :page-size="pageSize"
                    :total="total" layout="total, sizes, prev, pager, next, jumper" @size-change="handleSizeChange"
                    @current-change="handleCurrentChange" />
            </div>
            <div slot="footer">
            </div>
        </el-dialog>
        <!-- 预览对话框 -->
        <previewDialog v-if="preVisible" :row="data" @close="closeShareDialog" />
        <!-- 创建对话框 -->
        <reupload v-if="CreateVisible" :row="row" :is-list="isList" @close="close" @cancel="cancel"
            @confirm="confirm" />
    </div>
</template>

<script>

    import previewDialog from './previewDialog.vue'
    import { getModelList, downloadModel, deletePreModelVersion } from '@/api/modelManager.js'
    import reupload from './reupload.vue'
    import store from '@/store'
    export default {
        name: "VersionList",
        components: {
            reupload,
            previewDialog
        },
        props: {
            modelId: { type: String, default: "" },
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
                data: { modelId: undefined, version: undefined },
                timer: null,
                CreateVisible: false,
                isList: true

            }
        },
        created() {
            this.getList()
            this.timer = setInterval(() => { this.getList() }, 2000)

        },
        destroyed() {
            clearInterval(this.timer)
            this.timer = null
        },
        methods: {
            handlePreview(row) {
                this.preVisible = true
                this.data.modelId = row.modelId
                this.data.version = row.version
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
                deletePreModelVersion(data).then(response => {
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
                data.domain = this.GLOBAL.DOMAIN
                downloadModel(data).then(response => {
                    if (response.success) {
                        this.URLdownload(data, response.data.downloadUrl)
                    } else {
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
                link.setAttribute('download', data.modelName)
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
            },
            cancelShareDialog(val) {
                this.dialogVisible = val
            },
            confirmShareDialog(val) {
                this.dialogVisible = val
            },
            getList() {
                getModelList({ pageIndex: this.pageIndex, pageSize: this.pageSize, modelId: this.modelId }).then(response => {
                    if (response.success) {
                        this.tableData = response.data.modelVersions
                        this.tableData.forEach(item => {
                            if (sessionStorage.getItem(JSON.stringify(item.modelId + item.version))) {
                                item.progress = sessionStorage.getItem(JSON.stringify(item.modelId + item.version))
                            }

                        })
                        this.total = response.data.totalSize
                    } else {
                        this.$message({
                            message: this.getErrorMsg(response.error.subcode),
                            type: 'warning'
                        });
                    }
                })
            },
            // 删除确认
            open(val) {
                this.$confirm('此操作将永久删除该预置模型版本, 是否继续?', '提示', {
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
            // 版本列表状态
            fileStatus(val) {
                switch (val) {
                    case 0:
                        return '未上传'
                        break;
                    case 1:
                        return '制作中'
                        break;
                    case 2:
                        return '制作完成'
                        break;
                    case 3: '制作失败'
                    default:
                        return '上传失败'
                }
            },
            // 重新上传
            handleEdit(val) {
                this.isList = true
                this.CreateVisible = true
                this.row = val
                store.commit('user/SET_PROGRESSID', val.modelId)
            },
            cancel(val) {
                this.CreateVisible=val         
                this.getList()
            },
            confirm(val) {
                this.CreateVisible=val           
                this.getList()
            },
            close(val) {
                this.CreateVisible=val
                this.getList()
            },
        }

    }
</script>
<style lang="scss" scoped>
  .block {
    float: right;
    margin: 20px;
  }
</style>