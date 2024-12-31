<template>
    <div>
        <el-row>
            <el-col :span="12">
                <div>
                    任务名称:
                    <span>{{ profileInfo.name }}</span>
                </div>
            </el-col>
            <el-col :span="12">
                <div>
                    描述:
                    <span>{{ profileInfo.desc }}</span>
                </div>
            </el-col>
        </el-row>
        <el-row>
            <el-col :span="12">
                <div>
                    选用算法:
                    <span>{{
                        profileInfo.algorithmName +
                        ":" +
                        profileInfo.algorithmVersion
                    }}</span>
                </div>
            </el-col>
            <el-col :span="12">
                <div>
                    镜像选择:
                    <span>{{
                        profileInfo.imageName + ":" + profileInfo.imageVersion
                    }}</span>
                </div>
            </el-col>
        </el-row>
        <el-row>
            <el-col :span="12">
                <div>
                    选用数据集:
                    <span>{{ profileInfo.datasetShow }}</span>
                </div>
            </el-col>
            <el-col :span="12">
                <div>
                    选用模型:
                    <span>{{ profileInfo.modelShow }}</span>
                </div>
            </el-col>
        </el-row>
        <el-row>
            <el-col :span="12">
                <div>
                    是否分布式:
                    <span>{{
                        this.profileInfo.tasks.length > 1 ? "是" : "否"
                    }}</span>
                </div>
            </el-col>
            <el-col :span="12">
                <div>
                    资源规格:
                    <span>{{ profileInfo.resourceSpecName }}</span>
                </div>
            </el-col>
        </el-row>
        <el-row>
            <el-col :span="12">
                <div>
                    任务状态:
                    <span>{{ statusText[profileInfo.status][1] }}</span>
                </div>
            </el-col>
            <el-col :span="12">
                <div>
                    自定义启动命令:
                    <span>{{ profileInfo.command }}</span>
                </div>
            </el-col>
        </el-row>
        <el-row>
            <el-col :span="12">
                <div>
                    自动停止:
                    <span>{{
                        profileInfo.autoStopDuration == -1
                            ? "任务不会自动停止"
                            : profileInfo.autoStopDuration / 3600 + "小时"
                    }}</span>
                </div>
            </el-col>
            <el-col :span="12">
                <div>
                    <el-row :gutter="20">
                        <el-col :span="3"><div>访问配置</div></el-col>
                        <el-col :span="21">
                            <el-row :gutter="20">
                                <el-col :span="24" style="margin-top: 0px;text-align:center"
                                    ><div>
                                        <div>
                                            <el-col
                                                :span="6"
                                                style="margin: 0px"
                                                ><div>任务名称</div></el-col
                                            >
                                            <el-col
                                                :span="6"
                                                style="margin: 0px"
                                                ><div>访问路径</div></el-col
                                            >
                                            <el-col
                                                :span="6"
                                                style="margin: 0px"
                                                ><div>容器端口</div></el-col
                                            >
                                        </div>
                                    </div></el-col
                                >
                            </el-row>
                            <el-row :gutter="20">
                                <el-col :span="24" style="text-align:center"
                                    ><div>
                                        <div v-if="endpoint1 !== ''">
                                            <el-col
                                                :span="6"
                                                style="margin: 0px;font-weight:400"
                                                ><div>task0</div></el-col
                                            >
                                            <el-col
                                                :span="6"
                                                style="margin: 0px;font-weight:400"
                                                ><div>
                                                    /userendpoint/{{
                                                        endpoint1
                                                    }}
                                                </div></el-col
                                            >
                                            <el-col
                                                :span="6"
                                                style="margin: 0px;font-weight:400"
                                                ><div>{{ port1 }}</div></el-col
                                            >
                                        </div>
                                    </div></el-col
                                >
                            </el-row>
                            <el-row :gutter="20">
                                <el-col :span="24" style="text-align:center"
                                    ><div>
                                        <div v-if="endpoint2 !== ''">
                                            <el-col
                                                :span="6"
                                                style="margin: 0px;font-weight:400"
                                                ><div>task0</div></el-col
                                            >
                                            <el-col
                                                :span="6"
                                                style="margin: 0px;font-weight:400"
                                                ><div>
                                                    /userendpoint/{{
                                                        endpoint2
                                                    }}
                                                </div></el-col
                                            >
                                            <el-col
                                                :span="6"
                                                style="margin: 0px;font-weight:400"
                                                ><div>{{ port2 }}</div></el-col
                                            >
                                        </div>
                                    </div></el-col
                                >
                            </el-row>
                        </el-col>
                    </el-row>
                </div>
            </el-col>
        </el-row>
    </div>
</template>
<script>
export default {
    name: "NotebookProfile",
    props: {
        notebookData: {
            type: Object,
            default: () => {},
        },
    },
    data() {
        return {
            profileInfo: {},
            statusText: {
                preparing: ["status-ready", "初始中"],
                pending: ["status-agent", "等待中"],
                running: ["status-running", "运行中"],
                failed: ["status-danger", "失败"],
                succeeded: ["status-success", "成功"],
                stopped: ["status-stopping", "已停止"],
            },
            endpoint1: "",
            endpoint2: "",
            port1: "",
            port2: "",
        };
    },
    created() {
        this.profileInfo = this.notebookData;
        this.profileInfo.datasetShow = !this.profileInfo.datasetName
            ? ""
            : this.profileInfo.datasetName +
              ":" +
              this.profileInfo.datasetVersion;
        this.profileInfo.modelShow = !this.profileInfo.preTrainModelName
            ? ""
            : this.profileInfo.preTrainModelName +
              ":" +
              this.profileInfo.preTrainModelVersion;
        if (this.notebookData.tasks[0].endpoints != null) {
            this.endpoint1 = this.notebookData.tasks[0].endpoints[0].endpoint;
            this.port1 = this.notebookData.tasks[0].endpoints[0].port;
        }
        if (
            this.notebookData.tasks[1] &&
            this.notebookData.tasks[1].endpoints != null
        ) {
            this.endpoint2 = this.notebookData.tasks[1].endpoints[0].endpoint;
            this.port2 = this.notebookData.tasks[1].endpoints[0].port;
        }
    },
};
</script>
<style lang="scss" scoped>
.el-col {
    margin: 10px 0 20px 0;
    font-size: 15px;
    font-weight: 800;

    span {
        font-weight: 400;
        margin-left: 20px;
    }
}

// .taskList {
//   font-weight: 800;
// }

.block {
    float: right;
    margin: 20px;
}
.wrapper {
    // margin-top: 10px;
    // text-align: left;
    // // display: inline-block;
    // margin-left: 80px;
    // position: relative;
    // top: -27px;
}
</style>
