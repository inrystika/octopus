<template>
    <div>
        <el-dialog :title="userType==='user'?'编辑用户信息':'编辑群组信息'" width="35%" :visible.sync="CreateFormVisible"
            :before-close="handleDialogClose" :close-on-click-modal="false">
            <el-form ref="ruleForm" :model="ruleForm" :rules="rules" label-width="100px" class="demo-ruleForm">
                <el-form-item v-if="user" label="用户名称" :label-width="formLabelWidth">
                    <el-input v-model="ruleForm.fullName" disabled />
                </el-form-item>
                <el-form-item v-if="user&&password" label="用户新密码" :label-width="formLabelWidth" prop="password">
                    <el-input v-model="ruleForm.password" type="password" />
                </el-form-item>
                <el-form-item v-if="user&&password" label="密码确认" :label-width="formLabelWidth" prop="confirm">
                    <el-input v-model="ruleForm.confirm" type="password" />
                </el-form-item>
                <el-form-item v-if="user&&edite" label="电话" :label-width="formLabelWidth" prop="phone">
                    <el-input v-model="ruleForm.phone" />
                </el-form-item>
                <el-form-item label="外部存储" v-if="user&&edite" :label-width="formLabelWidth">
                    <el-radio-group v-model="mountExternalStorage">
                        <el-radio :label="true">允许</el-radio>
                        <el-radio :label="false">禁止</el-radio>
                    </el-radio-group>
                </el-form-item>
                <el-form-item label="备注" prop="desc" v-if="user&&edite" :label-width="formLabelWidth">
                    <el-input type="textarea" v-model="ruleForm.desc" maxlength="100" :show-word-limit="true">
                    </el-input>
                </el-form-item>
                <el-form-item v-if="group" label="群组名称" :label-width="formLabelWidth" prop="name">
                    <el-input v-model="ruleForm.name" disabled />
                </el-form-item>
                <el-form-item v-if="group" label="用户列表" :label-width="formLabelWidth" prop="userIds">
                    <el-select v-model="ruleForm.userIds" placeholder="请选择用户列表" multiple>
                        <el-option v-for="item in userOptions" :key="item.id" :label="item.fullName + ' ' + item.email"
                            :value="item.id" />
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
                <el-button type="primary" @click="confirm" v-preventReClick>确 定</el-button>
            </div>
        </el-dialog>
    </div>
</template>

