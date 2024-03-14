<template>
  <div>
    <el-form :model="form" ref="form" label-width="150px">
      <el-form-item :label="label" prop="containers"
                    :rules="label==='containers'?rules.containers:rules.initContainers">
        <template v-slot:label>
          {{ label }}
          <el-tooltip class="item" effect="dark"
                      content="containers"
                      placement="bottom">
            <span class="el-icon-info"></span>
          </el-tooltip>
        </template>
        <template slot="error" slot-scope="{ error }">
          <span style="color: #F56C6C" v-if="error">{{ error }}</span>
        </template>
        <el-collapse v-model="activeNames" @change="handleChange">
          <el-collapse-item v-for="(item,index) in form.containers" :name="item.name">
            <template slot="title" class="container-item">
              <el-button @click="form.containers.splice(index,1)"
                         size="small"
                         type="danger">
                delete
                <i class="el-icon-delete"></i>
              </el-button>
              <el-tag v-if="item.isInit" style="margin-left: 5px">
                name:{{ item.name }}(init)
              </el-tag>
              <el-tag v-else style="margin-left: 5px">
                name:{{ item.name }}
              </el-tag>
              <el-tag style="margin-left: 5px">
                image:{{ item.image === '' || item.image === undefined ? 'none' : item.image }}
              </el-tag>
            </template>
            <el-form-item label="image"
                          label-width="120px"
                          :prop="'containers.'+index+'.image'" :rules="rules.image">
              <el-input v-model="item.image" placeholder="please input image"></el-input>
            </el-form-item>
            <el-form-item
              label-width="120px"
              label="imagePullPolicy"
              :rules="rules.imagePullPolicy"
              :prop="'containers.'+index+'.imagePullPolicy'"
            >
              <el-select v-model="item.imagePullPolicy"
                         style="width: 100%"
                         placeholder="please input imagePullPolicy">
                <el-option v-for="item in define.imagePullPolicyOpts"
                           :label="item"
                           :key="item"
                           :value="item"></el-option>
              </el-select>
            </el-form-item>
            <el-form-item label-width="120px" label="workingDir">
              <el-input v-model="item.workingDir" placeholder="容器工作目录"></el-input>
            </el-form-item>
            <el-form-item label-width="120px" label="command">
              <el-row v-for="(cItem,cIndex) in item.command" class="for-item">
                <el-col :span="20">
                  <el-input v-model="item.command[cIndex]" placeholder="please input command"></el-input>
                </el-col>
                <el-col :span="4" style="text-align: center">
                  <el-button @click="form.containers[index].command.splice(cIndex,1)"
                             size="small"
                             type="danger">
                    delete
                    <i class="el-icon-delete"></i>
                  </el-button>
                </el-col>
              </el-row>
              <el-button @click="addCommand(index)"
                         size="mini"
                         type="primary">
                add
                <i class="el-icon-circle-plus"></i>
              </el-button>
            </el-form-item>
            <el-form-item label-width="120px" label="args">
              <el-row v-for="(argItem,argIndex) in form.containers[index].args" class="for-item">
                <el-col :span="20">
                  <el-input v-model="form.containers[index].args[argIndex]" placeholder="please input arg"></el-input>
                </el-col>
                <el-col :span="4" style="text-align: center">
                  <el-button @click="form.containers[index].args.splice(argIndex,1)"
                             size="small"
                             type="danger">
                    delete
                    <i class="el-icon-delete"></i>
                  </el-button>
                </el-col>
              </el-row>
              <el-button @click="addArgs(index)"
                         size="mini"
                         type="primary">
                add
                <i class="el-icon-circle-plus"></i>
              </el-button>
            </el-form-item>
            <el-form-item label-width="120px" label="tty">
              <el-switch v-model="item.tty"></el-switch>
            </el-form-item>
            <el-form-item label-width="120px" label="privileged">
              <el-switch v-model="item.privileged"></el-switch>
            </el-form-item>
            <el-form-item label-width="120px" label="envs">
              <el-row v-for="(itemEnv,indexEnv) in item.envs" class="for-item">
                <el-col :span="4" style="padding-right: 0.5rem">
                  <el-input v-model="itemEnv.name" placeholder="env name"></el-input>
                </el-col>
                <el-col :span="4" style="padding-right: 0.5rem">
                  <el-select v-model="itemEnv.type">
                    <el-option value="default"/>
                    <el-option value="configMap"/>
                    <el-option value="secret"/>
                  </el-select>
                </el-col>
                <el-col v-if="itemEnv.type!=='default'" :span="4" style="padding-right: 0.5rem">
                  <el-select v-model="itemEnv.refName" placeholder="envRef">
                    <el-option v-for="(itemEnvRef,_) in envRef[itemEnv.type]" :value="itemEnvRef"></el-option>
                  </el-select>
                </el-col>
                <el-col :span="4" style="padding-right: 0.5rem">
                  <el-input v-model="itemEnv.value" placeholder="env value/refKey"></el-input>
                </el-col>
                <el-col :span="2" style="text-align: center">
                  <el-button @click="form.containers[index].envs.splice(indexEnv,1)"
                             size="small"
                             type="danger">
                    delete
                    <i class="el-icon-delete"></i>
                  </el-button>
                </el-col>
              </el-row>
              <el-button @click="addEnv(index)"
                         size="mini"
                         type="primary">
                add
                <i class="el-icon-circle-plus"></i>
              </el-button>
            </el-form-item>
            <el-form-item label-width="120px" label="envsFrom">
              <el-row v-for="(itemEnvFrom,indexEnvFrom) in item.envsFrom" class="for-item">
                <el-col :span="4" style="padding-right: 0.5rem">
                  <el-input v-model="itemEnvFrom.prefix" placeholder="envFrom prefix"></el-input>
                </el-col>
                <el-col :span="4" style="padding-right: 0.5rem">
                  <el-select v-model="itemEnvFrom.refType">
                    <el-option value="configMap"/>
                    <el-option value="secret"/>
                  </el-select>
                </el-col>
                <el-col :span="4" style="padding-right: 0.5rem">
                  <el-select v-model="itemEnvFrom.name" placeholder="envRef">
                    <el-option v-for="(itemEnvFromRef,_) in envRef[itemEnvFrom.refType]"
                               :value="itemEnvFromRef"></el-option>
                  </el-select>
                </el-col>
                <el-col :span="2" style="text-align: center">
                  <el-button @click="form.containers[index].envsFrom.splice(indexEnv,1)"
                             size="small"
                             type="danger">
                    delete
                    <i class="el-icon-delete"></i>
                  </el-button>
                </el-col>
              </el-row>
              <el-button @click="addEnvFrom(index)"
                         size="mini"
                         type="primary">
                add
                <i class="el-icon-circle-plus"></i>
              </el-button>
            </el-form-item>
            <el-form-item label-width="120px" label="ports">
              <el-row v-for="(iitem,iitemIndex) in item.ports" class="for-item">
                <el-col :span="7">
                  <el-input v-model="iitem.name" placeholder="port name"></el-input>
                </el-col>
                <el-col :span="7" style="padding: 0 10px">
                  <el-input-number style="width: 100%" v-model="iitem.hostPort"
                                   placeholder="host port"></el-input-number>
                </el-col>
                <el-col :span="7" style="padding: 0 10px">
                  <el-input-number style="width: 100%" v-model="iitem.containerPort"
                                   placeholder="container port"></el-input-number>
                </el-col>
                <el-col :span="2" style="text-align: center">
                  <el-button @click="form.containers[index].ports.splice(iitemIndex,1)"
                             size="small"
                             type="danger">
                    delete
                    <i class="el-icon-delete"></i>
                  </el-button>
                </el-col>
              </el-row>
              <el-button @click="form.containers[index].ports.push({})"
                         size="mini"
                         type="primary">
                add
                <i class="el-icon-circle-plus"></i>
              </el-button>
            </el-form-item>
            <container-probe :probe="item.startupProbe" :index="index" label="startupProbe"
                             @probeChange="probeChange"/>
            <container-probe :probe="item.livenessProbe" :index="index" label="livenessProbe"
                             @probeChange="probeChange"/>
            <container-probe :probe="item.readinessProbe" :index="index"
                             label="readinessProbe"
                             @probeChange="probeChange"/>
            <container-resources @resourcesChange="resourcesChange" :resources="item.resources" :index="index"/>
            <el-form-item label="volumeMounts">
              <el-collapse v-model="item.volumeMountsActiveNames">
                <el-collapse-item v-for="(itemMount,itemMountIndex) in item.volumeMounts" :name="itemMount.id">
                  <template slot="title" class="container-item">
                    <!--                    卷为空时指定-->
                    <i class="el-icon-delete my-op-icon"
                       style="float: right;margin-right: 10px"
                       @click.stop="form.containers[index].volumeMounts.splice(itemMountIndex, 1)"
                    />
                    <span style="margin-right: 20px"><strong
                      style="margin-right: 10px">mountName</strong>{{
                        itemMount.mountName === undefined ? 'none' : itemMount.mountName
                      }}</span>
                    <span><strong style="margin-right: 10px">mountPath</strong>
                      {{ itemMount.mountPath === undefined ? 'none' : itemMount.mountPath }}</span>
                  </template>
                  <el-form-item label="volumeMountName">
                    <el-select v-model="itemMount.mountName" placeholder="选择卷" style="width: 100%">
                      <el-option v-for="volumeItem in volumes" v-if="volumeItem.name!==undefined"
                                 :label="volumeItem.name+'('+volumeItem.type+')'"
                                 :value="volumeItem.name"></el-option>
                    </el-select>
                  </el-form-item>
                  <el-form-item label="volumeMountPath">
                    <el-input v-model="itemMount.mountPath"></el-input>
                  </el-form-item>
                  <el-form-item label="readOnly">
                    <el-switch v-model="itemMount.readOnly"></el-switch>
                  </el-form-item>
                </el-collapse-item>
              </el-collapse>
              <el-form-item>
                <el-button @click="addVolumeMount(index)"
                           size="mini"
                           type="primary">
                  <i class="el-icon-circle-plus"></i>
                  add volumeMount
                </el-button>
              </el-form-item>
            </el-form-item>
          </el-collapse-item>
        </el-collapse>
        <el-form-item>
          <el-button
            @click="openAddContainerDialog"
            size="small" icon="el-icon-circle-plus"
            type="primary">add container
          </el-button>
        </el-form-item>
        <el-form-item>
          <el-divider/>
        </el-form-item>
      </el-form-item>
      <el-dialog
        title="add container"
        :visible.sync="dialogContainerVisible"
        width="30%">
        <create-container-dialog :containers.sync="form.containers" ref="containerDialog" @addContainer="addContainer"/>
      </el-dialog>
    </el-form>

  </div>
