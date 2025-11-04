<script setup lang="ts">
import { computed } from "vue";
import { useRouter } from "vue-router";
import { Calendar, Store } from "lucide-vue-next";
import type { OfferListItem } from "../../types/responses/offers";
import type { EnergyType } from "../../types/offer";

interface Props {
  offer: OfferListItem;
  actionButtonText?: string;
}

interface EnergyTypeConfig {
  gradient: string;
  badge: string;
  label: string;
}

const props = defineProps<Props>();
const router = useRouter();

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
  const quantityKwh = (quantityMwh * 1000).toFixed(0);
  return `${parseInt(quantityKwh).toLocaleString("pt-BR")} kWh`;
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
  deadline: formatDate(props.offer.period_end),
  location: props.offer.submarket,
  type: props.offer.energy_type,
}));

const typeConfig = computed(() => getEnergyTypeConfig(displayData.value.type));

const viewDetails = () => {
  router.push(`/offer/${props.offer.uuid}`);
};
</script>

<template>
  <div
    class="group relative overflow-hidden rounded-xl border border-neutral-200 bg-white shadow-sm transition-all duration-300 hover:shadow-lg"
  >
    <div :class="['h-2 w-full bg-gradient-to-r', typeConfig.gradient]"></div>

    <div class="p-6">
      <div class="flex items-start justify-between gap-4">
        <div class="flex-1">
          <div class="flex items-baseline gap-1">
            <span class="text-3xl font-bold text-neutral-900">
              {{ displayData.price }}
            </span>
            <span class="text-sm text-neutral-500">/MWh</span>
          </div>
          <p class="mt-1 text-sm font-medium text-neutral-600">
            {{ displayData.quantity }} disponíveis
          </p>
        </div>
        <span
          :class="[
            'rounded-full px-3 py-1 text-xs font-medium whitespace-nowrap',
            typeConfig.badge,
          ]"
        >
          {{ typeConfig.label || displayData.type }}
        </span>
      </div>

      <p class="mt-4 text-sm leading-relaxed text-neutral-600">
        {{ displayData.description }}
      </p>

      <div class="mt-6 flex flex-col gap-3 border-t border-neutral-100 pt-4">
        <div class="flex items-center gap-2 text-sm text-neutral-600">
          <Store :size="16" class="text-neutral-400" />
          <span>{{ displayData.location }}</span>
        </div>
        <div class="flex items-center gap-2 text-sm text-neutral-600">
          <Calendar :size="16" class="text-neutral-400" />
          <span>Válido até {{ displayData.deadline }}</span>
        </div>
      </div>

      <div class="mt-6 flex flex-col gap-2">
        <button
          @click="viewDetails"
          class="w-full rounded-lg border-2 border-emerald-500 bg-transparent py-2.5 text-sm font-medium text-emerald-600 transition-all duration-200 hover:bg-emerald-500 hover:text-white"
        >
          {{ actionButtonText || "Ver Detalhes" }}
        </button>

        <slot name="actions"></slot>
      </div>
    </div>
  </div>
</template>
