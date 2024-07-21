<script setup>
import {Coin, User} from "@element-plus/icons-vue";
import {onMounted} from "vue";
import router from "@/router/index.js";
import imgLogo from "@/assets/nats-logo.png";

onMounted(() => {
  isAuth()
})

const isAuth = () => {
  let auth = localStorage.getItem("isAuth")
  if (!auth) {
    router.push("/login")
  }
}
</script>

<template>
  <div class="common-layout">
    <el-container class="h-screen">
      <el-aside width="200px" class="border-gray-400 border-r-2">
        <div>
          <el-image :src="imgLogo" fit="contain" />
        </div>
        <el-menu
            router
            class="el-menu-vertical-demo"
            default-active="/stream"
        >
          <el-menu-item index="/stream">
            <el-icon><coin /></el-icon>
            <span>Stream 管理</span>
          </el-menu-item>
          <el-menu-item index="/consumer">
            <el-icon><user /></el-icon>
            <span>Consumer 管理</span>
          </el-menu-item>
        </el-menu>
      </el-aside>
      <el-main>
        <router-view></router-view>
      </el-main>
    </el-container>
  </div>
</template>

<style scoped>

</style>