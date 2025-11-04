<script setup lang="ts">
import type { Component } from "vue";
import { CheckCircle, Clock, XCircle } from "lucide-vue-next";

type TransactionStatus = "completed" | "pending" | "cancelled";

interface Transaction {
  id: number;
  title: string;
  date: string;
  amount: string;
  quantity: string;
  status: TransactionStatus;
  partner: string;
}

interface StatusConfig {
  icon: Component;
  label: string;
  bgColor: string;
  textColor: string;
  iconColor: string;
}

const STATUS_CONFIG: Record<TransactionStatus, StatusConfig> = {
  completed: {
    icon: CheckCircle,
    label: "Concluída",
    bgColor: "bg-emerald-100",
    textColor: "text-emerald-700",
    iconColor: "text-emerald-600",
  },
  pending: {
    icon: Clock,
    label: "Pendente",
    bgColor: "bg-amber-100",
    textColor: "text-amber-700",
    iconColor: "text-amber-600",
  },
  cancelled: {
    icon: XCircle,
    label: "Cancelada",
    bgColor: "bg-red-100",
    textColor: "text-red-700",
    iconColor: "text-red-600",
  },
};

const transactions: Transaction[] = [
  {
    id: 1,
    title: "Compra de Energia Solar",
    date: "20/10/2025",
    amount: "R$ 480,00",
    quantity: "1.000 kWh",
    status: "completed",
    partner: "EcoSolar Brasil",
  },
  {
    id: 2,
    title: "Compra de Energia Eólica",
    date: "18/10/2025",
    amount: "R$ 416,00",
    quantity: "800 kWh",
    status: "completed",
    partner: "VentoLimpo Energia",
  },
  {
    id: 3,
    title: "Compra de Energia Hidrelétrica",
    date: "15/10/2025",
    amount: "R$ 675,00",
    quantity: "1.500 kWh",
    status: "pending",
    partner: "AguaViva Energia",
  },
  {
    id: 4,
    title: "Compra de Energia Solar",
    date: "12/10/2025",
    amount: "R$ 240,00",
    quantity: "500 kWh",
    status: "completed",
    partner: "SolBrilhante",
  },
  {
    id: 5,
    title: "Compra de Energia Eólica",
    date: "08/10/2025",
    amount: "R$ 312,00",
    quantity: "600 kWh",
    status: "cancelled",
    partner: "VentoForte S.A.",
  },
];

const getStatusConfig = (status: TransactionStatus): StatusConfig => {
  return STATUS_CONFIG[status];
};

const formatTransactionId = (id: number): string => {
  return `#${String(id).padStart(4, "0")}`;
};
</script>

<template>
  <div class="flex flex-col gap-6">
    <div class="flex items-center justify-between">
      <div>
        <h2 class="text-2xl font-bold text-neutral-900">
          Histórico de Negociações
        </h2>
        <p class="mt-1 text-sm text-neutral-500">
          Acompanhe todas as suas transações
        </p>
      </div>
      <div class="flex gap-3">
        <select
          class="rounded-lg border border-neutral-300 bg-white px-4 py-2 text-sm text-neutral-700 shadow-sm transition-all duration-200 hover:border-neutral-400 focus:border-emerald-500 focus:ring-2 focus:ring-emerald-500/20 focus:outline-none"
        >
          <option value="all">Todos os status</option>
          <option value="completed">Concluídas</option>
          <option value="pending">Pendentes</option>
          <option value="cancelled">Canceladas</option>
        </select>
      </div>
    </div>

    <div
      class="overflow-hidden rounded-xl border border-neutral-200 bg-white shadow-sm"
    >
      <div class="overflow-x-auto">
        <table class="w-full">
          <thead class="bg-neutral-50">
            <tr>
              <th
                class="px-6 py-4 text-left text-xs font-semibold tracking-wider text-neutral-600 uppercase"
              >
                Transação
              </th>
              <th
                class="px-6 py-4 text-left text-xs font-semibold tracking-wider text-neutral-600 uppercase"
              >
                Parceiro
              </th>
              <th
                class="px-6 py-4 text-left text-xs font-semibold tracking-wider text-neutral-600 uppercase"
              >
                Quantidade
              </th>
              <th
                class="px-6 py-4 text-left text-xs font-semibold tracking-wider text-neutral-600 uppercase"
              >
                Valor
              </th>
              <th
                class="px-6 py-4 text-left text-xs font-semibold tracking-wider text-neutral-600 uppercase"
              >
                Data
              </th>
              <th
                class="px-6 py-4 text-left text-xs font-semibold tracking-wider text-neutral-600 uppercase"
              >
                Status
              </th>
            </tr>
          </thead>
          <tbody class="divide-y divide-neutral-100">
            <tr
              v-for="transaction in transactions"
              :key="transaction.id"
              class="transition-colors duration-150 hover:bg-neutral-50"
            >
              <td class="px-6 py-4">
                <div class="font-medium text-neutral-900">
                  {{ transaction.title }}
                </div>
                <div class="text-sm text-neutral-500">
                  ID: {{ formatTransactionId(transaction.id) }}
                </div>
              </td>
              <td class="px-6 py-4">
                <div class="text-sm text-neutral-700">
                  {{ transaction.partner }}
                </div>
              </td>
              <td class="px-6 py-4">
                <div class="text-sm font-medium text-neutral-900">
                  {{ transaction.quantity }}
                </div>
              </td>
              <td class="px-6 py-4">
                <div class="text-sm font-semibold text-neutral-900">
                  {{ transaction.amount }}
                </div>
              </td>
              <td class="px-6 py-4">
                <div class="text-sm text-neutral-600">
                  {{ transaction.date }}
                </div>
              </td>
              <td class="px-6 py-4">
                <span
                  :class="[
                    'inline-flex items-center gap-1.5 rounded-full px-3 py-1 text-xs font-medium',
                    getStatusConfig(transaction.status).bgColor,
                    getStatusConfig(transaction.status).textColor,
                  ]"
                >
                  <component
                    :is="getStatusConfig(transaction.status).icon"
                    :size="14"
                    :class="getStatusConfig(transaction.status).iconColor"
                  />
                  {{ getStatusConfig(transaction.status).label }}
                </span>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<style scoped></style>
