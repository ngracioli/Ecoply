<script setup lang="ts">
import { computed, onMounted, ref } from "vue";
import { Filter } from "lucide-vue-next";
import { useRouter } from "vue-router";
import api from "../../axios";
import { OFFER_ENDPOINTS } from "../../api/endpoints";
import type {
  OfferListItem,
  OffersListResponse,
} from "../../types/responses/offers";
import CreateOfferDialog from "./CreateOfferDialog.vue";
import EnergyOfferCard from "./EnergyOfferCard.vue";
import OffersFilterDialog from "./OffersFilterDialog.vue";
import Pagination from "../shared/Pagination.vue";
import { useStore } from "vuex";
import type { RootState } from "../../store";
import type { UserType as UserTypeEnum } from "../../types/user";
import { RouteNames } from "../../router/types";

interface FilterOptions {
  submarket?: string;
  energy_type?: string;
  period_start?: string;
  period_end?: string;
}

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

const router = useRouter();
const store = useStore<RootState>();
const userUserType = computed<UserTypeEnum | null>(
  () => store.getters["user/userType"] || null,
);

const ENERGY_TYPE_LABELS: Record<string, string> = {
  solar: "Solar",
  eolic: "Eólica",
  hydroelectric: "Hidrelétrica",
  geothermal: "Geotérmica",
};

const PAGE_SIZE = 10;

const showCreateDialog = ref(false);
const showFilterDialog = ref(false);
const filterDialogRef = ref<InstanceType<typeof OffersFilterDialog> | null>(
  null,
);
const offers = ref<OfferListItem[]>([]);
const loading = ref(false);
const error = ref<string | null>(null);

const currentPage = ref(1);
const hasNext = ref(false);
const hasPrev = ref(false);
const activeFilters = ref<FilterOptions>({});

const paginationInfo = computed<PaginationInfo>(() => {
  const start = (currentPage.value - 1) * PAGE_SIZE + 1;
  const end = start + offers.value.length - 1;
  return { start, end, total: offers.value.length };
});

const hasActiveFilters = computed(() => {
  return Object.values(activeFilters.value).some(
    (value) => value && value.trim(),
  );
});

const getEnergyTypeLabel = (type: string): string => {
  return ENERGY_TYPE_LABELS[type] || type;
};

const scrollToTop = () => {
  window.scrollTo({ top: 0, behavior: "smooth" });
};

const loadOffers = async () => {
  loading.value = true;
  error.value = null;

  try {
    const response = await api.get<OffersListResponse>(OFFER_ENDPOINTS.LIST, {
      params: {
        page_size: PAGE_SIZE,
        page: currentPage.value,
        ...activeFilters.value,
      },
    });

    console.log("Resposta da API:", response.data);
    offers.value = response.data.data;
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
    error.value = `Erro ao carregar ofertas: ${errorMessage}`;
  } finally {
    loading.value = false;
  }
};

const openCreateDialog = () => {
  showCreateDialog.value = true;
};

const openFilterDialog = () => {
  showFilterDialog.value = true;
};

const handleOfferCreated = () => {
  loadOffers();
};

const handleApplyFilters = (filters: FilterOptions) => {
  activeFilters.value = filters;
  currentPage.value = 1;
  loadOffers();
};

const clearFilters = () => {
  activeFilters.value = {};
  currentPage.value = 1;

  if (filterDialogRef.value) {
    filterDialogRef.value.clearFilters();
  }

  loadOffers();
};

const goToNextPage = () => {
  if (hasNext.value) {
    currentPage.value++;
    loadOffers();
    scrollToTop();
  }
};

const goToPrevPage = () => {
  if (hasPrev.value) {
    currentPage.value--;
    loadOffers();
    scrollToTop();
  }
};

const viewOfferDetails = (offerUuid: string) => {
  router.push({
    name: RouteNames.OFFER_DETAIL,
    params: { id: offerUuid },
  });
};

onMounted(() => {
  loadOffers();
});
</script>

