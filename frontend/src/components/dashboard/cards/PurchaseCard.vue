<script setup lang="ts">
import type { Component } from "vue";
import { CheckCircle, Clock, XCircle, X, Download } from "lucide-vue-next";
import Button from "primevue/button";
import type { PurchaseListItem } from "../../../types/responses/purchases";

type TransactionStatus = "completed" | "waiting" | "canceled";

interface StatusConfig {
  icon: Component;
  label: string;
  bgColor: string;
  textColor: string;
  iconColor: string;
}

interface PurchaseCardProps {
  purchase: PurchaseListItem;
  showCancelButton?: boolean;
  viewMode?: "compras" | "vendas";
}

interface PurchaseCardEmits {
  (e: "cancel", uuid: string): void;
  (e: "download-contract", uuid: string): void;
}

const props = withDefaults(defineProps<PurchaseCardProps>(), {
  showCancelButton: true,
  viewMode: "compras",
});
const emit = defineEmits<PurchaseCardEmits>();

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

const formatTransactionId = (uuid: string): string => {
  return `#${uuid.slice(0, 8).toUpperCase()}`;
};

const formatDate = (dateString: string): string => {
  const date = new Date(dateString);
  return date.toLocaleDateString("pt-BR");
};

const formatCurrency = (value: number): string => {
  return new Intl.NumberFormat("pt-BR", {
    style: "currency",
    currency: "BRL",
  }).format(value);
};

const formatQuantity = (value: number): string => {
  return `${value.toLocaleString("pt-BR")} MWh`;
};

const calculateTotalAmount = (
  quantity: number,
  pricePerMwh: number,
): number => {
  return quantity * pricePerMwh;
};

const getPartnerName = (): string => {
  if (props.viewMode === "compras") {
    return props.purchase.seller_name || "Não informado";
  } else {
    return props.purchase.buyer_name || "Não informado";
  }
};

const getPaymentMethodLabel = (method: string): string => {
  const labels: Record<string, string> = {
    pix: "PIX",
    card: "Cartão",
    billet: "Boleto",
  };
  return labels[method] || method;
};
</script>

