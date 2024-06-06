<script setup lang="ts">
import NotFound from './NotFound.vue';
import Entryboard from './components/Entryboard.vue';
import { ref, computed} from 'vue';

const routes = {
  '/entryboard': Entryboard,
}

const currentPath = ref(window.location.hash);

window.addEventListener('hashchange', () => {
  currentPath.value = window.location.hash
});

const currentView = computed(() => {
  //@ts-ignore routes[<string>] does not match literals provided above. That is OK because of the NotFound option.
  return routes[currentPath.value.slice(1) || '/'] || Entryboard
})

</script>

<template>
  <header>
  </header>

  <main>
    <component :is=currentView />
  </main>
</template>

<style scoped>
</style>
