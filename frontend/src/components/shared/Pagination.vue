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
    class="flex items-center justify-between rounded-lg border border-neutral-200 bg-white px-6 py-4 shadow-sm"
  >
    <button
      @click="$emit('prev')"
      :disabled="!hasPrev"
      :class="[
        'flex items-center gap-2 rounded-lg px-4 py-2 text-sm font-medium transition-all duration-200',
        hasPrev
          ? 'bg-emerald-500 text-white hover:bg-emerald-600 hover:shadow-md'
          : 'cursor-not-allowed bg-neutral-100 text-neutral-400',
      ]"
    >
      <svg
        xmlns="http://www.w3.org/2000/svg"
        width="16"
        height="16"
        viewBox="0 0 24 24"
        fill="none"
        stroke="currentColor"
        stroke-width="2"
        stroke-linecap="round"
        stroke-linejoin="round"
      >
        <polyline points="15 18 9 12 15 6"></polyline>
      </svg>
      Anterior
    </button>

    <div class="flex flex-col items-center gap-1">
      <div class="flex items-center gap-2">
        <span class="text-sm font-medium text-neutral-600">Página</span>
        <span
          class="rounded-lg bg-emerald-50 px-3 py-1 text-sm font-bold text-emerald-600"
        >
          {{ currentPage }}
        </span>
      </div>
      <span class="text-xs text-neutral-500">
        Mostrando {{ itemsStart }} - {{ itemsEnd }} {{ itemLabel }}
      </span>
    </div>

    <button
      @click="$emit('next')"
      :disabled="!hasNext"
      :class="[
        'flex items-center gap-2 rounded-lg px-4 py-2 text-sm font-medium transition-all duration-200',
        hasNext
          ? 'bg-emerald-500 text-white hover:bg-emerald-600 hover:shadow-md'
          : 'cursor-not-allowed bg-neutral-100 text-neutral-400',
      ]"
    >
      Próxima
      <svg
        xmlns="http://www.w3.org/2000/svg"
        width="16"
        height="16"
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
