<template>

    <div>
        <div class="function">
            <el-button type="primary" @click="addDialog = true">添加规格</el-button>
        </div>
        <el-table
            :data="tableData"
            style="width: 100%;font-size: 15px;"
            :header-cell-style="{'text-align':'left','color':'black'}"
            :cell-style="{'text-align':'left'}"
        >
            <el-table-column label="规格名称" align="center">
                <template slot-scope="scope">
                    <span>{{ scope.row.name }}</span>
                </template>
            </el-table-column>
            <el-table-column label="机时价格(1~10)" align="center">
                <template slot-scope="scope">
                    <span>{{ scope.row.price }}</span>
                </template>
            </el-table-column>
            <el-table-column label="资源数量" align="center">
                <template slot-scope="scope">
                    <span>{{ scope.row.resourceQuantity }}</span>
                </template>
            </el-table-column>
            <el-table-column label="操作" align="center">
                <template slot-scope="scope">
                    <el-button type="text" @click="open(scope.row)">删除</el-button>
                    <el-button type="text" @click="edit(scope.row)">编辑</el-button>
                </template>
            </el-table-column>
        </el-table>
        <!-- 操作对话框 -->
        <el-dialog title="添加规格" :visible.sync="addDialog" @close="closeDialog">
            <el-form ref="ruleForm" :model="ruleForm" :rules="rules" label-width="100px" class="demo-ruleForm">
                <el-form-item label="规格名称" prop="name">
                    <el-input v-model="ruleForm.name" />
                </el-form-item>
                <el-form-item label="机时价格" prop="price">
                    <el-input-number v-model="ruleForm.price" :min="0" label="描述文字" :precision="2" :step="0.01"/>
                    <span class="red">支持两位小数</span>
                </el-form-item>
                <el-form-item label="资源信息" prop="resourceQuantity">
                    <div v-for="(item, index) in ruleForm.resourceQuantity" :key="index">
                        <el-form-item style="margin-bottom:10px">
                            <el-select v-model="item.key" style="width: 20%;">
                                <el-option
                                    v-for="item in options"
                                    :key="item.name"
                                    :label="item.name"
                                    :value="item.name"
                                />
                            </el-select>
                            <span style="margin:0 10px 0 10px">=</span>
                            <el-input v-model="item.value" style="width: 20%;" />
                            <i class="el-icon-delete" @click="deleteItem(item, index, 'add')"></i>
                        </el-form-item>
                    </div>
                    <el-button type="primary" @click="addItem('add')">增加</el-button>
                </el-form-item>
            </el-form>
            <div slot="footer" class="dialog-footer">
                <el-button type="primary" @click="confirm" v-preventReClick>确认</el-button>
                <el-button @click="cancel">取消</el-button>
            </div>
        </el-dialog>

        <!-- 编辑对话框 -->
        <el-dialog title="编辑规格" :visible.sync="editDialog" @close="closeDialog">
            <el-form ref="eidtRuleForm" :model="eidtRuleForm" :rules="rules" label-width="100px" class="demo-ruleForm">
                <el-form-item label="规格名称" prop="name">
                    <el-input v-model="eidtRuleForm.name" />
                </el-form-item>
                <el-form-item label="机时价格" prop="price">
                    <el-input-number v-model="eidtRuleForm.price" :min="0" label="描述文字" :precision="2" :step="0.01"/>
                    <span class="red">支持两位小数</span>
                </el-form-item>
                <el-form-item label="资源信息" prop="resourceQuantity">
                    <div v-for="(item, index) in eidtRuleForm.resourceQuantity" :key="index">
                        <el-form-item style="margin-bottom:10px">
                            <el-select v-model="item.name" style="width: 20%;">
                                <el-option
                                    v-for="item in options"
                                    :key="item.name"
                                    :label="item.name"
                                    :value="item.name"
                                />
                            </el-select>
                            <span style="margin:0 10px 0 10px">=</span>
                            <el-input v-model="item.value" style="width: 20%;" />
                            <i class="el-icon-delete" @click="deleteItem(item, index, 'edit')"></i>
                        </el-form-item>
                    </div>
                    <el-button type="primary" @click="addItem('edit')">增加</el-button>
                </el-form-item>
            </el-form>
            <div slot="footer" class="dialog-footer">
                <el-button type="primary" @click="update" v-preventReClick>确认</el-button>
                <el-button @click="cancel">取消</el-button>
            </div>
        </el-dialog>

    </div>
</template>

