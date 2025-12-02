<script setup lang="ts">
import { ref, computed, onMounted } from "vue";
import { useStore } from "vuex";
import { useRouter } from "vue-router";
import {
  DollarSign,
  TrendingDown,
  TrendingUp,
  Zap,
  Clock,
  ShoppingCart,
  Package,
} from "lucide-vue-next";
import type { Component } from "vue";
import api from "../../../axios";
import { AUTH_ENDPOINTS } from "../../../api/endpoints";
import type { AnalyticsResponse } from "../../../types/analytics";
import type { RootState } from "../../../store/types";

interface MarketInsight {
  id: number;
  title: string;
  value: string;
  subtitle: string;
  trend: "up" | "down" | "neutral";
  icon: Component;
  colorClass: string;
  bgClass: string;
}

const store = useStore<RootState>();
const router = useRouter();

const loading = ref(true);
const analytics = ref<AnalyticsResponse["data"] | null>(null);

const isBuyer = computed(() => store.getters["user/isBuyer"]);
const isSupplier = computed(() => store.getters["user/isSupplier"]);

const formatCurrency = (value: number): string => {
  return new Intl.NumberFormat("pt-BR", {
    style: "currency",
    currency: "BRL",
  }).format(value);
};

const formatNumber = (value: number, decimals: number = 2): string => {
  return new Intl.NumberFormat("pt-BR", {
    minimumFractionDigits: decimals,
    maximumFractionDigits: decimals,
  }).format(value);
};

const goToOffer = (uuid: string) => {
  router.push(`/offer/${uuid}`);
};

const fetchAnalytics = async () => {
  try {
    loading.value = true;
    const response = await api.get<AnalyticsResponse>(AUTH_ENDPOINTS.ANALYTICS);
    analytics.value = response.data.data;
  } catch (error) {
  } finally {
    loading.value = false;
  }
};

const marketInsights = computed((): MarketInsight[] => {
  if (!analytics.value) return [];

  const insights: MarketInsight[] = [];

  if (isBuyer.value && analytics.value.buyer_info?.advantage_offer_uuid) {
    insights.push({
      id: 1,
      title: "Melhor Oferta Disponível",
      value: "Ver Oferta",
      subtitle: "Clique para visualizar a melhor oferta para você",
      trend: "down",
      icon: TrendingDown,
      colorClass: "text-emerald-600",
      bgClass: "bg-emerald-50",
    });
  }

  if (isSupplier.value && analytics.value.supplier_info) {
    const supplier = analytics.value.supplier_info;

    if (supplier.almost_expiring_offers > 0) {
      insights.push({
        id: 2,
        title: "Ofertas Expirando",
        value: `${supplier.almost_expiring_offers} ${supplier.almost_expiring_offers === 1 ? "oferta" : "ofertas"}`,
        subtitle: "Expiram nas próximas 48 horas",
        trend: "neutral",
        icon: Clock,
        colorClass: "text-amber-600",
        bgClass: "bg-amber-50",
      });
    }

    const priceDiff = supplier.user_price_avg - supplier.platform_price_avg;
    const priceDiffPercent = (
      (priceDiff / supplier.platform_price_avg) *
      100
    ).toFixed(1);
    const isAboveAverage = priceDiff > 0;

    insights.push({
      id: 3,
      title: "SEU PREÇO MÉDIO",
      value: formatCurrency(supplier.user_price_avg),
      subtitle: isAboveAverage
        ? `${priceDiffPercent}% acima da média da plataforma (${formatCurrency(supplier.platform_price_avg)})`
        : `${Math.abs(Number(priceDiffPercent))}% abaixo da média da plataforma (${formatCurrency(supplier.platform_price_avg)})`,
      trend: isAboveAverage ? "up" : "down",
      icon: isAboveAverage ? TrendingUp : TrendingDown,
      colorClass: isAboveAverage ? "text-red-600" : "text-emerald-600",
      bgClass: isAboveAverage ? "bg-red-50" : "bg-emerald-50",
    });
  }

  return insights;
});

