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
            </el-row>
            <el-row>
                <el-col :span="24">
                    <div>推理路径:<span>{{ data.serviceUrl }}</span></div>
                </el-col>
            </el-row>
            <el-row>
                <el-col :span="24">
                    <div>推理路径使用说明:<a href="https://octopus.openi.org.cn/docs/manual/infer" class="text">参考详细文档</a>
                    </div>
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
                statusText: { 'Preparing': ['status-ready', '初始中'], 'Creating': ['status-agent', '创建中'], 'Available': ['status-running', '运行中'], 'Failed': ['status-danger', '失败'], 'Stopped': ['status-stopping', '已停止'] },
            }
        },
        created() {
            this.data = this.row
            this.data.swaggerURL = this.data.serviceUrl.replace("predictions", "doc/")
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

    .text {
        font-weight: 400;
        margin-left: 10px;
        color:#3296fa;
        text-decoration: underline;
       
    }
</style>