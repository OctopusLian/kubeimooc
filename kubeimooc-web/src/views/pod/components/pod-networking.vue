<template>
  <div>
    <el-form-item label-width="120px" label="netWorking">
      <el-form-item label-width="120px" label="hostNetwork">
        <el-switch v-model="netWorking.hostNetwork"></el-switch>
      </el-form-item>
      <el-form-item label-width="120px" label="dnsPolicy">
        <el-select v-model="netWorking.dnsPolicy"
                   style="width: 100%"
                   placeholder="please input dnsPolicy">
          <el-option v-for="item in define.dnsPolicyOpts"
                     :label="item"
                     :key="item"
                     :value="item"></el-option>
        </el-select>
      </el-form-item>
      <el-form-item label-width="120px" label="dnsConfig">
        <el-form-item label-width="120px" label="nameservers">
          <el-row v-for="(cItem,cIndex) in netWorking.dnsConfig.nameservers" class="for-item">
            <el-col :span="20">
              <el-input v-model="netWorking.dnsConfig.nameservers[cIndex]"
                        placeholder="please input nameserver"></el-input>
            </el-col>
            <el-col :span="4" style="text-align: center">
              <el-button @click="netWorking.dnsConfig.nameservers.splice(cIndex,1)"
                         size="small"
                         type="danger">
                delete
                <i class="el-icon-delete"></i>
              </el-button>
            </el-col>
          </el-row>
          <el-button @click="netWorking.dnsConfig.nameservers.push('')"
                     size="mini"
                     type="primary">
            add
            <i class="el-icon-circle-plus"></i>
          </el-button>
        </el-form-item>
      </el-form-item>
      <el-form-item label-width="120px" label="hostName">
        <el-input v-model="netWorking.hostName"></el-input>
      </el-form-item>
      <el-form-item label-width="120px" label="hostAlias">
        <el-row v-for="(itemHostAlias,indexHostAlias) in netWorking.hostAliases"
                class="for-item">
          <el-col :span="9">
            <el-input v-model="itemHostAlias.key" placeholder="hostname"></el-input>
          </el-col>
          <el-col :span="1" style="text-align: center">-</el-col>
          <el-col :span="9">
            <el-input v-model="itemHostAlias.value"
                      placeholder="multiple IPs are separated by comma"></el-input>
          </el-col>
          <el-col :span="5" style="text-align: center">
            <el-button
              @click="netWorking.hostAliases.splice(indexHostAlias,1)"
              size="mini"
              type="danger">
              delete
              <i class="el-icon-delete"></i>
            </el-button>
          </el-col>
        </el-row>
        <el-button @click="netWorking.hostAliases.push({})"
                   size="mini"
                   type="primary">
          add
          <i class="el-icon-circle-plus"></i>
        </el-button>
      </el-form-item>
    </el-form-item>
  </div>
</template>

<script>
export default {
  props: {
    netWorking: {
      type: Object
    },
  },
  name: "PodNetworking",
  watch: {
    netWorking: {
      handler(val) {
        this.$emit("netWorkingChange", val)
      },
      deep: true
    }
  },
  data(){
    return {
      define: {
        dnsPolicyOpts: [
          "Default", "ClusterFirst", "ClusterFirstWithHostNet", "None"
        ],
      },
    }
  }
}
</script>

<style scoped>
.el-form-item {
  margin-bottom: 18px;
}

.for-item {
  margin-bottom: 4px;
}
</style>
