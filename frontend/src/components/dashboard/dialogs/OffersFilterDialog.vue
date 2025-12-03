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

const selectedSubmarket = ref<string>("");
const selectedEnergyType = ref<string>("");
const periodStart = ref<string>("");
const periodEnd = ref<string>("");

const submarketOptions = [
  { value: "", label: "Todos" },
  { value: "SE_CO", label: "Sudeste/Centro-Oeste (SE/CO)" },
  { value: "S", label: "Sul (S)" },
  { value: "NE", label: "Nordeste (NE)" },
  { value: "N", label: "Norte (N)" },
];

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

const handleBackdropClick = (event: MouseEvent) => {
  if ((event.target as HTMLElement).classList.contains("dialog-backdrop")) {
    closeDialog();
  }
};

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
          class="mobile-scrollable relative w-full max-w-2xl rounded-lg bg-white shadow-2xl sm:rounded-xl"
          @click.stop
        >
          <div
            class="flex items-center justify-between gap-3 border-b border-neutral-200 p-4 sm:p-6"
          >
            <div class="flex min-w-0 flex-1 items-center gap-2.5 sm:gap-3">
              <div class="shrink-0 rounded-lg bg-emerald-100 p-1.5 sm:p-2">
                <Filter class="text-emerald-600" :size="20" />
                <Filter class="hidden text-emerald-600 sm:block" :size="24" />
              </div>
              <div class="min-w-0 flex-1">
                <h2
                  class="truncate text-lg font-bold text-neutral-900 sm:text-xl"
                >
                  Filtrar Ofertas
                </h2>
                <p class="truncate text-xs text-neutral-500 sm:text-sm">
                  Refine sua busca por ofertas de energia
                </p>
              </div>
            </div>
            <button
              @click="closeDialog"
              class="shrink-0 touch-manipulation rounded-lg p-1.5 text-neutral-400 transition-colors hover:bg-neutral-100 hover:text-neutral-600 active:scale-95 sm:p-2"
            >
              <X :size="20" class="sm:hidden" />
              <X :size="24" class="hidden sm:block" />
            </button>
          </div>

          <div
            class="mobile-scrollable max-h-[60vh] space-y-5 overflow-y-auto p-4 sm:space-y-6 sm:p-6"
          >
            <div class="space-y-2">
              <label
                for="submarket"
                class="block text-xs font-medium text-neutral-700 sm:text-sm"
              >
                Submercado
              </label>
              <select
                id="submarket"
                v-model="selectedSubmarket"
                class="w-full touch-manipulation rounded-lg border border-neutral-300 bg-white px-3 py-2.5 text-sm text-neutral-900 transition-colors focus:border-emerald-500 focus:ring-2 focus:ring-emerald-500/20 focus:outline-none sm:px-4 sm:text-base"
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

            <div class="space-y-2">
              <label
                for="energy-type"
                class="block text-xs font-medium text-neutral-700 sm:text-sm"
              >
                Tipo de Energia
              </label>
              <select
                id="energy-type"
                v-model="selectedEnergyType"
                class="w-full touch-manipulation rounded-lg border border-neutral-300 bg-white px-3 py-2.5 text-sm text-neutral-900 transition-colors focus:border-emerald-500 focus:ring-2 focus:ring-emerald-500/20 focus:outline-none sm:px-4 sm:text-base"
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

            <div class="space-y-3 sm:space-y-4">
              <h3 class="text-xs font-medium text-neutral-700 sm:text-sm">
                Período de Fornecimento
              </h3>
              <div class="grid grid-cols-1 gap-3 sm:grid-cols-2 sm:gap-4">
                <div class="min-w-0 space-y-2">
                  <label
                    for="period-start"
                    class="block text-xs font-medium text-neutral-600 sm:text-sm"
                  >
                    Data Início
                  </label>
                  <div class="relative overflow-hidden">
                    <input
                      id="period-start"
                      v-model="periodStart"
                      type="date"
                      class="box-border w-full max-w-full min-w-0 touch-manipulation rounded-lg border border-neutral-300 bg-white px-2 py-2.5 pr-8 text-sm text-neutral-900 transition-colors focus:border-emerald-500 focus:ring-2 focus:ring-emerald-500/20 focus:outline-none sm:px-4 sm:text-base"
                      style="font-size: 14px"
                    />
                  </div>
                </div>

                <div class="min-w-0 space-y-2">
                  <label
                    for="period-end"
                    class="block text-xs font-medium text-neutral-600 sm:text-sm"
                  >
                    Data Fim
                  </label>
                  <div class="relative overflow-hidden">
                    <input
                      id="period-end"
                      v-model="periodEnd"
                      type="date"
                      class="box-border w-full max-w-full min-w-0 touch-manipulation rounded-lg border border-neutral-300 bg-white px-2 py-2.5 pr-8 text-sm text-neutral-900 transition-colors focus:border-emerald-500 focus:ring-2 focus:ring-emerald-500/20 focus:outline-none sm:px-4 sm:text-base"
                      style="font-size: 14px"
                    />
                  </div>
                </div>
              </div>
            </div>
          </div>

          <div
            class="flex flex-col items-stretch justify-between gap-3 border-t border-neutral-200 p-4 sm:flex-row sm:items-center sm:p-6"
          >
            <button
              @click="clearFilters"
              class="order-2 min-h-[44px] touch-manipulation rounded-lg border border-neutral-300 bg-white px-5 py-2.5 text-sm font-medium text-neutral-700 transition-colors hover:bg-neutral-50 active:scale-[0.98] sm:order-1 sm:px-6"
            >
              Limpar Filtros
            </button>
            <div class="order-1 flex gap-2.5 sm:order-2 sm:gap-3">
              <button
                @click="closeDialog"
                class="min-h-[44px] flex-1 touch-manipulation rounded-lg border border-neutral-300 bg-white px-4 py-2.5 text-sm font-medium text-neutral-700 transition-colors hover:bg-neutral-50 active:scale-[0.98] sm:flex-initial sm:px-6"
              >
                Cancelar
              </button>
              <button
                @click="applyFilters"
                class="min-h-[44px] flex-1 touch-manipulation rounded-lg bg-gradient-to-r from-emerald-500 to-teal-600 px-4 py-2.5 text-sm font-medium whitespace-nowrap text-white shadow-md transition-all hover:shadow-lg hover:brightness-110 active:scale-[0.98] sm:flex-initial sm:px-6"
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

input[type="date"] {
  appearance: none;
  -webkit-appearance: none;
  -moz-appearance: none;
}

input[type="date"]::-webkit-calendar-picker-indicator {
  position: absolute;
  right: 8px;
  width: 20px;
  height: 20px;
  cursor: pointer;
  flex-shrink: 0;
}

input[type="date"]::-webkit-date-and-time-value {
  text-align: left;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  margin-right: 28px;
}

@media (max-width: 639px) {
  input[type="date"] {
    font-size: 14px !important;
    padding-right: 32px !important;
  }
}
</style>
