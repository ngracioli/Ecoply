<script setup lang="ts">
import { ref } from "vue";
import { X, Filter } from "lucide-vue-next";

interface FilterOptions {
  submarket?: string;
  energy_type?: string;
  period_start?: string;
  period_end?: string;
}

interface Props {
  visible: boolean;
}

defineProps<Props>();

const emit = defineEmits<{
  (e: "update:visible", value: boolean): void;
  (e: "apply-filters", filters: FilterOptions): void;
}>();

// Estados dos filtros
const selectedSubmarket = ref<string>("");
const selectedEnergyType = ref<string>("");
const periodStart = ref<string>("");
const periodEnd = ref<string>("");

// Opções de submercado
const submarketOptions = [
  { value: "", label: "Todos" },
  { value: "SE_CO", label: "Sudeste/Centro-Oeste (SE/CO)" },
  { value: "S", label: "Sul (S)" },
  { value: "NE", label: "Nordeste (NE)" },
  { value: "N", label: "Norte (N)" },
];

// Opções de tipo de energia
const energyTypeOptions = [
  { value: "", label: "Todos" },
  { value: "solar", label: "Solar" },
  { value: "eolic", label: "Eólica" },
  { value: "hydroelectric", label: "Hidrelétrica" },
  { value: "geothermal", label: "Geotérmica" },
];

const closeDialog = () => {
  emit("update:visible", false);
};

const clearFilters = () => {
  selectedSubmarket.value = "";
  selectedEnergyType.value = "";
  periodStart.value = "";
  periodEnd.value = "";
};

const applyFilters = () => {
  const filters: FilterOptions = {};

  if (selectedSubmarket.value) {
    filters.submarket = selectedSubmarket.value;
  }

  if (selectedEnergyType.value) {
    filters.energy_type = selectedEnergyType.value;
  }

  if (periodStart.value) {
    filters.period_start = periodStart.value;
  }

  if (periodEnd.value) {
    filters.period_end = periodEnd.value;
  }

  emit("apply-filters", filters);
  closeDialog();
};

// Fecha o dialog ao clicar fora
const handleBackdropClick = (event: MouseEvent) => {
  if ((event.target as HTMLElement).classList.contains("dialog-backdrop")) {
    closeDialog();
  }
};

// Expõe a função clearFilters para o componente pai
defineExpose({
  clearFilters,
});
</script>