onMounted(() => {
  fetchAnalytics();
});
</script>

<template>
  <div class="flex flex-col gap-4 sm:gap-6">
    <div v-if="loading" class="grid grid-cols-1 gap-4 sm:gap-6 lg:grid-cols-2">
      <div
        v-for="i in 4"
        :key="i"
        class="h-48 animate-pulse rounded-lg bg-neutral-200 sm:h-64 sm:rounded-xl"
      ></div>
    </div>

    <div
      v-else-if="isBuyer && analytics?.buyer_info"
      class="space-y-4 sm:space-y-6"
    >
      <div
        v-if="
          analytics.buyer_info.advantage_offer_uuid &&
          analytics.buyer_info.advantage_offer_uuid !== ''
        "
        class="cursor-pointer rounded-lg border-2 border-emerald-300 bg-gradient-to-br from-emerald-50 to-green-50 p-4 shadow-md transition-all duration-300 hover:scale-[1.02] hover:border-emerald-400 hover:shadow-lg active:scale-[0.98] sm:rounded-xl sm:p-6"
        @click="goToOffer(analytics.buyer_info.advantage_offer_uuid)"
      >
        <div class="flex items-start gap-3 sm:gap-4">
          <div class="min-w-0 flex-1">
            <div class="flex items-center gap-1.5 text-emerald-700 sm:gap-2">
              <TrendingDown
                :size="16"
                :stroke-width="2.5"
                class="shrink-0 sm:hidden"
              />
              <TrendingDown
                :size="18"
                :stroke-width="2.5"
                class="hidden shrink-0 sm:block"
              />
              <p
                class="truncate text-xs font-bold tracking-wide uppercase sm:text-sm"
              >
                🎯 Melhor Oferta Disponível
              </p>
            </div>
            <h3
              class="mt-2 text-lg font-bold text-emerald-900 sm:mt-3 sm:text-2xl"
            >
              Ver Oferta Vantajosa
            </h3>
            <p class="mt-1.5 text-xs text-emerald-700 sm:mt-2 sm:text-sm">
              Encontramos a melhor oferta do seu submercado. Clique para
              visualizar os detalhes e aproveitar!
            </p>
            <div
              class="mt-3 inline-flex min-h-[40px] items-center gap-1.5 rounded-lg bg-emerald-600 px-3 py-2 text-xs font-semibold text-white transition-colors hover:bg-emerald-700 active:bg-emerald-800 sm:mt-4 sm:min-h-0 sm:gap-2 sm:px-4 sm:py-2 sm:text-sm"
            >
              <span>Visualizar Oferta</span>
              <TrendingDown :size="14" :stroke-width="2.5" class="sm:hidden" />
              <TrendingDown
                :size="16"
                :stroke-width="2.5"
                class="hidden sm:block"
              />
            </div>
          </div>
          <div
            class="flex h-12 w-12 shrink-0 items-center justify-center rounded-lg bg-emerald-600 text-white shadow-lg sm:h-16 sm:w-16 sm:rounded-xl"
          >
            <TrendingDown :size="24" :stroke-width="2.5" class="sm:hidden" />
            <TrendingDown
              :size="32"
              :stroke-width="2.5"
              class="hidden sm:block"
            />
          </div>
        </div>
      </div>

      <div class="grid grid-cols-1 gap-4 sm:gap-6 md:grid-cols-2">
        <div
          class="rounded-lg border border-neutral-200 bg-white p-4 shadow-sm transition-all duration-300 hover:shadow-md sm:rounded-xl sm:p-6"
        >
          <div class="flex items-start justify-between">
            <div class="min-w-0 flex-1">
              <div class="flex items-center gap-1.5 text-neutral-500 sm:gap-2">
                <ShoppingCart
                  :size="14"
                  :stroke-width="2"
                  class="shrink-0 sm:hidden"
                />
                <ShoppingCart
                  :size="16"
                  :stroke-width="2"
                  class="hidden shrink-0 sm:block"
                />
                <p class="truncate text-xs font-medium tracking-wide uppercase">
                  Total de Compras
                </p>
              </div>
              <p
                class="mt-2 text-2xl font-bold text-neutral-900 sm:mt-3 sm:text-3xl"
              >
                {{ analytics.buyer_info.purchases_count }}
              </p>
              <p class="mt-1 text-xs text-neutral-500">
                {{
                  analytics.buyer_info.purchases_count === 1
                    ? "Negociação concluída"
                    : "Negociações concluídas"
                }}
              </p>
            </div>
            <div
              class="flex h-10 w-10 shrink-0 items-center justify-center rounded-lg bg-blue-50 text-blue-600 sm:h-12 sm:w-12"
            >
              <ShoppingCart :size="20" :stroke-width="2" class="sm:hidden" />
              <ShoppingCart
                :size="24"
                :stroke-width="2"
                class="hidden sm:block"
              />
            </div>
          </div>
        </div>

        <div
          class="rounded-lg border border-neutral-200 bg-white p-4 shadow-sm transition-all duration-300 hover:shadow-md sm:rounded-xl sm:p-6"
        >
          <div class="flex items-start justify-between">
            <div class="min-w-0 flex-1">
              <div class="flex items-center gap-1.5 text-neutral-500 sm:gap-2">
                <Zap :size="14" :stroke-width="2" class="shrink-0 sm:hidden" />
                <Zap
                  :size="16"
                  :stroke-width="2"
                  class="hidden shrink-0 sm:block"
                />
                <p class="truncate text-xs font-medium tracking-wide uppercase">
                  Energia Movimentada
                </p>
              </div>
              <div class="mt-2 flex items-baseline gap-1.5 sm:mt-3 sm:gap-2">
                <p class="text-2xl font-bold text-neutral-900 sm:text-3xl">
                  {{ formatNumber(analytics.buyer_info.energy_transacted, 0) }}
                </p>
                <p class="text-xs text-neutral-500 sm:text-sm">kWh</p>
              </div>
              <p class="mt-1.5 text-xs text-neutral-400 sm:mt-2">
                Média de
                {{
                  formatNumber(
                    analytics.buyer_info.energy_transacted /
                      analytics.buyer_info.purchases_count,
                    0,
                  )
                }}
                kWh por compra
              </p>
            </div>
            <div
              class="flex h-10 w-10 shrink-0 items-center justify-center rounded-lg bg-amber-50 text-amber-600 sm:h-12 sm:w-12"
            >
              <Zap :size="20" :stroke-width="2" class="sm:hidden" />
              <Zap :size="24" :stroke-width="2" class="hidden sm:block" />
            </div>
          </div>
        </div>
      </div>
    </div>

    <div
      v-else-if="isSupplier && analytics?.supplier_info"
      class="space-y-4 sm:space-y-6"
    >
      <div
        v-if="
          analytics.buyer_info?.advantage_offer_uuid &&
          analytics.buyer_info.advantage_offer_uuid !== ''
        "
        class="cursor-pointer rounded-lg border-2 border-emerald-300 bg-gradient-to-br from-emerald-50 to-green-50 p-4 shadow-md transition-all duration-300 hover:scale-[1.02] hover:border-emerald-400 hover:shadow-lg active:scale-[0.98] sm:rounded-xl sm:p-6"
        @click="goToOffer(analytics.buyer_info.advantage_offer_uuid)"
      >
        <div class="flex items-start gap-3 sm:gap-4">
          <div class="min-w-0 flex-1">
            <div class="flex items-center gap-1.5 text-emerald-700 sm:gap-2">
              <TrendingDown
                :size="16"
                :stroke-width="2.5"
                class="shrink-0 sm:hidden"
              />
              <TrendingDown
                :size="18"
                :stroke-width="2.5"
                class="hidden shrink-0 sm:block"
              />
              <p
                class="truncate text-xs font-bold tracking-wide uppercase sm:text-sm"
              >
                🎯 Melhor Oferta Disponível
              </p>
            </div>
            <h3
              class="mt-2 text-lg font-bold text-emerald-900 sm:mt-3 sm:text-2xl"
            >
              Oportunidade de Compra Vantajosa
            </h3>
            <p class="mt-1.5 text-xs text-emerald-700 sm:mt-2 sm:text-sm">
              Encontramos a melhor oferta do seu submercado. Clique para
              visualizar os detalhes e aproveitar!
            </p>
            <div
              class="mt-3 inline-flex min-h-[40px] items-center gap-1.5 rounded-lg bg-emerald-600 px-3 py-2 text-xs font-semibold text-white transition-colors hover:bg-emerald-700 active:bg-emerald-800 sm:mt-4 sm:min-h-0 sm:gap-2 sm:px-4 sm:py-2 sm:text-sm"
            >
              <span>Visualizar Oferta</span>
              <TrendingDown :size="14" :stroke-width="2.5" class="sm:hidden" />
              <TrendingDown
                :size="16"
                :stroke-width="2.5"
                class="hidden sm:block"
              />
            </div>
          </div>
          <div
            class="flex h-12 w-12 shrink-0 items-center justify-center rounded-lg bg-emerald-600 text-white shadow-lg sm:h-16 sm:w-16 sm:rounded-xl"
          >
            <TrendingDown :size="24" :stroke-width="2.5" class="sm:hidden" />
            <TrendingDown
              :size="32"
              :stroke-width="2.5"
              class="hidden sm:block"
            />
          </div>
        </div>
      </div>

      <div
        class="grid grid-cols-1 gap-4 sm:grid-cols-2 sm:gap-6 lg:grid-cols-4"
      >
        <div
          class="rounded-lg border border-neutral-200 bg-white p-4 shadow-sm transition-all duration-300 hover:shadow-md sm:rounded-xl sm:p-6"
        >
          <div class="flex items-start justify-between">
            <div class="min-w-0 flex-1">
              <div class="flex items-center gap-1.5 text-neutral-500 sm:gap-2">
                <DollarSign
                  :size="14"
                  :stroke-width="2"
                  class="shrink-0 sm:hidden"
                />
                <DollarSign
                  :size="16"
                  :stroke-width="2"
                  class="hidden shrink-0 sm:block"
                />
                <p class="truncate text-xs font-medium tracking-wide uppercase">
                  Sua Receita Total
                </p>
              </div>
              <p
                class="mt-2 text-xl font-bold break-words text-neutral-900 sm:mt-3 sm:text-2xl"
              >
                {{ formatCurrency(analytics.supplier_info.money_earned) }}
              </p>
              <p class="mt-1.5 text-xs text-neutral-400 sm:mt-2">
                Média de
                {{
                  formatCurrency(
                    analytics.supplier_info.money_earned /
                      analytics.supplier_info.purchases_count,
                  )
                }}
                por venda
              </p>
            </div>
            <div
              class="flex h-10 w-10 shrink-0 items-center justify-center rounded-lg bg-emerald-50 text-emerald-600 sm:h-12 sm:w-12"
            >
              <DollarSign :size="20" :stroke-width="2" class="sm:hidden" />
              <DollarSign
                :size="24"
                :stroke-width="2"
                class="hidden sm:block"
              />
            </div>
          </div>
        </div>

        <div
          class="rounded-lg border border-neutral-200 bg-white p-4 shadow-sm transition-all duration-300 hover:shadow-md sm:rounded-xl sm:p-6"
        >
          <div class="flex items-start justify-between">
            <div class="min-w-0 flex-1">
              <div class="flex items-center gap-1.5 text-neutral-500 sm:gap-2">
                <ShoppingCart
                  :size="14"
                  :stroke-width="2"
                  class="shrink-0 sm:hidden"
                />
                <ShoppingCart
                  :size="16"
                  :stroke-width="2"
                  class="hidden shrink-0 sm:block"
                />
                <p class="truncate text-xs font-medium tracking-wide uppercase">
                  Suas Vendas Concluídas
                </p>
              </div>
              <p
                class="mt-2 text-2xl font-bold text-neutral-900 sm:mt-3 sm:text-3xl"
              >
                {{ analytics.supplier_info.purchases_count }}
              </p>
              <p class="mt-1 text-xs text-neutral-500">
                {{
                  analytics.supplier_info.purchases_count === 1
                    ? "Transação finalizada"
                    : "Transações finalizadas"
                }}
              </p>
            </div>
            <div
              class="flex h-10 w-10 shrink-0 items-center justify-center rounded-lg bg-blue-50 text-blue-600 sm:h-12 sm:w-12"
            >
              <ShoppingCart :size="20" :stroke-width="2" class="sm:hidden" />
              <ShoppingCart
                :size="24"
                :stroke-width="2"
                class="hidden sm:block"
              />
            </div>
          </div>
        </div>

        <div
          class="rounded-lg border border-neutral-200 bg-white p-4 shadow-sm transition-all duration-300 hover:shadow-md sm:rounded-xl sm:p-6"
        >
          <div class="flex items-start justify-between">
            <div class="min-w-0 flex-1">
              <div class="flex items-center gap-1.5 text-neutral-500 sm:gap-2">
                <Package
                  :size="14"
                  :stroke-width="2"
                  class="shrink-0 sm:hidden"
                />
                <Package
                  :size="16"
                  :stroke-width="2"
                  class="hidden shrink-0 sm:block"
                />
                <p class="truncate text-xs font-medium tracking-wide uppercase">
                  Suas Ofertas Ativas
                </p>
              </div>
              <p
                class="mt-2 text-2xl font-bold text-neutral-900 sm:mt-3 sm:text-3xl"
              >
                {{ analytics.supplier_info.active_offers }}
              </p>
              <p class="mt-1 text-xs text-neutral-500">
                {{
                  analytics.supplier_info.active_offers === 0
                    ? "Ofertas publicadas"
                    : analytics.supplier_info.active_offers === 1
                      ? "Oferta disponível"
                      : "Ofertas disponíveis"
                }}
              </p>
            </div>
            <div
              class="flex h-10 w-10 shrink-0 items-center justify-center rounded-lg bg-purple-50 text-purple-600 sm:h-12 sm:w-12"
            >
              <Package :size="20" :stroke-width="2" class="sm:hidden" />
              <Package :size="24" :stroke-width="2" class="hidden sm:block" />
            </div>
          </div>
        </div>

        <div
          class="rounded-lg border border-neutral-200 bg-white p-4 shadow-sm transition-all duration-300 hover:shadow-md sm:rounded-xl sm:p-6"
        >
          <div class="flex items-start justify-between">
            <div class="min-w-0 flex-1">
              <div class="flex items-center gap-1.5 text-neutral-500 sm:gap-2">
                <Clock
                  :size="14"
                  :stroke-width="2"
                  class="shrink-0 sm:hidden"
                />
                <Clock
                  :size="16"
                  :stroke-width="2"
                  class="hidden shrink-0 sm:block"
                />
                <p class="truncate text-xs font-medium tracking-wide uppercase">
                  Suas Ofertas Expirando
                </p>
              </div>
              <p
                class="mt-2 text-2xl font-bold sm:mt-3 sm:text-3xl"
                :class="
                  analytics.supplier_info.almost_expiring_offers > 0
                    ? 'text-amber-600'
                    : 'text-neutral-900'
                "
              >
                {{ analytics.supplier_info.almost_expiring_offers }}
              </p>
              <p
                class="mt-1 text-xs"
                :class="
                  analytics.supplier_info.almost_expiring_offers > 0
                    ? 'text-amber-500'
                    : 'text-neutral-500'
                "
              >
                {{
                  analytics.supplier_info.almost_expiring_offers === 0
                    ? "Ofertas expirando"
                    : analytics.supplier_info.almost_expiring_offers === 1
                      ? "Oferta requer atenção"
                      : "Ofertas requerem atenção"
                }}
              </p>
            </div>
            <div
              class="flex h-10 w-10 shrink-0 items-center justify-center rounded-lg sm:h-12 sm:w-12"
              :class="
                analytics.supplier_info.almost_expiring_offers > 0
                  ? 'bg-amber-50 text-amber-600'
                  : 'bg-neutral-50 text-neutral-400'
              "
            >
              <Clock :size="20" :stroke-width="2" class="sm:hidden" />
              <Clock :size="24" :stroke-width="2" class="hidden sm:block" />
            </div>
          </div>
        </div>
      </div>

      <div
        v-if="analytics.buyer_info"
        class="grid grid-cols-1 gap-4 sm:gap-6 md:grid-cols-2"
      >
        <div
          class="rounded-lg border border-neutral-200 bg-white p-4 shadow-sm transition-all duration-300 hover:shadow-md sm:rounded-xl sm:p-6"
        >
          <div class="flex items-start justify-between">
            <div class="min-w-0 flex-1">
              <div class="flex items-center gap-1.5 text-neutral-500 sm:gap-2">
                <ShoppingCart
                  :size="14"
                  :stroke-width="2"
                  class="shrink-0 sm:hidden"
                />
                <ShoppingCart
                  :size="16"
                  :stroke-width="2"
                  class="hidden shrink-0 sm:block"
                />
                <p class="truncate text-xs font-medium tracking-wide uppercase">
                  Suas Compras Realizadas
                </p>
              </div>
              <p
                class="mt-2 text-2xl font-bold text-neutral-900 sm:mt-3 sm:text-3xl"
              >
                {{ analytics.buyer_info.purchases_count }}
              </p>
              <p class="mt-1 text-xs text-neutral-500">
                {{
                  analytics.buyer_info.purchases_count === 1
                    ? "Compra concluída"
                    : "Compras concluídas"
                }}
              </p>
            </div>
            <div
              class="flex h-10 w-10 shrink-0 items-center justify-center rounded-lg bg-blue-50 text-blue-600 sm:h-12 sm:w-12"
            >
              <ShoppingCart :size="20" :stroke-width="2" class="sm:hidden" />
              <ShoppingCart
                :size="24"
                :stroke-width="2"
                class="hidden sm:block"
              />
            </div>
          </div>
        </div>

        <div
          class="rounded-lg border border-neutral-200 bg-white p-4 shadow-sm transition-all duration-300 hover:shadow-md sm:rounded-xl sm:p-6"
        >
          <div class="flex items-start justify-between">
            <div class="min-w-0 flex-1">
              <div class="flex items-center gap-1.5 text-neutral-500 sm:gap-2">
                <Zap :size="14" :stroke-width="2" class="shrink-0 sm:hidden" />
                <Zap
                  :size="16"
                  :stroke-width="2"
                  class="hidden shrink-0 sm:block"
                />
                <p class="truncate text-xs font-medium tracking-wide uppercase">
                  Sua Energia Adquirida
                </p>
              </div>
              <div class="mt-2 flex items-baseline gap-1.5 sm:mt-3 sm:gap-2">
                <p class="text-2xl font-bold text-neutral-900 sm:text-3xl">
                  {{ formatNumber(analytics.buyer_info.energy_transacted, 0) }}
                </p>
                <p class="text-xs text-neutral-500 sm:text-sm">kWh</p>
              </div>
              <p class="mt-1.5 text-xs text-neutral-400 sm:mt-2">
                Média de
                {{
                  formatNumber(
                    analytics.buyer_info.energy_transacted /
                      analytics.buyer_info.purchases_count,
                    0,
                  )
                }}
                kWh por compra
              </p>
            </div>
            <div
              class="flex h-10 w-10 shrink-0 items-center justify-center rounded-lg bg-amber-50 text-amber-600 sm:h-12 sm:w-12"
            >
              <Zap :size="20" :stroke-width="2" class="sm:hidden" />
              <Zap :size="24" :stroke-width="2" class="hidden sm:block" />
            </div>
          </div>
        </div>
      </div>
    </div>

    <div
      v-if="!loading && marketInsights.length > 0"
      class="rounded-lg border border-neutral-200 bg-white shadow-sm transition-all duration-300 hover:shadow-md sm:rounded-xl"
    >
      <div class="border-b border-neutral-200 p-4 sm:p-6">
        <div>
          <h2 class="text-base font-semibold text-neutral-900 sm:text-lg">
            Insights {{ isBuyer ? "de Compra" : "de Vendas" }}
          </h2>
          <p class="mt-1 text-xs text-neutral-500 sm:text-sm">
            Informações em tempo real para suas decisões
          </p>
        </div>
      </div>
      <div
        class="grid grid-cols-1 gap-3 p-4 sm:gap-4 sm:p-6 md:grid-cols-2 lg:grid-cols-3"
      >
        <div
          v-for="insight in marketInsights"
          :key="insight.id"
          class="group relative overflow-hidden rounded-lg border border-neutral-200 bg-white p-4 shadow-sm transition-all duration-300 hover:scale-[1.02] hover:shadow-md active:scale-[0.98] sm:rounded-xl sm:p-5"
          :class="{ 'cursor-pointer': insight.id === 1 && isBuyer }"
          @click="
            insight.id === 1 &&
            isBuyer &&
            analytics?.buyer_info?.advantage_offer_uuid
              ? goToOffer(analytics.buyer_info.advantage_offer_uuid)
              : null
          "
        >
          <div class="flex items-start gap-2.5 sm:gap-3">
            <div
              :class="[
                'flex h-10 w-10 flex-shrink-0 items-center justify-center rounded-lg transition-all duration-300 group-hover:scale-110 sm:h-12 sm:w-12 sm:rounded-xl',
                insight.bgClass,
                insight.colorClass,
              ]"
            >
              <component
                :is="insight.icon"
                :size="20"
                :stroke-width="2"
                class="sm:hidden"
              />
              <component
                :is="insight.icon"
                :size="24"
                :stroke-width="2"
                class="hidden sm:block"
              />
            </div>
            <div class="min-w-0 flex-1">
              <p class="truncate text-xs font-medium text-neutral-500">
                {{ insight.title }}
              </p>
              <h3
                :class="[
                  'mt-1 text-xl font-bold sm:text-2xl',
                  insight.colorClass,
                ]"
              >
                {{ insight.value }}
              </h3>
              <p
                class="mt-1.5 line-clamp-2 text-xs text-neutral-600 sm:mt-2 sm:text-sm"
              >
                {{ insight.subtitle }}
              </p>
              <div class="mt-2 flex items-center gap-1 sm:mt-3">
                <component
                  :is="
                    insight.trend === 'up'
                      ? TrendingUp
                      : insight.trend === 'down'
                        ? TrendingDown
                        : Clock
                  "
                  :size="12"
                  :class="[
                    'sm:hidden',
                    insight.trend === 'up'
                      ? 'text-red-500'
                      : insight.trend === 'down'
                        ? 'text-emerald-500'
                        : 'text-amber-500',
                  ]"
                />
                <component
                  :is="
                    insight.trend === 'up'
                      ? TrendingUp
                      : insight.trend === 'down'
                        ? TrendingDown
                        : Clock
                  "
                  :size="14"
                  :class="[
                    'hidden sm:block',
                    insight.trend === 'up'
                      ? 'text-red-500'
                      : insight.trend === 'down'
                        ? 'text-emerald-500'
                        : 'text-amber-500',
                  ]"
                />
                <span
                  :class="[
                    'text-xs font-medium',
                    insight.trend === 'up'
                      ? 'text-red-600'
                      : insight.trend === 'down'
                        ? 'text-emerald-600'
                        : 'text-amber-600',
                  ]"
                >
                  {{
                    insight.trend === "up"
                      ? "Em alta"
                      : insight.trend === "down"
                        ? "Oportunidade"
                        : "Atenção"
                  }}
                </span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped></style>
