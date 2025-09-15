<script setup lang="ts">
defineProps<{
  placeholder?: string;
  type?: string;
  id?: string;
}>();

const model = defineModel<string>();
</script>

<template>
  <div
    class="focus-within:border-primary-color flex w-full items-center rounded-lg border border-gray-500 transition-all duration-200 focus-within:ring-2 focus-within:ring-green-100"
  >
    <input
      :placeholder="placeholder"
      :type="type || 'text'"
      :id="id"
      v-model="model"
      class="flex-1 bg-transparent px-4 py-3 text-base placeholder-gray-600 outline-none"
    />
  </div>
</template>

<style scoped></style>
