<template>
  <div class="card-container">
    <div class="card-header">{{ title }}</div>
    <div class="card-body">
      <div class="app-container">
        <el-table :data="podList">
          <el-table-column width="25">
            <template slot-scope="{row}">
              <i v-if="row.status==='Running'" style="color: #67C23A" class="el-icon-success"></i>
              <i v-else style="color:#F56C6C" class="el-icon-warning"></i>
            </template>
          </el-table-column>
          <el-table-column label="name" width="235" prop="name">
          </el-table-column>
          <el-table-column label="node" width="100" prop="node"></el-table-column>
          <!--      <el-table-column label="images" :show-overflow-tooltip="true" prop="images"></el-table-column>-->
          <el-table-column label="ready" width="100" prop="ready"></el-table-column>
          <el-table-column label="restarts" width="100" prop="restarts"></el-table-column>
          <el-table-column label="IP" prop="IP"></el-table-column>
          <el-table-column label="age" prop="age">
            <template slot-scope="{row}" style="cursor: pointer">
              <timestamp :data="row.age"></timestamp>
            </template>
          </el-table-column>
        </el-table>
      </div>
    </div>
  </div>
</template>

<script>
import Timestamp from "@/components/Timestamp";

export default {
  name: "NodePodCard",
  components: {
    Timestamp
  },
  props: {
    title: {
      type: String
    },
    data: {
      type: Object
    }
  },
  data() {
    return {
      podList: []
    }
  },
  watch: {
    '$store.state.ns.nsName'(newVal) {
      this.getPodItemOrList()
    }
  },
  created() {
    this.getPodItemOrList()
  },
  methods: {
    getPodItemOrList() {
      let params = {
        namespace: this.$store.state.ns.nsName,
        nodeName: this.$route.query.name
      }
      this.$store.dispatch("pod/getPodItemOrList", params).then(res => {
        this.podList = res.data
      })
    },
  }
}
</script>

<style scoped>

</style>
