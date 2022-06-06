<template>
    <div>
        <div>
            <el-row>
                <el-col :span="12">
                    <div>服务名称:<span>{{ data.name }}</span></div>
                </el-col>
                <el-col :span="12">
                    <div>描述: <span>{{ data.desc }}</span></div>
                </el-col>
            </el-row>
            <el-row>
                <el-col :span="12">
                    <div>模型名称:<span>{{ data.modelName+ ":" + data.modelVersion }}</span></div>
                </el-col>
            </el-row>
            <el-row>
                <el-col :span="24">
                    <div>推理路径:<span>{{ data.serviceUrl }}</span></div>
                </el-col>
            </el-row>
            <el-row>
                <el-col :span="24">
                    <div>swagger推理路径: <a href="https://octopus.openi.org.cn/docs/manual/infer" target="_blank"
                            class="text">管理手册</a></div>
                </el-col>
            </el-row>
            <el-row>
                <el-col :span="12">
                    <div>任务状态:<span style="margin-left: 10px">{{ statusText[data.status][1] }}</span></div>
                </el-col>
                <el-col :span="12">
                    <div></div>
                </el-col>
            </el-row>
        </div>
    </div>
</template>

<script>
    export default {
        name: "deployCall",
        props: {
            row: {
                type: Object,
                default: () => { }
            }
        },
        data() {
            return {
                data: {},
                statusText: { 'Preparing': ['status-ready', '初始中'], 'Creating': ['status-agent', '创建中'], 'Available': ['status-running', '运行中'], 'Failed': ['status-danger', '失败'], 'Stopped': ['status-stopping', '已停止'] }
            }
        },
        created() {
            this.data = JSON.parse(JSON.stringify(this.row))
            this.data.swaggerURL = this.data.serviceUrl.replace("predictions", "doc/")
        },
        methods: {
        }
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

    .taskList {
        font-weight: 800;
    }

    .block {
        float: right;
        margin: 20px;
    }

    .text {
        font-weight: 400;
        margin-left: 10px;
        color: #3296fa;
        text-decoration: underline;

    }
</style>