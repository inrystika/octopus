<template>
    <div>
        <el-dialog
            :title="flag?'添加分布式任务':'编辑分布式任务'"
            width="50%"
            :visible.sync="CreateFormVisible"
            :before-close="handleDialogClose"
            append-to-body
            :close-on-click-modal="false"
        >
            <el-form
                ref="ruleForm"
                :model="ruleForm"
                :rules="rules"
                :label-width="formLabelWidth"
                class="demo-ruleForm"
            >
                <el-form-item label="任务名称" prop="name">
                    <el-input v-model="ruleForm.name" />
                </el-form-item>
                <el-form-item label="运行命令" prop="command">
                    <el-input v-model="ruleForm.command" type="textarea" />
                </el-form-item>
                <el-form-item label="运行参数">
                    <div v-for="(item, index) in ruleForm.parameters" :key="index">
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
                        <el-option
                            v-for="item in resourceOptions"
                            :key="item.id"
                            :label="item.label"
                            :value="item.value"
                        />
                    </el-select>
                </el-form-item>
                <el-form-item label="副本个数" prop="taskNumber">
                    <el-input v-model.number="ruleForm.taskNumber" />
                </el-form-item>
                <el-form-item label="最小副本成功数" prop="minSucceededTaskCount">
                    <el-input v-model.number="ruleForm.minSucceededTaskCount" />
                </el-form-item>
                <el-form-item label="最小副本失败数" prop="minFailedTaskCount">
                    <el-input v-model.number="ruleForm.minFailedTaskCount" />
                </el-form-item>
                <el-form-item label="是否是主任务" prop="isMainRole">
                    <el-select v-model="ruleForm.isMainRole" placeholder="是否是主任务">
                        <el-option label="是" :value="true" />
                        <el-option label="否" :value="false" />
                    </el-select>
                </el-form-item>
            </el-form>
            <div slot="footer" class="dialog-footer">
                <el-button @click="cancel">取 消</el-button>
                <el-button type="primary" @click="confirm">确 定</el-button>
            </div>
        </el-dialog>
    </div>
</template>

<script>
    import { getResourceList } from '@/api/trainingManager.js'
    export default {
        name: "DistributedTask",
        props: {
            row: {
                type: Object,
                default: () => { }
            },
            flag: {
                type: Boolean,
                default: false
            }
        },
        data() {
            return {
                ruleForm: {
                    name: '',
                    command: '',
                    resourceSpecId: '',
                    taskNumber: 1,
                    minSucceededTaskCount: 1,
                    minFailedTaskCount: 1,
                    isMainRole: '',
                    parameters: [{
                        key: "",
                        value: ""
                    }],
                },
                CreateFormVisible: true,
                resourceOptions: [],
                rules: {
                    name: [
                        { required: true, message: '请输入任务名称', trigger: 'blur' }

                    ],
                    command: [
                        { required: true, message: '请输入运行命令', trigger: 'blur' }

                    ],
                    resourceSpecId: [
                        { required: true, message: '请选择资源规格', trigger: 'change' }
                    ],
                    taskNumber: [
                        { required: true, message: '请输入副本个数', trigger: 'blur' },
                        { type: 'number', min: 1, message: '副本个数必须不小于1' }

                    ],

                    minSucceededTaskCount: [
                        { required: true, message: '请输入最小副本成功数', trigger: 'blur' },
                        { type: 'number', min: 1, message: '最小副本成功数必须不小于1' }

                    ],
                    minFailedTaskCount: [
                        { required: true, message: '请输入最小副本失败数', trigger: 'blur' },
                        { type: 'number', min: 1, message: '最小失败副本数数必须不小于1' }

                    ],
                    isMainRole: [
                        { required: true, message: '请选择是否为主任务', trigger: 'change' }
                    ]
                },
                formLabelWidth: '160px',
                argument: ''
            }
        },
        computed: {
            showArg: function() {
                let flag = true
                if (this.ruleForm.parameters.length === 0) {
                    return flag
                } else {
                    this.ruleForm.parameters.forEach(
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
        created() {
            this.ruleForm = this.row
            this.getResourceItem()
        },
        beforeDestroy() {
            this.ruleForm = {}
        },
        methods: {
            addItem() {
                this.ruleForm.parameters.push({
                    key: '',
                    value: ''
                })
            },
            deleteItem(item, index) {
                this.ruleForm.parameters.splice(index, 1)
            },
            cancel() {
                this.$emit('cancel', false)
            },
            confirm() {
                this.$refs['ruleForm'].validate((valid) => {
                    if (valid) {
                        if (this.ruleForm.minFailedTaskCount > this.ruleForm.taskNumber || this.ruleForm.minSucceededTaskCount > this.ruleForm.taskNumber) {
                            this.$message({
                                message: '最小副本成功数或者最小副本失败数不能大于副本数',
                                type: 'warning'
                            });
                            this.ruleForm.minFailedTaskCount = 1
                            this.ruleForm.minSucceededTaskCount = 1
                        } else {
                            this.$emit('subTasks', this.ruleForm)
                            this.$emit('confirm', false)
                            if (this.flag) {
                                this.$emit('flag', true)
                            } else { this.$emit('flag', false) }
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
            // 获取资源规格
            getResourceItem() {
              getResourceList(this.row.disResourcePool).then(response => {
                if (response.success) {
                        response.data.mapResourceSpecIdList.train.resourceSpecs.forEach(
                            item => {
                                this.resourceOptions.push({ label: item.name + ' ' + item.price + '机时/h', value: item.id })
                            }
                        )
                        this.ruleForm.resourceOptions = this.resourceOptions
                    } else {
                        this.$message({
                            message: this.getErrorMsg(response.error.subcode),
                            type: 'warning'
                        });
                    }
                })
            },
            // 运行参数预览
            open() {
                this.argument = ''
                const data = JSON.parse(JSON.stringify(this.ruleForm.parameters))
                if (data) {
                    data.forEach(
                        item => {
                            this.argument += '--' + item.key + '=' + item.value + ' '
                        }
                    )
                } else { this.argument = '' }
                this.$alert(this.argument, '运行参数', {
                    confirmButtonText: '确定',
                    callback: action => {
                    }
                });
            }

        }
    }
</script>
<style lang="scss" scoped>
</style>