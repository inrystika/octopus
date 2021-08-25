<template>
    <el-container>
        <el-main>
            <div class="content">
                <div><img src="../../assets/logo.svg" alt="" class="logo"></div>
                <div class="login-container">
                    <div class="grid-content">
                        <el-form :model="loginForm" :rules="rules" status-icon ref="loginForm" label-position="left"
                            label-width="0px" class="demo-ruleForm login-page">
                            <h3 class="title">系统登录</h3>
                            <el-form-item prop="email">
                                <el-input type="text" v-model="loginForm.email" auto-complete="off"
                                    placeholder="请输入管理员账号">
                                </el-input>
                            </el-form-item>
                            <el-form-item prop="password">
                                <el-input type="password" v-model="loginForm.password" auto-complete="off"
                                    placeholder="密码">
                                </el-input>
                            </el-form-item>
                            <!-- <el-checkbox v-model="checked" class="rememberme">记住密码</el-checkbox> -->
                            <el-form-item style="width:100%;">
                                <el-button type="primary" style="width:100%;" @click="handleLogin" :loading="logining">
                                    登录
                                </el-button>
                            </el-form-item>
                        </el-form>
                    </div>
                    <div class="note">Powered by Openl Octopus V2.1</div>
                    <div class="octopus"></div>
                </div>
            </div>
        </el-main>
    </el-container>

</template>

<script>
    export default {
        data() {
            // 邮箱类型验证        
            return {
                logining: false,
                loginForm: {
                    email: undefined,
                    password: undefined,
                },
                rules: {
                    email: [{ required: true, message: "请输入管理员账号", trigger: "blur" },
                    ],
                    password: [{ required: true, message: '请输入管理员密码', trigger: 'blur' }]
                },
                checked: false
            }
        },
        watch: {
            $route: {
                handler: function (route) {
                    this.redirect = route.query && route.query.redirect
                },
                immediate: true
            }
        },
        methods: {
            handleLogin() {
                this.$refs.loginForm.validate(valid => {
                    if (valid) {
                        this.loading = true
                        this.$store.dispatch('user/login', this.loginForm).then((res) => {
                            if (res === 'success') {
                                this.$router.push({ path: '/index' })
                                this.loading = false
                                this.$message({
                                    message: '登录成功',
                                    type: 'success'
                                });
                            }
                            else {
                                this.$message({
                                    message: '账号密码错误',
                                    type: 'warning'
                                });
                            }

                        }).catch(() => {
                            this.loading = false
                        })
                    } else {
                        console.log('error submit!!')
                        return false
                    }
                })
            }
        }
    };
</script>

<style scoped>
    .el-container {
        background-color: #e8edef;
        height: 100vh
    }

    .content {
        top:5%;
        position: relative;
        min-width: 800px;
        text-align: center;
    }

    .octopus {
        width: 100%;
        height: 280px;
        background-image: url('../../assets/background.svg');
        background-repeat: no-repeat;
        background-size: 300% 200%;
        background-position: 20% 45%;
        margin-bottom: 30px;
    }

    .logo {
        width: 200px;
        height: 100px;
    }

    .title {
        font-weight: bold;
        color: #996699;
        text-align: center;
    }

    .login-container {
        width: 100%;
        height: 100%;
        overflow: hidden;
    }

    .login-page {
        width: 400px;
        height: 320px;
        padding: 35px 35px 15px;
        background: #fff;
        border: 1px solid #996699;
        box-shadow: 0px 0px 15px #996699;
        margin: 0 auto;
        border-radius: 10px;


    }

    label.el-checkbox.rememberme {
        margin: 0px 0px 15px;
        text-align: left;
    }

    .bg-purple {
        width: 450px;
        height: 400px;
        background-image: url("../../assets/adminPic.svg");
        background-repeat: no-repeat;
        background-size: 100% 100%;
    }

    .note {
        font-weight: 600;
        color: #996699;
        margin-top: 40px;
    }
</style>