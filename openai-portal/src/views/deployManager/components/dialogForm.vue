<template>
    <div>
        <el-dialog :title="'编辑在线服务'" width="1200px" :visible.sync="CreateFormVisible" :before-close="handleDialogClose"
            :close-on-click-modal="false">
            <el-form :model="ruleForm" :rules="rules" ref="ruleForm" label-width="100px" class="demo-ruleForm">
                <el-form-item label="服务名称" prop="name">
                    <el-input v-model="ruleForm.name"></el-input>
                </el-form-item>
                <el-form-item label="文字描述" prop="desc">
                    <el-input type="textarea" v-model="ruleForm.desc"></el-input>
                </el-form-item>
                <el-form-item label="服务类型" prop="serviceType">
                    <el-select v-model="ruleForm.serviceType" placeholder="请选择服务类型">
                        <el-option label="HTTP模式" value="http"></el-option>
                        <el-option label="GRPC模式" value="grpc" disabled></el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="计算框架" prop="modelFrame">
                    <el-select v-model="ruleForm.modelFrame" placeholder="请选择服务类型">
                        <el-option label="PyTorch" value="pytorch"></el-option>
                        <el-option label="TensorFlow" value="tensorflow"></el-option>
                    </el-select>
                </el-form-item>
                <!-- 模型三级框 -->
                <div>
                    <el-form-item label="模型类型" prop="modelSource" style="display:inline-block;">
                        <el-select v-model="ruleForm.modelSource" placeholder="请选择模型类型" @change="changeModelSource"
                            :disabled="uncheckable">
                            <el-option label="我的模型" value="1" />
                            <el-option label="公共模型" value="2" />
                            <el-option label="预置模型" value="3" />
                        </el-select>
                    </el-form-item>
                    <el-form-item v-if="modelName" label="模型名称" prop="modelId" style="display:inline-block;">
                        <el-select v-model="ruleForm.modelId" v-loadmore="loadModelName" placeholder="请选择模型名称"
                            filterable remote :remote-method="remoteModel" @change="changeModelName"
                            @click.native="getModelItem" :disabled="uncheckable">
                            <el-option v-for="item in modelNameOption" :key="item.modelId+item.modelName"
                                :label="item.modelName" :value="item.modelId" />
                        </el-select>
                    </el-form-item>
                    <el-form-item v-if="modelVersion" label="模型版本" prop="modelVersion" style="display:inline-block;">
                        <el-select v-model="ruleForm.modelVersion" v-loadmore="loadModelVersion" placeholder="请选择模型版本"
                            :disabled="uncheckable">
                            <el-option v-for="item in modelVersionOption" :key="item.modelId+item.version"
                                :label="item.version" :value="item.version" />
                        </el-select>
                    </el-form-item>
                </div>
                <el-form-item label="资源类型" prop="resourceType">
                    <el-select v-model="ruleForm.resourceType" placeholder="请选择服务类型">
                        <el-option label="cpu" value="cpu"></el-option>
                        <el-option label="gpu" value="gpu" disabled></el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="资源规格" prop="resourceSpecId">
                    <el-select v-model="ruleForm.resourceSpecId" placeholder="请选择资源规格" style="width:35%">
                        <el-option v-for="item in resourceOptions" :key="item.id" :label="item.label"
                            :value="item.value" />
                    </el-select>
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" @click="submitForm('ruleForm')">提交服务</el-button>
                </el-form-item>
            </el-form>

        </el-dialog>
    </div>
