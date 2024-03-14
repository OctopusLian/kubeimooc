<template>
  <div class="card-container">
    <div class="card-header">
      {{ title }}
    </div>
    <div class="card-body">
      <el-row>    <!--:key必须写-->
        <el-col v-for="(value, key) in data" :key="key"
                :span="12" style="padding: 15px 0"
                v-if="key!=='labels' && key!=='taints'">
          <strong>{{ key }}:</strong>
          <span v-if="key==='age'">
            <timestamp :data="value"></timestamp>
          </span>
          <span v-else>
            {{ value }}
          </span>
        </el-col>
      </el-row>
    </div>
  </div>
</template>
<script>
import Timestamp from "@/components/Timestamp";

export default {
  components: {
    Timestamp
  },
  props: {
    title: {
      type: String,
    },
    data: {
      type: Object,
      default: () => {
      }
    }
  },
  name: "NodeInfoCard",
  created() {
    console.log(this.data)
  }
}
</script>

<style scoped>
</style>
