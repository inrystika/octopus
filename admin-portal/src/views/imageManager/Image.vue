<template>
    <div>
        <searchForm :search-form="searchForm" class="searchForm" :blur-name="'镜像名称/标签/描述 搜索'" @searchData="getSearchData" />
        <el-button v-if="!flag" type="primary" class="create" @click="create">创建</el-button>
        <el-table
            :data="tableData"
            style="width: 100%;font-size: 15px;"
            :header-cell-style="{'text-align':'left','color':'black'}"
            :cell-style="{'text-align':'left'}"
        >
            <el-table-column label="镜像名称" align="center">
                <template slot-scope="scope">
                    <span>{{ scope.row.imageName }}</span>
                </template>
            </el-table-column>
            <el-table-column label="镜像标签" align="center">
                <template slot-scope="scope">
                    <span>{{ scope.row.imageVersion }}</span>
                </template>
            </el-table-column>
            <el-table-column label="镜像描述" align="center" :show-overflow-tooltip="true">
                <template slot-scope="scope">
                    <span>{{ scope.row.imageDesc }}</span>
                </template>
            </el-table-column>
            <el-table-column v-if="flag" label="群组名">
                <template slot-scope="scope">
                    <span>{{ scope.row.spaceName===''?'默认群组':scope.row.spaceName }}</span>
                </template>
            </el-table-column>
            <el-table-column v-if="flag" label="提供者" align="center">
                <template slot-scope="scope">
                    <span>{{ scope.row.username }}</span>
                </template>
            </el-table-column>
            <el-table-column label="镜像地址" align="center" :show-overflow-tooltip="true">
                <template slot-scope="scope">
                    <span>{{ scope.row.sourceType===2?scope.row.imageAddr:'' }}</span>
                </template>
            </el-table-column>
            <el-table-column label="镜像类型" align="center">
                <template slot-scope="scope">
                    <span>{{ scope.row.imageType===1?'NoteBook类型':'训练类型' }}</span>
                </template>
            </el-table-column>
            <el-table-column label="上传类型" align="center">
                <template slot-scope="scope">
                    <span>{{ scope.row.sourceType===1?'上传':'远程' }}</span>
                </template>
            </el-table-column>
            <el-table-column label="状态" align="center">
                <template slot-scope="scope">
                    <span>{{ imageStatus(scope.row.imageStatus) }}</span>
                </template>
            </el-table-column>
            <el-table-column label="创建时间" align="center">
                <template slot-scope="scope">
                    <span>{{ parseTime(scope.row.createdAt) }}</span>
                </template>
            </el-table-column>
            <el-table-column v-if="!flag" label="操作" align="center" width="250">
                <template slot-scope="scope">
                    <el-button v-if="scope.row.imageStatus!==3" type="text" @click="handleEdit(scope.row)">重新上传
                    </el-button>
                    <el-button type="text" @click="open2(scope.row)">删除</el-button>
                    <el-button type="text" @click="open(scope.row)">修改描述</el-button>
                </template>
            </el-table-column>
        </el-table>
        <div class="block">
            <el-pagination
                :current-page="searchData.pageIndex"
                :page-sizes="[10, 20, 50, 80]"
                :page-size="searchData.pageSize"
                :total="total"
                layout="total, sizes, prev, pager, next, jumper"
                @size-change="handleSizeChange"
                @current-change="handleCurrentChange"
            />
        </div>
        <!-- 镜像对话框 -->
        <dialogForm
            v-if="FormVisible"
            :flag="Logo"
            :row="row"
            @cancel="cancel"
            @confirm="confirm"
            @close="close"
        />
    </div>
