<script setup lang="ts">
import type { OfferListItem } from "../../../types/responses/offers";
import type { PaymentMethod } from "../../../types/checkout";
import { Lock, ShieldCheck } from "lucide-vue-next";
import { useIconSize } from "../../../composables/useIconSize";

interface Props {
  offer: OfferListItem;
  quantityMwh: number | null;
  pricePerMwh: number;
  totalPrice: number;
  paymentMethod: PaymentMethod;
  installments: number;
  installmentPrice: number;
  isFormValid: boolean;
  processing: boolean;
  acceptedTerms: boolean;
}

interface Emits {
  (e: "submit"): void;
}

defineProps<Props>();
const emit = defineEmits<Emits>();
const { iconSize } = useIconSize();

const formatPrice = (price: number) => {
  return `R$ ${price.toFixed(2).replace(".", ",")}`;
};

const formatQuantity = (mwh: number) => {
  return `${mwh.toLocaleString("pt-BR")} MWh`;
};

const getTypeLabel = (type: string) => {
  const labels = {
    solar: "Energia Solar",
    eolic: "Energia Eólica",
    hydroelectric: "Energia Hidrelétrica",
    geothermal: "Energia Geotérmica",
  };
  return labels[type as keyof typeof labels] || type;
};

const handleSubmit = () => {
  emit("submit");
};
</script>

<template>
  <div class="sticky top-4 space-y-4 sm:top-8 sm:space-y-5 md:space-y-6">
    <div class="rounded-lg bg-white p-4 shadow-sm sm:rounded-xl sm:p-5 md:p-6">
      <h2
        class="mb-4 text-base font-semibold text-neutral-900 sm:mb-5 sm:text-lg md:mb-6"
      >
        Resumo do Pedido
      </h2>

      <div class="space-y-3 sm:space-y-4">
        <div class="rounded-lg bg-neutral-50 p-3 sm:p-4">
          <p
            class="text-xs font-medium break-words text-neutral-600 sm:text-sm"
          >
            {{ getTypeLabel(offer.energy_type) }}
          </p>
          <p class="mt-1 text-xs break-words text-neutral-500">
            {{ offer.description }}
          </p>
        </div>

        <div
          class="space-y-2 border-t border-neutral-200 pt-3 sm:space-y-3 sm:pt-4"
        >
          <div
            class="flex items-center justify-between gap-3 text-xs sm:text-sm"
          >
            <span class="shrink-0 text-neutral-600">Quantidade</span>
            <span
              class="min-w-0 text-right font-medium break-words text-neutral-900"
            >
              {{ quantityMwh ? formatQuantity(quantityMwh) : "—" }}
            </span>
          </div>

          <div
            class="flex items-center justify-between gap-3 text-xs sm:text-sm"
          >
            <span class="shrink-0 text-neutral-600">Preço por MWh</span>
            <span
              class="min-w-0 text-right font-medium break-words text-neutral-900"
            >
              {{ formatPrice(pricePerMwh) }}
            </span>
          </div>

          <div
            v-if="paymentMethod === 'card' && installments > 1"
            class="flex items-center justify-between gap-3 text-xs sm:text-sm"
          >
            <span class="shrink-0 text-neutral-600">Parcelas</span>
            <span
              class="min-w-0 text-right font-medium break-words text-neutral-900"
            >
              {{ installments }}x sem juros
            </span>
          </div>
        </div>

        <div
          class="flex items-center justify-between gap-3 border-t-2 border-neutral-200 pt-3 sm:pt-4"
        >
          <span
            class="shrink-0 text-base font-semibold text-neutral-900 sm:text-lg"
            >Total</span
          >
          <span
            class="min-w-0 text-right text-xl font-bold break-words text-emerald-600 sm:text-2xl"
          >
            {{ formatPrice(totalPrice) }}
          </span>
        </div>

        <div
          v-if="paymentMethod === 'card' && installments > 1"
          class="text-center text-xs text-neutral-500 sm:text-sm"
        >
          {{ installments }}x de {{ formatPrice(installmentPrice) }}
        </div>
      </div>
    </div>

    <div class="rounded-lg bg-emerald-50 p-4 sm:rounded-xl sm:p-5 md:p-6">
      <div class="flex items-start gap-2.5 sm:gap-3">
        <div class="shrink-0 rounded-lg bg-emerald-100 p-1.5 sm:p-2">
          <ShieldCheck class="text-emerald-600" :size="iconSize(18, 20)" />
        </div>
        <div class="min-w-0">
          <h3 class="text-sm font-semibold text-emerald-900 sm:text-base">
            Compra Segura
          </h3>
          <p class="mt-1 text-xs break-words text-emerald-700 sm:text-sm">
            Seus dados estão protegidos com criptografia de ponta a ponta
          </p>
        </div>
      </div>
    </div>

    <button
      @click="handleSubmit"
      :disabled="!isFormValid || processing"
      :class="[
        'min-h-[44px] w-full touch-manipulation rounded-lg py-3 text-base font-semibold text-white shadow-lg transition-all active:scale-[0.98] sm:rounded-xl sm:py-3.5 md:py-4',
        isFormValid && !processing
          ? 'bg-gradient-to-r from-emerald-500 to-green-500 hover:shadow-xl hover:brightness-110'
          : 'cursor-not-allowed bg-neutral-300',
      ]"
    >
      <span v-if="processing" class="flex items-center justify-center gap-2">
        <div
          class="h-4 w-4 animate-spin rounded-full border-2 border-white border-t-transparent sm:h-5 sm:w-5"
        ></div>
        Processando...
      </span>
      <span v-else class="flex items-center justify-center gap-2">
        <Lock :size="iconSize(18, 20)" />
        Finalizar Compra
      </span>
    </button>

    <div class="space-y-1.5 sm:space-y-2">
      <p
        v-if="!acceptedTerms"
        class="px-2 text-center text-xs font-medium break-words text-amber-600"
      >
        ⚠️ Você precisa aceitar os termos e condições para finalizar
      </p>
      <p v-else class="px-2 text-center text-xs break-words text-neutral-500">
        Ao finalizar, você concorda com os termos do contrato
      </p>
    </div>
  </div>
</template>
