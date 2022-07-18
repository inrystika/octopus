<template>
  <div class="login-container">
    <div class="login">
      <el-form
        ref="loginForm"
        :inline="true"
        :model="loginForm"
        :rules="loginRules"
        auto-complete="on"
        label-position="left"
        class="demo-form-inline"
      >
        <el-form-item class="button">
          <el-button :loading="loading" type="primary" @click.native.prevent="handleLogin">登录</el-button>
        </el-form-item>
        <el-form-item prop="password">
          <span class="svg-container">
            <i class=" el-icon-lock" />
          </span>
          <el-input
            :key="passwordType"
            ref="password"
            v-model="loginForm.password"
            :type="passwordType"
            placeholder="password"
            name="password"
            tabindex="2"
            auto-complete="on"
            @keyup.enter.native="handleLogin"
          />
          <span class=" show-pwd" @click="showPwd">
            <svg-icon :icon-class="passwordType === 'password' ? 'eye' : 'eye-open'" />
          </span>
        </el-form-item>
        <el-form-item prop="email">
          <span class="svg-container">
            <i class="el-icon-message" />
          </span>
          <el-input
            ref="email"
            v-model="loginForm.email"
            placeholder="email"
            name="email"
            type="text"
            tabindex="1"
            auto-complete="on"
          />
        </el-form-item>
      </el-form>
      <div class="title">
        <div>WELCOME</div>
        <div class="octopus">启智章鱼</div>
      </div>
    </div>
  </div>
</template>

<script>
  export default {
    name: 'Login',
    data() {
      var checkEmail = (rule, value, callback) => {
        const regEmail = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
        if (regEmail.test(value)) {
          return callback();
        }
        callback(new Error("请输入合法的邮箱"));
      };

      return {
        loginForm: {
          email: '',
          password: ''
        },
        loginRules: {
          email: [
            { required: true, message: "请输入邮箱", trigger: "blur" },
            { validator: checkEmail, trigger: "blur" }

          ], password: [
            { required: true, message: '请输入密码！', trigger: 'blur' },
            { min: 8, message: '密码长度不能小于8位', trigger: 'blur' }

          ]
        },
        loading: false,
        passwordType: 'password',
        redirect: undefined
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
      showPwd() {
        if (this.passwordType === 'password') {
          this.passwordType = ''
        } else {
          this.passwordType = 'password'
        }
        this.$nextTick(() => {
          this.$refs.password.focus()
        })
      },
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
              this.loading = false
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
  }
</script>

<style lang="scss">
  $light_gray:#fff;
  $cursor: #fff;
  $bg:#fff;
  $dark_gray:#889aa4;
  $light_gray:#eee;

  @supports (-webkit-mask: none) and (not (cater-color: $cursor)) {
    .login-container .el-input input {
      color: #fff;
    }
  }

  .login-container {
    .login {
      height: 200px;
      max-height: 200px;
      overflow: hidden;
      margin: 0 auto;

      .title {
        float: right;
        margin: 0px 50px 0 0;
        font-size: 1.5rem;
        font-family: Adobe Heiti Std;
        color: #d3e3ff;
        text-shadow: 0px 2px 2px rgb(64, 158, 255);

        .octopus {
          margin-top: 5px;
          margin-left: 30px;
          font-size: 1rem;
        }

      }

      .el-form {
        min-width: 1140px;
        padding: 80px 150px 0 150px
      }
    }

    .el-input {
      display: inline-block;
      height: 47px;
      width: 85%;

      input {
        background: transparent;
        border: 0px;
        -webkit-appearance: none;
        border-radius: 0px;
        padding: 12px 5px 12px 15px;
        color: #000;
        height: 47px;
        caret-color: #000;

        &:-webkit-autofill {

          box-shadow: 0 0 0px 1000px $bg inset;
          color: #000;
        }
      }
    }

    .el-form-item {
      border: 1px solid #000;
      border-radius: 5px;
      color: #454545;
      float: right;
    }

    .button {
      margin-top: 8px;
      background-color: #ffffff;
      border: 0px solid
    }

    min-height: 100%;
    width: 60%;
    margin: 0 auto;
    border-radius: 4px;
    border: 1px solid rgba(207, 206, 208, 0.8);
    box-shadow: 0px 6px 8px 0px rgb(141 141 141 / 24%);
    overflow: hidden;
    margin-top:10px;

    .login-form {
      position: relative;
      width: 520px;
      max-width: 100%;
      padding: 160px 35px 0;
      margin: 0 auto;
      overflow: hidden;

    }

    .svg-container {
      padding: 6px 5px 6px 15px;
      color: $dark_gray;
      vertical-align: middle;
      width: 30px;
      display: inline-block;
    }

    .show-pwd {
      position: absolute;
      right: 10px;
      top: 7px;
      font-size: 16px;
      color: $dark_gray;
      cursor: pointer;
      user-select: none;
    }
  }

  .el-form-item__content {
    min-width: 204px
  }
</style>