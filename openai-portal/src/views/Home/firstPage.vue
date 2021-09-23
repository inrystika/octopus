<template>
    <el-container>
        <el-aside>
            <div class="logo"></div>
        </el-aside>
        <el-main>
            <div class="content">
                <div class="login-container">
                    <div class="grid-content">
                        <el-form
                            ref="loginForm"
                            :model="loginForm"
                            :rules="rules"
                            status-icon
                            label-position="left"
                            label-width="0px"
                            class="demo-ruleForm login-page"
                          >
                            <div class="title"> <span class="welcome">欢迎使用</span><span class="octopus">启智章鱼</span></div>
                            <el-form-item prop="email">
                                <el-input
                                    v-model="loginForm.email"
                                    type="text"
                                    auto-complete="off"
                                    placeholder="请输入用户账号"
                                >
                                </el-input>
                            </el-form-item>
                            <el-form-item prop="password">
                                <el-input
                                    v-model="loginForm.password"
                                    type="password"
                                    auto-complete="off"
                                    placeholder="密码"
                                >
                                </el-input>
                            </el-form-item>
                            <el-form-item style="width:100%;">
                                <el-button type="primary" style="width:100%;" :loading="logining" @click="handleLogin">
                                    登录
                                </el-button>
                            </el-form-item>
                        </el-form>
                    </div>
                </div>
            </div>
        </el-main>
    </el-container>

</template>

<script>
    export default {
        data() {
            // 邮箱类型验证
            var checkEmail = (rule, value, callback) => {
                const regEmail = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
                if (regEmail.test(value)) {
                    return callback();
                }
                callback(new Error("请输入合法的邮箱"));
            };
            return {
                logining: false,
                loginForm: {
                    email: undefined,
                    password: undefined
                },
                rules: {
                    email: [{ required: true, message: "请输入用户账号", trigger: "blur" },
                    { validator: checkEmail, trigger: "blur" }
                    ],
                    password: [{ required: true, message: '请输入用户密码', trigger: 'blur' }]
                },
                checked: false
            }
        },
        watch: {
            $route: {
                handler: function(route) {
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
                            } else {
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
    .el-main {
        padding: 0px;
        width: 80%;
        min-width: 600px;
    }

    .el-aside {
        background-color: #1a1a23 !important;
        background: url('../../assets/octopus-login-pic.svg');
        /* background-position: 9vh 23vh; */
        background-position: 80px 200px;
        background-repeat: no-repeat;
        background-size: 100% 100%;
        width: 20%;
        min-width: 200px;
    }

    .content {
        text-align: center;
        min-height: 100vh;
        background: #e8edef;
        font-family: MicrosoftYaHei-Bold;
    }

    .logo {
        margin-top: 60px;
        height: 70px;
        background: url('../../assets/logo-w.svg');
        background-repeat: no-repeat;
        background-position-x: 20px;
        background-size: 80% 100%;
    }

    .title {
        font-size: 24px;
        text-align: left;
        margin-bottom: 66px;
    }

    .welcome {
        color: #996699;
    }

    .octopus {
        margin-left: 5px;
        font-weight: bold;
        color: #502374;
    }

    .login-container {
        width: 100%;
        height: 100vh;
        overflow: hidden;
        background: url('../../assets/octopus-login-bg.svg');
    }

    .login-page {
        margin-top: 100px;
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

    .grid-content {
        position: relative;
        top: 30%;
    }
</style>