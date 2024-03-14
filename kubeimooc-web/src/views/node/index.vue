<template>
  <div class="app-container">
    <div class="data-operation">
      <div class="data-operation-item" style="width: 20%">
        <el-input placeholder="输入关键词" v-model="keyword">
          <el-button slot="append" @click="filterNode()" icon="el-icon-search"></el-button>
        </el-input>
      </div>
    </div>
    <el-table :data="nodeList">
      <el-table-column width="25">
        <template slot-scope="{row}">
          <i v-if="row.status==='Ready'" style="color: #67C23A" class="el-icon-success"></i>
          <i v-else style="color:#F56C6C" class="el-icon-warning"></i>
        </template>
      </el-table-column>
      <el-table-column label="name" width="100" prop="name">
      </el-table-column>
      <el-table-column label="version" width="100" prop="version"></el-table-column>
      <!--      <el-table-column label="images" :show-overflow-tooltip="true" prop="images"></el-table-column>-->
      <el-table-column label="internalIp" width="120" prop="internalIp"></el-table-column>
      <el-table-column label="externalIp" width="120" prop="externalIp"></el-table-column>
      <el-table-column label="osImage" width="150" prop="osImage"></el-table-column>
      <el-table-column label="kernelVersion" width="150" prop="kernelVersion"></el-table-column>
      <el-table-column label="containerRuntime" width="150" prop="containerRuntime"></el-table-column>
      <el-table-column label="age" prop="age">
        <template slot-scope="{row}" style="cursor: pointer">
          <timestamp :data="row.age"></timestamp>
        </template>
      </el-table-column>
      <el-table-column label="operation">
        <template slot-scope="{row}">
          <a class="operation-class" @click="editNode(row.name)">编辑</a>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>
<script>
import Timestamp from "@/components/Timestamp";

export default {
  name: "Index",
  components: {
    Timestamp
  },
  data() {
    return {
      keyword: "",
      nodeList: []
    }
  },
  created() {
    this.getNodeItemOrList()
  },
  methods: {
    getNodeItemOrList() {
      let params = {
        namespace: this.$store.state.ns.nsName
      }
      this.$store.dispatch("node/getNodeItemOrList", params).then(res => {
        this.nodeList = res.data
      })
    },
    filterNode() {
      let params = {
        namespace: this.$store.state.ns.nsName,
        keyword: this.keyword
      }
      this.$store.dispatch("node/getNodeItemOrList", params).then(res => {
        this.nodeList = res.data
      })
    },
    editNode(name) {
      this.$store.dispatch("node/setNodeName", name)
      this.redirectTo("/node/edit?name=" + name)
    },
    redirectTo(url) {
      this.$router.push(
        {
          path: url
        }
      )
    },
  }
}
</script>
<style scoped>
.data-operation-item {
  display: inline-block;
}

.data-operation div:nth-child(n+2) {
  margin-left: 5px;
}

.operation-class {
  margin: 5px;
}

.operation-class:hover {
  color: #409EFF;
}
</style>
