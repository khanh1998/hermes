<template>
  <router-view></router-view>
</template>
<script lang="ts">
import { defineComponent } from 'vue'
export default defineComponent({
  name: 'App',
  data: function () {
    return {
      ws: {} as WebSocket,
    };
  },
  created() {
    const token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImtoYW5oIiwiaWQiOjEsInR5cGUiOiIxIiwiaWF0IjoxNjM0OTc0NzY2LCJleHAiOjE2MzQ5NzUwNjZ9.P-oXlKfgd28v6aQh8K25B7Lvb_5vNegjJkKhcsPjNPI"
    const host = import.meta.env.VITE_SOCKET_HOST;
    const ws = new WebSocket(`${host}?token=${token}`)
    this.ws = ws;
    let message = {
      name: "khanh",
      message: "ping ping ping ping",
      time: Date.now(),
    }
    ws.onopen = (ev: Event) => {
      console.log({ open: ev })
      ws.send(JSON.stringify(message))
    }
    ws.onmessage = (ev: MessageEvent<any>) => {
      console.log({ message: ev})
      message = {
        name: "khanh",
        message: "ping ping ping ping",
        time: Date.now(),
      }
      setTimeout(() => ws.send(JSON.stringify(message)), 1000)
    }
  },
});
</script>
