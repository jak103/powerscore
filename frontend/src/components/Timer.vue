<script setup lang="ts">
import { reactive } from 'vue';
const state = reactive({time: 5.04});
</script>

<template>
  <div class="timer">
    {{ state.time.toLocaleString(undefined, { minimumFractionDigits: 2 }) }}
  </div>
</template>

<style scoped>
</style>
