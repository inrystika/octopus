<template>
    <div>
        <div class="function">
            <el-button type="primary" @click="add">添加资源池</el-button>
        </div>
        <el-table v-loading="isLoading" :data="tableData" style="width: 100%;font-size: 15px;"
            :header-cell-style="{'text-align':'left','color':'black'}" :cell-style="{'text-align':'left'}">
            <el-table-column label="资源池名称" align="center">
                <template slot-scope="scope">
                    <span>{{ scope.row.name }}</span>
                </template>
            </el-table-column>
            <el-table-column label="默认资源池" align="center">
                <template slot-scope="scope">
                    <span>{{ scope.row.default?'是':'否' }}</span>
                </template>
            </el-table-column>
            <el-table-column label="描述" align="center">
                <template slot-scope="scope">
                    <span>{{ scope.row.desc }}</span>
                </template>
            </el-table-column>
            <el-table-column label="节点列表" align="center" show-overflow-tooltip>
                <template slot-scope="scope">
                    <span>{{ nodeList(scope.row.bindingNodes) }}</span>
                </template>
            </el-table-column>
            <el-table-column label="资源规格" align="center">
                <template slot-scope="scope">
                    <el-button type="text" @click="handleDetail(scope.row)">详情</el-button>
                </template>
            </el-table-column>
            <el-table-column label="操作">
                <template slot-scope="scope">
                    <el-button v-if="scope.row.default===false" type="text" @click="open(scope.row)">删除</el-button>
                    <el-button type="text" @click="handleEdit( scope.row)">编辑</el-button>
                </template>
            </el-table-column>
        </el-table>
        <!-- 节点详情对话框 -->
        <el-dialog title="详情信息" :visible.sync="detailDialog" width="30%" center :close-on-click-modal="false">
            <div class="wrapper">
                <div>NoteBook资源规格</div>
                <div>
                    <el-tag v-for="item in mapResourceSpecIdList.debug" :key="item.index" class="item">{{ item }}
                    </el-tag>
                </div>
            </div>
            <div class="wrapper">
                <div>训练资源规格</div>
                <div>
                    <el-tag v-for="item in mapResourceSpecIdList.train" :key="item.index" class="item">{{ item }}
                    </el-tag>
                </div>
            </div>
            <div class="wrapper">
                <div>部署资源规格</div>
                <div>
                    <el-tag v-for="item in mapResourceSpecIdList.deploy" :key="item.index" class="item">{{ item }}
                    </el-tag>
                </div>
            </div>
            <span slot="footer" class="dialog-footer">
                <el-button @click="detailDialog = false">取 消</el-button>
                <el-button type="primary" @click="detailDialog = false">确 定</el-button>
            </span>
        </el-dialog>
        <!-- 操作对话框 -->
        <el-dialog :title="flag?'增加资源池':'编辑'" :visible.sync="editeDialog" width="40%" :close-on-click-modal="false">
            <el-form ref="ruleForm" :model="ruleForm" :rules="rules" label-width="100px" class="demo-ruleForm">
                <el-form-item label="名称" :label-width="formLabelWidth" prop="name">
                    <el-input v-model="ruleForm.name" autocomplete="off" :disabled="disabled" />
                </el-form-item>
                <el-form-item label="描述" :label-width="formLabelWidth">
                    <el-input v-model="ruleForm.desc" type="textarea" />
                </el-form-item>
                <el-form-item label="节点列表" prop="bindingNodes" :label-width="formLabelWidth">
                    <el-select v-model="ruleForm.bindingNodes" multiple>
                        <el-option v-for="item in nodeOption" :key="item.name" :label="item.name" :value="item.name" />
                    </el-select>
                </el-form-item>
                <el-form-item label="NoteBook资源规格" :label-width="formLabelWidth">
                    <el-select v-model="ruleForm.mapResourceSpecIdList.debug" multiple>
                        <el-option v-for="item in resourceOption" :key="item.id" :label="item.name" :value="item.id" />
                    </el-select>
                </el-form-item>
                <el-form-item label="训练资源规格" :label-width="formLabelWidth">
                    <el-select v-model="ruleForm.mapResourceSpecIdList.train" multiple>
                        <el-option v-for="item in resourceOption" :key="item.id" :label="item.name" :value="item.id" />
                    </el-select>
                </el-form-item>
                <el-form-item label="部署资源规格" :label-width="formLabelWidth">
                    <el-select v-model="ruleForm.mapResourceSpecIdList.deploy" multiple>
                        <el-option v-for="item in resourceOption" :key="item.id" :label="item.name" :value="item.id" />
                    </el-select>
                </el-form-item>
            </el-form>
            <div slot="footer" class="dialog-footer">
                <el-button @click="editeDialog = false">取 消</el-button>
                <el-button type="primary" @click="confirm">确 定</el-button>
            </div>
        </el-dialog>

    </div>
