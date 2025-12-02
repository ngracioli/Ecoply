<script setup lang="ts">
import { computed, onMounted, ref } from "vue";
import { Filter } from "lucide-vue-next";
import { useRouter } from "vue-router";
import api from "../../../axios";
import { OFFER_ENDPOINTS } from "../../../api/endpoints";
import type {
  OfferListItem,
  OffersListResponse,
} from "../../../types/responses/offers";
import CreateOfferDialog from "../dialogs/create-offer/CreateOfferDialog.vue";
import EnergyOfferCard from "../cards/EnergyOfferCard.vue";
import OffersFilterDialog from "../dialogs/OffersFilterDialog.vue";
import Pagination from "../../shared/Pagination.vue";
import { useStore } from "vuex";
import type { RootState } from "../../../store";
import type { UserType as UserTypeEnum } from "../../../types/user";
import { RouteNames } from "../../../router/types";

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

    offers.value = response.data.data;
    hasNext.value = response.data.has_next;
    hasPrev.value = response.data.has_prev;
  } catch (err) {
    const apiError = err as ApiError;

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
  <div class="flex flex-col gap-4 sm:gap-6">
    <div
      class="flex flex-col gap-3 sm:flex-row sm:items-center sm:justify-between"
    >
      <div>
        <h2 class="text-xl font-bold text-neutral-900 sm:text-2xl">
          Ofertas de Energia
        </h2>
        <p class="mt-1 text-xs text-neutral-500 sm:text-sm">
          Encontre as melhores ofertas disponíveis
        </p>
      </div>
      <div class="flex flex-col gap-2 sm:flex-row sm:gap-3">
        <button
          @click="openFilterDialog"
          :class="[
            'flex min-h-[44px] items-center justify-center gap-2 rounded-lg px-4 py-3 text-xs font-medium shadow-md transition-all duration-200 hover:shadow-lg sm:min-h-0 sm:px-6 sm:py-2.5 sm:text-sm',
            hasActiveFilters
              ? 'bg-gradient-to-r from-emerald-500 to-teal-600 text-white hover:brightness-110 active:brightness-95'
              : 'border border-neutral-300 bg-white text-neutral-700 hover:bg-neutral-50 active:bg-neutral-100',
          ]"
        >
          <Filter :size="16" class="sm:hidden" />
          <Filter :size="18" class="hidden sm:block" />
          {{ hasActiveFilters ? "Filtros Ativos" : "Filtrar" }}
        </button>

        <button
          @click="openCreateDialog"
          class="min-h-[44px] rounded-lg bg-gradient-to-r from-emerald-500 to-teal-600 px-4 py-3 text-xs font-medium text-white shadow-md transition-all duration-200 hover:shadow-lg hover:brightness-110 active:brightness-95 sm:min-h-0 sm:px-6 sm:py-2.5 sm:text-sm"
          v-if="userUserType === 'supplier'"
        >
          Criar Nova Oferta
        </button>
      </div>
    </div>

    <div
      v-if="hasActiveFilters"
      class="flex flex-col gap-2 sm:flex-row sm:items-center"
    >
      <span class="shrink-0 text-xs text-neutral-600 sm:text-sm"
        >Filtros aplicados:</span
      >
      <div class="flex flex-wrap items-center gap-2">
        <span
          v-if="activeFilters.submarket"
          class="inline-flex items-center gap-1 rounded-full bg-emerald-100 px-2.5 py-1 text-xs font-medium text-emerald-700 sm:px-3"
        >
          Submercado: {{ activeFilters.submarket }}
        </span>
        <span
          v-if="activeFilters.energy_type"
          class="inline-flex items-center gap-1 rounded-full bg-emerald-100 px-2.5 py-1 text-xs font-medium text-emerald-700 sm:px-3"
        >
          Tipo: {{ getEnergyTypeLabel(activeFilters.energy_type) }}
        </span>
        <span
          v-if="activeFilters.period_start"
          class="inline-flex items-center gap-1 rounded-full bg-emerald-100 px-2.5 py-1 text-xs font-medium text-emerald-700 sm:px-3"
        >
          De: {{ activeFilters.period_start }}
        </span>
        <span
          v-if="activeFilters.period_end"
          class="inline-flex items-center gap-1 rounded-full bg-emerald-100 px-2.5 py-1 text-xs font-medium text-emerald-700 sm:px-3"
        >
          Até: {{ activeFilters.period_end }}
        </span>
        <button
          @click="clearFilters"
          class="min-h-[32px] rounded-full bg-red-100 px-2.5 py-1 text-xs font-medium text-red-700 transition-colors hover:bg-red-200 active:bg-red-300 sm:px-3"
        >
          Limpar filtros
        </button>
      </div>
    </div>

    <div v-if="loading" class="flex items-center justify-center py-8 sm:py-12">
      <div class="flex flex-col items-center gap-3">
        <div
          class="h-8 w-8 animate-spin rounded-full border-4 border-emerald-500 border-t-transparent sm:h-10 sm:w-10"
        ></div>
        <p class="text-xs text-neutral-500 sm:text-sm">Carregando ofertas...</p>
      </div>
    </div>

    <div
      v-else-if="error"
      class="rounded-lg border border-red-200 bg-red-50 p-4 text-center sm:p-6"
    >
      <p class="text-xs text-red-600 sm:text-sm">{{ error }}</p>
      <button
        @click="loadOffers"
        class="mt-3 min-h-[44px] rounded-lg bg-red-600 px-4 py-3 text-xs font-medium text-white hover:bg-red-700 active:bg-red-800 sm:min-h-0 sm:py-2 sm:text-sm"
      >
        Tentar Novamente
      </button>
    </div>

    <div
      v-else-if="offers.length > 0"
      class="grid grid-cols-1 gap-4 sm:gap-6 lg:grid-cols-2 xl:grid-cols-3"
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
      class="rounded-lg border border-neutral-200 bg-neutral-50 p-8 text-center sm:p-12"
    >
      <p class="text-sm text-neutral-600 sm:text-base">
        Nenhuma oferta disponível no momento.
      </p>
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
