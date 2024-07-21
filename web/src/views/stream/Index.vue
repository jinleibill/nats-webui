<script setup>
import {Delete, Edit, Search} from "@element-plus/icons-vue";
import {onMounted, ref} from "vue";
import request from "@/request.js";
import utils from "@/utils.js";
import Editor from "@/views/stream/Editor.vue";
import {ElMessage, ElMessageBox, ElNotification} from "element-plus";
import {useRouter} from "vue-router";

const router = useRouter()

const editorDialogVisible = ref(false)
const item = ref({})

const items = ref([])

onMounted(()=>{
  getItems()
})

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

const onAdd = () => {
  item.value = {}
  editorDialogVisible.value = true
}

const onShow = (row) => {
  router.push(`/stream/${row.name}`)
}

const onEdit = (row) => {
  item.value = row
  editorDialogVisible.value = true
}

const onDel = async (row) => {
  ElMessageBox.confirm(
      `确认删除 ${row.name} Stream`,
      '注意',
      {
        confirmButtonText: '确认',
        cancelButtonText: '取消',
        type: 'warning',
      }
  ).then(async () => {
      const response = await request.delete(`/streams/${row.name}`);
      if (response.code) return
      ElNotification({
        title: 'Success',
        message: "删除成功",
        type: 'success',
      })
      await getItems()
    }).catch(() => {
      ElMessage({
        type: 'info',
        message: 'Delete canceled',
      })
    })
}

const onChange = () => {
  getItems()
}

</script>

<template>
  <div>
    <div>
      Stream 管理
    </div>
    <el-divider />

    <el-button type="primary" @click="onAdd">添加</el-button>

    <el-table :data="items"  size="large" class="mt-3">
      <el-table-column prop="name" label="名称" width="150px"/>
      <el-table-column prop="subjects" label="主题" />
      <el-table-column prop="numReplicas" label="副本" width="100px"/>
      <el-table-column prop="messages" label="消息" width="150px"/>
      <el-table-column prop="bytes" label="字节" />
      <el-table-column prop="consumerCount" label="消费者" width="100px"/>
      <el-table-column prop="createAt" label="创建时间" />
      <el-table-column label="操作">
        <template #default="scope">
          <el-button type="info" :icon="Search" circle @click="onShow(scope.row)"/>
          <el-button type="primary" :icon="Edit" circle @click="onEdit(scope.row)"/>
          <el-button type="danger" :icon="Delete" circle @click="onDel(scope.row)"/>
        </template>
      </el-table-column>
    </el-table>

    <Editor v-model:dialogVisible="editorDialogVisible" v-model:item="item" @change="onChange"></Editor>
  </div>
</template>

<style scoped>

</style>