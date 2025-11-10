<script setup lang="ts">
import { computed, onMounted, ref } from "vue";
import { Trash2, Edit } from "lucide-vue-next";
import { useRouter } from "vue-router";
import api from "../../axios";
import { OFFER_ENDPOINTS } from "../../api/endpoints";
import type {
  OfferListItem,
  OffersListResponse,
} from "../../types/responses/offers";
import CreateOfferDialog from "./CreateOfferDialog.vue";
import EnergyOfferCard from "./EnergyOfferCard.vue";
import ConfirmDialog from "../shared/ConfirmDialog.vue";
import { RouteNames } from "../../router/types";

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

const router = useRouter();
const showCreateDialog = ref(false);
const showDeleteDialog = ref(false);
const offerToDelete = ref<string | null>(null);
const offerToEdit = ref<OfferListItem | null>(null);
const offers = ref<OfferListItem[]>([]);
const loading = ref(false);
const deleting = ref(false);
const error = ref<string | null>(null);

const currentPage = ref(1);
const hasNext = ref(false);
const hasPrev = ref(false);

const paginationInfo = computed<PaginationInfo>(() => {
  const start = (currentPage.value - 1) * PAGE_SIZE + 1;
  const end = start + offers.value.length - 1;
  return { start, end, total: offers.value.length };
});

const scrollToTop = () => {
  window.scrollTo({ top: 0, behavior: "smooth" });
};

