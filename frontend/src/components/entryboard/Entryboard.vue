<script setup lang="ts">
import { reactive } from 'vue';
import {ref} from 'vue'
import axios from 'axios'
import LockerRoomIdentifier from './LockerRoomIdentifier.vue';
import TeamInfo from './TeamInfo.vue';
import EditButton from './EditButton.vue';
import RefreshButton from './RefreshButton.vue'
import placeholderImage from "@/assets/img/hockeyClipart.jpg";

// TODO: version of this page explicitly for the entryway (grab everything from server without edit button)
// TODO: get editable version into main powerplay app

interface Team {
  id: number
  name: string,
  image?: HTMLImageElement
}

function getPlaceholderImage(): HTMLImageElement {
  var i = new Image();
  i.alt = "team logo";
  i.src = placeholderImage;
  i.className = "team-image";
  return i;
}

const teams = ref([])

async function refreshTeams() {
  // TODO: Authentication
  // TODO: Do this on startup and/or periodically
  await axios
    .get("http://localhost:3001/teams") // API endpoint here TODO: make this automatic somehow
    .then(response => {
        console.debug(response)
        teams.value = response.data
    })
}

function getTeams() : Team[] {
  return teams.value
}

const lockerRoomCount = 5;
const state = reactive({editMode: false});

function toggleEditMode() {
  state.editMode = !state.editMode;
}

</script>

<template>

<div class="entryboard">
  <EditButton @edit-toggle="toggleEditMode" />
  <RefreshButton @refresh-teams="refreshTeams" />
  <div class="assignment" v-for="i in 5">
    <LockerRoomIdentifier :number=i />
    <TeamInfo :teams=getTeams() :edit-mode=state.editMode />
  </div>
</div>

</template>

<style scoped>
</style>
