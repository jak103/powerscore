<script setup lang="ts">
import TeamScore from './TeamScore.vue';
import GameTime from './GameTime.vue';
</script>

<template>
  <TeamScore team="lakeville" score=7 /> <!--TODO: Make home team name and score reactive elements-->
  <GameTime/>
  <TeamScore team="eagan" score=2 /> <!--TODO: Make away team name and score reactive elements-->
</template>

<style scoped>
</style>