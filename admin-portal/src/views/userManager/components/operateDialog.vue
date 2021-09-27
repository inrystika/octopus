<template>
    <div>
        <el-dialog
            :title="userType==='user'?'重置密码':'编辑群组信息'"
            width="35%"
            :visible.sync="CreateFormVisible"
            :before-close="handleDialogClose"
            :close-on-click-modal="false"
        >
            <el-form ref="ruleForm" :model="ruleForm" :rules="rules" label-width="100px" class="demo-ruleForm">
                <el-form-item v-if="user" label="用户名称" :label-width="formLabelWidth">
                    <el-input v-model="ruleForm.fullName" disabled />
                </el-form-item>
                <el-form-item v-if="user" label="用户新密码" :label-width="formLabelWidth" prop="password">
                    <el-input v-model="ruleForm.password" />
                </el-form-item>
                <el-form-item v-if="user" label="密码确认" :label-width="formLabelWidth" prop="confirm">
                    <el-input v-model="ruleForm.confirm" />
                </el-form-item>
                <!-- <el-form-item label="验证码" :label-width="formLabelWidth" placeholder="请输入验证码" prop="code" v-if="user">
                    <el-input v-model="ruleForm.verifyCode" class="verifyCode"></el-input>
                    <VerificationCode :changeCode.sync='verifyCode'></VerificationCode>
                </el-form-item> -->
                <el-form-item v-if="group" label="群组名称" :label-width="formLabelWidth" prop="name">
                    <el-input v-model="ruleForm.name" disabled />
                </el-form-item>
                <el-form-item v-if="group" label="用户列表" :label-width="formLabelWidth" prop="userIds">
                    <el-select v-model="ruleForm.userIds" placeholder="请选择用户列表" multiple>
                        <el-option
                            v-for="item in userOptions"
                            :key="item.id"
                            :label="item.fullName + ' ' + item.email"
                            :value="item.id"
                        />
                    </el-select>
                </el-form-item>
                <el-form-item v-if="group" label="资源池" :label-width="formLabelWidth" prop="resourcePoolId">
                    <el-select v-model="ruleForm.resourcePoolId" placeholder="请选择资源池">
                        <el-option v-for="item in resourceOptions" :key="item.id" :label="item.name" :value="item.id" />
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
    import { editeUser, editeGroup, groupDetail, getUserList } from '@/api/userManager.js'
    import { getResourcePool } from '@/api/resourceManager.js'
    import { getErrorMsg } from '@/error/index'
    export default {
        name: "OperateDialog",
        props: {
            userType: {
                type: String,
                default: 'user'
            },
            row: {
                type: Object,
                default: () => { }
            },
            flag: {
                type: String,
                default: ""
            }
        },
        data() {
            return {
                fileList: [],
                ruleForm: {
                    fullName: '',
                    password: '',
                    fullname: '',
                    resourcePoolId: '',
                    name: '',
                    userIds: []
                },
                verifyCode: "",
                CreateFormVisible: true,
                user: false,
                group: false,
                resourceOptions: [],
                userOptions: [],
                rules: {
                    password: [
                        { required: true, message: '请输入密码', trigger: 'blur' }

                    ],
                    confirm: [
                        { required: true, message: '请再次输入密码', trigger: 'blur' }

                    ],
                    name: [
                        { required: true, message: '请输入用户名', trigger: 'blur' }

                    ],
                    userIds: [
                        { required: true, message: '请选择用户列表', trigger: 'change' }
                    ],
                    resourcePoolId: [
                        { required: true, message: '请选择资源池', trigger: 'change' }
                    ]
                },
                formLabelWidth: '120px',
                pageSize: 100,
                userCount: 1

            }
        },
        mounted() {
            if (this.userType === 'user') {
                this.user = true
                this.group = false
                this.ruleForm.fullName = this.row.fullName
                this.id = this.row.id
            } else {
                this.group = true
                this.user = false
                this.id = this.row.id
                groupDetail(this.id).then(
                    response => {
                        if (response.success) {
                            if (response.data && response.data.workspace !== null) {
                                this.ruleForm.name = response.data.workspace.name
                                this.ruleForm.resourcePoolId = response.data.workspace.resourcePoolId
                                if (response.data.users.length !== 0) {
                                    this.ruleForm.userName = []
                                    response.data.users.forEach(
                                        item => {
                                            this.ruleForm.userIds.push(item.id)
                                        }
                                    )
                                }
                            }
                        } else {
                            this.$message({
                                message: response.error.message,
                                type: 'error'
                            });
                        }
                    }
                )
                getResourcePool().then(response => {
                    if (response.success) {
                        if (response.data !== null && response.data.resourcePools !== null) {
                            this.resourceOptions = response.data.resourcePools
                        }
                    } else {
                        this.$message({
                            message: response.error.message,
                            type: 'error'
                        });
                    }
                })
                this.getUserList()
            }
        },
        beforeDestroy() {
            this.ruleForm = {}
        },
        methods: {
             // 错误码
             getErrorMsg(code) {
                return getErrorMsg(code)
            },
            cancel() {
                this.$emit('cancel', false)
            },
            confirm() {
                this.$refs['ruleForm'].validate((valid) => {
                    if (valid) {
                        if (this.userType === 'user') {
                            if (this.ruleForm.confirm === this.ruleForm.password) {
                                const data = { fullname: this.ruleForm.fullname, password: this.ruleForm.password, id: this.id }
                                editeUser(data).then(response => {
                                    if (response.success) {
                                        this.$message({
                                            message: '修改成功',
                                            type: 'success'
                                        });
                                    } else {
                                        this.$message({
                                            message: response.error.message,
                                            type: 'error'
                                        });
                                    }
                                })
                            } else {
                                this.$message({
                                    message: '输入密码不一致!',
                                    type: 'warning'
                                });
                            }
                        } else {
                            const data = { name: this.ruleForm.name, resourcePoolId: this.ruleForm.resourcePoolId, id: this.id, userIds: this.ruleForm.userIds }
                            editeGroup(data).then(response => {
                                if (response.success) {
                                    this.$message({
                                        message: '修改成功',
                                        type: 'success'
                                    });
                                } else {
                            this.$message({
                                message: this.getErrorMsg(response.error.subcode),
                                type: 'warning'
                            });
                        }
                            })
                        }
                        this.$emit('confirm', false)
                    } else {
                        console.log('error submit!!');
                        return false;
                    }
                });
            },
            handleDialogClose() {
                this.$emit('close', false)
            },
            getUserList() {
                getUserList({ pageSize: this.pageSize, pageIndex: this.userCount }).then(response => {
                    if (response.data !== null && response.data.users !== null && response.data.users.length) {
                        this.userOptions = this.userOptions.concat(response.data.users)
                        this.userCount = this.userCount + 1
                        this.getUserList()
                    }
                })
            }

        }
    }
</script>
<style lang="scss" scoped>
    .verifyCode {
        width: 25%;
        margin-right: 5px;
        position: relative;
        top: -15px
    }
</style>