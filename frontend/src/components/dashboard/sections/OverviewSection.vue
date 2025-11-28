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
      title: "Preço Médio",
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
  <div class="flex flex-col gap-6">
    <div v-if="loading" class="grid grid-cols-1 gap-6 lg:grid-cols-2">
      <div
        v-for="i in 4"
        :key="i"
        class="h-64 animate-pulse rounded-xl bg-neutral-200"
      ></div>
    </div>

    <div v-else-if="isBuyer && analytics?.buyer_info" class="space-y-6">
      <div
        v-if="
          analytics.buyer_info.advantage_offer_uuid &&
          analytics.buyer_info.advantage_offer_uuid !== ''
        "
        class="cursor-pointer rounded-xl border-2 border-emerald-300 bg-gradient-to-br from-emerald-50 to-green-50 p-6 shadow-md transition-all duration-300 hover:scale-[1.02] hover:border-emerald-400 hover:shadow-lg"
        @click="goToOffer(analytics.buyer_info.advantage_offer_uuid)"
      >
        <div class="flex items-start justify-between">
          <div class="flex-1">
            <div class="flex items-center gap-2 text-emerald-700">
              <TrendingDown :size="18" :stroke-width="2.5" />
              <p class="text-sm font-bold tracking-wide uppercase">
                🎯 Melhor Oferta Disponível
              </p>
            </div>
            <h3 class="mt-3 text-2xl font-bold text-emerald-900">
              Ver Oferta Vantajosa
            </h3>
            <p class="mt-2 text-sm text-emerald-700">
              Encontramos a melhor oferta do seu submercado. Clique para
              visualizar os detalhes e aproveitar!
            </p>
            <div
              class="mt-4 inline-flex items-center gap-2 rounded-lg bg-emerald-600 px-4 py-2 text-sm font-semibold text-white transition-colors hover:bg-emerald-700"
            >
              <span>Visualizar Oferta</span>
              <TrendingDown :size="16" :stroke-width="2.5" />
            </div>
          </div>
          <div
            class="flex h-16 w-16 items-center justify-center rounded-xl bg-emerald-600 text-white shadow-lg"
          >
            <TrendingDown :size="32" :stroke-width="2.5" />
          </div>
        </div>
      </div>

      <div class="grid grid-cols-1 gap-6 md:grid-cols-2">
        <div
          class="rounded-xl border border-neutral-200 bg-white p-6 shadow-sm transition-all duration-300 hover:shadow-md"
        >
          <div class="flex items-start justify-between">
            <div class="flex-1">
              <div class="flex items-center gap-2 text-neutral-500">
                <ShoppingCart :size="16" :stroke-width="2" />
                <p class="text-xs font-medium tracking-wide uppercase">
                  Total de Compras
                </p>
              </div>
              <p class="mt-3 text-3xl font-bold text-neutral-900">
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
              class="flex h-12 w-12 items-center justify-center rounded-lg bg-blue-50 text-blue-600"
            >
              <ShoppingCart :size="24" :stroke-width="2" />
            </div>
          </div>
        </div>

        <div
          class="rounded-xl border border-neutral-200 bg-white p-6 shadow-sm transition-all duration-300 hover:shadow-md"
        >
          <div class="flex items-start justify-between">
            <div class="flex-1">
              <div class="flex items-center gap-2 text-neutral-500">
                <Zap :size="16" :stroke-width="2" />
                <p class="text-xs font-medium tracking-wide uppercase">
                  Energia Movimentada
                </p>
              </div>
              <div class="mt-3 flex items-baseline gap-2">
                <p class="text-3xl font-bold text-neutral-900">
                  {{ formatNumber(analytics.buyer_info.energy_transacted, 0) }}
                </p>
                <p class="text-sm text-neutral-500">kWh</p>
              </div>
              <p class="mt-2 text-xs text-neutral-400">
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
              class="flex h-12 w-12 items-center justify-center rounded-lg bg-amber-50 text-amber-600"
            >
              <Zap :size="24" :stroke-width="2" />
            </div>
          </div>
        </div>
      </div>
    </div>

    <div v-else-if="isSupplier && analytics?.supplier_info" class="space-y-6">
      <div
        v-if="
          analytics.buyer_info?.advantage_offer_uuid &&
          analytics.buyer_info.advantage_offer_uuid !== ''
        "
        class="cursor-pointer rounded-xl border-2 border-emerald-300 bg-gradient-to-br from-emerald-50 to-green-50 p-6 shadow-md transition-all duration-300 hover:scale-[1.02] hover:border-emerald-400 hover:shadow-lg"
        @click="goToOffer(analytics.buyer_info.advantage_offer_uuid)"
      >
        <div class="flex items-start justify-between">
          <div class="flex-1">
            <div class="flex items-center gap-2 text-emerald-700">
              <TrendingDown :size="18" :stroke-width="2.5" />
              <p class="text-sm font-bold tracking-wide uppercase">
                🎯 Melhor Oferta Disponível
              </p>
            </div>
            <h3 class="mt-3 text-2xl font-bold text-emerald-900">
              Oportunidade de Compra Vantajosa
            </h3>
            <p class="mt-2 text-sm text-emerald-700">
              Encontramos a melhor oferta do seu submercado. Clique para
              visualizar os detalhes e aproveitar!
            </p>
            <div
              class="mt-4 inline-flex items-center gap-2 rounded-lg bg-emerald-600 px-4 py-2 text-sm font-semibold text-white transition-colors hover:bg-emerald-700"
            >
              <span>Visualizar Oferta</span>
              <TrendingDown :size="16" :stroke-width="2.5" />
            </div>
          </div>
          <div
            class="flex h-16 w-16 items-center justify-center rounded-xl bg-emerald-600 text-white shadow-lg"
          >
            <TrendingDown :size="32" :stroke-width="2.5" />
          </div>
        </div>
      </div>

      <div class="grid grid-cols-1 gap-6 md:grid-cols-2 lg:grid-cols-4">
        <div
          class="rounded-xl border border-neutral-200 bg-white p-6 shadow-sm transition-all duration-300 hover:shadow-md"
        >
          <div class="flex items-start justify-between">
            <div class="flex-1">
              <div class="flex items-center gap-2 text-neutral-500">
                <DollarSign :size="16" :stroke-width="2" />
                <p class="text-xs font-medium tracking-wide uppercase">
                  Receita Total
                </p>
              </div>
              <p class="mt-3 text-2xl font-bold text-neutral-900">
                {{ formatCurrency(analytics.supplier_info.money_earned) }}
              </p>
              <p class="mt-2 text-xs text-neutral-400">
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
              class="flex h-12 w-12 items-center justify-center rounded-lg bg-emerald-50 text-emerald-600"
            >
              <DollarSign :size="24" :stroke-width="2" />
            </div>
          </div>
        </div>

        <div
          class="rounded-xl border border-neutral-200 bg-white p-6 shadow-sm transition-all duration-300 hover:shadow-md"
        >
          <div class="flex items-start justify-between">
            <div class="flex-1">
              <div class="flex items-center gap-2 text-neutral-500">
                <ShoppingCart :size="16" :stroke-width="2" />
                <p class="text-xs font-medium tracking-wide uppercase">
                  Vendas Concluídas
                </p>
              </div>
              <p class="mt-3 text-3xl font-bold text-neutral-900">
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
              class="flex h-12 w-12 items-center justify-center rounded-lg bg-blue-50 text-blue-600"
            >
              <ShoppingCart :size="24" :stroke-width="2" />
            </div>
          </div>
        </div>

        <div
          class="rounded-xl border border-neutral-200 bg-white p-6 shadow-sm transition-all duration-300 hover:shadow-md"
        >
          <div class="flex items-start justify-between">
            <div class="flex-1">
              <div class="flex items-center gap-2 text-neutral-500">
                <Package :size="16" :stroke-width="2" />
                <p class="text-xs font-medium tracking-wide uppercase">
                  Ofertas Ativas
                </p>
              </div>
              <p class="mt-3 text-3xl font-bold text-neutral-900">
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
              class="flex h-12 w-12 items-center justify-center rounded-lg bg-purple-50 text-purple-600"
            >
              <Package :size="24" :stroke-width="2" />
            </div>
          </div>
        </div>

        <div
          class="rounded-xl border border-neutral-200 bg-white p-6 shadow-sm transition-all duration-300 hover:shadow-md"
        >
          <div class="flex items-start justify-between">
            <div class="flex-1">
              <div class="flex items-center gap-2 text-neutral-500">
                <Clock :size="16" :stroke-width="2" />
                <p class="text-xs font-medium tracking-wide uppercase">
                  Expirando em 48h
                </p>
              </div>
              <p
                class="mt-3 text-3xl font-bold"
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
              class="flex h-12 w-12 items-center justify-center rounded-lg"
              :class="
                analytics.supplier_info.almost_expiring_offers > 0
                  ? 'bg-amber-50 text-amber-600'
                  : 'bg-neutral-50 text-neutral-400'
              "
            >
              <Clock :size="24" :stroke-width="2" />
            </div>
          </div>
        </div>
      </div>

      <div
        v-if="analytics.buyer_info"
        class="grid grid-cols-1 gap-6 md:grid-cols-2"
      >
        <div
          class="rounded-xl border border-neutral-200 bg-white p-6 shadow-sm transition-all duration-300 hover:shadow-md"
        >
          <div class="flex items-start justify-between">
            <div class="flex-1">
              <div class="flex items-center gap-2 text-neutral-500">
                <ShoppingCart :size="16" :stroke-width="2" />
                <p class="text-xs font-medium tracking-wide uppercase">
                  Compras Realizadas
                </p>
              </div>
              <p class="mt-3 text-3xl font-bold text-neutral-900">
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
              class="flex h-12 w-12 items-center justify-center rounded-lg bg-blue-50 text-blue-600"
            >
              <ShoppingCart :size="24" :stroke-width="2" />
            </div>
          </div>
        </div>

        <div
          class="rounded-xl border border-neutral-200 bg-white p-6 shadow-sm transition-all duration-300 hover:shadow-md"
        >
          <div class="flex items-start justify-between">
            <div class="flex-1">
              <div class="flex items-center gap-2 text-neutral-500">
                <Zap :size="16" :stroke-width="2" />
                <p class="text-xs font-medium tracking-wide uppercase">
                  Energia Adquirida
                </p>
              </div>
              <div class="mt-3 flex items-baseline gap-2">
                <p class="text-3xl font-bold text-neutral-900">
                  {{ formatNumber(analytics.buyer_info.energy_transacted, 0) }}
                </p>
                <p class="text-sm text-neutral-500">kWh</p>
              </div>
              <p class="mt-2 text-xs text-neutral-400">
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
              class="flex h-12 w-12 items-center justify-center rounded-lg bg-amber-50 text-amber-600"
            >
              <Zap :size="24" :stroke-width="2" />
            </div>
          </div>
        </div>
      </div>
    </div>

    <div
      v-if="!loading && marketInsights.length > 0"
      class="rounded-xl border border-neutral-200 bg-white shadow-sm transition-all duration-300 hover:shadow-md"
    >
      <div class="border-b border-neutral-200 p-6">
        <div>
          <h2 class="text-lg font-semibold text-neutral-900">
            Insights {{ isBuyer ? "de Compra" : "de Vendas" }}
          </h2>
          <p class="mt-1 text-sm text-neutral-500">
            Informações em tempo real para suas decisões
          </p>
        </div>
      </div>
      <div class="grid grid-cols-1 gap-4 p-6 md:grid-cols-2 lg:grid-cols-3">
        <div
          v-for="insight in marketInsights"
          :key="insight.id"
          class="group relative overflow-hidden rounded-xl border border-neutral-200 bg-white p-5 shadow-sm transition-all duration-300 hover:scale-[1.02] hover:shadow-md"
          :class="{ 'cursor-pointer': insight.id === 1 && isBuyer }"
          @click="
            insight.id === 1 &&
            isBuyer &&
            analytics?.buyer_info?.advantage_offer_uuid
              ? goToOffer(analytics.buyer_info.advantage_offer_uuid)
              : null
          "
        >
          <div class="flex items-start gap-3">
            <div
              :class="[
                'flex h-12 w-12 flex-shrink-0 items-center justify-center rounded-xl transition-all duration-300 group-hover:scale-110',
                insight.bgClass,
                insight.colorClass,
              ]"
            >
              <component :is="insight.icon" :size="24" :stroke-width="2" />
            </div>
            <div class="min-w-0 flex-1">
              <p class="truncate text-xs font-medium text-neutral-500">
                {{ insight.title }}
              </p>
              <h3 :class="['mt-1 text-2xl font-bold', insight.colorClass]">
                {{ insight.value }}
              </h3>
              <p class="mt-2 line-clamp-2 text-sm text-neutral-600">
                {{ insight.subtitle }}
              </p>
              <div class="mt-3 flex items-center gap-1">
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
