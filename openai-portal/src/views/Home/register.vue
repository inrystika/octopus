<template>
    <div class="contain">
        <el-row type="flex" justify="center">
            <el-col :span="6">
                <div class="title" v-if="!show">登录</div>
                <div class="title" v-if="show">注册</div>
            </el-col>
        </el-row>
        <el-row type="flex" justify="center">
            <el-col :span="6">
                <div class="grid-content bg-purple-dark">
                    <el-form ref="loginForm" :model="loginForm" :rules="rules" label-width="80px">
                        <el-form-item prop="fullName" label="姓名" v-if="show">
                            <el-input v-model="loginForm.fullName" type="text" auto-complete="off"
                                placeholder="请输入姓名" />
                        </el-form-item>
                        <el-form-item prop="gender" label="姓名" v-if="show">
                            <el-radio-group v-model="loginForm.gender">
                                <el-radio :label="1">男</el-radio>
                                <el-radio :label="2">女</el-radio>
                            </el-radio-group>
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
                                <el-col :span="18" :offset="9" v-if="!show">
                                    <el-button type="primary" @click="login()">绑定并登录</el-button>
                                </el-col>
                                <el-col :span="18" :offset="9" v-if="show">
                                    <el-button type="primary" @click="register()">注册并登录</el-button>
                                </el-col>
                            </el-row>
                        </el-form-item>
                        <el-form-item>
                            <el-row type="flex" :gutter="20" v-if="hidden">
                                <el-col :span="18" :offset="9">
                                    <el-link type="primary" @click="goRegister">未注册?点击注册</el-link>
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
    import { login } from '@/api/Home.js'
    import { GetUrlParam } from '@/utils/index.js'
    import { setToken, getToken } from '@/utils/auth'
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
                hidden: true,
                loginForm: {
                    fullName: '',
                    username: undefined,
                    password: undefined,
                    gender: undefined,
                    bind: { platform: '', userId: '', userName: '' }
                },
                rules: {
                    username: [{ required: true, message: "请输入用户账号", trigger: "blur" },
                    { validator: checkEmail, trigger: "blur" }
                    ],
                    password: [{ required: true, message: '请输入用户密码', trigger: 'blur' }],
                    fullName: [{ required: true, message: '请输入姓名', trigger: 'blur' }],
                    gender: [
                        { required: true, message: '请选择性别', trigger: 'change' }
                    ],
                },
            }
        },
        created() {
            this.getThirdInfo()
            this.loginForm.bind.platform = sessionStorage.getItem("platform")
            this.loginForm.bind.userName = sessionStorage.getItem("thirdUserName")
            this.loginForm.bind.userId = sessionStorage.getItem("thirdUserId")
            console.log(this.loginForm)
        },
        computed: {
        },
        methods: {
            getThirdInfo() {
                if (GetUrlParam('token') !== '') {
                    setToken(GetUrlParam('token'))
                    console.log(getToken(),'token')
                    // this.$router.push({ path: '/index' })
                }
                else {
                    sessionStorage.setItem('thirdUserId', GetUrlParam('thirdUserId'))
                    if (GetUrlParam("thirdUserName")) {
                        let thirdUserName = GetUrlParam("thirdUserName")
                        sessionStorage.setItem('thirdUserName', thirdUserName)
                    }
                }

            },
            //去注册
            goRegister() {
                this.show = true
                this.hidden = false
            },
            // 注册并绑定
            register() {
                this.$refs['loginForm'].validate((valid) => {
                    if (valid) {
                        register(this.loginForm).then(
                            response => {
                                if (response.success) {
                                    this.$message({
                                        message: '注册成功',
                                        type: 'success'
                                    });
                                    setToken(getUrl("token", url))
                                    this.$router.push({ path: '/index' })
                                }
                                else {
                                    if (response.error.subcode == 16021) {
                                        this.$message({
                                            message: this.getErrorMsg(response.error.subcode),
                                            type: 'warning'
                                        });
                                    }
                                    else {
                                        this.$message({
                                            message: this.getErrorMsg(response.error.subcode),
                                            type: 'warning'
                                        });
                                    }

                                }
                            }
                        )
                    } else {
                        return false;
                    }
                });
            },
            // 登录并绑定
            login() {
                let loginForm = JSON.parse(JSON.stringify(this.loginForm));
                delete loginForm.fullName
                delete loginForm.gender
                login(loginForm).then((res) => {
                    if (res.success) {
                        setToken(getUrl("token", url))
                        this.$router.push({ path: '/index' })
                        this.$message({
                            message: '登录成功',
                            type: 'success'
                        });
                    } else {
                        this.$message({
                            message: this.getErrorMsg(res.error.subcode),
                            type: 'warning'
                        });
                        this.$router.push({ path: '/' })
                    }
                }).catch(() => {
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