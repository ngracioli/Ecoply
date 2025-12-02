<script setup lang="ts">
import { ChevronLeft, ChevronRight } from "lucide-vue-next";

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
    class="flex items-center justify-between rounded-xl bg-white p-3 shadow-sm sm:rounded-2xl sm:p-4"
  >
    <button
      @click="$emit('prev')"
      :disabled="!hasPrev"
      :class="[
        'flex h-9 w-9 shrink-0 touch-manipulation items-center justify-center rounded-full transition-all duration-200 sm:h-10 sm:w-10',
        hasPrev
          ? 'bg-emerald-500 text-white shadow-md hover:bg-emerald-600 active:scale-95 active:shadow-sm'
          : 'cursor-not-allowed bg-neutral-100 text-neutral-300',
      ]"
    >
      <ChevronLeft :size="18" class="sm:hidden" :stroke-width="2.5" />
      <ChevronLeft :size="20" class="hidden sm:block" :stroke-width="2.5" />
    </button>

    <div class="flex min-w-0 flex-1 flex-col items-center gap-0.5 px-3">
      <div class="flex items-center gap-1.5">
        <span class="text-[10px] font-medium text-neutral-500 sm:text-xs"
          >Página</span
        >
        <span
          class="rounded-full bg-emerald-50 px-2.5 py-0.5 text-xs font-semibold text-emerald-600 sm:px-3 sm:text-sm"
        >
          {{ currentPage }}
        </span>
      </div>
      <span class="truncate text-[10px] text-neutral-400 sm:text-xs">
        {{ itemsStart }}-{{ itemsEnd }} {{ itemLabel }}
      </span>
    </div>

    <button
      @click="$emit('next')"
      :disabled="!hasNext"
      :class="[
        'flex h-9 w-9 shrink-0 touch-manipulation items-center justify-center rounded-full transition-all duration-200 sm:h-10 sm:w-10',
        hasNext
          ? 'bg-emerald-500 text-white shadow-md hover:bg-emerald-600 active:scale-95 active:shadow-sm'
          : 'cursor-not-allowed bg-neutral-100 text-neutral-300',
      ]"
    >
      <ChevronRight :size="18" class="sm:hidden" :stroke-width="2.5" />
      <ChevronRight :size="20" class="hidden sm:block" :stroke-width="2.5" />
    </button>
  </div>
</template>
