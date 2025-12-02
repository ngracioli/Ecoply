<script setup lang="ts">
import { computed, ref } from "vue";

const props = withDefaults(
  defineProps<{
    modelValue: string | number | null;
    id?: string;
    label?: string;
    hint?: string;
    errorMessage?: string;
    required?: boolean;
    placeholder?: string;
    disabled?: boolean;
    type?: string;
    password?: boolean;
    showPasswordLabel?: string;
    hidePasswordLabel?: string;
  }>(),
  {
    type: "text",
    password: false,
    showPasswordLabel: "Mostrar senha",
    hidePasswordLabel: "Esconder senha",
  },
);

const emit = defineEmits<{
  (e: "update:modelValue", v: string | number | null): void;
  (e: "toggle:password", v: boolean): void;
}>();

const hasError = computed(() => !!props.errorMessage);

const wrapperClasses = computed(() => [
  "flex w-full items-center rounded-lg border transition-all duration-200 focus-within:ring-1 focus-within:ring-offset-0 sm:focus-within:ring-2",
  props.disabled
    ? "opacity-60 cursor-not-allowed"
    : "focus-within:ring-green-100",
  hasError.value
    ? "border-red-500 text-red-500 focus-within:border-red-500"
    : "border-gray-500 focus-within:border-primary-color",
]);

function onInput(e: Event) {
  const target = e.target as HTMLInputElement;
  emit("update:modelValue", target.value);
}

const show = ref(false);
const inputType = computed(() => {
  if (props.password) {
    return show.value ? "text" : "password";
  }
  return props.type;
});

function togglePassword() {
  show.value = !show.value;
  emit("toggle:password", show.value);
}
</script>

<template>
  <div class="w-full space-y-1 sm:space-y-1.5">
    <label
      v-if="label"
      :for="id"
      :class="{ 'text-red-500': hasError }"
      class="block text-xs font-medium text-gray-700 sm:text-sm"
    >
      {{ label }}<span v-if="required" class="text-red-500">*</span>
    </label>

    <div :class="wrapperClasses">
      <input
        :id="id"
        :disabled="disabled"
        :placeholder="placeholder"
        :type="inputType"
        :value="modelValue as any"
        @input="onInput"
        class="flex-1 bg-transparent px-3 py-2.5 text-sm placeholder-gray-600 outline-none sm:px-4 sm:py-3 sm:text-base"
        :aria-describedby="
          hasError ? `${id}-error` : hint ? `${id}-hint` : undefined
        "
        :aria-invalid="hasError ? 'true' : 'false'"
      />
      <button
        v-if="password"
        type="button"
        @click="togglePassword"
        :aria-label="show ? hidePasswordLabel : showPasswordLabel"
        class="flex cursor-pointer items-center justify-center px-2.5 text-gray-500 transition-colors duration-200 hover:text-gray-700 active:text-gray-800 sm:px-3"
      >
        <i
          :class="show ? 'pi pi-eye-slash' : 'pi pi-eye'"
          class="text-base leading-none sm:text-lg"
        ></i>
      </button>
      <slot v-if="!password" name="right" />
    </div>

    <p
      v-if="hasError"
      :id="`${id}-error`"
      class="mt-1 text-xs text-red-500 sm:mt-1.5"
    >
      {{ errorMessage }}
    </p>
    <p
      v-else-if="hint"
      :id="`${id}-hint`"
      class="mt-1 text-xs text-gray-500 sm:mt-1.5"
    >
      {{ hint }}
    </p>
  </div>
</template>
