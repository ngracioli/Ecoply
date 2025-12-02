<script setup lang="ts">
import { computed } from "vue";
import { useStore } from "vuex";
import { Menu } from "lucide-vue-next";
import type { RootState } from "../../../store";

type GreetingPeriod = "Bom dia" | "Boa tarde" | "Boa noite";

const emit = defineEmits<{
  "toggle:mobileMenu": [];
}>();

const store = useStore<RootState>();

const getGreetingByHour = (hour: number): GreetingPeriod => {
  if (hour < 12) return "Bom dia";
  if (hour < 18) return "Boa tarde";
  return "Boa noite";
};

const greeting = computed<GreetingPeriod>(() => {
  const currentHour = new Date().getHours();
  return getGreetingByHour(currentHour);
});

const userFullName = computed(() => store.getters["user/userName"] || "");
</script>

<template>
  <header
    class="flex items-center justify-between border-b border-neutral-200 bg-white px-4 py-3 sm:px-6 sm:py-4 md:px-8"
  >
    <div class="flex min-w-0 flex-1 items-center gap-2 sm:gap-3">
      <button
        @click="emit('toggle:mobileMenu')"
        class="flex shrink-0 touch-manipulation items-center justify-center rounded-lg border border-neutral-200 bg-white p-2 transition-colors hover:bg-neutral-50 active:bg-neutral-100 md:hidden"
      >
        <Menu :size="20" class="text-neutral-700" />
      </button>

      <h1
        class="min-w-0 text-base font-light text-neutral-800 sm:text-xl md:text-2xl"
      >
        <span class="inline sm:inline">{{ greeting }}, </span>
        <span
          class="inline-block max-w-[150px] truncate align-bottom font-semibold text-neutral-900 sm:max-w-none"
          >{{ userFullName }}</span
        >
      </h1>
      <span class="shrink-0 text-xl sm:text-2xl">👋</span>
    </div>
  </header>
</template>