</template>
<script>
    import dialogForm from "./components/dialogForm.vue";
    import { getUserImage, getPreImage, deletePreImage, editePreImage } from '@/api/imageManager.js'
    // import { groupDetail } from '@/api/userManager.js'
    import searchForm from '@/components/search/index.vue'
    import { parseTime } from '@/utils/index'
    import { getErrorMsg } from '@/error/index'
    export default {
        name: "PreImage",
        components: {
            dialogForm,
            searchForm
        },
        props: {
            imageTabType: { type: Number, default: undefined }
        },
        data() {
            return {
                tableData: [],
                row: {
                },
                total: undefined,
                FormVisible: false,
                flag: true,
                Logo: true,
                searchForm: [
                    {
                        type: 'Select', label: '状态', prop: 'imageStatus', placeholder: '请选择状态',
                        options: [{ label: '未制作', value: 1 }, { label: '制作中', value: 2 }, { label: '制作完成', value: 3 }, { label: '制作失败', value: 4 }]
                    },
                    { type: 'Input', label: '镜像名', prop: 'imageNameLike', placeholder: '请输入镜像名' },
                    {
                        type: 'Select', label: '镜像类型', prop: 'imageType', placeholder: '请选择镜像类型',
                        options: [{ label: 'NoteBook类型', value: 1 }, { label: '训练类型', value: 2 }]
                    },
                    {
                        type: 'Select', label: '来源类型', prop: 'sourceType', placeholder: '请选择来源类型',
                        options: [{ label: '上传', value: 1 }, { label: '远程', value: 2 }]
                    }],
                searchData: {
                    pageIndex: 1,
                    pageSize: 10
                }

            }
        },
        created() {
            this.getImage(this.searchData)
            if (this.imageTabType !== 1) {
                this.flag = false
            } else {
                this.searchForm.push(
                    { type: 'Input', label: '用户名', prop: 'userNameLike', placeholder: '请输入用户名' },
                    { type: 'Input', label: '群组名', prop: 'spaceNameLike', placeholder: '请输入群组名' }
                )
            }
        },
        methods: {
            // 错误码
            getErrorMsg(code) {
                return getErrorMsg(code)
            },
            getImage(data) {
                this.type = this.imageTabType
                if (this.type === 1) {
                    getUserImage(data).then(response => {
                        if (response.success) {
                            this.tableData = response.data.images;
                            this.total = response.data.totalSize
                        } else {
                            this.$message({
                                message: this.getErrorMsg(response.error.subcode),
                                type: 'warning'
                            });
                        }
                    })
                }
                if (this.type === 2) {
                    getPreImage(data).then(response => {
                        if (response.success) {
                            if (response.data !== null) {
                                this.total = parseInt(response.data.totalSize)
                                this.tableData = response.data.images
                            }
                        } else {
                            this.$message({
                                message: this.getErrorMsg(response.error.subcode),
                                type: 'warning'
                            });
                        }
                    })
                }
            },
            handleEdit(row) {
                this.row = row
                this.FormVisible = true
                this.Logo = false
            },
            handleDelete(row) {
                deletePreImage(row.id).then(response => {
                    if (response.success) {
                        this.$message({
                            message: '删除成功',
                            type: 'success'
                        });
                        this.getImage(this.searchData)
                    } else {
                        this.$message({
                            message: this.getErrorMsg(response.error.subcode),
                            type: 'warning'
                        });
                    }
                })
            },
            handleSizeChange(val) {
                this.searchData.pageSize = val
                this.getImage(this.searchData)
            },
            handleCurrentChange(val) {
                this.searchData.pageIndex = val
                this.getImage(this.searchData)
            },
            cancel(val) {
                this.FormVisible = val
                this.getImage(this.searchData)
            },
            confirm(val) {
                this.FormVisible = val
                this.getImage(this.searchData)
            },
            close(val) {
                this.FormVisible = val
                this.getImage(this.searchData)
            },
            create() {
                this.FormVisible = true; this.row = {}
                this.Logo = true
            },
            getSearchData(val) {
                this.searchData = { pageIndex: 1, pageSize: this.searchData.pageSize }
                this.searchData = Object.assign(val, this.searchData)
                this.getImage(this.searchData)
            },
            // 时间戳转换日期
            parseTime(val) {
                return parseTime(val)
            },
            // 镜像状态
            imageStatus(value) {
                switch (value) {
                    case 1:
                        return '未制作'
                    case 2:
                        return '制作中'
                    case 3:
                        return '制作完成'
                    case 4:
                        return '制作失败'
                }
            },
            // 修改描述
            open(val) {
                const data = val
                this.$prompt('请输入描述', '提示', {
                    confirmButtonText: '确定',
                    cancelButtonText: '取消'
                }).then(({ value }) => {
                    editePreImage({ id: data.id, imageName: data.imageName, imageVersion: data.imageVersion, imageType: data.imageType, imageAddr: data.imageAddr, imageDesc: value }).then(response => {
                        if (response.success) {
                            this.$message({
                                message: '编辑描述成功',
                                type: 'success'
                            });
                            this.getImage(this.searchData)
                            this.$emit('confirm', false)
                        } else {
                            this.$message({
                                message: this.getErrorMsg(response.error.subcode),
                                type: 'warning'
                            });
                        }
                    })
                }).catch(() => {
                    this.$message({
                        type: 'info',
                        message: '取消输入'
                    });
                });
            },
            // 删除确认
            open2(val) {
                this.$confirm('此操作将永久删除该镜像, 是否继续?', '提示', {
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
            }
            // 群组详情
            // getGroupDetail(id) {
            //     groupDetail(id).then(response => {
            //         if (response.success) {
            //             console.log(response.data.workspace.name)
            //             return response.data.workspace.name
            //         }
            //         else { return '' }

            //     })
            // }

        }
    }
</script>
<style lang="scss" scoped>
    .create {
        float: right;
    }

    .block {
        float: right;
        margin: 20px;
    }

    .searchForm {
        display: inline-block;
    }
</style>