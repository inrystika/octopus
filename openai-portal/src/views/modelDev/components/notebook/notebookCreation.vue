<template>
    <div>
        <el-dialog
            title="创建NoteBook"
            width="55%"
            :visible.sync="CreateFormVisible"
            :before-close="handleDialogClose"
            :close-on-click-modal="false"
              custom-class="dialog"
           
        >
            <el-form
                ref="ruleForm"
                :model="ruleForm"
                :rules="rules"
                :label-width="formLabelWidth"
                class="demo-ruleForm"
            >
                <el-form-item
                    label="名称"
                    :label-width="formLabelWidth"
                    prop="name"
                    class="name"
                >
                    <el-input
                        v-model="ruleForm.name"
                        placeholder="请输入NoteBook名称"
                    />
                </el-form-item>
                <div class="tip">
                    <i class="el-alert__icon el-icon-warning"></i>
                    算法存储在<span>/code</span>中，
                    数据集存储在<span>/dataset</span>中，
                    用户目录在<span>/userhome</span>中，
                    选择的模型存储在<span>/pretrainmodel</span>中
                </div>
                <el-form-item
                    label="描述"
                    :label-width="formLabelWidth"
                    prop="desc"
                >
                    <el-input
                        v-model="ruleForm.desc"
                        :autosize="{ minRows: 2, maxRows: 4 }"
                        placeholder="请输入NoteBook描述"
                        maxlength="300"
                        show-word-limit
                    />
                </el-form-item>
                <!-- 算法三级框 -->
                <div>
                    <el-form-item
                        label="算法名称"
                        prop="algorithmId"
                        style="display: inline-block"
                    >
                        <el-select
                            v-model="ruleForm.algorithmId"
                            v-loadmore="loadAlgorithmName"
                            placeholder="请选择算法名称"
                            filterable
                            remote
                            :remote-method="remoteAlgorithm"
                            @change="changeAlgorithmName"
                            @click.native="getAlgorithmItem"
                        >
                            <el-option
                                v-for="item in algorithmNameOption"
                                :key="item.algorithmId + item.algorithmName"
                                :label="item.algorithmName"
                                :value="item.algorithmId"
                            />
                        </el-select>
                        <el-tooltip
                            class="tipClass"
                            effect="dark"
                            :content="tipMessage"
                            placement="top-start"
                        >
                            <i class="el-icon-warning-outline"></i>
                        </el-tooltip>
                    </el-form-item>
                    <el-form-item
                        v-if="algorithmVersion"
                        label="算法版本"
                        prop="algorithmVersion"
                        style="display: inline-block"
                    >
                        <el-select
                            v-model="ruleForm.algorithmVersion"
                            v-loadmore="loadAlgorithmVersion"
                            placeholder="请选择算法版本"
                        >
                            <el-option
                                v-for="item in algorithmVersionOption"
                                :key="
                                    item.algorithmDetail.algorithmId +
                                    item.algorithmDetail.algorithmVersion
                                "
                                :label="item.algorithmDetail.algorithmVersion"
                                :value="item.algorithmDetail.algorithmVersion"
                            />
                        </el-select>
                    </el-form-item>
                </div>
                <!-- 镜像三级框 -->
                <div>
                    <el-form-item
                        label="镜像类型"
                        prop="imageSource"
                        :class="{ inline: imageName }"
                    >
                        <el-select
                            v-model="ruleForm.imageSource"
                            placeholder="请选择"
                            @change="changeimageSource"
                        >
                            <el-option label="我的镜像" value="my" />
                            <el-option label="预置镜像" value="pre" />
                            <el-option label="公共镜像" value="common" />
                        </el-select>
                    </el-form-item>
                    <el-form-item
                        v-if="imageName"
                        label="镜像名称"
                        prop="imageItem"
                        style="display: inline-block"
                    >
                        <el-select
                            v-model="ruleForm.imageItem"
                            v-loadmore="loadImageName"
                            value-key="id"
                            placeholder="请选择镜像"
                            filterable
                            remote
                            :remote-method="remoteImage"
                            @click.native="getImageItem"
                        >
                            <el-option
                                v-for="item in imageNameOption"
                                :key="item.id"
                                :label="
                                    item.imageName + ':' + item.imageVersion
                                "
                                :value="item"
                            />
                        </el-select>
                    </el-form-item>
                </div>
                <!-- 数据集三级框 -->
                <div>
                    <el-form-item
                        label="数据集类型"
                        prop="dataSetSource"
                        :class="{ inline: dataSetName }"
                    >
                        <el-select
                            v-model="ruleForm.dataSetSource"
                            clearable
                            placeholder="请选择"
                            @clear="clearDataSetVersionOption"
                            @change="changedataSetSource"
                        >
                            <el-option label="我的数据集" value="my" />
                            <el-option label="预置数据集" value="pre" />
                            <el-option label="公共数据集" value="common" />
                        </el-select>
                    </el-form-item>
                    <el-form-item
                        v-if="dataSetName"
                        label="数据集名称"
                        prop="dataSetId"
                        style="display: inline-block"
                    >
                        <el-select
                            v-model="ruleForm.dataSetId"
                            v-loadmore="loadDataSetName"
                            placeholder="请选择数据集名称"
                            filterable
                            remote
                            :remote-method="remoteDataSet"
                            @change="changeDataSetName"
                            @click.native="getDataSetItem"
                        >
                            <el-option
                                v-for="item in dataSetNameOption"
                                :key="item.id + item.name"
                                :label="item.name"
                                :value="item.id"
                            />
                        </el-select>
                    </el-form-item>
                    <el-form-item
                        v-if="dataSetVersion"
                        label="数据集版本"
                        prop="dataSetVersion"
                        style="display: inline-block"
                    >
                        <el-select
                            v-model="ruleForm.dataSetVersion"
                            v-loadmore="loadDataSetVersion"
                            placeholder="请选择数据集版本"
                        >
                            <el-option
                                v-for="item in dataSetVersionOption"
                                :key="item.datasetId + item.version"
                                :label="item.version"
                                :value="item.version"
                            />
                        </el-select>
                    </el-form-item>
                </div>
                <!-- 模型三级框 -->
                <div>
                    <el-form-item
                        label="模型类型"
                        prop="modelSource"
                        :class="{ inline: modelName }"
                    >
                        <el-select
                            v-model="ruleForm.modelSource"
                            clearable
                            placeholder="请选择"
                            @clear="clearModelVersionOption"
                            @change="changeModelSource"
                        >
                            <el-option label="我的模型" value="myModel" />
                            <el-option label="公共模型" value="commonModel" />
                            <el-option label="预置模型" value="preModel" />
                        </el-select>
                    </el-form-item>
                    <el-form-item
                        v-if="modelName"
                        label="模型名称"
                        prop="modelId"
                        style="display: inline-block"
                    >
                        <el-select
                            ref="modelNameRef"
                            v-model="ruleForm.modelId"
                            v-loadmore="loadModelName"
                            placeholder="请选择模型名称"
                            filterable
                            remote
                            :remote-method="remoteModel"
                            @change="changeModelName"
                            @click.native="getModelItem"
                        >
                            <el-option
                                v-for="item in modelNameOption"
                                :key="item.modelId"
                                :label="item.modelName"
                                :value="item.modelId"
                            />
                        </el-select>
                    </el-form-item>
                    <el-form-item
                        v-if="modelVersion"
                        label="模型版本"
                        prop="modelVersion"
                        style="display: inline-block"
                    >
                        <el-select
                            v-model="ruleForm.modelVersion"
                            v-loadmore="loadModelVersion"
                            placeholder="请选择数据集版本"
                        >
                            <el-option
                                v-for="item in modelVersionOption"
                                :key="
                                    item.versionDetail.modelId +
                                    item.versionDetail.version
                                "
                                :label="item.versionDetail.version"
                                :value="item.versionDetail.version"
                            />
                        </el-select>
                    </el-form-item>
                </div>
                <!-- 资源二级框 -->
                <div>
                    <el-form-item
                        label="资源池"
                        prop="resourcePool"
                        style="display: inline-block"
                    >
                        <el-select
                            v-model="ruleForm.resourcePool"
                            placeholder="请选择资源池"
                            @change="getResource"
                        >
                            <el-option
                                v-for="(item, index) in poolList"
                                :key="index"
                                :label="item"
                                :value="item"
                            />
                        </el-select>
                    </el-form-item>
                    <el-form-item
                        v-if="specificationVisible"
                        label="资源规格"
                        prop="specification"
                        style="display: inline-block"
                    >
                        <el-select
                            v-model="ruleForm.specification"
                            placeholder="请选择资源规格"
                        >
                            <el-option
                                v-for="(item, index) in resourceList"
                                :key="index"
                                :label="item.label"
                                :value="item.value"
                            />
                        </el-select>
                    </el-form-item>
                </div>
                <el-form-item label="自动停止" prop="autoStopDuration">
                    <el-switch
                        v-model="isAutoStop"
                        active-color="#3296fa"
                        inactive-color="gray"
                    >
                    </el-switch>
                    <span v-if="isAutoStop" style="margin-left: 10px">
                        <el-input-number
                            v-model="ruleForm.autoStopDuration"
                            :precision="1"
                            :step="1"
                            :min="1"
                        ></el-input-number>
                        <span style="color: #b3b3b3; margin-left: 10px"
                            >小时</span
                        >
                    </span>
                </el-form-item>
                <el-form-item>
                    <el-button type="text" @click="showMultitask"
                        >高级设置</el-button
                    >
                </el-form-item>
                <div v-if="isShowMultitask">
                    <el-form-item label="任务数" prop="taskNumber">
                        <el-select
                            v-model.number="ruleForm.taskNumber"
                            placeholder="请选择"
                            @change="changeNum"
                        >
                            <el-option label="1" value="1" />
                            <el-option label="2" value="2" />
                        </el-select>
                    </el-form-item>
                    <el-form-item
                        label="访问配置"
                        style="margin-bottom: 10px; font-weight: 800"
                    >
                        <el-row :gutter="20">
                            <el-col :span="4"><div>任务名称</div></el-col>
                            <el-col :span="5"><div>访问路径</div></el-col>
                            <el-col :span="5"><div>容器端口</div></el-col>
                        </el-row>
                    </el-form-item>
                    <el-form-item label="">
                        <div
                            class="wrapper"
                            v-if="
                                ruleForm.taskNumber == 1 ||
                                ruleForm.taskNumber == 2
                            "
                        >
                            <el-row :gutter="20">
                                <el-col :span="3"
                                    ><div>
                                        <span style="padding-left: 20px"
                                            >task0</span
                                        >
                                    </div></el-col
                                >
                                <el-col :span="6"
                                    ><div>
                                        <span>/userendpoint/</span>
                                        <el-input
                                            v-model="endpoint1"
                                            placeholder="路径"
                                            style="width: 40%"
                                        ></el-input>
                                        <span>/</span>
                                    </div></el-col
                                >
                                <el-col :span="6">
                                    <span></span
                                    ><el-input
                                        v-model="port1"
                                        placeholder="端口"
                                        style="width: 40%"
                                        oninput="value=value.replace(/[^0-9.]/g,'')"
                                    ></el-input
                                ></el-col>
                            </el-row>
                        </div>
                        <div
                            class="wrapper"
                            v-if="
                                ruleForm.taskNumber == 2 &&
                                !(ruleForm.taskNumber == 1)
                            "
                        >
                            <el-row :gutter="20">
                                <el-col :span="3"
                                    ><div>
                                        <span style="padding-left: 20px"
                                            >task1</span
                                        >
                                    </div></el-col
                                >
                                <el-col :span="6"
                                    ><div>
                                        <span>/userendpoint/</span>
                                        <el-input
                                            v-model="endpoint2"
                                            placeholder="路径"
                                            style="width: 40%"
                                        ></el-input
                                        ><span>/</span>
                                    </div></el-col
                                >
                                <el-col :span="6">
                                    <el-input
                                        v-model="port2"
                                        placeholder="端口"
                                        style="width: 40%"
                                        oninput="value=value.replace(/[^0-9.]/g,'')"
                                    ></el-input
                                ></el-col>
                            </el-row>
                        </div>
                    </el-form-item>
                    <!-- <el-form-item label="自定义启动命令" prop="command">
                    <el-input v-model="ruleForm.command" type="textarea"></el-input>
                  </el-form-item>
                    <div class="tip">
                      <i class="el-alert__icon el-icon-warning"></i>服务端口环境变量为<span>OCTOPUS_NOTEBOOK_PORT</span>，基础URL环境变量为<span>OCTOPUS_NOTEBOOK_BASE_URL</span>
                  </div>                 -->
                </div>
            </el-form>
            <div slot="footer" class="dialog-footer">
                <el-button @click="cancel">取 消</el-button>
                <el-button
                    type="primary"
                    @click="submit('ruleForm')"
                    v-preventReClick
                    >确 定</el-button
                >
            </div>
        </el-dialog>
    </div>
