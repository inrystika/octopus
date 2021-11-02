<template>
    <div>
        <div class="searchForm">
            <searchForm :search-form="searchForm" :blur-name="'镜像名称/标签/描述 搜索'" @searchData="getSearchData" />
        </div>
        <el-button v-if="flag" type="primary" class="create" @click="create">创建</el-button>
        <el-table :data="tableData" style="width: 100%;font-size: 15px;"
            :header-cell-style="{'text-align':'left','color':'black'}" :cell-style="{'text-align':'left'}">
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
            <el-table-column label="镜像地址" align="center" :show-overflow-tooltip="true">
                <template slot-scope="scope">
                    <span>{{ scope.row.sourceType===2?scope.row.imageAddr:'' }}</span>
                </template>
            </el-table-column>
            <el-table-column label="镜像描述" align="center" :show-overflow-tooltip="true">
                <template slot-scope="scope">
                    <span>{{ scope.row.imageDesc }}</span>
                </template>
            </el-table-column>
            <el-table-column label="镜像类型" align="center">
                <template slot-scope="scope">
                    <span>{{ imageType(scope.row.imageType) }}</span>
                </template>
            </el-table-column>
            <!-- <el-table-column label="镜像状态" align="center">
                <template slot-scope="scope">
                    <span>{{ imageStatus(scope.row.imageStatus) }}</span>
                </template>
            </el-table-column> -->
            <el-table-column label="创建时间" align="center">
                <template slot-scope="scope">
                    <span>{{ parseTime(scope.row.createdAt) }}</span>
                </template>
            </el-table-column>
            <el-table-column v-if="imageTabType==3" label="提供者" align="center">
                <template slot-scope="scope">
                    <span>{{ scope.row.username }}</span>
                </template>
            </el-table-column>
            <el-table-column label="镜像来源" align="center">
                <template slot-scope="scope">
                    <span>{{ sourceType(scope.row.sourceType) }}</span>
                </template>
            </el-table-column>
            <el-table-column label="状态" align="center" v-if="flag">
                <template slot-scope="scope">
                    <span v-if="!(scope.row.progress&&scope.row.progress!=0)">{{ imageStatus(scope.row.imageStatus)
                        }}</span>
                    <span v-if="scope.row.progress&&scope.row.progress!=0">{{ "上传中" }}</span>
                    <el-progress :percentage="parseInt(scope.row.progress-1)"
                        v-if="scope.row.progress&&scope.row.progress!=0"></el-progress>
                </template>
            </el-table-column>
            <el-table-column v-if="flag" label="操作" align="center" :width="250">
                <template slot-scope="scope">
                    <el-button v-if="scope.row.imageStatus==1||scope.row.imageStatus==4" type="text"
                        @click="handleEdit(scope.row)" :disabled="scope.row.progress&&scope.row.progress!=0">重新上传
                    </el-button>
                    <el-button type="text" @click="open2(scope.row)"
                        :disabled="scope.row.progress&&scope.row.progress!=0">删除</el-button>
                    <!-- <el-button @click="handleDelete(scope.row)" type="text">删除</el-button> -->
                    <el-button v-if="!scope.row.isShared&&scope.row.imageStatus===3" type="text"
                        @click="open3(scope.row)">分享
                    </el-button>
                    <el-button v-if="scope.row.isShared&&scope.row.imageStatus===3" type="text"
                        @click="open4(scope.row)">取消分享
                    </el-button>
                    <el-button type="text" :close-on-click-modal="false" @click="open(scope.row)">修改描述</el-button>
                </template>
            </el-table-column>
        </el-table>
        <div class="block">
            <el-pagination :current-page="searchData.pageIndex" :page-sizes="[10, 20, 50, 80]"
                :page-size="searchData.pageSize" layout="total, sizes, prev, pager, next, jumper" :total="total"
                @size-change="handleSizeChange" @current-change="handleCurrentChange" />
        </div>
        <!-- 镜像对话框 -->
        <dialogForm v-if="FormVisible" :row="row" :flag="Logo" @cancel="cancel" @confirm="confirm" @close="close" />

    </div>
</template>

