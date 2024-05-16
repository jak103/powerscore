<script setup lang="ts">
const props = defineProps<{
  teamName: string,
  score: number
}>();
</script>

<template>
  <div class="team-score">
    <div class="team-name">{{ props.teamName }}</div>
    <div class="score">{{ props.score }}</div>
  </div>
</template>

<style scoped>
</style>