<script setup>
import {ref} from "vue";
import request from "@/request.js";
import {ElNotification} from "element-plus";
import Selector from "@/views/consumer/Selector.vue";

const dialogVisible = defineModel("dialogVisible")
const item = defineModel("item")

const emit = defineEmits(["change"])

const title = ref("")
const submit = ref()
const submitText = ref("")
const isDisable = ref(false)

const editorFormRef = ref()
const editorForm = ref({
  name: "",
  streamName: "",
  durableName: "",
})
const editorFormRules = ref({
  name: [
    {required: true, message: "名称必填", trigger: "blur"},
  ],
})

const selectText = ref("请选择 stream")
const selectDialogVisible = ref(false)

const onInit = () => {
  if (Object.keys(item.value).length) {
    title.value = "编辑 Consumer"
    submit.value = onEdit
    submitText.value = "修改"
    editorForm.value.name = item.value.name
    editorForm.value.streamName = item.value.streamName
    editorForm.value.durableName = item.value.durableName
    isDisable.value = true
    selectText.value = item.value.streamName
  } else {
    title.value = "添加 Consumer"
    submit.value = onAdd
    submitText.value = "添加"
    editorForm.value.name = ""
    editorForm.value.streamName = ""
    editorForm.value.durableName = ""
    isDisable.value = false
    selectText.value = "请选择 stream"
  }
}

const onAdd = async (formEl) => {
  if (!formEl) return
  await formEl.validate(async (valid) => {
    if (valid) {
      if (selectText.value === "请选择 stream"){
        return
      }
      const response = await request.post("/consumers", {
        name: editorForm.value.name,
        stream_name: selectText.value,
        durable_name: editorForm.value.durableName,
      });
      if (response.code) return
      ElNotification({
        title: 'Success',
        message: response.message,
        type: 'success',
      })
      dialogVisible.value = false
      emit("change")
    }
  })
}

const onEdit = async(formEl) => {
  if (!formEl) return
  await formEl.validate(async (valid) => {
    if (valid) {
      const response = await request.put(`/consumers/${editorForm.value.name}`, {
        stream_name: selectText.value,
        durable_name: editorForm.value.durableName,
      });
      if (response.code) return
      ElNotification({
        title: 'Success',
        message: response.message,
        type: 'success',
      })
      dialogVisible.value = false
      emit("change")
    }
  })
}

const onReset = (formEl) => {
  formEl.resetFields()
  if (Object.keys(item.value).length) {
    editorForm.value.name = item.value.name
    editorForm.value.streamName = item.value.streamName
    editorForm.value.durableName = item.value.durableName
  }
}

const onSelect = () => {
  selectDialogVisible.value = true
}

const onChange = (select) => {
  selectText.value = select
  selectDialogVisible.value = false
}

</script>

<template>
  <div>
    <el-dialog v-model="dialogVisible" :title="title" width="600" @open="onInit">
      <el-form ref="editorFormRef" :model="editorForm" :rules="editorFormRules" label-width="auto" size="large">
        <el-form-item label="名称" prop="name">
          <el-input v-model="editorForm.name" :disabled="isDisable"/>
        </el-form-item>
        <el-form-item label="流名" prop="streamName">
          <el-button type="primary" @click="onSelect" :disabled="isDisable">{{ selectText }}</el-button>
        </el-form-item>
        <el-form-item label="存储名" prop="durableName">
          <el-input v-model="editorForm.durableName" :disabled="isDisable"/>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="onReset(editorFormRef)">重置</el-button>
          <el-button type="primary" @click="submit(editorFormRef)">
            {{submitText}}
          </el-button>
        </div>
      </template>
    </el-dialog>

    <Selector v-model:dialogVisible="selectDialogVisible" @change="onChange"></Selector>
  </div>
</template>

<style scoped>

</style>