<script setup>
import {Delete, Edit, Search} from "@element-plus/icons-vue";
import {onMounted, ref} from "vue";
import request from "@/request.js";
import utils from "@/utils.js";
import Editor from "@/views/consumer/Editor.vue";
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
  const response = await request.get("/consumers");
  if (response.code) return
  let tmpItems = []
  response.data.forEach(element => {
    tmpItems.push({
      name: element.name,
      durableName: element.durable_name,
      streamName: element.stream_name,
      numAckPending: element.num_ack_pending,
      numRedelivered: element.num_redelivered,
      numWaiting: element.num_waiting,
      numPending: element.num_pending,
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
  router.push({
    path: `/consumer/${row.name}`,
    query: {streamName: `${row.streamName}`},
  })
}

const onEdit = (row) => {
  item.value = row
  editorDialogVisible.value = true
}

const onDel = async (row) => {
  ElMessageBox.confirm(
      `确认删除 ${row.name} Consumer`,
      '注意',
      {
        confirmButtonText: '确认',
        cancelButtonText: '取消',
        type: 'warning',
      }
  ).then(async () => {
      const response = await request.delete(`/consumers/${row.name}?stream_name=${row.streamName}`);
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
      Consumer 管理
    </div>
    <el-divider />

    <el-button type="primary" @click="onAdd">添加</el-button>

    <el-table :data="items"  size="large" class="mt-3">
      <el-table-column prop="name" label="名称"/>
      <el-table-column prop="durableName" label="存储名"/>
      <el-table-column prop="streamName" label="流名"/>
      <el-table-column prop="numAckPending" label="待回复" width="80"/>
      <el-table-column prop="numRedelivered" label="已重新发送" width="130"/>
      <el-table-column prop="numWaiting" label="等待中" width="100"/>
      <el-table-column prop="numPending" label="待处理" width="100"/>
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