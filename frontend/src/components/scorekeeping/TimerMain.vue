<script setup lang="ts">

import Period from "@/components/scorekeeping/PeriodCounter.vue";

function setTimer() {
  alert("setTimer was called. Please implement me :)")
}

function pauseTimer() {
  alert("pauseTimer was called. Please implement me :)")
}

function startTimer() {
  alert("startTimer was called. Please implement me :)")
}

</script>

<template>
  <div class="bg-gray-50 p-2 m-2 rounded-md shadow-md text-lg md:text-6xl w-3/4 text-center">
    <div class="grid grid-cols-1">
        <div class="col-span-1 my-1">20:00</div>
        <div class="col-span-2 text-lg md:text-xl my-1 md:my-2 text-white">
          <button class="m-1 px-4 py-2 rounded-full bg-blue-800 ml-0.5 hover:bg-blue-900 hover:shadow-md" @click="setTimer()"> Set </button>
          <button class="m-1 px-4 py-2 rounded-full bg-red-700 ml-0.5 hover:bg-red-800 hover:shadow-md" @click="pauseTimer()"> Pause </button>
          <button class="m-1 px-4 py-2 rounded-full bg-green-800 ml-0.5 hover:bg-green-900 hover:shadow-md" @click="startTimer()"> Start </button>
        </div>
    </div>
    <Period/>
  </div>
</template>

<style scoped>

</style>