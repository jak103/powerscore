<script setup lang="ts">
const props = defineProps<{
  home: number,
  away: number
}>();
</script>

<template>
  <div class="shots-on-goal">
    <div class="shots">{{ home }}</div>
    <div class="shots">{{ away }}</div>
    <div class="shots-on-goal-text">Shots On Goal</div>
  </div>
</template>

<style scoped>
</style>
