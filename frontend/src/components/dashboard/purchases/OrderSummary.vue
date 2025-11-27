<script setup lang="ts">
import type { OfferListItem } from "../../../types/responses/offers";
import type { PaymentMethod } from "../../../types/checkout";
import { Lock, ShieldCheck } from "lucide-vue-next";

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
  <div class="sticky top-8 space-y-6">
    <div class="rounded-xl bg-white p-6 shadow-sm">
      <h2 class="mb-6 text-lg font-semibold text-neutral-900">
        Resumo do Pedido
      </h2>

      <div class="space-y-4">
        <div class="rounded-lg bg-neutral-50 p-4">
          <p class="text-sm font-medium text-neutral-600">
            {{ getTypeLabel(offer.energy_type) }}
          </p>
          <p class="mt-1 text-xs text-neutral-500">
            {{ offer.description }}
          </p>
        </div>

        <div class="space-y-3 border-t border-neutral-200 pt-4">
          <div class="flex justify-between text-sm">
            <span class="text-neutral-600">Quantidade</span>
            <span class="font-medium text-neutral-900">
              {{ quantityMwh ? formatQuantity(quantityMwh) : "—" }}
            </span>
          </div>

          <div class="flex justify-between text-sm">
            <span class="text-neutral-600">Preço por MWh</span>
            <span class="font-medium text-neutral-900">
              {{ formatPrice(pricePerMwh) }}
            </span>
          </div>

          <div
            v-if="paymentMethod === 'card' && installments > 1"
            class="flex justify-between text-sm"
          >
            <span class="text-neutral-600">Parcelas</span>
            <span class="font-medium text-neutral-900">
              {{ installments }}x sem juros
            </span>
          </div>
        </div>

        <div class="flex justify-between border-t-2 border-neutral-200 pt-4">
          <span class="text-lg font-semibold text-neutral-900">Total</span>
          <span class="text-2xl font-bold text-emerald-600">
            {{ formatPrice(totalPrice) }}
          </span>
        </div>

        <div
          v-if="paymentMethod === 'card' && installments > 1"
          class="text-center text-sm text-neutral-500"
        >
          {{ installments }}x de {{ formatPrice(installmentPrice) }}
        </div>
      </div>
    </div>

    <div class="rounded-xl bg-emerald-50 p-6">
      <div class="flex items-start gap-3">
        <div class="rounded-lg bg-emerald-100 p-2">
          <ShieldCheck class="text-emerald-600" :size="20" />
        </div>
        <div>
          <h3 class="font-semibold text-emerald-900">Compra Segura</h3>
          <p class="mt-1 text-sm text-emerald-700">
            Seus dados estão protegidos com criptografia de ponta a ponta
          </p>
        </div>
      </div>
    </div>

    <button
      @click="handleSubmit"
      :disabled="!isFormValid || processing"
      :class="[
        'w-full rounded-xl py-4 font-semibold text-white shadow-lg transition-all',
        isFormValid && !processing
          ? 'bg-gradient-to-r from-emerald-500 to-green-500 hover:shadow-xl hover:brightness-110'
          : 'cursor-not-allowed bg-neutral-300',
      ]"
    >
      <span v-if="processing" class="flex items-center justify-center gap-2">
        <div
          class="h-5 w-5 animate-spin rounded-full border-2 border-white border-t-transparent"
        ></div>
        Processando...
      </span>
      <span v-else class="flex items-center justify-center gap-2">
        <Lock :size="20" />
        Finalizar Compra
      </span>
    </button>

    <div class="space-y-2">
      <p
        v-if="!acceptedTerms"
        class="text-center text-xs font-medium text-amber-600"
      >
        ⚠️ Você precisa aceitar os termos e condições para finalizar
      </p>
      <p v-else class="text-center text-xs text-neutral-500">
        Ao finalizar, você concorda com os termos do contrato
      </p>
    </div>
  </div>
</template>
