<script setup lang="ts">
import { reactive } from 'vue';
import LockerRoomIdentifier from './LockerRoomIdentifier.vue';
import TeamInfo from './TeamInfo.vue';
import EditButton from './EditButton.vue';
import placeholderImage from "@/assets/img/hockeyClipart.jpg";


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

function getTeams() : Team[] {
  //TODO: This function should ask the backend for the data.
  //This is just dummy data.
  return [
    {
      id: 0,
      name: "No Assignment",
    },
    {
      id: 1,
      name: "Awesome Team",
      image: getPlaceholderImage(),
    },
    {
      id: 2,
      name: "Really Cool People With A Really Long Name",
      image: getPlaceholderImage(),
    },
    {
      id: 3,
      name: "Keyboard Warriors",
      image: getPlaceholderImage(),
    },
    {
      id: 4,
      name: "Parentheticals",
      image: getPlaceholderImage(),
    },
    {
      id: 5,
      name: "Rockets",
      image: getPlaceholderImage(),
    },
    {
      id: 6,
      name: "Balloons",
      image: getPlaceholderImage(),
    },
  ]
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
  <div class="assignment" v-for="i in 5">
    <LockerRoomIdentifier :number=i />
    <TeamInfo :teams=getTeams() :edit-mode=state.editMode />
  </div>
</div>

</template>

<style scoped>
</style>
