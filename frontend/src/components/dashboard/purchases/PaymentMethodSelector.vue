<script setup lang="ts">
import { computed, watch } from "vue";
import type { PaymentMethod, CardDetails } from "../../../types/checkout";
import { CreditCard, FileText, Wallet, Check } from "lucide-vue-next";
import { useIconSize } from "../../../composables/useIconSize";

interface Props {
  modelValue: PaymentMethod;
  installments: number;
  cardDetails: CardDetails;
  installmentPrice: number;
}

interface Emits {
  (e: "update:modelValue", value: PaymentMethod): void;
  (e: "update:installments", value: number): void;
  (e: "update:cardDetails", value: CardDetails): void;
}

const props = defineProps<Props>();
const emit = defineEmits<Emits>();
const { iconSize } = useIconSize();

const paymentMethod = computed({
  get: () => props.modelValue,
  set: (value) => emit("update:modelValue", value),
});

const selectedInstallments = computed({
  get: () => props.installments,
  set: (value) => emit("update:installments", value),
});

const card = computed({
  get: () => props.cardDetails,
  set: (value) => emit("update:cardDetails", value),
});

const availableInstallments = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12];

watch(paymentMethod, (newMethod) => {
  if (newMethod !== "card") {
    selectedInstallments.value = 1;
  }
});

const formatCardNumber = () => {
  let value = card.value.number.replace(/\s/g, "");
  value = value.replace(/\D/g, "");
  value = value.substring(0, 16);
  value = value.replace(/(\d{4})/g, "$1 ").trim();
  card.value = { ...card.value, number: value };
};

const formatExpiryDate = () => {
  let value = card.value.expiry_date.replace(/\D/g, "");
  if (value.length >= 2) {
    value = value.substring(0, 2) + "/" + value.substring(2, 4);
  }
  card.value = { ...card.value, expiry_date: value };
};

const formatCVV = () => {
  const cvv = card.value.cvv.replace(/\D/g, "").substring(0, 4);
  card.value = { ...card.value, cvv };
};

const formatPrice = (price: number) => {
  return `R$ ${price.toFixed(2).replace(".", ",")}`;
};
</script>

