<template>
  <el-select class="ns-select" v-model="nsVal" @change="getValue"  placeholder="Namespace">
    <el-option  v-for="(item) in nsOpts" :label="item" :value="item"></el-option>
  </el-select>
</template>

<script>
export default {
  name: "Namespace",
  data() {
    return {
      nsVal: "default",
      nsOpts: []
    }
  },
  created() {
    this.initNs()
  },
  methods: {
    initNs() {
      //加载当前命名空间
      this.$store.dispatch("ns/getCurNs").then(res => {
        this.nsVal = res
      })
      this.$store.dispatch("ns/getNamespaceList").then(res => {
        let data = res.data
        data.forEach((item) => {
          this.nsOpts.push(item.name)
        })
      })
    },
    getValue() {
      this.$store.dispatch("ns/updateNsName", this.nsVal)
    }
  }
}
</script>
<style scoped>
/*.ns-select {*/
/*  width: 100%;*/
/*  border: 0;*/
/*}*/
.ns-select{
  /*position:absolute;*/
  width:120px;
  height:28px;
  line-height:28px;
  font-size:15px;
  font-weight:bold;
}

</style>
