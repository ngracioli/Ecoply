<script setup lang="ts">
const props = withDefaults(
  defineProps<{
    as?: string;
    disabled?: boolean;
  }>(),
  {
    as: "form",
  },
);

const emit = defineEmits<{
  (e: "submit", ev: Event): void;
}>();

function onSubmit(ev: Event) {
  if (props.disabled) {
    ev.preventDefault();
    return;
  }
  emit("submit", ev);
}
</script>

<template>
  <component
    :is="props.as"
    @submit.prevent="onSubmit"
    class="space-y-4 sm:space-y-5 md:space-y-6"
  >
    <slot />
  </component>
</template>
