<template>
    <div>
        <div>
            <el-row>
                <el-col :span="12">
                    <div>名称:<span>{{ data.name }}</span></div>
                </el-col>
                <el-col :span="12">
                    <div>描述: <span>{{ data.desc }}</span></div>
                </el-col>
            </el-row>
            <el-row>
                <el-col :span="12">
                    <div>模型:<span>{{ data.modelName+ ":" + data.modelVersion }}</span></div>
                </el-col>
                <el-col :span="12">
                    <div>URL:<span>{{ data.serviceUrl }}</span></div>
                </el-col>
            </el-row>
            <el-row>
                <el-col :span="12">
                    <div>模型服务调用:</div>
                </el-col>
            </el-row>
            <el-row>
                <el-col :span="10">
                    <el-input type="textarea" placeholder="请输入内容" v-model="textarea" :rows="20" style="width: 100%">
                    </el-input>
                </el-col>
                <el-col :span="2" style="text-align: center;padding-top: 10%;">
                    <el-button round>生成</el-button>
                </el-col>
                <el-col :span="12">
                    <el-input type="textarea" v-model="textarea" :rows="20" style="width: 80%">
                    </el-input>
                </el-col>
            </el-row>
        </div>
        <div>
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
                textarea: '',
                statusText: { 'preparing': ['status-ready', '初始中'], 'pending': ['status-agent', '等待中'], 'running': ['status-running', '运行中'], 'failed': ['status-danger', '失败'], 'succeeded': ['status-success', '成功'], 'stopped': ['status-stopping', '已停止'] }
            }
        },
        created() {
            this.data = this.row
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
</style>