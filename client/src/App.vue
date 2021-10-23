<script setup lang="ts">
// This starter template is using Vue 3 <script setup> SFCs
// Check out https://v3.vuejs.org/api/sfc-script-setup.html#sfc-script-setup
import HelloWorld from './components/HelloWorld.vue'
</script>

<template>
  <img alt="Vue logo" src="./assets/logo.png" />
  <HelloWorld msg="Hello Vue 3 + TypeScript + Vite" />
</template>
<script lang="ts">
import { defineComponent } from 'vue'
export default defineComponent({
  name: 'App',
  components: {
    HelloWorld,
  },
  data: function () {
    return {
      ws: WebSocket,
    };
  },
  created() {
    const token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImtoYW5oIiwiaWQiOjEsInR5cGUiOiIxIiwiaWF0IjoxNjM0OTc0NzY2LCJleHAiOjE2MzQ5NzUwNjZ9.P-oXlKfgd28v6aQh8K25B7Lvb_5vNegjJkKhcsPjNPI"
    const ws = new WebSocket(`ws://localhost:8080?token=${token}`)
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
<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