const loadMyOffers = async () => {
  loading.value = true;
  error.value = null;

  try {
    const response = await api.get<OffersListResponse>(
      OFFER_ENDPOINTS.MY_OFFERS,
      {
        params: {
          page_size: PAGE_SIZE,
          page: currentPage.value,
        },
      },
    );

    offers.value = response.data.data;
    hasNext.value = response.data.has_next;
    hasPrev.value = response.data.has_prev;
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

const handleOfferCreated = () => {
  loadMyOffers();
};

const handleOfferUpdated = () => {
  loadMyOffers();
  offerToEdit.value = null;
};

const goToNextPage = () => {
  if (hasNext.value) {
    currentPage.value++;
    loadMyOffers();
    scrollToTop();
  }
};

const goToPrevPage = () => {
  if (hasPrev.value) {
    currentPage.value--;
    loadMyOffers();
    scrollToTop();
  }
};

const deleteOffer = async (offerUuid: string) => {
  offerToDelete.value = offerUuid;
  showDeleteDialog.value = true;
};

const confirmDelete = async () => {
  if (!offerToDelete.value) return;

  deleting.value = true;

  try {
    await api.delete(OFFER_ENDPOINTS.DELETE(offerToDelete.value));
    showDeleteDialog.value = false;
    offerToDelete.value = null;
    await loadMyOffers();
  } catch (err) {
    const apiError = err as ApiError;
    const serverMessage = apiError.response?.data?.message;
    const errorMessage = `Erro ao excluir oferta: ${serverMessage || "Tente novamente."}`;

    alert(errorMessage);
  } finally {
    deleting.value = false;
  }
};

const cancelDelete = () => {
  offerToDelete.value = null;
};

const viewOfferDetails = (offerUuid: string) => {
  router.push({
    name: RouteNames.OFFER_DETAIL,
    params: { id: offerUuid },
  });
};

onMounted(() => {
  loadMyOffers();
});
</script>

<template>
  <div class="flex flex-col gap-6">
    <div class="flex items-center justify-between">
      <div>
        <h2 class="text-2xl font-bold text-neutral-900">Minhas Ofertas</h2>
        <p class="mt-1 text-sm text-neutral-500">
          Gerencie todas as suas ofertas de energia
        </p>
      </div>
      <button
        @click="openCreateDialog"
        class="rounded-lg bg-gradient-to-r from-emerald-500 to-teal-600 px-6 py-2.5 text-sm font-medium text-white shadow-md transition-all duration-200 hover:shadow-lg hover:brightness-110"
      >
        Criar Nova Oferta
      </button>
    </div>

    <div v-if="loading" class="flex items-center justify-center py-12">
      <div class="flex flex-col items-center gap-3">
        <div
          class="h-10 w-10 animate-spin rounded-full border-4 border-emerald-500 border-t-transparent"
        ></div>
        <p class="text-sm text-neutral-500">Carregando suas ofertas...</p>
      </div>
    </div>

    <div
      v-else-if="error"
      class="rounded-lg border border-red-200 bg-red-50 p-6 text-center"
    >
      <p class="text-sm text-red-600">{{ error }}</p>
      <button
        @click="loadMyOffers"
        class="mt-3 rounded-lg bg-red-600 px-4 py-2 text-sm font-medium text-white hover:bg-red-700"
      >
        Tentar Novamente
      </button>
    </div>

    <div
      v-else-if="offers.length > 0"
      class="grid grid-cols-1 gap-6 lg:grid-cols-2 xl:grid-cols-3"
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
              @click="openEditDialog(offer)"
              class="flex flex-1 items-center justify-center gap-2 rounded-lg border-2 border-blue-500 bg-transparent py-2.5 text-sm font-medium text-blue-600 transition-all duration-200 hover:bg-blue-500 hover:text-white"
            >
              <Edit :size="16" />
              Editar
            </button>
            <button
              @click="deleteOffer(offer.uuid)"
              class="flex flex-1 items-center justify-center gap-2 rounded-lg border-2 border-red-500 bg-transparent py-2.5 text-sm font-medium text-red-600 transition-all duration-200 hover:bg-red-500 hover:text-white"
            >
              <Trash2 :size="16" />
              Excluir
            </button>
          </div>
        </template>
      </EnergyOfferCard>
    </div>

    <div
      v-if="offers.length > 0"
      class="flex items-center justify-between rounded-lg border border-neutral-200 bg-white px-6 py-4 shadow-sm"
    >
      <button
        @click="goToPrevPage"
        :disabled="!hasPrev"
        :class="[
          'flex items-center gap-2 rounded-lg px-4 py-2 text-sm font-medium transition-all duration-200',
          hasPrev
            ? 'bg-emerald-500 text-white hover:bg-emerald-600 hover:shadow-md'
            : 'cursor-not-allowed bg-neutral-100 text-neutral-400',
        ]"
      >
        <svg
          xmlns="http://www.w3.org/2000/svg"
          width="16"
          height="16"
          viewBox="0 0 24 24"
          fill="none"
          stroke="currentColor"
          stroke-width="2"
          stroke-linecap="round"
          stroke-linejoin="round"
        >
          <polyline points="15 18 9 12 15 6"></polyline>
        </svg>
        Anterior
      </button>

      <div class="flex flex-col items-center gap-1">
        <div class="flex items-center gap-2">
          <span class="text-sm font-medium text-neutral-600">Página</span>
          <span
            class="rounded-lg bg-emerald-50 px-3 py-1 text-sm font-bold text-emerald-600"
          >
            {{ currentPage }}
          </span>
        </div>
        <span class="text-xs text-neutral-500">
          Mostrando {{ paginationInfo.start }} -
          {{ paginationInfo.end }} ofertas
        </span>
      </div>

      <button
        @click="goToNextPage"
        :disabled="!hasNext"
        :class="[
          'flex items-center gap-2 rounded-lg px-4 py-2 text-sm font-medium transition-all duration-200',
          hasNext
            ? 'bg-emerald-500 text-white hover:bg-emerald-600 hover:shadow-md'
            : 'cursor-not-allowed bg-neutral-100 text-neutral-400',
        ]"
      >
        Próxima
        <svg
          xmlns="http://www.w3.org/2000/svg"
          width="16"
          height="16"
          viewBox="0 0 24 24"
          fill="none"
          stroke="currentColor"
          stroke-width="2"
          stroke-linecap="round"
          stroke-linejoin="round"
        >
          <polyline points="9 18 15 12 9 6"></polyline>
        </svg>
      </button>
    </div>

    <div
      v-else-if="!loading && offers.length === 0"
      class="rounded-lg border border-neutral-200 bg-neutral-50 p-12 text-center"
    >
      <p class="text-neutral-600">Você ainda não criou nenhuma oferta.</p>
      <button
        @click="openCreateDialog"
        class="mt-4 rounded-lg bg-emerald-500 px-4 py-2 text-sm font-medium text-white hover:bg-emerald-600"
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
      :loading="deleting"
      @confirm="confirmDelete"
      @cancel="cancelDelete"
    />
  </div>
</template>
