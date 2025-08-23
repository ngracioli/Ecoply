<script setup lang="ts">
defineProps<{
  icon: string;
}>();
</script>

<template>
  <div
    class="bg-primary-color flex h-20 w-20 items-center justify-center rounded-full"
  >
    <i :class="['pi', icon, 'text-5xl', 'text-white']"></i>
  </div>
</template>
