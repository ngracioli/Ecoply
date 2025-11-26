<script setup lang="ts">
import EnergyOfferCard from "../dashboard/EnergyOfferCard.vue";
import type { OfferListItem } from "../../types/responses/offers";

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
  <section id="offer-demonstration" class="bg-white px-8 py-24">
    <div class="mx-auto max-w-7xl">
      <div class="mb-20 text-center">
        <div
          class="mb-4 inline-flex items-center gap-2 rounded-full bg-emerald-50 px-4 py-2"
        >
          <div
            class="flex h-6 w-6 items-center justify-center rounded-full bg-emerald-500"
          >
            <i class="pi pi-shopping-cart text-xs text-white"></i>
          </div>
          <span class="text-sm font-medium text-emerald-700">
            Marketplace
          </span>
        </div>
        <h2 class="mb-4 text-5xl font-bold tracking-tight text-gray-900">
          Marketplace de Energia
        </h2>
        <p class="mx-auto max-w-2xl text-lg text-gray-600">
          Explore ofertas reais de energia renovável com preços transparentes e
          transações seguras
        </p>
      </div>

      <div class="mb-12 grid gap-8 md:grid-cols-2 lg:grid-cols-3">
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
        <p class="mb-6 text-sm text-gray-600">
          Mais de 234 ofertas disponíveis na plataforma
        </p>
        <button
          @click="scrollToDashboard"
          class="group inline-flex items-center gap-3 rounded-full bg-gradient-to-r from-emerald-600 to-emerald-700 px-8 py-4 text-base font-semibold text-white shadow-lg shadow-emerald-500/30 transition-all duration-300 hover:-translate-y-1 hover:shadow-xl hover:shadow-emerald-500/40"
        >
          Explorar Todas as Ofertas
          <i
            class="pi pi-arrow-right text-sm transition-transform duration-300 group-hover:translate-x-1"
          ></i>
        </button>
      </div>
    </div>
  </section>
</template>

<style scoped></style>
