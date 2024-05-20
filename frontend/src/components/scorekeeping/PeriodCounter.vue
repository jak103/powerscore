<script setup lang="ts">

import {ref} from "vue";

const period = ref(1)
</script>

<template>
  <div class="my-1">
    <div class="font-normal text-center text-lg md:text-2xl flex justify-between items-center">
      <div>Period <span class="font-semibold">{{period}}</span></div>
      <div>
        <button class="rounded-full bg-blue-700 px-4 py-2 ml-0.5 hover:bg-blue-800 hover:shadow-md text-white" @click="period--"> - </button>
        <button class="rounded-full bg-blue-800 px-4 py-2 ml-0.5 hover:bg-blue-900 hover:shadow-md text-white" @click="period++"> + </button>
      </div>
    </div>
  </div>
</template>

<style scoped>

</style>