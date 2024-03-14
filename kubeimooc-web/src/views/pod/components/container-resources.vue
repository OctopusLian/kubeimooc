<template>
  <div>
    <el-form-item label-width="120px" label="resources">
      <el-switch v-model="resources.enable" @change="changeResources()"></el-switch>
      <div v-if="resources.enable">
        <el-form-item label="cpuRequest">
          <el-slider
            v-model="resources.cpuRequest"
            :max="define.resourceCpuMax"
            :step="define.resourceCpuStep"
            show-tooltip
            show-input>
          </el-slider>
        </el-form-item>
        <el-form-item label="cpuLimit">
          <el-slider
            v-model="resources.cpuLimit"
            :max="define.resourceCpuMax"
            :step="define.resourceCpuStep"
            show-tooltip
            show-input>
          </el-slider>
        </el-form-item>
        <el-form-item label="memRequest">
          <el-slider
            v-model="resources.memRequest"
            :max="define.resourceMemMax"
            :step="define.resourceMemStep"
            show-tooltip
            show-input>
          </el-slider>
        </el-form-item>
        <el-form-item label="memLimit">
          <el-slider
            v-model="resources.memLimit"
            :max="define.resourceMemMax"
            :step="define.resourceMemStep"
            show-tooltip
            show-input>
          </el-slider>
        </el-form-item>
      </div>
    </el-form-item>
  </div>
</template>

<script>
// 容器资源定义
export default {
  name: "ContainerResources",
  props: {
    index: {
      type: Number
    },
    resources: {
      type: Object
    }
  },
  watch: {
    resources: {
      handler(val) {
        this.$emit("resourcesChange", val, this.index)
      },
      deep: true
    }
  },
  data() {
    return {
      define: {
        resourceCpuMax: 6000,
        resourceMemMax: 8192,
        resourceCpuStep: 10,
        resourceMemStep: 128
      }
    }
  },
  methods: {
    changeResources() {
      if (this.resources.enable) {
        this.resources = {
          enable: true,
          //设置cpu和内存的默认值
          cpuRequest: 1000,
          cpuLimit: 1000,
          memRequest: 512,
          memLimit: 512
        }
      }
    }
  }
}
</script>

<style scoped>

</style>
