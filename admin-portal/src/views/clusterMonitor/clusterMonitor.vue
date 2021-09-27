<template>
  <div>
    <iframe
      src="http://192.168.202.73/grafana/d/ft1oaQnWk/clustermetrics?orgId=1&refresh=10s&from=now-5m&to=now&var-Node=All"
      :height="iFrameHeight"
    >
    </iframe>
  </div>
</template>

<script>
  export default {
    name: "clusterMonitor",
    components: {

    },
    data() {
      return {
        iFrameHeight: ""
      }
    },
    mounted() {
    var parent = window.parent.document.documentElement;
    var scrollHeight = parent.scrollHeight;
    this.iFrameHeight = scrollHeight + "px"
  },
    methods: {

    }
  }
</script>
<style lang="scss" scoped>

</style>