<script setup lang="ts">
import { computed, onMounted, ref, watch } from "vue";
import { useStore } from "vuex";
import api from "../../../axios";
import { PURCHASE_ENDPOINTS } from "../../../api/endpoints";
import type {
  PurchaseListItem,
  PurchasesListResponse,
  PurchaseContractResponse,
} from "../../../types/responses/purchases";
import type { RootState } from "../../../store";
import type { UserType as UserTypeEnum } from "../../../types/user";
import PurchaseCard from "../cards/PurchaseCard.vue";
import Pagination from "../../shared/Pagination.vue";
import ConfirmDialog from "../../shared/ConfirmDialog.vue";
import { generateContract } from "../../../utils/contractGenerator";

interface PaginationInfo {
  start: number;
  end: number;
  total: number;
}

interface ApiError {
  response?: {
    data?: {
      message?: string;
    };
  };
  message?: string;
}

const PAGE_SIZE = 10;

const store = useStore<RootState>();
const userUserType = computed<UserTypeEnum | null>(
  () => store.getters["user/userType"] || null,
);

const activeView = ref<"compras" | "vendas">("compras");
const purchases = ref<PurchaseListItem[]>([]);
const loading = ref(false);
const error = ref<string | null>(null);

const currentPage = ref(1);
const hasNext = ref(false);
const hasPrev = ref(false);

const statusFilter = ref<"" | "completed" | "waiting" | "canceled">("");
const paymentMethodFilter = ref<"" | "pix" | "card" | "billet">("");
const priceOrder = ref<"asc" | "desc">("desc");
const quantityOrder = ref<"asc" | "desc">("desc");

const showCancelDialog = ref(false);
const purchaseToCancel = ref<string | null>(null);
const cancelLoading = ref(false);
const downloadingContract = ref(false);

const paginationInfo = computed<PaginationInfo>(() => {
  const start = (currentPage.value - 1) * PAGE_SIZE + 1;
  const end = start + purchases.value.length - 1;
  return { start, end, total: purchases.value.length };
});

const scrollToTop = () => {
  window.scrollTo({ top: 0, behavior: "smooth" });
};

const loadBuyerPurchases = async () => {
  loading.value = true;
  error.value = null;

  try {
    const response = await api.get<PurchasesListResponse>(
      PURCHASE_ENDPOINTS.LIST,
      {
        params: {
          page: currentPage.value,
          page_size: PAGE_SIZE,
          status: statusFilter.value,
          payment_method: paymentMethodFilter.value,
          order_price: priceOrder.value,
          order_quantity: quantityOrder.value,
        },
      },
    );

    console.log("Resposta da API (Compras):", response.data);

    purchases.value = response.data.data;
    hasNext.value = response.data.has_next;
    hasPrev.value = response.data.has_prev;
  } catch (err) {
    const apiError = err as ApiError;
    console.error("Erro completo:", err);
    console.error("Response:", apiError.response);

    const errorMessage =
      apiError.response?.data?.message ||
      apiError.message ||
      "Tente novamente.";
    error.value = `Erro ao carregar compras: ${errorMessage}`;
  } finally {
    loading.value = false;
  }
};

const loadSellerPurchases = async () => {
  loading.value = true;
  error.value = null;

  try {
    const response = await api.get<PurchasesListResponse>(
      PURCHASE_ENDPOINTS.SALES,
      {
        params: {
          page: currentPage.value,
          page_size: PAGE_SIZE,
          status: statusFilter.value,
          payment_method: paymentMethodFilter.value,
          order_price: priceOrder.value,
          order_quantity: quantityOrder.value,
        },
      },
    );

    console.log("Resposta da API (Vendas):", response.data);

    purchases.value = response.data.data;
    hasNext.value = response.data.has_next;
    hasPrev.value = response.data.has_prev;
  } catch (err) {
    const apiError = err as ApiError;
    console.error("Erro completo:", err);
    console.error("Response:", apiError.response);

    const errorMessage =
      apiError.response?.data?.message ||
      apiError.message ||
      "Tente novamente.";
    error.value = `Erro ao carregar vendas: ${errorMessage}`;
  } finally {
    loading.value = false;
  }
};

