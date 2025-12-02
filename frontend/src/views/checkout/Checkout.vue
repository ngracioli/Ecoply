<script setup lang="ts">
import { ref, onMounted } from "vue";
import { useRoute, useRouter } from "vue-router";
import api from "../../axios";
import { OFFER_ENDPOINTS, PURCHASE_ENDPOINTS } from "../../api/endpoints";
import type {
  OfferListItem,
  OfferDetailResponse,
} from "../../types/responses/offers";
import type { PurchaseContractResponse } from "../../types/responses/purchases";
import { ArrowLeft } from "lucide-vue-next";
import PurchaseForm from "../../components/dashboard/purchases/PurchaseForm.vue";
import PaymentProcessingModal from "../../components/dashboard/purchases/PaymentProcessingModal.vue";
import ContractDownloadDialog from "../../components/dashboard/purchases/ContractDownloadDialog.vue";
import ConfirmDialog from "../../components/shared/ConfirmDialog.vue";
import { generateContract } from "../../utils/contractGenerator";
import type { PurchaseFormData } from "../../components/dashboard/purchases/PurchaseForm.vue";

const route = useRoute();
const router = useRouter();
const offer = ref<OfferListItem | null>(null);
const loading = ref(true);
const processing = ref(false);
const error = ref<string | null>(null);
const paymentStatus = ref<"idle" | "processing" | "success" | "error">("idle");
const paymentError = ref<string>("");
const showContractDialog = ref(false);
const showPendingDialog = ref(false);
const purchaseFormData = ref<PurchaseFormData | null>(null);
const purchaseUuid = ref<string>("");
const contractDataFetched = ref(false);

const offerId = route.params.id as string;

const fetchOffer = async () => {
  try {
    loading.value = true;
    error.value = null;
    const response = await api.get<OfferDetailResponse>(
      OFFER_ENDPOINTS.DETAIL(offerId),
    );
    offer.value = response.data.data;
  } catch (err: any) {
    error.value =
      err.response?.data?.message || "Erro ao carregar os detalhes da oferta";
  } finally {
    loading.value = false;
  }
};

onMounted(() => {
  fetchOffer();
});

const goBack = () => {
  router.back();
};

const processPayment = async (formData: PurchaseFormData) => {
  paymentStatus.value = "processing";
  processing.value = true;
  purchaseFormData.value = formData;

  try {
    const payload = {
      quantity_mwh: formData.quantityMwh,
      payment_method: formData.paymentMethod,
    };

    const [purchaseResponse] = await Promise.all([
      api.post<{ data?: { uuid?: string } }>(
        OFFER_ENDPOINTS.PURCHASE(offerId),
        payload,
      ),
      new Promise((resolve) => setTimeout(resolve, 3500)),
    ]);

    if (purchaseResponse.data?.data?.uuid) {
      purchaseUuid.value = purchaseResponse.data.data.uuid;
    } else if ((purchaseResponse.data as any)?.uuid) {
      purchaseUuid.value = (purchaseResponse.data as any).uuid;
    }

    paymentStatus.value = "success";

    setTimeout(() => {
      paymentStatus.value = "idle";
      processing.value = false;
      setTimeout(() => {
        if (formData.paymentMethod === "pix") {
          showContractDialog.value = true;
        } else {
          showPendingDialog.value = true;
        }
      }, 300);
    }, 2000);
  } catch (err: any) {
    await new Promise((resolve) => setTimeout(resolve, 3500));

    paymentStatus.value = "error";
    paymentError.value =
      err.response?.data?.message || "Erro ao processar a compra";

    setTimeout(() => {
      paymentStatus.value = "idle";
      processing.value = false;
    }, 3000);
  }
};

const handleClosePaymentModal = () => {
  paymentStatus.value = "idle";
  processing.value = false;
};

const handleDownloadContract = async () => {
  if (contractDataFetched.value) {
    showContractDialog.value = false;
    router.push({ name: "Dashboard" });
    return;
  }

  if (purchaseUuid.value) {
    try {
      const response = await api.get<PurchaseContractResponse>(
        PURCHASE_ENDPOINTS.CONTRACT(purchaseUuid.value),
      );

      generateContract(response.data.data);
      contractDataFetched.value = true;

      showContractDialog.value = false;
      router.push({ name: "Dashboard" });
      return;
    } catch (err: any) {
      paymentError.value =
        err.response?.data?.message || "Erro ao buscar dados do contrato";
    }
  }

  showContractDialog.value = false;
  router.push({ name: "Dashboard" });
};

const handleSkipContract = () => {
  showContractDialog.value = false;
  router.push({ name: "Dashboard" });
};

const handlePendingClose = () => {
  showPendingDialog.value = false;
  router.push({ name: "Dashboard" });
};
</script>

<template>
  <div class="min-h-dvh bg-neutral-50">
    <header class="border-b border-neutral-200 bg-white">
      <div class="mx-auto max-w-7xl px-3 py-4 sm:px-6 sm:py-6 lg:px-8">
        <div class="flex items-center gap-2 sm:gap-4">
          <button
            @click="goBack"
            class="rounded-lg p-1.5 text-neutral-600 transition-colors hover:bg-neutral-100 sm:p-2"
          >
            <ArrowLeft :size="20" class="sm:hidden" />
            <ArrowLeft :size="24" class="hidden sm:block" />
          </button>
          <h1
            class="text-lg font-semibold text-neutral-900 sm:text-xl lg:text-2xl"
          >
            Finalizar Compra
          </h1>
        </div>
      </div>
    </header>

    <main class="mx-auto max-w-7xl px-3 py-4 sm:px-6 sm:py-8 lg:px-8">
      <div
        v-if="loading"
        class="flex items-center justify-center py-12 sm:py-20"
      >
        <div class="text-center">
          <div
            class="mx-auto h-10 w-10 animate-spin rounded-full border-4 border-neutral-200 border-t-emerald-500 sm:h-12 sm:w-12"
          ></div>
          <p class="mt-3 text-sm text-neutral-600 sm:mt-4 sm:text-base">
            Carregando...
          </p>
        </div>
      </div>

      <div
        v-else-if="error"
        class="rounded-lg border border-red-200 bg-red-50 p-4 text-center sm:rounded-xl sm:p-6"
      >
        <p class="text-sm text-red-700 sm:text-base">{{ error }}</p>
        <button
          @click="fetchOffer"
          class="mt-3 rounded-lg bg-red-600 px-4 py-2 text-sm text-white transition-colors hover:bg-red-700 sm:mt-4 sm:text-base"
        >
          Tentar novamente
        </button>
      </div>

      <PurchaseForm
        v-else-if="offer"
        :offer="offer"
        :processing="processing"
        @submit="processPayment"
      />
    </main>

    <PaymentProcessingModal
      :status="paymentStatus"
      :error-message="paymentError"
      @close="handleClosePaymentModal"
    />

    <ContractDownloadDialog
      v-model:visible="showContractDialog"
      @download="handleDownloadContract"
      @skip="handleSkipContract"
    />

    <ConfirmDialog
      v-model:visible="showPendingDialog"
      title="Compra Registrada com Sucesso!"
      message="Sua compra está pendente aguardando a confirmação do pagamento. O contrato estará disponível para download na aba 'Histórico de Negociações' assim que a compra for concluída."
      confirm-text="Ir para Dashboard"
      variant="info"
      @confirm="handlePendingClose"
    />
  </div>
</template>
