<script setup lang="ts">
import {
  DollarSign,
  TrendingDown,
  TrendingUp,
  Zap,
  TrendingDown as ChartDown,
  AlertCircle,
  CheckCircle,
  Clock,
} from "lucide-vue-next";
import type { Component } from "vue";

interface StatCard {
  title: string;
  value: string;
  change: string;
  isPositive: boolean;
  icon: Component;
}

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

const stats: StatCard[] = [
  {
    title: "Economia Total",
    value: "R$ 127.450",
    change: "+12.5%",
    isPositive: true,
    icon: DollarSign,
  },
  {
    title: "Energia Economizada",
    value: "2.430 kWh",
    change: "+8.2%",
    isPositive: true,
    icon: Zap,
  },
  {
    title: "Negociações Ativas",
    value: "18",
    change: "-2",
    isPositive: false,
    icon: TrendingUp,
  },
  {
    title: "Taxa Média",
    value: "R$ 0,52/kWh",
    change: "+3.1%",
    isPositive: true,
    icon: TrendingDown,
  },
];

const marketInsights: MarketInsight[] = [
  {
    id: 1,
    title: "Melhor Oferta Disponível",
    value: "R$ 0,48/kWh",
    subtitle: "Energia Solar • 850 kWh disponíveis",
    trend: "down",
    icon: TrendingDown,
    colorClass: "text-emerald-600",
    bgClass: "bg-emerald-50",
  },
  {
    id: 2,
    title: "Ofertas Expirando",
    value: "3 ofertas",
    subtitle: "Expiram nas próximas 48 horas",
    trend: "neutral",
    icon: Clock,
    colorClass: "text-amber-600",
    bgClass: "bg-amber-50",
  },
  {
    id: 3,
    title: "Taxa de Sucesso",
    value: "94%",
    subtitle: "Das suas negociações foram concluídas",
    trend: "up",
    icon: CheckCircle,
    colorClass: "text-blue-600",
    bgClass: "bg-blue-50",
  },
  {
    id: 4,
    title: "Alerta de Mercado",
    value: "Alta Demanda",
    subtitle: "Preços 7% acima da média semanal",
    trend: "up",
    icon: AlertCircle,
    colorClass: "text-purple-600",
    bgClass: "bg-purple-50",
  },
];
</script>

<template>
  <div class="flex flex-col gap-6">
    <div class="grid grid-cols-1 gap-6 md:grid-cols-2 lg:grid-cols-4">
      <div
        v-for="stat in stats"
        :key="stat.title"
        class="group relative overflow-hidden rounded-xl border border-neutral-200 bg-white p-6 shadow-sm transition-all duration-300 hover:shadow-md"
      >
        <div class="flex items-start justify-between">
          <div class="flex-1">
            <p class="text-sm font-medium text-neutral-500">{{ stat.title }}</p>
            <h3 class="mt-2 text-3xl font-bold text-neutral-900">
              {{ stat.value }}
            </h3>
            <div class="mt-3 flex items-center gap-1">
              <component
                :is="stat.isPositive ? TrendingUp : TrendingDown"
                :size="16"
                :class="[
                  'transition-colors',
                  stat.isPositive ? 'text-emerald-500' : 'text-red-500',
                ]"
              />
              <span
                :class="[
                  'text-sm font-medium',
                  stat.isPositive ? 'text-emerald-600' : 'text-red-600',
                ]"
              >
                {{ stat.change }}
              </span>
              <span class="text-sm text-neutral-400">vs. mês anterior</span>
            </div>
          </div>
          <div
            class="flex h-12 w-12 items-center justify-center rounded-lg bg-gradient-to-br from-emerald-50 to-teal-50 text-emerald-600 transition-all duration-300 group-hover:scale-110"
          >
            <component :is="stat.icon" :size="24" :stroke-width="2" />
          </div>
        </div>
      </div>
    </div>

    <div
      class="rounded-xl border border-neutral-200 bg-white shadow-sm transition-all duration-300 hover:shadow-md"
    >
      <div class="border-b border-neutral-200 p-6">
        <div class="flex items-center justify-between">
          <div>
            <h2 class="text-lg font-semibold text-neutral-900">
              Insights do Mercado
            </h2>
            <p class="mt-1 text-sm text-neutral-500">
              Informações em tempo real para suas decisões
            </p>
          </div>
          <div
            class="flex items-center gap-2 rounded-lg bg-emerald-50 px-3 py-2"
          >
            <div
              class="h-2 w-2 animate-pulse rounded-full bg-emerald-500"
            ></div>
            <span class="text-sm font-medium text-emerald-700">Ao vivo</span>
          </div>
        </div>
      </div>
      <div class="grid grid-cols-1 gap-4 p-6 md:grid-cols-2 lg:grid-cols-4">
        <div
          v-for="insight in marketInsights"
          :key="insight.id"
          class="group relative overflow-hidden rounded-xl border border-neutral-200 bg-white p-5 shadow-sm transition-all duration-300 hover:scale-[1.02] hover:shadow-md"
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
