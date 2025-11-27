<script setup lang="ts">
import { computed } from "vue";

interface Props {
  currentValue: number;
  totalValue: number;
  showPercentage?: boolean;
  showLabel?: boolean;
  label?: string;
  height?: "sm" | "md" | "lg";
  gradientFrom?: string;
  gradientTo?: string;
}

const props = withDefaults(defineProps<Props>(), {
  showPercentage: true,
  showLabel: false,
  height: "sm",
  gradientFrom: "emerald-500",
  gradientTo: "green-500",
});

const percentage = computed(() => {
  if (!props.currentValue || !props.totalValue) return "0";
  const value = (props.currentValue / props.totalValue) * 100;

  if (value < 1 && value > 0) {
    return value.toFixed(2);
  }

  if (value < 10) {
    return value.toFixed(1);
  }

  return value.toFixed(0);
});

const heightClass = computed(() => {
  const heights = {
    sm: "h-2",
    md: "h-3",
    lg: "h-4",
  };
  return heights[props.height];
});
</script>

<template>
  <div>
    <div
      v-if="showLabel && label"
      class="mb-2 text-sm font-medium text-neutral-600"
    >
      {{ label }}
    </div>
    <div class="flex items-center gap-3">
      <div
        :class="[
          'flex-1 overflow-hidden rounded-full bg-neutral-200',
          heightClass,
        ]"
      >
        <div
          :style="{ width: `${percentage}%` }"
          :class="[
            'h-full bg-gradient-to-r transition-all',
            `from-${gradientFrom} to-${gradientTo}`,
          ]"
        ></div>
      </div>
      <span
        v-if="showPercentage"
        class="text-sm font-semibold text-neutral-700"
      >
        {{ percentage }}%
      </span>
    </div>
  </div>
</template>
