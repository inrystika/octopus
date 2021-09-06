<template>
    <div>
        <el-dialog :title="title" width="55%" :visible.sync="CreateFormVisible" :before-close="handleDialogClose"
            :close-on-click-modal="false">
            <el-form :model="ruleForm" :rules="rules" ref="ruleForm" :label-width="formLabelWidth"
                class="demo-ruleForm">
                <el-form-item label="任务名称" :label-width="formLabelWidth" placeholder="请输入镜像名称" prop="name">
                    <el-input v-model="ruleForm.name" maxlength="30" show-word-limit></el-input>
                </el-form-item>
                <el-form-item label="任务描述" :label-width="formLabelWidth">
                    <el-input type="textarea" v-model="ruleForm.desc" maxlength="300" show-word-limit></el-input>
                </el-form-item>
                <!-- 算法三级框 -->
                <div>
                    <el-form-item label="算法类型" prop="algorithmSource" :class="{inline:algorithmName}">
                        <el-select v-model="ruleForm.algorithmSource" placeholder="请选择" @change="changealgorithmSource">
                            <el-option label="我的算法" value="my"></el-option>
                            <el-option label="预置算法" value="pre"></el-option>
                            <el-option label="公共算法" value="common"></el-option>
                        </el-select>
                    </el-form-item>
                    <el-form-item label="算法名称" prop="algorithmId" v-if="algorithmName" style="display:inline-block;">
                        <el-select v-model="ruleForm.algorithmId" placeholder="请选择算法名称" v-loadmore='loadAlgorithmName'
                            @change="changeAlgorithmName" filterable remote :remote-method="remoteAlgorithm">
                            <el-option v-for="item in algorithmNameOption" :key="item.algorithmId"
                                :label="item.algorithmName" :value='item.algorithmId'>
                            </el-option>
                        </el-select>
                    </el-form-item>
                    <el-form-item label="算法版本" prop="algorithmVersion" v-if="algorithmVersion"
                        style="display:inline-block;">
                        <el-select v-model="ruleForm.algorithmVersion" placeholder="请选择算法版本"
                            v-loadmore='loadAlgorithmVersion'>
                            <el-option v-for="item in algorithmVersionOption"
                                :key="item.algorithmDetail.algorithmId+item.algorithmDetail.algorithmVersion"
                                :label="item.algorithmDetail.algorithmVersion"
                                :value='item.algorithmDetail.algorithmVersion'>
                            </el-option>
                        </el-select>
                    </el-form-item>
                </div>
                <!-- 镜像三级框 -->
                <div>
                    <el-form-item label="镜像类型" prop="imageSource" :class="{inline:imageName}">
                        <el-select v-model="ruleForm.imageSource" @change="changeimageSource" placeholder="请选择">
                            <el-option label="我的镜像" value="my"></el-option>
                            <el-option label="预置镜像" value="pre"></el-option>
                            <el-option label="公共镜像" value="common"></el-option>
                        </el-select>
                    </el-form-item>
                    <el-form-item label="镜像名称" prop="imageId" v-if="imageName" style="display: inline-block;">
                        <el-select v-model="ruleForm.imageId" placeholder="请选择镜像名称" v-loadmore='loadImageName'
                            filterable remote :remote-method="remoteImage">
                            <el-option v-for="item in imageNameOption" :key="item.id"
                                :label="item.imageName+':'+item.imageVersion" :value="item.id">
                            </el-option>
                        </el-select>
                    </el-form-item>
                </div>
                <!-- 数据集三级框 -->
                <div>
                    <el-form-item label="数据集类型" prop="dataSetSource" :class="{inline:dataSetName}">
                        <el-select v-model="ruleForm.dataSetSource" @change="changedataSetSource" placeholder="请选择">
                            <el-option label="我的数据集" value="my"></el-option>
                            <el-option label="预置数据集" value="pre"></el-option>
                            <el-option label="公共数据集" value="common"></el-option>
                        </el-select>
                    </el-form-item>
                    <el-form-item label="数据集名称" prop="dataSetId" v-if="dataSetName" style="display: inline-block;">
                        <el-select v-model="ruleForm.dataSetId" placeholder="请选择数据集名称" v-loadmore='loadDataSetName'
                            @change="changeDataSetName" filterable remote :remote-method="remoteDataSet">
                            <el-option v-for="item in dataSetNameOption" :key="item.id+item.name" :label="item.name"
                                :value='item.id'>
                            </el-option>
                        </el-select>
                    </el-form-item>
                    <el-form-item label="数据集版本" prop="dataSetVersion" v-if="dataSetVersion"
                        style="display: inline-block;">
                        <el-select v-model="ruleForm.dataSetVersion" placeholder="请选择数据集版本"
                            v-loadmore='loadDataSetVersion'>
                            <el-option v-for="item in dataSetVersionOption" :key="item.datasetId+item.version"
                                :label="item.version" :value='item.version'>
                            </el-option>
                        </el-select>
                    </el-form-item>
                </div>
                <el-divider></el-divider>
                <el-form-item label="分布式" prop="distributed ">
                    <el-select v-model="ruleForm.isDistributed">
                        <el-option label="是" :value="true"></el-option>
                        <el-option label="否" :value="false"></el-option>
                    </el-select>
                </el-form-item>
                <div v-if="show">
                    <el-form-item label="运行命令" prop="command">
                        <el-input type="textarea" v-model="ruleForm.command"></el-input>
                    </el-form-item>
                    <el-form-item label="运行参数">
                        <div v-for="(item, index) in ruleForm.config[0].parameters" :key="index">
                            <el-form-item style="margin-bottom:10px">
                                <el-input placeholder="key" v-model="item.key" style="width: 20%;">
                                </el-input>
                                <span style="margin:0 10px 0 10px">=</span>
                                <el-input placeholder="value" v-model="item.value" style="width: 20%;">
                                </el-input>
                                <i class="el-icon-delete" @click="deleteItem(item, index)"></i>
                            </el-form-item>
                        </div>
                        <el-button @click="addItem" type="primary">增加</el-button>
                        <el-button type="text" @click="open" :disabled="showArg">预览</el-button>
                    </el-form-item>
                    <el-form-item label="资源规格" prop="resourceSpecId">
                        <el-select v-model="ruleForm.resourceSpecId" placeholder="请选择资源规格" style="width:35%">
                            <el-option v-for="(item,index) in resourceOptions" :key="index" :label="item.label"
                                :value="item.value">
                            </el-option>
                        </el-select>
                    </el-form-item>
                </div>
                <div v-if="!show">
                    <traningList @tableData="getTableData" :Table="table" :resource="resourceOptions"></traningList>
                </div>
            </el-form>
            <div slot="footer" class="dialog-footer">
                <el-button type="success" @click="traningAndSave('traning')" v-if="showTraning">开始训练</el-button>
                <el-button type="primary" @click="traningAndSave('save')" v-if="showTemplate">保存模板</el-button>
                <el-button type="warning" @click="cancel">取消</el-button>
            </div>
        </el-dialog>
    </div>
