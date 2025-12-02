<script setup lang="ts"></script>

<template>
  <div
    class="relative w-full max-w-2xl overflow-hidden rounded-xl border border-gray-200 bg-white px-4 py-4 shadow-sm sm:rounded-2xl sm:px-6 sm:py-5"
  >
    <div
      class="bg-accent-color absolute top-0 left-0 h-full w-1 rounded-l-xl sm:rounded-l-2xl"
    ></div>

    <div
      class="flex h-full flex-col items-stretch gap-3 sm:flex-row sm:gap-4 md:gap-6"
    >
      <div class="flex flex-1 flex-col justify-between gap-2 sm:gap-3">
        <div class="flex items-start justify-between">
          <div class="flex flex-wrap items-center gap-2 sm:gap-3">
            <span
              class="bg-accent-color rounded-full px-3 py-1 text-xs font-semibold text-white sm:px-4 sm:py-2 sm:text-sm"
            >
              Oferta
            </span>

            <div
              class="flex items-center gap-1.5 text-xs text-gray-700 sm:gap-2 sm:text-sm"
            >
              <i class="pi pi-clock text-[10px] text-gray-700 sm:text-xs"></i>
              <span>2 dias</span>
            </div>
          </div>
        </div>

        <div>
          <h1
            class="text-xl leading-tight font-semibold text-gray-900 sm:text-2xl md:text-3xl"
          >
            Energia Solar Residencial
          </h1>

          <p
            class="mt-1.5 text-sm text-gray-600 sm:mt-2 sm:text-base md:text-xl"
          >
            Solução completa para residências — painel, inversor e instalação.
          </p>
        </div>

        <div
          class="flex items-start justify-start gap-4 text-xs text-gray-700 sm:gap-6 sm:text-sm"
        >
          <div class="flex flex-col text-xs text-gray-700 sm:text-sm">
            <div class="flex flex-wrap items-center gap-1.5 sm:gap-2">
              <i class="pi pi-user text-[10px] text-gray-400 sm:text-xs"></i>
              <span class="font-medium">João Silva</span>
              <span class="text-gray-300">|</span>
              <i
                class="pi pi-map-marker text-[10px] text-gray-400 sm:text-xs"
              ></i>
              <span class="text-xs text-gray-700 sm:text-sm"
                >São Paulo, SP</span
              >
            </div>

            <div
              class="mt-1 flex items-center gap-1.5 text-base font-semibold text-gray-700 sm:gap-2 sm:text-lg md:text-xl"
            >
              <i class="pi pi-bolt text-accent-color text-sm sm:text-base"></i>
              <span>150 kWh</span>
            </div>
          </div>
        </div>
      </div>

      <div
        class="flex w-full flex-col items-end justify-between gap-3 sm:w-44 sm:gap-4"
      >
        <div class="text-right">
          <div class="text-xs text-gray-500 sm:text-sm">Preço da Unidade</div>
          <div class="text-accent-color text-xl font-bold sm:text-2xl">
            R$ 0,45/kWh
          </div>
        </div>

        <div class="w-full sm:w-auto">
          <button
            class="text-accent-color border-accent-color hover:bg-accent-color focus:ring-accent-color/30 w-full rounded-lg border-2 bg-white px-4 py-2 text-xs font-semibold transition-all duration-300 ease-out hover:border-transparent hover:text-white hover:shadow-md focus:ring-2 focus:outline-none active:scale-[0.98] sm:w-40 sm:text-sm"
          >
            Ver Oferta
          </button>
        </div>
      </div>
    </div>
  </div>
</template>
