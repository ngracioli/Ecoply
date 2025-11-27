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
    small: "text-xs px-2 py-1",
    medium: "text-sm px-3 py-1.5",
    large: "text-base px-4 py-2",
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
    class="inline-flex items-center gap-1.5"
    :class="[badgeConfig.class, sizeClasses]"
  >
    <i :class="badgeConfig.icon" class="text-xs"></i>
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
}

.tooltip-success .p-tooltip-arrow {
  border-top-color: rgb(240 255 250) !important;
}

.tooltip-danger .p-tooltip-text {
  background-color: rgb(254 242 242) !important;
  color: rgb(185 28 28) !important;
  border: 1px solid rgb(252 165 165) !important;
  box-shadow: 0 4px 12px rgba(239, 68, 68, 0.15) !important;
}

.tooltip-danger .p-tooltip-arrow {
  border-top-color: rgb(254 242 242) !important;
}
</style>
