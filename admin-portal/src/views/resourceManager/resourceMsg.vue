<template>
    <el-tabs v-model="activeName" class="Wrapper" @tab-click="handleClick">
        <el-tab-pane label="系统资源" name="first">
            <resource :resource-tab-type="1" />
        </el-tab-pane>
        <el-tab-pane label="自定义资源" name="second">
            <resource :resource-tab-type="2" />
        </el-tab-pane>
    </el-tabs>
</template>
<script>
    import resource from "./components/resource.vue";
    export default {
        components: {
            resource
        },
        data() {
            return {
                activeName: 'first'
            }
        },
        mounted() {
            window.addEventListener('beforeunload', e => {
                sessionStorage.clear()
            });

        },
        destroyed() {
            window.removeEventListener('beforeunload', e => {
                sessionStorage.clear()
            })
        },
        methods: {
            handleClick(tab, event) {

            }
        }
    }
</script>
<style lang="scss" scoped>
    .Wrapper {
        margin: 15px !important;
        background-color: #fff;
        padding: 20px;
        min-height: 900px;
    }
</style>