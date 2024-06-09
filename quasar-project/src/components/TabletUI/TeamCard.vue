<script setup lang="ts">
import placeholderImage from 'assets/img/hockeyClipart.jpg'; // TODO: need to make this dynamically load team logos
defineProps<{
  teamName?: string
  homeTeam?: boolean
}>()
</script>

<template>
  <div class="container">
    <div class="row items-center justify-between text-weight-medium">
      {{teamName + " (" + (homeTeam ? "home)" : "away)")}}
      <img alt="team logo" class="logoStyle" :src="placeholderImage" />
    </div>
    <div class="row items-center justify-between">
<!--      TODO: Make this dynamic and poll for updates. -->
      Current score is 0
    </div>
  </div>
</template>

<style scoped>
.logoStyle {
  height: 2.5rem;
  width: 2.5rem;
  border-radius: 100%;
}
.container {
  background-color: #f5f5f5;
  padding: 0.5rem 0.5rem 0.5rem 1rem;
  margin: 0 0.5rem 0 0.5rem;
  border-radius: 18px;
  font-size: 1.125rem;
  line-height: 1.75rem;
}
</style>
