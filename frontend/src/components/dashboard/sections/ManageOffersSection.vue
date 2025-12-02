<script setup lang="ts">
import { onMounted, ref } from "vue";
import { Trash2, Edit } from "lucide-vue-next";
import { useRouter } from "vue-router";
import api from "../../../axios";
import { OFFER_ENDPOINTS } from "../../../api/endpoints";
import type {
  OfferListItem,
  OffersListResponse,
} from "../../../types/responses/offers";
import CreateOfferDialog from "../dialogs/create-offer/CreateOfferDialog.vue";
import EnergyOfferCard from "../cards/EnergyOfferCard.vue";
import ConfirmDialog from "../../shared/ConfirmDialog.vue";
import DeleteOfferModal from "../dialogs/DeleteOfferModal.vue";
import { RouteNames } from "../../../router/types";

interface ApiError {
  response?: {
    data?: {
      message?: string;
    };
  };
  message?: string;
}

const router = useRouter();
const showCreateDialog = ref(false);
const showDeleteDialog = ref(false);
const offerToDelete = ref<string | null>(null);
const offerToEdit = ref<OfferListItem | null>(null);
const offers = ref<OfferListItem[]>([]);
const loading = ref(false);
const error = ref<string | null>(null);
const showDeleteErrorModal = ref(false);

const loadMyOffers = async () => {
  loading.value = true;
  error.value = null;

  try {
    const response = await api.get<OffersListResponse>(
      OFFER_ENDPOINTS.MY_OFFERS,
    );

    offers.value = response.data.data;
  } catch (err) {
    const apiError = err as ApiError;
    const errorMessage = apiError.message || "Erro desconhecido";
    const serverMessage = apiError.response?.data?.message;

    error.value = `Erro ao carregar suas ofertas: ${serverMessage || errorMessage}`;
  } finally {
    loading.value = false;
  }
};

const openCreateDialog = () => {
  offerToEdit.value = null;
  showCreateDialog.value = true;
};

const openEditDialog = (offer: OfferListItem) => {
  offerToEdit.value = offer;
  showCreateDialog.value = true;
};

const isOfferEditable = (offer: OfferListItem): boolean => {
  return offer.status === "fresh";
};

const handleOfferCreated = () => {
  loadMyOffers();
};

const handleOfferUpdated = () => {
  loadMyOffers();
  offerToEdit.value = null;
};

const deleteOffer = async (offerUuid: string) => {
  offerToDelete.value = offerUuid;
  showDeleteDialog.value = true;
};

const confirmDelete = async () => {
  if (!offerToDelete.value) return;

  showDeleteDialog.value = false;

  try {
    await api.delete(OFFER_ENDPOINTS.DELETE(offerToDelete.value));
    offerToDelete.value = null;
    await loadMyOffers();
  } catch (err) {
    showDeleteErrorModal.value = true;
  }
};

const cancelDelete = () => {
  offerToDelete.value = null;
  showDeleteDialog.value = false;
};

const closeDeleteErrorModal = () => {
  showDeleteErrorModal.value = false;
  offerToDelete.value = null;
};

const viewOfferDetails = (offerUuid: string) => {
  router.push({
    name: RouteNames.MANAGE_OFFER,
    params: { id: offerUuid },
  });
};

onMounted(() => {
  loadMyOffers();
});
</script>

