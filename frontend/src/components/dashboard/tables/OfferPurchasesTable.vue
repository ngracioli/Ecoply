<script setup lang="ts">
import type { Component } from "vue";
import { computed } from "vue";
import { CheckCircle, Clock, XCircle } from "lucide-vue-next";
import DataTable from "primevue/datatable";
import Column from "primevue/column";
import type { PurchaseListItem } from "../../../types/responses/purchases";

type TransactionStatus = "completed" | "waiting" | "canceled";

interface StatusConfig {
  icon: Component;
  label: string;
  bgColor: string;
  textColor: string;
  iconColor: string;
}

interface PurchaseWithTotal extends PurchaseListItem {
  total: number;
}

interface Props {
  purchases: PurchaseListItem[];
  loading?: boolean;
}

const props = defineProps<Props>();

const purchasesWithTotal = computed<PurchaseWithTotal[]>(() => {
  return props.purchases
    .map((purchase) => ({
      ...purchase,
      total: purchase.quantity_mwh * purchase.price_per_mwh,
    }))
    .sort(
      (a, b) =>
        new Date(b.created_at).getTime() - new Date(a.created_at).getTime(),
    );
});

const STATUS_CONFIG: Record<TransactionStatus, StatusConfig> = {
  completed: {
    icon: CheckCircle,
    label: "Concluída",
    bgColor: "bg-emerald-100",
    textColor: "text-emerald-700",
    iconColor: "text-emerald-600",
  },
  waiting: {
    icon: Clock,
    label: "Pendente",
    bgColor: "bg-amber-100",
    textColor: "text-amber-700",
    iconColor: "text-amber-600",
  },
  canceled: {
    icon: XCircle,
    label: "Cancelada",
    bgColor: "bg-red-100",
    textColor: "text-red-700",
    iconColor: "text-red-600",
  },
};

const getStatusConfig = (status: TransactionStatus): StatusConfig => {
  return STATUS_CONFIG[status];
};

const formatCurrency = (value: number): string => {
  return `R$ ${value.toLocaleString("pt-BR", { minimumFractionDigits: 2, maximumFractionDigits: 2 })}`;
};

const formatQuantity = (value: number): string => {
  return `${value.toLocaleString("pt-BR")} MWh`;
};

const formatDate = (dateString: string): string => {
  const date = new Date(dateString);
  return date.toLocaleDateString("pt-BR", {
    day: "2-digit",
    month: "2-digit",
    year: "numeric",
    hour: "2-digit",
    minute: "2-digit",
  });
};

const getPaymentMethodLabel = (method: string): string => {
  const methodLabels: Record<string, string> = {
    card: "Cartão",
    pix: "PIX",
    billet: "Boleto",
  };
  return methodLabels[method] || method;
};

const calculateTotal = (purchase: PurchaseListItem): string => {
  return formatCurrency(purchase.quantity_mwh * purchase.price_per_mwh);
};
</script>

