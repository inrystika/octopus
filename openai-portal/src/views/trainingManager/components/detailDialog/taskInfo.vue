<template>
    <div>
      <div v-html="this.initInfo"></div>
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
            initInfo: "",
          }
        },
        created() {
            let taskInfoString = this.row.initInfo ? this.row.initInfo.replace(/\n/g, "<br>") : ''
            let taskInfoData = JSON.parse(taskInfoString)
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
                this.initInfo = message
            }
        }
    }
</script>