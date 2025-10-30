<script setup lang="ts">
import { DollarSign, TrendingDown, TrendingUp, Zap } from "lucide-vue-next";
import type { Component } from "vue";

type ActivityType = "offer" | "success" | "payment";

interface StatCard {
  title: string;
  value: string;
  change: string;
  isPositive: boolean;
  icon: Component;
}

interface Activity {
  id: number;
  title: string;
  description: string;
  time: string;
  type: ActivityType;
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

const recentActivity: Activity[] = [
  {
    id: 1,
    title: "Nova oferta disponível",
    description: "Energia solar - 500 kWh",
    time: "Há 2 horas",
    type: "offer",
  },
  {
    id: 2,
    title: "Negociação concluída",
    description: "Contrato #1234 - R$ 2.500",
    time: "Há 5 horas",
    type: "success",
  },
  {
    id: 3,
    title: "Pagamento processado",
    description: "Fatura #789 - R$ 1.800",
    time: "Ontem",
    type: "payment",
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
        <h2 class="text-lg font-semibold text-neutral-900">
          Atividades Recentes
        </h2>
        <p class="mt-1 text-sm text-neutral-500">
          Acompanhe suas últimas movimentações
        </p>
      </div>
      <div class="divide-y divide-neutral-100">
        <div
          v-for="activity in recentActivity"
          :key="activity.id"
          class="group flex items-center gap-4 p-6 transition-all duration-200 hover:bg-neutral-50"
        >
          <div
            :class="[
              'flex h-10 w-10 flex-shrink-0 items-center justify-center rounded-lg transition-all duration-200',
              activity.type === 'offer'
                ? 'bg-blue-50 text-blue-600 group-hover:bg-blue-100'
                : activity.type === 'success'
                  ? 'bg-emerald-50 text-emerald-600 group-hover:bg-emerald-100'
                  : 'bg-purple-50 text-purple-600 group-hover:bg-purple-100',
            ]"
          >
            <Zap :size="20" :stroke-width="2" />
          </div>
          <div class="flex-1">
            <h4 class="font-medium text-neutral-900">{{ activity.title }}</h4>
            <p class="mt-0.5 text-sm text-neutral-500">
              {{ activity.description }}
            </p>
          </div>
          <span class="text-sm text-neutral-400">{{ activity.time }}</span>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped></style>
