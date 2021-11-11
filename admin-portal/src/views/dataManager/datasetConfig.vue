<template>
    <div>
        <el-tag :key="tag.id" v-for="tag in dynamicTags" closable :disable-transitions="false" @close="handleClose(tag)"
            @click="changeValue(tag)">
            {{tag.typeDesc }}
        </el-tag>
        <el-input class="input-new-tag" v-if="inputVisible" v-model="inputValue" ref="saveTagInput" size="small"
            @keyup.enter.native="handleInputConfirm" @blur="handleInputConfirm">
        </el-input>
        <el-button v-else class="button-new-tag" size="small" @click="showInput">+ 新增类型</el-button>
    </div>
</template>
<script>
    import { datasetType, addDatasetType, deleteDatasetType, updateDatasetType } from "@/api/dataManager"
    import { getErrorMsg } from '@/error/index'
    export default {
        name: 'datasetConfig',
        data() {
            return {
                dynamicTags: [],
                inputVisible: false,
                inputValue: '',
                tempTag: '',
                // 是否改变原来的值
                isChange: false,
                tempId: ''
            };
        },
        created() {
            this.getType();
        },
        methods: {
            getErrorMsg(code) {
                return getErrorMsg(code)
            },
            handleClose(tag) {
                deleteDatasetType(tag.id).then(response => {
                    if (response.success) {
                        this.getType()
                    }
                    else {
                        this.$message({
                            message: this.getErrorMsg(response.error.subcode),
                            type: 'warning'
                        });
                    }
                })
            },
            showInput() {
                this.tempTag = ''
                this.inputVisible = true;
                this.inputValue = ''
                this.isChange = false
                this.$nextTick(_ => {
                    this.$refs.saveTagInput.$refs.input.focus();
                });
            },
            handleInputConfirm() {
                let inputValue = this.inputValue;
                inputValue = inputValue.replace(/^\s\s*/, '').replace(/\s\s*$/, '')
                if (this.isChange) {
                    updateDatasetType({ id: this.tempId, typeDesc: inputValue }).then(
                        response => {
                            if (response.success) {
                                this.getType()
                            }
                            else {
                                this.$message({
                                    message: this.getErrorMsg(response.error.subcode),
                                    type: 'warning'
                                });
                            }
                        }
                    )

                }
                // 点击添加时，追加
                if (inputValue) {
                    if (inputValue) {
                        addDatasetType(inputValue).then(response => {
                            if (response.success) {
                                this.getType()
                            }
                            else {
                                this.$message({
                                    message: this.getErrorMsg(response.error.subcode),
                                    type: 'warning'
                                });
                            }
                        })

                    }
                }
                this.inputVisible = false;
                this.inputValue = '';
            },
            getType() {
                datasetType({ pageIndex: 1, pageSize: 20 }).then(
                    response => {
                        if (response.success) {
                            this.dynamicTags = response.data.datasetTypes
                            if (this.dynamicTags == null) {
                                this.dynamicTags = []
                            }
                        }
                        else {
                            this.$message({
                                message: this.getErrorMsg(response.error.subcode),
                                type: 'warning'
                            });
                        }
                    }
                )
            },
            changeValue(tag) {
                this.inputVisible = true
                this.$nextTick(_ => {
                    this.$refs.saveTagInput.$refs.input.focus();
                });
                this.inputValue = tag.typeDesc
                this.tempTag = tag.typeDesc
                this.tempId = tag.id
                this.isChange = true
            }
        }
    }
</script>
<style>
    .el-tag+.el-tag {
        margin-left: 10px;
    }

    .button-new-tag {
        margin-left: 10px;
        height: 32px;
        line-height: 30px;
        padding-top: 0;
        padding-bottom: 0;
    }

    .input-new-tag {
        width: 90px;
        margin-left: 10px;
        vertical-align: bottom;
    }
</style>