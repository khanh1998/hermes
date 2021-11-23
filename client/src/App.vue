<template>
  <reload-prompt />
  <router-view></router-view>
</template>
<script setup lang="ts">
import { defineComponent } from "@vue/runtime-core";
import ReloadPrompt from "./atoms/ReloadPrompt.vue";
</script>