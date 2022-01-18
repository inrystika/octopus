<template>
    <div>
        <div>
            <el-input v-model="subTaskInfo" type="textarea" :readonly="true" :rows="20" />
        </div>
        <div class="block">
            <el-pagination :current-page="pageIndex" :page-sizes="[10, 20, 50, 80]" :page-size="pageSize"
                layout="total, sizes, prev, pager, next, jumper" :total="total" @size-change="handleSizeChange"
                @current-change="handleCurrentChange" />
        </div>

        <div slot="footer" class="dialog-footer" />
    </div>
</template>
<script>
    import { deployEvent } from "@/api/deployManager";
    export default {
        name: "deployEvent",
        props: {
            id: {
                type: String,
                default: ''
            }
        },
        data() {
            return {
                subTaskInfo: "",
                total: 0,
                pageIndex: 1,
                pageSize: 10,
            }
        },
        created() {
            this.deployEvent()
        },
        methods: {
            deployEvent() {
                const param = {
                    id: this.id,
                    pageIndex: this.pageIndex,
                    pageSize: this.pageSize
                }
                deployEvent(param).then(response => {
                    if (response.success) {
                        this.total = response.payload.totalSize
                        let infoMessage = ""
                        response.data.depEvents.forEach(function (element) {
                            const title = element.reason
                            const message = element.message
                            infoMessage += "\n" + "[" + title + "]"
                            infoMessage += "\n" + "[" + message + "]" + "\n"
                        })
                        this.subTaskInfo = infoMessage
                    } else {
                        this.subTaskInfo = "暂无相关运行信息"
                    }
                }).catch(err => {
                    console.log("err:", err)
                    this.$message({
                        message: "未知错误",
                        type: 'warning'
                    });
                });
            },
            handleDialogClose() {
                this.$emit('close', false)
            },
            handleSizeChange(val) {
                this.pageSize = val
                this.deployEvent()
            },
            handleCurrentChange(val) {
                this.pageIndex = val
                this.deployEvent()
            },
            selectLog() {
                this.deployEvent()
            }
        },
    }
</script>

<style lang="scss" scoped>
    .el-col {
        margin: 10px 0 20px 0;
        font-size: 15px;
        font-weight: 800;

        span {
            font-weight: 400;
            margin-left: 20px
        }
    }


    .block {
        float: right;
        margin: 20px;
    }
</style>