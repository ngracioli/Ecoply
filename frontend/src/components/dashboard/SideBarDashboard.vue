<script setup lang="ts">
import {
  Leaf,
  Home,
  FileText,
  CreditCard,
  User,
  HelpCircle,
} from "lucide-vue-next";

const menuItems = [
  { label: "Dashboard", icon: Home },
  { label: "Ofertas", icon: FileText },
  { label: "Transações", icon: CreditCard },
  { label: "Perfil", icon: User },
  { label: "Suporte", icon: HelpCircle },
];
const menuItemClass =
  "cursor-pointer w-full text-gray-700 rounded-sm px-2 py-1 hover:bg-neutral-200 hover:text-black transition-colors duration-200";
</script>

<template>
  <nav
    class="flex h-screen w-full max-w-1/8 flex-col border-r border-black/10 bg-white"
  >
    <header
      class="flex flex-row items-center justify-center gap-4 border-b border-black/10 py-6"
    >
      <div
        class="bg-primary-dark-color flex h-8 w-8 items-center justify-center rounded-lg"
      >
        <Leaf class="h-5 w-5 text-white" />
      </div>
      <div class="flex flex-col">
        <h1 class="text-primary-dark-color text-xl font-bold">Ecoply</h1>
        <p class="text-sm text-neutral-500">Oferte Energia</p>
      </div>
    </header>
    <section class="flex flex-col gap-2 p-6">
      <h1 class="mb-2 px-2 text-sm font-medium text-neutral-500">
        Navegação Principal
      </h1>
      <ul class="flex flex-col gap-2">
        <li
          v-for="item in menuItems"
          :key="item.label"
          :class="menuItemClass"
          class="flex items-center gap-2"
        >
          <component :is="item.icon" class="h-5 w-5 shrink-0" />
          {{ item.label }}
        </li>
      </ul>
    </section>
  </nav>
</template>

<style scoped></style>
