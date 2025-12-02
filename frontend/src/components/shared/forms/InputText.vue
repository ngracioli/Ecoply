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
    class="focus-within:border-primary-color flex w-full items-center rounded-lg border border-gray-500 transition-all duration-200 focus-within:ring-1 focus-within:ring-green-100 focus-within:ring-offset-0 sm:focus-within:ring-2"
  >
    <input
      :placeholder="placeholder"
      :type="type || 'text'"
      :id="id"
      v-model="model"
      class="flex-1 bg-transparent px-3 py-2.5 text-sm placeholder-gray-600 outline-none sm:px-4 sm:py-3 sm:text-base"
    />
  </div>
</template>

<style scoped></style>
