<template>
    <div>
        <el-dialog
            title="详情"
            width="80%"
            :visible.sync="CreateFormVisible"
            :before-close="handleDialogClose"
            :append-to-body="true"
            custom-class="dialog"
            :close-on-click-modal="false"
        >
            <el-tabs v-model="activeName" class="Wrapper" @tab-click="handleTabClick">
                <el-tab-pane label="任务简况" name="menu1">
                    <taskProfile v-if="tabRefresh.menu1" :row="data" />
                </el-tab-pane>
                <el-tab-pane label="任务日志" name="menu2">
                    <taskLog v-if="tabRefresh.menu2" :row="data" />
                </el-tab-pane>
                <el-tab-pane label="任务负载" name="menu3">
                    <taskLoad v-if="tabRefresh.menu3" :row="data" />
                </el-tab-pane>
                <el-tab-pane label="运行信息" name="menu4">
                    <taskInfo v-if="tabRefresh.menu4" :row="data" />
                </el-tab-pane>
            </el-tabs>
        </el-dialog>
    </div>
</template>
<script>
    import taskLoad from './taskLoad.vue'
    import taskLog from './taskLog.vue'
    import taskProfile from './taskProfile.vue'
    import taskInfo from './taskInfo.vue'
    export default {
        name: "detailDialog",
        components: { taskLoad, taskLog, taskProfile, taskInfo },
        props: {
            data: {
                type: Object,
                default: () => { }
            }

        },
        data() {
            return {
                CreateFormVisible: true,
                row: {},
                activeName: 'menu1',
                tabRefresh: {
                    menu1: true,
                    menu2: false,
                    menu3: false,
                    menu4: false
                }
            }
        },
        created() {
            this.row = this.data
        },
        beforeDestroy() {

        },
        methods: {
            handleDialogClose() {
                this.$emit('close', false)
            },
            handleTabClick(tab) {
                this.activeName = tab.name
                switch (this.activeName) {
                    case 'menu1':
                        this.switchTab('menu1')
                        break
                    case 'menu2':
                        this.switchTab('menu2')
                        break
                    case 'menu3':
                        this.switchTab('menu3')
                        break
                    case 'menu4':
                        this.switchTab('menu4')
                        break
                }
            },
            switchTab(tab) {
                for (const key in this.tabRefresh) {
                    if (key === tab) {
                        this.tabRefresh[key] = true
                    } else {
                        this.tabRefresh[key] = false
                    }
                }
            }
        }
    }
</script>
<style lang="scss">
    .dialog {
        min-height: 800px
    }
</style>