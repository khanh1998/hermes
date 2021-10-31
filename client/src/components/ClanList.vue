<template>
  <div class="flex flex-col border border-green-400 rounded p-2">
		<span>Clan</span>
    <select @change="handleChange">
      <option disabled value="">Please select one</option>
      <option v-for="clan in clans" :key="clan.id" v-bind:value="clan.id">
        {{ clan.name }}
      </option>
    </select>
  </div>
</template>
<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "ClanList",
  props: {
    clans: {
      type: Array,
      required: true,
      default: () => [],
    },
    selectedClanId: {
      type: Number,
    },
  },
  emits: ["change-clan"],
  computed: {},
  methods: {
    handleChange(ev: Event) {
      const clanId = Number((ev?.target as any).value);
			this.$emit("change-clan", clanId);
    },
  },
});
</script>