</template>

<script>
    import traningList from './traningList.vue'
    import { createTask, saveTemplate, getResourceList, editeTemplate } from '@/api/trainingManager'
    import { getPresetAlgorithmList, getPublicAlgorithmList, getMyAlgorithmList, getAlgorithmVersionList } from '@/api/modelDev'
    import { getMyImage, getPublicImage, getPreImage } from '@/api/imageManager'
    import { getMyDatasetList, getPublicDatasetList, getPresetDatasetList, getVersionList } from '@/api/datasetManager'
    import { getErrorMsg } from '@/error/index'
    export default {
        name: "dialogCreateForm",
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

            }
        },
        computed: {
            title: function () {
                switch (this.flag) {
                    case 2:
                        this.showTemplate = true
                        this.showTraning = false
                        return '创建任务模板'
                        break
                    default:
                        this.showTraning = true
                        this.showTemplate = true
                        return '创建训练任务'
                }
            },
            showArg: function () {
                let flag = true
                if (this.ruleForm.config[0].parameters.length === 0) {
                    return flag
                }
                else {
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
            },
        },
        watch: {
            'ruleForm.isDistributed': {
                deep: true,
                handler: function (newV, oldV) {
                    if (newV === true && oldV === false) { this.show = false; }
                    else if (newV === false && oldV === true) { this.show = true; }
                }

            },
        },
        data() {
            return {
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
                        ],

                    }],
                    resourceSpecId: "",
                    command: ''
                },
                CreateFormVisible: true,
                rules: {
                    name: [
                        { required: true, message: '请输入任务名称', trigger: 'blur' },
                    ],
                    childName: [
                        { required: true, message: '请输入任务名称', trigger: 'blur' },
                    ],
                    algorithmSource: [
                        { required: true, message: '请选择算法类型', trigger: 'change' }
                    ],
                    algorithmId: [
                        { required: true, message: '请选择算法名称', trigger: 'change' }
                    ],
                    algorithmVersion: [{ required: true, message: '请选择算法版本', trigger: 'change' }],
                    imageSource: [
                        { required: true, message: '请选择镜像类型', trigger: 'change' }
                    ],
                    imageId: [
                        { required: true, message: '请选择镜像名称', trigger: 'change' }
                    ],
                    dataSetSource: [
                        { required: true, message: '请选择数据集类型', trigger: 'change' }
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
                    ],
                },
                formLabelWidth: '120px',
                // 算法三级框
                algorithmChange: false,
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
                resourceOptions: [],
                data: {},
                temp: { algorithmId: '' },
                argument: '',


            }
        },
        created() {
            // 判断是创建训练任务还是创建模板还是创建模板
            // 1创建训练任务2创建训练模板3其他页面跳转    
            this.getResourceList()
            if (this.flag === 3) {
                let temp = JSON.parse(JSON.stringify(this.row))
                this.temp.algorithmId = temp.algorithmId
                this.ruleForm.algorithmSource = temp.algorithmSource
                this.ruleForm.algorithmVersion = temp.algorithmVersion
                this.ruleForm.algorithmId = temp.algorithmName
                this.algorithmName = true
                this.algorithmVersion = true
            }

        },
        directives: {
            loadmore: {
                inserted: function (el, binding) {
                    const SELECTWRAP_DOM = el.querySelector('.el-select-dropdown .el-select-dropdown__wrap');
                    SELECTWRAP_DOM.addEventListener('scroll', function () {
                        const CONDITION = this.scrollHeight - this.scrollTop <= this.clientHeight;
                        if (CONDITION) {
                            binding.value();
                        }
                    })
                }
            }
        },
        methods: {
            // 错误码
            getErrorMsg(code) {
                return getErrorMsg(code)
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

                    }

                })
            },
            addItem() {
                this.ruleForm.config[0].parameters.push({
                    key: '',
                    value: ''
                })
            },
            deleteItem(item, index) {
                this.ruleForm.config[0].parameters.splice(index, 1)

            },
            cancel() {
                let msg = ''
                if (this.flag == 2) { msg = '任务模板' }
                else { msg = '训练任务' }
                this.$confirm('此操作将取消创建' + msg + ', 是否继续?', '提示', {
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
                let isRepeat = function (arr) {
                    var hash = {};
                    for (var i in arr) {
                        if (hash[arr[i]])
                            return true;
                        hash[arr[i]] = true;
                    }

                    return false;
                }
                let data = []
                val.config.forEach(
                    item => {
                        data.push(item.name)
                    }
                )
                return isRepeat(data)

            },
            //判断子任务是否重名或者分布式任务是否存在子任务
            isSubmit(data) {
                let isSubmit = true
                let isRepeat = true
                if (data.config.length > 1) {
                    isRepeat = this.nameIsRepeat(data)
                }
                else { isRepeat = false }
                if (isRepeat) {
                    isSubmit = false
                    this.$message({
                        message: '子任务名称不能重复',
                        type: 'warning'
                    });
                }
                let isChildTask = false
                if (data.isDistributed) {
                    data.config.forEach(item => {
                        if (item.name === '' || item.command === '' || !item.name || !item.command) {
                            isSubmit = false
                            isChildTask = true
                        }
                    })
                }
                if (data.config.length === 0) {
                    isSubmit = false
                    isChildTask = true

                }
                if (isChildTask) {
                    this.$message({
                        message: '请完善子任务信息',
                        type: 'warning'
                    });
                }
                return isSubmit
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
                        }
                        var data = JSON.parse(JSON.stringify(this.ruleForm))
                        delete data.command;
                        delete data.resourceSpecId
                        delete data.algorithmSource
                        delete data.imageSource
                        delete data.dataSetSource
                        if (this.flag === 3) {
                            if (!this.algorithmChange) {
                                data.algorithmId = this.temp.algorithmId
                            }

                        }
                        if (this.isSubmit(data)) {
                            if (val === 'traning') {
                                createTask(data).then(response => {
                                    if (response.success) {
                                        this.$message({
                                            message: '创建成功',
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
                            if (val === 'save') {
                                saveTemplate(data).then(response => {
                                    if (response.success) {
                                        this.$message({
                                            message: '保存成功',
                                            type: 'success'
                                        });
                                        if (this.flag === 2) {
                                            this.$emit('confirm', false)
                                        }

                                    } else {
                                        this.$message({
                                            message: this.getErrorMsg(response.error.subcode),
                                            type: 'warning'
                                        });
                                    }

                                })

                            }


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
                if (val.length === 0) {
                    val = [{
                        parameters: [
                            {
                                key: "",
                                value: ""
                            }
                        ]
                    }]
                }
                this.ruleForm.config = val

            },
            // 算法三级对话框实现
            changealgorithmSource() {
                this.algorithmName = true
                this.algorithmNameCount = 1
                this.algorithmNameOption = [],
                    this.ruleForm.algorithmId = '',
                    this.ruleForm.algorithmVersion = ''
                this.algorithmChange = true
                this.getAlgorithmNameList()

            },
            changeAlgorithmName() {
                this.algorithmVersion = true
                this.algorithmVersionCount = 1
                this.algorithmVersionOption = [],
                this.ruleForm.algorithmVersion = ''
                this.getAlgorithmVersionList()
            },
            getAlgorithmNameList(searchKey) {
                if (this.ruleForm.algorithmSource === 'my') {
                    getMyAlgorithmList({ pageIndex: this.algorithmNameCount, pageSize: 10, searchKey: searchKey }).then(response => {
                        if (searchKey && searchKey.length != 0) {
                            this.algorithmNameOption = response.data.algorithms
                        }
                        else {
                            if (response.data.algorithms.length !== 0) {
                                this.algorithmNameOption = this.algorithmNameOption.concat(response.data.algorithms);
                                this.algorithmNameTotal = response.data.totalSize

                            }
                        }

                    })
                }
                if (this.ruleForm.algorithmSource === 'pre') {
                    getPresetAlgorithmList({ pageIndex: this.algorithmNameCount, pageSize: 10, searchKey: searchKey }).then(response => {
                        if (searchKey && searchKey.length != 0) {
                            this.algorithmNameOption = response.data.algorithms
                        }
                        else {
                            if (response.data.algorithms.length !== 0) {
                                this.algorithmNameOption = this.algorithmNameOption.concat(response.data.algorithms)
                                this.algorithmNameTotal = response.data.totalSize
                            }
                        }
                    })
                }
                if (this.ruleForm.algorithmSource === 'common') {
                    getPublicAlgorithmList({ pageIndex: this.algorithmNameCount, pageSize: 10, searchKey: searchKey }).then(response => {
                        if (searchKey && searchKey.length != 0) {
                            this.algorithmNameOption = response.data.algorithms
                        }
                        else {
                            if (response.data.algorithms.length !== 0) {
                                this.algorithmNameOption = this.algorithmNameOption.concat(response.data.algorithms);
                                this.algorithmNameTotal = response.data.totalSize
                            }
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
                    this.getAlgorithmNameList()
                }

            },
            loadAlgorithmVersion() {
                this.algorithmVersionCount = this.algorithmVersionCount + 1
                if (this.algorithmVersionOption.length < this.algorithmVersionTotal) {
                    this.getAlgorithmNameList()
                }
            },
            // 镜像二级对话框实现
            changeimageSource() {
                this.imageName = true
                this.imageNameCount = 1
                this.imageNameOption = [],
                    this.ruleForm.imageId = '',
                    this.getImageNameList()
            },
            getImageNameList(searchKey) {
                if (this.ruleForm.imageSource === 'my') {
                    getMyImage({ pageIndex: this.imageNameCount, pageSize: 10, imageStatus: 3, imageType: 2, searchKey: searchKey }).then(response => {
                        if (searchKey && searchKey.length != 0) {
                            let data = response.data.images;
                            let tableData = [];
                            data.forEach(item => {
                                tableData.push({ ...item.image, isShared: item.isShared })
                            })
                            this.imageNameOption = tableData
                        }
                        else {
                            if (response.data.images.length !== 0) {
                                let data = response.data.images;
                                let tableData = [];
                                this.imageNameTotal = response.data.totalSize
                                data.forEach(item => {
                                    tableData.push({ ...item.image, isShared: item.isShared })
                                })
                                this.imageNameOption = this.imageNameOption.concat(tableData)
                            }
                        }

                    })
                }
                if (this.ruleForm.imageSource === 'pre') {
                    getPreImage({ pageIndex: this.imageNameCount, pageSize: 10, imageStatus: 3, imageType: 2, searchKey: searchKey }).then(response => {
                        if (searchKey && searchKey.length != 0) {
                            this.imageNameOption = response.data.images
                        }
                        else {
                            if (response.data.images.length !== 0) { this.imageNameOption = this.imageNameOption.concat(response.data.images); this.imageNameTotal = response.data.totalSize }
                        }

                    })
                }
                if (this.ruleForm.imageSource === 'common') {
                    getPublicImage({ pageIndex: this.imageNameCount, pageSize: 10, imageStatus: 3, imageType: 2, searchKey: searchKey }).then(response => {
                        if (searchKey && searchKey.length != 0) {
                            this.imageNameOption = response.data.images
                        }
                        else {
                            if (response.data.images.length !== 0) { this.imageNameOption = this.imageNameOption.concat(response.data.images); this.imageNameTotal = response.data.totalSize }
                        }

                    })
                }

            },
            loadImageName() {
                this.imageNameCount = this.imageNameCount + 1
                if (this.imageNameOption.length < thisimageNameTotal) {
                    this.getImageNameList()
                }

            },
            // 数据集三级对话框
            changedataSetSource() {
                this.dataSetName = true
                this.dataSetNameCount = 1
                this.dataSetNameOption = [],
                    this.ruleForm.dataSetId = '',
                    this.ruleForm.dataSetVersion = ''
                this.getDataSetNameList()
            },
            changeDataSetName() {
                this.dataSetVersion = true
                this.dataSetVersionCount = 1
                this.dataSetVersionOption = [],
                    this.ruleForm.dataSetVersion = '',
                    this.getDataSetVersionList()
            },
            getDataSetNameList(searchKey) {
                if (this.ruleForm.dataSetSource === 'my') {
                    getMyDatasetList({ pageIndex: this.dataSetNameCount, pageSize: 10, searchKey: searchKey }).then(response => {
                        if (response.data.datasets === null) {
                            response.data.datasets = []
                        }
                        if (searchKey && searchKey.length != 0) {
                            this.dataSetNameOption = response.data.datasets
                        }
                        else {
                            this.dataSetNameOption = this.dataSetNameOption.concat(response.data.datasets)
                            this.dataSetNameTotal = response.data.totalSize
                        }


                    })
                }
                if (this.ruleForm.dataSetSource === 'pre') {
                    getPresetDatasetList({ pageIndex: this.dataSetNameCount, pageSize: 10, searchKey: searchKey }).then(response => {
                        if (response.data.datasets === null) {
                            response.data.datasets = []
                        }
                        if (searchKey && searchKey.length != 0) {
                            this.dataSetNameOption = response.data.datasets
                        }
                        else {
                            this.dataSetNameOption = this.dataSetNameOption.concat(response.data.datasets); this.dataSetNameTotal = response.data.totalSize
                        }
                    })
                }
                if (this.ruleForm.dataSetSource === 'common') {
                    getPublicDatasetList({ pageIndex: this.dataSetNameCount, pageSize: 10, searchKey: searchKey }).then(response => {
                        if (response.data.datasets === null) {
                            response.data.datasets = []
                        }
                        if (searchKey && searchKey.length != 0) {
                            this.dataSetNameOption = response.data.datasets
                        }
                        else {           
                            this.dataSetNameOption = this.dataSetNameOption.concat(response.data.datasets); this.dataSetNameTotal = response.data.totalSize

                        }

                    })
                }
            },
            getDataSetVersionList() {
                let data = {}
                data.datasetId = this.ruleForm.dataSetId
                data.pageIndex = this.dataSetVersionCount
                data.pageSize = 10
                data.status = 3
                getVersionList(data).then(response => {
                    if (response.data.versions !== null) { this.dataSetVersionOption = this.dataSetVersionOption.concat(response.data.versions); this.dataSetVersionTotal = response.data.totalSize }

                })
            },
            loadDataSetName() {
                this.dataSetNameCount = this.dataSetNameCount + 1
                if (this.dataSetNameOption.length < this.dataSetNameTotal) {
                    this.getDataSetNameList()
                }

            },
            loadDataSetVersion() {
                this.dataSetVersionCount = this.dataSetVersionCount + 1
                if (this.dataSetVersionOption.length < this.dataSetVersionTotal) {
                    this.getDataSetVersionList()
                }
            },
            //运行参数预览
            open() {
                this.argument = ''
                let data = JSON.parse(JSON.stringify(this.ruleForm.config[0].parameters))
                data.forEach(
                    item => {
                        this.argument += '--' + item.key + '=' + item.value + " "
                    }
                )
                this.$alert(this.argument, '运行参数', {
                    confirmButtonText: '确定',
                    callback: action => {
                    }
                });
            },
            //远程请求算法名称
            remoteAlgorithm(a) {
                this.getAlgorithmNameList(a)
            },
            // 远程请求镜像名称
            remoteImage(a) {
                this.getImageNameList(a)
            },
            //远程请求数据集名称
            remoteDataSet(a) {
                this.getDataSetNameList(a)
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