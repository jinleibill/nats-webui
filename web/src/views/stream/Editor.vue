<script setup>
import {ref} from "vue";
import request from "@/request.js";
import {ElNotification} from "element-plus";

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
  subjects: "",
})
const editorFormRules = ref({
  name: [
    {required: true, message: "名称必填", trigger: "blur"},
  ],
  subjects: [
    {required: true, message: "主题必填", trigger: "blur"},
  ],
})

const onInit = () => {
  if (Object.keys(item.value).length) {
    title.value = "编辑 Stream"
    submit.value = onEdit
    submitText.value = "修改"
    editorForm.value.name = item.value.name
    editorForm.value.subjects = item.value.subjects.join(",")
    isDisable.value = true
  } else {
    title.value = "添加 Stream"
    submit.value = onAdd
    submitText.value = "添加"
    editorForm.value.name = ""
    editorForm.value.subjects = ""
    isDisable.value = false
  }
}

const onAdd = async (formEl) => {
  if (!formEl) return
  await formEl.validate(async (valid) => {
    if (valid) {
      const response = await request.post("/streams", {
        name: editorForm.value.name,
        subjects: editorForm.value.subjects.split(","),
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
      const response = await request.put(`/streams/${editorForm.value.name}`, {
        subjects: editorForm.value.subjects.split(","),
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
    editorForm.value.subjects = item.value.subjects.join(",")
  }
}

</script>

<template>
  <div>
    <el-dialog v-model="dialogVisible" :title="title" width="600" @open="onInit">
      <el-form ref="editorFormRef" :model="editorForm" :rules="editorFormRules" label-width="auto" size="large">
        <el-form-item label="名称" prop="name">
          <el-input v-model="editorForm.name" :disabled="isDisable"/>
        </el-form-item>
        <el-form-item label="主题" prop="subjects">
          <el-input v-model="editorForm.subjects" placeholder="多个主题逗号分隔"/>
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
  </div>
</template>

<style scoped>

</style>