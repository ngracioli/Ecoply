<script setup lang="ts">
import { ref, onMounted, computed } from "vue";
import { useRoute, useRouter } from "vue-router";
import api from "../../axios";
import { OFFER_ENDPOINTS } from "../../api/endpoints";
import type {
  OfferListItem,
  OfferDetailResponse,
} from "../../types/responses/offers";
import {
  Calendar,
  MapPin,
  TrendingUp,
  Package,
  Clock,
  ArrowLeft,
  User,
} from "lucide-vue-next";
import ProgressBar from "../../components/shared/ProgressBar.vue";
import SubmarketBadge from "../../components/shared/SubmarketBadge.vue";

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
    const response = await api.get<OfferDetailResponse>(
      OFFER_ENDPOINTS.DETAIL(offerId.value),
    );
    offer.value = response.data.data;
  } catch (err: any) {
    error.value =
      err.response?.data?.message || "Erro ao carregar os detalhes da oferta";
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

const formatQuantity = (quantityMwh: number) => {
  return `${quantityMwh.toLocaleString("pt-BR")} MWh`;
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

const goToCheckout = () => {
  router.push({
    name: "Checkout",
    params: { id: offerId.value },
  });
};
</script>

<template>
  <div class="min-h-dvh bg-gradient-to-br from-neutral-50 to-neutral-100">
    <header class="border-b border-neutral-200 bg-white shadow-sm">
      <div class="mx-auto max-w-7xl px-3 py-4 sm:px-6 sm:py-6 lg:px-8">
        <div class="flex items-center justify-between">
          <div class="flex items-center gap-2 sm:gap-4">
            <button
              @click="goBack"
              class="rounded-lg p-1.5 text-neutral-600 transition-colors hover:bg-neutral-100 hover:text-neutral-900 sm:p-2"
            >
              <ArrowLeft :size="20" class="sm:hidden" />
              <ArrowLeft :size="24" class="hidden sm:block" />
            </button>
            <h1
              class="text-xl font-bold text-neutral-900 sm:text-2xl lg:text-3xl"
            >
              Detalhes da Oferta
            </h1>
          </div>
        </div>
      </div>
    </header>

    <main class="mx-auto max-w-7xl px-3 py-4 sm:px-6 sm:py-8 lg:px-8">
      <div
        v-if="loading"
        class="flex items-center justify-center py-8 sm:py-12"
      >
        <div class="text-center">
          <div
            class="mx-auto h-10 w-10 animate-spin rounded-full border-4 border-neutral-200 border-t-emerald-500 sm:h-12 sm:w-12"
          ></div>
          <p class="mt-3 text-sm text-neutral-600 sm:mt-4 sm:text-base">
            Carregando detalhes...
          </p>
        </div>
      </div>

      <div
        v-else-if="error"
        class="rounded-lg border border-red-200 bg-red-50 p-4 text-center sm:p-6"
      >
        <p class="text-sm text-red-700 sm:text-base">{{ error }}</p>
        <button
          @click="fetchOffer"
          class="mt-3 rounded-lg bg-red-600 px-4 py-2 text-sm text-white transition-colors hover:bg-red-700 sm:mt-4 sm:text-base"
        >
          Tentar novamente
        </button>
      </div>

      <div v-else-if="offer" class="space-y-4 sm:space-y-6">
        <div
          class="overflow-hidden rounded-lg border border-neutral-200 bg-white shadow-lg sm:rounded-xl"
        >
          <div
            :class="[
              'h-2 w-full bg-gradient-to-r sm:h-3',
              getTypeColor(offer.energy_type),
            ]"
          ></div>

          <div class="p-4 sm:p-6 lg:p-8">
            <div
              class="mb-4 flex flex-col gap-3 sm:mb-6 sm:flex-row sm:flex-wrap sm:items-start sm:justify-between sm:gap-4"
            >
              <div class="flex-1">
                <div
                  class="flex flex-col gap-2 sm:flex-row sm:flex-wrap sm:items-center sm:gap-3"
                >
                  <h2
                    class="text-lg font-bold text-neutral-900 sm:text-xl lg:text-2xl"
                  >
                    {{ offer.description }}
                  </h2>
                  <span
                    :class="[
                      'inline-flex w-fit rounded-full border px-3 py-1 text-xs font-medium capitalize sm:px-4 sm:py-1.5 sm:text-sm',
                      getTypeBadgeColor(offer.energy_type),
                    ]"
                  >
                    {{ getTypeLabel(offer.energy_type) }}
                  </span>
                </div>
              </div>
              <div class="text-left sm:text-right">
                <div class="flex items-baseline gap-1">
                  <span
                    class="text-2xl font-bold text-neutral-900 sm:text-3xl lg:text-4xl"
                  >
                    {{ formatPrice(offer.price_per_mwh) }}
                  </span>
                  <span class="text-base text-neutral-500 sm:text-lg"
                    >/MWh</span
                  >
                </div>
              </div>
            </div>

            <div
              class="grid gap-3 sm:gap-4 md:grid-cols-2 md:gap-6 lg:grid-cols-3"
            >
              <div
                class="rounded-lg border border-neutral-200 bg-neutral-50 p-4 sm:p-5 lg:p-6"
              >
                <div class="mb-2 flex items-center gap-2 sm:mb-3 sm:gap-3">
                  <div class="rounded-lg bg-emerald-100 p-1.5 sm:p-2">
                    <Package class="text-emerald-600" :size="20" />
                  </div>
                  <h3
                    class="text-sm font-semibold text-neutral-700 sm:text-base"
                  >
                    Quantidade Disponível
                  </h3>
                </div>
                <p class="text-xl font-bold text-neutral-900 sm:text-2xl">
                  {{ formatQuantity(offer.remaining_quantity_mwh) }}
                </p>
                <div class="mt-2 sm:mt-3">
                  <ProgressBar
                    :current-value="offer.remaining_quantity_mwh"
                    :total-value="offer.initial_quantity_mwh"
                    :show-percentage="true"
                    :show-label="true"
                    label="Disponível"
                    height="sm"
                    gradient-from="emerald-500"
                    gradient-to="green-500"
                  />
                </div>
              </div>

              <div
                class="rounded-lg border border-neutral-200 bg-neutral-50 p-4 sm:p-5 lg:p-6"
              >
                <div class="mb-2 flex items-center gap-2 sm:mb-3 sm:gap-3">
                  <div class="rounded-lg bg-blue-100 p-1.5 sm:p-2">
                    <TrendingUp class="text-blue-600" :size="20" />
                  </div>
                  <h3
                    class="text-sm font-semibold text-neutral-700 sm:text-base"
                  >
                    Quantidade Inicial
                  </h3>
                </div>
                <p class="text-xl font-bold text-neutral-900 sm:text-2xl">
                  {{ formatQuantity(offer.initial_quantity_mwh) }}
                </p>
                <p class="mt-1.5 text-xs text-neutral-600 sm:mt-2 sm:text-sm">
                  Total ofertado inicialmente
                </p>
              </div>

              <div
                class="rounded-lg border border-neutral-200 bg-neutral-50 p-4 sm:p-5 lg:p-6"
              >
                <div class="mb-2 flex items-center gap-2 sm:mb-3 sm:gap-3">
                  <div class="rounded-lg bg-purple-100 p-1.5 sm:p-2">
                    <MapPin class="text-purple-600" :size="20" />
                  </div>
                  <h3
                    class="text-sm font-semibold text-neutral-700 sm:text-base"
                  >
                    Submercado
                  </h3>
                </div>
                <SubmarketBadge
                  :offer-submarket="offer.submarket"
                  size="medium"
                />
                <p class="mt-1.5 text-xs text-neutral-600 sm:mt-2 sm:text-sm">
                  Região de fornecimento
                </p>
              </div>

              <div
                class="rounded-lg border border-neutral-200 bg-neutral-50 p-4 sm:p-5 lg:p-6"
              >
                <div class="mb-2 flex items-center gap-2 sm:mb-3 sm:gap-3">
                  <div class="rounded-lg bg-amber-100 p-1.5 sm:p-2">
                    <Calendar class="text-amber-600" :size="20" />
                  </div>
                  <h3
                    class="text-sm font-semibold text-neutral-700 sm:text-base"
                  >
                    Início do Período
                  </h3>
                </div>
                <p
                  class="text-base font-bold text-neutral-900 sm:text-lg lg:text-xl"
                >
                  {{ formatDate(offer.period_start) }}
                </p>
              </div>

              <div
                class="rounded-lg border border-neutral-200 bg-neutral-50 p-4 sm:p-5 lg:p-6"
              >
                <div class="mb-2 flex items-center gap-2 sm:mb-3 sm:gap-3">
                  <div class="rounded-lg bg-red-100 p-1.5 sm:p-2">
                    <Calendar class="text-red-600" :size="20" />
                  </div>
                  <h3
                    class="text-sm font-semibold text-neutral-700 sm:text-base"
                  >
                    Fim do Período
                  </h3>
                </div>
                <p
                  class="text-base font-bold text-neutral-900 sm:text-lg lg:text-xl"
                >
                  {{ formatDate(offer.period_end) }}
                </p>
              </div>

              <div
                class="rounded-lg border border-neutral-200 bg-neutral-50 p-4 sm:p-5 lg:p-6"
              >
                <div class="mb-2 flex items-center gap-2 sm:mb-3 sm:gap-3">
                  <div class="rounded-lg bg-indigo-100 p-1.5 sm:p-2">
                    <Clock class="text-indigo-600" :size="20" />
                  </div>
                  <h3
                    class="text-sm font-semibold text-neutral-700 sm:text-base"
                  >
                    Criada em
                  </h3>
                </div>
                <p class="text-base font-bold text-neutral-900 sm:text-lg">
                  {{ formatDateTime(offer.created_at) }}
                </p>
              </div>
            </div>

            <div
              class="mt-4 rounded-lg border border-neutral-200 bg-neutral-50 p-4 sm:mt-6 sm:p-5 lg:p-6"
            >
              <div class="mb-2 flex items-center gap-2 sm:mb-3 sm:gap-3">
                <div class="rounded-lg bg-neutral-200 p-1.5 sm:p-2">
                  <User class="text-neutral-600" :size="20" />
                </div>
                <h3 class="text-sm font-semibold text-neutral-700 sm:text-base">
                  Informações do Vendedor
                </h3>
              </div>
              <p class="text-xs break-all text-neutral-600 sm:text-sm">
                ID do Agente:
                <span class="font-mono text-neutral-900">{{
                  offer.seller_agent_uuid
                }}</span>
              </p>
              <p
                class="mt-1.5 text-xs break-all text-neutral-600 sm:mt-2 sm:text-sm"
              >
                ID da Oferta:
                <span class="font-mono text-neutral-900">{{ offer.uuid }}</span>
              </p>
            </div>

            <div class="mt-5 flex flex-wrap gap-3 sm:mt-8 sm:gap-4">
              <button
                @click="goToCheckout"
                class="w-full flex-1 rounded-lg bg-gradient-to-r from-emerald-500 to-green-500 px-5 py-2.5 text-sm font-semibold text-white shadow-lg transition-all hover:shadow-xl hover:brightness-110 sm:px-6 sm:py-3 sm:text-base"
              >
                Comprar Energia
              </button>
            </div>
          </div>
        </div>
      </div>
    </main>
  </div>
</template>
