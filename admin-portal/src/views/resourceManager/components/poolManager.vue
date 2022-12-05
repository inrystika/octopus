<template>
    <div>
        <div class="function">
            <el-button type="primary" @click="add">添加资源池</el-button>
        </div>
        <el-table v-loading="isLoading" :data="tableData" style="width: 100%;font-size: 15px;" :span-method="listSpanMethod"
            :header-cell-style="{'text-align':'center','color':'black'}" :cell-style="{'text-align':'left'}">
            <el-table-column label="资源池名称">
                <template slot-scope="scope">
                    <span>{{ scope.row.name }}</span>
                </template>
            </el-table-column>
            <el-table-column label="默认资源池">
                <template slot-scope="scope">
                    <span>{{ scope.row.default?'是':'否' }}</span>
                </template>
            </el-table-column>
            <el-table-column label="描述">
                <template slot-scope="scope">
                    <span>{{ scope.row.desc }}</span>
                </template>
            </el-table-column>
            <el-table-column label="节点列表" show-overflow-tooltip>
                <template slot-scope="scope">
                    <span>{{  nodeList(scope.row.bindingNodes) }}</span>
                </template>
            </el-table-column>
            <el-table-column label="资源规格">
                <template slot-scope="scope">
                    <el-button type="text" @click="handleDetail(scope.row)">详情</el-button>
                </template>
            </el-table-column>
            <el-table-column label="资源信息">
                <el-table-column label="名称">
                    <template slot-scope="scope">
                        <span style="color: #409eff">
                            {{ scope.row.childName }}
                        </span>
                    </template>
                </el-table-column>
                <el-table-column label="已分配">
                    <template slot-scope="scope">
                        <span style="color: #409eff;">
                            {{ scope.row.use }}
                        </span>
                    </template>
                </el-table-column>
                <el-table-column label="总量">
                    <template slot-scope="scope">
                        <span style="color: #409eff;">
                            {{ scope.row.total }}
                        </span>
                    </template>
                </el-table-column>
                <el-table-column label="已分配百分比" width="120px">
                    <template slot-scope="scope">
                        <div class="circleBox" v-if="!scope.row.children">
                            <el-progress color="#409EFF" type="circle" :show-text="false"
                                :percentage="scope.row.percentage" :width="60" :height="60">
                            </el-progress>
                            <div class="circleCenter">
                                <div style=" font-weight: bold; font-size: 12px;"> {{scope.row.percentage?scope.row.percentage:0}}%</div>
                            </div>
                        </div>
                    </template>
                </el-table-column>
            </el-table-column>
            <el-table-column label="操作">
                <template slot-scope="scope">
                    <el-button v-if="scope.row.default===false" type="text" @click="handleDeletePool(scope.row)">删除</el-button>
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
        <el-dialog :title="flag?'增加资源池':'编辑'" :visible.sync="editDialog" width="40%" :close-on-click-modal="false">
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
                <el-button @click="editDialog = false">取 消</el-button>
                <el-button type="primary" @click="confirm" v-preventReClick>确 定</el-button>
            </div>
        </el-dialog>

    </div>
</template>

<script>
    import { getResourcePool, deleteResourcePool, createResourcePool, updateResourcePool, getNodeList, getResource } from '@/api/resourceManager.js'
    import { formatSize } from '@/utils/index.js'
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
                editDialog: false,
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
          listSpanMethod({ row, column, rowIndex, columnIndex }) {
            if (columnIndex < 5 || columnIndex > 8) {
                    if (row.span_num > 0) {
                        return {
                            rowspan: row.span_num,
                            colspan: 1
                        };

                    }
                    else {
                        return { rowspan: 0, colspan: 0 }
                    }
                }
            },
            handleTableData(data) {
                let arr = [];
                let on = 0;
                let spanNum = 0;
                for (let i = 0; i < data.length; i++) {
                    let node_info = data[i].children
                    on++;
                    for (let j = 0; j < node_info.length; j++) {
                        let info = {
                            on: on,
                            span_num: j === 0 ? node_info.length : 0,
                            childName: node_info[j].childName,
                            use: node_info[j].use,
                            total: node_info[j].total,
                            percentage: node_info[j].percentage,
                            name: data[i].name,
                            default: data[i].default,
                            desc: data[i].desc,
                            bindingNodes: data[i].bindingNodes,
                            mapResourceSpecIdList: data[i].mapResourceSpecIdList,
                            id: data[i].id,
                        }
                        arr.push(info)
                    }
                }
                return arr
            },
            getDetail(val) {
                let data = []
                if (Object.getOwnPropertyNames(val.resourceAllocated).length && Object.getOwnPropertyNames(val.resourceCapacity).length) {
                    for (const key1 in val.resourceAllocated) {
                        for (const key2 in val.resourceCapacity) {
                            if (key1 === key2) {
                                let percentage
                                if (parseInt(val.resourceAllocated[key1]) === 0) {
                                    0
                                } else if ((/^\d+$/.test(val.resourceAllocated[key1])) && (/^\d+$/.test(val.resourceCapacity[key1]))) {
                                  percentage = val.resourceAllocated[key1] / val.resourceCapacity[key1] * 100
                                  percentage = parseFloat(percentage.toFixed(2))
                                } else {                               
                                    percentage = formatSize(val.resourceAllocated[key1]) / formatSize(val.resourceCapacity[key1])
                                    percentage = percentage * 100
                                    percentage = parseFloat(percentage.toFixed(2))
                                }
                                data.push({ childName: key1, use: val.resourceAllocated[key1], total: val.resourceCapacity[key1], percentage: percentage})
                            }
                        }
                    }
                } else {
                  data.push(
                    {
                        childName: "cpu",
                        id: "",
                        percentage: 0,
                        total: "0",
                        use: "0"
                    },
                    {
                        childName: "memory",
                        id: "",
                        percentage: 0,
                        total: "0",
                        use: "0"
                    },
                  )
                }
                return data
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
                this.editDialog = true
                this.disabled = false
            },
            handleEdit(val) {
                this.flag = false
                this.editDialog = true
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
                                    this.editDialog = false
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
                                    this.editDialog = false
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
                            this.isLoading = false
                            response.data.resourcePools.forEach(
                                item => {
                                    // item.id = Math.random()
                                    if (this.getDetail(item) !== []) {
                                        item.children = this.getDetail(item)
                                    }
                                    else { item.children = [] }
                                }
                            )
                            this.tableData = this.handleTableData(response.data.resourcePools)
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
            handleDeletePool(val) {
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
            nodeList(val){
              if(val){
                  return val.toString()
              }
              else{
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
    .circleBox {
        position: relative;
        text-align: center;
        top:20px
    }

    .circleCenter {
        position: relative;
        top: -45px;

    }
</style>