<template>
    <div>
        <el-dialog title="高级搜索" :visible.sync="dialogFormVisible" :close-on-click-modal="false">
            <el-form
                ref="searchForm"
                :inline="true"
                class="demo-form-inline"
                :label-position="labelPosition"
                label-width="100px"
            >
                <el-form-item v-for="item in searchForm" :key="item.props" :label="item.label">
                    <el-input
                        v-if="item.type==='Input'"
                        v-model="searchData[item.prop]"
                        :placeholder="item.placeholder"
                    />
                    <el-select v-if="item.type==='Select'" v-model="searchData[item.prop]">
                        <el-option v-for="op in item.options" :key="op.value" :label="op.label" :value="op.value" />
                    </el-select>
                    <el-date-picker
                        v-if="item.type==='Time'"
                        v-model="searchData[item.prop]"
                        type="datetimerange"
                        range-separator="至"
                        start-placeholder="开始日期"
                        end-placeholder="结束日期"
                        value-format="timestamp"
                        class="data"
                    />
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
                labelPosition: 'left'

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
            reset() { this.searchData = { searchKey: '' } }
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
    .data{margin-right:40px;}
</style>