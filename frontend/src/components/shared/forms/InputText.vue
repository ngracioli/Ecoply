<script setup lang="ts">
defineProps<{
  placeholder?: string;
  type?: string;
  id?: string;
}>();
</script>

<template>
  <div
    class="w-full flex items-center rounded-lg border border-gray-500 transition-all duration-200 focus-within:border-green-500 focus-within:ring-2 focus-within:ring-green-100"
  >
    <input
      :placeholder="placeholder"
      :type="type || 'text'"
      :id="id"
      class="flex-1 bg-transparent px-4 py-3 text-base placeholder-gray-600 outline-none"
    />
  </div>
</template>

<style scoped></style>
