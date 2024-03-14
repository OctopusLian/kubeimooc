<template>
  <div>
    <el-form ref="form" :model="form" :rules="rules">
      <el-form-item label="container name" prop="name">
        <el-input v-model="form.name" placeholder="please input container name"/>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSubmit('form')">submit</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
//容器重名校验
export default {
  name: "CreateContainerDialog",
  props: {
    containers:
      {
        type: Array,
        default: () => []
      }
  },
  data() {
    const duplicateValidate = (rule, value, callback) => {
      //同步数据的写法，保持父子组件数据同步
      let containers = JSON.parse(JSON.stringify(this.containers))
      for(let i=0;i<containers.length;i++){
        let item = containers[i]
        if(item.name===value){
          callback(new Error("duplicate define container"))
          return
        }
      }
      callback()
    }
    return {
      form: {
        name: "",
        isInit: false
      },
      rules: {
        name: [
          {required: true, message: 'please input container name!', trigger: 'blur'},
          {
            type: 'string',
            required: true,
            validator: duplicateValidate,
            trigger: ['blur']
          }
        ]
      }
    }
  },
  methods: {
    resetVal() {
      this.form.name = ""
    },
    onSubmit(formName) {
      //1 参数校验
      this.$refs[formName].validate((valid) => {
        if (!valid) {
          return false
        } else {
          //2 调用父组件方法，传参+清空数据
          this.$emit("addContainer", this.form.name)
          this.resetVal()
        }
      })
    },
    close() {
      this.$emit("closeContainerDialog")
    }
  }
}
</script>

<style scoped>

</style>
