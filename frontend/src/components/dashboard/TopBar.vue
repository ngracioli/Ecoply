<script setup lang="ts">
import { computed } from "vue";
import { useStore } from "vuex";
import { Bell } from "lucide-vue-next";
import type { RootState } from "../../store";

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

const notificationCount = computed(() => 3);

const hasNotifications = computed(() => notificationCount.value > 0);
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

    <div class="flex items-center gap-4">
      <button
        class="relative rounded-lg p-2 text-neutral-600 transition-all duration-200 hover:bg-neutral-100 hover:text-neutral-900"
        aria-label="Notificações"
      >
        <Bell :size="22" :stroke-width="2" />
        <span
          v-if="hasNotifications"
          class="absolute top-1 right-1 flex h-5 w-5 items-center justify-center rounded-full bg-red-500 text-[10px] font-bold text-white shadow-md"
        >
          {{ notificationCount }}
        </span>
      </button>
    </div>
  </header>
</template>
