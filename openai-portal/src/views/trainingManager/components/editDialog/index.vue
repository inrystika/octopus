<template>
    <div>
        <el-dialog :title="title" width="55%" :visible.sync="CreateFormVisible" :before-close="handleDialogClose"
            :close-on-click-modal="false">
            <el-form ref="ruleForm" :model="ruleForm" :rules="rules" :label-width="formLabelWidth"
                class="demo-ruleForm">
                <el-form-item :label="name" :label-width="formLabelWidth" placeholder="请输入镜像名称" prop="name">
                    <el-input v-model="ruleForm.name" maxlength="30" show-word-limit />
                </el-form-item>
                <el-form-item label="任务描述" :label-width="formLabelWidth">
                    <el-input v-model="ruleForm.desc" type="textarea" maxlength="300" show-word-limit />
                </el-form-item>
                <!-- 算法三级框 -->
                <div>
                    <el-form-item label="算法类型" :class="{inline:algorithmName}">
                        <el-select v-model="ruleForm.algorithmSource" placeholder="请选择" @change="changealgorithmSource">
                            <el-option label="我的算法" value="my" />
                            <el-option label="预置算法" value="pre" />
                            <el-option label="公共算法" value="common" />
                        </el-select>
                    </el-form-item>
                    <el-form-item v-if="algorithmName" label="算法名称" prop="algorithmId" style="display:inline-block;">
                        <el-select v-model="ruleForm.algorithmId" v-loadmore="loadAlgorithmName" placeholder="请选择算法名称"
                            filterable remote :remote-method="remoteAlgorithm" @change="changeAlgorithmName">
                            <el-option v-for="item in algorithmNameOption" :key="item.algorithmId"
                                :label="item.algorithmName" :value="item.algorithmId" />
                        </el-select>
                    </el-form-item>
                    <el-form-item v-if="algorithmVersion" label="算法版本" prop="algorithmVersion"
                        style="display:inline-block;">
                        <el-select v-model="ruleForm.algorithmVersion" v-loadmore="loadAlgorithmVersion"
                            placeholder="请选择算法版本">
                            <el-option v-for="item in algorithmVersionOption"
                                :key="item.algorithmDetail.algorithmId+item.algorithmDetail.algorithmVersion"
                                :label="item.algorithmDetail.algorithmVersion"
                                :value="item.algorithmDetail.algorithmVersion" />
                        </el-select>
                    </el-form-item>
                </div>
                <!-- 镜像三级框 -->
                <div>
                    <el-form-item label="镜像类型" :class="{inline:imageName}">
                        <el-select v-model="ruleForm.imageSource" placeholder="请选择" @change="changeimageSource">
                            <el-option label="我的镜像" value="my" />
                            <el-option label="预置镜像" value="pre" />
                            <el-option label="公共镜像" value="common" />
                        </el-select>
                    </el-form-item>
                    <el-form-item v-if="imageName" label="镜像名称" prop="imageId" style="display: inline-block;">
                        <el-select v-model="ruleForm.imageId" v-loadmore="loadImageName" placeholder="请选择镜像名称"
                            filterable remote :remote-method="remoteImage" @change="changeimageName">
                            <el-option v-for="item in imageNameOption" :key="item.id"
                                :label="item.imageName+':'+item.imageVersion" :value="item.id" />
                        </el-select>
                    </el-form-item>
                </div>
                <!-- 数据集三级框 -->
                <div>
                    <el-form-item label="数据集类型" :class="{inline:dataSetName}">
                        <el-select v-model="ruleForm.dataSetSource" placeholder="请选择" @change="changedataSetSource">
                            <el-option label="我的数据集" value="my" />
                            <el-option label="预置数据集" value="pre" />
                            <el-option label="公共数据集" value="common" />
                        </el-select>
                    </el-form-item>
                    <el-form-item v-if="dataSetName" label="数据集名称" prop="dataSetId" style="display: inline-block;">
                        <el-select v-model="ruleForm.dataSetId" v-loadmore="loadDataSetName" placeholder="请选择数据集名称"
                            filterable remote :remote-method="remoteDataSet" @change="changeDataSetName">
                            <el-option v-for="item in dataSetNameOption" :key="item.id+item.name" :label="item.name"
                                :value="item.id" />
                        </el-select>
                    </el-form-item>
                    <el-form-item v-if="dataSetVersion" label="数据集版本" prop="dataSetVersion"
                        style="display: inline-block;">
                        <el-select v-model="ruleForm.dataSetVersion" v-loadmore="loadDataSetVersion"
                            placeholder="请选择数据集版本">
                            <el-option v-for="item in dataSetVersionOption" :key="item.datasetId+item.version"
                                :label="item.version" :value="item.version" />
                        </el-select>
                    </el-form-item>
                </div>
                <el-divider />
                <el-form-item label="分布式" prop="distributed ">
                    <el-select v-model="ruleForm.isDistributed" :disabled="true">
                        <el-option label="是" :value="true" />
                        <el-option label="否" :value="false" />
                    </el-select>
                </el-form-item>
                <div v-if="show">
                    <el-form-item label="运行命令" prop="command">
                        <el-input v-model="ruleForm.command" type="textarea" />
                    </el-form-item>
                    <el-form-item label="运行参数">
                        <div v-for="(item, index) in ruleForm.config[0].parameters" :key="index">
                            <el-form-item style="margin-bottom:10px">
                                <el-input v-model="item.key" placeholder="key" style="width: 20%;" />
                                <span style="margin:0 10px 0 10px">=</span>
                                <el-input v-model="item.value" placeholder="value" style="width: 20%;" />
                                <i class="el-icon-delete" @click="deleteItem(item, index)"></i>
                            </el-form-item>
                        </div>
                        <el-button type="primary" @click="addItem">增加</el-button>
                        <el-button type="text" :disabled="showArg" @click="open">预览</el-button>
                    </el-form-item>
                    <el-form-item label="资源规格" prop="resourceSpecId">
                        <el-select v-model="ruleForm.resourceSpecId" placeholder="请选择资源规格" style="width:35%">
                            <el-option v-for="(item,index) in resourceOptions" :key="index" :label="item.label"
                                :value="item.value" />
                        </el-select>
                    </el-form-item>
                </div>
                <div v-if="!show">
                    <traningList :training-table="table" :resource="resourceOptions" @tableData="getTableData" />
                </div>
            </el-form>
            <div slot="footer" class="dialog-footer">
                <el-button v-if="showTraning" type="success" @click="traningAndSave('traning')" v-preventReClick>开始训练
                </el-button>
                <el-button v-if="!showTemplate" type="primary" @click="traningAndSave('save')" v-preventReClick>保存模板
                </el-button>
                <el-button type="warning" @click="cancel">取消</el-button>
            </div>
        </el-dialog>
    </div>
