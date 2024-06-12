<script setup lang="ts">
import { ref, defineProps, onMounted } from 'vue';
import axios from 'axios';

const props = defineProps<{
  homeTeam?: boolean;
}>();

onMounted(() => {
  fetchScores();
});

let score = ref(0);

const incrementScore = async () => {
  score.value++;
  await updateScore();
};

const decrementScore = async () => {
  if(score.value > 0){
    score.value--;
    await updateScore();
  }
};

const updateScore = async () => {
  console.log(props.homeTeam)
  try {
    if (props.homeTeam) {
      const response = await axios.post('http://localhost:8080/api/update-score', {
      home: score.value,
    });
    console.log('Score updated successfully:', response.data);

    } else {
      const response = await axios.post('http://localhost:8080/api/update-score', {
      away: score.value,
    });
    console.log('Score updated successfully:', response.data);
    }

  } catch (error) {
    console.error('Error updating score:', error);
  }
};

async function fetchScores() {
  try {
    const response = await fetch('http://localhost:8080/api/score');
    const scores = await response.json();
    if (props.homeTeam) {
      score.value = scores.home;
    } else {
      score.value = scores.away;
    }
    
  } catch (error) {
    console.error('Error fetching scores:', error);
  }
}
</script>

<template>
  <div class="grid grid-cols-1 md:grid-cols-2 align-bottom">
    <div class="col-span-1 font-normal">Current score is {{ score }}</div>
    <div class="col-span-1 flex justify-end items-center text-white">
      <button class="rounded-full bg-blue-700 px-4 py-2 ml-0.5 hover:bg-blue-800 hover:shadow-md" @click="decrementScore"> - </button>
      <button class="rounded-full bg-blue-800 px-4 py-2 ml-0.5 hover:bg-blue-900 hover:shadow-md" @click="incrementScore"> + </button>
    </div>
  </div>
</template>

<style scoped>

</style>
