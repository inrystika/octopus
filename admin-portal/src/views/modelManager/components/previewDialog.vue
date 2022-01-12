<template>
    <div>
        <el-dialog title="预览" :visible.sync="dialogTableVisible" :before-close="handleDialogClose"
            :close-on-click-modal="false">
            <el-table :data="tableData" height="300">
                <el-table-column property="name" label="模型名称" />
                <el-table-column property="contentType" label="内容类型" />
                <el-table-column property="size" label="模型大小" />
                <el-table-column label="最后修改时间">
                    <template slot-scope="scope">
                        <span style="margin-left: 10px">{{ scope.row.lastModified | parseTime }}</span>
                    </template>
                </el-table-column>
            </el-table>
        </el-dialog>
    </div>
</template>

<script>
    import { preview } from '@/api/modelManager.js'
    export default {
        name: "PreviewDialog",
        props: {
            row: { type: Object, default: () => { } }
        },
        data() {
            return {
                centerDialogVisible: true,
                tableData: undefined,
                dialogTableVisible: true,
                data: undefined
            }
        },
        created() {
            this.data = this.row
            this.getPreList()
        },
        beforeDestroy() {

        },
        methods: {
            getPreList() {
                preview(this.row).then(response => {
                    if (response.success) {
                        this.tableData = response.data.modelInfoList
                    } else {
                        this.$message({
                            message: this.getErrorMsg(response.error.subcode),
                            type: 'warning'
                        });
                    }
                })
            },
            handleDialogClose() {
                this.$emit('close', false)
            }
        }
    }
</script>
<style lang="scss" scoped>
    .el-dialog--center .el-dialog__body {
        text-align: center;
    }
</style>