</template>

<script>
    import traningList from './traningList.vue'
    import { createTask, getResourceList, editTemplate } from '@/api/trainingManager'
    import { getPresetAlgorithmList, getPublicAlgorithmList, getMyAlgorithmList, getAlgorithmVersionList } from '@/api/modelDev'
    import { getMyImage, getPublicImage, getPreImage } from '@/api/imageManager'
    import { getMyDatasetList, getPublicDatasetList, getPresetDatasetList, getVersionList } from '@/api/datasetManager'
    export default {
        name: "DialogEditForm",
        components: {
            traningList

        },
        props: {
            row: {
                type: Object,
                default: () => { }
            },
            flag: {
                type: Number,
                default: undefined
            }
        },
        data() {
            return {
                // 模板id
                id: '',
                show: true,
                showTraning: true,
                showTemplate: false,
                table: [],
                ruleForm: {
                    name: '',
                    desc: '',
                    algorithmSource: '',
                    algorithmId: '',
                    algorithmVersion: '',
                    imageSource: '',
                    imageId: "",
                    dataSetSource: '',
                    dataSetId: '',
                    dataSetVersion: '',
                    isDistributed: false,
                    config: [{
                        name: '',
                        command: "",
                        resourceSpecId: undefined,
                        parameters: [
                            {
                                key: "",
                                value: ""
                            }
                        ]

                    }],
                    resourceSpecId: "",
                    command: ''
                },

                CreateFormVisible: true,
                rules: {
                    name: [
                        { required: true, message: '请输入任务名称', trigger: 'blur' }
                    ],
                    childName: [
                        { required: true, message: '请输入任务名称', trigger: 'blur' }
                    ],

                    algorithmId: [
                        { required: true, message: '请选择算法名称', trigger: 'change' }
                    ],
                    algorithmVersion: [{ required: true, message: '请选择算法版本', trigger: 'change' }],

                    imageId: [
                        { required: true, message: '请选择镜像名称', trigger: 'change' }
                    ],

                    dataSetId: [
                        { required: true, message: '请选择数据集名称', trigger: 'change' }
                    ],
                    dataSetVersion: [
                        { required: true, message: '请选择数据集版本', trigger: 'change' }
                    ],
                    isDistributed: [
                        { required: true, message: '请选择是否为分布式', trigger: 'change' }
                    ],
                    command: [
                        { required: true, message: '请填写运行命令', trigger: 'blur' }
                    ],
                    resourceSpecId: [
                        { required: true, message: '请选择活资源规格', trigger: 'change' }
                    ]
                },
                formLabelWidth: '120px',
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
                resourceOptions: [],
                dataSetNameTotal: undefined,
                dataSetVersionTotal: undefined,
                // 临时镜像版本号
                imageVersion: '',
                // 临时temp
                temp: { algorithmId: '', imageId: '', dataSetId: '', imageVersion: '' },
                algorithmChange: false,
                imageChange: false,
                dataSetChange: false,
                tip: false,
                algorithmNameTemp: '',
                imageTemp: '',
                dataSetTemp: '',
                name: ''

            }
        },
        computed: {
            title: function () {
                switch (this.flag) {
                    case 1:
                        this.showTraning = false
                        this.name = '模版名称'
                        return '编辑任务模板'
                        break
                    case 2:
                        this.showTraning = true
                        this.showTemplate = true
                        this.name = '任务名称'
                        return '创建训练任务'
                        break
                }
            },
            showArg: function () {
                let flag = true
                if (this.ruleForm.config[0].parameters === null) {
                    return flag
                } else {
                    this.ruleForm.config[0].parameters.forEach(
                        item => {
                            if (item.key !== "" && item.value !== "") {
                                flag = false
                                return flag
                            }
                        }
                    )

                    return flag
                }
            }

        },
        watch: {
            'ruleForm.isDistributed': {
                deep: true,
                handler: function (newV, oldV) {
                    if (newV === true && oldV === false) { this.show = false; } else if (newV === false && oldV === true) { this.show = true; }
                }

            }

        },
        created() {
            this.algorithmName = true
            this.algorithmVersion = true
            this.imageName = true
            this.dataSetName = true
            this.dataSetVersion = true
            // 获取模板信息
            // 1编辑模块2创建训练任务
            this.getResourceList()
        },
        methods: {
            // 获取资源规格
            getResourceList() {
                getResourceList().then(response => {
                    if (response.success) {
                        response.data.mapResourceSpecIdList.train.resourceSpecs.forEach(
                            item => {
                                this.resourceOptions.push({ label: item.name + ' ' + item.price + '机时/h', value: item.id })
                            }
                        )
                        const data = JSON.parse(JSON.stringify(this.row))
                        this.temp.imageVersion = data.imageVersion
                        this.temp.algorithmId = data.algorithmId
                        this.temp.imageId = data.imageId
                        this.temp.dataSetId = data.dataSetId
                        this.id = data.id
                        delete this.row.createdAt
                        delete this.row.workspaceId
                        delete this.row.userId
                        delete this.row.updatedAt
                        delete this.row.id
                        delete this.row.imageVersion
                        this.ruleForm = this.row
                        this.ruleForm.imageId = this.row.imageName + ':' + this.temp.imageVersion
                        this.ruleForm.algorithmId = this.row.algorithmName
                        this.ruleForm.dataSetId = this.row.dataSetName
                        delete this.ruleForm.algorithmName
                        delete this.ruleForm.imageName
                        delete this.ruleForm.dataSetName
                        this.$set(this.ruleForm, "command", data.config[0].command)
                        this.$set(this.ruleForm, "resourceSpecId", data.config[0].resourceSpecId)
                        if (this.ruleForm.isDistributed) {
                            this.table = this.ruleForm.config
                        }
                        if (this.flag === 2) {
                            this.ruleForm.name = ''
                        }
                    }
                })
            },
            addItem() {
                if (this.ruleForm.config[0].parameters === null) {
                    this.ruleForm.config[0].parameters = [{
                        key: '',
                        value: ''
                    }]
                } else {
                    this.ruleForm.config[0].parameters.push({
                        key: '',
                        value: ''
                    })
                }
            },
            deleteItem(item, index) {
                this.ruleForm.config[0].parameters.splice(index, 1)
            },
            cancel() {
                let message = ''
                if (this.flag === 1) { message = '此操作将放弃编辑任务模板, 是否继续?' } else { message = '此操作将放弃创建训练任务，是否继续' }
                this.$confirm(message, '提示', {
                    confirmButtonText: '确定',
                    cancelButtonText: '取消',
                    type: 'warning'
                }).then(() => {
                    this.$emit('cancel', false)
                }).catch(() => {
                    this.$message({
                        type: 'info',
                        message: '已取消操作'
                    });
                });
            },
            nameIsRepeat(val) {
                const isRepeat = function (arr) {
                    var hash = {};
                    for (var i in arr) {
                        if (hash[arr[i]]) { return true; }
                        hash[arr[i]] = true;
                    }

                    return false;
                }
                const data = []
                val.config.forEach(
                    item => {
                        data.push(item.name)
                    }
                )
                return isRepeat(data)
            },
            traningAndSave(val) {
                this.$refs['ruleForm'].validate((valid) => {
                    if (valid) {
                        if (!this.ruleForm.isDistributed) {
                            this.ruleForm.config = [this.ruleForm.config[0]]
                            this.ruleForm.config[0].command = this.ruleForm.command
                            this.ruleForm.config[0].resourceSpecId = this.ruleForm.resourceSpecId
                            this.ruleForm.config[0].name = this.ruleForm.name
                            this.ruleForm.config[0].taskNumber = 1
                            this.ruleForm.config[0].minFailedTaskCount = 1
                            this.ruleForm.config[0].minSucceededTaskCount = 1
                            delete this.ruleForm.config[0].isMainRole
                            delete this.ruleForm.config[0].name
                        }
                        var data = JSON.parse(JSON.stringify(this.ruleForm))
                        if (!this.algorithmChange) {
                            data.algorithmId = this.temp.algorithmId
                        }
                        if (!this.imageChange) {
                            data.imageId = this.temp.imageId
                        }
                        if (!this.dataSetChange) {
                            data.dataSetId = this.temp.dataSetId
                        }
                        delete data.command;
                        delete data.resourceSpecId
                        delete data.algorithmSource
                        delete data.imageSource
                        delete data.dataSetSource
                        if (this.ruleForm.isDistributed && this.tip) {
                            this.$message({
                                message: '请添加子任务',
                                type: 'warning'
                            });
                            return
                        }
                        let isRepeat = true
                        if (data.config.length > 1) {
                            isRepeat = this.nameIsRepeat(data)
                        } else { isRepeat = false }
                        if (!isRepeat) {
                            if (val === 'traning') {
                                createTask(data).then(response => {
                                    if (response.success) {
                                        this.$message({
                                            message: '创建成功',
                                            type: 'success'
                                        });
                                        this.$emit('confirm', false)
                                        this.$emit('createTraning')
                                    } else {
                                        this.$message({
                                            message: this.getErrorMsg(response.error.subcode),
                                            type: 'warning'
                                        });
                                    }
                                })
                            }
                            if (val === 'save') {
                                editTemplate({ data: data, id: this.id }).then(response => {
                                    if (response.success) {
                                        this.$message({
                                            message: '保存成功',
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
                            }
                        } else {
                            this.$message({
                                message: '子任务名称不能重复',
                                type: 'warning'
                            });
                        }
                    } else {
                        console.log('error submit!!');
                        return false;
                    }
                });
            },
            handleDialogClose() {
                this.$emit('close', false)
            },
            getTableData(val) {
                this.ruleForm.config = val
                if (this.ruleForm.config.length === 0) {
                    this.tip = true
                } else { this.tip = false }
            },
            // 算法三级对话框实现
            changealgorithmSource() {
                this.algorithmName = true
                this.algorithmNameCount = 1
                this.algorithmNameOption = []
                this.ruleForm.algorithmId = ''
                this.ruleForm.algorithmVersion = ''
                this.algorithmChange = true
                this.getAlgorithmNameList()
            },
            changeAlgorithmName() {
                this.algorithmVersion = true
                this.algorithmVersionCount = 1
                this.algorithmVersionOption = []
                this.ruleForm.algorithmVersion = ''
                this.getAlgorithmVersionList()
            },
            getAlgorithmNameList(searchKey) {
                if (this.ruleForm.algorithmSource === 'my') {
                    getMyAlgorithmList({ pageIndex: this.algorithmNameCount, pageSize: 10, nameLike: searchKey }).then(response => {
                        if (response.data.algorithms.length !== 0) {
                            this.algorithmNameOption = this.algorithmNameOption.concat(response.data.algorithms);
                            this.algorithmNameTotal = response.data.totalSize
                        }
                    })
                }
                if (this.ruleForm.algorithmSource === 'pre') {
                    getPresetAlgorithmList({ pageIndex: this.algorithmNameCount, pageSize: 10, nameLike: searchKey }).then(response => {
                        if (response.data.algorithms.length !== 0) {
                            this.algorithmNameOption = this.algorithmNameOption.concat(response.data.algorithms)
                            this.algorithmNameTotal = response.data.totalSize
                        }
                    })
                }
                if (this.ruleForm.algorithmSource === 'common') {
                    getPublicAlgorithmList({ pageIndex: this.algorithmNameCount, pageSize: 10, nameLike: searchKey }).then(response => {
                        if (response.data.algorithms.length !== 0) {
                            this.algorithmNameOption = this.algorithmNameOption.concat(response.data.algorithms);
                            this.algorithmNameTotal = response.data.totalSize
                        }
                    })
                }
            },
            getAlgorithmVersionList() {
                getAlgorithmVersionList({ pageIndex: this.algorithmVersionCount, pageSize: 10, algorithmId: this.ruleForm.algorithmId, fileStatus: 3 }).then(response => {
                    if (response.success) {
                        this.algorithmVersionOption = this.algorithmVersionOption.concat(response.data.algorithms)
                        this.algorithmVersionTotal = response.data.totalSize
                    }
                })
            },
            loadAlgorithmName() {
                this.algorithmNameCount = this.algorithmNameCount + 1
                if (this.algorithmNameOption.length < this.algorithmNameTotal) {
                    this.getAlgorithmNameList(this.algorithmNameTemp)
                }
            },
            loadAlgorithmVersion() {
                this.algorithmVersionCount = this.algorithmVersionCount + 1
                if (this.algorithmVersionOption.length < this.algorithmVersionTotal) {
                    this.getAlgorithmNameList
                }
            },
            // 镜像二级对话框实现
            changeimageSource() {
                this.imageName = true
                this.imageNameCount = 1
                this.imageNameOption = []
                this.ruleForm.imageId = ''
                this.ruleForm.imageVersion = ''
                this.imageChange = true
                this.getImageNameList()
            },
            // 镜像三级对话框实现
            changeimageName() {
                this.imageVersion = true
                this.ruleForm.imageVersion = ''
                this.imageNameOption.forEach(item => {
                    if (item.id === this.ruleForm.imageId) {
                        this.ruleForm.imageVersion = item.imageVersion
                    }
                })
            },
            getImageNameList(searchKey) {
                if (this.ruleForm.imageSource === 'my') {
                    getMyImage({ pageIndex: this.imageNameCount, pageSize: 10, imageStatus: 3, imageType: 2, nameVerLike: searchKey }).then(response => {
                        if (response.data.images.length !== 0) {
                            const data = response.data.images;
                            const tableData = [];
                            this.imageNameTotal = response.data.totalSize
                            data.forEach(item => {
                                tableData.push({ ...item.image, isShared: item.isShared })
                            })
                            this.imageNameOption = this.imageNameOption.concat(tableData)
                        }
                    })
                }
                if (this.ruleForm.imageSource === 'pre') {
                    getPreImage({ pageIndex: this.imageNameCount, pageSize: 10, imageStatus: 3, imageType: 2, nameVerLike: searchKey }).then(response => {
                        if (response.data.images.length !== 0) { this.imageNameOption = this.imageNameOption.concat(response.data.images); this.imageNameTotal = response.data.totalSize }
                    })
                }
                if (this.ruleForm.imageSource === 'common') {
                    getPublicImage({ pageIndex: this.imageNameCount, pageSize: 10, imageStatus: 3, imageType: 2, nameVerLike: searchKey }).then(response => {
                        if (response.data.images.length !== 0) { this.imageNameOption = this.imageNameOption.concat(response.data.images); this.imageNameTotal = response.data.totalSize }
                    })
                }
            },
            loadImageName() {
                this.imageNameCount = this.imageNameCount + 1
                if (this.imageNameOption.length < this.imageNameTotal) {
                    this.getImageNameList(this.imageTemp)
                }
            },
            // 数据集三级对话框
            changedataSetSource() {
                this.dataSetName = true
                this.dataSetNameCount = 1
                this.dataSetNameOption = []
                this.ruleForm.dataSetId = ''
                this.ruleForm.dataSetVersion = ''
                this.dataSetChange = true
                this.getDataSetNameList()
            },
            changeDataSetName() {
                this.dataSetVersion = true
                this.dataSetVersionCount = 1
                this.dataSetVersionOption = []
                this.ruleForm.dataSetVersion = ''
                this.getDataSetVersionList()
            },
            getDataSetNameList(searchKey) {
                if (this.ruleForm.dataSetSource === 'my') {
                    getMyDatasetList({ pageIndex: this.dataSetNameCount, pageSize: 10, nameLike: searchKey }).then(response => {
                        if (response.data.datasets === null) {
                            response.data.datasets = []
                        }
                        this.dataSetNameOption = this.dataSetNameOption.concat(response.data.datasets)
                        this.dataSetNameTotal = response.data.totalSize
                    })
                }
                if (this.ruleForm.dataSetSource === 'pre') {
                    getPresetDatasetList({ pageIndex: this.dataSetNameCount, pageSize: 10, nameLike: searchKey }).then(response => {
                        if (response.data.datasets === null) {
                            response.data.datasets = []
                        }
                        this.dataSetNameOption = this.dataSetNameOption.concat(response.data.datasets); this.dataSetNameTotal = response.data.totalSize
                    })
                }
                if (this.ruleForm.dataSetSource === 'common') {
                    getPublicDatasetList({ pageIndex: this.dataSetNameCount, pageSize: 10, nameLike: searchKey }).then(response => {
                        if (response.data.datasets === null) {
                            response.data.datasets = []
                        }
                        this.dataSetNameOption = this.dataSetNameOption.concat(response.data.datasets); this.dataSetNameTotal = response.data.totalSize
                    })
                }
            },
            getDataSetVersionList() {
                const data = {}
                data.datasetId = this.ruleForm.dataSetId
                data.pageIndex = this.dataSetVersionCount
                data.pageSize = 10
                data.status = 3
                getVersionList(data).then(response => {
                    if (response.data.versions !== null) {
                        this.dataSetVersionOption = this.dataSetVersionOption.concat(response.data.versions)
                        this.dataSetVersionCount = response.data.totalSize
                    }
                })
            },
            loadDataSetName() {
                this.dataSetNameCount = this.dataSetNameCount + 1
                if (this.dataSetNameOption.length < this.dataSetNameTotal) {
                    this.getDataSetNameList(this.dataSetTemp)
                }
            },
            loadDataSetVersion() {
                this.dataSetVersionCount = this.dataSetVersionCount + 1
                if (this.dataSetVersionOption.length < this.dataSetVersionTotal) {
                    this.getDataSetVersionList()
                }
            },
            // 运行参数预览
            open() {
                const data = JSON.parse(JSON.stringify(this.ruleForm.config[0].parameters))
                let argument = ''
                if (data) {
                    data.forEach(
                        item => {
                            argument += '--' + item.key + '=' + item.value + ' '
                        }
                    )
                } else { argument = '' }
                this.$alert(argument, '运行参数', {
                    confirmButtonText: '确定',
                    callback: action => {
                    }
                });
            },
            // 远程请求算法名称
            remoteAlgorithm(a) {
                if (a == '') { this.algorithmNameTemp = '' } else { this.algorithmNameTemp = a }
                this.algorithmNameOption = []
                this.algorithmNameCount = 1
                this.getAlgorithmNameList(this.algorithmNameTemp)
            },
            // 远程请求镜像名称
            remoteImage(a) {
                if (a == '') { this.imageTemp = '' } else { this.imageTemp = a }
                this.imageNameOption = []
                this.imageNameCount = 1
                this.getImageNameList(this.imageTemp)
            },
            // 远程请求数据集名称
            remoteDataSet(a) {
                if (a == '') { this.dataSetTemp = '' } else { this.dataSetTemp = a }
                this.dataSetNameOption = []
                this.dataSetNameCount = 1
                this.getDataSetNameList(this.dataSetTemp)
            }

        }
    }
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
</style>