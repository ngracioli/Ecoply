<script setup lang="ts">
import { computed, onMounted, ref } from "vue";
import { useRoute, useRouter } from "vue-router";
import { ArrowLeft, Edit, Trash2, Calendar } from "lucide-vue-next";
import api from "../../axios";
import { OFFER_ENDPOINTS } from "../../api/endpoints";
import type { Offer } from "../../types/responses/offer";
import type { PurchaseListItem } from "../../types/responses/purchases";
import type { EnergyType } from "../../types/offer";
import OfferStatsCards from "../../components/dashboard/cards/OfferStatsCards.vue";
import OfferPurchasesTable from "../../components/dashboard/tables/OfferPurchasesTable.vue";
import CreateOfferDialog from "../../components/dashboard/dialogs/create-offer/CreateOfferDialog.vue";
import ConfirmDialog from "../../components/shared/ConfirmDialog.vue";
import DeleteOfferModal from "../../components/dashboard/dialogs/DeleteOfferModal.vue";
import ProgressBar from "../../components/shared/ProgressBar.vue";
import SubmarketBadge from "../../components/shared/SubmarketBadge.vue";
import { RouteNames } from "../../router/types";
import type { OfferListItem } from "../../types/responses/offers";

interface EnergyTypeConfig {
  gradient: string;
  badge: string;
  label: string;
}

interface ApiError {
  response?: {
    data?: {
      message?: string;
    };
  };
  message?: string;
}

const route = useRoute();
const router = useRouter();

const offer = ref<Offer | null>(null);
const loading = ref(true);
const error = ref<string | null>(null);
const showEditDialog = ref(false);
const showDeleteDialog = ref(false);
const showDeleteErrorModal = ref(false);
const purchasesLoading = ref(false);
const purchases = ref<PurchaseListItem[]>([]);

const ENERGY_TYPE_CONFIG: Record<EnergyType, EnergyTypeConfig> = {
  solar: {
    gradient: "from-yellow-400 to-orange-500",
    badge: "bg-yellow-100 text-yellow-700",
    label: "Solar",
  },
  eolic: {
    gradient: "from-sky-400 to-blue-500",
    badge: "bg-sky-100 text-sky-700",
    label: "Eólica",
  },
  hydroelectric: {
    gradient: "from-blue-500 to-indigo-600",
    badge: "bg-blue-100 text-blue-700",
    label: "Hidrelétrica",
  },
  geothermal: {
    gradient: "from-orange-500 to-red-600",
    badge: "bg-orange-100 text-orange-700",
    label: "Geotérmica",
  },
};

const DEFAULT_CONFIG: EnergyTypeConfig = {
  gradient: "from-gray-400 to-gray-600",
  badge: "bg-gray-100 text-gray-700",
  label: "",
};

const normalizeEnergyType = (type: string): EnergyType | null => {
  const normalized = type?.toLowerCase() || "";
  return normalized in ENERGY_TYPE_CONFIG ? (normalized as EnergyType) : null;
};

const getEnergyTypeConfig = (type: string): EnergyTypeConfig => {
  const normalizedType = normalizeEnergyType(type);
  return normalizedType ? ENERGY_TYPE_CONFIG[normalizedType] : DEFAULT_CONFIG;
};

const typeConfig = computed(() => {
  return offer.value
    ? getEnergyTypeConfig(offer.value.energy_type)
    : DEFAULT_CONFIG;
});

const formatPrice = (price: number): string => {
  return `R$ ${price.toFixed(2).replace(".", ",")}`;
};

const formatQuantity = (quantityMwh: number): string => {
  return `${quantityMwh.toLocaleString("pt-BR")} MWh`;
};

const formatDate = (dateString: string): string => {
  const [year, month, day] = dateString.split("-");
  return `${day}/${month}/${year}`;
};

const stats = computed(() => {
  if (!offer.value) {
    return {
      totalSales: 0,
      totalRevenue: 0,
      totalQuantitySold: 0,
      remainingQuantity: 0,
    };
  }

  const completedPurchases = purchases.value.filter(
    (p) => p.status === "completed",
  );
  const totalSales = completedPurchases.length;
  const totalQuantitySold = completedPurchases.reduce(
    (acc, p) => acc + p.quantity_mwh,
    0,
  );
  const totalRevenue = completedPurchases.reduce(
    (acc, p) => acc + p.quantity_mwh * p.price_per_mwh,
    0,
  );

  return {
    totalSales,
    totalRevenue,
    totalQuantitySold,
    remainingQuantity: offer.value.remaining_quantity_mwh,
  };
});

