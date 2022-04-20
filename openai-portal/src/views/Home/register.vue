<template>
    <div class="contain">
        <el-row type="flex" justify="center">
            <el-col :span="6">
                <div class="title">注册</div>
            </el-col>
        </el-row>
        <el-row type="flex" justify="center">
            <el-col :span="6">
                <div class="grid-content bg-purple-dark">
                    <el-form ref="loginForm" :model="loginForm" :rules="rules" label-width="80px">
                        <el-form-item prop="fullName" label="姓名">
                            <el-input v-model="loginForm.fullName" type="text" auto-complete="off"
                                placeholder="请输入姓名" />
                        </el-form-item>
                        <el-form-item prop="username" label="邮箱">
                            <el-input v-model="loginForm.username" type="text" auto-complete="off"
                                placeholder="请输入用户账号" />
                        </el-form-item>
                        <el-form-item prop="password" label="密码">
                            <el-input v-model="loginForm.password" type="password" auto-complete="off"
                                placeholder="密码" />
                        </el-form-item>
                        <el-form-item>
                            <el-row type="flex" :gutter="20">
                                <el-col :span="18" :offset="10">
                                    <el-button type="primary" @click="submitForm()">注册</el-button>
                                </el-col>
                                <el-col :span="6" :offset="2" class="login">
                                    <el-link type="primary" v-show="show" @click="login()">去登录</el-link>
                                </el-col>
                            </el-row>
                        </el-form-item>
                    </el-form>
                </div>
            </el-col>
        </el-row>
    </div>
</template>
<script>
    import { register } from '@/api/themeChange.js'
    import { mapGetters } from 'vuex'
    import store from '@/store'
    export default {
        data() {
            var checkEmail = (rule, value, callback) => {
                const regEmail = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
                if (regEmail.test(value)) {
                    return callback();
                }
                callback(new Error("请输入合法的邮箱"));
            };
            return {
                show: false,
                loginForm: {
                    fullName: '',
                    username: undefined,
                    password: undefined,
                    gender: 2,
                    bind: { platform: '', userId: '', userName: '' }
                },
                rules: {
                    username: [{ required: true, message: "请输入用户账号", trigger: "blur" },
                    { validator: checkEmail, trigger: "blur" }
                    ],
                    password: [{ required: true, message: '请输入用户密码', trigger: 'blur' }],
                    fullName: [{ required: true, message: '请输入姓名', trigger: 'blur' }]
                },
            }
        },
        created() {
            this.loginForm.bind.platform = sessionStorage.getItem("platform")
            this.loginForm.bind.userName = sessionStorage.getItem("thirdUserName")
            this.loginForm.bind.userId = sessionStorage.getItem("thirdUserId")
        },
        computed: {
            ...mapGetters([
                'platform',
                'userId',
                'userName'
            ])
        },
        methods: {
            submitForm() {
                this.$refs['loginForm'].validate((valid) => {
                    if (valid) {
                        register(this.loginForm).then(
                            response => {
                                if (response.success) {
                                    this.$message({
                                        message: '注册成功',
                                        type: 'success'
                                    });
                                    this.show = true
                                }
                                else {
                                    this.show = true
                                    this.$message({
                                        message: this.getErrorMsg(response.error.subcode),
                                        type: 'warning'
                                    });
                                }
                            }
                        )
                    } else {
                        return false;
                    }
                });
            },
            login() {
                this.$router.push({
                    path: "/login"
                })
            },
        }
    }
</script>
<style scoped>
    .title {
        text-align: center;
        font-weight: 800;
        margin: 150px 0 30px 80px;
    }

    .login {
        text-align: right;
    }

    .contain {
        min-width: 1400px;
    }
</style>