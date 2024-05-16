<script setup lang="ts">
import { reactive } from 'vue';
import Penalties from './Penalties.vue';
import ShotsOnGoal from './ShotsOnGoal.vue';

const state = reactive({
  homePenalties: [
    {
      player: 20,
      time: "1:42"
    },
    {
      player: 7,
      time: "1:42"
    },
  ],
  awayPenalties: [
    {
      player: 11,
      time: "1:47"
    },
    {
      player: 62,
      time: ":55"
    },
  ],
  homeShotsOnGoal: 30,
  awayShotsOnGoal: 24
})
//TODO: where do penalties and shots on goal come from?
</script>

<template>
  <div class="secondary-info">
    <Penalties team="home" :penalties=state.homePenalties />
    <ShotsOnGoal :home=state.homeShotsOnGoal :away=state.awayShotsOnGoal />
    <Penalties team="away" :penalties=state.awayPenalties />
  </div>
</template>

<style scoped>
</style>
