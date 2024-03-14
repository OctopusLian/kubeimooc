<template>
  <div class="card-container">
    <div class="card-header">{{ title }}</div>
    <div class="card-body">
      <el-row v-for="(item,index) in data" class="for-item">
        <el-col :span="8" style="padding-right: 5px">
          <el-input v-model="item.key" placeholder="taint key"></el-input>
        </el-col>
        <el-col :span="4">
          <el-input v-model="item.value" placeholder="taint value"></el-input>
        </el-col>
        <el-col :span="4">
          <el-select v-model="item.effect" style="width: 100%;margin-left: 0.5rem">
            <el-option value="NoSchedule"></el-option>
            <el-option value="PreferNoSchedule"></el-option>
            <el-option value="NoExecute"></el-option>
          </el-select>
        </el-col>
        <el-col :span="6" style="text-align: left;padding-left: 1em">
          <el-button @click="data.splice(index,1)"
                     size="small"
                     type="danger">
            delete
            <i class="el-icon-delete"></i>
          </el-button>
          <el-button v-if="index===data.length-1" @click="data.push({
          effect:'NoSchedule'
          })"
                     size="small"
                     type="primary">
            add
            <i class="el-icon-circle-plus"></i>
          </el-button>
        </el-col>
      </el-row>
      <el-button v-if="data!==null && data.length!==0" @click="submit()"
                 size="small"
                 type="primary">
        提交
        <i class="el-icon-success"></i>
      </el-button>
      <el-button v-else @click="firstAdd()"
                 size="small"
                 type="primary">
        add
        <i class="el-icon-circle-plus"></i>
      </el-button>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    title: {
      type: String
    },
    data: {
      type: Array
    }
  },
  name: "NodeTaintCard",
  created() {
  },
  methods: {
    firstAdd() {
      this.data = []
      this.data.push({
        effect: 'NoSchedule'
      })
    },
    submit() {
      this.$emit("updateTaint")
    }
  }
}
</script>

<style scoped>

</style>