</template>

<script>
    import { getResourcePool, deleteResourcePool, createResourcePool, updateResourcePool, getNodeList, getResource } from '@/api/resourceManager.js'
    import { getErrorMsg } from '@/error/index'
    export default {
        name: "ResourcePool",
        data() {
            var checkName = (rule, value, callback) => {
                const regName = /^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$/;
                if (regName.test(value)) {
                    return callback();
                }
                callback(new Error("请输入合法的名称:首字符和结尾必须为小写英文或者数字其余为英文数字或者-"));
            };
            return {
                nodeDetail: false,
                tableData: [],
                detailDialog: false,
                editeDialog: false,
                formLabelWidth: '220px',
                flag: false,
                mapResourceSpecIdList: { debug: [], train: [], deploy: [] },
                ruleForm: { name: "", desc: "", bindingNodes: [], mapResourceSpecIdList: { debug: [], train: [], deploy: [] } },
                rules: {
                    name: [
                        { required: true, message: '请输入资源池名称', trigger: 'blur' },
                        { validator: checkName, trigger: "blur" }

                    ],
                    bindingNodes: [
                        { required: true, message: '请选择节点', trigger: 'change' }
                    ]

                },
                nodeOption: [],
                resourceOption: [],
                disabled: false,
                // 资源池ID
                id: undefined,
                isLoading: true,
                timer: null

            }
        },
        created() {
            this.getNodeList()
            this.getResource()
        },
        mounted() {
            this.getResourcePool()
        },
        beforeDestroy() {
            clearInterval(this.timer);
        },
        methods: {
            // 错误码
            getErrorMsg(code) {
                return getErrorMsg(code)
            },
            handleDetail(val) {
                this.detailDialog = true
                const mapResourceSpecIdList = JSON.parse(JSON.stringify(val.mapResourceSpecIdList))
                this.mapResourceSpecIdList = { debug: [], train: [], deploy: [] }
                if (mapResourceSpecIdList.debug.resourceSpecIds !== null) {
                    mapResourceSpecIdList.debug.resourceSpecIds.forEach(
                        item => {
                            this.resourceOption.forEach(Item => {
                                if (item === Item.id) {
                                    this.mapResourceSpecIdList.debug.push(
                                        Item.name
                                    )
                                }
                            })
                        }
                    )
                }
                if (mapResourceSpecIdList.train.resourceSpecIds !== null) {
                    mapResourceSpecIdList.train.resourceSpecIds.forEach(
                        item => {
                            this.resourceOption.forEach(Item => {
                                if (item === Item.id) {
                                    this.mapResourceSpecIdList.train.push(
                                        Item.name
                                    )
                                }
                            })
                        }
                    )
                }
                if (mapResourceSpecIdList.deploy.resourceSpecIds !== null) {
                    mapResourceSpecIdList.deploy.resourceSpecIds.forEach(
                        item => {
                            this.resourceOption.forEach(Item => {
                                if (item === Item.id) {
                                    this.mapResourceSpecIdList.deploy.push(
                                        Item.name
                                    )
                                }
                            })
                        }
                    )
                }
            },
            add() {
                this.ruleForm = { name: "", desc: "", bindingNodes: [], mapResourceSpecIdList: { debug: [], train: [], deploy: [] } }
                this.flag = true
                this.editeDialog = true
                this.disabled = false
            },
            handleEdit(val) {
                this.flag = false
                this.editeDialog = true
                this.disabled = true
                this.id = val.id
                this.ruleForm.name = val.name
                this.ruleForm.desc = val.desc
                this.ruleForm.bindingNodes = val.bindingNodes
                this.ruleForm.mapResourceSpecIdList.debug = []
                this.ruleForm.mapResourceSpecIdList.deploy = []
                this.ruleForm.mapResourceSpecIdList.train = []
                if (val.mapResourceSpecIdList.debug.resourceSpecIds === null) { val.mapResourceSpecIdList.debug.resourceSpecIds = [] }
                if (val.mapResourceSpecIdList.deploy.resourceSpecIds === null) { val.mapResourceSpecIdList.deploy.resourceSpecIds = [] }
                if (val.mapResourceSpecIdList.train.resourceSpecIds === null) { val.mapResourceSpecIdList.train.resourceSpecIds = [] }
                val.mapResourceSpecIdList.debug.resourceSpecIds.forEach(item => {
                    this.resourceOption.forEach(Item => {
                        if (item === Item.id) {
                            this.ruleForm.mapResourceSpecIdList.debug.push(Item.id)
                        }
                    })
                })
                val.mapResourceSpecIdList.deploy.resourceSpecIds.forEach(item => {
                    this.resourceOption.forEach(Item => {
                        if (item === Item.id) {
                            this.ruleForm.mapResourceSpecIdList.deploy.push(Item.id)
                        }
                    })
                })
                val.mapResourceSpecIdList.train.resourceSpecIds.forEach(item => {
                    this.resourceOption.forEach(Item => {
                        if (item === Item.id) {
                            this.ruleForm.mapResourceSpecIdList.train.push(Item.id)
                        }
                    })
                })
            },
            confirm() {
                this.$refs['ruleForm'].validate((valid) => {
                    if (valid) {
                        const data = { name: '', desc: "", bindingNodes: [], mapResourceSpecIdList: { debug: { resourceSpecIds: [] }, train: { resourceSpecIds: [] }, deploy: { resourceSpecIds: [] } } }
                        data.name = this.ruleForm.name
                        data.desc = this.ruleForm.desc
                        data.bindingNodes = this.ruleForm.bindingNodes
                        data.mapResourceSpecIdList.debug.resourceSpecIds = this.ruleForm.mapResourceSpecIdList.debug
                        data.mapResourceSpecIdList.deploy.resourceSpecIds = this.ruleForm.mapResourceSpecIdList.deploy
                        data.mapResourceSpecIdList.train.resourceSpecIds = this.ruleForm.mapResourceSpecIdList.train
                        if (this.flag) {
                            createResourcePool(data).then(response => {
                                if (response.success) {
                                    this.$message({
                                        message: '创建成功',
                                        type: 'success'
                                    });
                                    this.getResourcePool()
                                    this.editeDialog = false
                                } else {
                                    this.$message({
                                        message: this.getErrorMsg(response.error.subcode),
                                        type: 'warning'
                                    });
                                }
                            })
                        } else {
                            data.id = this.id
                            delete data.name
                            updateResourcePool(data).then(response => {
                                if (response.success) {
                                    this.$message({
                                        message: '编辑成功',
                                        type: 'success'
                                    });
                                    this.getResourcePool()
                                    this.editeDialog = false
                                } else {
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
            handleDelete(val) {
                deleteResourcePool(val.id).then(response => {
                    if (response.success) {
                        this.$message({
                            message: '删除成功',
                            type: 'success'
                        });
                        this.isLoading = true
                        this.timer = setInterval(this.getResourcePool(), 3000);
                    } else {
                        this.$message({
                            message: this.getErrorMsg(response.error.subcode),
                            type: 'warning'
                        });
                    }
                })
            },
            getResourcePool() {
                getResourcePool().then(response => {
                    if (response.success) {
                        if (response.data !== null && response.data.resourcePools !== null) {
                            this.tableData = response.data.resourcePools
                            this.isLoading = false
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
            getNodeList() {
                getNodeList().then(response => {
                    if (response.success) {
                        if (response.data !== null && response.data.nodes !== null) {
                            this.nodeOption = response.data.nodes
                        } else {
                            this.nodeOption = []
                        }
                    } else {
                        this.$message({
                            message: this.getErrorMsg(response.error.subcode),
                            type: 'warning'
                        });
                    }
                })
            },
            getResource() {
                getResource().then(response => {
                    if (response.success) {
                        if (response.data !== null && response.data.resourceSpecs !== null) {
                            this.resourceOption = response.data.resourceSpecs
                        } else {
                            this.resourceOption = []
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
            nodeList(val) {
                if (val) {
                    return val.toString()
                }
                else {
                    return ''
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

    .wrapper {
        vertical-align: text-top;
        margin: 10px 0 10px 0;

        div {
            margin-top: 5px;
        }
    }

    .item {
        margin: 5px;
    }
</style>