<script>
    import { editUser, editGroup, groupDetail, getUserList } from '@/api/userManager.js'
    import { getResourcePool, getGroupResourcePool } from '@/api/resourceManager.js'
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
            },
            type: {
                type: String,
                default: 'edite'
            }
        },
        data() {
            var checkPhone = (rule, value, callback) => {
                if (value === '') {
                    callback();
                } else {
                    let reg = /^(13|14|15|17|18|19)[0-9]{9}$/
                    if (reg.test(value)) {
                        callback();
                    }
                    else {
                        callback("请输入正确手机号码");
                    }
                }
            };
            var validatePass = (rule, value, callback) => {
                if (value === '') {
                    callback(new Error('请输入密码'));
                } else {
                    if (this.ruleForm.confirm !== '') {
                        this.$refs.ruleForm.validateField('confirm');
                    }
                    callback();
                }
            };
            var validatePass2 = (rule, value, callback) => {
                if (value === '' || value == undefined) {
                    callback(new Error('请再次输入密码'));
                } else if (value !== this.ruleForm.password) {
                    callback(new Error('两次输入密码不一致!'));
                } else {
                    callback();
                }
            };
            return {
                fileList: [],
                ruleForm: {
                    fullName: '',
                    password: '',
                    fullname: '',
                    resourcePoolId: '',
                    name: '',
                    userIds: [],
                    email: undefined,
                    desc: '',
                    phone: ''
                },
                verifyCode: "",
                CreateFormVisible: true,
                user: false,
                group: false,
                resourceOptions: [],
                userOptions: [],
                rules: {
                    password: [
                        { required: true, message: '请输入密码', trigger: 'blur' },
                        { min: 8, message: '密码长度不得少于8位', trigger: 'blur' },
                        { validator: validatePass, trigger: 'blur' }

                    ],
                    confirm: [
                        { validator: validatePass2, trigger: 'blur' }
                    ],
                    name: [
                        { required: true, message: '请输入用户名', trigger: 'blur' }

                    ],
                    userIds: [
                        { required: true, message: '请选择用户列表', trigger: 'change' }
                    ],
                    resourcePoolId: [
                        { required: true, message: '请选择资源池', trigger: 'change' }
                    ],
                    phone: [
                        { required: false, message: '请输入电话' },
                        { validator: checkPhone, trigger: "blur" }
                    ],
                },
                formLabelWidth: '120px',
                pageSize: 100,
                userCount: 1,
                edite: false,
                password: false,
                mountExternalStorage: false

            }
        },
        mounted() {
            if (this.userType === 'user') {
                this.user = true
                this.group = false
                this.ruleForm.fullName = this.row.fullName
                this.ruleForm.phone = this.row.phone
                this.ruleForm.desc = this.row.desc
                this.id = this.row.id
                if (this.row.permission) {
                    this.mountExternalStorage = this.row.permission.mountExternalStorage
                } else { this.mountExternalStorage = false }

                if (this.type == 'password') {
                    this.password = true
                    this.edite = false
                } else {
                    this.password = false
                    this.edite = true
                }
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
                            this.resourceOptions = response.data.resourcePools.filter(item => {
                                if (!item.default) {
                                    return item
                                }
                            })
                        }
                        getGroupResourcePool().then(response => {
                            if (response.success) {
                                if (response.data.workspaces && response.data.workspaces.length != 0) {
                                    this.resourceOptions.forEach(
                                        item => {
                                            response.data.workspaces.forEach(
                                                Item => {
                                                    if (item.id == Item.resourcePoolId) {
                                                        item.flag = true
                                                    }
                                                }
                                            )
                                        }
                                    )
                                    this.resourceOptions = this.resourceOptions.filter(
                                        item => {
                                            if (!item.flag) {
                                                return item
                                            }
                                        }
                                    )
                                }
                            } else {
                                this.$message({
                                    message: this.getErrorMsg(response.error.subcode),
                                    type: 'warning'
                                });
                            }

                        })

                    } else {
                        this.$message({
                            message: this.getErrorMsg(response.error.subcode),
                            type: 'warning'
                        });
                    }
                })
                    ,
                    this.getUserList()
            }
        },
        beforeDestroy() {
            this.ruleForm = {}
        },
        methods: {
            cancel() {
                this.$emit('cancel', false)
            },
            confirm() {
                this.$refs['ruleForm'].validate((valid) => {
                    if (valid) {
                        if (this.userType === 'user') {
                            let data = {}
                            if (this.password) { data = { fullname: this.ruleForm.fullname, password: this.ruleForm.password, id: this.id } }
                            else {
                                data = { id: this.id, phone: this.ruleForm.phone.toString(), desc: this.ruleForm.desc, permission: { mountExternalStorage: this.mountExternalStorage } }
                            }
                            editUser(data).then(response => {
                                if (response.success) {
                                    this.$message({
                                        message: '修改成功',
                                        type: 'success'
                                    });
                                    this.$emit('confirm', false)
                                } else {
                                    this.$message({
                                        message: response.error.message,
                                        type: 'error'
                                    });
                                }
                            })
                        } else {
                            const data = { name: this.ruleForm.name, resourcePoolId: this.ruleForm.resourcePoolId, id: this.id, userIds: this.ruleForm.userIds }
                            editGroup(data).then(response => {
                                if (response.success) {
                                    this.$message({
                                        message: '修改成功',
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