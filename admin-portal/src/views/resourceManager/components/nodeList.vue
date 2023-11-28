<template>
    <div>
        <el-table :data="tableData" style="width: 100%;font-size: 15px" :header-cell-style="{'text-align':'center','color':'black'}"
            :span-method="listSpanMethod" :row-style="{height:'5px'}" :cell-style="{padding:'5px 0'}">
            <el-table-column label="节点名字">
                <template slot-scope="scope">
                    <span>{{ scope.row.name }}</span>
                </template>
            </el-table-column>
            <el-table-column label="IP">
                <template slot-scope="scope">
                    <span>{{ scope.row.ip }}</span>
                </template>
            </el-table-column>
            <el-table-column label="节点状态">
                <template slot-scope="scope">
                    <span>{{ scope.row.status }}</span>
                </template>
            </el-table-column>
            <el-table-column label="所属资源池" show-overflow-tooltip>
                <template slot-scope="scope">
                    <span>{{ scope.row.resourcePools }}</span>
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
                                <!-- <div style="  font-size: 10px;">使用率 </div> -->
                            </div>
                        </div>
                    </template>
                </el-table-column>
            </el-table-column>
        </el-table>

       <div style="float:right">
            <el-pagination
                @size-change="handleSizeChange"
                @current-change="handleCurrentChange"
                :current-page="currentPage"
                :page-sizes="[20, 50, 100]"
                :page-size="pageSize"
                layout="total, sizes, prev, pager, next, jumper"
                :total="totalSize">
            </el-pagination>
       </div>
    </div>
</template>
<script>
    import { getNodeList } from '@/api/resourceManager.js'
    import { formatSize } from '@/utils/index.js'
    export default {
        name: "NodeList",
        components: {

        },
        data() {
            return {
                input: '',
                totalData: [],               
                tableData: [],
                currentPage: 1,
                pageSize: 20,
                totalSize: 0,
                tooldata: []
            }
        },
        created() {
            this.getNodeList()
        },
        methods: {
            tool() {
              
              for(let i = 0;i<1000;i++) {
                let param = {
                  "name": "amax"+i,
                  "ip": "192.168.202.76",
                  "status": "NotReady",
                  "resourcePools": null,
                  "capacity": {
                      "cpu": "80",
                      "memory": "98637700Ki"
                  },
                  "allocated": {
                      "cpu": "550m",
                      "memory": "600Mi"
                  },
                  "children":
                    [
                      {childName:'cpu',id:'1',perentage:0.59,total:"80",use:"550m"},
                      {childName:'cpu',id:'1',perentage:0.59,total:"80",use:"550m"}
                    ]
                  
              }
                this.tooldata.push(param)
              }
                  this.totalSize = this.tooldata.length
                  this.totalData = this.handleTableData(this.tooldata)
            },
            handleSizeChange(val) {
                this.pageSize = val
                this.showCurrentdata()
            },
            handleCurrentChange(val) {
                this.currentPage = val
                this.showCurrentdata()
            },
            showCurrentdata() {
                let start = (this.currentPage-1)*this.pageSize*2
                let end = this.currentPage*this.pageSize*2-1
                this.tableData = this.totalData.slice(start,end)
            },
            getDetail(val) {
                let data = []
                for (const key1 in val.allocated) {
                    for (const key2 in val.capacity) {
                        if (key1 === key2) {
                            let percentage
                            if (parseInt(val.allocated[key1]) === 0) {
                                0
                            } else if ((/^\d+$/.test(val.allocated[key1])) && (/^\d+$/.test(val.capacity[key1]))) {
                                percentage = val.allocated[key1] / val.capacity[key1] * 100
                                percentage = parseFloat(percentage.toFixed(2))
                            } else {
                                percentage = formatSize(val.allocated[key1]) / formatSize(val.capacity[key1])
                                percentage = percentage * 100
                                percentage = parseFloat(percentage.toFixed(2))
                            }
                            data.push({ childName: key1, use: val.allocated[key1], total: val.capacity[key1], percentage: percentage, id: Math.random() })
                        }
                    }
                }
                return data
            },
            getNodeList(val) {
                getNodeList(val).then(response => {
                    if (response.success) {
                        if (response.data !== null && response.data.nodes !== null) {
                            response.data.nodes.forEach(
                                item => {
                                    if (item.resourcePools !== null) {
                                        item.resourcePools = item.resourcePools.toString();
                                    }
                                }
                            )
                            response.data.nodes.forEach(
                                item => {
                                    item.id = Math.random()
                                    if (this.getDetail(item) !== []) {
                                        item.children = this.getDetail(item)
                                    }
                                    else { item.children = [] }
                                }
                            )
                            // this.tool()

                            this.totalSize = response.data.nodes.length
                            this.totalData = this.handleTableData(response.data.nodes)
                            this.showCurrentdata()
                        }
                    } else {
                        this.$message({
                            message: this.getErrorMsg(response.error.subcode),
                            type: 'warning'
                        });
                    }
                })
            },
            //合并列
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
                            ip: data[i].ip,
                            status: data[i].status,
                            resourcePools: data[i].resourcePools
                        }
                        arr.push(info)
                    }
                }
                return arr
            },
            listSpanMethod({ row, column, rowIndex, columnIndex }) {
                if (columnIndex < 4) {
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

    .circleBox {
        position: relative;
        text-align: center;
        top:20px
    }

    .circleCenter {
        position: relative;
        top: -45px;

    }

    .el-progress-circle__track {
        stroke: #409EFF;
    }
</style>