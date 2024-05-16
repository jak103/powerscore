<script setup lang="ts">
import Timer from './Timer.vue';
import Period from './Period.vue';
import { reactive } from 'vue';

const state = reactive({period: 3});

</script>

<template>
  <div class="time">
    <Timer/>
    <Period :period=state.period />
  </div>
</template>

<style scoped>
</style>