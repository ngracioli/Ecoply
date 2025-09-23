<script setup lang="ts">
import { ref } from "vue";
import SideBarDashboard from "../../components/dashboard/SideBarDashboard.vue";
import TopBarDashboard from "../../components/dashboard/TopBarDashboard.vue";

const isSidebarOpen = ref(true);
</script>

<template>
  <main class="flex h-screen w-full bg-neutral-700">
    <SideBarDashboard v-if="isSidebarOpen" />
    <div class="flex flex-1 flex-col">
      <TopBarDashboard v-model:sidebarOpen="isSidebarOpen" />
      <div class="flex-1 p-6 text-gray-800">
        <p>Bem-vindo ao seu dashboard!</p>
      </div>
    </div>
  </main>
</template>

<style scoped></style>
