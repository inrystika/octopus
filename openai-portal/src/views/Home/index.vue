<template>
    <div class="firstPage">
        <header1 />
        <banner1 />
        <login />
        <description />
        <banner2 class="banner" />
        <description2 />
    </div>
</template>

<script>
    import header1 from "./header1.vue";
    import banner1 from "./banner1.vue";
    import banner2 from "./banner2.vue";
    import login from "./login.vue";
    import description from './description.vue'
    import description2 from './description2.vue';

    export default {
        name: "FirstPage",
        components: {
            header1,
            banner1,
            banner2,
            login,
            description,
            description2
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