<template>
    <div>
        <el-dialog title="详情" width="80%" :visible.sync="CreateFormVisible" :before-close="handleDialogClose"
            :append-to-body="true" custom-class="dialog" :close-on-click-modal="false">
            <el-tabs v-model="activeName" class="Wrapper" @tab-click="handleTabClick">
                <el-tab-pane label="部署调用" name="menu1">
                    <deployCall v-if="tabRefresh.menu1" :row="data" />
                </el-tab-pane>
                <el-tab-pane label="事件记录" name="menu2">
                    <deployRecord v-if="tabRefresh.menu2" :id="id" />
                </el-tab-pane>
            </el-tabs>
        </el-dialog>
    </div>
</template>
<script>
    import deployCall from './deployCall.vue'
    import deployRecord from './deployRecord.vue'
    export default {
        name: "DetailDialog",
        components: { deployCall, deployRecord },
        props: {
            row: {
                type: Object,
                default: () => { }
            }

        },
        data() {
            return {
                CreateFormVisible: true,
                activeName: 'menu1',
                id: '',
                tabRefresh: {
                    menu1: true,
                    menu2: false,
                    menu3: false,
                    menu4: false
                },
                data: {}
            }
        },
        mounted() {
            this.data = this.row
            this.id = this.row.id
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