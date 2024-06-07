<script setup lang="ts">
import { ref } from 'vue'

// TODO: Make options values dynamic on load, taking into account which team scored
// TODO: Create event listeners to pop up this form automatically when a team scores
// TODO: Hook up submit button to the backend API
const prompt = ref(false)
const goal = ref('')
const optionsGoal = [
  '11 - John',
  '15 - Jane',
  '21 - Eric',
  '46 - Randy',
]

const assistOne = ref('')
const optionsAssistOne = [
  '11 - John',
  '15 - Jane',
  '21 - Eric',
  '46 - Randy',
]

const assistTwo = ref('')
const optionsAssistTwo = [
  '11 - John',
  '15 - Jane',
  '21 - Eric',
  '46 - Randy',
]

</script>

<template>
  <div class="q-pa-md q-gutter-sm">
  <q-btn label="New Goal" color="primary" @click="prompt = true" />

  <q-dialog v-model="prompt" persistent>
    <q-card style="min-width: 350px">
      <q-card-section>
        <div class="text-h6">New Goal</div>
      </q-card-section>

      <q-card-section class="q-pt-none">
        <q-select v-model="goal" :options="optionsGoal" label="Goal" />
      </q-card-section>
      <q-card-section class="q-pt-none">
        <q-select v-model="assistOne" :options="optionsAssistOne" label="Assist 1" />
      </q-card-section>
      <q-card-section class="q-pt-none">
        <q-select v-model="assistTwo" :options="optionsAssistTwo" label="Assist 2" />
      </q-card-section>

      <q-card-actions align="right" class="text-primary">
        <q-btn flat label="Cancel" v-close-popup />
        <q-btn flat label="Submit" v-close-popup />
      </q-card-actions>
    </q-card>
  </q-dialog>
  </div>
</template>

<style scoped>

</style>