</template>

<script>
import {
    createNotebook,
    getMyAlgorithmList,
    getAlgorithmVersionList,
} from "@/api/modelDev";
import {
    getMyDatasetList,
    getPublicDatasetList,
    getPresetDatasetList,
    getVersionList,
} from "@/api/datasetManager";
import { getMyImage, getPublicImage, getPreImage } from "@/api/imageManager";
import {
    getMyModel,
    getPreModel,
    getPublicModel,
    getPublicList,
    getNoPublicList,
} from "@/api/modelManager";
import { getResourceList } from "@/api/trainingManager";
import { mapGetters } from "vuex";
import { randomName } from "@/utils/index";
export default {
    name: "NotebookCreation",
    directives: {
        loadmore: {
            inserted: function (el, binding) {
                const SELECTWRAP_DOM = el.querySelector(
                    ".el-select-dropdown .el-select-dropdown__wrap"
                );
                SELECTWRAP_DOM.addEventListener("scroll", function () {
                    const CONDITION =
                        this.scrollHeight - this.scrollTop <= this.clientHeight;
                    binding.value();
                });
            },
        },
    },
    data() {
        var checkVersion = (idField, errorMessage) => {
            return (rule, value, callback) => {
                if (this.ruleForm[idField] && !value) {
                    callback(new Error(errorMessage));
                }
                return callback();
            };
        };
        return {
            isAutoStop: true,
            specificationVisible: false,
            poolList: [],
            endpoint1: "",
            endpoint2: "",
            port1: "",
            port2: "",
            ruleForm: {
                name: "",
                desc: "",
                algorithmSource: "my",
                algorithmId: "",
                algorithmVersion: "",
                imageSource: "",
                imageItem: "",
                dataSetSource: "",
                dataSetId: "",
                dataSetVersion: "",
                modelSource: "",
                modelId: "",
                modelVersion: "",
                taskNumber: 1,
                resourcePool: "",
                specification: "",
                command: "",
                autoStopDuration: 4,
            },
            rules: {
                name: [
                    {
                        required: true,
                        message: "请输入NoteBook名称",
                        trigger: "blur",
                    },
                ],
                algorithmSource: [
                    {
                        required: true,
                        message: "请选择算法类型",
                        trigger: "change",
                    },
                ],
                algorithmId: [
                    {
                        required: true,
                        message: "请选择算法名称",
                        trigger: "change",
                    },
                ],
                algorithmVersion: [
                    {
                        required: true,
                        message: "请选择算法版本",
                        trigger: "change",
                    },
                ],
                imageSource: [
                    {
                        required: true,
                        message: "请选择镜像类型",
                        trigger: "change",
                    },
                ],
                imageItem: [
                    {
                        required: true,
                        message: "请选择镜像名称",
                        trigger: "change",
                    },
                ],
                dataSetVersion: [
                    {
                        validator: checkVersion(
                            "dataSetId",
                            "请选择数据集版本"
                        ),
                        trigger: "blur",
                    },
                ],
                modelVersion: [
                    {
                        validator: checkVersion("modelId", "请选择模型版本"),
                        trigger: "blur",
                    },
                ],
                specification: [
                    {
                        required: true,
                        message: "请选择资源规格",
                        trigger: "blur",
                    },
                ],
                resourcePool: [
                    {
                        required: true,
                        message: "请选择资源池",
                        trigger: "blur",
                    },
                ],
            },
            CreateFormVisible: true,
            formLabelWidth: "120px",
            pageIndex: 1,
            pageSize: 20,
            isShowMultitask: false,
            // 算法三级框
            algorithmName: false,
            algorithmVersion: false,
            algorithmNameOption: [],
            algorithmVersionOption: [],
            algorithmNameCount: 1,
            algorithmVersionCount: 1,
            algorithmNameTotal: undefined,
            algorithmVersionTotal: undefined,
            // 镜像二级框
            imageName: false,
            imageNameOption: [],
            imageNameCount: 1,
            imageNameTotal: undefined,
            // 数据集三级框
            dataSetName: false,
            dataSetVersion: false,
            dataSetNameOption: [],
            dataSetVersionOption: [],
            dataSetNameCount: 1,
            dataSetVersionCount: 1,
            dataSetNameTotal: undefined,
            dataSetVersionTotal: undefined,
            // 模型三级框
            modelName: false,
            modelVersion: false,
            modelNameOption: [],
            modelVersionOption: [],
            modelNameCount: 1,
            modelVersionCount: 1,
            modelNameTotal: undefined,
            modelVersionTotal: undefined,
            data: {},
            resourceList: [],
            algorithmNameTemp: "",
            imageTemp: "",
            dataSetTemp: "",
            modelTemp: "",
            tipMessage:
                "创建Notebook任务时，算法只能从‘我的算法’中选择；请先在‘我的算法’中上传算法。",
        };
    },
    created() {
        // this.getResource();
        this.getSpacePools();
        this.ruleForm.name = this.randomName("notebook");
    },
    computed: {
        ...mapGetters(["workspaces"]),
    },
    methods: {
        clearModelVersionOption() {
            this.modelVersionOption = [];
        },
        changeModelSource() {
            this.modelName = true;
            this.modelNameCount = 1;
            this.modelNameOption = [];
            this.ruleForm.modelId = "";
            this.ruleForm.modelVersion = "";
            this.modelChange = true;
            this.getModelNameList();
        },
        changeModelName() {
            this.modelVersion = true;
            this.modelVersionCount = 1;
            this.modelVersionOption = [];
            this.ruleForm.modelVersion = "";
            this.getModelVersionList();
        },
        getModelNameList(searchKey) {
            if (this.ruleForm.modelSource === "myModel") {
                getMyModel({
                    pageIndex: this.modelNameCount,
                    pageSize: 10,
                    nameLike: searchKey,
                }).then((response) => {
                    if (response.data.models === null) {
                        response.data.models = [];
                    } else {
                        this.modelNameOption = this.modelNameOption.concat(
                            response.data.models
                        );
                        this.modelNameTotal = response.data.totalSize;
                    }
                });
            } else if (this.ruleForm.modelSource === "preModel") {
                getPreModel({
                    pageIndex: this.modelNameCount,
                    pageSize: 10,
                    nameLike: searchKey,
                }).then((response) => {
                    if (response.data.models !== null) {
                        this.modelNameOption = this.modelNameOption.concat(
                            response.data.models
                        );
                        this.modelNameTotal = response.data.totalSize;
                    } else {
                        response.data.models = [];
                    }
                });
            } else if (this.ruleForm.modelSource === "commonModel") {
                getPublicModel({
                    pageIndex: this.modelNameCount,
                    pageSize: 10,
                    nameLike: searchKey,
                }).then((response) => {
                    if (response.data.models !== null) {
                        this.modelNameOption = this.modelNameOption.concat(
                            response.data.models
                        );
                        this.modelNameTotal = response.data.totalSize;
                    } else {
                        response.data.models = [];
                    }
                });
            }
        },
        getModelVersionList() {
            const data = {};
            data.modelId = this.ruleForm.modelId;
            data.pageIndex = this.modelVersionCount;
            data.pageSize = 10;
            if (this.ruleForm.modelSource === "commonModel") {
                getPublicList(data).then((response) => {
                    if (response.data.modelVersions !== null) {
                        this.modelVersionOption =
                            this.modelVersionOption.concat(
                                response.data.modelVersions
                            );
                        this.modelVersionTotal = response.data.totalSize;
                    }
                });
            } else {
                getNoPublicList(data).then((response) => {
                    if (response.data.modelVersions !== null) {
                        this.modelVersionOption =
                            this.modelVersionOption.concat(
                                response.data.modelVersions
                            );
                        this.modelVersionTotal = response.data.totalSize;
                    }
                });
                console.log(
                    "this.modelVersionOption:",
                    this.modelVersionOption
                );
            }
        },
        loadModelName() {
            this.modelNameCount = this.modelNameCount + 1;
            if (this.modelNameOption.length < this.modelNameTotal) {
                this.getModelNameList(this.modelTemp);
            }
        },
        loadModelVersion() {
            this.modelVersionCount = this.modelVersionCount + 1;
            if (this.modelVersionOption.length < this.modelVersionTotal) {
                this.getModelVersionList();
            }
        },
        // 远程请求数据集名称
        remoteModel(searchKey) {
            if (searchKey === "") {
                this.modelTemp = "";
            } else {
                this.modelTemp = searchKey;
            }
            this.modelNameOption = [];
            this.modelNameCount = 1;
            this.getModelNameList(this.modelTemp);
        },
        getModelItem() {
            this.modelTemp = "";
            this.modelNameCount = 1;
            if (this.ruleForm.modelSource === "myModel") {
                getMyModel({
                    pageIndex: this.modelNameCount,
                    pageSize: 10,
                }).then((response) => {
                    if (response.data.models === null) {
                        response.data.models = [];
                    } else {
                        this.modelNameOption = response.data.models;
                        this.modelNameTotal = response.data.totalSize;
                    }
                });
            } else if (this.ruleForm.modelSource === "preModel") {
                getPreModel({
                    pageIndex: this.modelNameCount,
                    pageSize: 10,
                }).then((response) => {
                    if (response.data.models !== null) {
                        this.modelNameOption = response.data.models;
                        this.modeltNameTotal = response.data.totalSize;
                    } else {
                        response.data.models = [];
                    }
                });
            } else if (this.ruleForm.modelSource === "commonModel") {
                getPublicModel({
                    pageIndex: this.modelNameCount,
                    pageSize: 10,
                }).then((response) => {
                    if (response.data.models !== null) {
                        this.modelNameOption = response.data.models;
                        this.modelNameTotal = response.data.totalSize;
                    } else {
                        response.data.models = [];
                    }
                });
            }
        },
        randomName(val) {
            return randomName(val);
        },
        clearDataSetVersionOption() {
            this.dataSetVersionOption = [];
        },
        getSpacePools() {
            let workspaceName = JSON.parse(
                sessionStorage.getItem("space")
            ).workspaceName;
            this.workspaces.forEach((item) => {
                // 获取当前群组绑定资源池列表
                if (item.name == workspaceName) {
                    this.poolList = item.resourcePools;
                }
            });
        },
        handleDialogClose() {
            this.$emit("close", false);
        },
        cancel() {
            this.$confirm("此操作将被取消，是否继续?", "提示", {
                confirmButtonText: "确定",
                cancelButtonText: "取消",
                type: "warning",
            })
                .then(() => {
                    this.$emit("cancel", false);
                })
                .catch(() => {
                    this.$message({
                        type: "info",
                        message: "已中断取消操作",
                    });
                });
        },
        getResource() {
            this.specificationVisible = true;
            this.ruleForm.specification = "";
            this.resourceList = [];
            getResourceList(this.ruleForm.resourcePool).then((response) => {
                if (response.success) {
                    response.data.mapResourceSpecIdList.debug.resourceSpecs.forEach(
                        (item) => {
                            this.resourceList.push({
                                label: item.name + " " + item.price + "机时/h",
                                value: item.id,
                            });
                        }
                    );
                } else {
                    this.$message({
                        message: this.getErrorMsg(response.error.subcode),
                        type: "warning",
                    });
                }
            });
        },
        showMultitask() {
            this.isShowMultitask = !this.isShowMultitask;
        },
        submit(formName) {
            this.$refs[formName].validate((valid) => {
                if (
                    this.endpoint1 !== "" &&
                    this.endpoint2 !== "" &&
                    this.endpoint1 == this.endpoint2
                ) {
                    this.$message({
                        message: "路径不能相同",
                        type: "warning",
                    });
                    return;
                }
                // let taskConfigs = [
                //     {
                //         endpoints: [
                //             {
                //                 endpoint: this.endpoint1,
                //                 port:
                //                     this.port1 == "" ? 0 : parseInt(this.port1),
                //             },
                //         ],
                //     },
                //     {
                //         endpoints: [
                //             {
                //                 endpoint: this.endpoint2,
                //                 port:
                //                     this.port2 == "" ? 0 : parseInt(this.port2),
                //             },
                //         ],
                //     },
                // ];
                if (
                    (this.endpoint1 !== "" && this.port1 == "") ||
                    (this.port1 !== "" && this.endpoint1 == "")
                ) {
                    this.$message({
                        message: "路径和端口必须同时填写或者不同不填写",
                        type: "warning",
                    });
                    return;
                }
                if (
                    (this.endpoint2 !== "" && this.port2 == "") ||
                    (this.port2 !== "" && this.endpoint2 == "")
                ) {
                    this.$message({
                        message: "路径和端口必须同时填写或者不同不填写",
                        type: "warning",
                    });
                    return;
                }
                let flag = true;
                let data;
                if (this.endpoint1 == "" && this.endpoint2 == "") {
                    flag = false;
                } else {
                    flag = true;
                }
                if (this.endpoint1 != "" && this.endpoint2 == "") {
                    data = [
                        {
                            endpoints: [
                                {
                                    endpoint: this.endpoint1,
                                    port:
                                        this.port1 == ""
                                            ? 0
                                            : parseInt(this.port1),
                                },
                            ],
                        },
                        {},
                    ];
                }
                if (this.endpoint2 != "" && this.endpoint1 == "") {
                    data = [
                        {},
                        {
                            endpoints: [
                                {
                                    endpoint: this.endpoint2,
                                    port:
                                        this.port2 == ""
                                            ? 0
                                            : parseInt(this.port2),
                                },
                            ],
                        },
                    ];
                }
                if (this.endpoint2 != "" && this.endpoint1 != "") {
                    data = [
                        {
                            endpoints: [
                                {
                                    endpoint: this.endpoint1,
                                    port:
                                        this.port1 == ""
                                            ? 0
                                            : parseInt(this.port1),
                                },
                            ],
                        },
                        {
                            endpoints: [
                                {
                                    endpoint: this.endpoint2,
                                    port:
                                        this.port2 == ""
                                            ? 0
                                            : parseInt(this.port2),
                                },
                            ],
                        },
                    ];
                }
                // console.log(taskConfigs);
                if (valid) {
                    this.showUpload = true;
                    const param = {
                        name: this.ruleForm.name,
                        desc: this.ruleForm.desc,
                        imageId: this.ruleForm.imageItem.id,
                        imageVersion: this.ruleForm.imageItem.imageVersion,
                        resourceSpecId: this.ruleForm.specification,
                        algorithmId: this.ruleForm.algorithmId || "",
                        algorithmVersion: this.ruleForm.algorithmVersion || "",
                        datasetId: this.ruleForm.dataSetId || "",
                        datasetVersion: this.ruleForm.dataSetVersion || "",
                        preTrainModelId: this.ruleForm.modelId || "",
                        preTrainModelName:
                            this.$refs.modelNameRef?.selected.label || "",
                        preTrainModelVersion: this.ruleForm.modelVersion || "",
                        taskNumber: this.ruleForm.taskNumber,
                        resourcePool: this.ruleForm.resourcePool,
                        command: this.ruleForm.command,
                        autoStopDuration: this.ruleForm.autoStopDuration * 3600,
                        // taskConfigs: taskConfigs,
                    };
                    if (flag) {
                        param.taskConfigs = data;
                    }
                    if (!this.isAutoStop) {
                        param.autoStopDuration = -1;
                    }
                    const confirmInfo = this.$createElement;
                    this.$confirm("温馨提示", {
                        title: "温馨提示",
                        message: confirmInfo("div", [
                            confirmInfo(
                                "p",
                                "NoteBook 任务用于调试程序,使用 Jupyterlab 代码编辑器调试程序"
                            ),
                            confirmInfo("br", ""),
                            confirmInfo(
                                "p",
                                "调试代码保存路径默认为 Linux 系统的/code"
                            ),
                            confirmInfo("br", ""),
                            confirmInfo(
                                "p",
                                "NoteBook 任务达到管理员设置的运行时间后,会自动停止并释放资源"
                            ),
                        ]),
                    })
                        .then(() => {
                            createNotebook(param).then((response) => {
                                if (response.success) {
                                    this.$message.success("创建成功");
                                    this.$emit("confirm", false);
                                } else {
                                    this.$message({
                                        message: this.getErrorMsg(
                                            response.error.subcode
                                        ),
                                        type: "warning",
                                    });
                                }
                            });
                        })
                        .catch(() => {
                            this.$message({
                                type: "info",
                                message: "已取消",
                            });
                        });
                } else {
                    return false;
                }
            });
        },
        // 算法三级对话框实现
        changeAlgorithmName() {
            this.algorithmVersion = true;
            this.algorithmVersionCount = 1;
            this.algorithmVersionOption = [];
            this.ruleForm.algorithmVersion = "";
            this.getAlgorithmVersionList();
        },
        getAlgorithmNameList(searchKey) {
            getMyAlgorithmList({
                pageIndex: this.algorithmNameCount,
                pageSize: 10,
                nameLike: searchKey,
            }).then((response) => {
                this.algorithmNameOption = this.algorithmNameOption.concat(
                    response.data.algorithms
                );
                this.algorithmNameTotal = response.data.totalSize;
            });
        },
        getAlgorithmVersionList() {
            getAlgorithmVersionList({
                pageIndex: this.algorithmVersionCount,
                pageSize: 10,
                algorithmId: this.ruleForm.algorithmId,
                fileStatus: 3,
            }).then((response) => {
                if (response.success) {
                    this.algorithmVersionOption =
                        this.algorithmVersionOption.concat(
                            response.data.algorithms
                        );
                    this.algorithmVersionTotal = response.data.totalSize;
                }
            });
        },
        loadAlgorithmName() {
            this.algorithmNameCount = this.algorithmNameCount + 1;
            if (this.algorithmNameOption.length < this.algorithmNameTotal) {
                this.getAlgorithmNameList(this.algorithmNameTemp);
            }
        },
        loadAlgorithmVersion() {
            this.algorithmVersionCount = this.algorithmVersionCount + 1;
            if (
                this.algorithmVersionOption.length < this.algorithmVersionTotal
            ) {
                this.getAlgorithmNameList();
            }
        },
        // 镜像二级对话框实现
        changeimageSource() {
            this.imageName = true;
            this.imageNameCount = 1;
            this.imageNameOption = [];
            this.ruleForm.imageItem = "";
            this.getImageNameList();
        },
        getImageNameList(searchKey) {
            if (this.ruleForm.imageSource === "my") {
                getMyImage({
                    pageIndex: this.imageNameCount,
                    pageSize: 10,
                    imageStatus: 3,
                    imageType: 1,
                    nameVerLike: searchKey,
                }).then((response) => {
                    if (response.data.images.length !== 0) {
                        const data = response.data.images;
                        const tableData = [];
                        this.imageNameTotal = response.data.totalSize;
                        data.forEach((item) => {
                            tableData.push({
                                ...item.image,
                                isShared: item.isShared,
                            });
                        });
                        this.imageNameOption =
                            this.imageNameOption.concat(tableData);
                    }
                });
            }
            if (this.ruleForm.imageSource === "pre") {
                getPreImage({
                    pageIndex: this.imageNameCount,
                    pageSize: 10,
                    imageStatus: 3,
                    imageType: 1,
                    nameVerLike: searchKey,
                }).then((response) => {
                    if (response.data.images.length !== 0) {
                        this.imageNameOption = this.imageNameOption.concat(
                            response.data.images
                        );
                        this.imageNameTotal = response.data.totalSize;
                    }
                });
            }
            if (this.ruleForm.imageSource === "common") {
                getPublicImage({
                    pageIndex: this.imageNameCount,
                    pageSize: 10,
                    imageStatus: 3,
                    imageType: 1,
                    nameVerLike: searchKey,
                }).then((response) => {
                    if (response.data.images.length !== 0) {
                        this.imageNameOption = this.imageNameOption.concat(
                            response.data.images
                        );
                        this.imageNameTotal = response.data.totalSize;
                    }
                });
            }
        },
        loadImageName() {
            this.imageNameCount = this.imageNameCount + 1;
            if (this.imageNameOption.length < this.imageNameTotal) {
                this.getImageNameList(this.imageTemp);
            }
        },
        // 数据集三级对话框
        changedataSetSource() {
            this.dataSetName = true;
            this.dataSetNameCount = 1;
            this.dataSetNameOption = [];
            this.ruleForm.dataSetId = "";
            this.ruleForm.dataSetVersion = "";
            this.dataSetChange = true;
            this.getDataSetNameList();
        },
        changeDataSetName() {
            this.dataSetVersion = true;
            this.dataSetVersionCount = 1;
            this.dataSetVersionOption = [];
            this.ruleForm.dataSetVersion = "";
            this.getDataSetVersionList();
        },
        getDataSetNameList(searchKey) {
            if (this.ruleForm.dataSetSource === "my") {
                getMyDatasetList({
                    pageIndex: this.dataSetNameCount,
                    pageSize: 10,
                    nameLike: searchKey,
                }).then((response) => {
                    if (response.data.datasets === null) {
                        response.data.datasets = [];
                    } else {
                        this.dataSetNameOption = this.dataSetNameOption.concat(
                            response.data.datasets
                        );
                        this.dataSetNameTotal = response.data.totalSize;
                    }
                });
            }
            if (this.ruleForm.dataSetSource === "pre") {
                getPresetDatasetList({
                    pageIndex: this.dataSetNameCount,
                    pageSize: 10,
                    nameLike: searchKey,
                }).then((response) => {
                    if (response.data.datasets !== null) {
                        this.dataSetNameOption = this.dataSetNameOption.concat(
                            response.data.datasets
                        );
                        this.dataSetNameTotal = response.data.totalSize;
                    } else {
                        response.data.datasets = [];
                    }
                });
            }
            if (this.ruleForm.dataSetSource === "common") {
                getPublicDatasetList({
                    pageIndex: this.dataSetNameCount,
                    pageSize: 10,
                    nameLike: searchKey,
                }).then((response) => {
                    if (response.data.datasets !== null) {
                        this.dataSetNameOption = this.dataSetNameOption.concat(
                            response.data.datasets
                        );
                        this.dataSetNameTotal = response.data.totalSize;
                    } else {
                        response.data.datasets = [];
                    }
                });
            }
        },
        getDataSetVersionList() {
            const data = {};
            data.datasetId = this.ruleForm.dataSetId;
            data.pageIndex = this.dataSetVersionCount;
            data.pageSize = 10;
            data.status = 3;
            getVersionList(data).then((response) => {
                if (response.data.versions !== null) {
                    this.dataSetVersionOption =
                        this.dataSetVersionOption.concat(
                            response.data.versions
                        );
                    this.dataSetVersionTotal = response.data.totalSize;
                }
            });
        },
        loadDataSetName() {
            this.dataSetNameCount = this.dataSetNameCount + 1;
            if (this.dataSetNameOption.length < this.dataSetNameTotal) {
                this.getDataSetNameList(this.dataSetTemp);
            }
        },
        loadDataSetVersion() {
            this.dataSetVersionCount = this.dataSetVersionCount + 1;
            if (this.dataSetVersionOption.length < this.dataSetVersionTotal) {
                this.getDataSetVersionList();
            }
        },
        // 远程请求算法名称
        remoteAlgorithm(searchKey) {
            if (searchKey === "") {
                this.algorithmNameTemp = "";
            } else {
                this.algorithmNameTemp = searchKey;
            }
            this.algorithmNameOption = [];
            this.algorithmNameCount = 1;
            this.getAlgorithmNameList(this.algorithmNameTemp);
        },
        getAlgorithmItem() {
            this.algorithmNameTemp = "";
            this.algorithmNameCount = 1;
            getMyAlgorithmList({
                pageIndex: this.algorithmNameCount,
                pageSize: 10,
            }).then((response) => {
                this.algorithmNameOption = response.data.algorithms;
                this.algorithmNameTotal = response.data.totalSize;
            });
        },
        // 远程请求镜像名称
        remoteImage(searchKey) {
            if (searchKey === "") {
                this.imageTemp = "";
            } else {
                this.imageTemp = searchKey;
            }
            this.imageNameOption = [];
            this.imageNameCount = 1;
            this.getImageNameList(this.imageTemp);
        },
        getImageItem() {
            this.imageTemp = "";
            this.imageNameCount = 1;
            if (this.ruleForm.imageSource === "my") {
                getMyImage({
                    pageIndex: this.imageNameCount,
                    pageSize: 10,
                    imageStatus: 3,
                    imageType: 1,
                }).then((response) => {
                    if (response.data.images.length !== 0) {
                        const data = response.data.images;
                        const tableData = [];
                        this.imageNameTotal = response.data.totalSize;
                        data.forEach((item) => {
                            tableData.push({
                                ...item.image,
                                isShared: item.isShared,
                            });
                        });
                        this.imageNameOption = tableData;
                    }
                });
            }
            if (this.ruleForm.imageSource === "pre") {
                getPreImage({
                    pageIndex: this.imageNameCount,
                    pageSize: 10,
                    imageStatus: 3,
                    imageType: 1,
                }).then((response) => {
                    if (response.data.images.length !== 0) {
                        this.imageNameOption = response.data.images;
                        this.imageNameTotal = response.data.totalSize;
                    }
                });
            }
            if (this.ruleForm.imageSource === "common") {
                getPublicImage({
                    pageIndex: this.imageNameCount,
                    pageSize: 10,
                    imageStatus: 3,
                    imageType: 1,
                }).then((response) => {
                    if (response.data.images.length !== 0) {
                        this.imageNameOption = response.data.images;
                        this.imageNameTotal = response.data.totalSize;
                    }
                });
            }
        },
        // 远程请求数据集名称
        remoteDataSet(searchKey) {
            if (searchKey === "") {
                this.dataSetTemp = "";
            } else {
                this.dataSetTemp = searchKey;
            }
            this.dataSetNameOption = [];
            this.dataSetNameCount = 1;
            this.getDataSetNameList(this.dataSetTemp);
        },
        getDataSetItem() {
            this.dataSetTemp = "";
            this.dataSetNameCount = 1;
            if (this.ruleForm.dataSetSource === "my") {
                getMyDatasetList({
                    pageIndex: this.dataSetNameCount,
                    pageSize: 10,
                }).then((response) => {
                    if (response.data.datasets === null) {
                        response.data.datasets = [];
                    } else {
                        this.dataSetNameOption = response.data.datasets;
                        this.dataSetNameTotal = response.data.totalSize;
                    }
                });
            }
            if (this.ruleForm.dataSetSource === "pre") {
                getPresetDatasetList({
                    pageIndex: this.dataSetNameCount,
                    pageSize: 10,
                }).then((response) => {
                    if (response.data.datasets !== null) {
                        this.dataSetNameOption = response.data.datasets;
                        this.dataSetNameTotal = response.data.totalSize;
                    } else {
                        response.data.datasets = [];
                    }
                });
            }
            if (this.ruleForm.dataSetSource === "common") {
                getPublicDatasetList({
                    pageIndex: this.dataSetNameCount,
                    pageSize: 10,
                }).then((response) => {
                    if (response.data.datasets !== null) {
                        this.dataSetNameOption = response.data.datasets;
                        this.dataSetNameTotal = response.data.totalSize;
                    } else {
                        response.data.datasets = [];
                    }
                });
            }
        },
        changeNum() {
            if (this.ruleForm.taskNumber == 1) {
                this.endpoint2 = "";
                this.port2 = "";
            }
        },
    },
};
</script>

<style lang="scss" scoped>
.line {
    text-align: center;
}

.inline {
    display: inline-block !important;
}

.block {
    display: block !important;
}

.tipClass {
    margin-left: 5px;
    font-size: 16px;
    color: #409eff;
}

.name {
    margin-bottom: 0px;
}

.tip {
    margin: 16px 0 16px 120px;
    color: #b3b3b3;
}

.tip span {
    color: #000;
    font-weight: 600;
}

.el-alert__icon {
    color: orange;
}
.interval {
    display: inline-block;
    width: 20px;
    text-align: center;
    font-size: 12px;
    font-weight: 800;
}
.wrapper {
    margin-top: 5px;
    // min-width: 400px;
}
// .el-dialog {
//   min-width: 1300px; /* 你想要的最小宽度 */
// }
.dialog {
    // min-height: 800px;
    min-width: 1000px;
}
</style>