<template>
  <div class="rounded-lg bg-white p-4 shadow-sm sm:rounded-xl sm:p-5 md:p-6">
    <h2
      class="mb-4 text-base font-semibold text-neutral-900 sm:mb-5 sm:text-lg md:mb-6"
    >
      Forma de Pagamento
    </h2>

    <div class="space-y-2.5 sm:space-y-3">
      <label
        :class="[
          'flex cursor-pointer touch-manipulation items-center gap-3 rounded-lg border-2 p-3 transition-all active:scale-[0.99] sm:gap-4 sm:p-4',
          paymentMethod === 'pix'
            ? 'border-emerald-500 bg-emerald-50'
            : 'border-neutral-200 hover:border-neutral-300',
        ]"
      >
        <input
          v-model="paymentMethod"
          type="radio"
          value="pix"
          class="h-5 w-5 shrink-0 touch-manipulation text-emerald-500 focus:ring-emerald-500"
        />
        <div class="flex min-w-0 flex-1 items-center gap-2.5 sm:gap-3">
          <div
            :class="[
              'shrink-0 rounded-lg p-1.5 sm:p-2',
              paymentMethod === 'pix' ? 'bg-emerald-100' : 'bg-neutral-100',
            ]"
          >
            <Wallet
              :size="iconSize(18, 20)"
              :class="
                paymentMethod === 'pix'
                  ? 'text-emerald-600'
                  : 'text-neutral-600'
              "
            />
          </div>
          <div class="min-w-0">
            <p class="text-sm font-medium text-neutral-900 sm:text-base">PIX</p>
            <p class="text-xs text-neutral-500 sm:text-sm">
              Aprovação instantânea
            </p>
          </div>
        </div>
        <Check
          v-if="paymentMethod === 'pix'"
          :size="iconSize(18, 20)"
          class="shrink-0 text-emerald-500"
        />
      </label>

      <label
        :class="[
          'flex cursor-pointer touch-manipulation items-center gap-3 rounded-lg border-2 p-3 transition-all active:scale-[0.99] sm:gap-4 sm:p-4',
          paymentMethod === 'card'
            ? 'border-emerald-500 bg-emerald-50'
            : 'border-neutral-200 hover:border-neutral-300',
        ]"
      >
        <input
          v-model="paymentMethod"
          type="radio"
          value="card"
          class="h-5 w-5 shrink-0 touch-manipulation text-emerald-500 focus:ring-emerald-500"
        />
        <div class="flex min-w-0 flex-1 items-center gap-2.5 sm:gap-3">
          <div
            :class="[
              'shrink-0 rounded-lg p-1.5 sm:p-2',
              paymentMethod === 'card' ? 'bg-emerald-100' : 'bg-neutral-100',
            ]"
          >
            <CreditCard
              :size="iconSize(18, 20)"
              :class="
                paymentMethod === 'card'
                  ? 'text-emerald-600'
                  : 'text-neutral-600'
              "
            />
          </div>
          <div class="min-w-0">
            <p class="text-sm font-medium text-neutral-900 sm:text-base">
              Cartão de Crédito
            </p>
            <p class="text-xs text-neutral-500 sm:text-sm">
              Em até 12x sem juros
            </p>
          </div>
        </div>
        <Check
          v-if="paymentMethod === 'card'"
          :size="iconSize(18, 20)"
          class="shrink-0 text-emerald-500"
        />
      </label>

      <label
        :class="[
          'flex cursor-pointer touch-manipulation items-center gap-3 rounded-lg border-2 p-3 transition-all active:scale-[0.99] sm:gap-4 sm:p-4',
          paymentMethod === 'billet'
            ? 'border-emerald-500 bg-emerald-50'
            : 'border-neutral-200 hover:border-neutral-300',
        ]"
      >
        <input
          v-model="paymentMethod"
          type="radio"
          value="billet"
          class="h-5 w-5 shrink-0 touch-manipulation text-emerald-500 focus:ring-emerald-500"
        />
        <div class="flex min-w-0 flex-1 items-center gap-2.5 sm:gap-3">
          <div
            :class="[
              'shrink-0 rounded-lg p-1.5 sm:p-2',
              paymentMethod === 'billet' ? 'bg-emerald-100' : 'bg-neutral-100',
            ]"
          >
            <FileText
              :size="iconSize(18, 20)"
              :class="
                paymentMethod === 'billet'
                  ? 'text-emerald-600'
                  : 'text-neutral-600'
              "
            />
          </div>
          <div class="min-w-0">
            <p class="text-sm font-medium text-neutral-900 sm:text-base">
              Boleto Bancário
            </p>
            <p class="text-xs text-neutral-500 sm:text-sm">
              Vencimento em 3 dias úteis
            </p>
          </div>
        </div>
        <Check
          v-if="paymentMethod === 'billet'"
          :size="iconSize(18, 20)"
          class="shrink-0 text-emerald-500"
        />
      </label>
    </div>

    <div
      v-if="paymentMethod === 'card'"
      class="mt-4 space-y-3 border-t border-neutral-200 pt-4 sm:mt-5 sm:space-y-4 sm:pt-5 md:mt-6 md:pt-6"
    >
      <div>
        <label
          class="mb-1.5 block text-xs font-medium text-neutral-700 sm:mb-2 sm:text-sm"
        >
          Número do Cartão
        </label>
        <input
          v-model="card.number"
          @input="formatCardNumber"
          type="text"
          inputmode="numeric"
          placeholder="0000 0000 0000 0000"
          maxlength="19"
          class="w-full min-w-0 touch-manipulation rounded-lg border border-neutral-300 px-3 py-2.5 text-base focus:border-emerald-500 focus:ring-2 focus:ring-emerald-500/20 focus:outline-none sm:px-4 sm:py-3"
        />
      </div>

      <div>
        <label
          class="mb-1.5 block text-xs font-medium text-neutral-700 sm:mb-2 sm:text-sm"
        >
          Nome do Titular
        </label>
        <input
          v-model="card.holder_name"
          type="text"
          placeholder="Nome como está no cartão"
          class="w-full min-w-0 touch-manipulation rounded-lg border border-neutral-300 px-3 py-2.5 text-base uppercase focus:border-emerald-500 focus:ring-2 focus:ring-emerald-500/20 focus:outline-none sm:px-4 sm:py-3"
        />
      </div>

      <div class="grid grid-cols-2 gap-3 sm:gap-4">
        <div>
          <label
            class="mb-1.5 block text-xs font-medium text-neutral-700 sm:mb-2 sm:text-sm"
          >
            Validade
          </label>
          <input
            v-model="card.expiry_date"
            @input="formatExpiryDate"
            type="text"
            inputmode="numeric"
            placeholder="MM/AA"
            maxlength="5"
            class="w-full min-w-0 touch-manipulation rounded-lg border border-neutral-300 px-3 py-2.5 text-base focus:border-emerald-500 focus:ring-2 focus:ring-emerald-500/20 focus:outline-none sm:px-4 sm:py-3"
          />
        </div>
        <div>
          <label
            class="mb-1.5 block text-xs font-medium text-neutral-700 sm:mb-2 sm:text-sm"
          >
            CVV
          </label>
          <input
            v-model="card.cvv"
            @input="formatCVV"
            type="text"
            inputmode="numeric"
            placeholder="123"
            maxlength="4"
            class="w-full min-w-0 touch-manipulation rounded-lg border border-neutral-300 px-3 py-2.5 text-base focus:border-emerald-500 focus:ring-2 focus:ring-emerald-500/20 focus:outline-none sm:px-4 sm:py-3"
          />
        </div>
      </div>

      <div>
        <label
          class="mb-1.5 block text-xs font-medium text-neutral-700 sm:mb-2 sm:text-sm"
        >
          Número de Parcelas
        </label>
        <select
          v-model.number="selectedInstallments"
          class="w-full min-w-0 touch-manipulation rounded-lg border border-neutral-300 px-3 py-2.5 text-base focus:border-emerald-500 focus:ring-2 focus:ring-emerald-500/20 focus:outline-none sm:px-4 sm:py-3"
        >
          <option v-for="n in availableInstallments" :key="n" :value="n">
            {{ n }}x de {{ formatPrice(installmentPrice / n) }}
            {{ n === 1 ? "à vista" : "sem juros" }}
          </option>
        </select>
      </div>
    </div>
  </div>
</template>
