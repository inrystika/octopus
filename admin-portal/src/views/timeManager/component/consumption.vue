<template>
    <div>
        <searchForm :search-form="searchForm" blur-name="任务名称 搜索" @searchData="getSearchData" />
        <el-table :data="tableData" style="width: 100%;font-size: 15px;"
            :header-cell-style="{'text-align':'left','color':'black'}" :cell-style="{'text-align':'left'}">
            <el-table-column :label="type=='user'?'用户名':'群组名'" align="center">
                <template slot-scope="scope">
                    <span v-if="type=='user'">{{ scope.row.userName }}</span>
                    <span v-if="type=='group'">{{ scope.row.spaceName }}</span>
                </template>
            </el-table-column>
            <el-table-column prop="userName" label="用户名" v-if="type=='group'"> </el-table-column>
            <el-table-column prop="title" label="任务名称"> </el-table-column>
            <el-table-column label="消费机时(h)" align="center">
                <template slot-scope="scope">
                    <span>{{ scope.row.amount }}</span>
                </template>
            </el-table-column>
            <el-table-column label="开始时间" align="center">
                <template slot-scope="scope">
                    <span>{{ parseTime(scope.row.createdAt) }}</span>
                </template>
            </el-table-column>
            <el-table-column label="结束时间" align="center">
                <template slot-scope="scope">
                    <span>{{ parseTime(scope.row.endedAt) }}</span>
                </template>
            </el-table-column>
            <el-table-column label="类型" align="center" v-if="type=='user'">
                <template slot-scope="scope">
                    <span>{{ scope.row.bizType==1?'训练':'notebook' }}</span>
                </template>
            </el-table-column>
        </el-table>
        <div class="block">
            <el-pagination :current-page="pageIndex" :page-sizes="[10, 20, 50, 80]" :page-size="pageSize" :total="total"
                layout="total, sizes, prev, pager, next, jumper" @size-change="handleSizeChange"
                @current-change="handleCurrentChange" />
        </div>
    </div>
</template>
<script>
    import { getUserPay, getGroupPay } from '@/api/machineManager.js'
    import searchForm from '@/components/search/index.vue'
    import { getErrorMsg } from '@/error/index'
    import { parseTime, formatDuring } from '@/utils/index'
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
                pageIndex: 1,
                pageSize: 10,
                total: undefined,
                tableData: [],
                formLabelWidth: '120px',
                flag: undefined,
                form: { userName: '', userId: '', spaceName: '', spaceId: '', amount: undefined },
                searchForm: [
                    // { type: 'Time', label: '开始时间', prop: 'time', placeholder: '请选择时间段' },
                    // { type: 'Input', label: '用户名', prop: 'userNameLike', placeholder: '请输入用户名' }
                ],
                type: ''

            }
        },

        created() {
            this.getPay()
            if (this.consumptionTabType === 1) {
                this.type = 'user'
            } else {
                this.type = 'group'
            }
        },

        methods: {
            // 错误码
            getErrorMsg(code) {
                return getErrorMsg(code)
            },
            handleSizeChange(val) {
                this.pageSize = val
                this.getPay()
            },
            handleCurrentChange(val) {
                this.pageIndex = val
                this.getPay()
            },
            getPay(data) {
                if (!data) { data = { pageIndex: this.pageIndex, pageSize: this.pageSize } }
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
            // 时间戳转换日期
            parseTime(val) {
                return parseTime(val)
            },

            getSearchData(val) {
                let data = {}
                data = Object.assign(val, { pageIndex: this.pageIndex, pageSize: this.pageSize })
                this.getPay(data)
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