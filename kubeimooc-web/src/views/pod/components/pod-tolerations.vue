<template>
  <el-form-item label="tolerations">
    <el-row v-for="(item,index) in tolerations" class="for-item">
      <el-col :span="3">
        <el-input v-model="item.key" placeholder="key"></el-input>
      </el-col>
      <el-col :span="3">
        <el-select v-model="item.operator" placeholder="operator" style="width: 100%">
          <el-option value="Exists" label="Exists"></el-option>
          <el-option value="Equal" label="Equal"></el-option>
        </el-select>
      </el-col>
      <el-col :span="3">
        <el-input v-model="item.value" placeholder="value"></el-input>
      </el-col>
      <el-col :span="3">
        <el-select v-model="item.effect">
          <el-option value="NoSchedule"></el-option>
          <el-option value="PreferNoSchedule"></el-option>
          <el-option value="NoExecute"></el-option>
        </el-select>
      </el-col>
      <el-col v-if="item.effect==='NoExecute'" :span="6">
        <el-input-number v-model="item.tolerationSeconds" placeholder="tolerationSeconds"></el-input-number>
      </el-col>
      <el-button @click="tolerations.splice(index,1)"
                 size="small"
                 type="danger">
        delete
        <i class="el-icon-delete"></i>
      </el-button>
      <!--      <el-col :span="3">-->
      <!--        <el-button @click="tolerations.splice(index,1)"-->
      <!--                   size="small"-->
      <!--                   type="danger">-->
      <!--          delete-->
      <!--          <i class="el-icon-delete"></i>-->
      <!--        </el-button>-->
      <!--      </el-col>-->
    </el-row>
    <el-button @click="tolerations.push({})"
               size="small"
               type="primary">
      add toleration
      <i class="el-icon-circle-plus"></i>
    </el-button>
  </el-form-item>
</template>

<script>
export default {
  props: {
    tolerations: {
      type: Array
    }
  },
  watch: {
    tolerations: {
      handler(val) {
        this.$emit("tolerationsChange", val)
      },
      deep: true
    }
  },
  name: "PodTolerations",
}
</script>

<style scoped>

</style>
