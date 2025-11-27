<script setup lang="ts">
import { computed, watch } from "vue";
import type { PaymentMethod, CardDetails } from "../../../types/checkout";
import { CreditCard, FileText, Wallet, Check } from "lucide-vue-next";

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
  <div class="rounded-xl bg-white p-6 shadow-sm">
    <h2 class="mb-6 text-lg font-semibold text-neutral-900">
      Forma de Pagamento
    </h2>

    <div class="space-y-3">
      <label
        :class="[
          'flex cursor-pointer items-center gap-4 rounded-lg border-2 p-4 transition-all',
          paymentMethod === 'pix'
            ? 'border-emerald-500 bg-emerald-50'
            : 'border-neutral-200 hover:border-neutral-300',
        ]"
      >
        <input
          v-model="paymentMethod"
          type="radio"
          value="pix"
          class="h-5 w-5 text-emerald-500 focus:ring-emerald-500"
        />
        <div class="flex flex-1 items-center gap-3">
          <div
            :class="[
              'rounded-lg p-2',
              paymentMethod === 'pix' ? 'bg-emerald-100' : 'bg-neutral-100',
            ]"
          >
            <Wallet
              :size="20"
              :class="
                paymentMethod === 'pix'
                  ? 'text-emerald-600'
                  : 'text-neutral-600'
              "
            />
          </div>
          <div>
            <p class="font-medium text-neutral-900">PIX</p>
            <p class="text-sm text-neutral-500">Aprovação instantânea</p>
          </div>
        </div>
        <Check
          v-if="paymentMethod === 'pix'"
          :size="20"
          class="text-emerald-500"
        />
      </label>

      <label
        :class="[
          'flex cursor-pointer items-center gap-4 rounded-lg border-2 p-4 transition-all',
          paymentMethod === 'card'
            ? 'border-emerald-500 bg-emerald-50'
            : 'border-neutral-200 hover:border-neutral-300',
        ]"
      >
        <input
          v-model="paymentMethod"
          type="radio"
          value="card"
          class="h-5 w-5 text-emerald-500 focus:ring-emerald-500"
        />
        <div class="flex flex-1 items-center gap-3">
          <div
            :class="[
              'rounded-lg p-2',
              paymentMethod === 'card'
                ? 'bg-emerald-100'
                : 'bg-neutral-100',
            ]"
          >
            <CreditCard
              :size="20"
              :class="
                paymentMethod === 'card'
                  ? 'text-emerald-600'
                  : 'text-neutral-600'
              "
            />
          </div>
          <div>
            <p class="font-medium text-neutral-900">Cartão de Crédito</p>
            <p class="text-sm text-neutral-500">Em até 12x sem juros</p>
          </div>
        </div>
        <Check
          v-if="paymentMethod === 'card'"
          :size="20"
          class="text-emerald-500"
        />
      </label>

      <label
        :class="[
          'flex cursor-pointer items-center gap-4 rounded-lg border-2 p-4 transition-all',
          paymentMethod === 'billet'
            ? 'border-emerald-500 bg-emerald-50'
            : 'border-neutral-200 hover:border-neutral-300',
        ]"
      >
        <input
          v-model="paymentMethod"
          type="radio"
          value="billet"
          class="h-5 w-5 text-emerald-500 focus:ring-emerald-500"
        />
        <div class="flex flex-1 items-center gap-3">
          <div
            :class="[
              'rounded-lg p-2',
              paymentMethod === 'billet' ? 'bg-emerald-100' : 'bg-neutral-100',
            ]"
          >
            <FileText
              :size="20"
              :class="
                paymentMethod === 'billet'
                  ? 'text-emerald-600'
                  : 'text-neutral-600'
              "
            />
          </div>
          <div>
            <p class="font-medium text-neutral-900">Boleto Bancário</p>
            <p class="text-sm text-neutral-500">Vencimento em 3 dias úteis</p>
          </div>
        </div>
        <Check
          v-if="paymentMethod === 'billet'"
          :size="20"
          class="text-emerald-500"
        />
      </label>
    </div>

    <div
      v-if="paymentMethod === 'card'"
      class="mt-6 space-y-4 border-t border-neutral-200 pt-6"
    >
      <div>
        <label class="mb-2 block text-sm font-medium text-neutral-700">
          Número do Cartão
        </label>
        <input
          v-model="card.number"
          @input="formatCardNumber"
          type="text"
          placeholder="0000 0000 0000 0000"
          maxlength="19"
          class="w-full rounded-lg border border-neutral-300 px-4 py-3 focus:border-emerald-500 focus:ring-2 focus:ring-emerald-500/20 focus:outline-none"
        />
      </div>

      <div>
        <label class="mb-2 block text-sm font-medium text-neutral-700">
          Nome do Titular
        </label>
        <input
          v-model="card.holder_name"
          type="text"
          placeholder="Nome como está no cartão"
          class="w-full rounded-lg border border-neutral-300 px-4 py-3 uppercase focus:border-emerald-500 focus:ring-2 focus:ring-emerald-500/20 focus:outline-none"
        />
      </div>

      <div class="grid grid-cols-2 gap-4">
        <div>
          <label class="mb-2 block text-sm font-medium text-neutral-700">
            Validade
          </label>
          <input
            v-model="card.expiry_date"
            @input="formatExpiryDate"
            type="text"
            placeholder="MM/AA"
            maxlength="5"
            class="w-full rounded-lg border border-neutral-300 px-4 py-3 focus:border-emerald-500 focus:ring-2 focus:ring-emerald-500/20 focus:outline-none"
          />
        </div>
        <div>
          <label class="mb-2 block text-sm font-medium text-neutral-700">
            CVV
          </label>
          <input
            v-model="card.cvv"
            @input="formatCVV"
            type="text"
            placeholder="123"
            maxlength="4"
            class="w-full rounded-lg border border-neutral-300 px-4 py-3 focus:border-emerald-500 focus:ring-2 focus:ring-emerald-500/20 focus:outline-none"
          />
        </div>
      </div>

      <div>
        <label class="mb-2 block text-sm font-medium text-neutral-700">
          Número de Parcelas
        </label>
        <select
          v-model.number="selectedInstallments"
          class="w-full rounded-lg border border-neutral-300 px-4 py-3 focus:border-emerald-500 focus:ring-2 focus:ring-emerald-500/20 focus:outline-none"
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
