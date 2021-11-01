<template>
    <div>
        <el-row>
            <el-col :span="12">
                <div>任务名称:<span>{{ data.name }}</span></div>
            </el-col>
            <el-col :span="12">
                <div>是否分布式:<span>{{ data.isDistributed?'是':'否' }}</span></div>
            </el-col>
        </el-row>
        <el-row>
            <el-col v-if="show" :span="10">
                <div>子任务名:<el-select v-model="value" placeholder="请选择" class="select" @change="selectLog">
                        <el-option
                            v-for="item in options"
                            :key="item.value+item.label"
                            :label="item.label"
                            :value="item.value"
                        />
                    </el-select>
                </div>
            </el-col>
        </el-row>
        <div>
            <el-row>
                <el-col v-if="showLog" :span="6">
                    <div>
                        任务日志:
                        <a
                            :href="href+'/log/user/trainjob/'+data.id+'/'+subName+'/index.log'"
                            download="日志.text"
                            class="download"
                            target="_blank"
                        >
                            下载
                        </a>
                    </div>
                </el-col>
            </el-row>

        </div>
        <div>
            <el-row>
                <el-input v-model="textarea" type="textarea" :rows="20" />
            </el-row>
        </div>
    </div>
</template>

<script>
    import { showLog, getTraningDetail } from '@/api/trainingManager.js'
    import { getErrorMsg } from '@/error/index'
    export default {
        name: "TaskLog",
        props: {
            row: {
                type: Object,
                default: () => { }
            }

        },
        data() {
            return {
                data: {},
                options: [],
                value: '',
                textarea: '',
                // logName: '',
                subName: '',
                flag: true,
                resource: [],
                showLog: false,
                href: '',
                timer: null,
                timer2: null,
                status: "running"

            }
        },
        computed: {
            show: function() {
                if (this.data.isDistributed === true) {
                    return true
                } else {
                    return false
                }
            }
        },
        created() {
            this.data = JSON.parse(JSON.stringify(this.row))
            this.href = window.location.protocol + '//' + window.location.host
            if (!this.data.isDistributed) {
                // this.logName = this.data.config[0].name
                this.subName = 'task0/0'
                this.getState()
            } else {
                for (let i = 0; i < this.data.config.length; i++) {
                    for (let j = 0; j < this.data.config[i].taskNumber; j++) {
                        this.options.push({ label: this.data.config[i].replicaStates[j].key, value: 'task' + i + "/" + j })
                    }
                }
            }
        },
        destroyed() {
            this.flag = false
            clearInterval(this.timer2);
            this.timer2 = null;
            clearTimeout(this.timer);
            this.timer = null;
        },
        methods: {
            // 错误码
            getErrorMsg(code) {
                return getErrorMsg(code)
            },
            selectLog() {
                this.subName = this.value
                this.getState()
            },
            getLog(status) {
                const data = { jobId: this.data.id, subName: this.subName }
                if (status === 'stopped' || status == 'succeeded') {
                    showLog(data).then(response => {
                        if (response) {
                            this.showLog = true
                            this.textarea = response.data
                        } else {
                            this.textarea = ''
                        }
                    })
                } else {
                    showLog(data).then(response => {
                        if (response) {
                            this.showLog = true
                            this.textarea = response.data
                            if (this.timer !== null) {
                                clearTimeout(this.timer);
                                this.timer = null;
                            }
                            if (this.timer2 !== null) {
                                clearInterval(this.timer2);
                                this.timer2 = null;
                            }

                            if (this.flag && this.status === 'running') {
                                this.timer = setTimeout(this.getLog, 1000);
                                this.timer2 = setInterval(this.getState, 1000);
                            }
                        } else {
                            clearInterval(this.timer2);
                            this.timer2 = null;
                            clearTimeout(this.timer);
                            this.timer = null;
                            this.textarea = ''
                        }
                    })
                }
            },
            getState() {
                getTraningDetail(this.data.id).then(response => {
                    if (response.success) {
                        this.status = response.data.trainJob.status
                        this.getLog(this.status)
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
    .el-col {
        margin: 10px 0 20px 0;
        font-size: 15px;
        font-weight: 800;

        span {
            font-weight: 400;
            margin-left: 20px
        }
    }

    .taskList {
        font-weight: 800;
    }

    .download {
        color: #409EFF;
        margin-left: 10px;
    }

    .select {
        margin-left: 5px;
    }
</style>