<script setup lang="ts">
defineProps<{
  icon: string;
}>();
</script>

<template>
  <div
    class="bg-secondary-color flex items-center justify-center rounded-xl p-3 transition-transform duration-300 group-hover:scale-110 sm:rounded-2xl sm:p-4"
  >
    <i :class="['pi', icon, 'text-primary-color', 'text-xl sm:text-2xl']"></i>
  </div>
</template>
