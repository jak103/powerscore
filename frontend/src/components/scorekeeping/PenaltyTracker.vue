<script setup lang="ts">
import {reactive} from "vue";

defineProps<{
  team?: string
}>()

// TODO: Refactor to not use dummy data, either define props to be passed in or use computed. Also need UI/logic for adding/removing penalties
const penalties = reactive({
    0: {
      id: 12341234,
      player: 15,
      penalty: "2:00"
    },
    1: {
      id: 2345234534,
      player: 16,
      penalty: "1:32"
    }
})

</script>

<template>
  <div class="bg-gray-50 p-2 m-2 rounded-md shadow-md text-lg">
    <div class="font-semibold">{{ team }} Penalties Placeholder </div>
    <div class="grid grid-cols-2 gap-1">
      <div class="grid grid-cols-1 col-span-1">
        <h3>Player</h3>
        <div v-for="penalty in penalties" :key="penalty.id">{{penalty.player}}</div>
      </div>
      <div class="grid grid-cols-1 col-span-1">
        <h3>Penalty</h3>
        <div v-for="penalty in penalties" :key="penalty.id">{{penalty.penalty}}</div>
      </div>
    </div>
  </div>
</template>

<style scoped>

</style>