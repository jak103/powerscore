<script setup lang="ts">
import ScoreCounter from "@/components/scorekeeping/ScoreCounter.vue";
import placeholderImage from "@/assets/img/hockeyClipart.jpg";
import ShotsOnGoalCounter from "@/components/scorekeeping/ShotsOnGoalCounter.vue"; // TODO: need to make this dynamically load team logos
import {defineProps } from 'vue';

defineProps<{
  teamName?: string
  homeTeam?: boolean
}>()

</script>

<template>
  <div class="bg-gray-50 p-2 m-2 rounded-md shadow-md text-lg">
    <div class="flex justify-between items-start mb-1 md:mb-2 font-semibold">
      {{teamName + " (" + (homeTeam ? "home)" : "away)")}}
      <img alt="team logo" class="h-10 w-10 md:h-12 md:w-12 rounded-full ring-2 ring-gray-300" :src="placeholderImage" />
    </div>
    <ScoreCounter :current-score="0" :homeTeam="homeTeam"/>
    <ShotsOnGoalCounter :homeTeam="homeTeam"/>
  </div>
</template>

<style scoped>

</style>