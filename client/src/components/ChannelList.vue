<template>
  <div class="flex flex-col border border-green-400 rounded p-2">
    <span>Channel</span>
    <div
      v-for="chan in extendedChannels"
      :key="chan.id"
      class="my-1 p-1 rounded hover:bg-green-700 cursor-pointer"
      :class="[chan.background]"
      @click="changeChannel(chan.id)"
    >
      {{ chan.name }}
    </div>
  </div>
</template>
<script lang="ts">
import { defineComponent } from "vue";
import { Channel } from "../store/clan";
interface ExtendedChannel extends Channel {
  selected: boolean;
}
export default defineComponent({
  name: "ChannelList",
  props: {
    channels: {
      type: Array,
      default: () => [],
      required: true,
    },
    currentChannelId: Number,
  },
  emits: ["change-channel"],
  computed: {
    extendedChannels() {
      return this.channels.map((item: any) => {
        return {
          ...item,
          background: this.currentChannelId === item.id ? "bg-gray-400" : "",
        };
      });
    },
  },
  methods: {
    changeChannel(channelId: number) {
      this.$emit("change-channel", channelId);
    },
  },
});
</script>
