<template>
    <div>
        <searchForm :search-form="searchForm" blur-name="充值说明 搜索" @searchData="getSearchData" />
        <el-table :data="tableData" style="width: 100%;font-size: 15px;"
            :header-cell-style="{'text-align':'left','color':'black'}" :cell-style="{'text-align':'left'}">
            <el-table-column :label="type=='user'?'用户名':'群组名'" align="center">
                <template slot-scope="scope">
                    <span v-if="type=='user'">{{ scope.row.userName }}</span>
                    <span v-if="type=='group'">{{ scope.row.spaceName }}</span>
                </template>
            </el-table-column>
            <el-table-column label="充值机时(h)" align="center">
                <template slot-scope="scope">
                    <span>{{ scope.row.amount }}</span>
                </template>
            </el-table-column>
            <el-table-column label="充值时间" align="center">
                <template slot-scope="scope">
                    <span>{{ scope.row.createdAt | parseTime }}</span>
                </template>
            </el-table-column>
            <el-table-column prop="title" label="充值说明"> </el-table-column>
        </el-table>
        <div class="block">
            <el-pagination :current-page="pageIndex" :page-sizes="[10, 20, 50, 80]" :page-size="pageSize" :total="total"
                layout="total, sizes, prev, pager, next, jumper" @size-change="handleSizeChange"
                @current-change="handleCurrentChange" />
        </div>
    </div>
</template>
<script>
    import { getUserRecharge, getGroupRecharge } from '@/api/machineManager.js'
    import searchForm from '@/components/search/index.vue'
    export default {
        name: "UserMachineTime",
        components: {
            searchForm
        },
        props: {
            rechangeTabType: { type: Number, default: undefined }
        },
        data() {
            return {
                pageIndex: 1,
                pageSize: 10,
                total: undefined,
                tableData: [],
                formLabelWidth: '120px',
                flag: undefined,
                form: { userName: '', userId: '', spaceName: '', spaceId: '', amount: undefined },
                searchForm: [


                ],
                type: ''
                // timer: null

            }
        },

        created() {
            this.Recharge()
            if (this.rechangeTabType === 1) {
                this.type = 'user'
                this.searchForm = [{ type: 'InputSelectUser', label: '用户名', prop: 'userId', placeholder: '请输入用户名' }]
            } else {
                this.type = 'group',
                    this.searchForm = [{ type: 'InputSelectGroup', label: '群组名', prop: 'spaceId', placeholder: '请输入群组名' }]
            }
        },
        methods: {
            handleSizeChange(val) {
                this.pageSize = val
                this.Recharge()
            },
            handleCurrentChange(val) {
                this.pageIndex = val
                this.Recharge()
            },
            Recharge(data) {
                if (!data) { data = { pageIndex: this.pageIndex, pageSize: this.pageSize } }
                if (this.rechangeTabType === 1) {
                    getUserRecharge(data).then(response => {
                        if (response.success) {
                            this.total = parseInt(response.data.totalSize)
                            this.tableData = response.data.records
                        } else {
                            this.$message({
                                message: this.getErrorMsg(response.error.subcode),
                                type: 'warning'
                            });
                        }
                    })
                } else {
                    getGroupRecharge(data).then(response => {
                        if (response.success) {
                            this.total = parseInt(response.data.totalSize)
                            this.tableData = response.data.records
                        } else {
                            this.$message({
                                message: this.getErrorMsg(response.error.subcode),
                                type: 'warning'
                            });
                        }
                    })
                }
            },
            getSearchData(val) {
                let data = {}
                data = Object.assign(val, { pageIndex: this.pageIndex, pageSize: this.pageSize })
                this.Recharge(data)
            },
        }
    }
</script>
<style lang="scss" scoped>
    .function {
        float: right;
        margin: 10px;
    }

    .block {
        float: right;
        margin: 20px;
    }
</style>