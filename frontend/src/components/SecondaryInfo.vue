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

async function fetchShots() {
  try {
    const response = await fetch('http://localhost:8080/api/shots');
    const shots = await response.json();
    state.homeShotsOnGoal = shots.home;
    state.awayShotsOnGoal = shots.away;
    
  } catch (error) {
    console.error('Error fetching scores:', error);
  }
}

fetchShots()

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


