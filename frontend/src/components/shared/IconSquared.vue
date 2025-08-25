<script setup lang="ts">
defineProps<{
  icon: string;
}>();
</script>

<template>
  <div
    class="bg-primary-color/20 flex items-center justify-center rounded-2xl p-4 transition-transform duration-300 group-hover:scale-110"
  >
    <i :class="['pi', icon, 'text-primary-color']"></i>
  </div>
</template>
