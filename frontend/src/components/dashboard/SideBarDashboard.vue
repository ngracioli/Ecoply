<script setup lang="ts">
const menuItemClass =
  "cursor-pointer w-full text-gray-700 rounded-sm px-2 py-1 hover:bg-neutral-200 hover:text-black transition-colors duration-200";
</script>

<template>
  <nav class="flex h-screen w-full max-w-1/8 flex-col gap-6 bg-white">
    <section class="flex flex-col gap-2 px-4 py-5">
      <h1 class="px-2">PRINCIPAL</h1>
      <ul>
        <li :class="menuItemClass">Dashboard</li>
        <li :class="menuItemClass">Ofertas</li>
        <li :class="menuItemClass">Transações</li>
      </ul>
    </section>
    <section class="flex flex-col gap-2 px-4 py-5">
      <h1 class="px-2">PERFIL</h1>
      <ul>
        <li :class="menuItemClass">Perfil</li>
        <li :class="menuItemClass">Suporte</li>
      </ul>
    </section>
  </nav>
</template>

<style scoped></style>
