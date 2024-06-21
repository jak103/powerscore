<script setup lang="ts">
import {ref, onMounted} from "vue";
import axios from 'axios';

const props = defineProps<{
  homeTeam?: boolean;
}>();

onMounted(() => {
  fetchShots();
});

const shotsOnGoal = ref(0)

const incrementShots = async () => {
  shotsOnGoal.value++;
  await updateShots();
};

const decrementShots = async () => {
  if (shotsOnGoal.value > 0){
    shotsOnGoal.value--;
    await updateShots();
  }
};

const updateShots = async () => {
  console.log(props.homeTeam)
  try {
    if (props.homeTeam) {
      const response = await axios.post('http://localhost:8080/api/update-shots', {
      home: shotsOnGoal.value,
    });
    console.log('Shots updated successfully:', response.data);

    } else {
      const response = await axios.post('http://localhost:8080/api/update-shots', {
      away: shotsOnGoal.value,
    });
    console.log('Shots updated successfully:', response.data);
    }

  } catch (error) {
    console.error('Error updating Shots:', error);
  }
};

async function fetchShots() {
  try {
    const response = await fetch('http://localhost:8080/api/shots');
    const shots = await response.json();
    if (props.homeTeam) {
      shotsOnGoal.value = shots.home;
    } else {
      shotsOnGoal.value = shots.away;
    }
    
  } catch (error) {
    console.error('Error fetching scores:', error);
  }
}

</script>

<template>
  <div class="grid grid-cols-1 md:grid-cols-2 align-bottom my-1 md:my-2">
    <div class="col-span-1 font-normal">Shots on goal {{shotsOnGoal}}</div>
    <div class="col-span-1 flex justify-end items-center text-white">
      <button class="rounded-full bg-blue-700 px-4 py-2 ml-0.5 hover:bg-blue-800 hover:shadow-md" @click="decrementShots"> - </button>
      <button class="rounded-full bg-blue-800 px-4 py-2 ml-0.5 hover:bg-blue-900 hover:shadow-md" @click="incrementShots"> + </button>
    </div>
  </div>
</template>

<style scoped>

</style>