const offerForEdit = computed((): OfferListItem | null => {
  if (!offer.value) return null;

  return {
    uuid: offer.value.uuid,
    price_per_mwh: offer.value.price_per_mwh,
    initial_quantity_mwh: offer.value.initial_quantity_mwh,
    remaining_quantity_mwh: offer.value.remaining_quantity_mwh,
    description: offer.value.description,
    period_start: offer.value.period_start,
    period_end: offer.value.period_end,
    status: offer.value.status,
    energy_type: offer.value.energy_type,
    submarket: offer.value.submarket,
    seller_agent_uuid: offer.value.seller_agent_uuid,
    created_at: offer.value.created_at,
  };
});

const loadOffer = async () => {
  loading.value = true;
  error.value = null;

  try {
    const offerId = route.params.id as string;
    const response = await api.get<{ data: Offer }>(
      OFFER_ENDPOINTS.DETAIL(offerId),
    );
    offer.value = response.data.data;
  } catch (err) {
    const apiError = err as ApiError;
    error.value = `Erro ao carregar oferta: ${apiError.response?.data?.message || apiError.message || "Erro desconhecido"}`;
  } finally {
    loading.value = false;
  }
};

const loadPurchases = async () => {
  purchasesLoading.value = true;

  try {
    const offerId = route.params.id as string;
    const response = await api.get<{ data: PurchaseListItem[] }>(
      OFFER_ENDPOINTS.PURCHASE(offerId),
    );
    purchases.value = response.data.data;
  } catch (err) {
  } finally {
    purchasesLoading.value = false;
  }
};

const isOfferEditable = computed(() => {
  if (!offer.value) return false;
  return offer.value.status === "fresh";
});

const goBack = () => {
  router.push({ name: RouteNames.DASHBOARD });
};

const openEditDialog = () => {
  showEditDialog.value = true;
};

const handleOfferUpdated = () => {
  loadOffer();
  loadPurchases();
};

const openDeleteDialog = () => {
  showDeleteDialog.value = true;
};

const confirmDelete = async () => {
  if (!offer.value) return;

  showDeleteDialog.value = false;

  try {
    await api.delete(OFFER_ENDPOINTS.DELETE(offer.value.uuid));
    router.push({ name: RouteNames.DASHBOARD });
  } catch (err) {
    showDeleteErrorModal.value = true;
  }
};

const cancelDelete = () => {
  showDeleteDialog.value = false;
};

const closeDeleteErrorModal = () => {
  showDeleteErrorModal.value = false;
};

onMounted(() => {
  loadOffer();
  loadPurchases();
});
</script>

