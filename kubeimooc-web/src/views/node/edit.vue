<template>
  <div class="app-container">
    <node-info-card :title="'节点信息'" :data="nodeInfo"/>
    <node-label-card :title="'节点标签(key/value)'" :data="nodeLabels" @updateLabel="updateLabel"
                     ref="nodeLabel"></node-label-card>
    <node-taint-card :title="'节点污点(key/value/effect)'" :data="nodeTaints" @updateTaint="updateTaint"
                     ref="nodeTaint"></node-taint-card>
    <node-pod-card :title="'节点Pod'" :data="nodeInfo"></node-pod-card>
  </div>
</template>

<script>
import NodeInfoCard from './components/NodeInfoCard'
import NodeLabelCard from "./components/NodeLabelCard";
import NodeTaintCard from "./components/NodeTaintCard";
import NodePodCard from "./components/NodePodCard";
import {Message} from "element-ui";

export default {
  name: "Edit",
  components: {
    NodeInfoCard,
    NodeLabelCard,
    NodeTaintCard,
    NodePodCard
  },
  data() {
    return {
      nodeInfo: {},
      nodeLabels: [],
      nodeTaints: []
    }
  },
  created() {
    this.getNodeItemOrList()
  },
  methods: {
    getNodeItemOrList() {
      let params = {
        name: this.$route.query.name
      }
      this.$store.dispatch("node/getNodeItemOrList", params).then(res => {
        this.nodeInfo = res.data
        this.nodeLabels = res.data.labels
        this.nodeTaints = res.data.taints
      })
    },
    updateLabel() {
      //提交数据到后台
      let data = {
        name: this.nodeInfo.name,
        labels: this.$refs.nodeLabel.data
      }
      this.$store.dispatch("node/updateNodeLabel", data).then(res => {
        Message.success(res.msg)
      })
    },
    updateTaint() {
      //提交数据到后台
      let data = {
        name: this.nodeInfo.name,
        taints: this.$refs.nodeTaint.data
      }
      this.$store.dispatch("node/updateNodeTaint", data).then(res => {
        Message.success(res.msg)
      })
    }
  }
}
</script>

<style scoped>
</style>
