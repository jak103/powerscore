<script setup lang="ts">
import {ref} from "vue";

const shotsOnGoal = ref(0)
</script>

<template>
  <div class="grid grid-cols-1 md:grid-cols-2 align-bottom my-1 md:my-2">
    <div class="col-span-1 font-normal">Shots on goal {{shotsOnGoal}}</div>
    <div class="col-span-1 flex justify-end items-center text-white">
      <button class="rounded-full bg-blue-700 px-4 py-2 ml-0.5 hover:bg-blue-800 hover:shadow-md" @click="shotsOnGoal--"> - </button>
      <button class="rounded-full bg-blue-800 px-4 py-2 ml-0.5 hover:bg-blue-900 hover:shadow-md" @click="shotsOnGoal++"> + </button>
    </div>
  </div>
</template>

<style scoped>

</style>