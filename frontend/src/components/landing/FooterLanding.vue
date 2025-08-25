<script setup lang="ts">
import IconSquared from "../shared/IconSquared.vue";
</script>

<template>
  <footer class="border-t border-gray-200 bg-white">
    <!-- Container principal usando FLEX -->
    <div
      class="mx-auto flex max-w-6xl items-start justify-between space-x-16 px-8 py-10"
    >
      <!-- Coluna 1 - SunShare -->
      <div class="flex max-w-xs flex-col items-start">
        <div class="mb-2">
          <h1 class="text-primary-color text-2xl font-extrabold">Ecoply</h1>
        </div>
        <p class="mt-2 text-base text-gray-700">
          Conectando produtores e consumidores de energia renovável para um
          futuro mais sustentável e descentralizado.
        </p>
        <div class="mt-6 flex space-x-4 text-2xl">
          <a
            href="#"
            class="transform transition hover:scale-110 hover:text-green-500"
          >
            <IconSquared :icon="'pi-facebook'" />
          </a>
          <a
            href="#"
            class="transform transition hover:scale-110 hover:text-green-500"
          >
            <IconSquared :icon="'pi-twitter'" />
          </a>
          <a
            href="#"
            class="transform transition hover:scale-110 hover:text-green-500"
          >
            <IconSquared :icon="'pi-instagram'" />
          </a>
          <a
            href="#"
            class="transform transition hover:scale-110 hover:text-green-500"
          >
            <IconSquared :icon="'pi-linkedin'" />
          </a>
        </div>
      </div>

      <div class="flex min-w-[180px] flex-col">
        <h2 class="text-primary-color mb-3 text-lg font-semibold">Navegação</h2>
        <ul class="space-y-2 text-base text-gray-700">
          <li>
            <a href="#" class="transition hover:text-green-500">Início</a>
          </li>
          <li>
            <a href="#" class="transition hover:text-green-500"
              >Como Funciona</a
            >
          </li>
          <li>
            <a href="#" class="transition hover:text-green-500">Benefícios</a>
          </li>
          <li>
            <a href="#" class="transition hover:text-green-500">Impacto</a>
          </li>
          <li>
            <a href="#" class="transition hover:text-green-500">Contato</a>
          </li>
        </ul>
      </div>

      <div class="flex min-w-[200px] flex-col items-start">
        <h2 class="text-primary-color mb-3 text-lg font-semibold">Contato</h2>
        <ul class="space-y-2 text-base text-gray-700">
          <li>
            <a
              href="mailto:contato@sunshare.com.br"
              class="transition hover:text-green-500"
              >contato@ecoply.com.br</a
            >
          </li>
          <li>
            <a href="tel:+551999999999" class="transition hover:text-green-500"
              >(19) 9999-9999</a
            >
          </li>
          <li>Rio Claro, SP</li>
        </ul>
      </div>
    </div>

    <div
      class="mx-auto mt-8 max-w-3/4 border-t border-gray-200 py-4 text-center text-sm text-gray-500"
    >
      © 2025 <span class="text-primary-color font-bold">Ecoply</span>. Todos os
      direitos reservados.
    </div>
  </footer>
</template>

<style scoped>
footer {
  font-family: "Inter", "Roboto", Arial, sans-serif;
}
</style>
