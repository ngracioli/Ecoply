<script setup lang="ts">
import { ref, onMounted, computed } from "vue";
import { useRoute, useRouter } from "vue-router";
import axios from "../axios";
import type { OfferListItem } from "../types/responses/offers";
import {
  Calendar,
  MapPin,
  TrendingUp,
  Package,
  Clock,
  ArrowLeft,
  User,
} from "lucide-vue-next";

const route = useRoute();
const router = useRouter();
const offer = ref<OfferListItem | null>(null);
const loading = ref(true);
const error = ref<string | null>(null);

const offerId = computed(() => route.params.id as string);

const fetchOffer = async () => {
  try {
    loading.value = true;
    error.value = null;
    const response = await axios.get(`/api/v1/offers/${offerId.value}`);
    offer.value = response.data.data;
  } catch (err: any) {
    error.value =
      err.response?.data?.message || "Erro ao carregar os detalhes da oferta";
    console.error("Erro ao buscar oferta:", err);
  } finally {
    loading.value = false;
  }
};

onMounted(() => {
  fetchOffer();
});

const goBack = () => {
  router.back();
};

const getTypeColor = (type: string) => {
  const colors = {
    solar: "from-amber-500 to-orange-500",
    eolic: "from-cyan-500 to-blue-500",
    hydroelectric: "from-blue-600 to-indigo-600",
    geothermal: "from-red-500 to-orange-600",
  };
  return colors[type as keyof typeof colors] || colors.solar;
};

const getTypeBadgeColor = (type: string) => {
  const colors = {
    solar: "bg-amber-100 text-amber-700 border-amber-300",
    eolic: "bg-cyan-100 text-cyan-700 border-cyan-300",
    hydroelectric: "bg-blue-100 text-blue-700 border-blue-300",
    geothermal: "bg-red-100 text-red-700 border-red-300",
  };
  return colors[type as keyof typeof colors] || colors.solar;
};

const getTypeLabel = (type: string) => {
  const labels = {
    solar: "Energia Solar",
    eolic: "Energia Eólica",
    hydroelectric: "Energia Hidrelétrica",
    geothermal: "Energia Geotérmica",
  };
  return labels[type as keyof typeof labels] || "Energia Solar";
};

const formatPrice = (price: number) => {
  return `R$ ${price.toFixed(2).replace(".", ",")}`;
};

const formatQuantity = (mwh: number) => {
  const kwh = (mwh * 1000).toFixed(0);
  return `${parseInt(kwh).toLocaleString("pt-BR")} kWh`;
};

const formatDate = (dateString: string) => {
  return new Date(dateString).toLocaleDateString("pt-BR", {
    day: "2-digit",
    month: "long",
    year: "numeric",
  });
};

const formatDateTime = (dateString: string) => {
  return new Date(dateString).toLocaleString("pt-BR", {
    day: "2-digit",
    month: "2-digit",
    year: "numeric",
    hour: "2-digit",
    minute: "2-digit",
  });
};

const getStatusBadge = (status: string) => {
  const statusConfig = {
    active: {
      label: "Ativa",
      class: "bg-green-100 text-green-700 border-green-300",
    },
    inactive: {
      label: "Inativa",
      class: "bg-gray-100 text-gray-700 border-gray-300",
    },
    expired: {
      label: "Expirada",
      class: "bg-red-100 text-red-700 border-red-300",
    },
  };
  return (
    statusConfig[status as keyof typeof statusConfig] || statusConfig.active
  );
};

const calculatePercentageRemaining = () => {
  if (!offer.value) return 0;
  return (
    (offer.value.remaining_quantity_mwh / offer.value.initial_quantity_mwh) *
    100
  ).toFixed(0);
};
</script>

