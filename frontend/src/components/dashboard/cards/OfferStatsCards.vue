<script setup lang="ts">
import { computed } from "vue";
import { TrendingUp, DollarSign, ShoppingCart, Package } from "lucide-vue-next";

interface Props {
  totalSales: number;
  totalRevenue: number;
  totalQuantitySold: number;
  remainingQuantity: number;
}

const props = defineProps<Props>();

const formatCurrency = (value: number): string => {
  return `R$ ${value.toLocaleString("pt-BR", { minimumFractionDigits: 2, maximumFractionDigits: 2 })}`;
};

const formatQuantity = (value: number): string => {
  return `${value.toLocaleString("pt-BR")} MWh`;
};

const stats = computed(() => [
  {
    title: "Vendas Totais",
    value: props.totalSales.toString(),
    icon: ShoppingCart,
    gradient: "from-blue-500 to-blue-600",
    bg: "bg-blue-50",
    iconColor: "text-blue-600",
  },
  {
    title: "Receita Total",
    value: formatCurrency(props.totalRevenue),
    icon: DollarSign,
    gradient: "from-emerald-500 to-emerald-600",
    bg: "bg-emerald-50",
    iconColor: "text-emerald-600",
  },
  {
    title: "Quantidade Vendida",
    value: formatQuantity(props.totalQuantitySold),
    icon: TrendingUp,
    gradient: "from-purple-500 to-purple-600",
    bg: "bg-purple-50",
    iconColor: "text-purple-600",
  },
  {
    title: "Estoque Restante",
    value: formatQuantity(props.remainingQuantity),
    icon: Package,
    gradient: "from-orange-500 to-orange-600",
    bg: "bg-orange-50",
    iconColor: "text-orange-600",
  },
]);
</script>

<template>
  <div class="grid grid-cols-1 gap-4 sm:grid-cols-2 lg:grid-cols-4">
    <div
      v-for="stat in stats"
      :key="stat.title"
      class="group relative overflow-hidden rounded-xl border border-neutral-200 bg-white p-6 shadow-sm transition-all duration-300 hover:shadow-md"
    >
      <div
        :class="[
          'absolute top-0 right-0 h-24 w-24 translate-x-8 -translate-y-8 rounded-full bg-gradient-to-br opacity-10',
          stat.gradient,
        ]"
      ></div>

      <div class="relative">
        <div class="flex items-start justify-between">
          <div class="flex-1">
            <p class="text-sm font-medium text-neutral-500">{{ stat.title }}</p>
            <p class="mt-2 text-2xl font-bold text-neutral-900">
              {{ stat.value }}
            </p>
          </div>
          <div :class="['rounded-lg p-2', stat.bg]">
            <component :is="stat.icon" :size="20" :class="stat.iconColor" />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
