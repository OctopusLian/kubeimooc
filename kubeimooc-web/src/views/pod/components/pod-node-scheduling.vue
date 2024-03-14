<template>
  <div>
    <el-form-item label-width="120px" label="nodeScheduling">
      <el-radio v-model="nodeScheduling.type" label="nodeAny">nodeAny</el-radio>
      <el-radio v-model="nodeScheduling.type" label="nodeName">nodeName</el-radio>
      <el-radio v-model="nodeScheduling.type" label="nodeSelector">nodeSelector</el-radio>
      <el-radio v-model="nodeScheduling.type" label="nodeAffinity">nodeAffinity</el-radio>
      <el-form-item v-if="nodeScheduling.type==='nodeName'" label="nodeName">
        <el-select v-model="nodeScheduling.nodeName" placeholder="select node" style="width: 100%">
          <el-option v-for="(item,_) in nodeList" :value="item" :label="item"/>
        </el-select>
      </el-form-item>
      <el-form-item v-if="nodeScheduling.type==='nodeSelector'" label="nodeSelector">
        <el-row v-for="(item,index) in nodeScheduling.nodeSelector" class="for-item">
          <el-col :span="10" style="padding-right: 5px">
            <el-input v-model="item.key" placeholder="label key"></el-input>
          </el-col>
          <el-col :span="1" style="text-align: center">-</el-col>
          <el-col :span="10">
            <el-input v-model="item.value" placeholder="label value"></el-input>
          </el-col>
          <el-col :span="3" style="text-align: center">
            <el-button @click="nodeScheduling.nodeSelector.splice(index,1)"
                       size="small"
                       type="danger">
              delete
              <i class="el-icon-delete"></i>
            </el-button>
          </el-col>
        </el-row>
        <el-button @click="addNodeSelector()"
                   size="small"
                   type="primary">
          add label
          <i class="el-icon-circle-plus"></i>
        </el-button>
      </el-form-item>
      <el-form-item v-if="nodeScheduling.type==='nodeAffinity'" label="nodeAffinity">
        <el-row v-for="(item,index) in nodeScheduling.nodeAffinity" class="for-item">
          <el-col :span="5" style="margin-right: 0.5rem">
            <el-input v-model="item.key" placeholder="key"></el-input>
          </el-col>
          <el-col :span="2" style="margin-right: 0.5rem">
            <el-select v-model="item.operator" placeholder="operator">
              <el-option value="In" label="In"></el-option>
              <el-option value="NotIn" label="NotIn"></el-option>
              <el-option value="Exists" label="Exists"></el-option>
              <el-option value="DoesNotExist" label="DoesNotExist"></el-option>
              <el-option value="Gt" label="Gt"></el-option>
              <el-option value="Lt" label="Lt"></el-option>
            </el-select>
          </el-col>
          <el-col :span="5">
            <el-input v-model="item.values" placeholder="values"></el-input>
          </el-col>
          <el-col :span="3" style="text-align: center">
            <el-button @click="nodeScheduling.nodeAffinity.splice(index,1)"
                       size="small"
                       type="danger">
              delete
              <i class="el-icon-delete"></i>
            </el-button>
          </el-col>
        </el-row>
        <el-button @click="nodeScheduling.nodeAffinity.push({})"
                   size="small"
                   type="primary">
          add term
          <i class="el-icon-circle-plus"></i>
        </el-button>
      </el-form-item>
    </el-form-item>
  </div>
</template>

<script>
export default {
  name: "podNodeScheduling",
  props: {
    nodeScheduling: {
      type: Object
    }
  },
  watch: {
    nodeScheduling: {
      handler(val) {
        this.$emit("nodeSchedulingChange", val)
      },
      deep: true
    }
  },
  data() {
    return {
      nodeList: [],
    }
  },
  created() {
    this.$store.dispatch("node/getNodeItemOrList", {}).then(res => {
      for (let i = 0; i < res.data.length; i++) {
        this.nodeList.push(res.data[i].name)
      }
    })
  },
  methods: {
    addNodeSelector() {
      if (this.nodeScheduling.nodeSelector == null) {
        this.nodeScheduling.nodeSelector = []
      }
      this.nodeScheduling.nodeSelector.push({})
    }
  }
}
</script>

<style scoped>

</style>
