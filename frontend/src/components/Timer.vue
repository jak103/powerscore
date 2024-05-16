<script setup lang="ts">
import { reactive } from 'vue';
const state = reactive({time: 5.04}) //TODO: where does the time come from?
</script>

<template>
  <div class="timer">
    {{ state.time.toLocaleString(undefined, { minimumFractionDigits: 2 }) }}
  </div>
</template>

<style scoped>
</style>
