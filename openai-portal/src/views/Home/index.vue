<template>
    <div class="firstPage">
        <login />
    </div>
</template>

<script>
    import login from "./login.vue";

    export default {
        name: "FirstPage",
        components: {
            login,
        },
        data() {
            return {
                dialogFormVisible: false,
                ruleForm: {
                    pass: '',
                    name: ''

                },
                formLabelWidth: '120px'
            }
        },
        mounted() {
            window.addEventListener("beforeunload", (e) => {
                clearProgress()
            });
        },
        destroyed() {
            window.removeEventListener("beforeunload", (e) => {
                clearProgress()
            });
        },
        methods: {
            close() { this.dialogFormVisible = false }
        }

    }
</script>

<style lang="scss" scoped>
    .firstPage {
        min-width: 800px;
    }

    .loginButton {
        float: right;
        margin-right: 20px;
        margin-top: 15px;
    }

    .img {
        width: 40px;
        height: 40px;
        margin: 10px;
    }

    .el-dialog__header {
        text-align: center;
    }

    .el-form-item__label {
        text-align: center;
    }

    .banner {
        margin: 0 auto;
        width: 60%;
    }
</style>