<template>
  <div class="app-container">
    <el-form ref="form" :model="form" :rules="rules" label-position="left" label-width="150px" class="my-form">
      <el-form-item label="name" prop="base.name">
        <el-input :disabled="disabled" v-model="form.base.name" placeholder="pod name"/>
      </el-form-item>
      <el-form-item label="labels">
        <el-row v-for="(item,index) in form.base.labels" class="for-item">
          <el-col :span="10" style="padding-right: 5px">
            <el-input v-model="item.key" placeholder="label key"></el-input>
          </el-col>
          <el-col :span="1" style="text-align: center">-</el-col>
          <el-col :span="10">
            <el-input v-model="item.value" placeholder="label value"></el-input>
          </el-col>
          <el-col :span="3" style="text-align: center">
            <el-button @click="form.base.labels.splice(index,1)"
                       size="small"
                       type="danger">
              delete
              <i class="el-icon-delete"></i>
            </el-button>
          </el-col>
        </el-row>
        <el-button @click="form.base.labels.push({})"
                   size="small"
                   type="primary">
          add label
          <i class="el-icon-circle-plus"></i>
        </el-button>
      </el-form-item>
      <el-form-item label="imagePullSecrets" prop="base.imagePullSecrets">
        <el-select multiple placeholder="please select secret" v-model="form.base.imagePullSecrets" style="width: 100%">
          <el-option v-for="item in imagePullSecrets"
                     :key="item"
                     :label="item"
                     :value="item">
          </el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="restartPolicy" prop="base.restartPolicy">
        <el-select v-model="form.base.restartPolicy" placeholder="please select restart policy" style="width: 100%">
          <el-option v-for="item in restartPolicyOpts"
                     :key="item"
                     :value="item"
                     :label=item></el-option>
        </el-select>
      </el-form-item>
      <pod-networking :net-working="form.netWorking" @netWorkingChange="netWorkingChange"/>
      <pod-node-scheduling :node-scheduling="form.nodeScheduling"
                           @nodeSchedulingChange="nodeSchedulingChange"></pod-node-scheduling>
      <pod-tolerations :tolerations="form.tolerations" @tolerationsChange="tolerationsChange"></pod-tolerations>
      <el-form-item label="volumes">
        <el-row v-for="(item,index) in form.volumes" class="for-item">
          <el-col :span="10" style="padding-right: 5px">
            <el-select v-model="item.type" placeholder="volume type" style="display: block">
              <el-option v-for="item in define.volumeOpts"
                         :label="item"
                         :key="item"
                         :value="item"></el-option>
            </el-select>
          </el-col>
          <el-col :span="1" style="text-align: center">-</el-col>
          <el-col :span="10">
            <el-input v-model="item.name" placeholder="volume name"></el-input>
          </el-col>
          <el-col :span="3" style="text-align: center">
            <el-button @click="form.volumes.splice(index,1)"
                       size="small"
                       type="danger">
              delete
              <i class="el-icon-delete"></i>
            </el-button>
          </el-col>
        </el-row>
        <el-button @click.stop="form.volumes.push({})"
                   size="small"
                   type="primary">
          add volume
          <i class="el-icon-circle-plus"></i>
        </el-button>
      </el-form-item>
      <el-form-item>
        <el-divider/>
      </el-form-item>
      <container-list
        :label="'initContainers'"
        :base="form.base"
        :volumes="form.volumes"
        :data="form.initContainers"
        ref="initContainers"
        :define="define"/>
      <container-list
        :label="'containers'"
        :base="form.base"
        :volumes="form.volumes"
        :data="form.containers"
        ref="containers"
        :define="define"/>
      <el-form-item>
        <el-button @click="onSubmit('form')" type="primary">
          submit
        </el-button>
      </el-form-item>
    </el-form>
  </div>
