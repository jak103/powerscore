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
  editMode?: boolean,
}>();


const select = ref<HTMLSelectElement | null>(null);
const name = ref<HTMLDivElement | null>(null);
const placeholderTeam = {id: 0, name: "No Assignent"};
const state: {team: Team} = reactive({team: placeholderTeam});

onMounted(() => {
  select.value?.focus();
  name.value?.focus();
})

function setTeam() {
  console.log(select.value?.selectedIndex);
  console.log(select.value?.options);
  var id = select.value?.options[select.value?.selectedIndex].value || 0;
  console.log(id);
  state.team = props.teams.find((t) => t.id == id) || {id: 70, name: "AH"};
}

</script>

<template>
  <div class="assignment-team-info">
    <div class="assignment-team-picture" v-html=state.team.image?.outerHTML></div>
    <div class="assignment-team-name" @change="setTeam">
      <select ref="select" name="team" v-if="props.editMode">
        <option v-for="optionTeam in props.teams" :value=optionTeam.id >{{ optionTeam.name }}</option> 
      </select>
      <div ref="name" v-if="!props.editMode">
        {{ state.team.name }}
      </div>
    </div>
    
  </div>
</template>

<style scoped>
</style>