<template>
  <div class="min-h-screen bg-gradient-to-br from-neutral-50 to-neutral-100">
    <!-- Header -->
    <header class="border-b border-neutral-200 bg-white shadow-sm">
      <div class="mx-auto max-w-7xl px-4 py-6 sm:px-6 lg:px-8">
        <div class="flex items-center justify-between">
          <div class="flex items-center gap-4">
            <button
              @click="goBack"
              class="rounded-lg p-2 text-neutral-600 transition-colors hover:bg-neutral-100 hover:text-neutral-900"
            >
              <ArrowLeft :size="24" />
            </button>
            <h1 class="text-3xl font-bold text-neutral-900">
              Detalhes da Oferta
            </h1>
          </div>
        </div>
      </div>
    </header>

    <!-- Content -->
    <main class="mx-auto max-w-7xl px-4 py-8 sm:px-6 lg:px-8">
      <!-- Loading State -->
      <div v-if="loading" class="flex items-center justify-center py-12">
        <div class="text-center">
          <div
            class="mx-auto h-12 w-12 animate-spin rounded-full border-4 border-neutral-200 border-t-emerald-500"
          ></div>
          <p class="mt-4 text-neutral-600">Carregando detalhes...</p>
        </div>
      </div>

      <!-- Error State -->
      <div
        v-else-if="error"
        class="rounded-lg border border-red-200 bg-red-50 p-6 text-center"
      >
        <p class="text-red-700">{{ error }}</p>
        <button
          @click="fetchOffer"
          class="mt-4 rounded-lg bg-red-600 px-4 py-2 text-white transition-colors hover:bg-red-700"
        >
          Tentar novamente
        </button>
      </div>

      <!-- Offer Details -->
      <div v-else-if="offer" class="space-y-6">
        <!-- Main Card -->
        <div
          class="overflow-hidden rounded-xl border border-neutral-200 bg-white shadow-lg"
        >
          <!-- Gradient Header -->
          <div
            :class="[
              'h-3 w-full bg-gradient-to-r',
              getTypeColor(offer.energy_type),
            ]"
          ></div>

          <div class="p-8">
            <!-- Header Info -->
            <div class="mb-6 flex flex-wrap items-start justify-between gap-4">
              <div class="flex-1">
                <div class="mb-2 flex flex-wrap items-center gap-3">
                  <span
                    :class="[
                      'rounded-full border px-4 py-1.5 text-sm font-medium capitalize',
                      getTypeBadgeColor(offer.energy_type),
                    ]"
                  >
                    {{ getTypeLabel(offer.energy_type) }}
                  </span>
                  <span
                    :class="[
                      'rounded-full border px-4 py-1.5 text-sm font-medium',
                      getStatusBadge(offer.status).class,
                    ]"
                  >
                    {{ getStatusBadge(offer.status).label }}
                  </span>
                </div>
                <h2 class="text-2xl font-bold text-neutral-900">
                  {{ offer.description }}
                </h2>
              </div>
              <div class="text-right">
                <div class="flex items-baseline gap-1">
                  <span class="text-4xl font-bold text-neutral-900">
                    {{ formatPrice(offer.price_per_mwh) }}
                  </span>
                  <span class="text-lg text-neutral-500">/MWh</span>
                </div>
              </div>
            </div>

            <!-- Stats Grid -->
            <div class="grid gap-6 md:grid-cols-2 lg:grid-cols-3">
              <!-- Quantidade Disponível -->
              <div
                class="rounded-lg border border-neutral-200 bg-neutral-50 p-6"
              >
                <div class="mb-3 flex items-center gap-3">
                  <div class="rounded-lg bg-emerald-100 p-2">
                    <Package class="text-emerald-600" :size="24" />
                  </div>
                  <h3 class="font-semibold text-neutral-700">
                    Quantidade Disponível
                  </h3>
                </div>
                <p class="text-2xl font-bold text-neutral-900">
                  {{ formatQuantity(offer.remaining_quantity_mwh) }}
                </p>
                <div class="mt-3">
                  <div
                    class="flex items-center justify-between text-sm text-neutral-600"
                  >
                    <span>Disponível</span>
                    <span>{{ calculatePercentageRemaining() }}%</span>
                  </div>
                  <div
                    class="mt-2 h-2 overflow-hidden rounded-full bg-neutral-200"
                  >
                    <div
                      :style="{ width: `${calculatePercentageRemaining()}%` }"
                      class="h-full bg-gradient-to-r from-emerald-500 to-green-500 transition-all"
                    ></div>
                  </div>
                </div>
              </div>

              <!-- Quantidade Inicial -->
              <div
                class="rounded-lg border border-neutral-200 bg-neutral-50 p-6"
              >
                <div class="mb-3 flex items-center gap-3">
                  <div class="rounded-lg bg-blue-100 p-2">
                    <TrendingUp class="text-blue-600" :size="24" />
                  </div>
                  <h3 class="font-semibold text-neutral-700">
                    Quantidade Inicial
                  </h3>
                </div>
                <p class="text-2xl font-bold text-neutral-900">
                  {{ formatQuantity(offer.initial_quantity_mwh) }}
                </p>
                <p class="mt-2 text-sm text-neutral-600">
                  Total ofertado inicialmente
                </p>
              </div>

              <!-- Localização -->
              <div
                class="rounded-lg border border-neutral-200 bg-neutral-50 p-6"
              >
                <div class="mb-3 flex items-center gap-3">
                  <div class="rounded-lg bg-purple-100 p-2">
                    <MapPin class="text-purple-600" :size="24" />
                  </div>
                  <h3 class="font-semibold text-neutral-700">Submercado</h3>
                </div>
                <p class="text-2xl font-bold text-neutral-900">
                  {{ offer.submarket }}
                </p>
                <p class="mt-2 text-sm text-neutral-600">
                  Região de fornecimento
                </p>
              </div>

              <!-- Período de Início -->
              <div
                class="rounded-lg border border-neutral-200 bg-neutral-50 p-6"
              >
                <div class="mb-3 flex items-center gap-3">
                  <div class="rounded-lg bg-amber-100 p-2">
                    <Calendar class="text-amber-600" :size="24" />
                  </div>
                  <h3 class="font-semibold text-neutral-700">
                    Início do Período
                  </h3>
                </div>
                <p class="text-xl font-bold text-neutral-900">
                  {{ formatDate(offer.period_start) }}
                </p>
              </div>

              <!-- Período de Fim -->
              <div
                class="rounded-lg border border-neutral-200 bg-neutral-50 p-6"
              >
                <div class="mb-3 flex items-center gap-3">
                  <div class="rounded-lg bg-red-100 p-2">
                    <Calendar class="text-red-600" :size="24" />
                  </div>
                  <h3 class="font-semibold text-neutral-700">Fim do Período</h3>
                </div>
                <p class="text-xl font-bold text-neutral-900">
                  {{ formatDate(offer.period_end) }}
                </p>
              </div>

              <!-- Data de Criação -->
              <div
                class="rounded-lg border border-neutral-200 bg-neutral-50 p-6"
              >
                <div class="mb-3 flex items-center gap-3">
                  <div class="rounded-lg bg-indigo-100 p-2">
                    <Clock class="text-indigo-600" :size="24" />
                  </div>
                  <h3 class="font-semibold text-neutral-700">Criada em</h3>
                </div>
                <p class="text-lg font-bold text-neutral-900">
                  {{ formatDateTime(offer.created_at) }}
                </p>
              </div>
            </div>

            <!-- Additional Info -->
            <div
              class="mt-6 rounded-lg border border-neutral-200 bg-neutral-50 p-6"
            >
              <div class="mb-3 flex items-center gap-3">
                <div class="rounded-lg bg-neutral-200 p-2">
                  <User class="text-neutral-600" :size="24" />
                </div>
                <h3 class="font-semibold text-neutral-700">
                  Informações do Vendedor
                </h3>
              </div>
              <p class="text-sm text-neutral-600">
                ID do Agente:
                <span class="font-mono text-neutral-900">{{
                  offer.seller_agent_uuid
                }}</span>
              </p>
              <p class="mt-2 text-sm text-neutral-600">
                ID da Oferta:
                <span class="font-mono text-neutral-900">{{ offer.uuid }}</span>
              </p>
            </div>

            <!-- Action Buttons -->
            <div class="mt-8 flex flex-wrap gap-4">
              <button
                class="flex-1 rounded-lg bg-gradient-to-r from-emerald-500 to-green-500 px-6 py-3 font-semibold text-white shadow-lg transition-all hover:shadow-xl"
              >
                Comprar Energia
              </button>
              <button
                class="rounded-lg border border-neutral-300 bg-white px-6 py-3 font-semibold text-neutral-700 transition-colors hover:bg-neutral-50"
              >
                Entrar em Contato
              </button>
            </div>
          </div>
        </div>
      </div>
    </main>
  </div>
</template>
