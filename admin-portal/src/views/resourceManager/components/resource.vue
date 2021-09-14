<template>
    <div>
        <div class="function">
            <el-button v-if="customize" type="primary" @click="add">增加资源</el-button>
        </div>
        <div>
            <el-table 
                :data="tableData" 
                style="width: 100%;font-size: 15px;"
                :header-cell-style="{'text-align':'left','color':'black'}" 
                :cell-style="{'text-align':'left'}"
            >
                <el-table-column label="资源名称" align="center">
                    <template slot-scope="scope">
<<<<<<< HEAD
                        <span>{{ scope.row.name }}</span>
=======
                        <span >{{ scope.row.name }}</span>
>>>>>>> 1cc371d52e40272c728853b54fb6e132110a588a
                    </template>
                </el-table-column>
                <el-table-column v-if="customize" label="引用" align="center">
                    <template slot-scope="scope">
<<<<<<< HEAD
                        <span>{{ scope.row.resourceRef }}</span>
=======
                        <span >{{ scope.row.resourceRef }}</span>
>>>>>>> 1cc371d52e40272c728853b54fb6e132110a588a
                    </template>
                </el-table-column>
                <el-table-column v-if="customize" label="绑定节点" align="center" show-overflow-tooltip>
                    <template slot-scope="scope">
<<<<<<< HEAD
                        <span>{{ scope.row.bindingNodes }}</span>
=======
                        <span >{{ scope.row.bindingNodes }}</span>
>>>>>>> 1cc371d52e40272c728853b54fb6e132110a588a
                    </template>
                </el-table-column>
                <el-table-column label="备注" align="center">
                    <template slot-scope="scope">
<<<<<<< HEAD
                        <span>{{ scope.row.desc }}</span>
=======
                        <span >{{ scope.row.desc }}</span>
>>>>>>> 1cc371d52e40272c728853b54fb6e132110a588a
                    </template>
                </el-table-column>
                <el-table-column label="操作" align="center">
                    <template slot-scope="scope">
                        <el-button type="text" @click="handleEdite(scope.row)">编辑</el-button>
                        <el-button v-if="customize" type="text" @click="handleDelete(scope.row)">删除</el-button>
                    </template>
                </el-table-column>
            </el-table>
        </div>
        <!-- 增加和编辑自定义资源对话框 -->
        <el-dialog :title="flag?'增加资源':'编辑'" :visible.sync="dialogVisible" width="40%" :close-on-click-modal="false">
            <el-form :model="ruleForm" :rules="rules" ref="ruleForm" label-width="100px" class="demo-ruleForm">
                <el-form-item label="资源名称" :label-width="formLabelWidth" prop="name">
                    <el-input v-model="ruleForm.name" autocomplete="off" :disabled="disabled">
                    </el-input>
                </el-form-item>
                <el-form-item v-if="customize" label="引用" prop="resourceRef" :label-width="formLabelWidth">
                    <el-select v-model="ruleForm.resourceRef" placeholder="请选择引用">
                        <el-option v-for="item in resourceOption" :key="item.name" :label="item.name"
                            :value="item.name">
                        </el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="备注" :label-width="formLabelWidth">
                    <el-input v-model="ruleForm.desc" autocomplete="off"></el-input>
                </el-form-item>
                <el-form-item v-if="customize" label="绑定节点" prop="bindingNodes" :label-width="formLabelWidth">
                    <el-select v-model="ruleForm.bindingNodes" placeholder="请选择绑定节点" multiple>
                        <el-option v-for="item in nodeOption" :key="item.name" :label="item.name" :value="item.name">
                        </el-option>
                    </el-select>
                </el-form-item>
            </el-form>
            <div slot="footer" class="dialog-footer">
                <el-button @click="dialogVisible = false">取 消</el-button>
                <el-button type="primary" @click="confirm">确 定</el-button>
            </div>
        </el-dialog>
    </div>
