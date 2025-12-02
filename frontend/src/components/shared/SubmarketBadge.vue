<script setup lang="ts">
import { computed } from "vue";
import { useStore } from "vuex";
import type { RootState } from "../../store";

interface Props {
  offerSubmarket: string;
  size?: "small" | "medium" | "large";
}

const props = withDefaults(defineProps<Props>(), {
  size: "medium",
});

const store = useStore<RootState>();
const userSubmarket = computed<string | null>(
  () => store.getters["user/userSubmarket"],
);

const isMatchingSubmarket = computed(() => {
  if (!userSubmarket.value) return false;
  return props.offerSubmarket === userSubmarket.value;
});

const badgeConfig = computed(() => {
  if (isMatchingSubmarket.value) {
    return {
      severity: "success" as const,
      tooltipText: "Seu agente CCEE é deste submercado",
      icon: "pi pi-check-circle",
      class: "submarket-badge-success",
    };
  }
  return {
    severity: "danger" as const,
    tooltipText: "Seu agente CCEE não é deste submercado",
    icon: "pi pi-exclamation-triangle",
    class: "submarket-badge-danger",
  };
});

const sizeClasses = computed(() => {
  const sizes = {
    small: "text-xs px-2 py-0.5 sm:py-1",
    medium: "text-xs px-2.5 py-1 sm:text-sm sm:px-3 sm:py-1.5",
    large: "text-sm px-3 py-1.5 sm:text-base sm:px-4 sm:py-2",
  };
  return sizes[props.size];
});
</script>

<template>
  <div
    v-tooltip.top="{
      value: badgeConfig.tooltipText,
      showDelay: 300,
      hideDelay: 100,
      class: isMatchingSubmarket ? 'tooltip-success' : 'tooltip-danger',
    }"
    class="inline-flex items-center gap-1 sm:gap-1.5"
    :class="[badgeConfig.class, sizeClasses]"
  >
    <i :class="badgeConfig.icon" class="text-[10px] sm:text-xs"></i>
    <span class="font-medium">{{ offerSubmarket }}</span>
  </div>
</template>

<style scoped>
.submarket-badge-success {
  background-color: rgb(240 253 244);
  color: rgb(21 128 61);
  border: 1px solid rgb(187 247 208);
  border-radius: 9999px;
}

.submarket-badge-success:hover {
  background-color: rgb(220 252 231);
  border-color: rgb(134 239 172);
}

.submarket-badge-danger {
  background-color: rgb(254 242 242);
  color: rgb(185 28 28);
  border: 1px solid rgb(254 202 202);
  border-radius: 9999px;
}

.submarket-badge-danger:hover {
  background-color: rgb(254 226 226);
  border-color: rgb(252 165 165);
}
</style>

<style>
.tooltip-success .p-tooltip-text {
  background-color: rgb(240 255 250) !important;
  color: rgb(21 128 61) !important;
  border: 1px solid rgb(200 247 218) !important;
  box-shadow: 0 4px 12px rgba(34, 197, 94, 0.15) !important;
  font-size: 0.75rem !important;
  padding: 0.5rem 0.75rem !important;
}

@media (min-width: 640px) {
  .tooltip-success .p-tooltip-text {
    font-size: 0.875rem !important;
    padding: 0.625rem 0.875rem !important;
  }
}

.tooltip-success .p-tooltip-arrow {
  border-top-color: rgb(240 255 250) !important;
}

.tooltip-danger .p-tooltip-text {
  background-color: rgb(254 242 242) !important;
  color: rgb(185 28 28) !important;
  border: 1px solid rgb(252 165 165) !important;
  box-shadow: 0 4px 12px rgba(239, 68, 68, 0.15) !important;
  font-size: 0.75rem !important;
  padding: 0.5rem 0.75rem !important;
}

@media (min-width: 640px) {
  .tooltip-danger .p-tooltip-text {
    font-size: 0.875rem !important;
    padding: 0.625rem 0.875rem !important;
  }
}

.tooltip-danger .p-tooltip-arrow {
  border-top-color: rgb(254 242 242) !important;
}
</style>
