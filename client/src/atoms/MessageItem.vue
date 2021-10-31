<template>
  <div class="m-1 flex flex-row" :class="[justify]">
    <div
      class="rounded-full bg-green-200 border border-green-400 w-8 h-8 text-center font-bold text-white align-text mr-1"
      :class="[display]"
    >
      {{ username }}
    </div>
    <div class="flex flex-col items-end">
      <span
        class="border border-green-600 p-1 rounded"
        :class="[messageBackground]"
      >
        {{ message }}
      </span>
      <span class="text-xs">{{ timeStr }}</span>
    </div>
  </div>
</template>
<script lang="ts">
import { defineComponent } from "vue";
import { Side } from "../constants/Chatbox";
export default defineComponent({
  name: "LeftAlignMessage",
  props: {
    message: String,
    timeStr: String,
    username: String,
    side: {
      type: Number,
      validator(val) {
        return Side.LEFT === val || Side.RIGHT === val;
      },
    },
    showAvatar: Boolean,
    isSender: Boolean,
  },
  computed: {
    justify(): string {
      return this.side === Side.LEFT ? "justify-start" : "justify-end	";
    },
    display(): string {
      return this.showAvatar ? "" : "hidden";
    },
    messageBackground(): string {
      return this.isSender ? "bg-white" : "bg-green-100"; 
    }
  },
});
</script>