<template>
  <div
    class="rounded-lg border border-neutral-200 bg-white shadow-sm sm:rounded-xl"
  >
    <div class="border-b border-neutral-200 p-4 sm:p-6">
      <h3 class="text-base font-semibold text-neutral-900 sm:text-lg">
        Histórico de Compras
      </h3>
      <p class="mt-1 text-xs text-neutral-500 sm:text-sm">
        Todas as transações concluidas nesta oferta
      </p>
    </div>

    <div class="hidden lg:block">
      <DataTable
        :value="purchasesWithTotal"
        :loading="loading"
        dataKey="uuid"
        stripedRows
        showGridlines
        sortMode="single"
        :sortField="'created_at'"
        :sortOrder="-1"
        class="text-sm"
      >
        <template #empty>
          <div class="py-12 text-center">
            <p class="text-neutral-500">Nenhuma compra realizada ainda</p>
          </div>
        </template>

        <template #loading>
          <div class="flex items-center justify-center py-12">
            <div class="flex flex-col items-center gap-3">
              <div
                class="h-8 w-8 animate-spin rounded-full border-4 border-emerald-500 border-t-transparent"
              ></div>
              <p class="text-sm text-neutral-500">Carregando compras...</p>
            </div>
          </div>
        </template>

        <Column field="created_at" header="Data" sortable>
          <template #body="slotProps">
            <span class="font-medium text-neutral-700">
              {{ formatDate(slotProps.data.created_at) }}
            </span>
          </template>
        </Column>

        <Column field="quantity_mwh" header="Quantidade" sortable>
          <template #body="slotProps">
            <span class="font-semibold text-neutral-900">
              {{ formatQuantity(slotProps.data.quantity_mwh) }}
            </span>
          </template>
        </Column>

        <Column field="price_per_mwh" header="Preço/MWh" sortable>
          <template #body="slotProps">
            <span class="text-neutral-700">
              {{ formatCurrency(slotProps.data.price_per_mwh) }}
            </span>
          </template>
        </Column>

        <Column field="total" header="Total" sortable>
          <template #body="slotProps">
            <span class="font-bold text-emerald-600">
              {{ calculateTotal(slotProps.data) }}
            </span>
          </template>
        </Column>

        <Column field="payment_method" header="Pagamento" sortable>
          <template #body="slotProps">
            <span class="text-neutral-600">
              {{ getPaymentMethodLabel(slotProps.data.payment_method) }}
            </span>
          </template>
        </Column>

        <Column field="status" header="Status">
          <template #body="slotProps">
            <span
              :class="[
                'inline-flex items-center gap-1.5 rounded-full px-3 py-1 text-xs font-medium',
                getStatusConfig(slotProps.data.status).bgColor,
                getStatusConfig(slotProps.data.status).textColor,
              ]"
            >
              <component
                :is="getStatusConfig(slotProps.data.status).icon"
                :size="14"
                :class="getStatusConfig(slotProps.data.status).iconColor"
              />
              {{ getStatusConfig(slotProps.data.status).label }}
            </span>
          </template>
        </Column>
      </DataTable>
    </div>

    <div class="lg:hidden">
      <div v-if="loading" class="flex items-center justify-center py-12">
        <div class="flex flex-col items-center gap-3">
          <div
            class="h-8 w-8 animate-spin rounded-full border-4 border-emerald-500 border-t-transparent"
          ></div>
          <p class="text-sm text-neutral-500">Carregando compras...</p>
        </div>
      </div>

      <div
        v-else-if="purchasesWithTotal.length === 0"
        class="py-12 text-center"
      >
        <p class="text-sm text-neutral-500">Nenhuma compra realizada ainda</p>
      </div>

      <div v-else class="divide-y divide-neutral-200">
        <div
          v-for="purchase in purchasesWithTotal"
          :key="purchase.uuid"
          class="p-4 transition-colors hover:bg-neutral-50"
        >
          <div class="mb-3 flex items-start justify-between">
            <div class="flex-1">
              <p class="text-xs font-medium text-neutral-500">Data da Compra</p>
              <p class="mt-0.5 text-sm font-semibold text-neutral-900">
                {{ formatDate(purchase.created_at) }}
              </p>
            </div>
            <span
              :class="[
                'inline-flex items-center gap-1.5 rounded-full px-2.5 py-1 text-xs font-medium',
                getStatusConfig(purchase.status).bgColor,
                getStatusConfig(purchase.status).textColor,
              ]"
            >
              <component
                :is="getStatusConfig(purchase.status).icon"
                :size="12"
                :class="getStatusConfig(purchase.status).iconColor"
              />
              {{ getStatusConfig(purchase.status).label }}
            </span>
          </div>

          <div class="grid grid-cols-2 gap-3">
            <div>
              <p class="text-xs font-medium text-neutral-500">Quantidade</p>
              <p class="mt-0.5 text-sm font-semibold text-neutral-900">
                {{ formatQuantity(purchase.quantity_mwh) }}
              </p>
            </div>

            <div>
              <p class="text-xs font-medium text-neutral-500">Preço/MWh</p>
              <p class="mt-0.5 text-sm text-neutral-700">
                {{ formatCurrency(purchase.price_per_mwh) }}
              </p>
            </div>

            <div>
              <p class="text-xs font-medium text-neutral-500">Pagamento</p>
              <p class="mt-0.5 text-sm text-neutral-700">
                {{ getPaymentMethodLabel(purchase.payment_method) }}
              </p>
            </div>

            <div>
              <p class="text-xs font-medium text-neutral-500">Total</p>
              <p class="mt-0.5 text-sm font-bold text-emerald-600">
                {{ calculateTotal(purchase) }}
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
:deep(.p-datatable .p-datatable-thead > tr > th) {
  background-color: #f9fafb;
  color: #374151;
  font-weight: 600;
  padding: 0.875rem;
  border-bottom: 2px solid #e5e7eb;
  font-size: 0.875rem;
}

@media (min-width: 1024px) {
  :deep(.p-datatable .p-datatable-thead > tr > th) {
    padding: 1rem;
  }
}

:deep(.p-datatable .p-datatable-tbody > tr > td) {
  padding: 0.75rem 0.875rem;
  font-size: 0.875rem;
}

@media (min-width: 1024px) {
  :deep(.p-datatable .p-datatable-tbody > tr > td) {
    padding: 0.875rem 1rem;
  }
}

:deep(.p-datatable .p-datatable-tbody > tr:hover) {
  background-color: #f9fafb;
}
</style>
