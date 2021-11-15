<!-- <template>
    <div>
        <el-tag :key="tag.id" v-for="tag in dynamicTags" closable :disable-transitions="false" @close="handleClose(tag)"
            @click="changeValue(tag)">
            {{tag.typeDesc }}
        </el-tag>
        <el-input class="input-new-tag" v-if="inputVisible" v-model="inputValue" ref="saveTagInput" size="small"
            @keyup.enter.native="handleInputConfirm" @blur="handleInputConfirm">
        </el-input>
        <el-button v-else class="button-new-tag" size="small" @click="showInput">+ 新增类型</el-button>
        <star-input-tag v-model=" dynamicTags" theme="添加新标签" />
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
                tempId: '',

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
</style> -->
<template>
    <div>
        <el-tag v-for="(tag,index) in dynamicTags" :key="index" closable :disable-transitions="false"
            @click="editTag(tag,index)" @close="handleClose(tag)">
            <span v-if="index!=num"> {{tag.typeDesc }}</span>
            <input class="custom_input" type="text" v-model="inputValue" v-if="index==num" ref="editInput"
                @keyup.enter.native="handleInput(tag)" @blur="handleInput(tag)">
        </el-tag>
        <el-input class="input-new-tag" v-if="inputVisible" v-model="inputValue" ref="saveTagInput" size="small"
            @keyup.enter.native="handleInputConfirm" @blur="handleInputConfirm">
        </el-input>
        <el-button v-else class="button-new-tag" size="small" @click="showInput">{{'+ 新标签'}}</el-button>
    </div>
</template>
<script>
    import { datasetType, addDatasetType, deleteDatasetType, updateDatasetType } from "@/api/dataManager"
    import { getErrorMsg } from '@/error/index'
    export default {
        name: 'star-input-tag',
        created() {
            this.getType();
        },
        data() {
            return {
                inputVisible: false,
                inputValue: '',
                id: '',
                num: -1,
                dynamicTags: []
            }
        },

        methods: {
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
                this.inputVisible = true;
                this.$nextTick(_ => {
                    this.$refs.saveTagInput.$refs.input.focus();
                });

            },

            handleInputConfirm() {
                let inputValue = this.inputValue;
                inputValue = inputValue.replace(/^\s\s*/, '').replace(/\s\s*$/, '')
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
            editTag(tag, index) {
                this.num = index;
                this.$nextTick(_ => {
                    this.$refs.editInput[0].focus();
                });
                this.inputValue = tag.typeDesc;
                this.id = tag.id
            },
            handleInput(tag) {
                updateDatasetType({ id: this.id, typeDesc: this.inputValue }).then(
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
                        this.inputValue = '';
                        this.num = -1;
                    }
                )

            },
            getErrorMsg(code) {
                return getErrorMsg(code)
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
        }
    }
</script>
<style scoped lang="scss">
    .el-tag+.el-tag {
        margin-left: 15px;
    }

    .button-new-tag {
        margin-left: 15px;
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

    .custom_input {
        width: 80px;
        height: 16px;
        outline: none;
        border: transparent;
        background-color: transparent;
        font-size: 12px;
        color: #E6A23C;
    }
</style>