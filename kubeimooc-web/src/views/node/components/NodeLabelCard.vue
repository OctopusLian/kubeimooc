<template>
  <div class="card-container">
    <div class="card-header">{{ title }}</div>
    <div class="card-body">
      <el-row v-for="(item,index) in data" class="for-item">
        <el-col :span="8" style="padding-right: 5px">
          <el-input v-model="item.key" placeholder="label key"></el-input>
        </el-col>
        <!--        <el-col :span="1" class="intev-line">-</el-col>-->
        <el-col :span="8">
          <el-input v-model="item.value" placeholder="label value"></el-input>
        </el-col>
        <el-col :span="6" style="text-align: left;padding-left: 1em">
          <el-button @click="data.splice(index,1)"
                     size="small"
                     type="danger">
            delete
            <i class="el-icon-delete"></i>
          </el-button>
          <el-button v-if="index===data.length-1" @click="data.push({})"
                     size="small"
                     type="primary">
            add
            <i class="el-icon-circle-plus"></i>
          </el-button>
        </el-col>
      </el-row>
      <el-button @click="submit()"
                 size="small"
                 type="primary">
        提交
        <i class="el-icon-success"></i>
      </el-button>
    </div>
  </div>
</template>

<script>
import {Message} from 'element-ui'
export default {
  props: {
    title: {
      type: String
    },
    data: {
      type: Array
    }
  },
  name: "NodeLabelCard",
  methods: {
    submit() {
      this.$emit("updateLabel")
    }
  }
}
</script>

<style scoped>

</style>
