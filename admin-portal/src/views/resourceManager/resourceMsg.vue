<template>
    <el-tabs v-model="activeName" class="Wrapper" @tab-click="handleClick">
        <el-tab-pane label="系统资源" name="first">
            <resource :Type="1" />
        </el-tab-pane>
        <el-tab-pane label="自定义资源" name="second">
            <resource :Type="2" />
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
        methods: {
            handleClick(tab, event) {

            }
        }
    }
</script>
<style lang="scss" scoped>
    .Wrapper {
    margin: 15px!important;
    background-color:#fff;
    padding: 20px;
    min-height: 900px;
  }
</style>