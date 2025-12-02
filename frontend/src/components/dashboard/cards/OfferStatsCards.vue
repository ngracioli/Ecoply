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
  <div
    class="grid grid-cols-1 gap-3 px-4 sm:grid-cols-2 sm:gap-4 sm:px-0 lg:grid-cols-4"
  >
    <div
      v-for="stat in stats"
      :key="stat.title"
      class="group relative overflow-hidden rounded-lg border border-neutral-200 bg-white p-4 shadow-sm transition-all duration-300 hover:shadow-md sm:rounded-xl sm:p-6"
    >
      <div
        :class="[
          'absolute top-0 right-0 h-20 w-20 translate-x-6 -translate-y-6 rounded-full bg-gradient-to-br opacity-10 sm:h-24 sm:w-24 sm:translate-x-8 sm:-translate-y-8',
          stat.gradient,
        ]"
      ></div>

      <div class="relative">
        <div class="flex items-start justify-between gap-3">
          <div class="min-w-0 flex-1">
            <p class="text-xs font-medium text-neutral-500 sm:text-sm">
              {{ stat.title }}
            </p>
            <p
              class="mt-1.5 text-xl font-bold break-words text-neutral-900 sm:mt-2 sm:text-2xl"
            >
              {{ stat.value }}
            </p>
          </div>
          <div :class="['shrink-0 rounded-lg p-1.5 sm:p-2', stat.bg]">
            <component
              :is="stat.icon"
              :size="18"
              class="sm:hidden"
              :class="stat.iconColor"
            />
            <component
              :is="stat.icon"
              :size="20"
              class="hidden sm:block"
              :class="stat.iconColor"
            />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
