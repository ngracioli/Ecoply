<script setup lang="ts">
import { computed } from "vue";
import EnergyOfferCard from "../dashboard/cards/EnergyOfferCard.vue";
import type { OfferListItem } from "../../types/responses/offers";
import type { PlatformAnalytics } from "../../types/analytics";

interface Props {
  platformData: PlatformAnalytics;
  loading: boolean;
}

const props = defineProps<Props>();

const formatNumber = (num: number): string => {
  if (num >= 1000) {
    return (num / 1000).toFixed(1) + "k";
  }
  return num.toString();
};

const offersText = computed(() => {
  return props.loading ? "..." : formatNumber(props.platformData.active_offers);
});

const mockOffers: OfferListItem[] = [
  {
    uuid: "demo-1",
    description:
      "Energia solar de alta qualidade proveniente de painéis fotovoltaicos de última geração.",
    price_per_mwh: 245.0,
    initial_quantity_mwh: 20.0,
    remaining_quantity_mwh: 15.5,
    period_start: "2025-01-01",
    period_end: "2025-12-31",
    status: "active",
    submarket: "Sudeste",
    energy_type: "solar",
    seller_agent_uuid: "seller-1",
    created_at: "2025-01-01T00:00:00Z",
  },
  {
    uuid: "demo-2",
    description:
      "Energia eólica limpa e renovável com certificação de origem sustentável.",
    price_per_mwh: 189.5,
    initial_quantity_mwh: 35.0,
    remaining_quantity_mwh: 28.3,
    period_start: "2025-01-01",
    period_end: "2025-11-30",
    status: "active",
    submarket: "Sul",
    energy_type: "eolic",
    seller_agent_uuid: "seller-2",
    created_at: "2025-01-01T00:00:00Z",
  },
  {
    uuid: "demo-3",
    description:
      "Energia hidrelétrica proveniente de pequenas centrais com mínimo impacto ambiental.",
    price_per_mwh: 198.75,
    initial_quantity_mwh: 50.0,
    remaining_quantity_mwh: 42.8,
    period_start: "2025-01-01",
    period_end: "2025-12-15",
    status: "active",
    submarket: "Nordeste",
    energy_type: "hydroelectric",
    seller_agent_uuid: "seller-3",
    created_at: "2025-01-01T00:00:00Z",
  },
];

const scrollToDashboard = () => {
  window.scrollTo({ top: 0, behavior: "smooth" });
};
</script>

<template>
  <section
    id="offer-demonstration"
    class="bg-white px-4 py-12 sm:px-6 sm:py-16 md:px-8 lg:py-24"
  >
    <div class="mx-auto max-w-7xl">
      <div class="mb-12 text-center sm:mb-16 lg:mb-20">
        <div
          class="mb-3 inline-flex items-center gap-1.5 rounded-full bg-emerald-50 px-3 py-1.5 sm:mb-4 sm:gap-2 sm:px-4 sm:py-2"
        >
          <div
            class="flex h-5 w-5 items-center justify-center rounded-full bg-emerald-500 sm:h-6 sm:w-6"
          >
            <i
              class="pi pi-shopping-cart text-[10px] text-white sm:text-xs"
            ></i>
          </div>
          <span class="text-xs font-medium text-emerald-700 sm:text-sm">
            Marketplace
          </span>
        </div>
        <h2
          class="mb-3 text-3xl font-bold tracking-tight text-gray-900 sm:mb-4 sm:text-4xl lg:text-5xl"
        >
          Marketplace de Energia
        </h2>
        <p
          class="mx-auto max-w-2xl px-4 text-base text-gray-600 sm:px-0 sm:text-lg"
        >
          Explore ofertas reais de energia renovável com preços transparentes e
          transações seguras
        </p>
      </div>

      <div
        class="mb-8 grid gap-4 sm:mb-10 sm:gap-6 md:grid-cols-2 md:gap-8 lg:mb-12 lg:grid-cols-3"
      >
        <div
          v-for="offer in mockOffers"
          :key="offer.uuid"
          class="transition-all duration-300 hover:-translate-y-2"
        >
          <EnergyOfferCard
            :offer="offer"
            action-button-text="Ver Detalhes"
            @click="scrollToDashboard"
          />
        </div>
      </div>

      <div class="text-center">
        <p class="mb-4 text-xs text-gray-600 sm:mb-6 sm:text-sm">
          Mais de {{ offersText }} ofertas disponíveis na plataforma
        </p>
        <button
          @click="scrollToDashboard"
          class="group inline-flex items-center gap-2 rounded-full bg-gradient-to-r from-emerald-600 to-emerald-700 px-6 py-3 text-sm font-semibold text-white shadow-lg shadow-emerald-500/30 transition-all duration-300 hover:-translate-y-1 hover:shadow-xl hover:shadow-emerald-500/40 active:scale-95 sm:gap-3 sm:px-8 sm:py-4 sm:text-base"
        >
          <span class="hidden sm:inline">Explorar Todas as Ofertas</span>
          <span class="sm:hidden">Ver Ofertas</span>
          <i
            class="pi pi-arrow-right text-xs transition-transform duration-300 group-hover:translate-x-1 sm:text-sm"
          ></i>
        </button>
      </div>
    </div>
  </section>
</template>

<style scoped></style>
