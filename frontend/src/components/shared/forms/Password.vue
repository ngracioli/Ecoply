<script setup lang="ts">
import { ref } from "vue";

defineProps<{
  placeholder?: string;
  id?: string;
}>();

const showPassword = ref(false);

const togglePassword = () => {
  showPassword.value = !showPassword.value;
};
</script>

<template>
  <div
    class="flex w-full items-center rounded-lg border border-gray-500 transition-all duration-200 focus-within:border-green-500 focus-within:ring-2 focus-within:ring-green-100"
  >
    <input
      :type="showPassword ? 'text' : 'password'"
      :placeholder="placeholder"
      :id="id"
      class="flex-1 bg-transparent px-4 py-3 text-base placeholder-gray-600 outline-none"
    />
    <button
      type="button"
      @click="togglePassword"
      class="flex cursor-pointer items-center justify-center px-3 text-gray-500 transition-colors duration-200 hover:text-gray-700"
    >
      <i
        :class="showPassword ? 'pi pi-eye-slash' : 'pi pi-eye'"
        class="text-lg leading-none"
      ></i>
    </button>
  </div>
</template>

<style scoped></style>