<template>
  <div class="min-h-screen bg-neutral-50">
    <div class="mx-auto max-w-7xl p-8">
      <button
        @click="goBack"
        class="mb-6 flex items-center gap-2 text-sm font-medium text-neutral-600 transition-colors hover:text-neutral-900"
      >
        <ArrowLeft :size="16" />
        Voltar ao Dashboard
      </button>

      <div v-if="loading" class="flex items-center justify-center py-20">
        <div class="flex flex-col items-center gap-3">
          <div
            class="h-12 w-12 animate-spin rounded-full border-4 border-emerald-500 border-t-transparent"
          ></div>
          <p class="text-sm text-neutral-500">Carregando oferta...</p>
        </div>
      </div>

      <div
        v-else-if="error"
        class="rounded-xl border border-red-200 bg-red-50 p-8 text-center"
      >
        <p class="text-sm text-red-600">{{ error }}</p>
        <button
          @click="loadOffer"
          class="mt-4 rounded-lg bg-red-600 px-6 py-2.5 text-sm font-medium text-white hover:bg-red-700"
        >
          Tentar Novamente
        </button>
      </div>

      <div v-else-if="offer" class="flex flex-col gap-6">
        <div
          class="overflow-hidden rounded-xl border border-neutral-200 bg-white shadow-sm"
        >
          <div
            :class="['h-3 w-full bg-gradient-to-r', typeConfig.gradient]"
          ></div>

          <div class="p-8">
            <div class="flex items-center gap-3">
              <h1 class="text-3xl font-bold text-neutral-900">
                {{ offer.description }}
              </h1>
              <span
                :class="[
                  'rounded-full px-4 py-1.5 text-sm font-medium',
                  typeConfig.badge,
                ]"
              >
                {{ typeConfig.label || offer.energy_type }}
              </span>
            </div>

            <div class="mt-6 flex items-baseline gap-2">
              <span class="text-4xl font-bold text-neutral-900">
                {{ formatPrice(offer.price_per_mwh) }}
              </span>
              <span class="text-lg text-neutral-500">/MWh</span>
            </div>

            <div class="mt-6">
              <div
                class="flex items-center justify-between text-sm text-neutral-600"
              >
                <span
                  >Disponível:
                  <span class="font-semibold">{{
                    formatQuantity(offer.remaining_quantity_mwh)
                  }}</span></span
                >
                <span
                  >Total:
                  <span class="font-semibold">{{
                    formatQuantity(offer.initial_quantity_mwh)
                  }}</span></span
                >
              </div>
              <div class="mt-2">
                <ProgressBar
                  :current-value="offer.remaining_quantity_mwh"
                  :total-value="offer.initial_quantity_mwh"
                  :show-percentage="true"
                  height="md"
                  gradient-from="emerald-500"
                  gradient-to="green-500"
                />
              </div>
            </div>

            <div
              class="mt-6 flex flex-wrap items-center gap-6 border-t border-neutral-100 pt-6"
            >
              <div class="flex items-center gap-2 text-sm">
                <span class="font-medium text-neutral-600">Submercado:</span>
                <SubmarketBadge
                  :offer-submarket="offer.submarket"
                  size="small"
                />
              </div>
              <div class="flex items-center gap-2 text-sm text-neutral-600">
                <Calendar :size="18" class="text-neutral-400" />
                <span
                  ><span class="font-medium">Início:</span>
                  {{ formatDate(offer.period_start) }}</span
                >
              </div>
              <div class="flex items-center gap-2 text-sm text-neutral-600">
                <Calendar :size="18" class="text-neutral-400" />
                <span
                  ><span class="font-medium">Término:</span>
                  {{ formatDate(offer.period_end) }}</span
                >
              </div>
            </div>

            <div class="mt-6 flex gap-3 border-t border-neutral-100 pt-6">
              <button
                v-tooltip.top="{
                  value: isOfferEditable
                    ? ''
                    : 'Você só pode editar ofertas que ainda não tiveram interações',
                  showDelay: 300,
                  hideDelay: 100,
                  class: 'tooltip-info',
                  disabled: isOfferEditable,
                }"
                @click="openEditDialog"
                :disabled="!isOfferEditable"
                :class="[
                  'flex items-center justify-center gap-2 rounded-lg border-2 px-6 py-2.5 text-sm font-medium transition-all duration-200',
                  isOfferEditable
                    ? 'border-blue-500 bg-transparent text-blue-600 hover:bg-blue-500 hover:text-white'
                    : 'cursor-not-allowed border-neutral-300 bg-neutral-100 text-neutral-400',
                ]"
              >
                <Edit :size="16" />
                Editar Oferta
              </button>
              <button
                v-tooltip.top="{
                  value: isOfferEditable
                    ? ''
                    : 'Você só pode excluir ofertas que ainda não tiveram interações',
                  showDelay: 300,
                  hideDelay: 100,
                  class: 'tooltip-info',
                  disabled: isOfferEditable,
                }"
                @click="openDeleteDialog"
                :disabled="!isOfferEditable"
                :class="[
                  'flex items-center justify-center gap-2 rounded-lg border-2 px-6 py-2.5 text-sm font-medium transition-all duration-200',
                  isOfferEditable
                    ? 'border-red-500 bg-transparent text-red-600 hover:bg-red-500 hover:text-white'
                    : 'cursor-not-allowed border-neutral-300 bg-neutral-100 text-neutral-400',
                ]"
              >
                <Trash2 :size="16" />
                Excluir Oferta
              </button>
            </div>
          </div>
        </div>

        <OfferStatsCards
          :total-sales="stats.totalSales"
          :total-revenue="stats.totalRevenue"
          :total-quantity-sold="stats.totalQuantitySold"
          :remaining-quantity="stats.remainingQuantity"
        />

        <OfferPurchasesTable
          :purchases="purchases"
          :loading="purchasesLoading"
        />
      </div>

      <CreateOfferDialog
        v-model:visible="showEditDialog"
        :offer="offerForEdit"
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