<template>
  <div class="flex flex-col gap-4 sm:gap-6">
    <div
      class="flex flex-col gap-3 sm:flex-row sm:items-center sm:justify-between"
    >
      <div>
        <h2 class="text-xl font-bold text-neutral-900 sm:text-2xl">
          Minhas Ofertas
        </h2>
        <p class="mt-1 text-xs text-neutral-500 sm:text-sm">
          Gerencie todas as suas ofertas de energia
        </p>
      </div>
      <button
        @click="openCreateDialog"
        class="min-h-[44px] shrink-0 rounded-lg bg-gradient-to-r from-emerald-500 to-teal-600 px-4 py-3 text-xs font-medium text-white shadow-md transition-all duration-200 hover:shadow-lg hover:brightness-110 active:brightness-95 sm:min-h-0 sm:px-6 sm:py-2.5 sm:text-sm"
      >
        Criar Nova Oferta
      </button>
    </div>

    <div v-if="loading" class="flex items-center justify-center py-8 sm:py-12">
      <div class="flex flex-col items-center gap-3">
        <div
          class="h-8 w-8 animate-spin rounded-full border-4 border-emerald-500 border-t-transparent sm:h-10 sm:w-10"
        ></div>
        <p class="text-xs text-neutral-500 sm:text-sm">
          Carregando suas ofertas...
        </p>
      </div>
    </div>

    <div
      v-else-if="error"
      class="rounded-lg border border-red-200 bg-red-50 p-4 text-center sm:p-6"
    >
      <p class="text-xs text-red-600 sm:text-sm">{{ error }}</p>
      <button
        @click="loadMyOffers"
        class="mt-3 min-h-[44px] rounded-lg bg-red-600 px-4 py-3 text-xs font-medium text-white hover:bg-red-700 active:bg-red-800 sm:min-h-0 sm:py-2 sm:text-sm"
      >
        Tentar Novamente
      </button>
    </div>

    <div v-else-if="offers.length > 0" class="flex flex-col gap-4 sm:gap-6">
      <div
        class="grid grid-cols-1 gap-4 sm:gap-6 lg:grid-cols-2 xl:grid-cols-3"
      >
        <EnergyOfferCard
          v-for="offer in offers"
          :key="offer.uuid"
          :offer="offer"
          actionButtonText="Visualizar Oferta"
          @click="viewOfferDetails(offer.uuid)"
        >
          <template #actions>
            <div class="flex gap-2">
              <button
                v-tooltip.top="{
                  value: isOfferEditable(offer)
                    ? ''
                    : 'Você só pode editar ofertas que ainda não tiveram interações',
                  showDelay: 300,
                  hideDelay: 100,
                  class: 'tooltip-info',
                  disabled: isOfferEditable(offer),
                }"
                @click="openEditDialog(offer)"
                :disabled="!isOfferEditable(offer)"
                :class="[
                  'flex min-h-[44px] flex-1 items-center justify-center gap-1.5 rounded-lg border-2 py-3 text-xs font-medium transition-all duration-200 sm:min-h-0 sm:gap-2 sm:py-2.5 sm:text-sm',
                  isOfferEditable(offer)
                    ? 'border-blue-500 bg-transparent text-blue-600 hover:bg-blue-500 hover:text-white active:bg-blue-600'
                    : 'cursor-not-allowed border-neutral-300 bg-neutral-100 text-neutral-400',
                ]"
              >
                <Edit :size="14" class="sm:hidden" />
                <Edit :size="16" class="hidden sm:block" />
                <span>Editar</span>
              </button>
              <button
                v-tooltip.top="{
                  value: isOfferEditable(offer)
                    ? ''
                    : 'Você só pode excluir ofertas que ainda não tiveram interações',
                  showDelay: 300,
                  hideDelay: 100,
                  class: 'tooltip-info',
                  disabled: isOfferEditable(offer),
                }"
                @click="deleteOffer(offer.uuid)"
                :disabled="!isOfferEditable(offer)"
                :class="[
                  'flex min-h-[44px] flex-1 items-center justify-center gap-1.5 rounded-lg border-2 py-3 text-xs font-medium transition-all duration-200 sm:min-h-0 sm:gap-2 sm:py-2.5 sm:text-sm',
                  isOfferEditable(offer)
                    ? 'border-red-500 bg-transparent text-red-600 hover:bg-red-500 hover:text-white active:bg-red-600'
                    : 'cursor-not-allowed border-neutral-300 bg-neutral-100 text-neutral-400',
                ]"
              >
                <Trash2 :size="14" class="sm:hidden" />
                <Trash2 :size="16" class="hidden sm:block" />
                <span>Excluir</span>
              </button>
            </div>
          </template>
        </EnergyOfferCard>
      </div>
    </div>

    <div
      v-else-if="!loading && offers.length === 0"
      class="rounded-lg border border-neutral-200 bg-neutral-50 p-8 text-center sm:p-12"
    >
      <p class="text-sm text-neutral-600 sm:text-base">
        Você ainda não criou nenhuma oferta.
      </p>
      <button
        @click="openCreateDialog"
        class="mt-4 min-h-[44px] rounded-lg bg-emerald-500 px-4 py-3 text-xs font-medium text-white hover:bg-emerald-600 active:bg-emerald-700 sm:min-h-0 sm:px-4 sm:py-2 sm:text-sm"
      >
        Criar Primeira Oferta
      </button>
    </div>

    <CreateOfferDialog
      v-model:visible="showCreateDialog"
      :offer="offerToEdit"
      @offer-created="handleOfferCreated"
      @offer-updated="handleOfferUpdated"
    />

    <ConfirmDialog
      v-model:visible="showDeleteDialog"
      title="Excluir Oferta"
      message="Tem certeza que deseja excluir esta oferta? Esta ação não pode ser desfeita."
      confirm-text="Excluir"
      cancel-text="Cancelar"
      variant="danger"
      @confirm="confirmDelete"
      @cancel="cancelDelete"
    />

    <DeleteOfferModal
      :visible="showDeleteErrorModal"
      @close="closeDeleteErrorModal"
    />
  </div>
</template>

<style>
.tooltip-info .p-tooltip-text {
  background-color: rgb(240 255 250) !important;
  color: rgb(21 128 61) !important;
  border: 1px solid rgb(200 247 218) !important;
  box-shadow: 0 4px 12px rgba(34, 197, 94, 0.15) !important;
  font-size: 0.875rem !important;
  padding: 0.5rem 0.75rem !important;
}

.tooltip-info .p-tooltip-arrow {
  border-top-color: rgb(240 255 250) !important;
}
</style>
