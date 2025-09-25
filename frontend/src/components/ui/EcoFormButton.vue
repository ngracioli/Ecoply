<script setup lang="ts">
import { computed } from "vue";

const props = withDefaults(
  defineProps<{
    type?: "button" | "submit" | "reset";
    variant?: "primary" | "secondary" | "outline" | "danger";
    size?: "sm" | "md" | "lg";
    loading?: boolean;
    disabled?: boolean;
    full?: boolean; // full width
  }>(),
  {
    type: "button",
    variant: "primary",
    size: "md",
  },
);

const emit = defineEmits<{ (e: "click", ev: MouseEvent): void }>();

const base =
  "relative inline-flex items-center justify-center font-semibold rounded-xl transition-colors duration-200 focus:outline-none disabled:cursor-not-allowed disabled:opacity-60 cursor-pointer";

const variants: Record<string, string> = {
  primary: "bg-primary-color text-white hover:bg-primary-dark-color",
  secondary: "bg-gray-800 text-white hover:bg-gray-700",
  outline: "border border-gray-400 text-gray-800 hover:bg-gray-100",
  danger: "bg-red-600 text-white hover:bg-red-500",
};

const sizes: Record<string, string> = {
  sm: "text-sm px-3 py-2",
  md: "text-sm px-5 py-3",
  lg: "text-base px-6 py-4",
};

const classes = computed(() => [
  base,
  variants[props.variant],
  sizes[props.size],
  props.full ? "w-full" : "",
]);

function onClick(ev: MouseEvent) {
  if (props.disabled || props.loading) {
    ev.preventDefault();
    return;
  }
  emit("click", ev);
}
</script>

<template>
  <button
    :type="type"
    :class="classes"
    :disabled="disabled || loading"
    @click="onClick"
  >
    <span v-if="loading" class="absolute inset-y-0 left-2 flex items-center">
      <svg
        class="h-4 w-4 animate-spin text-current"
        xmlns="http://www.w3.org/2000/svg"
        fill="none"
        viewBox="0 0 24 24"
      >
        <circle
          class="opacity-25"
          cx="12"
          cy="12"
          r="10"
          stroke="currentColor"
          stroke-width="4"
        />
        <path
          class="opacity-75"
          fill="currentColor"
          d="M4 12a8 8 0 018-8v4a4 4 0 00-4 4H4z"
        />
      </svg>
    </span>
    <span :class="loading ? 'opacity-0' : 'opacity-100'">
      <slot />
    </span>
  </button>
</template>
