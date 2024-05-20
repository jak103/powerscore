<script setup lang="ts">
  import { reactive } from 'vue'
  import TeamCard from "@/components/scorekeeping/TeamCard.vue";
  import Timer from "@/components/scorekeeping/TimerMain.vue";
  import PenaltyTracker from "@/components/scorekeeping/PenaltyTracker.vue";

  const dummyData = reactive({
    teamOne: {
      name: "The Orange Slices",
      logo: "pathToLogoPlaceholder",
    },
    teamTwo: {
      name: "The Cherry Tomatoes",
      logo: "pathToLogoPlaceholder",
    },
  })
</script>

<template>
<!--  NOTE: just trying to get a basic structure here to split up the work a bit, we can make it prettier and move/fine tune things as we go -->
  <div id="score-keeper-root" class="grid grid-cols-1">
    <h1 class="p-2 text-lg">Powerscore live game workspace (work in progress)</h1>

    <!--    Teams     -->
    <div class="grid grid-cols-1 md:grid-cols-2 ">
      <TeamCard :team-name="dummyData.teamOne.name" :home-team="true" class="col-span-1"/>
      <TeamCard :team-name="dummyData.teamTwo.name" :home-team="false" class="col-span-1"/>
    </div>

    <!--    Penalty sections -->
    <div class="grid grid-cols-1 md:grid-cols-2">
        <PenaltyTracker :team="dummyData.teamOne.name"/>
        <PenaltyTracker :team="dummyData.teamTwo.name"/>
    </div>

    <!-- Timer (and period)  -->
    <div class="flex justify-center">
      <Timer/>
    </div>
  </div>
</template>

<style scoped>

</style>