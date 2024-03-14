<template>
  <div class="app-container">
    <div class="data-operation">
      <div class="data-operation-item" style="width: 20%">
        <el-input placeholder="输入关键词" v-model="keyword">
          <el-button slot="append" @click="filterPod()" icon="el-icon-search"></el-button>
        </el-input>
      </div>
      <div class="data-operation-item">
        <el-button type="primary" @click="redirectTo('/pod/create')">新增</el-button>
      </div>
    </div>
    <el-table :data="podList">
      <el-table-column width="25">
        <template slot-scope="{row}">
          <i v-if="row.status==='Running'" style="color: #67C23A" class="el-icon-success"></i>
          <i v-else style="color:#F56C6C" class="el-icon-warning"></i>
        </template>
      </el-table-column>
      <el-table-column label="name" width="235" prop="name">
      </el-table-column>
      <!--      <el-table-column label="labels" width="200" prop="labels">-->
      <!--        <template slot-scope="{row}">-->
      <!--          <el-tag class="my-tag" v-for="(value,key,index) in row.labels">{{ key }}:{{ value }}</el-tag>-->
      <!--        </template>-->
      <!--      </el-table-column>-->
      <!--      <el-table-column label="namespace" prop="namespace"></el-table-column>-->
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
      <el-table-column label="operation">
        <template slot-scope="{row}">
          <a class="operation-class" @click="editPod(row.name)">编辑</a>
          <a class="operation-class" @click="deletePod(row.name)">删除</a>
        </template>

        <!--        <a class="operation-class">日志</a>-->
        <!--        <a class="operation-class">进入</a>-->

      </el-table-column>
    </el-table>
  </div>
</template>
<script>
import Timestamp from "@/components/Timestamp";
import {Message} from "element-ui";

export default {
  name: "PodIndex",
  components: {
    Timestamp
  },
  data() {
    return {
      keyword: "",
      podList: []
    }
  },
  created() {
    this.getPodItemOrList()
  },
  watch: {
    '$store.state.ns.nsName'(newVal) {
      this.getPodItemOrList()
    }
  },
  methods: {
    filterPod() {
      let params = {
        namespace: this.$store.state.ns.nsName,
        keyword: this.keyword
      }
      this.$store.dispatch("pod/getPodItemOrList", params).then(res => {
        this.podList = res.data
      })
    },
    getPodItemOrList() {
      let params = {
        namespace: this.$store.state.ns.nsName
      }
      this.$store.dispatch("pod/getPodItemOrList", params).then(res => {
        this.podList = res.data
      })
    },
    deletePod(name) {
      let params = {
        namespace: this.$store.state.ns.nsName,
        name: name
      }
      this.$confirm('此操作将永久删除该Pod, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        this.$store.dispatch("pod/deletePod", params).then(res => {
          Message.success(res.msg)
          //刷新当先页面(延迟3秒)
          this.getPodItemOrList()
        })
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '已取消删除'
        });
      });
    },
    editPod(name) {
      this.redirectTo("/pod/create?name=" + name)
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