</template>
<script>
    import { getResourceList, deleteResource, updateResource, getNodeList, createCustomizeResource } from '@/api/resourceManager.js'
    import { getErrorMsg } from '@/error/index'
    export default {
        name: "resource",
        props: {
            Type: { type: Number, default: undefined }
        },
        created() {
            if (this.Type === 1) { this.system = true; this.customize = false, this.disabled = true; this.rules = {} }
            else {
                this.customize = true
                this.system = false
            }
            this.getResourceList()
            this.getNodeList()
        },
        data() {
            var checkName = (rule, value, callback) => {
                const regName = /^(([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9])$/;
                if (regName.test(value)) {
                    return callback();
                }
                callback(new Error("请输入合法的名称:首字符和结尾必须为数字或者英文字母，其余字符支持数字英文-,_或者."));
            };
            return {
                flag: true,
                disabled: false,
                dialogVisible: false,
                systemVisible: false,
                tableData: [],
                ruleForm: { name: '', resourceRef: '', desc: '', bindingNodes: [], id: undefined },
                rules: {
                    name: [
                        { required: true, message: '请输入资源名称', trigger: 'blur' },
                        { validator: checkName, trigger: "blur" }


                    ],
                    resourceRef: [
                        { required: true, message: '请选择引用', trigger: 'change' }
                    ],
                    bindingNodes: [
                        { required: true, message: '请选择绑定节点', trigger: 'change' }
                    ]

                },
                formLabelWidth: '120px',
                system: false,
                customize: false,
                nodeOption: [],
                resourceOption: [],



            }
        },
        methods: {
            // 错误码
            getErrorMsg(code) {
                return getErrorMsg(code)
            },
            handleEdite(val) {
                let temp = val.bindingNodes.split(',')
                this.ruleForm = JSON.parse(JSON.stringify(val))
                this.ruleForm.bindingNodes = temp
                this.flag = false
                this.dialogVisible = true



            },
            confirm() {
                this.$refs['ruleForm'].validate((valid) => {
                    if (valid) {
                        if (!this.flag) {
                            updateResource(this.ruleForm).then(response => {
                                if (response.success) {
                                    this.$message({
                                        message: '编辑成功',
                                        type: 'success'
                                    });
                                    this.getResourceList();
                                    this.dialogVisible = false
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
                            createCustomizeResource(this.ruleForm).then(response => {
                                if (response.success) {
                                    this.$message({
                                        message: '创建成功',
                                        type: 'success'
                                    });
                                    this.getResourceList()
                                    this.dialogVisible = false
                                }
                                else {
                                    this.$message({
                                        message: this.getErrorMsg(response.error.subcode),
                                        type: 'warning'
                                    });
                                }
                            })
                        }


                    } else {
                        console.log('error submit!!');
                        return false;
                    }
                });

            },
            handleDelete(row) {
                deleteResource(row.id).then(response => {
                    if (response.success) {
                        this.$message({
                            message: '删除成功',
                            type: 'success'
                        });
                        this.getResourceList()
                    }
                    else {
                        this.$message({
                            message: this.getErrorMsg(response.error.subcode),
                            type: 'warning'
                        });
                    }
                })
            },
            add() {
                this.ruleForm = {},
                    this.flag = true
                this.dialogVisible = true
            },
            getResourceList() {
                getResourceList().then(response => {
                    if (response.success) {
                        if (this.system && response.data !== null && response.data.resources !== null) {
                            this.tableData = response.data.resources.filter(item => item.resourceRef.length === 0)
                        }
                        if (this.customize && response.data !== null && response.data.resources !== null) {
                            this.tableData = response.data.resources.filter(item => item.resourceRef.length !== 0)
                            this.resourceOption = response.data.resources.filter(item => item.resourceRef.length === 0)

                        }
                        this.tableData.forEach(item => {
                            if (item.bindingNodes !== null) { item.bindingNodes = item.bindingNodes.toString() }

                        })
                    }
                    else {
                        this.$message({
                            message: this.getErrorMsg(response.error.subcode),
                            type: 'warning'
                        });
                    }

                })
            },
            getNodeList() {
                getNodeList().then(response => {
                    if (response.success) {
                        if (response.data !== null && response.data.nodes !== null) {
                            this.nodeOption = response.data.nodes
                        }
                    }
                    else {
                        this.$message({
                            message: this.getErrorMsg(response.error.subcode),
                            type: 'warning'
                        });
                    }


                })
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