</template>
<script>
    import { createDeploy } from '@/api/deployManager.js'
    import { getMyModel, getPreModel, getPublicModel, getPublicList, getNoPublicList } from '@/api/modelManager.js'
    import { getResourceList } from '@/api/trainingManager.js'
    import upload from '@/components/upload/index.vue'
    export default {
        name: "DialogCreateForm",
        components: {
            upload
        },
        props: {
            row: {
                type: Object,
                default: () => { }
            },
        },
        data() {
            return {
                CreateFormVisible: true,
                ruleForm: {
                    name: '',
                    serviceType: '',
                    desc: '',
                    modelSource: '',
                    modelId: '',
                    modelVersion: '',
                    resourceType: '',
                    resourceSpecId: '',
                    domain: this.GLOBAL.DOMAIN
                },
                rules: {
                    name: [
                        { required: true, message: '请输入服务名称', trigger: 'blur' },
                    ],
                    desc: [
                        { required: true, message: '请填写文字描述', trigger: 'blur' }
                    ],
                    serviceType: [
                        { required: true, message: '请选择服务类型', trigger: 'change' }
                    ],
                    modelFrame: [{ required: true, message: '请选择计算框架', trigger: 'change' }],
                    modelSource: [
                        { required: true, message: '请选择模型类型', trigger: 'change' }
                    ],
                    modelId: [
                        { required: true, message: '请选择模型名称', trigger: 'change' }
                    ],
                    modelVersion: [
                        { required: true, message: '请选择模型版本', trigger: 'change' }
                    ],
                    resourceType: [
                        { required: true, message: '请选择资源类型', trigger: 'change' }
                    ],
                    resourceSpecId: [
                        { required: true, message: '请选择资源规格', trigger: 'change' }
                    ],

                },
                resourceOptions: [],
                // 模型类型
                modelType: undefined,
                // 模型名称
                modelName: false,
                modelNameCount: 1,
                modelNameOption: [],
                modelNameTotal: undefined,
                modelNameTemp: '',
                // 模型版本
                modelVersion: false,
                modelVersionCount: 1,
                modelVersionOption: [],
                modelVersionTotal: undefined,
                modelVersionTemp: '',
                uncheckable: false,
                //临时模型名称ID
                tempId: undefined,
                flag:false

            }
        },
        watch: {},
        created() {
            if (JSON.stringify(this.row) !== '{}') {
                this.uncheckable = true
                this.modelName = true
                this.modelVersion = true
                this.ruleForm.modelSource = this.row.type.toString()       
                this.ruleForm.modelId = this.row.modelName
                this.ruleForm.modelVersion = this.row.version
                this.tempId = this.row.modelId
                this.flag=true
            }
            this.getResourceList()
        },
        beforeDestroy() {

        },
        methods: {
            // 模型三级对话框实现
            changeModelSource() {
                this.modelName = true
                this.modelNameCount = 1
                this.modelNameOption = []
            },
            getModel(searchKey) {
                if (this.ruleForm.modelSource == 1) {
                    getMyModel({ pageIndex: this.modelNameCount, pageSize: 10, nameLike: searchKey }).then(response => {
                        if (response.success) {
                            if (response.data.models !== null) {
                                this.modelNameTotal = response.data.totalSize
                                this.modelNameOption = this.modelNameOption.concat(response.data.models);
                            }
                        } else {
                            this.$message({
                                message: this.getErrorMsg(response.error.subcode),
                                type: 'warning'
                            });
                        }
                    })
                }
                if (this.ruleForm.modelSource == 2) {
                    getPublicModel({ pageIndex: this.modelNameCount, pageSize: 10, nameLike: searchKey }).then(response => {
                        if (response.success) {
                            if (response.data.models !== null) {
                                this.modelNameTotal = response.data.totalSize
                                this.modelNameOption = this.modelNameOption.concat(response.data.models);
                            }
                        } else {
                            this.$message({
                                message: this.getErrorMsg(response.error.subcode),
                                type: 'warning'
                            });
                        }
                    })
                }
                if (this.ruleForm.modelSource == 3) {
                    getPreModel({ pageIndex: this.modelNameCount, pageSize: 10, nameLike: searchKey }).then(response => {
                        if (response.success) {
                            if (response.data.models !== null) {
                                this.modelNameTotal = response.data.totalSize
                                this.modelNameOption = this.modelNameOption.concat(response.data.models);
                            }
                        } else {
                            this.$message({
                                message: this.getErrorMsg(response.error.subcode),
                                type: 'warning'
                            });
                        }
                    })
                }
            },
            loadModelName() {
                this.modelNameCount = this.modelNameCount + 1
                if (this.modelNameOption.length < this.modelNameTotal) {
                    this.getModel(this.modelNameTemp)
                }
            },
            remoteModel(searchName) {
                if (searchName == '') {
                    this.modelNameTemp = ''
                } else {
                    this.modelNameTemp = searchName
                }
                this.modelNameOption = []
                this.modelNameTotal = 1
                this.getModel(this.modelNameTemp)
            },
            getModelItem() {
                this.modelNameTemp = ''
                this.modelNameCount = 1
                if (this.ruleForm.modelSource == 1) {
                    getMyModel({ pageIndex: this.modelNameCount, pageSize: 10, }).then(response => {
                        if (response.success) {
                            if (response.data.models !== null) {
                                this.modelNameTotal = response.data.totalSize
                                this.modelNameOption = this.modelNameOption.concat(response.data.models);
                            }
                        } else {
                            this.$message({
                                message: this.getErrorMsg(response.error.subcode),
                                type: 'warning'
                            });
                        }
                    })
                }
                if (this.ruleForm.modelSource == 2) {
                    getPublicModel({ pageIndex: this.modelNameCount, pageSize: 10, }).then(response => {
                        if (response.success) {
                            if (response.data.models !== null) {
                                this.modelNameTotal = response.data.totalSize
                                this.modelNameOption = this.modelNameOption.concat(response.data.models);
                            }
                        } else {
                            this.$message({
                                message: this.getErrorMsg(response.error.subcode),
                                type: 'warning'
                            });
                        }
                    })
                }
                if (this.ruleForm.modelSource == 3) {
                    getPreModel({ pageIndex: this.modelNameCount, pageSize: 10, }).then(response => {
                        if (response.success) {
                            if (response.data.models !== null) {
                                this.modelNameTotal = response.data.totalSize
                                this.modelNameOption = this.modelNameOption.concat(response.data.models);
                            }
                        } else {
                            this.$message({
                                message: this.getErrorMsg(response.error.subcode),
                                type: 'warning'
                            });
                        }
                    })
                }
            },
            changeModelName() {
                this.modelVersion = true
                this.modelVersionCount = 1
                this.modelVersionOption = []
                this.ruleForm.modelVersion = ''
                this.getModelVersionList()
            },
            getModelVersionList() {
                if (this.ruleForm.modelSource !== 2) {
                    getNoPublicList({ pageIndex: this.modelVersionCount, pageSize: 1, modelId: this.ruleForm.modelId }).then(response => {
                        if (response.success) {
                            if (response.data.modelVersions !== null) {
                                this.modelVersionTotal = response.data.totalSize
                                const data = response.data.modelVersions
                                const tableData = []
                                data.forEach(item => {
                                    tableData.push({ ...item.versionDetail })
                                })
                                this.modelVersionOption = this.modelVersionOption.concat(tableData);
                            }
                        } else {
                            this.$message({
                                message: this.getErrorMsg(response.error.subcode),
                                type: 'warning'
                            });
                        }
                    })
                } else {
                    getPublicList({ pageIndex: this.modelVersionCount, pageSize: 1, modelId: this.ruleForm.modelId }).then(response => {
                        if (response.success) {
                            if (response.data.modelVersions !== null) {
                                this.modelVersionTotal = response.data.totalSize
                                this.modelVersionOption = this.modelVersionOption.concat(response.data.models);
                                console.log(this.modelVersionOption, "OPPO")
                            }
                        } else {
                            this.$message({
                                message: this.getErrorMsg(response.error.subcode),
                                type: 'warning'
                            });
                        }
                    })
                }
            },
            loadModelVersion() {
                this.modelVersionCount = this.modelVersionCount + 1
                if (this.modelVersionOption.length < this.modelVersionTotal) {
                    this.getModelVersionList()
                }
            },
            // 获取资源规格
            getResourceList() {
                getResourceList().then(response => {
                    if (response.success) {
                        response.data.mapResourceSpecIdList.train.resourceSpecs.forEach(
                            item => {
                                this.resourceOptions.push({ label: item.name + ' ' + item.price + '机时/h', value: item.id })
                            }
                        )
                    } else {
                        this.$message({
                            message: this.getErrorMsg(response.error.subcode),
                            type: 'warning'
                        });
                    }
                })
            },
            submitForm(formName) {
                this.$refs[formName].validate((valid) => {
                    if (valid) {
                        if(this.flag){
                            this.ruleForm.modelId=this.tempId
                        }
                        createDeploy(this.ruleForm).then(response => {
                            if (response.success) {
                                this.$message({
                                    message: '提交成功',
                                    type: 'success'
                                });
                                this.$emit('confirm', false)
                            } else {
                                this.$message({
                                    message: this.getErrorMsg(response.error.subcode),
                                    type: 'warning'
                                });
                            }
                        })

                    } else {
                        console.log('error submit!!');
                        return false;
                    }
                });
            },
            handleDialogClose() {
                this.$emit('close', false)
            },
            confirm(val) { this.$emit('confirm', val) },
            cancel() {
                this.$emit('cancel', false)
            },
            isCloseX(val) {
                this.close = val
            }
        }
    }
</script>
<style lang="scss" scoped>
</style>