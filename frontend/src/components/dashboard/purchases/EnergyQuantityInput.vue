<script setup lang="ts">
import { computed } from "vue";

interface Props {
  modelValue: number | null;
  maxQuantity: number;
}

interface Emits {
  (e: "update:modelValue", value: number | null): void;
}

const props = defineProps<Props>();
const emit = defineEmits<Emits>();

const quantity = computed({
  get: () => props.modelValue,
  set: (value) => {
    if (value && value > props.maxQuantity) {
      emit("update:modelValue", props.maxQuantity);
    } else {
      emit("update:modelValue", value);
    }
  },
});

const formatQuantity = (mwh: number) => {
  return `${mwh.toLocaleString("pt-BR")} MWh`;
};
</script>

<template>
  <div class="rounded-xl bg-white p-6 shadow-sm">
    <h2 class="mb-6 text-lg font-semibold text-neutral-900">
      Quantidade de Energia
    </h2>

    <div class="space-y-4">
      <div>
        <label class="mb-2 block text-sm font-medium text-neutral-700">
          Quantidade (MWh)
        </label>
        <input
          v-model.number="quantity"
          type="number"
          step="0.1"
          :max="maxQuantity"
          min="0.1"
          placeholder="Digite a quantidade desejada"
          class="w-full rounded-lg border border-neutral-300 px-4 py-3 text-lg focus:border-emerald-500 focus:ring-2 focus:ring-emerald-500/20 focus:outline-none"
        />
        <p class="mt-2 text-sm text-neutral-500">
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
