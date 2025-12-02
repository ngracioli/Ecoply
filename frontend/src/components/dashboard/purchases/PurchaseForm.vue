<script setup lang="ts">
import { ref, computed } from "vue";
import type { OfferListItem } from "../../../types/responses/offers";
import type { PaymentMethod, CardDetails } from "../../../types/checkout";
import EnergyQuantityInput from "./EnergyQuantityInput.vue";
import PaymentMethodSelector from "./PaymentMethodSelector.vue";
import ContractTerms from "./ContractTerms.vue";
import OrderSummary from "./OrderSummary.vue";

interface Props {
  offer: OfferListItem;
  processing: boolean;
}

interface Emits {
  (e: "submit", data: PurchaseFormData): void;
}

export interface PurchaseFormData {
  quantityMwh: number;
  paymentMethod: PaymentMethod;
  installments: number;
  cardDetails: CardDetails;
}

const props = defineProps<Props>();
const emit = defineEmits<Emits>();

const quantityMwh = ref<number | null>(null);
const paymentMethod = ref<PaymentMethod>("pix");
const installments = ref<number>(1);
const acceptedTerms = ref(false);
const isQuantityValid = ref(false);
const cardDetails = ref<CardDetails>({
  number: "",
  holder_name: "",
  expiry_date: "",
  cvv: "",
});

const maxQuantity = computed(() => {
  return props.offer?.remaining_quantity_mwh || 0;
});

const totalPrice = computed(() => {
  if (!props.offer || !quantityMwh.value) return 0;
  return props.offer.price_per_mwh * quantityMwh.value;
});

const installmentPrice = computed(() => {
  return totalPrice.value / installments.value;
});

const handleQuantityValidation = (isValid: boolean) => {
  isQuantityValid.value = isValid;
};

const isFormValid = computed(() => {
  if (!quantityMwh.value || !isQuantityValid.value) return false;

  if (!acceptedTerms.value) return false;

  if (paymentMethod.value === "card") {
    return (
      cardDetails.value.number.replace(/\s/g, "").length === 16 &&
      cardDetails.value.holder_name.trim().length > 0 &&
      cardDetails.value.expiry_date.length === 5 &&
      cardDetails.value.cvv.length >= 3
    );
  }

  return true;
});

const handleSubmit = () => {
  if (!isFormValid.value || !quantityMwh.value) return;

  emit("submit", {
    quantityMwh: quantityMwh.value,
    paymentMethod: paymentMethod.value,
    installments: installments.value,
    cardDetails: cardDetails.value,
  });
};
</script>

<template>
  <div class="grid gap-4 sm:gap-6 lg:grid-cols-3 lg:gap-8">
    <div class="order-1 lg:order-1 lg:col-span-2">
      <div class="space-y-4 sm:space-y-6">
        <EnergyQuantityInput
          v-model="quantityMwh"
          :max-quantity="maxQuantity"
          @validation-change="handleQuantityValidation"
        />

        <PaymentMethodSelector
          v-model="paymentMethod"
          v-model:installments="installments"
          v-model:card-details="cardDetails"
          :installment-price="installmentPrice"
        />

        <ContractTerms
          v-model="acceptedTerms"
          :offer="offer"
          :quantity-mwh="quantityMwh"
          :price-per-mwh="offer.price_per_mwh"
          :total-price="totalPrice"
          :payment-method="paymentMethod"
          :installments="installments"
          :installment-price="installmentPrice"
        />
      </div>
    </div>

    <div class="order-2 lg:order-2 lg:col-span-1">
      <OrderSummary
        :offer="offer"
        :quantity-mwh="quantityMwh"
        :price-per-mwh="offer.price_per_mwh"
        :total-price="totalPrice"
        :payment-method="paymentMethod"
        :installments="installments"
        :installment-price="installmentPrice"
        :is-form-valid="isFormValid"
        :processing="processing"
        :accepted-terms="acceptedTerms"
        @submit="handleSubmit"
      />
    </div>
  </div>
</template>
