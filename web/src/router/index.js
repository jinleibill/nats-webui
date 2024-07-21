import { createRouter, createWebHistory } from 'vue-router'
import Login from "@/views/Login.vue";
import StreamIndex from "@/views/stream/Index.vue";
import Layout from "@/views/Layout.vue";
import StreamShow from "@/views/stream/Show.vue";
import ConsumerIndex from "@/views/consumer/Index.vue";
import ConsumerShow from "@/views/consumer/Show.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: Layout,
      redirect: "/stream",
      children: [
        {
          path: '/stream',
          name: 'stream',
          component: StreamIndex
        },
        {
          path: '/stream/:stream',
          name: 'streamDetail',
          component: StreamShow
        },
        {
          path: '/consumer',
          name: 'consumer',
          component: ConsumerIndex
        },
        {
          path: '/consumer/:consumer',
          name: 'consumerDetail',
          component: ConsumerShow
        },
      ]
    },
    {
      path: '/login',
      name: 'login',
      component: Login
    },
  ]
})

export default router
