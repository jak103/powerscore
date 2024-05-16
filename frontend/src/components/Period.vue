<script setup lang="ts">
const props = defineProps<{
  period: number,
}>();
</script>

<template>
  <div class="period">
    <div class="period-text">Period</div>
    <div class="period-number"> {{ props.period }}</div>
  </div>
</template>

<style scoped>
</style>
