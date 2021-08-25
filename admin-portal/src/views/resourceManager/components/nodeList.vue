<template>
    <div>
        <el-table 
            :data="tableData" 
            style="width: 100%;font-size: 15px;"
            :header-cell-style="{'text-align':'left','color':'black'}" 
            :cell-style="{'text-align':'left'}"
          >
            <el-table-column label="节点名字" align="center">
                <template slot-scope="scope">
                    <span style="margin-left: 10px">{{ scope.row.name }}</span>
                </template>
            </el-table-column>
            <el-table-column label="IP" align="center">
                <template slot-scope="scope">
                    <span style="margin-left: 10px">{{ scope.row.ip }}</span>
                </template>
            </el-table-column>
            <el-table-column label="节点状态" align="center">
                <template slot-scope="scope">
                    <span style="margin-left: 10px">{{ scope.row.status }}</span>
                </template>
            </el-table-column>
            <el-table-column label="所属资源池" align="center" show-overflow-tooltip>
                <template slot-scope="scope">
                    <span style="margin-left: 10px">{{ scope.row.resourcePools }}</span>
                </template>
            </el-table-column>
            <el-table-column label="节点详情" align="center">
                <template slot-scope="scope">
                    <span @mouseover="handleDetail(scope.row)" class="detail">详情</span>
                </template>
            </el-table-column>
        </el-table>
        <!-- 节点详情对话框 -->
        <el-dialog :title="'节点详情/' + title" :visible.sync="nodeDetail" :close-on-click-modal="false">
            <el-table :data="data">
                <el-table-column label="名称">
                    <template slot-scope="scope">
                        <span style="margin-left: 10px">
                            {{ scope.row.name }}
                        </span>
                    </template>
                </el-table-column>
                <el-table-column label="平台使用量">
                    <template slot-scope="scope">
                        <span style="margin-left: 10px">
                            {{ scope.row.use }}
                        </span>
                    </template>
                </el-table-column>
                <el-table-column label="总量">
                    <template slot-scope="scope">
                        <span style="margin-left: 10px">
                            {{ scope.row.total }}
                        </span>
                    </template>
                </el-table-column>
                <el-table-column label="使用百分比">
                    <template slot-scope="scope">
                        <el-progress type="circle" :percentage="scope.row.percentage" :width="50" :height="50">
                        </el-progress>
                    </template>
                </el-table-column>


            </el-table>
            <div slot="footer" class="dialog-footer">
                <el-button @click="nodeDetail = false">取 消</el-button>
                <el-button type="primary" @click="nodeDetail = false">确 定</el-button>
            </div>
        </el-dialog>

    </div>
</template>
<script>
    import { getNodeList } from '@/api/resourceManager.js'
    import { formatSize } from '@/utils/index.js'
    import { getErrorMsg } from '@/error/index'
    export default {
        name: "nodeList",
        components: {


        },
        data() {
            return {
                input: '',
                nodeDetail: false,
                tableData: [],
                data: [],
                title: ""

            }
        },
        created() {
            this.getNodeList()
        },
        methods: {
            // 错误码
            getErrorMsg(code) {
                return getErrorMsg(code)
            },
            handleDetail(val) {
                this.title = val.name
                this.data = []
                for (let key1 in val.allocated) {
                    for (let key2 in val.capacity) {
                        if (key1 === key2) {
                            let percentage
                            if (val.allocated[key1] === 0) { percentage: 0 }
                            else if ((/^\d+$/.test(val.allocated[key1])) && (/^\d+$/.test(val.capacity[key1]))) {
                                percentage = val.allocated[key1] / val.capacity[key1] * 100
                                percentage = parseFloat(percentage.toFixed(2))
                            }
                            else {
                                percentage = formatSize(val.allocated[key1]) / formatSize(val.capacity[key1])
                                percentage = percentage * 100
                                percentage = parseFloat(percentage.toFixed(2))
                            }
                            this.data.push({ name: key1, use: val.allocated[key1], total: val.capacity[key1], percentage: percentage })
                        }
                    }

                }
                this.nodeDetail = true
            },
            getNodeList(val) {
                getNodeList(val).then(response => {
                    if (response.success) {
                        if (response.data != null && response.data.nodes != null) {
                            response.data.nodes.forEach(
                                item => {
                                    if (item.resourcePools !== null) {
                                        item.resourcePools = item.resourcePools.toString();
                                    }

                                }
                            )
                            this.tableData = response.data.nodes
                        }
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


    .detail {
        color: #409eff;
    }
</style>