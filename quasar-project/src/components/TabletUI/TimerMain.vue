<script setup lang="ts">
import {ref} from 'vue';

const period = ref(1)
// TODO: Create a method to poll the rink every so often to synchronize current time.

</script>

<template>
  <div class="timer-container">
    <div class="grid grid-cols-1">
        <div class="col-span-1 my-1">20:00</div>
    </div>
    <div>Period <span class="font-semibold">{{period}}</span></div>
  </div>
</template>

<style scoped>
.timer-container {
  background-color: #f5f5f5;
  padding: 1rem 1rem 1rem 1rem;
  margin: 1rem 1rem 1rem 1rem;
  border-radius: 18px;
  font-size: 36px;
  text-align: center;
  width: 50%;
}
</style>