<template>
  <div class="flex flex-col gap-6">
    <div class="flex items-center justify-between">
      <div>
        <h2 class="text-2xl font-bold text-neutral-900">Ofertas de Energia</h2>
        <p class="mt-1 text-sm text-neutral-500">
          Encontre as melhores ofertas disponíveis
        </p>
      </div>
      <div class="flex gap-3">
        <button
          @click="openFilterDialog"
          :class="[
            'flex items-center gap-2 rounded-lg px-6 py-2.5 text-sm font-medium shadow-md transition-all duration-200 hover:shadow-lg',
            hasActiveFilters
              ? 'bg-gradient-to-r from-emerald-500 to-teal-600 text-white hover:brightness-110'
              : 'border border-neutral-300 bg-white text-neutral-700 hover:bg-neutral-50',
          ]"
        >
          <Filter :size="18" />
          {{ hasActiveFilters ? "Filtros Ativos" : "Filtrar" }}
        </button>

        <button
          @click="openCreateDialog"
          class="rounded-lg bg-gradient-to-r from-emerald-500 to-teal-600 px-6 py-2.5 text-sm font-medium text-white shadow-md transition-all duration-200 hover:shadow-lg hover:brightness-110"
          v-if="userUserType === 'supplier'"
        >
          Criar Nova Oferta
        </button>
      </div>
    </div>

    <div v-if="hasActiveFilters" class="flex items-center gap-2">
      <span class="text-sm text-neutral-600">Filtros aplicados:</span>
      <div class="flex flex-wrap items-center gap-2">
        <span
          v-if="activeFilters.submarket"
          class="inline-flex items-center gap-1 rounded-full bg-emerald-100 px-3 py-1 text-xs font-medium text-emerald-700"
        >
          Submercado: {{ activeFilters.submarket }}
        </span>
        <span
          v-if="activeFilters.energy_type"
          class="inline-flex items-center gap-1 rounded-full bg-emerald-100 px-3 py-1 text-xs font-medium text-emerald-700"
        >
          Tipo: {{ getEnergyTypeLabel(activeFilters.energy_type) }}
        </span>
        <span
          v-if="activeFilters.period_start"
          class="inline-flex items-center gap-1 rounded-full bg-emerald-100 px-3 py-1 text-xs font-medium text-emerald-700"
        >
          De: {{ activeFilters.period_start }}
        </span>
        <span
          v-if="activeFilters.period_end"
          class="inline-flex items-center gap-1 rounded-full bg-emerald-100 px-3 py-1 text-xs font-medium text-emerald-700"
        >
          Até: {{ activeFilters.period_end }}
        </span>
        <button
          @click="clearFilters"
          class="bg-red-100 px-3 py-1 text-xs font-medium text-red-700 transition-colors hover:bg-red-200"
        >
          Limpar filtros
        </button>
      </div>
    </div>

    <div v-if="loading" class="flex items-center justify-center py-12">
      <div class="flex flex-col items-center gap-3">
        <div
          class="h-10 w-10 animate-spin rounded-full border-4 border-emerald-500 border-t-transparent"
        ></div>
        <p class="text-sm text-neutral-500">Carregando ofertas...</p>
      </div>
    </div>

    <div
      v-else-if="error"
      class="rounded-lg border border-red-200 bg-red-50 p-6 text-center"
    >
      <p class="text-sm text-red-600">{{ error }}</p>
      <button
        @click="loadOffers"
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
        @click="viewOfferDetails(offer.uuid)"
      />
    </div>

    <Pagination
      v-if="offers.length > 0"
      :current-page="currentPage"
      :has-next="hasNext"
      :has-prev="hasPrev"
      :items-start="paginationInfo.start"
      :items-end="paginationInfo.end"
      item-label="ofertas"
      @prev="goToPrevPage"
      @next="goToNextPage"
    />

    <div
      v-else-if="!loading && offers.length === 0"
      class="rounded-lg border border-neutral-200 bg-neutral-50 p-12 text-center"
    >
      <p class="text-neutral-600">Nenhuma oferta disponível no momento.</p>
    </div>

    <CreateOfferDialog
      v-model:visible="showCreateDialog"
      @offer-created="handleOfferCreated"
    />

    <OffersFilterDialog
      ref="filterDialogRef"
      v-model:visible="showFilterDialog"
      @apply-filters="handleApplyFilters"
    />
  </div>
</template>

<style scoped></style>
