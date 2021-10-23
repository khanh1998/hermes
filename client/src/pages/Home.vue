<template>
	<h1>This is our home!</h1>
</template>
<script lang="ts">
import { defineComponent } from 'vue'
export default defineComponent({
  name: 'Home',
  data: function () {
    return {
      ws: {} as WebSocket,
    };
  },
  created() {
    const token = localStorage.getItem('ws_token');
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