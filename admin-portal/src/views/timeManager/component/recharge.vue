<template>
    <div>
        <searchForm :search-form="searchForm" blur-name="充值说明 搜索" @searchData="getSearchData" />
        <el-table :data="tableData" style="width: 100%;font-size: 15px;"
            :header-cell-style="{'text-align':'left','color':'black'}" :cell-style="{'text-align':'left'}">
            <el-table-column :label="type=='user'?'用户名':'群组名'" align="center">
                <template slot-scope="scope">
                    <el-tooltip trigger="hover" :content="scope.row.userEmail" placement="top">
                        <span v-if="type=='user'">{{ scope.row.userName }}</span>
                    </el-tooltip>
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
            <el-pagination :current-page="searchData.pageIndex" :page-sizes="[10, 20, 50, 80]"
                :page-size="searchData.pageSize" :total="total" layout="total, sizes, prev, pager, next, jumper"
                @size-change="handleSizeChange" @current-change="handleCurrentChange" />
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
                searchData: {
                    pageIndex: 1,
                    pageSize: 10
                },
                total: undefined,
                tableData: [],
                formLabelWidth: '120px',
                flag: undefined,
                form: { userName: '', userId: '', spaceName: '', spaceId: '', amount: undefined },
                searchForm: [


                ],
                type: '',
                searchKey: ''
                // timer: null

            }
        },

        created() {
            this.Recharge(this.searchData)
            if (this.rechangeTabType === 1) {
                this.type = 'user'
                this.searchForm = [{ type: 'InputSelectUser', label: '用户', prop: 'userId', placeholder: '请输入用户名' }]
            } else {
                this.type = 'group',
                    this.searchForm = [{ type: 'InputSelectGroup', label: '群组', prop: 'spaceId', placeholder: '请输入群组名' }]
            }
        },
        methods: {
            handleSizeChange(val) {
                this.searchData.pageSize = val
                this.searchData.searchKey=this.searchKey
                this.Recharge(this.searchData)
            },
            handleCurrentChange(val) {
                this.searchData.pageIndex = val
                this.searchData.searchKey=this.searchKey
                this.Recharge(this.searchData)
            },
            Recharge(data) {
                if (!data) { data = { pageIndex: this.searchData.pageIndex, pageSize: this.searchData.pageSize } }
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
                data = Object.assign(val, { pageIndex: 1, pageSize: this.searchData.pageSize })
                this.Recharge(data)
                if (val.searchKey) {
                    this.searchKey = val.searchKey
                }

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