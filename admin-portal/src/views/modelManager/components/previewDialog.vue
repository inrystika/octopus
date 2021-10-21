<template>
    <div>
        <el-dialog
            title="预览"
            :visible.sync="dialogTableVisible"
            :before-close="handleDialogClose"
            :close-on-click-modal="false"
        >
            <el-table :data="tableData" height="300">
                <el-table-column property="name" label="模型名称" />
                <el-table-column property="contentType" label="内容类型" />
                <el-table-column property="size" label="模型大小" />
                <el-table-column label="最后修改时间">
                    <template slot-scope="scope">
                        <span style="margin-left: 10px">{{ parseTime(scope.row.lastModified) }}</span>
                    </template>
                </el-table-column>
            </el-table>
        </el-dialog>
    </div>
</template>

<script>
    import { preview } from '@/api/modelManager.js'
    import { parseTime } from '@/utils/index'
    import { getErrorMsg } from '@/error/index'
    export default {
        name: "PreviewDialog",
        props: {
            row: { type: Object, default: () => {} }
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
            // 错误码
            getErrorMsg(code) {
                return getErrorMsg(code)
            },
            getPreList() {
                preview(this.row).then(response => {
                    if (response.success) {
                        this.tableData = response.data.modelInfoList
                    } else {
                        this.$message({
                            message: '目前暂无数据',
                            type: 'success'
                        });
                    }
                })
            },
            handleDialogClose() {
                this.$emit('close', false)
            },
            // 时间戳转换日期
            parseTime(val) {
                return parseTime(val)
            }

        }
    }
</script>
<style lang="scss" scoped>
    .el-dialog--center .el-dialog__body {
        text-align: center;
    }
</style>