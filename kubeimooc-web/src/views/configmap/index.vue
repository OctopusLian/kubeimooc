<template>
  <div class="app-container">
    <div class="data-operation">
      <div class="data-operation-item" style="width: 20%">
        <el-input placeholder="输入关键词" v-model="keyword">
          <el-button slot="append" @click="filterNode()" icon="el-icon-search"></el-button>
        </el-input>
      </div>
      <div class="data-operation-item">
        <el-button type="primary" @click="redirectTo('/volume/configmap-edit')">新增</el-button>
      </div>
    </div>
    <el-table :data="list">
      <el-table-column width="25">
        <template slot-scope="{row}">
          <i style="color: #67C23A" class="el-icon-success"></i>
        </template>
      </el-table-column>
      <el-table-column label="name" width="150" prop="name">
      </el-table-column>
      <el-table-column label="data" width="100" prop="dataNum"></el-table-column>
      <el-table-column label="age" width="100" prop="age">
        <template slot-scope="{row}" style="cursor: pointer">
          <timestamp :data="row.age"></timestamp>
        </template>
      </el-table-column>
      <el-table-column label="operation">
        <template slot-scope="{row}">
          <a class="operation-class" @click="goEdit(row.name)">编辑</a>
          <a class="operation-class" @click="doDel(row.name)">删除</a>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>
<script>
import Timestamp from "@/components/Timestamp";
import {Message} from "element-ui";

export default {
  name: "Index",
  components: {
    Timestamp
  },
  data() {
    return {
      keyword: "",
      list: []
    }
  },
  created() {
    this.getNodeItemOrList()
  },
  methods: {
    doDel(name) {
      let params = {
        name: name,
        namespace: this.$store.state.ns.nsName
      }
      this.$confirm('此操作将永久删除该ConfigMap, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        this.$store.dispatch("cm/deleteCm", params).then(res => {
          Message.success(res.msg)
          //刷新当先页面(延迟3秒)
          this.getNodeItemOrList()
        })
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '已取消删除'
        });
      });
    },
    getNodeItemOrList() {
      let params = {
        namespace: this.$store.state.ns.nsName
      }
      this.$store.dispatch("cm/getCmItemOrList", params).then(res => {
        this.list = res.data
      })
    },
    filterNode() {
      let params = {
        namespace: this.$store.state.ns.nsName,
        keyword: this.keyword
      }
      this.$store.dispatch("cm/getCmItemOrList", params).then(res => {
        this.list = res.data
      })
    },
    goEdit(name) {
      this.redirectTo("/volume/configmap-edit?name=" + name)
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
