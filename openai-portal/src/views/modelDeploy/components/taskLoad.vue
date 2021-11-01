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
            <el-col v-if="show" :span="12">
                <div>子任务名:<el-select v-model="value" placeholder="请选择" class="select" @change="selectLoad">
                        <el-option v-for="item in options" :key="item.value" :label="item.label" :value="item.value" />
                    </el-select>
                </div>
            </el-col>
        </el-row>
        <div>
            <el-row v-if="value!==''||!data.isDistributed">
                <el-col :span="24">
                    <div class="loadName">任务负载</div>
                    <iframe :src="loadHref" class="load" frameBorder="0" scrolling="no" />
                </el-col>
            </el-row>

        </div>
    </div>
</template>
<script>
    export default {
        name: "TaskLoad",
        components: {
        },
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
                pod: '',
                jobid: '',
                loadHref: '',
                href: ''

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
            // 本地调试

            // if (this.href == 'localhost') {
            //     this.href = 'http://192.168.202.73'
            // }
            if (!this.data.isDistributed) {
                this.pod = this.data.id + '-task0-0'
                this.loadHref = this.href + '/grafana/d/TK8iV8nWk/taskmetrics?orgId=1&refresh=10s&var-pod=' + this.pod + '&var-pod_name=' + this.pod
            } else {
                this.options = []
                for (let i = 0; i < this.data.config.length; i++) {
                    // this.data.config[i].subName = 'task' + i
                    for (let j = 0; j < this.data.config[i].taskNumber; j++) {
                        this.data.config[i].subName = '-task' + i + '-' + j
                        this.options.push({ label: this.data.config[i].replicaStates[j].key, value: this.data.config[i].subName })
                    }
                }
            }
        },
        methods: {
            selectLoad() {
                this.pod = this.data.id + this.value
                this.loadHref = this.href + '/grafana/d/TK8iV8nWk/taskmetrics?orgId=1&refresh=10s&var-pod=' + this.pod + '&var-pod_name=' + this.pod
                // alert(this.loadHref)
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

    .load {
        width: 100%;
        min-height: 1000px;
        margin-top: 30px;
        overflow: hidden;

    }

    .loadName {
        text-align: left;
        font-size: 18px;
        font-weight: 1000;
    }

    .select {
        margin-left: 5px;
    }
</style>