<script setup lang="ts">
import { computed } from "vue";
import { Calendar } from "lucide-vue-next";
import type { OfferListItem } from "../../../types/responses/offers";
import type { EnergyType } from "../../../types/offer";
import ProgressBar from "../../shared/ProgressBar.vue";
import SubmarketBadge from "../../shared/SubmarketBadge.vue";

interface Props {
  offer: OfferListItem;
  actionButtonText?: string;
}

interface Emits {
  (e: "click"): void;
}

interface EnergyTypeConfig {
  gradient: string;
  badge: string;
  label: string;
}

const props = defineProps<Props>();
const emit = defineEmits<Emits>();

const ENERGY_TYPE_CONFIG: Record<EnergyType, EnergyTypeConfig> = {
  solar: {
    gradient: "from-yellow-400 to-orange-500",
    badge: "bg-yellow-100 text-yellow-700",
    label: "Solar",
  },
  eolic: {
    gradient: "from-sky-400 to-blue-500",
    badge: "bg-sky-100 text-sky-700",
    label: "Eólica",
  },
  hydroelectric: {
    gradient: "from-blue-500 to-indigo-600",
    badge: "bg-blue-100 text-blue-700",
    label: "Hidrelétrica",
  },
  geothermal: {
    gradient: "from-orange-500 to-red-600",
    badge: "bg-orange-100 text-orange-700",
    label: "Geotérmica",
  },
};

const DEFAULT_CONFIG: EnergyTypeConfig = {
  gradient: "from-gray-400 to-gray-600",
  badge: "bg-gray-100 text-gray-700",
  label: "",
};

const formatPrice = (price: number): string => {
  return `R$ ${price.toFixed(2).replace(".", ",")}`;
};

const formatQuantity = (quantityMwh: number): string => {
  return `${quantityMwh.toLocaleString("pt-BR")} MWh`;
};

const formatDate = (dateString: string): string => {
  const [year, month, day] = dateString.split("-");
  return `${day}/${month}/${year}`;
};

const normalizeEnergyType = (type: string): EnergyType | null => {
  const normalized = type?.toLowerCase() || "";
  return normalized in ENERGY_TYPE_CONFIG ? (normalized as EnergyType) : null;
};

const getEnergyTypeConfig = (type: string): EnergyTypeConfig => {
  const normalizedType = normalizeEnergyType(type);
  return normalizedType ? ENERGY_TYPE_CONFIG[normalizedType] : DEFAULT_CONFIG;
};

const displayData = computed(() => ({
  description: props.offer.description,
  price: formatPrice(props.offer.price_per_mwh),
  quantity: formatQuantity(props.offer.remaining_quantity_mwh),
  initialQuantity: formatQuantity(props.offer.initial_quantity_mwh),
  deadline: formatDate(props.offer.period_end),
  location: props.offer.submarket,
  type: props.offer.energy_type,
}));

const typeConfig = computed(() => getEnergyTypeConfig(displayData.value.type));

const handleActionClick = () => {
  emit("click");
};
</script>

<template>
  <div
    class="group relative overflow-hidden rounded-lg border border-neutral-200 bg-white shadow-sm transition-all duration-300 hover:shadow-lg sm:rounded-xl"
  >
    <div
      :class="['h-1.5 w-full bg-gradient-to-r sm:h-2', typeConfig.gradient]"
    ></div>

    <div class="p-4 sm:p-6">
      <div class="flex items-start justify-between gap-3 sm:gap-4">
        <div class="min-w-0 flex-1">
          <div class="flex items-baseline gap-1">
            <span class="text-2xl font-bold text-neutral-900 sm:text-3xl">
              {{ displayData.price }}
            </span>
            <span class="text-xs text-neutral-500 sm:text-sm">/MWh</span>
          </div>
          <p class="mt-1 text-xs font-medium text-neutral-600 sm:text-sm">
            Disponível:
            <span class="font-semibold">{{ displayData.quantity }}</span> de
            <span class="font-semibold">{{ displayData.initialQuantity }}</span>
          </p>
        </div>
        <span
          :class="[
            'shrink-0 rounded-full px-2.5 py-1 text-xs font-medium whitespace-nowrap sm:px-3',
            typeConfig.badge,
          ]"
        >
          {{ typeConfig.label || displayData.type }}
        </span>
      </div>

      <div class="mt-3 sm:mt-4">
        <ProgressBar
          :current-value="offer.remaining_quantity_mwh"
          :total-value="offer.initial_quantity_mwh"
          :show-percentage="true"
          height="sm"
          gradient-from="emerald-500"
          gradient-to="green-500"
        />
      </div>

      <div
        class="mt-4 flex flex-col gap-2.5 border-t border-neutral-100 pt-3 sm:mt-6 sm:gap-3 sm:pt-4"
      >
        <div class="flex items-center gap-2">
          <span class="shrink-0 text-xs text-neutral-500 sm:text-sm"
            >Submercado:</span
          >
          <SubmarketBadge
            :offer-submarket="displayData.location"
            size="small"
          />
        </div>
        <div
          class="flex items-center gap-2 text-xs text-neutral-600 sm:text-sm"
        >
          <Calendar :size="14" class="shrink-0 text-neutral-400 sm:hidden" />
          <Calendar
            :size="16"
            class="hidden shrink-0 text-neutral-400 sm:block"
          />
          <span>Válido até {{ displayData.deadline }}</span>
        </div>
      </div>

      <div class="mt-4 flex flex-col gap-2 sm:mt-6">
        <button
          @click="handleActionClick"
          class="min-h-[44px] w-full touch-manipulation rounded-lg border-2 border-emerald-500 bg-transparent py-2.5 text-sm font-medium text-emerald-600 transition-all duration-200 hover:bg-emerald-500 hover:text-white active:scale-[0.98] sm:py-2.5"
        >
          {{ actionButtonText || "Ver Detalhes" }}
        </button>

        <slot name="actions"></slot>
      </div>
    </div>
  </div>
</template>