<template>
  <tr
    class="hidden transition-colors duration-150 hover:bg-neutral-50 md:table-row"
  >
    <td class="px-6 py-4">
      <div class="font-medium text-neutral-900">
        {{ props.viewMode === "compras" ? "Compra" : "Venda" }} de Energia
      </div>
      <div class="text-sm text-neutral-500">
        ID: {{ formatTransactionId(purchase.uuid) }}
      </div>
      <div class="mt-1 text-xs text-neutral-400">
        {{ getPaymentMethodLabel(purchase.payment_method) }}
      </div>
    </td>
    <td class="px-6 py-4">
      <div class="text-sm text-neutral-700">
        {{ getPartnerName() }}
      </div>
      <div class="text-xs text-neutral-500">
        {{ props.viewMode === "compras" ? "Vendedor" : "Comprador" }}
      </div>
    </td>
    <td class="px-6 py-4">
      <div class="text-sm font-medium text-neutral-900">
        {{ formatQuantity(purchase.quantity_mwh) }}
      </div>
    </td>
    <td class="px-6 py-4">
      <div class="text-sm font-semibold text-neutral-900">
        {{
          formatCurrency(
            calculateTotalAmount(purchase.quantity_mwh, purchase.price_per_mwh),
          )
        }}
      </div>
      <div class="text-xs text-neutral-500">
        {{ formatCurrency(purchase.price_per_mwh) }}/MWh
      </div>
    </td>
    <td class="px-6 py-4">
      <div class="text-sm text-neutral-600">
        {{ formatDate(purchase.created_at) }}
      </div>
    </td>
    <td class="px-6 py-4">
      <span
        :class="[
          'inline-flex items-center gap-1.5 rounded-full px-3 py-1 text-xs font-medium',
          getStatusConfig(purchase.status).bgColor,
          getStatusConfig(purchase.status).textColor,
        ]"
      >
        <component
          :is="getStatusConfig(purchase.status).icon"
          :size="14"
          :class="getStatusConfig(purchase.status).iconColor"
        />
        {{ getStatusConfig(purchase.status).label }}
      </span>
    </td>
    <td v-if="props.showCancelButton" class="px-6 py-4">
      <Button
        v-if="purchase.status === 'completed'"
        @click="emit('download-contract', purchase.uuid)"
        label="Baixar Contrato"
        severity="success"
        size="small"
        class="action-button download-button"
      >
        <template #icon>
          <Download :size="16" />
        </template>
      </Button>
      <Button
        v-else
        @click="emit('cancel', purchase.uuid)"
        :disabled="purchase.status === 'canceled'"
        label="Cancelar"
        severity="danger"
        size="small"
        class="action-button cancel-button"
      >
        <template #icon>
          <X :size="16" />
        </template>
      </Button>
    </td>
  </tr>

  <tr class="mb-4 block md:hidden">
    <td colspan="100%" class="block p-0">
      <div
        class="mx-4 overflow-hidden rounded-lg border border-neutral-200 bg-white shadow-sm"
      >
        <div class="space-y-3 p-4">
          <div class="flex items-start justify-between gap-3">
            <div class="min-w-0 flex-1">
              <div class="text-sm font-medium text-neutral-900">
                {{ props.viewMode === "compras" ? "Compra" : "Venda" }} de
                Energia
              </div>
              <div class="mt-0.5 text-xs text-neutral-500">
                ID: {{ formatTransactionId(purchase.uuid) }}
              </div>
            </div>
            <span
              :class="[
                'inline-flex shrink-0 items-center gap-1 rounded-full px-2.5 py-1 text-xs font-medium',
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

          <div class="grid grid-cols-2 gap-3 text-sm">
            <div>
              <div class="mb-1 text-xs text-neutral-500">
                {{ props.viewMode === "compras" ? "Vendedor" : "Comprador" }}
              </div>
              <div class="font-medium break-words text-neutral-900">
                {{ getPartnerName() }}
              </div>
            </div>
            <div>
              <div class="mb-1 text-xs text-neutral-500">Quantidade</div>
              <div class="font-medium text-neutral-900">
                {{ formatQuantity(purchase.quantity_mwh) }}
              </div>
            </div>
          </div>

          <div class="space-y-2 border-t border-neutral-100 pt-3">
            <div class="flex items-center justify-between">
              <span class="text-xs text-neutral-500">Valor Total</span>
              <span class="text-base font-bold text-neutral-900">
                {{
                  formatCurrency(
                    calculateTotalAmount(
                      purchase.quantity_mwh,
                      purchase.price_per_mwh,
                    ),
                  )
                }}
              </span>
            </div>
            <div class="flex items-center justify-between">
              <span class="text-xs text-neutral-500">Preço/MWh</span>
              <span class="text-xs text-neutral-600">
                {{ formatCurrency(purchase.price_per_mwh) }}
              </span>
            </div>
            <div class="flex items-center justify-between">
              <span class="text-xs text-neutral-500">Data</span>
              <span class="text-xs text-neutral-600">
                {{ formatDate(purchase.created_at) }}
              </span>
            </div>
            <div class="flex items-center justify-between">
              <span class="text-xs text-neutral-500">Pagamento</span>
              <span class="text-xs font-medium text-neutral-600">
                {{ getPaymentMethodLabel(purchase.payment_method) }}
              </span>
            </div>
          </div>

          <div
            v-if="props.showCancelButton"
            class="border-t border-neutral-100 pt-3"
          >
            <Button
              v-if="purchase.status === 'completed'"
              @click="emit('download-contract', purchase.uuid)"
              label="Baixar Contrato"
              severity="success"
              size="small"
              class="mobile-action-button download-button w-full touch-manipulation"
            >
              <template #icon>
                <Download :size="16" />
              </template>
            </Button>
            <Button
              v-else
              @click="emit('cancel', purchase.uuid)"
              :disabled="purchase.status === 'canceled'"
              label="Cancelar"
              severity="danger"
              size="small"
              class="mobile-action-button cancel-button w-full touch-manipulation"
            >
              <template #icon>
                <X :size="16" />
              </template>
            </Button>
          </div>
        </div>
      </div>
    </td>
  </tr>
</template>

<style scoped>
.action-button {
  min-width: 160px !important;
  width: 160px !important;
}

.mobile-action-button {
  min-height: 44px !important;
  justify-content: center !important;
}

.cancel-button {
  background-color: #ef4444 !important;
  border-color: #ef4444 !important;
  color: white !important;
}

.cancel-button:hover {
  background-color: #dc2626 !important;
  border-color: #dc2626 !important;
}

.cancel-button:active {
  transform: scale(0.98);
}

.cancel-button:disabled {
  opacity: 0.5;
  cursor: not-allowed !important;
  background-color: #f5f5f5 !important;
  border-color: #d4d4d4 !important;
  color: #a3a3a3 !important;
}

.cancel-button:disabled:hover {
  background-color: #f5f5f5 !important;
  border-color: #d4d4d4 !important;
  color: #a3a3a3 !important;
}

.cancel-button:disabled:active {
  transform: none;
}

.download-button {
  background-color: #10b981 !important;
  border-color: #10b981 !important;
  color: white !important;
}

.download-button:hover {
  background-color: #059669 !important;
  border-color: #059669 !important;
}

.download-button:active {
  transform: scale(0.98);
}
</style>