</template>


<script>
import ContainerProbe from "@/views/pod/components/container-probe";
import CreateContainerDialog from "@/views/pod/components/create-container-dialog";
import ContainerResources from "@/views/pod/components/container-resources";

export default {
  components: {
    CreateContainerDialog,
    ContainerProbe,
    ContainerResources,
  },
  props: {
    base: {
      type: Object,
      default: () => {
      }
    },
    volumes: {
      type: Array,
      default: () => []
    },
    define: {
      type: Object,
      default: () => {
      }
    },
    label: {
      type: String,
      default: ""
    },
    data: {
      type: Array,
      default: () => []
    }
  },
  name: "Containers",
  created() {
    this.loadRef()
  },
  data() {
    const containersValidator = (rule, value, callback) => {
      //校验容器定义是否为空
      if (value.length === 0) {
        callback(new Error("container is empty"))
      } else {
        callback()
      }
    }
    return {
      formValidatePass: false,
      activeNames: [],
      form: {
        containers: [],
      },
      envRef: {
        configMap: [],
        secret: [],
        configMapKeys: {},
        secretKeys: {}
      },
      rules: {
        //init-containers 可以不定义
        initContainers: [],
        containers: [
          {required: true, message: 'please add container', trigger: 'blur'},
          {
            type: 'string',
            required: true,
            validator: containersValidator,
            trigger: ['blur']
          }
        ],
        image: [
          {required: true, message: 'please input container image', trigger: 'blur'}
        ],
        imagePullPolicy: [
          {required: true, message: 'please input container imagePullPolicy', trigger: 'blur'}
        ]
      },
      dialogContainerVisible: false
    }
  },
  watch: {
    data(val) {
      this.form.containers = val
      //打开折叠窗口
      for (let i = 0; i < val.length; i++) {
        this.activeNames.push(val[i].name)
      }
    }
  },
  methods: {
    handleChange(val) {
      // console.log(val);
    },
    openAddContainerDialog() {
      this.dialogContainerVisible = true
    },
    addVolumeMount(indexContainer) {
      //计算该卷的挂载id
      let volumeMountId = this.form.containers[indexContainer].volumeMounts.length
      this.form.containers[indexContainer].volumeMounts.push({
        id: volumeMountId
      })
      this.form.containers[indexContainer].volumeMountsActiveNames.push(volumeMountId)
    },
    addContainer(name) {
      this.form.containers.push({
        name: name,
        imagePullPolicy: "IfNotPresent",
        workingDir: "",
        tty: false,
        privileged: false,
        command: [],
        args: [],
        envs: [],
        envsFrom: [],
        //容器端口
        ports: [],
        volumeMounts: [],
        volumeMountsActiveNames: [],
        resources: {
          enable: false,
          //设置cpu和内存的默认值
          cpuRequest: 1000,
          cpuLimit: 1000,
          memRequest: 512,
          memLimit: 512
        },
        startupProbe: {
          enable: false,
          type: "http",
          httpGet: {
            path: "/",
            port: 8080,
            scheme: "HTTP",
            httpHeaders: []
          },
          tcpSocket: {
            port: 8080
          },
          exec: {
            command: []
          },
          initialDelaySeconds: 5,
          periodSeconds: 10,
          timeoutSeconds: 1,
          successThreshold: 1,
          failureThreshold: 3
        },
        livenessProbe: {
          enable: false,
          type: "http",
          initialDelaySeconds: 5,
          periodSeconds: 10,
          timeoutSeconds: 1,
          successThreshold: 1,
          failureThreshold: 3,
          tcpSocket: {
            port: 8080
          },
          exec: {
            command: []
          },
          httpGet: {
            path: "/",
            port: 8080,
            scheme: "HTTP",
            httpHeaders: []
          },
        },
        readinessProbe: {
          enable: false,
          type: "http",
          initialDelaySeconds: 5,
          periodSeconds: 10,
          timeoutSeconds: 1,
          successThreshold: 1,
          failureThreshold: 3,
          tcpSocket: {
            port: 8080
          },
          exec: {
            command: []
          },
          httpGet: {
            path: "/",
            port: 8080,
            scheme: "HTTP",
            httpHeaders: []
          },
        }
      })
      this.dialogContainerVisible = false
      //打开折叠窗口
      this.activeNames.push(name)
      //清除之前的表单验证
      this.$refs["form"].clearValidate()
    },
    onSubmit(formName) {
      this.$refs[formName].validate((valid) => {
        if (!valid) {
          //打开所有的折叠窗口
          for (let i = 0; i < this.form.containers.length; i++) {
            if (this.activeNames.indexOf(this.form.containers[i].name) === -1) {
              this.activeNames.push(this.form.containers[i].name)
            }
          }
          this.formValidatePass = false
          return false
        }
        this.formValidatePass = true
      })
    },
    resourcesChange(val, index) {
      this.form.containers[index].resources = val
    },
    addArgs(index) {
      if (this.form.containers[index].args == null) {
        this.form.containers[index].args = []
      }
      this.form.containers[index].args.push('')
    },
    addEnv(index) {
      this.form.containers[index].envs.push({
        name: "",
        refName: "",
        type: "default",
        value: ""
      })
    },
    addEnvFrom(index) {
      this.form.containers[index].envsFrom.push({
        prefix: "",
        refType: "configMap",
        name: ""
      })
    },
    loadRef() {
      //加载configmap
      this.envRef.configMap = []
      let params = {
        namespace: this.$store.state.ns.nsName,
      }
      this.$store.dispatch("cm/getCmItemOrList", params).then(res => {
        let data = res.data
        for (let i = 0; i < data.length; i++) {
          this.envRef.configMap.push(data[i].name)
          //加载keys
          params.name = data[i].name
          this.$store.dispatch("cm/getCmItemOrList", params).then(resDetail => {
            let dataList = resDetail.data.dataList
            let keys = []
            for (let j = 0; j < dataList.length; j++) {
              keys.push(dataList[j].key)
            }
            this.$set(this.envRef.configMapKeys, resDetail.data.name, keys)
          })
        }
      })
      //加载secret
      let params_secret = {
        namespace: this.$store.state.ns.nsName,
      }
      this.envRef.secret = []
      this.$store.dispatch("secret/getSecretItemOrList", params_secret).then(res => {
        let data = res.data
        for (let i = 0; i < data.length; i++) {
          this.envRef.secret.push(data[i].name)
          //加载keys
          params_secret.name = data[i].name
          this.$store.dispatch("secret/getSecretItemOrList", params_secret).then(resDetail => {
            let dataList = resDetail.data.dataList
            let keys = []
            for (let j = 0; j < dataList.length; j++) {
              keys.push(dataList[j].key)
            }
            this.$set(this.envRef.secretKeys, resDetail.data.name, keys)
            console.log(
              "xxxxxxx"
            )
            console.log(this.envRef.secretKeys)
          })
        }
      })
    },
    addCommand(index) {
      if (this.form.containers[index].command == null) {
        this.form.containers[index].command = []
      }
      this.form.containers[index].command.push('')
    },
    //监听子组件 探针数据变化情况，并赋值给父组件
    probeChange(val, label, index) {
      switch (label) {
        case "startupProbe": {
          this.form.containers[index].startupProbe = val
          break
        }
        case "livenessProbe": {
          this.form.containers[index].livenessProbe = val
          break
        }
        case "readinessProbe": {
          this.form.containers[index].readinessProbe = val
          break
        }
      }
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
