<script setup lang="ts">
import { ref, onMounted, computed, watch } from "vue";
import { useRoute, useRouter } from "vue-router";
import api from "../axios";
import { OFFER_ENDPOINTS } from "../api/endpoints";
import type {
  OfferListItem,
  OfferDetailResponse,
} from "../types/responses/offers";
import type { PaymentMethod, CardDetails } from "../types/checkout";
import {
  ArrowLeft,
  CreditCard,
  FileText,
  Wallet,
  Check,
  Lock,
  Zap,
  ShieldCheck,
} from "lucide-vue-next";

const route = useRoute();
const router = useRouter();
const offer = ref<OfferListItem | null>(null);
const loading = ref(true);
const processing = ref(false);
const error = ref<string | null>(null);

// Form data
const quantityMwh = ref<number>(1);
const paymentMethod = ref<PaymentMethod>("pix");
const installments = ref<number>(1);
const cardDetails = ref<CardDetails>({
  number: "",
  holder_name: "",
  expiry_date: "",
  cvv: "",
});

const offerId = computed(() => route.params.id as string);

const fetchOffer = async () => {
  try {
    loading.value = true;
    error.value = null;
    const response = await api.get<OfferDetailResponse>(
      OFFER_ENDPOINTS.DETAIL(offerId.value),
    );
    offer.value = response.data.data;
  } catch (err: any) {
    error.value =
      err.response?.data?.message || "Erro ao carregar os detalhes da oferta";
    console.error("Erro ao buscar oferta:", err);
  } finally {
    loading.value = false;
  }
};

onMounted(() => {
  fetchOffer();
});

const goBack = () => {
  router.back();
};

const formatPrice = (price: number) => {
  return `R$ ${price.toFixed(2).replace(".", ",")}`;
};

const formatQuantity = (mwh: number) => {
  const kwh = (mwh * 1000).toFixed(0);
  return `${parseInt(kwh).toLocaleString("pt-BR")} kWh`;
};

const maxQuantity = computed(() => {
  return offer.value?.remaining_quantity_mwh || 0;
});

const totalPrice = computed(() => {
  if (!offer.value) return 0;
  return offer.value.price_per_mwh * quantityMwh.value;
});

const installmentPrice = computed(() => {
  return totalPrice.value / installments.value;
});

const availableInstallments = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12];

// Watch payment method changes
watch(paymentMethod, (newMethod) => {
  if (newMethod !== "credit_card") {
    installments.value = 1;
  }
});

// Quantity validation
const handleQuantityChange = () => {
  if (quantityMwh.value < 0.1) {
    quantityMwh.value = 0.1;
  }
  if (quantityMwh.value > maxQuantity.value) {
    quantityMwh.value = maxQuantity.value;
  }
};

// Card number formatting
const formatCardNumber = () => {
  let value = cardDetails.value.number.replace(/\s/g, "");
  value = value.replace(/\D/g, "");
  value = value.substring(0, 16);
  value = value.replace(/(\d{4})/g, "$1 ").trim();
  cardDetails.value.number = value;
};

// Expiry date formatting
const formatExpiryDate = () => {
  let value = cardDetails.value.expiry_date.replace(/\D/g, "");
  if (value.length >= 2) {
    value = value.substring(0, 2) + "/" + value.substring(2, 4);
  }
  cardDetails.value.expiry_date = value;
};

// CVV formatting
const formatCVV = () => {
  cardDetails.value.cvv = cardDetails.value.cvv
    .replace(/\D/g, "")
    .substring(0, 4);
};

const isFormValid = computed(() => {
  if (!quantityMwh.value || quantityMwh.value <= 0) return false;

  if (paymentMethod.value === "credit_card") {
    return (
      cardDetails.value.number.replace(/\s/g, "").length === 16 &&
      cardDetails.value.holder_name.trim().length > 0 &&
      cardDetails.value.expiry_date.length === 5 &&
      cardDetails.value.cvv.length >= 3
    );
  }

  return true;
});

