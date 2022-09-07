<template>
    <div>
        <searchForm :search-form="searchForm" class="searchForm" :blur-name="user?'用户/邮箱 搜索':'群组 搜索'"
            @searchData="getSearchData" />
        <div class="create">
            <el-button v-if="user" type="primary" @click="create">创建用户</el-button>
            <el-button v-if="group" type="primary" @click="create">创建群组</el-button>
        </div>
        <el-table :data="tableData" style="width: 100%;font-size: 15px;"
            :header-cell-style="{'text-align':'left','color':'black'}" :cell-style="{'text-align':'left'}">
            <el-table-column v-if="user" label="用户名称" align="center">
                <template slot-scope="scope">
                    <span>{{ scope.row.fullName }}</span>
                </template>
            </el-table-column>
            <el-table-column v-if="user" label="用户邮箱" align="center">
                <template slot-scope="scope">
                    <span>{{ scope.row.email }}</span>
                </template>
            </el-table-column>
            <el-table-column v-if="user" label="电话" align="center">
                <template slot-scope="scope">
                    <span>{{ scope.row.phone }}</span>
                </template>
            </el-table-column>
            <el-table-column v-if="group" label="群组名称" align="center">
                <template slot-scope="scope">
                    <span>{{ scope.row.name }}</span>
                </template>
            </el-table-column>
            <el-table-column v-if="group" label="资源池" align="center">
                <template slot-scope="scope">
                    <span>{{ scope.row.resourcePoolId }}</span>
                </template>
            </el-table-column>
            <el-table-column v-if="user" label="状态" align="center">
                <template slot-scope="scope">
                    <span>{{ scope.row.status===1?'已冻结':'已激活' }}</span>
                </template>
            </el-table-column>
            <el-table-column v-if="user" label="备注" align="center">
                <template slot-scope="scope">
                    <span>{{ scope.row.desc }}</span>
                </template>
            </el-table-column>
            <el-table-column label="创建时间" align="center">
                <template slot-scope="scope">
                    <span>{{ scope.row.createdAt | parseTime }}</span>
                </template>
            </el-table-column>
            <el-table-column label="修改时间" align="center">
                <template slot-scope="scope">
                    <span>{{ scope.row.updatedAt | parseTime }}</span>
                </template>
            </el-table-column>
            <el-table-column label="操作" align="center" width="300">
                <template slot-scope="scope">
                    <el-button v-if="user" type="text" @click="handleUserEdit(scope.row)">资源池</el-button>
                    <el-button v-if="scope.row.status===2 && user" type="text" @click="handleFreeze(scope.row)">冻结
                    </el-button>
                    <el-button v-if="scope.row.status===1 && user" type="text" @click="handleThaw( scope.row)">激活
                    </el-button>
                    <el-button v-if="user" type="text" @click="handleReset(scope.row)">编辑</el-button>
                    <el-button v-if="group" type="text" @click="handleEdit(scope.row)">编辑</el-button>
                    <!-- <el-button @click="handleDelete(scope.row)" type="text" v-if="group">删除</el-button> -->
                    <el-button type="text" @click="handleDetail(scope.row)">{{ user?'所属群组':'用户列表' }}</el-button>
                    <el-button v-if="user" type="text" @click="handleUserConfig(scope.row)">用户配置</el-button>
                </template>
            </el-table-column>
        </el-table>
        <div class="block">
            <el-pagination :current-page="searchData.pageIndex" :page-sizes="[10, 20, 50, 80]"
                :page-size="searchData.pageSize" :total="total" layout="total, sizes, prev, pager, next, jumper"
                @size-change="handleSizeChange" @current-change="handleCurrentChange" />
        </div>
        <!-- 新增对话框 -->
        <addDialog v-if="CreateVisible" :flag="flag" @cancel="cancel" @confirm="confirm" @close="close" />
        <!-- 创修改信息对话框 -->
        <operateDialog v-if="operateVisible" :row="row" :user-type="change" @cancel="cancel" @confirm="confirm"
            @close="close" />
        <!-- 用户配置对话框 -->
        <userConfig v-if="userConfigVisible" :conKey="conKey" :conValue="conValue" :row="row" @cancel="cancel"
            @confirm="confirm" @close="close">
        </userConfig>
        <!-- 资源池绑定对话框 -->
        <userEdit v-if="userEdit" :userResourcePoolList="userResourcePoolList" :row="row" @cancel="cancel"
            @confirm="confirm" @close="close">
        </userEdit>
        <!-- 详情对话框 -->
        <el-dialog :title="user?'用户名' + userName:'群组名' + groupName" :visible.sync="detailVisible" width="30%" center
            class="title" :close-on-click-modal="false">
            <div v-if="user">
                <el-tag v-for="item in detailData" :key="item.index" class="detailData">{{ item.name }}</el-tag>
            </div>
            <div v-if="group">
                <el-tag v-for="item in detailData" :key="item.index" class="detailData">{{ item.fullName }}</el-tag>
            </div>
            <span slot="footer" class="dialog-footer">
                <el-button type="primary" @click="detailVisible = false">确 定</el-button>
            </span>
        </el-dialog>

    </div>
