<script setup lang="ts">
import { ref, watch } from "vue";

interface Props {
  modelValue: number | null;
  maxQuantity: number;
}

interface Emits {
  (e: "update:modelValue", value: number | null): void;
  (e: "validationChange", isValid: boolean): void;
}

const props = defineProps<Props>();
const emit = defineEmits<Emits>();

const hasError = ref(false);
const errorMessage = ref("");
const inputValue = ref<string>("");

watch(
  () => props.modelValue,
  (newValue) => {
    if (newValue === null || newValue === undefined) {
      inputValue.value = "";
    } else {
      inputValue.value = String(newValue);
    }
  },
  { immediate: true },
);

const validateAndEmit = (value: string) => {
  if (value === "" || value === null) {
    hasError.value = false;
    errorMessage.value = "";
    emit("update:modelValue", null);
    emit("validationChange", false);
    return;
  }

  const numValue = parseFloat(value);

  if (isNaN(numValue)) {
    hasError.value = true;
    errorMessage.value = "Digite um número válido";
    emit("validationChange", false);
    return;
  }

  if (numValue < 0.1) {
    hasError.value = true;
    errorMessage.value = "A quantidade mínima é 0.1 MWh";
    emit("validationChange", false);
    emit("update:modelValue", numValue);
    return;
  }

  if (numValue > props.maxQuantity) {
    hasError.value = true;
    errorMessage.value = `A quantidade máxima disponível é ${formatQuantity(props.maxQuantity)}`;
    emit("validationChange", false);
    emit("update:modelValue", numValue);
    return;
  }

  hasError.value = false;
  errorMessage.value = "";
  emit("update:modelValue", numValue);
  emit("validationChange", true);
};

const handleInput = (event: Event) => {
  const target = event.target as HTMLInputElement;
  inputValue.value = target.value;
  validateAndEmit(target.value);
};

const formatQuantity = (mwh: number) => {
  return `${mwh.toLocaleString("pt-BR")} MWh`;
};
</script>

<template>
  <div class="rounded-lg bg-white p-4 shadow-sm sm:rounded-xl sm:p-5 md:p-6">
    <h2
      class="mb-4 text-base font-semibold text-neutral-900 sm:mb-5 sm:text-lg md:mb-6"
    >
      Quantidade de Energia
    </h2>

    <div class="space-y-3 sm:space-y-4">
      <div>
        <label
          class="mb-1.5 block text-xs font-medium text-neutral-700 sm:mb-2 sm:text-sm"
        >
          Quantidade (MWh)
        </label>
        <input
          :value="inputValue"
          @input="handleInput"
          type="number"
          step="0.1"
          :max="maxQuantity"
          min="0.1"
          placeholder="Digite a quantidade desejada"
          :class="[
            'w-full min-w-0 touch-manipulation rounded-lg border px-3 py-2.5 text-base transition-colors focus:outline-none sm:px-4 sm:py-3 sm:text-lg',
            hasError
              ? 'border-red-500 focus:border-red-500 focus:ring-2 focus:ring-red-500/20'
              : 'border-neutral-300 focus:border-emerald-500 focus:ring-2 focus:ring-emerald-500/20',
          ]"
        />
        <p
          v-if="hasError"
          class="mt-1.5 text-xs font-medium break-words text-red-600 sm:mt-2 sm:text-sm"
        >
          {{ errorMessage }}
        </p>
        <p
          v-else
          class="mt-1.5 text-xs break-words text-neutral-500 sm:mt-2 sm:text-sm"
        >
          Disponível: {{ formatQuantity(maxQuantity) }} | Mínimo: 0.1 MWh
        </p>
      </div>
    </div>
  </div>
</template>

<style scoped>
input[type="number"]::-webkit-inner-spin-button,
input[type="number"]::-webkit-outer-spin-button {
  -webkit-appearance: none;
  margin: 0;
}

input[type="number"] {
  -moz-appearance: textfield;
  appearance: textfield;
}
</style>
