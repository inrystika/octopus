<template>
    <div>
        <el-dialog 
            :title="'版本列表/' + modelName" 
            width="70%" 
            :visible.sync="CreateFormVisible"
            :before-close="handleDialogClose" 
            :close-on-click-modal="false"
        >
            <el-table 
              :data="tableData" 
              style="width: 100%" 
              height="500"
            >
                <el-table-column prop="version" label="算法版本" align="center">
                </el-table-column>
                <el-table-column prop="descript" label="模型描述" align="center">
                </el-table-column>
                <el-table-column label="状态" align="center">
                    <template slot-scope="scope">
                        <span style="margin-left: 10px">{{ fileStatus(scope.row.fileStatus) }}</span>
                    </template>
                </el-table-column>
                <el-table-column label="创建时间" align="center">
                    <template slot-scope="scope">
                        <span style="margin-left: 10px">{{ parseTime(scope.row.createdAt) }}</span>
                    </template>
                </el-table-column>
                <el-table-column label="操作" align="center" width="350px">
                    <template slot-scope="scope">
                        <el-button type="text" @click="handlePreview(scope.row)" :disabled="scope.row.fileStatus!==2">预览
                        </el-button>
                        <el-button v-if="Type===3" type="text" @click="open(scope.row)">删除</el-button>
                        <el-button type="text" @click="handledDownload(scope.row)" :disabled="scope.row.fileStatus!==2">
                            下载</el-button>
                    </template>
                </el-table-column>
            </el-table>
            <div class="block">
                <el-pagination 
                  :current-page="pageIndex" 
                  :page-sizes="[10, 20, 50, 80]" 
                  :page-size="pageSize"
                  :total="total"
                  layout="total, sizes, prev, pager, next, jumper" 
                  @size-change="handleSizeChange" 
                  @current-change="handleCurrentChange"
                />
            </div>
            <div slot="footer">
            </div>
        </el-dialog>
        <!-- 预览对话框 -->
        <previewDialog v-if="preVisible" @close="closeShareDialog" :row="data"></previewDialog>
    </div>
</template>

<script>

    import previewDialog from './previewDialog.vue'
    import { getModelList, downloadModel, deletePreModelVersion } from '@/api/modelManager.js'
    import { parseTime } from '@/utils/index'
    import { getErrorMsg } from '@/error/index'
    export default {
        name: "versionList",
        components: {

            previewDialog
        },
        props: {
            modelId: {
                type: String,
            },
            Type: { type: Number, default: undefined },
            modelName: { type: String }
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
                data: { modelId: undefined, version: undefined }

            }
        },
        created() {
            this.getList()
        },
        beforeDestroy() {

        },
        methods: {
            // 错误码
            getErrorMsg(code) {
                return getErrorMsg(code)
            },
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
                let data = JSON.parse(JSON.stringify(row));
                data.version = row.version
                data.modelId = row.modelId
                deletePreModelVersion(data).then(response => {
                    if (response.success) {
                        this.$message({
                            message: '删除成功',
                            type: 'success'
                        });
                        this.getList()
                    }
                    else {
                        this.$message({
                            message: this.getErrorMsg(response.error.subcode),
                            type: 'warning'
                        });
                    }

                })
            },
            handledDownload(row) {
                let data = {}
                data.version = row.version
                data.modelId = row.modelId
                data.domain = this.GLOBAL.DOMAIN
                downloadModel(data).then(response => {
                    if (response.success) { this.URLdownload(data, response.data.downloadUrl) }
                    else {
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
                this.tableData = []
                getModelList({ pageIndex: this.pageIndex, pageSize: this.pageSize, modelId: this.modelId }).then(response => {
                    if (response.success) {
                        this.total = response.data.totalSize
                        this.tableData = response.data.modelVersions
                        this.$message({
                            message: "获取列表成功",
                            type: 'success'
                        });
                    }
                    else {
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
            }
            ,//时间戳转换日期
            parseTime(val) {
                return parseTime(val)
            },
            // 版本列表状态
            fileStatus(val) {
                switch (val) {
                    case 0:
                        return '初始态'
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
            }
        },

    }
</script>
<style lang="scss" scoped>
</style>