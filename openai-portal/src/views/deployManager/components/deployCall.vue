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
            <!-- <el-row>
                <el-col :span="12">
                    <div>模型服务调用:</div>
                </el-col>
            </el-row>
            <el-row>
                <el-col :span="10">
                    <el-input type="textarea" placeholder="请输入内容" v-model="textareaInput" :rows="20"
                        style="width: 100%">
                    </el-input>
                </el-col>
                <el-col :span="2" style="text-align: center;padding-top: 10%;">
                    <el-button round @click="generate">生成</el-button>
                </el-col>
                <el-col :span="12">
                    <el-input type="textarea" v-model="textareaOutput" :rows="20" style="width: 80%" disabled>
                    </el-input>
                </el-col>
            </el-row> -->
        </div>
        <div>
        </div>
    </div>
</template>

<script>
    import { startDeploy } from "@/api/deployManager";
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
                textareaInput: '',
                textareaOutput: '',
                statusText: { 'Preparing': ['status-ready', '初始中'], 'Pending': ['status-agent', '等待中'], 'Running': ['status-running', '运行中'], 'Failed': ['status-danger', '失败'], 'Succeeded': ['status-success', '成功'], 'Stopped': ['status-stopping', '已停止'] }
            }
        },
        created() {
            this.data = this.row
        },
        methods: {
            generate() {
                let data = {}
                data.serviceUrl = this.data.serviceUrl
                data.data = this.textareaInput
                startDeploy(data).then(
                    response => { }
                )
            }
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