const processPayment = async () => {
  if (!isFormValid.value) return;

  processing.value = true;

  // Simulate payment processing
  setTimeout(() => {
    processing.value = false;
    alert("Compra realizada com sucesso! (Simulação)");
    router.push({ name: "Dashboard" });
  }, 2000);
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
</script>

<template>
  <div class="min-h-screen bg-neutral-50">
    <!-- Header -->
    <header class="border-b border-neutral-200 bg-white">
      <div class="mx-auto max-w-7xl px-4 py-6 sm:px-6 lg:px-8">
        <div class="flex items-center gap-4">
          <button
            @click="goBack"
            class="rounded-lg p-2 text-neutral-600 transition-colors hover:bg-neutral-100"
          >
            <ArrowLeft :size="24" />
          </button>
          <h1 class="text-2xl font-semibold text-neutral-900">
            Finalizar Compra
          </h1>
        </div>
      </div>
    </header>

    <!-- Content -->
    <main class="mx-auto max-w-7xl px-4 py-8 sm:px-6 lg:px-8">
      <!-- Loading State -->
      <div v-if="loading" class="flex items-center justify-center py-20">
        <div class="text-center">
          <div
            class="mx-auto h-12 w-12 animate-spin rounded-full border-4 border-neutral-200 border-t-emerald-500"
          ></div>
          <p class="mt-4 text-neutral-600">Carregando...</p>
        </div>
      </div>

      <!-- Error State -->
      <div
        v-else-if="error"
        class="rounded-xl border border-red-200 bg-red-50 p-6 text-center"
      >
        <p class="text-red-700">{{ error }}</p>
        <button
          @click="fetchOffer"
          class="mt-4 rounded-lg bg-red-600 px-4 py-2 text-white transition-colors hover:bg-red-700"
        >
          Tentar novamente
        </button>
      </div>

      <!-- Checkout Content -->
      <div v-else-if="offer" class="grid gap-8 lg:grid-cols-3">
        <!-- Left Column - Form -->
        <div class="lg:col-span-2">
          <div class="space-y-6">
            <!-- Quantity Selection -->
            <div class="rounded-xl bg-white p-6 shadow-sm">
              <h2 class="mb-6 text-lg font-semibold text-neutral-900">
                Quantidade de Energia
              </h2>

              <div class="space-y-4">
                <div>
                  <label
                    class="mb-2 block text-sm font-medium text-neutral-700"
                  >
                    Quantidade (MWh)
                  </label>
                  <input
                    v-model.number="quantityMwh"
                    @input="handleQuantityChange"
                    type="number"
                    step="0.1"
                    :max="maxQuantity"
                    min="0.1"
                    class="w-full rounded-lg border border-neutral-300 px-4 py-3 text-lg focus:border-emerald-500 focus:ring-2 focus:ring-emerald-500/20 focus:outline-none"
                  />
                  <p class="mt-2 text-sm text-neutral-500">
                    Disponível: {{ formatQuantity(maxQuantity) }} ({{
                      maxQuantity
                    }}
                    MWh)
                  </p>
                </div>

                <div
                  class="flex items-center gap-2 rounded-lg bg-emerald-50 p-3 text-sm text-emerald-700"
                >
                  <Zap :size="16" />
                  <span>{{ formatQuantity(quantityMwh) }}</span>
                </div>
              </div>
            </div>

            <!-- Payment Method -->
            <div class="rounded-xl bg-white p-6 shadow-sm">
              <h2 class="mb-6 text-lg font-semibold text-neutral-900">
                Forma de Pagamento
              </h2>

              <div class="space-y-3">
                <!-- PIX -->
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
                        paymentMethod === 'pix'
                          ? 'bg-emerald-100'
                          : 'bg-neutral-100',
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
                      <p class="text-sm text-neutral-500">
                        Aprovação instantânea
                      </p>
                    </div>
                  </div>
                  <Check
                    v-if="paymentMethod === 'pix'"
                    :size="20"
                    class="text-emerald-500"
                  />
                </label>

                <!-- Credit Card -->
                <label
                  :class="[
                    'flex cursor-pointer items-center gap-4 rounded-lg border-2 p-4 transition-all',
                    paymentMethod === 'credit_card'
                      ? 'border-emerald-500 bg-emerald-50'
                      : 'border-neutral-200 hover:border-neutral-300',
                  ]"
                >
                  <input
                    v-model="paymentMethod"
                    type="radio"
                    value="credit_card"
                    class="h-5 w-5 text-emerald-500 focus:ring-emerald-500"
                  />
                  <div class="flex flex-1 items-center gap-3">
                    <div
                      :class="[
                        'rounded-lg p-2',
                        paymentMethod === 'credit_card'
                          ? 'bg-emerald-100'
                          : 'bg-neutral-100',
                      ]"
                    >
                      <CreditCard
                        :size="20"
                        :class="
                          paymentMethod === 'credit_card'
                            ? 'text-emerald-600'
                            : 'text-neutral-600'
                        "
                      />
                    </div>
                    <div>
                      <p class="font-medium text-neutral-900">
                        Cartão de Crédito
                      </p>
                      <p class="text-sm text-neutral-500">
                        Em até 12x sem juros
                      </p>
                    </div>
                  </div>
                  <Check
                    v-if="paymentMethod === 'credit_card'"
                    :size="20"
                    class="text-emerald-500"
                  />
                </label>

                <!-- Boleto -->
                <label
                  :class="[
                    'flex cursor-pointer items-center gap-4 rounded-lg border-2 p-4 transition-all',
                    paymentMethod === 'boleto'
                      ? 'border-emerald-500 bg-emerald-50'
                      : 'border-neutral-200 hover:border-neutral-300',
                  ]"
                >
                  <input
                    v-model="paymentMethod"
                    type="radio"
                    value="boleto"
                    class="h-5 w-5 text-emerald-500 focus:ring-emerald-500"
                  />
                  <div class="flex flex-1 items-center gap-3">
                    <div
                      :class="[
                        'rounded-lg p-2',
                        paymentMethod === 'boleto'
                          ? 'bg-emerald-100'
                          : 'bg-neutral-100',
                      ]"
                    >
                      <FileText
                        :size="20"
                        :class="
                          paymentMethod === 'boleto'
                            ? 'text-emerald-600'
                            : 'text-neutral-600'
                        "
                      />
                    </div>
                    <div>
                      <p class="font-medium text-neutral-900">
                        Boleto Bancário
                      </p>
                      <p class="text-sm text-neutral-500">
                        Vencimento em 3 dias úteis
                      </p>
                    </div>
                  </div>
                  <Check
                    v-if="paymentMethod === 'boleto'"
                    :size="20"
                    class="text-emerald-500"
                  />
                </label>
              </div>

              <!-- Credit Card Details -->
              <div
                v-if="paymentMethod === 'credit_card'"
                class="mt-6 space-y-4 border-t border-neutral-200 pt-6"
              >
                <div>
                  <label
                    class="mb-2 block text-sm font-medium text-neutral-700"
                  >
                    Número do Cartão
                  </label>
                  <input
                    v-model="cardDetails.number"
                    @input="formatCardNumber"
                    type="text"
                    placeholder="0000 0000 0000 0000"
                    maxlength="19"
                    class="w-full rounded-lg border border-neutral-300 px-4 py-3 focus:border-emerald-500 focus:ring-2 focus:ring-emerald-500/20 focus:outline-none"
                  />
                </div>

                <div>
                  <label
                    class="mb-2 block text-sm font-medium text-neutral-700"
                  >
                    Nome do Titular
                  </label>
                  <input
                    v-model="cardDetails.holder_name"
                    type="text"
                    placeholder="Nome como está no cartão"
                    class="w-full rounded-lg border border-neutral-300 px-4 py-3 uppercase focus:border-emerald-500 focus:ring-2 focus:ring-emerald-500/20 focus:outline-none"
                  />
                </div>

                <div class="grid grid-cols-2 gap-4">
                  <div>
                    <label
                      class="mb-2 block text-sm font-medium text-neutral-700"
                    >
                      Validade
                    </label>
                    <input
                      v-model="cardDetails.expiry_date"
                      @input="formatExpiryDate"
                      type="text"
                      placeholder="MM/AA"
                      maxlength="5"
                      class="w-full rounded-lg border border-neutral-300 px-4 py-3 focus:border-emerald-500 focus:ring-2 focus:ring-emerald-500/20 focus:outline-none"
                    />
                  </div>
                  <div>
                    <label
                      class="mb-2 block text-sm font-medium text-neutral-700"
                    >
                      CVV
                    </label>
                    <input
                      v-model="cardDetails.cvv"
                      @input="formatCVV"
                      type="text"
                      placeholder="123"
                      maxlength="4"
                      class="w-full rounded-lg border border-neutral-300 px-4 py-3 focus:border-emerald-500 focus:ring-2 focus:ring-emerald-500/20 focus:outline-none"
                    />
                  </div>
                </div>

                <div>
                  <label
                    class="mb-2 block text-sm font-medium text-neutral-700"
                  >
                    Número de Parcelas
                  </label>
                  <select
                    v-model.number="installments"
                    class="w-full rounded-lg border border-neutral-300 px-4 py-3 focus:border-emerald-500 focus:ring-2 focus:ring-emerald-500/20 focus:outline-none"
                  >
                    <option
                      v-for="n in availableInstallments"
                      :key="n"
                      :value="n"
                    >
                      {{ n }}x de {{ formatPrice(installmentPrice / n) }}
                      {{ n === 1 ? "à vista" : "sem juros" }}
                    </option>
                  </select>
                </div>
              </div>
            </div>

            <!-- Contract -->
            <div class="rounded-xl bg-white p-6 shadow-sm">
              <h2 class="mb-4 text-lg font-semibold text-neutral-900">
                Contrato de Compra de Energia
              </h2>

              <div
                class="max-h-96 space-y-4 overflow-y-auto rounded-lg bg-neutral-50 p-6 text-sm text-neutral-700"
              >
                <h3 class="font-semibold text-neutral-900">
                  CONTRATO DE COMPRA E VENDA DE ENERGIA ELÉTRICA
                </h3>

                <p>
                  <strong>1. DAS PARTES</strong><br />
                  Pelo presente instrumento particular de contrato de compra e
                  venda de energia elétrica, as partes:
                </p>

                <p>
                  <strong>VENDEDOR:</strong> Agente fornecedor identificado pelo
                  UUID {{ offer.seller_agent_uuid }}, regularmente inscrito e
                  autorizado pela CCEE (Câmara de Comercialização de Energia
                  Elétrica).
                </p>

                <p>
                  <strong>COMPRADOR:</strong> O usuário identificado e
                  autenticado nesta plataforma, doravante denominado
                  "COMPRADOR".
                </p>

                <p>
                  <strong>2. DO OBJETO</strong><br />
                  O presente contrato tem por objeto a compra e venda de energia
                  elétrica proveniente de fonte
                  {{ getTypeLabel(offer.energy_type) }}, nas seguintes
                  condições:
                </p>

                <ul class="list-inside list-disc space-y-1 pl-4">
                  <li>
                    Quantidade: {{ quantityMwh }} MWh ({{
                      formatQuantity(quantityMwh)
                    }})
                  </li>
                  <li>
                    Preço unitário: {{ formatPrice(offer.price_per_mwh) }}/MWh
                  </li>
                  <li>Valor total: {{ formatPrice(totalPrice) }}</li>
                  <li>Submercado: {{ offer.submarket }}</li>
                  <li>
                    Período de fornecimento: {{ offer.period_start }} até
                    {{ offer.period_end }}
                  </li>
                </ul>

                <p>
                  <strong>3. DO PREÇO E CONDIÇÕES DE PAGAMENTO</strong><br />
                  O COMPRADOR pagará ao VENDEDOR o valor total de
                  {{ formatPrice(totalPrice) }}, através de
                  <span v-if="paymentMethod === 'pix'"
                    >PIX com aprovação instantânea</span
                  >
                  <span v-else-if="paymentMethod === 'credit_card'">
                    cartão de crédito em {{ installments }}x de
                    {{ formatPrice(installmentPrice) }}
                  </span>
                  <span v-else
                    >boleto bancário com vencimento em 3 dias úteis</span
                  >.
                </p>

                <p>
                  <strong>4. DAS OBRIGAÇÕES DO VENDEDOR</strong><br />
                  O VENDEDOR obriga-se a fornecer a energia elétrica contratada,
                  nas quantidades e prazos estabelecidos, garantindo a qualidade
                  e continuidade do fornecimento conforme regulamentação da
                  ANEEL.
                </p>

                <p>
                  <strong>5. DAS OBRIGAÇÕES DO COMPRADOR</strong><br />
                  O COMPRADOR obriga-se a efetuar o pagamento nas condições
                  estabelecidas e a utilizar a energia elétrica adquirida em
                  conformidade com as normas técnicas e regulamentares
                  aplicáveis.
                </p>

                <p>
                  <strong>6. DA VIGÊNCIA</strong><br />
                  O presente contrato tem vigência pelo período de fornecimento
                  estabelecido, iniciando-se em {{ offer.period_start }} e
                  encerrando-se em {{ offer.period_end }}.
                </p>

                <p>
                  <strong>7. DA RESCISÃO</strong><br />
                  O presente contrato poderá ser rescindido por qualquer das
                  partes mediante notificação prévia de 30 (trinta) dias, sem
                  prejuízo das obrigações já assumidas até a data da rescisão.
                </p>

                <p>
                  <strong>8. DAS DISPOSIÇÕES GERAIS</strong><br />
                  As partes elegem o foro da comarca de [LOCAL] para dirimir
                  quaisquer controvérsias oriundas do presente contrato, com
                  renúncia expressa a qualquer outro, por mais privilegiado que
                  seja.
                </p>

                <p class="text-xs text-neutral-500 italic">
                  Este é um contrato fictício para fins de demonstração da
                  plataforma. Nenhuma transação real será realizada.
                </p>
              </div>

              <div class="mt-4 flex items-start gap-3">
                <input
                  id="accept-contract"
                  type="checkbox"
                  class="mt-1 h-5 w-5 rounded border-neutral-300 text-emerald-500 focus:ring-emerald-500"
                />
                <label for="accept-contract" class="text-sm text-neutral-700">
                  Li e aceito os termos do contrato de compra de energia
                  elétrica
                </label>
              </div>
            </div>
          </div>
        </div>

        <!-- Right Column - Summary -->
        <div class="lg:col-span-1">
          <div class="sticky top-8 space-y-6">
            <!-- Order Summary -->
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
                      {{ formatQuantity(quantityMwh) }}
                    </span>
                  </div>

                  <div class="flex justify-between text-sm">
                    <span class="text-neutral-600">Preço por MWh</span>
                    <span class="font-medium text-neutral-900">
                      {{ formatPrice(offer.price_per_mwh) }}
                    </span>
                  </div>

                  <div
                    v-if="paymentMethod === 'credit_card' && installments > 1"
                    class="flex justify-between text-sm"
                  >
                    <span class="text-neutral-600">Parcelas</span>
                    <span class="font-medium text-neutral-900">
                      {{ installments }}x sem juros
                    </span>
                  </div>
                </div>

                <div
                  class="flex justify-between border-t-2 border-neutral-200 pt-4"
                >
                  <span class="text-lg font-semibold text-neutral-900"
                    >Total</span
                  >
                  <span class="text-2xl font-bold text-emerald-600">
                    {{ formatPrice(totalPrice) }}
                  </span>
                </div>

                <div
                  v-if="paymentMethod === 'credit_card' && installments > 1"
                  class="text-center text-sm text-neutral-500"
                >
                  {{ installments }}x de {{ formatPrice(installmentPrice) }}
                </div>
              </div>
            </div>

            <!-- Security Badge -->
            <div class="rounded-xl bg-emerald-50 p-6">
              <div class="flex items-start gap-3">
                <div class="rounded-lg bg-emerald-100 p-2">
                  <ShieldCheck class="text-emerald-600" :size="20" />
                </div>
                <div>
                  <h3 class="font-semibold text-emerald-900">Compra Segura</h3>
                  <p class="mt-1 text-sm text-emerald-700">
                    Seus dados estão protegidos com criptografia de ponta a
                    ponta
                  </p>
                </div>
              </div>
            </div>

            <!-- Action Button -->
            <button
              @click="processPayment"
              :disabled="!isFormValid || processing"
              :class="[
                'w-full rounded-xl py-4 font-semibold text-white shadow-lg transition-all',
                isFormValid && !processing
                  ? 'bg-gradient-to-r from-emerald-500 to-green-500 hover:shadow-xl hover:brightness-110'
                  : 'cursor-not-allowed bg-neutral-300',
              ]"
            >
              <span
                v-if="processing"
                class="flex items-center justify-center gap-2"
              >
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

            <p class="text-center text-xs text-neutral-500">
              Ao finalizar, você concorda com os termos do contrato
            </p>
          </div>
        </div>
      </div>
    </main>
  </div>
</template>

<style scoped>
/* Custom scrollbar for contract */
.overflow-y-auto::-webkit-scrollbar {
  width: 8px;
}

.overflow-y-auto::-webkit-scrollbar-track {
  background: #f5f5f5;
  border-radius: 4px;
}

.overflow-y-auto::-webkit-scrollbar-thumb {
  background: #d4d4d4;
  border-radius: 4px;
}

.overflow-y-auto::-webkit-scrollbar-thumb:hover {
  background: #a3a3a3;
}

/* Input number arrows removal for cleaner look */
input[type="number"]::-webkit-inner-spin-button,
input[type="number"]::-webkit-outer-spin-button {
  -webkit-appearance: none;
  margin: 0;
}

input[type="number"] {
  -moz-appearance: textfield;
  appearance: textfield;
}
</style>
