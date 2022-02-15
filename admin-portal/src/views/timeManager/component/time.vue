<template>
    <div>
        <div>
            <el-select v-model="userId" filterable :filter-method="getUserOptions" v-loadmore="loadUserName"
                @focus='userClick' v-if="type=='user'" placeholder="用户">
                <el-option v-for="op in userOptions" :key="op.id" :label="op.fullName+'('+op.email+')'"
                    :value="op.id" />
            </el-select>
            <el-select v-model="spaceId" v-loadmore="loadGroupName"
            @focus='groupClick' v-if="type=='group'" placeholder="群组">
                <el-option v-for="op in groupOptions" :key="op.id" :label="op.name" :value="op.id" />
            </el-select>
            <el-button type="primary" @click="reset" class="reset">重置</el-button>
            <el-button type="primary" @click="search" class="reset">搜索</el-button>
        </div>
        <el-table :data="tableData" style="width: 100%;font-size: 15px;"
            :header-cell-style="{'text-align':'left','color':'black'}" :cell-style="{'text-align':'left'}">
            <el-table-column :label="type==='user'?'用户名称':'群组名称'" align="center">
                <template slot-scope="scope">
                    <el-tooltip trigger="hover" :content="scope.row.userEmail" placement="top">
                        <span v-if="type==='user'" style="margin-left: 10px">{{ scope.row.userName }}</span>
                    </el-tooltip>
                    <span v-if="type==='group'" style="margin-left: 10px">{{ scope.row.spaceName }}</span>
                </template>
            </el-table-column>
            <el-table-column label="当前机时剩余(小时)" align="center">
                <template slot-scope="scope">
                    <span style="margin-left: 10px">{{ scope.row.amount }}</span>
                </template>
            </el-table-column>
            <el-table-column label="操作" align="center">
                <template slot-scope="scope">
                    <el-button type="text" @click="addTime(scope.row)">增加机时</el-button>
                    <el-button type="text" @click="deleteTime(scope.row)">减少机时</el-button>
                </template>
            </el-table-column>
        </el-table>
        <div class="block">
            <el-pagination :current-page="searchData.pageIndex" :page-sizes="[10, 20, 50, 80]"
                :page-size="searchData.pageSize" :total="total" layout="total, sizes, prev, pager, next, jumper"
                @size-change="handleSizeChange" @current-change="handleCurrentChange" />
        </div>
        <!-- 对话框 -->
        <el-dialog :title="title" :visible.sync="dialogFormVisible" width="25%" :close-on-click-modal="false">
            <el-form :model="form">
                <el-form-item v-if="timeTabType===1" label="用户名称" :label-width="formLabelWidth">
                    <span>{{ form.userName }}</span>
                </el-form-item>
                <el-form-item v-if="timeTabType===2" label="群组名称" :label-width="formLabelWidth">
                    <span>{{ form.spaceName }}</span>
                </el-form-item>
                <el-form-item v-if="flag===0" label="增加机时" :label-width="formLabelWidth">
                    <el-input v-model="form.amount" autocomplete="off" style="width: 40%;"
                        onkeyup="this.value = this.value.replace(/[^\d.]/g,'');" />
                    <span>小时</span>
                </el-form-item>
                <el-form-item v-if="flag===1" label="减少机时" :label-width="formLabelWidth">
                    <el-input v-model="form.amount" autocomplete="off" style="width: 40%;"
                        onkeyup="this.value = this.value.replace(/[^\d.]/g,'');" />
                    <span>小时</span>
                </el-form-item>
                <el-form-item label="充值说明" :label-width="formLabelWidth">
                    <el-input v-model="form.title" style="width: 40%;"></el-input>
                </el-form-item>
            </el-form>
            <div slot="footer" class="dialog-footer">
                <el-button @click="cancel">取 消</el-button>
                <el-button type="primary" @click="confirm" v-preventReClick>确 定</el-button>
            </div>
        </el-dialog>
    </div>
