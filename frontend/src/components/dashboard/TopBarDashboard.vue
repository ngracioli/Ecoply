<script setup lang="ts">
import { computed } from "vue";

const sidebarOpen = defineModel<boolean>("sidebarOpen", { required: true });

const buttonIcon = computed(() =>
  sidebarOpen.value ? "pi-arrow-left" : "pi-arrow-right",
);

const toggleSidebar = () => {
  sidebarOpen.value = !sidebarOpen.value;
};
</script>

<template>
  <header
    class="flex h-16 w-full flex-row items-center justify-between border-b border-black/10 bg-white px-6"
  >
    <section class="flex flex-1 flex-row items-center gap-4">
      <button @click="toggleSidebar">
        <i :class="'pi ' + buttonIcon"></i>
      </button>
      <div class="flex flex-col">
        <h1 class="text-lg font-semibold text-gray-800">
          Boa tarde, João Silva! 🌱
        </h1>
        <p class="text-sm text-gray-600">
          Vamos fazer a diferença com energia limpa hoje
        </p>
      </div>
    </section>
    <section class="flex flex-row items-center gap-4">
      <button><i class="pi pi-bell"></i></button>
      <button><i class="pi pi-cog"></i></button>
    </section>
  </header>
</template>

<style scoped>
header > section > button > i {
  color: var(--color-gray-500);
}

header > section > button {
  cursor: pointer;
  padding: 0.75rem;
  margin: 0;
  border: none;
  background: none;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  border-radius: 0.5rem;
  transition: background-color 0.2s ease;
}

header > section > button:hover {
  background-color: rgba(0, 0, 0, 0.08);
}
</style>
