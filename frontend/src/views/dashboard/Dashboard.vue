<script setup lang="ts">
import SideBarDashboard from "../../components/dashboard/SideBarDashboard.vue";

import { useStore } from "vuex";
import { useRouter } from "vue-router";

const store = useStore();
const router = useRouter();

const logout = async () => {
  await store.dispatch("auth/logout");
  router.push({ name: "Login" });
};
</script>

<template>
  <main class="flex h-screen w-full bg-neutral-700">
    <SideBarDashboard />
    <button @click="logout">Logout</button>
  </main>
</template>

<style scoped></style>
