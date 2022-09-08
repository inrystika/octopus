<template>
    <div>
        <el-dialog :title="flag==='user'?'添加用户':'添加群组'" width="35%" :visible.sync="CreateFormVisible"
            :before-close="handleDialogClose" :close-on-click-modal="false">
            <el-form ref="ruleForm" :model="ruleForm" :rules="rules" label-width="100px" class="demo-ruleForm">
                <el-form-item v-if="user" label="用户邮箱" :label-width="formLabelWidth" prop="email">
                    <el-input v-model="ruleForm.email" />
                </el-form-item>
                <el-form-item v-if="user" label="用户密码" :label-width="formLabelWidth" prop="password">
                    <el-input v-model="ruleForm.password" type="password" />
                </el-form-item>
                <el-form-item v-if="user" label="密码确认" :label-width="formLabelWidth" prop="confirm">
                    <el-input v-model="ruleForm.confirm" type="password" />
                </el-form-item>
                <el-form-item v-if="user" label="电话" :label-width="formLabelWidth" prop="phone">
                    <el-input v-model="ruleForm.phone" />
                </el-form-item>
                <el-form-item v-if="group" label="群组名称" :label-width="formLabelWidth" prop="name">
                    <el-input v-model.trim="ruleForm.name" />
                </el-form-item>
                <el-form-item v-if="group" label="用户列表" :label-width="formLabelWidth" prop="userIds">
                    <el-select v-model="ruleForm.userIds" v-loadmore="loadUser" placeholder="请选择用户列表" multiple>
                        <el-option v-for="item in userOptions" :key="item.id" :label="item.fullName + ' ' + item.email"
                            :value="item.id" />
                    </el-select>
                </el-form-item>
                <el-form-item v-if="group" label="资源池" :label-width="formLabelWidth" prop="resourcePoolId">
                    <el-select v-model="ruleForm.resourcePoolId" placeholder="请选择资源池">
                        <el-option v-for="item in resourceOptions" :key="item.id" :label="item.name" :value="item.id" />
                    </el-select>
                </el-form-item>
                <el-form-item v-if="user" label="姓名" :label-width="formLabelWidth" prop="fullname">
                    <el-input v-model.trim="ruleForm.fullname" />
                </el-form-item>
                <el-form-item label="备注" prop="desc" v-if="user" :label-width="formLabelWidth">
                    <el-input type="textarea" v-model="ruleForm.desc"></el-input>
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
    import { createUser, createGroup, getUserList } from '@/api/userManager.js'
    import { getResourcePool, getGroupResourcePool } from '@/api/resourceManager.js'
    export default {
        name: "CreateDialog",
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
        props: {
            flag: {
                type: String,
                default: ""
            }
        },
        data() {
            // 邮箱类型验证
            var checkEmail = (rule, value, callback) => {
                const regEmail = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
                if (regEmail.test(value)) {
                    return callback();
                }
                callback(new Error("请输入合法的邮箱地址"));
            }
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
                if (value === '') {
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
                    fullname: '',
                    password: '',
                    confirm: '',
                    phone: '',
                    resourcePoolId: '',
                    name: '',
                    userIds: [],
                    email: undefined,
                    desc: ''
                },
                CreateFormVisible: true,
                user: false,
                group: false,
                rules: {
                    fullname: [
                        { required: true, message: '请输入用户名称', trigger: 'blur' }

                    ],
                    email: [
                        { required: true, message: "请输入邮箱", trigger: "blur" },
                        { validator: checkEmail, trigger: "blur" }

                    ],
                    password: [
                        { required: true, message: '请输入密码', trigger: 'blur' },
                        { min: 8, message: '密码长度不得少于8位', trigger: 'blur' },
                        { validator: validatePass, trigger: 'blur' }

                    ],
                    confirm: [
                        { validator: validatePass2, trigger: 'blur' }
                    ],
                    phone: [
                        { required: false, message: '请输入电话' },
                        { validator: checkPhone, trigger: "blur" }
                    ],
                    name: [
                        { required: true, message: '请输入群组名', trigger: 'blur' }

                    ],
                    userIds: [
                        { required: true, message: '请选择用户列表', trigger: 'change' }
                    ],
                    resourcePoolId: [
                        { required: true, message: '请选择资源池', trigger: 'change' }
                    ]
                },
                formLabelWidth: '120px',
                userOptions: [],
                resourceOptions: [],
                pageSize: 10,
                userCount: 1,
                total: undefined
            }
        },
        mounted() {
            if (this.flag === 'user') {
                this.user = true
                this.group = false
            } else {
                this.group = true
                this.user = false
                this.getResourcePool()
                this.getUserList()
            }
        },
        beforeDestroy() {
            this.ruleForm = {}
        },
        methods: {
            getUserList() {
                getUserList({ pageSize: this.pageSize, pageIndex: this.userCount }).then(response => {
                    if (response.success) {
                        if (response.data !== null && response.data.users !== null) {
                            this.userOptions = this.userOptions.concat(response.data.users)
                            this.total = response.data.totalSize
                        }
                    } else {
                        this.$message({
                            message: this.getErrorMsg(response.error.subcode),
                            type: 'warning'
                        });
                    }
                })
            },
            getResourcePool() {
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
            },
            cancel() {
                this.$emit('cancel', false)
            },
            confirm() {
                this.$refs['ruleForm'].validate((valid) => {
                    if (valid) {
                        if (this.user) {
                            const data = { fullname: this.ruleForm.fullname, password: this.ruleForm.password, email: this.ruleForm.email, gender: 1, phone: this.ruleForm.phone.toString(), desc: this.ruleForm.desc }
                            createUser(data).then(response => {
                                if (response.success === true) {
                                    this.$message({
                                        message: '新增用户成功',
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
                        } else if (this.group) {
                            const data = { name: this.ruleForm.name, resourcePoolId: this.ruleForm.resourcePoolId, userIds: this.ruleForm.userIds }
                            createGroup(data).then(response => {
                                if (response.success === true) {
                                    this.$message({
                                        message: '新增群组成功',
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
            loadUser() {
                this.userCount = this.userCount + 1
                if (this.userOptions.length < this.total) {
                    this.getUserList()
                }
            }

        }
    }
</script>
<style lang="scss" scoped>
</style>