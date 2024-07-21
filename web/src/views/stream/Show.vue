<script setup>
import {useRoute, useRouter} from "vue-router";
import {onMounted, ref} from "vue";
import request from "@/request.js";
import {Warning} from "@element-plus/icons-vue";

const route = useRoute()
const router = useRouter()
const stream = route.params.stream

const column = ref(3)

const item = ref({})

const isLoaded = ref(false)

onMounted(()=>{
  getItem()
})

const goBack = () => {
  router.go(-1)
}

const getItem = async () => {
  const response = await request.get(`/streams/${stream}`);
  if (response.code) return
  item.value = response.data
  isLoaded.value = true
}

</script>

<template>
  <div>
    <el-page-header @back="goBack">
      <template #content>
        <span class="text-large font-600 mr-3"> {{ stream }} 详情 </span>
      </template>
    </el-page-header>

    <el-divider />

    <div v-if="isLoaded">
      <el-descriptions size="large" :column="column" border title="配置" class="mt-5">
        <el-descriptions-item width="200px">
          <template #label>
            name
            <el-popover
                placement="right"
                trigger="click"
                content="名称"
            >
              <template #reference>
                <el-icon><Warning /></el-icon>
              </template>
            </el-popover>
          </template>
          {{item.config.name}}
        </el-descriptions-item>
        <el-descriptions-item width="200px">
          <template #label>
            subjects
            <el-popover
                placement="right"
                trigger="click"
                content="stream 正在侦听的主题的列表, 支持通配符"
            >
              <template #reference>
                <el-icon><Warning /></el-icon>
              </template>
            </el-popover>
          </template>
          {{item.config.subjects}}
        </el-descriptions-item>
        <el-descriptions-item width="200px">
          <template #label>
            retention
            <el-popover
                placement="right"
                trigger="click"
                content="stream 消息保留策略"
            >
              <template #reference>
                <el-icon><Warning /></el-icon>
              </template>
            </el-popover>
          </template>
          {{item.config.retention}}
        </el-descriptions-item>
        <el-descriptions-item width="200px">
          <template #label>
            max_consumers
            <el-popover
                placement="right"
                trigger="click"
                content="允许的最大使用者数；-1 不限制"
            >
              <template #reference>
                <el-icon><Warning /></el-icon>
              </template>
            </el-popover>
          </template>
          {{item.config.max_consumers}}
        </el-descriptions-item>
        <el-descriptions-item width="200px">
          <template #label>
            max_msgs
            <el-popover
                placement="right"
                trigger="click"
                content="允许的最大消息数；-1 不限制"
            >
              <template #reference>
                <el-icon><Warning /></el-icon>
              </template>
            </el-popover>
          </template>
          {{item.config.max_msgs}}
        </el-descriptions-item>
        <el-descriptions-item width="200px">
          <template #label>
            max_bytes
            <el-popover
                placement="right"
                trigger="click"
                content="允许的最大字节数；-1 不限制"
            >
              <template #reference>
                <el-icon><Warning /></el-icon>
              </template>
            </el-popover>
          </template>
          {{item.config.max_bytes}}
        </el-descriptions-item>
        <el-descriptions-item width="200px">
          <template #label>
            discard
            <el-popover
                placement="right"
                trigger="click"
                content="当 stream 达到其消息数或总字节数限制时处理消息的策略"
            >
              <template #reference>
                <el-icon><Warning /></el-icon>
              </template>
            </el-popover>
          </template>
          {{item.config.discard}}
        </el-descriptions-item>
        <el-descriptions-item width="200px">
          <template #label>
            max_age
            <el-popover
                placement="right"
                trigger="click"
                content="stream 将保留的消息的最大期限"
            >
              <template #reference>
                <el-icon><Warning /></el-icon>
              </template>
            </el-popover>
          </template>
          {{item.config.max_age}}
        </el-descriptions-item>
        <el-descriptions-item width="200px">
          <template #label>
            max_msgs_per_subject
            <el-popover
                placement="right"
                trigger="click"
                content="stream 将保留的每个主题的最大消息数"
            >
              <template #reference>
                <el-icon><Warning /></el-icon>
              </template>
            </el-popover>
          </template>
          {{item.config.max_msgs_per_subject}}
        </el-descriptions-item>
        <el-descriptions-item width="200px">
          <template #label>
            max_msg_size
            <el-popover
                placement="right"
                trigger="click"
                content="stream 中任何单个消息的最大大小"
            >
              <template #reference>
                <el-icon><Warning /></el-icon>
              </template>
            </el-popover>
          </template>
          {{item.config.max_msg_size}}
        </el-descriptions-item>
        <el-descriptions-item width="200px">
          <template #label>
            storage
            <el-popover
                placement="right"
                trigger="click"
                content="stream 的存储后端的类型"
            >
              <template #reference>
                <el-icon><Warning /></el-icon>
              </template>
            </el-popover>
          </template>
          {{item.config.storage}}
        </el-descriptions-item>
        <el-descriptions-item width="200px">
          <template #label>
            num_replicas
            <el-popover
                placement="right"
                trigger="click"
                content="stream 副本数"
            >
              <template #reference>
                <el-icon><Warning /></el-icon>
              </template>
            </el-popover>
          </template>
          {{item.config.num_replicas}}
        </el-descriptions-item>
        <el-descriptions-item width="200px">
          <template #label>
            duplicate_window
            <el-popover
                placement="right"
                trigger="click"
                content="用于跟踪重复消息的范围"
            >
              <template #reference>
                <el-icon><Warning /></el-icon>
              </template>
            </el-popover>
          </template>
          {{item.config.duplicate_window}}
        </el-descriptions-item>
        <el-descriptions-item width="200px">
          <template #label>
            compression
            <el-popover
                placement="right"
                trigger="click"
                content="指定的信息存储压缩算法"
            >
              <template #reference>
                <el-icon><Warning /></el-icon>
              </template>
            </el-popover>
          </template>
          {{item.config.compression}}
        </el-descriptions-item>
        <el-descriptions-item width="200px">
          <template #label>
            allow_direct
            <el-popover
                placement="right"
                trigger="click"
                content="允许使用直接获取API直接访问单个消息"
            >
              <template #reference>
                <el-icon><Warning /></el-icon>
              </template>
            </el-popover>
          </template>
          {{item.config.allow_direct}}
        </el-descriptions-item>
        <el-descriptions-item width="200px">
          <template #label>
            created
            <el-popover
                placement="right"
                trigger="click"
                content="创建时间"
            >
              <template #reference>
                <el-icon><Warning /></el-icon>
              </template>
            </el-popover>
          </template>
          {{item.created}}
        </el-descriptions-item>
      </el-descriptions>
      <el-descriptions size="large" :column="column" border title="状态" class="mt-5">
        <el-descriptions-item width="200px">
          <template #label>
            messages
            <el-popover
                placement="right"
                trigger="click"
                content="存储在流中的消息数"
            >
              <template #reference>
                <el-icon><Warning /></el-icon>
              </template>
            </el-popover>
          </template>
          {{item.state.messages}}
        </el-descriptions-item>
        <el-descriptions-item width="200px">
          <template #label>
            bytes
            <el-popover
                placement="right"
                trigger="click"
                content="存储在流中的字节数"
            >
              <template #reference>
                <el-icon><Warning /></el-icon>
              </template>
            </el-popover>
          </template>
          {{item.state.bytes}}
        </el-descriptions-item>
        <el-descriptions-item width="200px">
          <template #label>
            first_seq
            <el-popover
                placement="right"
                trigger="click"
                content="流中第一个消息的序列号"
            >
              <template #reference>
                <el-icon><Warning /></el-icon>
              </template>
            </el-popover>
          </template>
          {{item.state.first_seq}}
        </el-descriptions-item>
        <el-descriptions-item width="200px">
          <template #label>
            first_ts
            <el-popover
                placement="right"
                trigger="click"
                content="流中第一条消息的时间戳"
            >
              <template #reference>
                <el-icon><Warning /></el-icon>
              </template>
            </el-popover>
          </template>
          {{item.state.first_ts}}
        </el-descriptions-item>
        <el-descriptions-item width="200px">
          <template #label>
            last_seq
            <el-popover
                placement="right"
                trigger="click"
                content="流中最后一个消息的序列号"
            >
              <template #reference>
                <el-icon><Warning /></el-icon>
              </template>
            </el-popover>
          </template>
          {{item.state.last_seq}}
        </el-descriptions-item>
        <el-descriptions-item width="200px">
          <template #label>
            last_ts
            <el-popover
                placement="right"
                trigger="click"
                content="流中最后一条消息的时间戳。"
            >
              <template #reference>
                <el-icon><Warning /></el-icon>
              </template>
            </el-popover>
          </template>
          {{item.state.last_ts}}
        </el-descriptions-item>
        <el-descriptions-item width="200px">
          <template #label>
            consumer_count
            <el-popover
                placement="right"
                trigger="click"
                content="流上的消费者数量。"
            >
              <template #reference>
                <el-icon><Warning /></el-icon>
              </template>
            </el-popover>
          </template>
          {{item.state.consumer_count}}
        </el-descriptions-item>
        <el-descriptions-item width="200px">
          <template #label>
            num_deleted
            <el-popover
                placement="right"
                trigger="click"
                content="从流中删除的消息数。"
            >
              <template #reference>
                <el-icon><Warning /></el-icon>
              </template>
            </el-popover>
          </template>
          {{item.state.num_deleted}}
        </el-descriptions-item>
        <el-descriptions-item width="200px">
          <template #label>
            num_subjects
            <el-popover
                placement="right"
                trigger="click"
                content="流已接收到其上的消息的唯一主题数。"
            >
              <template #reference>
                <el-icon><Warning /></el-icon>
              </template>
            </el-popover>
          </template>
          {{item.state.num_subjects}}
        </el-descriptions-item>
      </el-descriptions>
    </div>

  </div>
</template>

<style scoped>

</style>