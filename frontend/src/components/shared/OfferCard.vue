<script setup lang="ts"></script>

<template>
  <div
    class="relative flex w-full max-w-md flex-col gap-4 rounded-2xl border border-gray-200 bg-white p-6 shadow-sm"
  >
    <div
      class="absolute top-0 left-0 h-full w-1 rounded-l-2xl bg-green-500"
    ></div>

    <div class="gap flex items-center justify-between">
      <span
        class="rounded-full bg-green-100 px-3 py-1 text-xs font-medium text-green-600"
      >
        Oferta
      </span>

      <div class="flex items-center gap-1 text-xs text-gray-400">
        <i class="pi pi-clock text-gray-400"></i>
        <span>2 dias</span>
      </div>
    </div>

    <h1 class="text-lg font-semibold text-gray-800">
      Energia Solar Residencial
    </h1>

    <div class="flex items-center gap-1 text-sm text-gray-500">
      <i class="pi pi-map-marker"></i>
      <span>São Paulo, SP</span>
    </div>

    <div class="mt-1 flex items-center justify-between">
      <div
        class="flex items-center gap-1 text-base font-semibold text-gray-700"
      >
        <i class="pi pi-bolt text-green-500"></i>
        <span>150 kWh</span>
      </div>
      <span class="font-semibold text-green-600">R$ 0,45/kWh</span>
    </div>

    <span class="text-sm text-gray-400">Vendedor: João Silva</span>

    <button
      class="mt-2 w-full rounded-lg border border-gray-200 py-2 text-sm font-medium text-gray-700 transition hover:border-green-500 hover:text-green-600"
    >
      Ver Oferta
    </button>
  </div>
</template>

<style scoped></style>
