<script setup lang="ts">
import TeamScore from './TeamScore.vue';
import GameTime from './GameTime.vue';
</script>

<template>
  <div class="primary-info">
    <TeamScore team-name="lakeville" :score=7 /> <!--TODO: Make home team name and score reactive elements-->
    <GameTime/>
    <TeamScore team-name="eagan" :score=2 /> <!--TODO: Make away team name and score reactive elements-->
  </div>
</template>

<style scoped>
</style>