<script setup lang="ts"></script>

<template>
  <div
    class="relative aspect-[16/10] w-full max-w-2xl overflow-hidden rounded-2xl border border-gray-200 bg-white px-6 py-5 shadow-sm"
  >
    <div
      class="bg-accent-color absolute top-0 left-0 h-full w-1 rounded-l-2xl"
    ></div>

    <div class="flex h-full flex-col items-stretch gap-4 sm:flex-row sm:gap-6">
      <div class="flex flex-1 flex-col justify-between gap-3">
        <div class="flex items-start justify-between">
          <div class="flex items-center gap-3">
            <span
              class="bg-accent-color rounded-full px-4 py-2 text-sm font-semibold text-white"
            >
              Oferta
            </span>

            <div class="flex items-center gap-2 text-sm text-gray-700">
              <i class="pi pi-clock text-gray-700"></i>
              <span>2 dias</span>
            </div>
          </div>
        </div>

        <div>
          <h1 class="text-3xl leading-tight font-semibold text-gray-900">
            Energia Solar Residencial
          </h1>

          <p class="mt-2 text-xl text-gray-600">
            Solução completa para residências — painel, inversor e instalação.
          </p>
        </div>

        <div class="flex items-start justify-start gap-6 text-sm text-gray-700">
          <div class="flex flex-col text-sm text-gray-700">
            <div class="flex items-center gap-2">
              <i class="pi pi-user text-gray-400"></i>
              <span class="font-medium">João Silva</span>
              <span class="text-gray-300">|</span>
              <i class="pi pi-map-marker text-gray-400"></i>
              <span class="text-sm text-gray-700">São Paulo, SP</span>
            </div>

            <div
              class="mt-1 flex items-center gap-2 text-xl font-semibold text-gray-700"
            >
              <i class="pi pi-bolt text-accent-color"></i>
              <span>150 kWh</span>
            </div>
          </div>
        </div>
      </div>

      <div class="flex w-full flex-col items-end justify-between gap-4 sm:w-44">
        <div class="text-right">
          <div class="text-sm text-gray-500">Preço da Unidade</div>
          <div class="text-accent-color text-2xl font-bold">R$ 0,45/kWh</div>
        </div>

        <div class="w-full sm:w-auto">
          <button
            class="text-accent-color border-accent-color hover:bg-accent-color focus:ring-accent-color/30 w-full rounded-lg border-2 bg-white px-4 py-2 text-sm font-semibold transition-colors duration-300 ease-out hover:border-transparent hover:text-white hover:shadow-md focus:ring-2 focus:outline-none sm:w-40"
          >
            Ver Oferta
          </button>
        </div>
      </div>
    </div>
  </div>
</template>
