<script setup lang="ts">
import Penalties from './Penalties.vue';
import ShotsOnGoal from './ShotsOnGoal.vue';
</script>

<template>
  <Penalties team="home"/>
  <ShotsOnGoal/>
  <Penalties team="away"/>
</template>

<style scoped>
</style>
