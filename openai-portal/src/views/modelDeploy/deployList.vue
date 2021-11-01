<template>
    <div>
        <div>
            <searchForm :search-form="searchForm" :blur-name="'搜索'" @searchData="getSearchData" />
        </div>
        <div class="index">
            <el-table :data="tableData" style="width: 100%;font-size: 15px;"
                :header-cell-style="{'text-align':'left','color':'black'}" :cell-style="{'text-align':'left'}">
                <el-table-column prop="name" label="模型名称" align="center" />
                <el-table-column prop="version" label="版本" align="center" />
                <el-table-column prop="descript" label="模型描述" align="center" />
                <el-table-column prop="status" label="状态" align="center" />
                <el-table-column label="创建时间" align="center">
                    <template slot-scope="scope">
                        <span style="margin-left: 10px">{{ parseTime(scope.row.createdAt) }}</span>
                    </template>
                </el-table-column>
                <el-table-column label="操作" align="center">
                    <template slot-scope="scope">
                        <el-button type="text" @click="handleDetail(scope.row)">详情</el-button>
                    </template>
                </el-table-column>
            </el-table>
        </div>
        <div class="block">
            <el-pagination :current-page="searchData.pageIndex" :page-sizes="[10, 20, 50,80]"
                :page-size="searchData.pageSize" layout="total, sizes, prev, pager, next, jumper" :total="total"
                @size-change="handleSizeChange" @current-change="handleCurrentChange" />
        </div>
        <!-- 部署详情对话框 -->
        <detailDialog v-if="detailVisible" @close="close" @cancel="cancel" @confirm="confirm" />
    </div>
</template>

<script>
    import detailDialog from './components/index.vue'
    import { parseTime } from '@/utils/index'
    import searchForm from '@/components/search/index.vue'
    import { getErrorMsg } from '@/error/index'
    export default {
        name: "MyModel",
        components: {
            detailDialog,
            searchForm
        },

        data() {
            return {
                searchData: {
                    pageIndex: 1,
                    pageSize: 10
                },
                total:undefined,
                tableData: [{ name: '预制模型一', version: '1.1', descript: '这是一段模型描述', status: '运行中', createdAt: '111110000' }],
                detailVisible: false,
                searchForm:[]
            }
        },
        created() {
            console.log(this.$route.params.data)
        },
        methods: {
            // 错误码
            getErrorMsg(code) {
                return getErrorMsg(code)
            },
            handleSizeChange(val) {
                this.searchData.pageSize = val
            },
            handleCurrentChange(val) {
                this.searchData.pageIndex = val
            },
            close(val) {
                this.detailVisible = val;
                this.detailVisible = val
            },
            cancel(val) {
                this.detailVisible = val;
                this.detailVisible = val
            },
            confirm(val) {
                this.detailVisible = val;
                this.detailVisible = val
            },

            handleDetail(val) {
                this.detailVisible = true
            },
            getSearchData(){},
            // 时间戳转换日期
            parseTime(val) {
                return parseTime(val)
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