const loadPurchases = async () => {
  if (activeView.value === "compras") {
    await loadBuyerPurchases();
  } else {
    await loadSellerPurchases();
  }
};

const goToNextPage = () => {
  if (hasNext.value) {
    currentPage.value++;
    loadPurchases();
    scrollToTop();
  }
};

const goToPrevPage = () => {
  if (hasPrev.value) {
    currentPage.value--;
    loadPurchases();
    scrollToTop();
  }
};

const handleCancelPurchase = (purchaseUuid: string) => {
  purchaseToCancel.value = purchaseUuid;
  showCancelDialog.value = true;
};

const confirmCancelPurchase = async () => {
  if (!purchaseToCancel.value) return;

  cancelLoading.value = true;

  try {
    await api.post(
      `${PURCHASE_ENDPOINTS.LIST}/${purchaseToCancel.value}/cancel`,
    );
    await loadPurchases();
    showCancelDialog.value = false;
    purchaseToCancel.value = null;
  } catch (err) {
    const apiError = err as ApiError;
    console.error("Erro ao cancelar compra:", err);

    const errorMessage =
      apiError.response?.data?.message ||
      apiError.message ||
      "Não foi possível cancelar a compra.";
    error.value = errorMessage;
  } finally {
    cancelLoading.value = false;
  }
};

const handleCancelDialogClose = () => {
  showCancelDialog.value = false;
  purchaseToCancel.value = null;
};

const handleDownloadContract = async (purchaseUuid: string) => {
  downloadingContract.value = true;

  try {
    const response = await api.get<PurchaseContractResponse>(
      `${PURCHASE_ENDPOINTS.LIST}/${purchaseUuid}/contract`,
    );

    const contractData = response.data.data;
    generateContract(contractData);
  } catch (err) {
    const apiError = err as ApiError;
    console.error("Erro ao baixar contrato:", err);

    const errorMessage =
      apiError.response?.data?.message ||
      apiError.message ||
      "Não foi possível baixar o contrato.";
    error.value = errorMessage;
  } finally {
    downloadingContract.value = false;
  }
};

watch(activeView, () => {
  currentPage.value = 1;
  statusFilter.value = "";
  paymentMethodFilter.value = "";
  priceOrder.value = "desc";
  quantityOrder.value = "desc";
  loadPurchases();
});

watch([statusFilter, paymentMethodFilter, priceOrder, quantityOrder], () => {
  currentPage.value = 1;
  loadPurchases();
});

