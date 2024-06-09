<script setup lang="ts">
import { ref } from 'vue'

// TODO: Make options values dynamic on load
// TODO: Create event listeners to pop up this form automatically when a penalty is triggered from the rink device
// TODO: Hook up submit button to the backend API
const promptPenalty = ref(false)

const offender = ref('')
const optionsOffender = [
  '11 - John',
  '15 - Jane',
  '21 - Eric',
  '46 - Randy',
  '76 - Jacques',
  '65 - Melinda',
  '17 - Jack'
]

const penalty = ref('')
const optionsPenalty = [
  'Cross-Checking',
  'Roughing',
  'Interference',
  'Made Up Penalty'
]

</script>

<template>
  <div class="q-pa-md q-gutter-sm">
    <q-btn push class="big-button" label="New Penalty" color="red" @click="promptPenalty = true" />

    <q-dialog v-model="promptPenalty" persistent>
      <q-card style="min-width: 350px">
        <q-card-section>
          <div class="text-h6">New Penalty</div>
        </q-card-section>

        <q-card-section class="q-pt-none">
          <q-select v-model="offender" :options="optionsOffender" label="Offender" />
        </q-card-section>
        <q-card-section class="q-pt-none">
          <q-select v-model="penalty" :options="optionsPenalty" label="Associated Penalty" />
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
.big-button {
  font-size: 24px;
}
</style>
