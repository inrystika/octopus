<template>
    <div>
      <div v-html="this.initInfo"></div>
    </div>
</template>

<script>
    export default {
        name: "TaskInfo",
        props: {
            row: {
                type: Object,
                default: () => { }
            }
        },
        data() {
          return {
            initInfo: ""
          }
        },
        created() {
            const taskInfoString = this.row.initInfo ? this.row.initInfo.replace(/\n/g, "<br>") : ''
            const taskInfoData = JSON.parse(taskInfoString)
            for (const pid in taskInfoData['podEvents']) {
                const eventList = taskInfoData['podEvents'][pid]
                const roleName = taskInfoData['podRoleName'][pid]
                if (roleName == "") {
                    continue
                }
                let message = ""
                for (const key in eventList) {
                    const event = eventList[key]
                    if (event['reason'] == "" && event['message'] == "") {
                        continue
                    }
                    message += "[" + event['reason'] + "]" + "<br>"
                    message += event['message'] + "<br><br>"
                }
                for (const key in taskInfoData['extras']) {
                    const event = taskInfoData['extras'][key]
                    if (event['reason'] == "" && event['message'] == "") {
                        continue
                    }
                    message += "[" + event['reason'] + "]" + "<br>"
                    message += event['message'] + "<br><br>"
                }
                message += "<br>"
                this.initInfo = message
            }
        }
    }
</script>