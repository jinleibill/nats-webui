<script setup>
import {useRoute, useRouter} from "vue-router";
import {onMounted, ref} from "vue";
import request from "@/request.js";
import {Warning} from "@element-plus/icons-vue";

const route = useRoute()
const router = useRouter()
const consumer = route.params.consumer
const streamName = route.query.streamName

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
  const response = await request.get(`/consumers/${consumer}?stream_name=${streamName}`);
  if (response.code) return
  item.value = response.data
  isLoaded.value = true
}

</script>

<template>
  <div>
    <el-page-header @back="goBack">
      <template #content>
        <span class="text-large font-600 mr-3"> {{ consumer }} 详情 </span>
      </template>
    </el-page-header>

    <el-divider />

    <div v-if="isLoaded">
      <el-descriptions size="large" :column="column" border title="基础信息" class="mt-5">
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
          {{item.info.name}}
        </el-descriptions-item>`
        <el-descriptions-item width="200px">
          <template #label>
            stream_name
            <el-popover
                placement="right"
                trigger="click"
                content="流名称"
            >
              <template #reference>
                <el-icon><Warning /></el-icon>
              </template>
            </el-popover>
          </template>
          {{item.info.stream_name}}
        </el-descriptions-item>`
        <el-descriptions-item width="200px">
          <template #label>
            num_ack_pending
            <el-popover
                placement="right"
                trigger="click"
                content="the number of messages that have been delivered but not yet acknowledged."
            >
              <template #reference>
                <el-icon><Warning /></el-icon>
              </template>
            </el-popover>
          </template>
          {{item.info.num_ack_pending}}
        </el-descriptions-item>`
        <el-descriptions-item width="200px">
          <template #label>
            num_pending
            <el-popover
                placement="right"
                trigger="click"
                content="the number of messages that match the consumer's filter, but have not been delivered yet."
            >
              <template #reference>
                <el-icon><Warning /></el-icon>
              </template>
            </el-popover>
          </template>
          {{item.info.num_pending}}
        </el-descriptions-item>`
        <el-descriptions-item width="200px">
          <template #label>
            num_redelivered
            <el-popover
                placement="right"
                trigger="click"
                content="counts the number of messages that have been redelivered and not yet acknowledged"
            >
              <template #reference>
                <el-icon><Warning /></el-icon>
              </template>
            </el-popover>
          </template>
          {{item.info.num_redelivered}}
        </el-descriptions-item>`
        <el-descriptions-item width="200px">
          <template #label>
            num_waiting
            <el-popover
                placement="right"
                trigger="click"
                content="the count of active pull requests. It is only relevant for pull-based consumers."
            >
              <template #reference>
                <el-icon><Warning /></el-icon>
              </template>
            </el-popover>
          </template>
          {{item.info.num_waiting}}
        </el-descriptions-item>`
        <el-descriptions-item width="200px">
          <template #label>
            num_waiting
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
          {{item.info.created}}
        </el-descriptions-item>`
      </el-descriptions>
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
          {{item.info.config.name}}
        </el-descriptions-item>`
        <el-descriptions-item width="200px">
          <template #label>
            ack_policy
            <el-popover
                placement="right"
                trigger="click"
                content="defines the acknowledgement policy for the consumer"
            >
              <template #reference>
                <el-icon><Warning /></el-icon>
              </template>
            </el-popover>
          </template>
          {{item.info.config.ack_policy}}
        </el-descriptions-item>`
        <el-descriptions-item width="200px">
          <template #label>
            ack_wait
            <el-popover
                placement="right"
                trigger="click"
                content="defines how long the server will wait for an acknowledgement before resending a message"
            >
              <template #reference>
                <el-icon><Warning /></el-icon>
              </template>
            </el-popover>
          </template>
          {{item.info.config.ack_wait}}
        </el-descriptions-item>`
        <el-descriptions-item width="200px">
          <template #label>
            deliver_policy
            <el-popover
                placement="right"
                trigger="click"
                content="defines from which point to start delivering messages from the stream."
            >
              <template #reference>
                <el-icon><Warning /></el-icon>
              </template>
            </el-popover>
          </template>
          {{item.info.config.deliver_policy}}
        </el-descriptions-item>`
        <el-descriptions-item width="200px">
          <template #label>
            durable_name
            <el-popover
                placement="right"
                trigger="click"
                content="an optional durable name for the consumer."
            >
              <template #reference>
                <el-icon><Warning /></el-icon>
              </template>
            </el-popover>
          </template>
          {{item.info.config.durable_name}}
        </el-descriptions-item>`
        <el-descriptions-item width="200px">
          <template #label>
            filter_subjects
            <el-popover
                placement="right"
                trigger="click"
                content="can be used to filter messages delivered from the stream"
            >
              <template #reference>
                <el-icon><Warning /></el-icon>
              </template>
            </el-popover>
          </template>
          {{item.info.config.filter_subjects}}
        </el-descriptions-item>`
        <el-descriptions-item width="200px">
          <template #label>
            max_ack_pending
            <el-popover
                placement="right"
                trigger="click"
                content="a maximum number of outstanding unacknowledged messages"
            >
              <template #reference>
                <el-icon><Warning /></el-icon>
              </template>
            </el-popover>
          </template>
          {{item.info.config.max_ack_pending}}
        </el-descriptions-item>`
        <el-descriptions-item width="200px">
          <template #label>
            max_deliver
            <el-popover
                placement="right"
                trigger="click"
                content="defines the maximum number of delivery attempts for a message"
            >
              <template #reference>
                <el-icon><Warning /></el-icon>
              </template>
            </el-popover>
          </template>
          {{item.info.config.max_deliver}}
        </el-descriptions-item>`
        <el-descriptions-item width="200px">
          <template #label>
            max_waiting
            <el-popover
                placement="right"
                trigger="click"
                content="a maximum number of pull requests waiting to be fulfilled."
            >
              <template #reference>
                <el-icon><Warning /></el-icon>
              </template>
            </el-popover>
          </template>
          {{item.info.config.max_waiting}}
        </el-descriptions-item>`
        <el-descriptions-item width="200px">
          <template #label>
            num_replicas
            <el-popover
                placement="right"
                trigger="click"
                content="the number of replicas for the consumer's state"
            >
              <template #reference>
                <el-icon><Warning /></el-icon>
              </template>
            </el-popover>
          </template>
          {{item.info.config.num_replicas}}
        </el-descriptions-item>`
        <el-descriptions-item width="200px">
          <template #label>
            replay_policy
            <el-popover
                placement="right"
                trigger="click"
                content="defines the rate at which messages are sent to the consumer"
            >
              <template #reference>
                <el-icon><Warning /></el-icon>
              </template>
            </el-popover>
          </template>
          {{item.info.config.replay_policy}}
        </el-descriptions-item>`
      </el-descriptions>
      <el-descriptions size="large" :column="column" border title="ack floor" class="mt-5">
        <el-descriptions-item width="200px">
          <template #label>
            consumer_seq
            <el-popover
                placement="right"
                trigger="click"
                content="消费者序列号"
            >
              <template #reference>
                <el-icon><Warning /></el-icon>
              </template>
            </el-popover>
          </template>
          {{item.info.ack_floor.consumer_seq}}
        </el-descriptions-item>`
        <el-descriptions-item width="200px">
          <template #label>
            stream_seq
            <el-popover
                placement="right"
                trigger="click"
                content="流序列号"
            >
              <template #reference>
                <el-icon><Warning /></el-icon>
              </template>
            </el-popover>
          </template>
          {{item.info.ack_floor.stream_seq}}
        </el-descriptions-item>`
      </el-descriptions>
    </div>

  </div>
</template>

<style scoped>

</style>