</template>

<script>
    import { getUserList, getGroupList, freeze, activation, deleteGroup, userDetail, groupDetail } from '@/api/userManager.js'
    import { getResourcePool } from '@/api/resourceManager.js'
    import { getUserConfigKey, getUserConfig } from '@/api/userManager.js'
    import operateDialog from "./components/operateDialog.vue";
    import addDialog from "./components/addDialog.vue";
    import userConfig from "./components/userConfig.vue";
    import userEdit from "./components/userEdit.vue";
    import searchForm from '@/components/search/index.vue'
    export default {
        name: "UserList",
        components: {
            operateDialog,
            addDialog,
            searchForm,
            userConfig,
            userEdit
        },
        props: {
            userTabType: { type: Number, default: undefined }
        },
        data() {
            return {
                tableData: [],
                row: {},
                CreateVisible: false,
                operateVisible: false,
                detailVisible: false,
                userConfigVisible: false,
                userEdit: false,
                userResourcePoolList: [],
                show: false,
                user: false,
                group: false,
                total: undefined,
                change: undefined,
                flag: undefined,
                searchForm: {},
                // 用户id
                id: undefined,
                detailData: [],
                userName: '',
                groupName: '',
                // timer: null,
                searchData: {
                    pageIndex: 1,
                    pageSize: 10
                },
                conKey: [],
                conValue: {},
                time: null
            }
        },
        created() {
            if (this.userTabType === 1) {
                this.user = true; this.group = false
                this.change = 'user'
                this.searchForm = [
                    {
                        type: 'Select', label: '状态', prop: 'status', placeholder: '请选择状态',
                        options: [{ label: '已冻结', value: 1 }, { label: '已激活', value: 2 }]
                    }
                ]
                this.flag = 'user'

            } else {
                this.user = false
                this.group = true
                this.change = 'group'
                this.flag = 'group'
                this.searchForm = []
            }
            this.getList(this.searchData)

        },
        beforeDestroy() {
            clearTimeout(this.timer);
            this.timer = null;
        },
        methods: {
            handleUserEdit(row) {
                this.row = row
                getResourcePool().then(response => {
                    if (response.success) {
                        if (response.data !== null && response.data.resourcePools !== null) {
                            this.userEdit = true
                            this.userResourcePoolList = response.data.resourcePools
                        } else {
                            this.userResourcePoolList = []
                        }
                    } else {
                        this.$message({
                            message: this.getErrorMsg(response.error.subcode),
                            type: 'warning'
                        });
                    }
                })
            },
            handleFreeze(row) {
                freeze(row.id).then(response => {
                    if (response.success) {
                        this.$message({
                            message: '冻结成功',
                            type: 'success'
                        });
                        this.getList(this.searchData)
                    } else {
                        this.$message({
                            message: this.getErrorMsg(response.error.subcode),
                            type: 'warning'
                        });
                    }
                })
            },
            handleThaw(row) {
                activation(row.id).then(response => {
                    if (response.success) {
                        this.$message({
                            message: '激活成功',
                            type: 'success'
                        });
                        this.getList(this.searchData)
                    } else {
                        this.$message({
                            message: this.getErrorMsg(response.error.subcode),
                            type: 'warning'
                        });
                    }
                })
            },
            handleEdit(row) {
                this.row = row
                this.operateVisible = true
            },
            handleDelete(row) {
                deleteGroup(row.id).then(response => {
                    if (response.success) {
                        this.$message({
                            message: '删除成功',
                            type: 'success'
                        });
                        this.getList(this.searchData)
                    } else {
                        this.$message({
                            message: this.getErrorMsg(response.error.subcode),
                            type: 'warning'
                        });
                    }
                })
            },
            handleReset(row) {
                this.row = row
                this.operateVisible = true
            },
            handleUserConfig(row) {
                this.row = row
                this.getUserConfigKey(row)
                // this.userConfigVisible = true
            },
            handleDetail(row) {
                if (this.user) {
                    this.id = row.id
                    this.userName = row.fullName
                    userDetail(this.id).then(response => {
                        this.detailData = []
                        this.detailData = response.data.workspaces
                        this.detailData.unshift({ name: '默认群组' })
                        this.detailVisible = true
                    })
                } else {
                    this.id = row.id
                    this.groupName = row.name
                    groupDetail(this.id).then(response => {
                        if (response.success) {
                            if (response.data !== null && response.data.users.length !== 0) {
                                this.detailData = response.data.users
                                this.detailVisible = true
                            }
                        } else {
                            this.$message({
                                message: this.getErrorMsg(response.error.subcode),
                                type: 'warning'
                            });
                        }
                    })
                }
            },
            handleSizeChange(val) {
                this.searchData.pageSize = val
                this.getList(this.searchData)
            },
            handleCurrentChange(val) {
                this.searchData.pageIndex = val
                this.getList(this.searchData)
            },
            cancel(val) {
                this.CreateVisible = val
                this.operateVisible = val
                this.userConfigVisible = val
                this.userEdit = val
                this.getList(this.searchData)
            },
            confirm(val) {
                this.CreateVisible = val
                this.operateVisible = val
                this.userConfigVisible = val
                this.userEdit = val
                this.timer = setTimeout(this.getList, 500)

            },
            close(val) {
                this.CreateVisible = val
                this.operateVisible = val
                this.userConfigVisible = val
                this.userEdit = val
                this.getList(this.searchData)
            },
            create() {
                this.CreateVisible = true
            },
            getList(data) {
                if (!data) {
                    data = this.searchData
                }
                if (this.userTabType === 1) {
                    getUserList(data).then(response => {
                        if (response.success) {
                            this.total = response.data.totalSize
                            this.tableData = response.data.users
                        } else {
                            this.$message({
                                message: this.getErrorMsg(response.error.subcode),
                                type: 'warning'
                            });
                        }
                    })
                } else {
                    getGroupList(data).then(response => {
                        if (response.success) {
                            this.total = response.data.totalSize
                            this.tableData = response.data.workspaces
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
                this.searchData = { pageIndex: 1, pageSize: this.searchData.pageSize }
                this.searchData = Object.assign(val, this.searchData)
                this.getList(this.searchData)
            },
            // 获取用户配置信息
            getUserConfigKey(row) {
                getUserConfigKey().then(response => {
                    if (response.success) {
                        this.conKey = response.data.configKeys
                        getUserConfig(row.id).then(response => {
                            if (response.success) {
                                this.conValue = response.data.config
                                this.userConfigVisible = true
                            } else {
                                this.$message({
                                    message: this.getErrorMsg(response.error.subcode),
                                    type: 'warning'
                                });
                            }
                        })
                    } else {
                        this.$message({
                            message: this.getErrorMsg(response.error.subcode),
                            type: 'warning'
                        });
                    }
                })
            }
        }
    }
</script>
<style lang="scss" scoped>
    .block {
        float: right;
        margin: 20px;
    }

    .create {
        float: right;
    }

    .searchForm {
        display: inline-block;
    }

    .name {
        font-size: 14px;
        margin: 10px;
    }

    .detailData {
        margin: 5px;
    }

    .title {
        font-size: 16px;
        font-weight: 600;
    }
</style>