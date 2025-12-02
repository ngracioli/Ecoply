<script setup lang="ts">
defineProps<{
  icon: string;
}>();
</script>

<template>
  <div
    class="bg-primary-color flex h-16 w-16 items-center justify-center rounded-full sm:h-20 sm:w-20"
  >
    <i
      :class="['pi', icon, 'text-4xl sm:text-5xl', 'text-secondary-color']"
    ></i>
  </div>
</template>
