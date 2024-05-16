<script setup lang="ts">
import { reactive } from 'vue';

const props = defineProps<{
  team: "home" | "away",
  penalties?: {player: number, time: string}[]
}>();
</script>

<template>
  <div class="penalties">
    <table>
      <tbody>
        <tr>
          <th>PLYR</th>
          <th>Penalty</th>
        </tr>
        <tr v-for="penalty in props.penalties">
          <td class="penalty-player">{{ penalty.player }}</td>
          <td class="penalty-time"> {{ penalty.time }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<style scoped>
</style>
