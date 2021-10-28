<template>
  <div class="m-1 flex flex-row" :class="[justify]">
    <div class="rounded-full bg-green-900 w-8 h-8 text-center align-text mr-1">
      {{ username }}
    </div>
    <div class="flex flex-col items-end">
      <span class="border-green-600 bg-green-100 p-1 rounded">
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
        return Side.LEFT === val || Side.RIGHT === val
      },
    },
	},
  computed: {
    justify(): string {
      return this.side === Side.LEFT ? "justify-start" : "justify-end	";
    }
  }
});
</script>
