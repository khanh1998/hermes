<template>
  <div class="flex flex-col border border-green-400 rounded p-2">
    <h1 class="mb-4">Search result for: {{ searchString }}</h1>
    <div v-if="!loading">
      <div v-if="!result">No records for {{ searchString }}</div>
      <div v-for="mess in result" :key="mess.time" class="py-1">
        <div v-html="mess.message"></div>
      </div>
    </div>
    <div v-if="loading">Searching...</div>
    <router-link to="/" class="mt-4">Back to chat box</router-link>
  </div>
</template>
<script lang="ts">
import { defineComponent } from "vue";
import { Message, MessageState, SearchResult } from "../store/message";

export default defineComponent({
  name: "SearchResult",
  props: {
    searchString: {
      type: String,
      default: "",
    },
  },
  emits: [""],
  watch: {
    searchString(newVal: string) {
      console.log("hrer");

      if (newVal) {
        this.$store.dispatch("message/search", {
          clanId: 1,
          channelId: 1,
          searchString: this.searchString,
        });
      }
    },
  },
  computed: {
    messageState(): MessageState {
      return this.$store.state.message;
    },
    loading(): boolean {
      return this.messageState.loading;
    },
    result(): Array<Message> | null {
      const latest = this.messageState.searches.at(0);
      if (latest && latest.searchString === this.searchString) {
        return latest.results.map((item: Message) => {
          item.message = item.message.replaceAll(
            this.searchString,
            `<strong> ${this.searchString} </strong>`
          );
          item.message = `<p>${item.message}</p>`;
          return item;
        });
      }
      return null;
    },
  },
  methods: {},
  created() {
    this.$store.dispatch("message/search", {
      clanId: 1,
      channelId: 1,
      searchString: this.searchString,
    });
  },
});
</script>
