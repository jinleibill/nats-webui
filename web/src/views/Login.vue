<script setup>
import {Close, Delete, Edit, Link} from "@element-plus/icons-vue";
import {onMounted, ref, watch} from "vue";
import {ElNotification} from "element-plus";
import request from "@/request.js";
import {useRouter} from "vue-router";
import utils from "@/utils.js";
import imgLogo from "@/assets/nats-logo.png";

const addConnectSpan = ref(24)
const isAddConnect = ref(false)
watch(isAddConnect, function(newVal) {
  if (newVal) {
    connectForm.value.name = ""
    connectForm.value.host = ""
    connectForm.value.port = ""
    addConnectSpan.value = 16
  } else {
    addConnectSpan.value = 24
  }
})

const connectData = ref([])
const connectFormRef = ref()
const connectForm = ref({
  name: "",
  host: "",
  port: ""
})
const connectFormRules = ref({
  name: [
    {required: true, message: "名称必填", trigger: "blur"},
  ],
  host: [
    {required: true, message: "地址必填", trigger: "blur"},
  ],
  port: [
    {required: true, message: "端口必填", trigger: "blur"},
  ],
})

const router = useRouter()

onMounted(()=>{
  getConnects()
})

const getConnects = async () => {
  const response = await request.get("/nodes");
  if (response.code) return
  let connects = []
  response.data.forEach(element => {
    connects.push({
      name: element.name,
      host: element.host,
      port: element.port,
      createAt: utils.formatData(element.create_at),
    })
  })
  connectData.value = connects
}

const onAuth = async (item) => {
  const response = await request.post("/auths", {
    url: `${item.host}:${item.port}`,
  });
  if (response.code) return
  localStorage.setItem("isAuth", true)
  await router.push("/")
}

const onConnect = async (formEl) => {
  if (!formEl) return
  await formEl.validate(async (valid) => {
    if (valid) {
      const response = await request.get(`/connect?host=${connectForm.value.host}&port=${connectForm.value.port}`);
      if (response.code) return
      ElNotification({
        title: 'Success',
        message: response.message,
        type: 'success',
      })
    }
  })
}

const onCancel = (formEl) => {
  if (!formEl) return
  formEl.resetFields()
}

const onSubmit = async (formEl) => {
  if (!formEl) return
  await formEl.validate(async (valid) => {
    if (valid) {
      const response = await request.post("/nodes", {
        name: connectForm.value.name,
        host: connectForm.value.host,
        port: connectForm.value.port,
      });
      if (response.code) return
      ElNotification({
        title: 'Success',
        message: response.message,
        type: 'success',
      })
      await getConnects()
    }
  })
}
</script>

<template>
  <div class="common-layout">
    <el-container class="h-screen">
      <el-header height="100px" class="border-gray-400 border-b-2">
        <div class="w-1/5 m-auto">
          <el-image :src="imgLogo" fit="contain" />
        </div>
      </el-header>
      <el-main>
        <el-row class="h-full" :gutter="20">
          <el-col :span="addConnectSpan">
            <el-button type="primary" @click="isAddConnect = !isAddConnect">添加节点</el-button>

            <el-table :data="connectData"  size="large" class="mt-3">
              <el-table-column prop="name" label="名称" />
              <el-table-column prop="host" label="地址" />
              <el-table-column prop="port" label="端口" />
              <el-table-column prop="createAt" label="创建时间" />
              <el-table-column label="操作">
                <template #default="scope">
                  <el-button type="success" :icon="Link" circle @click="onAuth(scope.row)"/>
                  <el-button type="primary" :icon="Edit" circle disabled/>
                  <el-button type="danger" :icon="Delete" circle disabled/>
                </template>
              </el-table-column>
            </el-table>
          </el-col>
          <el-col :span="8" class="border-blue-300 border-2 rounded" v-if="isAddConnect">
            <div class="flex justify-end my-3">
              <el-icon @click="isAddConnect = !isAddConnect"><Close /></el-icon>
            </div>
            <el-form ref="connectFormRef" :model="connectForm" :rules="connectFormRules" label-width="auto" label-position="top" size="large">
              <el-form-item label="名称" prop="name">
                <el-input v-model="connectForm.name" />
              </el-form-item>
              <el-form-item label="地址" prop="host">
                <el-input v-model="connectForm.host" />
              </el-form-item>
              <el-form-item label="端口" prop="port">
                <el-input v-model="connectForm.port" type="number" />
              </el-form-item>
              <el-form-item>
                <el-button text @click="onConnect(connectFormRef)">测试</el-button>
                <el-button @click="onCancel(connectFormRef)">取消</el-button>
                <el-button type="primary" @click="onSubmit(connectFormRef)">添加</el-button>
              </el-form-item>
            </el-form>
          </el-col>
        </el-row>
      </el-main>
    </el-container>
  </div>
</template>

<style scoped>

</style>