</template>
<script>
    import { groupList, userList, groupRecharge, userRecharge } from '@/api/machineManager.js'
    import { getUserList, getGroupList } from '@/api/userManager.js'
    export default {
        name: "UserMachineTime",
        props: {
            timeTabType: { type: Number, default: undefined }
        },
        data() {
            return {
                searchData: {
                    pageIndex: 1,
                    pageSize: 10
                },
                total: undefined,
                tableData: [],
                dialogFormVisible: false,
                formLabelWidth: '120px',
                flag: undefined,
                form: { userName: '', userId: '', spaceName: '', spaceId: '', amount: undefined, title: '' },
                searchForm: [{ type: 'Input', label: 'ID', prop: 'id', placeholder: '请输入ID' }],
                type: '',
                userId: '',
                spaceId: '',
                userOptions: [],
                groupOptions: [],
                usersCount: 1,
                usersTotal: undefined,
                groupCount: 1,
                groupTotal: undefined,
                userTemp: '',
                groupTemp: ''

            }
        },
        computed: {
            title: function () {
                if (this.flag === 0) {
                    return '增加机时'
                }
                if (this.flag === 1) {
                    return '减少机时'
                }
            }
        },
        created() {
            this.getTime()
            if (this.timeTabType === 1) {
                this.type = 'user'
            } else {
                this.type = 'group'
            }

        },
        methods: {
            handleSizeChange(val) {
                this.searchData.pageSize = val
                let data = {}
                if (this.type == 'user') {
                    data = Object.assign(this.searchData, { userId: this.userId })
                }
                
                else { data = Object.assign(this.searchData, { spaceId: this.spaceId }) }
                this.getTime(data)
            },
            handleCurrentChange(val) {
                this.searchData.pageIndex = val
                let data = {}
                if (this.type == 'user') {
                    data = Object.assign(this.searchData, { userId: this.userId })
                }
                else { data = Object.assign(this.searchData, { spaceId: this.spaceId }) }
                this.getTime(data)
            },
            getTime(data) {
                if (!data) { data = { pageIndex: this.searchData.pageIndex, pageSize: this.searchData.pageSize } }
                if (this.timeTabType === 1) {
                    userList(data).then(response => {
                        if (response.success) {
                            this.total = parseInt(response.data.totalSize)
                            this.tableData = response.data.billingUsers
                        } else {
                            this.$message({
                                message: this.getErrorMsg(response.error.subcode),
                                type: 'warning'
                            });
                        }
                    })
                } else {
                    groupList(data).then(response => {
                        if (response.success) {
                            this.total = parseInt(response.data.totalSize)
                            this.tableData = response.data.billingSpaces
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
                data = Object.assign(val, { pageIndex: this.searchData.pageIndex, pageSize: this.searchData.pageSize })
                if (this.timeTabType === 1) {
                    data.userId = data.id
                } else {
                    data.spaceId = data.id
                }
                delete data.id
                this.getTime(data)
            },
            addTime(val) {
                this.dialogFormVisible = true
                this.form.amount = ''
                this.form.title = ''
                if (this.timeTabType === 1) {
                    this.form.userName = val.userName; this.form.userId = val.userId
                } else {
                    this.form.spaceName = val.spaceName; this.form.spaceId = val.spaceId
                }
                this.flag = 0
            },
            deleteTime(val) {
                this.dialogFormVisible = true
                this.form.amount = ''
                this.form.title = ''
                if (this.timeTabType === 1) {
                    this.form.userName = val.userName; this.form.userId = val.userId
                } else {
                    this.form.spaceName = val.spaceName; this.form.spaceId = val.spaceId
                }
                this.flag = 1
            },
            cancel() { this.dialogFormVisible = false },
            confirm() {
                const data = JSON.parse(JSON.stringify(this.form))
                if (this.flag === 1) {
                    data.amount = -data.amount
                } else {
                    data.amount = +data.amount
                }
                if (this.timeTabType === 1) {
                    delete data.userName
                    delete data.spaceName
                    delete data.spaceId
                    userRecharge(data).then(response => {
                        if (response.success) {
                            this.$message({
                                message: '修改成功',
                                type: 'success'
                            });
                            this.getTime()
                        } else {
                            this.$message({
                                message: this.getErrorMsg(response.error.subcode),
                                type: 'warning'
                            });
                        }
                    })
                } else {
                    groupRecharge(data).then(response => {
                        delete data.userName
                        delete data.spaceName
                        delete data.userId
                        if (response.success) {
                            this.$message({
                                message: '修改成功',
                                type: 'success'
                            });
                            this.getTime()
                        } else {
                            this.$message({
                                message: this.getErrorMsg(response.error.subcode),
                                type: 'warning'
                            });
                        }
                    })
                }

                this.dialogFormVisible = false
            },
            getUserOptions(val) {
                this.usersCount = 1
                if (val != '') {
                    this.userTemp = val
                    this.userOptions = []
                    getUserList({
                        pageIndex: this.usersCount,
                        pageSize: 10,
                        searchKey: val
                    }).then(response => {
                        if (response.success) {
                            this.usersTotal = response.data.totalSize
                            this.userOptions = response.data.users
                        }
                    })
                }
                else { this.userOptions = [] }
            },
            groupClick() { this.getGroupOptions() },
            getGroupOptions() {
                this.groupCount = 1
                this.groupOptions = [{ name: '默认群组', id: 'default-workspace' }]
                getGroupList({
                    pageIndex: this.groupCount,
                    pageSize: 10,
                }).then(response => {
                    if (response.success) {
                        this.groupTotal = response.data.totalSize
                        this.groupOptions = this.groupOptions.concat(response.data.workspaces)
                    }
                })
            },
            loadUserName() {
                this.usersCount = this.usersCount + 1
                if (this.userOptions.length < this.usersTotal) {
                    getUserList({
                        pageIndex: this.usersCount,
                        pageSize: 10,
                        searchKey: this.userTemp
                    }).then(response => {
                        if (response.success) {
                            this.usersTotal = response.data.totalSize
                            this.userOptions = this.userOptions.concat(response.data.users)

                        }
                    })
                }
            },
            loadGroupName() {
                this.groupCount = this.groupCount + 1
                if (this.groupOptions.length < this.groupTotal + 1) {
                    getGroupList({
                        pageIndex: this.groupCount,
                        pageSize: 10,
                    }).then(response => {
                        if (response.success) {
                            this.groupTotal = response.data.totalSize
                            this.groupOptions = this.groupOptions.concat(response.data.workspaces)
                        }
                    })
                }
            },
            userClick() { this.userOptions = [] },
            search() {
                let data = {}
                if (this.type == 'user') {
                    data = { pageIndex: this.searchData.pageIndex, pageSize: this.searchData.pageSize, userId: this.userId }
                }
                else { data = { pageIndex: this.searchData.pageIndex, pageSize: this.searchData.pageSize, spaceId: this.spaceId } }
                this.getTime(data)
            },
            reset() {
                this.userId = ''
                this.spaceId = ''
            }

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

    .reset {
        margin-left: 10px;
    }
</style>