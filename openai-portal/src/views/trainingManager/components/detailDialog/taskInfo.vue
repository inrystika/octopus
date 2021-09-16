<template>
    <div>
        <div v-for="(item,key) in initInfo" :key="key">
            <h3>{{key}}</h3>
            <div v-html="item"></div>
        </div>
    </div>
</template>

<script>
    export default {
        name: "taskInfo",
        props: {
            row: {
                type: Object,
                default: () => { }
            },
        },
        data() {
          return {
            initInfo: {},         
          }
        },
        created() {
            let taskInfoString = this.row.initInfo ? this.row.initInfo.replace(/\n/g, "<br>") : ''
            let taskInfoData = JSON.parse(taskInfoString)
            let tempTaskInfoData = {}
            for(let pid in taskInfoData['podEvents']){       
                const eventList = taskInfoData['podEvents'][pid]
                const roleName = taskInfoData['podRoleName'][pid]
                if (roleName == "") {
                    continue
                }
                let message = ""
                for (let key in eventList) {
                    let event = eventList[key]
                    if (event['reason'] == "" && event['message'] == "") {
                        continue
                    }
                    message += "[" + event['reason'] + "]" + "<br>"
                    message += event['message'] + "<br><br>"
                }
                for (let key in taskInfoData['extras']) {               
                    let event = taskInfoData['extras'][key]               
                    if (event['reason'] == "" && event['message'] == "") {
                        continue
                    }
                    message +=  "[" + event['reason'] + "]" + "<br>"
                    message += event['message'] + "<br><br>"
                }
                message += "<br>"
                tempTaskInfoData[roleName] = message
            }
            let obj = {}
            Object.keys(tempTaskInfoData).sort().forEach(function(key) {
                obj[key] = tempTaskInfoData[key];
            });
            this.initInfo = obj 
        }
    }
</script>

<style lang="scss" scoped>
    .select {
        margin-left: 5px;
    }
</style>