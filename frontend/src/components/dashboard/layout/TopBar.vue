<script setup lang="ts">
import { computed } from "vue";
import { useStore } from "vuex";
import type { RootState } from "../../../store";

type GreetingPeriod = "Bom dia" | "Boa tarde" | "Boa noite";

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
    class="flex items-center justify-between border-b border-neutral-200 bg-white px-8 py-4"
  >
    <div class="flex items-center gap-2">
      <h1 class="text-2xl font-light text-neutral-800">
        {{ greeting }},
        <span class="font-semibold text-neutral-900">{{ userFullName }}</span>
      </h1>
      <span class="text-2xl">👋</span>
    </div>
  </header>
</template>
