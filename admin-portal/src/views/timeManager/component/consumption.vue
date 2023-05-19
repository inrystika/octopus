<template>
    <div>
        <searchForm :search-form="searchForm" blur-name="任务名称 搜索" @searchData="getSearchData" />
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
            <el-table-column label="用户名" v-if="type=='group'">
                <template slot-scope="scope">
                    <el-tooltip trigger="hover" :content="scope.row.userEmail" placement="top">
                        <span>{{ scope.row.userName }}</span>
                    </el-tooltip>
                </template>
            </el-table-column>
            <el-table-column prop="title" label="任务名称"> </el-table-column>
            <el-table-column label="消费机时(h)" align="center">
                <template slot-scope="scope">
                    <span>{{ scope.row.amount }}</span>
                </template>
            </el-table-column>
            <el-table-column label="开始时间" align="center">
                <template slot-scope="scope">
                    <span>{{ scope.row.startedAt | parseTime }}</span>
                </template>
            </el-table-column>
            <el-table-column label="结束时间" align="center">
                <template slot-scope="scope">
                    <span>{{ scope.row.endedAt | parseTime }}</span>
                </template>
            </el-table-column>
            <el-table-column label="类型" align="center">
                <template slot-scope="scope">
                    <span>{{ changeType(scope.row.bizType) }}</span>
                </template>
            </el-table-column>
        </el-table>
        <div class="block">
            <el-pagination :current-page="searchData.pageIndex" :page-sizes="[10, 20, 50, 80]"
                :page-size="searchData.pageSize" :total="total" layout="total, sizes, prev, pager, next, jumper"
                @size-change="handleSizeChange" @current-change="handleCurrentChange" />
        </div>
    </div>
</template>
<script>
    import { getUserPay, getGroupPay } from '@/api/machineManager.js'
    import searchForm from '@/components/search/index.vue'
    export default {
        name: "UserMachineTime",
        components: {
            searchForm
        },
        props: {
            consumptionTabType: { type: Number, default: undefined }
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
                searchForm: [],
                type: '',
                searchKey: ''

            }
        },

        created() {
            this.getPay(this.searchData)
            if (this.consumptionTabType === 1) {
                this.type = 'user'
                this.searchForm = [{ type: 'InputSelectUser', label: '用户', prop: 'userId', placeholder: '请输入用户名' }]
            } else {
                this.type = 'group'
                this.searchForm = [{ type: 'InputSelectGroup', label: '群组', prop: 'spaceId', placeholder: '请输入群组名' }]
            }
        },

        methods: {
            changeType(value) {
                switch (value) {
                    case 0:
                        return ''
                    case 1:
                        return '训练'
                    case 2:
                        return 'NoteBook'
                    default:
                        return '模型部署'
                }
            },
            handleSizeChange(val) {
                this.searchData.pageSize = val
                this.searchData.searchKey = this.searchKey
                this.getPay(this.searchData)
            },
            handleCurrentChange(val) {
                this.searchData.pageIndex = val
                this.searchData.searchKey = this.searchKey
                this.getPay(this.searchData)
            },
            getPay(data) {
                if (!data) { data = { pageIndex: this.searchData.pageIndex, pageSize: this.searchData.pageSize } }
                if (data.time && data.time.length !== 0) {
                    data.startedAtGte = data.time[0] / 1000
                    data.startedAtLt = data.time[1] / 1000
                    delete data.time
                }
                if (this.consumptionTabType === 1) {
                    getUserPay(data).then(response => {
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
                    getGroupPay(data).then(response => {
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
                this.searchData.pageIndex = 1
                data = Object.assign(val, { pageIndex: 1, pageSize: this.searchData.pageSize })
                this.getPay(data)
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