<template>
  <div class="w-full">
    <input
      type="text"
      placeholder="Search..."
      class="
        w-full
        rounded
        border border-light-blue-500
        p-1
        focus:border-blue-500
        ring-indigo-300
        focus:outline-none focus:ring
      "
      v-model="search"
      @keyup="pressEnterToSearch"
    />
  </div>
</template>
<script lang="ts">
import { defineComponent } from "vue";
const emits = {
  SearchFor: "search-for"
}
export default defineComponent({
  name: "SearchBox",
  data: () => ({
    search: "",
  }),
  props: {},
  emits: [emits.SearchFor],
  computed: {},
  methods: {
    pressEnterToSearch(ev: KeyboardEvent) {
      if (ev.key === "Enter") {
        this.emitSearchFor()
      }
    },
    emitSearchFor() {
      if (this.search.length > 0) {
        this.$emit(emits.SearchFor, this.search)
      }
    },
  },
});
</script>