<template>
  <Transition
    name="dialog"
    enter-active-class="transition-opacity duration-300 ease-out"
    leave-active-class="transition-opacity duration-200 ease-in"
    enter-from-class="opacity-0"
    enter-to-class="opacity-100"
    leave-from-class="opacity-100"
    leave-to-class="opacity-0"
  >
    <div
      v-if="visible"
      class="dialog-backdrop fixed inset-0 z-50 flex items-center justify-center bg-black/50 p-4"
      @click="handleBackdropClick"
    >
      <Transition
        name="dialog-content"
        enter-active-class="transition-all duration-300 ease-out"
        leave-active-class="transition-all duration-200 ease-in"
        enter-from-class="opacity-0 scale-95 translate-y-4"
        enter-to-class="opacity-100 scale-100 translate-y-0"
        leave-from-class="opacity-100 scale-100 translate-y-0"
        leave-to-class="opacity-0 scale-95 translate-y-4"
      >
        <div
          v-if="visible"
          class="relative w-full max-w-2xl rounded-xl bg-white shadow-2xl"
          @click.stop
        >
          <!-- Header -->
          <div
            class="flex items-center justify-between border-b border-neutral-200 p-6"
          >
            <div class="flex items-center gap-3">
              <div class="rounded-lg bg-emerald-100 p-2">
                <Filter class="text-emerald-600" :size="24" />
              </div>
              <div>
                <h2 class="text-xl font-bold text-neutral-900">
                  Filtrar Ofertas
                </h2>
                <p class="text-sm text-neutral-500">
                  Refine sua busca por ofertas de energia
                </p>
              </div>
            </div>
            <button
              @click="closeDialog"
              class="rounded-lg p-2 text-neutral-400 transition-colors hover:bg-neutral-100 hover:text-neutral-600"
            >
              <X :size="24" />
            </button>
          </div>

          <!-- Content -->
          <div class="space-y-6 p-6">
            <!-- Submercado -->
            <div class="space-y-2">
              <label
                for="submarket"
                class="block text-sm font-medium text-neutral-700"
              >
                Submercado
              </label>
              <select
                id="submarket"
                v-model="selectedSubmarket"
                class="w-full rounded-lg border border-neutral-300 bg-white px-4 py-2.5 text-sm text-neutral-900 transition-colors focus:border-emerald-500 focus:ring-2 focus:ring-emerald-500/20 focus:outline-none"
              >
                <option
                  v-for="option in submarketOptions"
                  :key="option.value"
                  :value="option.value"
                >
                  {{ option.label }}
                </option>
              </select>
            </div>

            <!-- Tipo de Energia -->
            <div class="space-y-2">
              <label
                for="energy-type"
                class="block text-sm font-medium text-neutral-700"
              >
                Tipo de Energia
              </label>
              <select
                id="energy-type"
                v-model="selectedEnergyType"
                class="w-full rounded-lg border border-neutral-300 bg-white px-4 py-2.5 text-sm text-neutral-900 transition-colors focus:border-emerald-500 focus:ring-2 focus:ring-emerald-500/20 focus:outline-none"
              >
                <option
                  v-for="option in energyTypeOptions"
                  :key="option.value"
                  :value="option.value"
                >
                  {{ option.label }}
                </option>
              </select>
            </div>

            <!-- Período -->
            <div class="space-y-4">
              <h3 class="text-sm font-medium text-neutral-700">
                Período de Fornecimento
              </h3>
              <div class="grid grid-cols-2 gap-4">
                <!-- Data Início -->
                <div class="space-y-2">
                  <label
                    for="period-start"
                    class="block text-sm font-medium text-neutral-600"
                  >
                    Data Início
                  </label>
                  <input
                    id="period-start"
                    v-model="periodStart"
                    type="date"
                    class="w-full rounded-lg border border-neutral-300 bg-white px-4 py-2.5 text-sm text-neutral-900 transition-colors focus:border-emerald-500 focus:ring-2 focus:ring-emerald-500/20 focus:outline-none"
                  />
                </div>

                <!-- Data Fim -->
                <div class="space-y-2">
                  <label
                    for="period-end"
                    class="block text-sm font-medium text-neutral-600"
                  >
                    Data Fim
                  </label>
                  <input
                    id="period-end"
                    v-model="periodEnd"
                    type="date"
                    class="w-full rounded-lg border border-neutral-300 bg-white px-4 py-2.5 text-sm text-neutral-900 transition-colors focus:border-emerald-500 focus:ring-2 focus:ring-emerald-500/20 focus:outline-none"
                  />
                </div>
              </div>
            </div>
          </div>

          <!-- Footer -->
          <div
            class="flex items-center justify-between border-t border-neutral-200 p-6"
          >
            <button
              @click="clearFilters"
              class="rounded-lg border border-neutral-300 bg-white px-6 py-2.5 text-sm font-medium text-neutral-700 transition-colors hover:bg-neutral-50"
            >
              Limpar Filtros
            </button>
            <div class="flex gap-3">
              <button
                @click="closeDialog"
                class="rounded-lg border border-neutral-300 bg-white px-6 py-2.5 text-sm font-medium text-neutral-700 transition-colors hover:bg-neutral-50"
              >
                Cancelar
              </button>
              <button
                @click="applyFilters"
                class="rounded-lg bg-gradient-to-r from-emerald-500 to-teal-600 px-6 py-2.5 text-sm font-medium text-white shadow-md transition-all hover:shadow-lg hover:brightness-110"
              >
                Aplicar Filtros
              </button>
            </div>
          </div>
        </div>
      </Transition>
    </div>
  </Transition>
</template>

<style scoped>
/* Animações customizadas para melhor UX */
.dialog-enter-active,
.dialog-leave-active {
  transition: opacity 0.3s ease;
}

.dialog-enter-from,
.dialog-leave-to {
  opacity: 0;
}

.dialog-content-enter-active,
.dialog-content-leave-active {
  transition: all 0.3s ease;
}

.dialog-content-enter-from,
.dialog-content-leave-to {
  opacity: 0;
  transform: scale(0.95) translateY(1rem);
}
</style>
