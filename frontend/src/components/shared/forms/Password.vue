<script setup lang="ts">
import { ref } from "vue";

defineProps<{
  placeholder?: string;
  id?: string;
}>();

const showPassword = ref(false);

const model = defineModel<string>();

const togglePassword = () => {
  showPassword.value = !showPassword.value;
};
</script>

<template>
  <div
    class="focus-within:border-primary-color flex w-full items-center rounded-lg border border-gray-500 transition-all duration-200 focus-within:ring-1 focus-within:ring-green-100 focus-within:ring-offset-0 sm:focus-within:ring-2"
  >
    <input
      :type="showPassword ? 'text' : 'password'"
      :placeholder="placeholder"
      :id="id"
      v-model="model"
      class="flex-1 bg-transparent px-3 py-2.5 text-sm placeholder-gray-600 outline-none sm:px-4 sm:py-3 sm:text-base"
    />
    <button
      type="button"
      @click="togglePassword"
      class="flex cursor-pointer items-center justify-center px-2.5 text-gray-500 transition-colors duration-200 hover:text-gray-700 active:text-gray-800 sm:px-3"
    >
      <i
        :class="showPassword ? 'pi pi-eye-slash' : 'pi pi-eye'"
        class="text-base leading-none sm:text-lg"
      ></i>
    </button>
  </div>
</template>

<style scoped></style>
