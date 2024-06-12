<script setup lang="ts">
import TeamScore from './TeamScore.vue';
import GameTime from './GameTime.vue';
import { reactive } from 'vue';

const state = reactive({
  homeTeamName: "lakeville",
  homeTeamScore: 7,
  awayTeamName: "eagan",
  awayTeamScore: 2,
});

async function fetchScores() {
  try {
    const response = await fetch('http://localhost:8080/api/score');
    const scores = await response.json();
    state.homeTeamScore = scores.home;
    state.awayTeamScore = scores.away;
    
  } catch (error) {
    console.error('Error fetching scores:', error);
  }
}

fetchScores();

</script>

<template>
  <div class="primary-info">
    <TeamScore :team-name=state.homeTeamName :score=state.homeTeamScore />
    <GameTime/>
    <TeamScore :team-name=state.awayTeamName :score=state.awayTeamScore />
  </div>
</template>

<style scoped>
</style>