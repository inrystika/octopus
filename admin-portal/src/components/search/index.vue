<template>
    <div>
        <el-dialog title="高级搜索" :visible.sync="dialogFormVisible" :close-on-click-modal="false">
            <el-form ref="searchForm" :inline="true" class="demo-form-inline" :label-position="labelPosition"
                label-width="100px">
                <el-form-item v-for="item in searchForm" :key="item.props" :label="item.label">
                    <el-input v-if="item.type==='Input'" v-model="searchData[item.prop]"
                        :placeholder="item.placeholder" />
                    <!-- 用户可搜索下拉框 -->
                    <el-select v-if="item.type==='InputSelectUser'" v-model="searchData[item.prop]" filterable
                        :filter-method="getUserOptions" v-loadmore="loadUserName" @focus='userClick'>
                        <el-option v-for="op in userOptions" :key="op.id" :label="op.fullName+'('+op.email+')'"
                            :value="op.id" />
                    </el-select>
                    <!-- 群组可搜索下拉框 -->
                    <el-select v-if="item.type==='InputSelectGroup'" v-model="searchData[item.prop]"
                        v-loadmore="loadGroupName" @focus='groupClick'>
                        <el-option v-for="op in groupOptions" :key="op.id" :label="op.name" :value="op.id" />
                    </el-select>
                    <el-select v-if="item.type==='Select'" v-model="searchData[item.prop]">
                        <el-option v-for="op in item.options" :key="op.value" :label="op.label" :value="op.value" />
                    </el-select>
                    <el-date-picker v-if="item.type==='Time'" v-model="searchData[item.prop]" type="datetimerange"
                        range-separator="至" start-placeholder="开始日期" end-placeholder="结束日期" value-format="timestamp"
                        class="data" />
                </el-form-item>
            </el-form>
            <div class="buttonWrapper">
                <el-button type="primary" @click="reset">重置</el-button>
                <el-button type="primary" @click="onSubmit">搜索</el-button>
            </div>
        </el-dialog>
        <el-form :inline="true" class="demo-form-inline">
            <el-form-item>
                <el-input v-model.trim="searchData.searchKey" :placeholder="blurName" />
            </el-form-item>
            <el-form-item>
                <el-button type="primary" @click="onSubmit">搜索</el-button>
                <el-button v-if="searchForm.length!==0" type="text" @click="changeSearchType()">高级搜索</el-button>
            </el-form-item>
        </el-form>
    </div>
</template>
<script>
    import { getUserList, getGroupList } from '@/api/userManager.js'
    export default {
        props: {
            searchForm: { type: Array, default: () => [] },
            blurName: { type: String, default: '' }
        },
        data() {
            return {
                advanced: false,
                searchData: { searchKey: '' },
                dialogFormVisible: false,
                labelPosition: 'left',
                userOptions: [],
                groupOptions: [],
                usersCount: 1,
                usersTotal: undefined,
                groupCount: 1,
                groupTotal: undefined,
                userTemp: ''

            }
        },
        methods: {
            onSubmit() {
                if (this.dialogFormVisible) { this.searchData.searchKey = "" }
                if (!this.dialogFormVisible) { this.searchData = { searchKey: this.searchData.searchKey } }
                if (this.dialogFormVisible) { this.dialogFormVisible = !this.dialogFormVisible }
                this.$emit('searchData', this.searchData)
            },
            changeSearchType() {
                this.dialogFormVisible = !this.dialogFormVisible
            },
            reset() { this.searchData = { searchKey: '' }, this.userOptions = [], this.groupOptions = [] },
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
            userClick() { this.getUserOptions() },
            groupClick() { this.getGroupOptions() }
        }
    }
</script>
<style lang="scss" scoped>
    .buttonWrapper {
        text-align: center;
    }

    .el-select {
        width: 188px;
    }

    .el-select>.el-input {
        max-width: 185px !important;
    }

    .data {
        margin-right: 40px;
    }
</style>