onMounted(() => {
  loadPurchases();
});
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
    </div>

    <div v-if="userUserType === 'supplier'" class="flex justify-start">
      <div class="inline-flex rounded-lg bg-neutral-100 p-1">
        <button
          @click="activeView = 'compras'"
          :class="[
            'rounded-md px-6 py-2.5 text-sm font-medium transition-all duration-200',
            activeView === 'compras'
              ? 'bg-white text-emerald-600 shadow-sm'
              : 'text-neutral-600 hover:text-neutral-900',
          ]"
        >
          Minhas Compras
        </button>
        <button
          @click="activeView = 'vendas'"
          :class="[
            'rounded-md px-6 py-2.5 text-sm font-medium transition-all duration-200',
            activeView === 'vendas'
              ? 'bg-white text-emerald-600 shadow-sm'
              : 'text-neutral-600 hover:text-neutral-900',
          ]"
        >
          Minhas Vendas
        </button>
      </div>
    </div>

    <div class="rounded-lg border border-neutral-200 bg-white p-4 shadow-sm">
      <div class="grid grid-cols-1 gap-4 md:grid-cols-4">
        <div>
          <label
            for="status-filter"
            class="mb-1 block text-sm font-medium text-neutral-700"
          >
            Status
          </label>
          <select
            id="status-filter"
            v-model="statusFilter"
            class="w-full rounded-lg border border-neutral-300 px-3 py-2 text-sm focus:border-emerald-500 focus:ring-1 focus:ring-emerald-500 focus:outline-none"
          >
            <option value="">Todos</option>
            <option value="completed">Concluído</option>
            <option value="waiting">Aguardando</option>
            <option value="canceled">Cancelado</option>
          </select>
        </div>

        <div>
          <label
            for="payment-filter"
            class="mb-1 block text-sm font-medium text-neutral-700"
          >
            Método de Pagamento
          </label>
          <select
            id="payment-filter"
            v-model="paymentMethodFilter"
            class="w-full rounded-lg border border-neutral-300 px-3 py-2 text-sm focus:border-emerald-500 focus:ring-1 focus:ring-emerald-500 focus:outline-none"
          >
            <option value="">Todos</option>
            <option value="pix">PIX</option>
            <option value="card">Cartão</option>
            <option value="billet">Boleto</option>
          </select>
        </div>

        <div>
          <label
            for="price-order"
            class="mb-1 block text-sm font-medium text-neutral-700"
          >
            Ordenar por Preço
          </label>
          <select
            id="price-order"
            v-model="priceOrder"
            class="w-full rounded-lg border border-neutral-300 px-3 py-2 text-sm focus:border-emerald-500 focus:ring-1 focus:ring-emerald-500 focus:outline-none"
          >
            <option value="desc">Maior para Menor</option>
            <option value="asc">Menor para Maior</option>
          </select>
        </div>

        <div>
          <label
            for="quantity-order"
            class="mb-1 block text-sm font-medium text-neutral-700"
          >
            Ordenar por Quantidade
          </label>
          <select
            id="quantity-order"
            v-model="quantityOrder"
            class="w-full rounded-lg border border-neutral-300 px-3 py-2 text-sm focus:border-emerald-500 focus:ring-1 focus:ring-emerald-500 focus:outline-none"
          >
            <option value="desc">Maior para Menor</option>
            <option value="asc">Menor para Maior</option>
          </select>
        </div>
      </div>
    </div>

    <div class="flex flex-col gap-6">
      <div v-if="loading" class="flex items-center justify-center py-12">
        <div class="flex flex-col items-center gap-3">
          <div
            class="h-10 w-10 animate-spin rounded-full border-4 border-emerald-500 border-t-transparent"
          ></div>
          <p class="text-sm text-neutral-500">Carregando histórico...</p>
        </div>
      </div>

      <div
        v-else-if="error"
        class="rounded-lg border border-red-200 bg-red-50 p-6 text-center"
      >
        <p class="text-sm text-red-600">{{ error }}</p>
        <button
          @click="loadPurchases"
          class="mt-3 rounded-lg bg-red-600 px-4 py-2 text-sm font-medium text-white hover:bg-red-700"
        >
          Tentar Novamente
        </button>
      </div>

      <div
        v-else-if="purchases.length > 0"
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
                <th
                  v-if="activeView === 'compras'"
                  class="px-6 py-4 text-left text-xs font-semibold tracking-wider text-neutral-600 uppercase"
                >
                  Ações
                </th>
              </tr>
            </thead>
            <tbody class="divide-y divide-neutral-100">
              <PurchaseCard
                v-for="purchase in purchases"
                :key="purchase.uuid"
                :purchase="purchase"
                :show-cancel-button="activeView === 'compras'"
                :view-mode="activeView"
                @cancel="handleCancelPurchase"
                @download-contract="handleDownloadContract"
              />
            </tbody>
          </table>
        </div>
      </div>

      <Pagination
        v-if="purchases.length > 0"
        :current-page="currentPage"
        :has-next="hasNext"
        :has-prev="hasPrev"
        :items-start="paginationInfo.start"
        :items-end="paginationInfo.end"
        item-label="transações"
        @prev="goToPrevPage"
        @next="goToNextPage"
      />

      <div
        v-else-if="!loading && purchases.length === 0"
        class="rounded-lg border border-neutral-200 bg-neutral-50 p-12 text-center"
      >
        <p class="text-neutral-600">Nenhuma transação encontrada.</p>
      </div>
    </div>

    <ConfirmDialog
      v-model:visible="showCancelDialog"
      title="Cancelar Compra"
      message="Tem certeza que deseja cancelar esta compra? Esta ação não pode ser desfeita."
      confirm-text="Sim, cancelar"
      cancel-text="Não, manter"
      variant="danger"
      :loading="cancelLoading"
      @confirm="confirmCancelPurchase"
      @cancel="handleCancelDialogClose"
    />
  </div>
</template>
