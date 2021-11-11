<template>
    <div>
        <div class="title">算法类型</div>
        <div>
            <el-tag :key="tag.id" v-for="tag in typeTags" closable :disable-transitions="false"
                @close="handleClose(tag,'type')" @click="changeValue(tag,'type')">
                {{tag.typeDesc }}
            </el-tag>
            <el-input class="input-new-tag" v-if="inputVisibleType" v-model="inputTypeValue" ref="saveTypeInput"
                size="small" @keyup.enter.native="handleInputConfirm" @blur="handleInputConfirm('type')">
            </el-input>
            <el-button v-else class="button-new-tag" size="small" @click="showInput('type')">+ 新增类型</el-button>
        </div>
        <el-divider></el-divider>
        <div class="title">算法框架</div>
        <div>
            <el-tag :key="tag.id" v-for="tag in frameTags" closable :disable-transitions="false"
                @close="handleClose(tag,'frame')" @click="changeValue(tag,'frame')">
                {{tag.frameworkDesc }}
            </el-tag>
            <el-input class="input-new-tag" v-if="inputVisibleFrame" v-model="inputFrameValue" ref="saveFrameInput"
                size="small" @keyup.enter.native="handleInputConfirm" @blur="handleInputConfirm('frame')">
            </el-input>
            <el-button v-else class="button-new-tag" size="small" @click="showInput('frame')">+ 新增类型</el-button>
        </div>
    </div>
</template>
<script>
    import { algorithmType, addAlgorithmType, deleteAlgorithmType, updateAlgorithmType, frameType, addFrameType, deleteFrameType, updateFrameType } from "@/api/modelDev"
    import { getErrorMsg } from '@/error/index'
    export default {
        name: 'datasetConfig',
        data() {
            return {
                typeTags: [],
                frameTags: [],
                inputVisibleType: false,
                inputVisibleFrame: false,
                inputTypeValue: '',
                tempType: '',
                // 是否改变原来的值
                isTypeChange: false,
                tempTypeId: '',
                inputFrameValue: false,
                tempFrame: "",
                tempFrameId: '',
                isFrameChange: false
            };
        },
        created() {
            this.getType();
            this.getFrameType()
        },
        methods: {
            getErrorMsg(code) {
                return getErrorMsg(code)
            },
            handleClose(tag, val) {
                if (val === 'type') {
                    deleteAlgorithmType(tag.id).then(response => {
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
                else {
                    deleteFrameType(tag.id).then(response => {
                        if (response.success) {
                            this.getFrameType()
                        }
                        else {
                            this.$message({
                                message: this.getErrorMsg(response.error.subcode),
                                type: 'warning'
                            });
                        }
                    })
                }

            },
            showInput(val) {
                if (val === 'type') {
                    this.tempType = ''
                    this.inputVisibleType = true;
                    this.inputTypeValue = ''
                    this.isTypeChange = false
                    this.$nextTick(_ => {
                        this.$refs.saveTypeInput.$refs.input.focus();
                    });
                }
                else {

                    this.tempFrame = ''
                    this.inputVisibleFrame = true;
                    this.inputFrameValue = ''
                    this.isFrameChange = false
                    this.$nextTick(_ => {
                        this.$refs.saveFrameInput.$refs.input.focus();
                    });

                }


            },
            handleInputConfirm(val) {
                if (val === 'type') {
                    let inputTypeValue = this.inputTypeValue;
                    inputTypeValue = inputTypeValue.replace(/^\s\s*/, '').replace(/\s\s*$/, '')
                    if (this.isTypeChange) {
                        updateAlgorithmType({ id: this.tempTypeId, typeDesc: inputTypeValue }).then(
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
                    if (inputTypeValue) {
                        if (inputTypeValue) {
                            addAlgorithmType(inputTypeValue).then(response => {
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
                    this.inputVisibleType = false;
                    this.inputTypeValue = '';
                }
                else {
                    let inputFrameValue = this.inputFrameValue;
                    inputFrameValue = inputFrameValue.replace(/^\s\s*/, '').replace(/\s\s*$/, '')
                    if (this.isFrameChange) {
                        updateFrameType({ id: this.tempFrameId, frameworkDesc: inputFrameValue }).then(
                            response => {
                                if (response.success) {
                                    this.getFrameType()
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
                    if (inputFrameValue) {
                        if (inputFrameValue) {
                            addFrameType(inputFrameValue).then(response => {
                                if (response.success) {
                                    this.getFrameType()
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
                    this.inputVisibleFrame = false;
                    this.inputFrameValue = '';
                }

            },
            getType() {
                algorithmType({ pageIndex: 1, pageSize: 20 }).then(
                    response => {
                        if (response.success) {
                            this.typeTags = response.data.algorithmTypes
                            if (this.typeTags == null) {
                                this.typeTags = []
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
            getFrameType() {
                frameType({ pageIndex: 1, pageSize: 20 }).then(
                    response => {
                        if (response.success) {
                            this.frameTags = response.data.algorithmFrameworks
                            if (this.frameTags == null) {
                                this.frameTags = []
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
            changeValue(tag, val) {
                if (val === 'type') {
                    this.inputVisibleType = true
                    this.$nextTick(_ => {
                        this.$refs.saveTypeInput.$refs.input.focus();
                    });
                    this.inputTypeValue = tag.typeDesc
                    this.tempType = tag.typeDesc
                    this.tempTypeId = tag.id
                    this.isTypeChange = true
                }
                else {
                    this.inputVisibleFrame = true
                    this.$nextTick(_ => {
                        this.$refs.saveFrameInput.$refs.input.focus();
                    });
                    this.inputFrameValue = tag.frameworkDesc
                    this.tempFrame = tag.frameworkDesc
                    this.tempFrameId = tag.id
                    this.isFrameChange = true
                }

            }
        }
    }
</script>
<style>
    .title{margin-bottom: 20px;font-size: 16px;font-weight: 800;}
    .el-tag+.el-tag {
        margin-left: 25px;
    }

    .button-new-tag {
        margin-left: 25px;
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