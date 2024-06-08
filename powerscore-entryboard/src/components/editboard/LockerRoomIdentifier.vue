<script setup lang="ts">
const props = defineProps<{
  number: number
}>();
</script>

<template>
  <div class="locker-room-identifier">
    Locker Room #{{ props.number }}
  </div>
</template>

<style scoped>
</style>