<script>
    import { getResource, deleteSpecification, createResource, getResourceList, updateSpecification } from '@/api/resourceManager.js'
    export default {
        name: "Resource",
        components: {},
        data() {
            return {
                addDialog: false,
                editDialog: false,
                tableData: [],
                msg: '',
                eidtRuleForm: {
                    id: '',
                    name: '',
                    price: undefined,
                    resourceQuantity: []
                },
                ruleForm: {
                    name: '',
                    price: undefined,
                    resourceQuantity: []
                },
                rules: {
                    name: [
                        { required: true, message: '请输入规格名称', trigger: 'blur' }

                    ],
                    price: [
                        { required: true, message: '请输入机时价格', trigger: 'blur' }
                    ],
                    resourceQuantity: { required: true, message: '请输入资源信息', trigger: ['change', 'blur'] }
                },
                options: []

            }
        },
        created() {
            this.getResource()
            this.getResourceList()
        },
        methods: {
            handleDelete(val) {
                deleteSpecification(val.id).then(response => {
                    if (response.success) {
                        this.$message({
                            message: '删除成功',
                            type: 'success'
                        });
                        this.getResource()
                    } else {
                        this.$message({
                            message: this.getErrorMsg(response.error.subcode),
                            type: 'warning'
                        });
                    }
                })
            },
            addItem(name) {
                if(name == 'add') {
                    this.ruleForm.resourceQuantity.push({
                        key: '',
                        value: ''
                    })
                    return
                }
                this.eidtRuleForm.resourceQuantity.push({
                    key: '',
                    value: ''
                })
            },
            deleteItem(item, index, name) {
                if(name == 'add') {
                    this.ruleForm.resourceQuantity.splice(index, 1)
                    return 
                }
                this.eidtRuleForm.resourceQuantity.splice(index, 1)
            },
            confirm() {
                this.$refs['ruleForm'].validate((valid) => {
                    if (valid) {
                        const obj = {}
                        const data = JSON.parse(JSON.stringify(this.ruleForm))
                        data.resourceQuantity.forEach(
                            item => {
                                obj[item.key] = item.value
                            }
                        )
                        data.resourceQuantity = obj
                        let flag = true
                        for (var key in data.resourceQuantity) {
                            if (key === '' || data.resourceQuantity[key] === '') {
                                flag = false
                            }
                        }
                        if (!flag) {
                            this.$message({
                                message: '资源信息不能为空',
                                type: 'warning'
                            });
                            return
                        } else {
                            createResource(data).then(response => {
                                if (response.success) {
                                    this.$message({
                                        message: '添加规格成功',
                                        type: 'success'
                                    })
                                    this.ruleForm = {
                                        name: '',
                                        price: undefined,
                                        resourceQuantity: []
                                    }
                                    this.getResource()
                                    this.addDialog = false
                                } else {
                                    this.$message({
                                        message: this.getErrorMsg(response.error.subcode),
                                        type: 'warning'
                                    });
                                }
                            });
                        }
                    } else {
                        console.log('error submit!!');
                        return false;
                    }
                });
            },
            cancel() { 
                this.addDialog = false
                this.editDialog = false
                this.eidtRuleForm = {
                    id: '',
                    name: '',
                    price: undefined,
                    resourceQuantity: []
                }
                this.ruleForm = {
                    name: '',
                    price: undefined,
                    resourceQuantity: []
                }
            },
            getResource() {
                getResource().then(response => {
                    if (response.success) {
                        if (response.data !== null && response.data.resourceSpecs !== null) {
                            response.data.resourceSpecs.forEach(
                                item => {
                                    item.resourceQuantity = JSON.stringify(item.resourceQuantity).replace(/\"/g, '').replace(/\{|}/g, '')
                                }
                            )
                            this.tableData = response.data.resourceSpecs
                        } else {
                          this.tableData = []
                        }
                    } else {
                        this.$message({
                            message: this.getErrorMsg(response.error.subcode),
                            type: 'warning'
                        });
                    }
                })
            },
            handleSizeChange(val) {
                this.pageSize = val
                this.getResource()
            },
            handleCurrentChange(val) {
                this.pageIndex = val
                this.getResource()
            },
            getResourceList() {
                getResourceList().then(response => {
                    if (response.success) {
                        if (response.data !== null && response.data.resources !== null) {
                            this.options = response.data.resources
                        }
                    } else {
                        this.$message({
                            message: this.getErrorMsg(response.error.subcode),
                            type: 'warning'
                        });
                    }
                })
            },
            // 删除确认
            open(val) {
                this.$confirm('此操作将永久删除该资源规格, 是否继续?', '提示', {
                    confirmButtonText: '确定',
                    cancelButtonText: '取消',
                    type: 'warning'
                }).then(() => {
                    this.handleDelete(val)
                }).catch(() => {
                    this.$message({
                        type: 'info',
                        message: '已取消删除'
                    });
                });
            },
            update() {
                this.$refs['eidtRuleForm'].validate((valid) => {                 
                    if (valid) {
                        let obj = {}
                        let param = {}
                        obj.name = this.eidtRuleForm.name
                        obj.id = this.eidtRuleForm.id
                        obj.price = this.eidtRuleForm.price
                        this.eidtRuleForm.resourceQuantity.forEach(
                            item => {
                                param[item.name] = item.value
                            }
                        )
                        obj.resourceQuantity = param
                        updateSpecification(obj).then(response => {
                            if (response.success) {
                                this.$message({
                                    message: '更新规格成功',
                                    type: 'success'
                                })
                                this.eidtRuleForm = {
                                    id: '',
                                    name: '',
                                    price: undefined,
                                    resourceQuantity: []
                                }
                                this.getResource()
                                this.editDialog = false
                            } else {
                                this.$message({
                                    message: this.getErrorMsg(response.error.subcode),
                                    type: 'warning'
                                });
                            }
                        })                            
                    } else {
                      console.log("update error")
                      return false;
                    }
                })
            },
            edit(val) {
                let tempArr = val.resourceQuantity.split(",")
                tempArr.forEach((item) => {
                    let arr = item.split(":")
                    let obj = {}
                    obj.name = arr[0]
                    obj.value = arr[1]
                    this.eidtRuleForm.resourceQuantity.push(obj)
                })
                this.editDialog = true
                this.eidtRuleForm.id = val.id
                this.eidtRuleForm.name = val.name
                this.eidtRuleForm.price = val.price
            },
            closeDialog() {
                this.addDialog = false
                this.editDialog = false
                this.eidtRuleForm = {
                    name: '',
                    price: undefined,
                    resourceQuantity: []
                }
                this.ruleForm = {
                    name: '',
                    price: undefined,
                    resourceQuantity: []
                }

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

    .line {
        text-align: center;
    }
    .red{color:#409EFF;margin-left: 10px;font-weight: 800;}
</style>