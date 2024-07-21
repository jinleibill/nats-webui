<script setup>
import {ref} from "vue";
import request from "@/request.js";
import utils from "@/utils.js";

const dialogVisible = defineModel("dialogVisible")
const emit = defineEmits(["change"])

const items = ref([])
const radio = ref("")
const currentSelect = ref("")

const getItems = async() => {
  const response = await request.get("/streams");
  if (response.code) return
  let tmpItems = []
  response.data.forEach(element => {
    tmpItems.push({
      name: element.name,
      subjects: element.subjects,
      numReplicas: element.num_replicas,
      messages: element.messages,
      consumerCount: element.consumer_count,
      bytes: element.bytes,
      createAt: utils.formatData(element.create_at),
    })
  })
  items.value = tmpItems
}

const getCurrentRow = (item) => {
  currentSelect.value = item.name
}

const submit = () => {
  emit("change", currentSelect.value)
}

</script>

<template>
  <el-dialog v-model="dialogVisible" title="请选择 Stream" width="1000" @open="getItems">
    <el-table :data="items"  size="large" class="mt-3">
      <el-table-column label="选择" width="65">
        <template #default="scope">
          <el-radio :label="scope.row.name" v-model="radio" @change.native="getCurrentRow(scope.row)">&nbsp;</el-radio>
        </template>
      </el-table-column>
      <el-table-column prop="name" label="名称" width="150px"/>
      <el-table-column prop="subjects" label="主题" />
      <el-table-column prop="numReplicas" label="副本" width="100px"/>
      <el-table-column prop="messages" label="消息" width="150px"/>
      <el-table-column prop="bytes" label="字节" />
      <el-table-column prop="consumerCount" label="消费者" width="100px"/>
      <el-table-column prop="createAt" label="创建时间" />
    </el-table>

    <template #footer>
      <div class="dialog-footer">
        <el-button type="primary" @click="submit()">
          确定
        </el-button>
      </div>
    </template>
  </el-dialog>

</template>

<style scoped>

</style>