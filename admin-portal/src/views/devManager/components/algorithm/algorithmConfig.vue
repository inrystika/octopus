<template>
    <div>
        <div class="title">模型类别</div>
        <div>
            <el-tag v-for="(tag,index) in dynamicType" :key="index" :closable="tag.sourceType!==1"
                :disable-transitions="false" @click="editTag(tag,index,'TYPE')" @close="handleClose(tag,'TYPE')">
                <span v-if="index!=typeNum"> {{tag.lableDesc }}</span>
                <input class="custom_input" type="text" v-model="typeValue" v-if="index==typeNum" ref="editeTypeInput"
                    @keyup.enter.native="handleInput(tag,'TYPE')" @blur="handleInput(tag,'TYPE')">
            </el-tag>
            <el-input class="input-new-tag" v-if="inputTypeVisible" v-model="typeValue" ref="saveTypeInput" size="small"
                @keyup.enter.native="handleInputConfirm('TYPE')" @blur="handleInputConfirm('TYPE')">
            </el-input>
            <el-button v-else class="button-new-tag" size="small" @click="showInput('TYPE')">{{'+ 新标签'}}</el-button>
        </div>
        <el-divider></el-divider>
        <div class="title">框架类型</div>
        <div>
            <el-tag v-for="(tag,index) in dynamicFrame" :key="index" :closable="tag.sourceType!==1"
                :disable-transitions="false" @click="editTag(tag,index,'FRAME')" @close="handleClose(tag,'FRAME')">
                <span v-if="index!=frameNum"> {{tag.lableDesc }}</span>
                <input class="custom_input" type="text" v-model="frameValue" v-if="index==frameNum"
                    ref="editeFrameInput" @keyup.enter.native="handleInput(tag,'FRAME')"
                    @blur="handleInput(tag,'FRAME')">
            </el-tag>
            <el-input class="input-new-tag" v-if="inputFrameVisible" v-model="frameValue" ref="saveFrameInput"
                size="small" @keyup.enter.native="handleInputConfirm" @blur="handleInputConfirm">
            </el-input>
            <el-button v-else class="button-new-tag" size="small" @click="showInput">{{'+ 新标签'}}</el-button>
        </div>
    </div>
</template>
<script>
    import { algorithmType, addAlgorithmType, deleteAlgorithmType, updateAlgorithmType, frameType, addFrameType, deleteFrameType, updateFrameType } from "@/api/modelDev"
    import { getErrorMsg } from '@/error/index'
    export default {
        name: 'star-input-tag',
        created() {
            this.algorithmType();
            this.frameType()
        },
        data() {
            return {
                //算法类型
                inputTypeVisible: false,
                typeValue: '',
                typeId: '',
                typeNum: -1,
                dynamicType: [],
                //算法框架
                inputFrameVisible: false,
                frameValue: '',
                frameId: '',
                frameNum: -1,
                dynamicFrame: []
            }
        },

        methods: {
            handleClose(tag, val) {
                if (val === 'TYPE') {
                    deleteAlgorithmType(tag.id).then(response => {
                        if (response.success) {
                            this.algorithmType()
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
                            this.frameType()
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
                if (val === 'TYPE') {
                    this.inputTypeVisible = true;
                    this.$nextTick(_ => {
                        this.$refs.saveTypeInput.$refs.input.focus();
                    });
                }
                else {
                    this.inputFrameVisible = true;
                    this.$nextTick(_ => {
                        this.$refs.saveFrameInput.$refs.input.focus();
                    });
                }


            },

            handleInputConfirm(val) {
                if (val === 'TYPE') {
                    let typeValue = this.typeValue;
                    typeValue = typeValue.replace(/^\s\s*/, '').replace(/\s\s*$/, '')
                    // 点击添加时，追加
                    if (typeValue) {
                        if (typeValue) {
                            addAlgorithmType(typeValue).then(response => {
                                if (response.success) {
                                    this.algorithmType()
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
                    this.inputTypeVisible = false;
                    this.typeValue = '';
                }
                else {
                    let frameValue = this.frameValue;
                    frameValue = frameValue.replace(/^\s\s*/, '').replace(/\s\s*$/, '')
                    // 点击添加时，追加
                    if (frameValue) {
                        if (frameValue) {
                            addFrameType(frameValue).then(response => {
                                if (response.success) {
                                    this.frameType()
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
                    this.inputFrameVisible = false;
                    this.frameValue = '';
                }
            },
            editTag(tag, index, val) {
                if (tag.sourceType == 1) {
                    return
                } else {
                    if (val === 'TYPE') {
                        this.typeNum = index;
                        this.$nextTick(_ => {
                            this.$refs.editeTypeInput[0].focus();
                        });
                        this.typeValue = tag.lableDesc;
                        this.typeId = tag.id
                    }
                    else {
                        this.frameNum = index;
                        this.$nextTick(_ => {
                            this.$refs.editeFrameInput[0].focus();
                        });
                        this.frameValue = tag.lableDesc;
                        this.frameId = tag.id
                    }
                }

            },
            handleInput(tag, val) {
                if (val === 'TYPE') {
                    updateAlgorithmType({ id: this.typeId, lableDesc: this.typeValue }).then(
                        response => {
                            if (response.success) {
                                this.algorithmType()
                            }
                            else {
                                this.$message({
                                    message: this.getErrorMsg(response.error.subcode),
                                    type: 'warning'
                                });
                            }
                            this.typeValue = '';
                            this.typeNum = -1;
                        }
                    )
                }
                else {
                    updateFrameType({ id: this.frameId, lableDesc: this.frameValue }).then(
                        response => {
                            if (response.success) {
                                this.frameType()
                            }
                            else {
                                this.$message({
                                    message: this.getErrorMsg(response.error.subcode),
                                    type: 'warning'
                                });
                            }
                            this.frameValue = '';
                            this.frameNum = -1;
                        }
                    )
                }

            },
            getErrorMsg(code) {
                return getErrorMsg(code)
            },
            algorithmType() {
                algorithmType({ pageIndex: 1, pageSize: 20 }).then(
                    response => {
                        if (response.success) {
                            this.dynamicType = response.data.lables
                            if (this.dynamicType == null) {
                                this.dynamicType = []
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
            frameType() {
                frameType({ pageIndex: 1, pageSize: 20 }).then(
                    response => {
                        if (response.success) {
                            this.dynamicFrame = response.data.lables
                            if (this.dynamicFrame == null) {
                                this.dynamicFrame = []
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

    .title {
        margin-bottom: 20px;
        font-size: 16px;
        font-weight: 800;
    }
</style>