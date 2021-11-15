<template>
  <div
    class="
      flex flex-col
      border border-green-400
      text-green-900
      rounded
      p-2
      capitalize
			font-medium
    "
  >
    <p>{{ user?.fullname }}</p>
  </div>
</template>
<script lang="ts">
import { defineComponent } from "vue";
import { Channel } from "../store/clan";
import { UserData } from "../store/user";

export default defineComponent({
  name: "UserInfo",
  props: {
    user: {
      type: Object,
      default: () => ({}),
      required: true,
    },
  },
});
</script>