</template>
<script>
import CreateContainerDialog from "./components/create-container-dialog";
import PodNetworking from "@/views/pod/components/pod-networking";
import PodNodeScheduling from "@/views/pod/components/pod-node-scheduling";
import PodTolerations from "@/views/pod/components/pod-tolerations";
import ContainerList from "./components/container"
import {Message} from 'element-ui'

export default {
  name: "PodCreate",
  components: {
    CreateContainerDialog,
    ContainerList,
    PodNetworking,
    PodNodeScheduling,
    PodTolerations
  },
  data() {
    return {
      //为true 表示为更新操作 名字禁止修改
      disabled: false,
      define: {
        volumeOpts: [
          "emptyDir"
        ],
        imagePullPolicyOpts: [
          "Always", "IfNotPresent", "Never"
        ],
      },
      restartPolicyOpts: [
        "Always", "On-Failure", "Never"
      ],
      imagePullSecrets: [],
      dialogLabelVisible: false,
      dialogContainerVisible: false,
      dialogEnvVisible: false,
      activeNames: [],
      form: {
        base: {
          name: "",
          labels: [],
          imagePullSecrets: [],
          restartPolicy: "Always",
          namespace: ""
        },
        volumes: [],
        nodeScheduling: {
          type: "nodeAny",
          nodeName: "",
          nodeSelector: [],
          nodeAffinity: [],
        },
        tolerations: [],
        netWorking: {
          hostNetwork: false,
          dnsConfig: {
            nameservers: []
          },
          dnsPolicy: "Default",
          hostAliases: [],
          hostName: ""
        },
        initContainers: [],
        containers: [],
      },
      rules: {
        "base.name": [
          {required: true, message: 'please input pod name', trigger: 'blur'}
        ],
        "base.restartPolicy": [
          {required: true, message: 'please input restartPolicy', trigger: 'blur'}
        ],
      }
    }
  },
  created() {
    this.load()
  },
  methods: {
    load() {
      //如果是编辑，则提前加载参数
      //获取url pod名称
      let name = this.$route.query.name
      if (name !== undefined) {
        this.disabled = true
        //请求详情
        let params = {
          namespace: this.$store.state.ns.nsName,
          name: name
        }
        this.$store.dispatch("pod/getPodItemOrList", params).then(res => {
          // this.form = res.data
          this.form = res.data
        })
      }
    },
    onSubmit(formName) {
      this.$refs[formName].validate((valid) => {
        if (!valid) {
          //校验容器子表单
          Message.error("提交失败，请检查必填型！")
          return false
        }
        //校验container表单，init-container可以不填写的
        this.$refs.initContainers.onSubmit("form")
        this.$refs.containers.onSubmit("form")
        //获取子组件变量校验是否表单验证通过
        if (!this.$refs.initContainers.formValidatePass || !this.$refs.containers.formValidatePass) {
          Message.error("提交失败，请检查必填项！")
          return false
        }
        //组件数据提交到后台
        this.form.initContainers = this.$refs.initContainers.form.containers
        this.form.containers = this.$refs.containers.form.containers
        this.form.base.namespace = this.$store.state.ns.nsName
        this.$store.dispatch("pod/createPod", this.form).then(response => {
          Message.success(response.msg)
          this.$router.push(
            {
              path: "/pod/list"
            }
          )
        })
      })
    },
    close() {
      this.dialogVisible = false
    },
    closeContainerDialog() {
      this.dialogContainerVisible = false
    },
    handleChange(val) {
      console.log(val);
    },
    netWorkingChange(val) {
      this.form.netWorking = val
    },
    nodeSchedulingChange(val) {
      console.log(val)
      this.form.nodeScheduling = val
    },
    tolerationsChange(val) {
      console.log(val)
      this.form.tolerations = val
    }
  }
}
</script>

<style scoped>
.my-form {
  width: 90%;
  margin-left: 5%;
}


.el-form-item {
  margin-bottom: 22px;
}

.for-item {
  margin-bottom: 5px;
}
</style>