<script>
    import searchForm from '@/components/search/index.vue'
    import dialogForm from "./components/dialogForm.vue";
    import { getMyImage, getPublicImage, getPreImage, deleteImage, shareImage, editeImage, cancelImage } from '@/api/imageManager.js'
    import { parseTime } from '@/utils/index'
    import { getErrorMsg } from '@/error/index'
    import store from '@/store'
    export default {
        name: "PreImage",
        components: {
            dialogForm,
            searchForm
        },
        props: {
            imageTabType: { type: Number, default: undefined },
            status: { type: Boolean },
            image: { type: Boolean, default: false }
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
                searchForm: [{ type: 'Input', label: '镜像名称', prop: 'imageNameLike', placeholder: '请输入镜像名称' },
                {
                    type: 'Select', label: '镜像类型', prop: 'imageType', placeholder: '请选择镜像类型',
                    options: [{ label: 'Notebook', value: 1 }, { label: '训练类型', value: 2 }]
                },
                {
                    type: 'Select', label: '状态', prop: 'imageStatus', placeholder: '请输入状态',
                    options: [{ label: '未制作', value: 1 }, { label: '制作中', value: 2 }, { label: '制作完成', value: 3 }, { label: '制作失败', value: 4 }]
                },
                { type: 'Input', label: '镜像标签', prop: 'imageVersion', placeholder: '请输入镜像标签' },
                {
                    type: 'Select', label: '上传类型', prop: 'sourceType', placeholder: '请选择镜像版本',
                    options: [{ label: '上传', value: 1 }, { label: '远程', value: 2 }]
                }

                ],
                searchData: {
                    pageIndex: 1,
                    pageSize: 10
                },
                timer: null
            }
        },
        created() {
            this.getImage(this.searchData)
            if (this.imageTabType !== 1) {
                this.flag = false

            }
            if (this.imageTabType == 1) {
                this.timer = setInterval(() => { this.getImage(this.searchData) }, 1000)
            }
            if (this.image) {
                this.FormVisible = true
            }
        },
        mounted() {
            window.addEventListener('beforeunload', e => {
                sessionStorage.clear()
            });

        },
        destroyed() {
            window.removeEventListener('beforeunload', e => {
                sessionStorage.clear()
            })
            clearInterval(this.timer)
            this.timer = null
        },
        methods: {
            // 错误码
            getErrorMsg(code) {
                return getErrorMsg(code)
            },
            getImage(data) {
                this.type = this.imageTabType
                if (this.type === 1) {
                    getMyImage(data).then(response => {
                        if (response.success) {
                            if (response.data.images !== null) {
                                this.total = response.data.totalSize
                                const data = response.data.images
                                this.tableData = []
                                data.forEach(item => {
                                    this.tableData.push({ ...item.image, isShared: item.isShared })
                                })
                                this.tableData.forEach(item => {
                                    if (sessionStorage.getItem(JSON.stringify(item.id))) {
                                        item.progress = sessionStorage.getItem(JSON.stringify(item.id))
                                    }

                                })
                            }
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
                            if (response.data.images != null) {
                                this.total = response.data.totalSize
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
                if (this.type === 3) {
                    getPublicImage(data).then(response => {
                        if (response.success) {
                            if (response.data.images !== null) {
                                this.total = response.data.totalSize
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
            imageType(value) {
                switch (value) {
                    case 1:
                        return 'NoteBook类型'
                    default:
                        return '训练类型'
                }
            },
            sourceType(value) {
                switch (value) {
                    case 1:
                        return '上传'
                    default:
                        return '远程'
                }
            },
            imageStatus(value) {
                switch (value) {
                    case 1:
                        return '未上传'
                    case 2:
                        return '制作中'
                    case 3:
                        return '制作完成'
                    case 4:
                        return '制作失败'
                }
            },
            handleEdit(row) {
                this.row = row
                this.FormVisible = true
                this.Logo = false
                store.commit('user/SET_PROGRESSID', row.id)
            },
            handleDelete(row) {
                deleteImage(row.id).then(response => {
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
            handleShare(row) {
                shareImage(row.id).then(response => {
                    if (response.success) {
                        this.$message({
                            message: '分享成功',
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
            // 修改描述
            open(val) {
                const data = val
                this.$prompt('请输入描述', '提示', {
                    confirmButtonText: '确定',
                    cancelButtonText: '取消',
                    closeOnClickModal: false
                }).then(({ value }) => {
                    editeImage({ id: data.id, imageName: data.imageName, imageVersion: data.imageVersion, imageType: data.imageType, imageAddr: data.imageAddr, imageDesc: value }).then(response => {
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
                this.$confirm('此操作将永久删除该镜像（如该镜像已分享，则分享的镜像也会被删除)，是否继续?', '提示', {
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
            // 分享取消分享
            open3(val) {
                this.$confirm('此操作将分享镜像, 是否继续?', '提示', {
                    confirmButtonText: '确定',
                    cancelButtonText: '取消',
                    type: 'warning',
                    center: true
                }).then(() => {
                    this.handleShare(val)
                }).catch(() => {
                    this.$message({
                        type: 'info',
                        message: '已取消分享'
                    });
                });
            },
            // 取消分享
            open4(val) {
                this.$confirm('此操作将取消镜像分享, 是否继续?', '提示', {
                    confirmButtonText: '确定',
                    cancelButtonText: '取消',
                    type: 'warning',
                    center: true
                }).then(() => {
                    this.cancelShare(val)
                }).catch(() => {
                    this.$message({
                        type: 'info',
                        message: '已取消分享'
                    });
                });
            },
            cancelShare(val) {
                cancelImage(val.id).then(response => {                
                    if (response.success) {
                        this.$message({
                            message: '取消分享成功',
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

            }

        }
    }
</script>
<style lang="scss" scoped>
    .block {
        float: right;
        margin: 20px;
    }

    .create {
        float: right;
    }

    .searchForm {
        display: inline-block;
    }
</style>