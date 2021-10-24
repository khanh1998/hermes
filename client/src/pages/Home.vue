<template>
  <chat-box class="m-2"/>
</template>
<script lang="ts">
import { defineComponent } from 'vue'
import ChatBox from '../components/ChatBox.vue';
export default defineComponent({
  components: { ChatBox },
  name: 'Home',
});
</script>