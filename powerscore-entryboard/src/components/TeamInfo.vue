<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue';

//TODO: common location for interfaces.
interface Team {
  id: number,
  name: string,
  image?: HTMLImageElement
}

const props = defineProps<{
  teams: Team[],
  id: number,
}>();

const placeholderTeam = {id: 0, name: "Waiting For Connection..."};
const state: {team: Team} = reactive({team: placeholderTeam});

function setTeam() {
  state.team = props.teams.find((t) => t.id == props.id) || {id: 70, name: "Waiting for Assignment..."};
}

onMounted( () => {
    setTeam()
    setInterval( () => {
    setTeam()
  }, 10000)
})

</script>

<template>
  <div class="assignment-team-info">
    <div class="assignment-team-picture" v-html=state.team.image?.outerHTML></div>
    <div class="assignment-team-name">
      <div ref="name">
        {{ state.team.name }}
      </div>
    </div>
    
  </div>
</template>

<style scoped>
</style>
