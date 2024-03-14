<template>
  <div>
    <el-form-item label-width="120px" :label="label">
      <el-switch v-model="probe.enable" @change="changeProbe()"></el-switch>
      <div v-if="probe.enable">
        <el-form-item label="startupProbeType">
          <el-select style="width: 100%"
                     v-model="probe.type"
                     placeholder="please select probe type">
            <el-option v-for="probeTypeOpt in define.probeTypeOpts"
                       :label="probeTypeOpt"
                       :value="probeTypeOpt"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item v-if="probe.type==='exec'" label="exec">
          <el-form-item label="command">
            <el-row v-for="(cItem,cIndex) in probe.exec.command">
              <el-col :span="16" style="padding: 0 10px">
                <el-input v-model="probe.exec.command[cIndex]" class="for-item"></el-input>
              </el-col>
              <el-col :span="4" style="text-align: center">
                <el-button
                  @click="probe.exec.command.splice(cIndex,1)"
                  size="mini"
                  type="danger">
                  delete
                  <i class="el-icon-delete"></i>
                </el-button>
              </el-col>
            </el-row>
            <el-button @click="addExecCommand()"
                       size="mini"
                       type="primary">
              add command
              <i class="el-icon-circle-plus"></i>
            </el-button>
          </el-form-item>
        </el-form-item>
        <el-form-item v-if="probe.type==='tcp'" label="tcpSocket">
          <el-form-item label="host">
            <el-input v-model="probe.tcpSocket.host"></el-input>
          </el-form-item>
          <el-form-item label="port">
            <el-input-number v-model="probe.tcpSocket.port"></el-input-number>
          </el-form-item>
        </el-form-item>
        <el-form-item v-if="probe.type==='http'" label="httpGet">
          <el-form-item label="scheme">
            <el-select v-model="probe.httpGet.scheme">
              <el-option v-for="item in define.httpTypeOpts" :value="item" :label="item"></el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="path">
            <el-input v-model="probe.httpGet.path"></el-input>
          </el-form-item>
          <el-form-item label="port">
            <el-input-number v-model="probe.httpGet.port"></el-input-number>
          </el-form-item>
          <el-form-item label="httpHeaders">
            <el-row v-for="(itemHttpGetHeader,indexHttpGetHeader) in probe.httpGet.httpHeaders"
                    class="for-item">
              <el-col :span="9">
                <el-input v-model="itemHttpGetHeader.key"></el-input>
              </el-col>
              <el-col :span="1" style="text-align: center">-</el-col>
              <el-col :span="9">
                <el-input v-model="itemHttpGetHeader.value"></el-input>
              </el-col>
              <el-col :span="5" style="text-align: center">
                <el-button
                  @click="probe.httpGet.httpHeaders.splice(indexHttpGetHeader,1)"
                  size="mini"
                  type="danger">
                  delete
                  <i class="el-icon-delete"></i>
                </el-button>
              </el-col>
            </el-row>
            <el-button @click="probe.httpGet.httpHeaders.push({})"
                       size="mini"
                       type="primary">
              add
              <i class="el-icon-circle-plus"></i>
            </el-button>
          </el-form-item>
        </el-form-item>
        <el-form-item label="initialDelaySeconds">
          <el-slider
            v-model="probe.initialDelaySeconds"
            :max="define.probeMax"
            :step="define.probeStep"
            show-tooltip
            show-input>
          </el-slider>
        </el-form-item>
        <el-form-item label="periodSeconds">
          <el-slider
            v-model="probe.periodSeconds"
            :max="define.probeMax"
            :step="define.probeStep"
            show-tooltip
            show-input>
          </el-slider>
        </el-form-item>
        <el-form-item label="timeoutSeconds">
          <el-slider
            v-model="probe.timeoutSeconds"
            :max="define.probeMax"
            :step="define.probeStep"
            show-tooltip
            show-input>
          </el-slider>
        </el-form-item>
        <el-form-item label="successThreshold">
          <el-slider
            v-model="probe.successThreshold"
            :max="define.probeMax"
            :step="define.probeStep"
            show-tooltip
            show-input>
          </el-slider>
        </el-form-item>
        <el-form-item label="failureThreshold">
          <el-slider
            v-model="probe.failureThreshold"
            :max="define.probeMax"
            :step="define.probeStep"
            show-tooltip
            show-input>
          </el-slider>
        </el-form-item>
      </div>
    </el-form-item>
  </div>
</template>

<script>
// 容器探针组件
export default {
  props: {
    probe: {
      type: Object,
      default: () => {
      }
    },
    label: {
      type: String
    },
    //容器索引坐标
    index: {
      type: Number
    }
  },
  name: "ContainerProbe",
  data() {
    return {
      define: {
        probeMax: 300,
        probeStep: 1,
        probeTypeOpts: [
          "tcp", "http", "exec"
        ],
        httpTypeOpts: [
          "HTTP", "HTTPS"
        ]
      }
    }
  },
  watch: {
    probe: {
      handler(val) {
        this.$emit("probeChange", val, this.label, this.index)
      },
      deep: true
    }
  },
  methods: {
    changeProbe() {
      if (this.probe.enable) {
        this.probe = {
          enable: true,
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
        }
      }
    },
    addExecCommand() {
      if (this.probe.exec.command == null) {
        this.probe.exec.command = []
      }
      this.probe.exec.command.push('')
    }
  }
}
</script>

<style scoped>
.for-item {
  margin-bottom: 4px;
}

.el-form-item {
  margin-bottom: 18px;
}
</style>
