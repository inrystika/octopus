<template>
  <div class="wrapper">
    <iframe :src="grafanaUri" :height="iFrameHeight" frameBorder="0" scrolling="no" class="clusterDashboard">
    </iframe>
  </div>
</template>
<script>
  // import clusterMonitor from "./clusterMonitor.vue";
  // import groupMonitor from "./groupMonitor.vue";
  export default {
    // components: {
    //   clusterMonitor,
    //   groupMonitor
    // },
    data() {
      return {
        activeName: 'first',
        iFrameHeight: "",
        grafanaUri: this.GLOBAL.DOMAIN + "/grafana/d/ft1oaQnWk/clustermetrics?orgId=1&refresh=10s&from=now-5m&to=now&var-Node=All"
      }
    },
    mounted() {
      var parent = window.parent.document.documentElement;
      var scrollHeight = parent.scrollHeight;
      this.iFrameHeight = scrollHeight + "px"
      window.addEventListener('beforeunload', e => {
        sessionStorage.clear()
      });

    },
    destroyed() {
      window.removeEventListener('beforeunload', e => {
        sessionStorage.clear()
      })
    },
    // mounted() {
    //   var parent = window.parent.document.documentElement;
    //   var scrollHeight = parent.scrollHeight;
    //   this.iFrameHeight = scrollHeight + "px"
    // },
    methods: {
      handleClick(tab, event) {
      }
    }
  }
</script>
<style lang="scss" scoped>
  .wrapper {
    background-color: black;
    min-height: 100vh;
    padding-right: 50px;
  }

  .clusterDashboard {
    width: 100%;
    border: none;
    margin-left: 10px;
  }
</style>