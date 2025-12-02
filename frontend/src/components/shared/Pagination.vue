<script setup lang="ts">
interface PaginationProps {
  currentPage: number;
  hasNext: boolean;
  hasPrev: boolean;
  itemsStart: number;
  itemsEnd: number;
  itemLabel?: string;
}

interface PaginationEmits {
  (e: "prev"): void;
  (e: "next"): void;
}

withDefaults(defineProps<PaginationProps>(), {
  itemLabel: "itens",
});

defineEmits<PaginationEmits>();
</script>

<template>
  <div
    class="flex flex-col items-stretch justify-between gap-3 rounded-lg border border-neutral-200 bg-white px-3 py-3 shadow-sm sm:flex-row sm:items-center sm:gap-4 sm:px-6 sm:py-4"
  >
    <button
      @click="$emit('prev')"
      :disabled="!hasPrev"
      :class="[
        'flex items-center justify-center gap-1.5 rounded-lg px-3 py-2 text-xs font-medium transition-all duration-200 active:scale-95 sm:gap-2 sm:px-4 sm:text-sm',
        hasPrev
          ? 'bg-emerald-500 text-white hover:bg-emerald-600 hover:shadow-md'
          : 'cursor-not-allowed bg-neutral-100 text-neutral-400',
      ]"
    >
      <svg
        xmlns="http://www.w3.org/2000/svg"
        class="h-3.5 w-3.5 sm:h-4 sm:w-4"
        viewBox="0 0 24 24"
        fill="none"
        stroke="currentColor"
        stroke-width="2"
        stroke-linecap="round"
        stroke-linejoin="round"
      >
        <polyline points="15 18 9 12 15 6"></polyline>
      </svg>
      <span class="hidden sm:inline">Anterior</span>
      <span class="sm:hidden">Ant.</span>
    </button>

    <div class="flex flex-col items-center gap-1 sm:gap-1">
      <div class="flex items-center gap-1.5 sm:gap-2">
        <span class="text-xs font-medium text-neutral-600 sm:text-sm"
          >Página</span
        >
        <span
          class="rounded-lg bg-emerald-50 px-2.5 py-0.5 text-xs font-bold text-emerald-600 sm:px-3 sm:py-1 sm:text-sm"
        >
          {{ currentPage }}
        </span>
      </div>
      <span class="text-[10px] text-neutral-500 sm:text-xs">
        {{ itemsStart }} - {{ itemsEnd }} {{ itemLabel }}
      </span>
    </div>

    <button
      @click="$emit('next')"
      :disabled="!hasNext"
      :class="[
        'flex items-center justify-center gap-1.5 rounded-lg px-3 py-2 text-xs font-medium transition-all duration-200 active:scale-95 sm:gap-2 sm:px-4 sm:text-sm',
        hasNext
          ? 'bg-emerald-500 text-white hover:bg-emerald-600 hover:shadow-md'
          : 'cursor-not-allowed bg-neutral-100 text-neutral-400',
      ]"
    >
      <span class="hidden sm:inline">Próxima</span>
      <span class="sm:hidden">Prox.</span>
      <svg
        xmlns="http://www.w3.org/2000/svg"
        class="h-3.5 w-3.5 sm:h-4 sm:w-4"
        viewBox="0 0 24 24"
        fill="none"
        stroke="currentColor"
        stroke-width="2"
        stroke-linecap="round"
        stroke-linejoin="round"
      >
        <polyline points="9 18 15 12 9 6"></polyline>
      </svg>
    </button>
  </div>
</template>

<style scoped></style>
