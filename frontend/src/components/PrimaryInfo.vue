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