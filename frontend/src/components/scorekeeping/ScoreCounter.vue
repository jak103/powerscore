<script setup lang="ts">
import {ref} from "vue";

defineProps<{
  currentScore?: number
}>()
// TODO: make this initialize to the loaded score value from the prop, add validation, maybe a confirmation before actually changing the value?
const score = ref(0)
</script>

<template>
  <div class="grid grid-cols-1 md:grid-cols-2 align-bottom ">
    <div class="col-span-1 font-normal">Current score is {{score}}</div>
    <div class="col-span-1 flex justify-end items-center text-white">
      <button class="rounded-full bg-blue-700 px-4 py-2 ml-0.5 hover:bg-blue-800 hover:shadow-md" @click="score--"> - </button>
      <button class="rounded-full bg-blue-800 px-4 py-2 ml-0.5 hover:bg-blue-900 hover:shadow-md" @click="score++"> + </button>
    </div>
  </div>
</template>

<style scoped>

</style>