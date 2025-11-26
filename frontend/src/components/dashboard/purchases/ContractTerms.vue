<script setup lang="ts">
import { computed } from "vue";
import type { OfferListItem } from "../../../types/responses/offers";
import type { PaymentMethod } from "../../../types/checkout";
import SubmarketBadge from "../../shared/SubmarketBadge.vue";

interface Props {
  modelValue: boolean;
  offer: OfferListItem;
  quantityMwh: number | null;
  pricePerMwh: number;
  totalPrice: number;
  paymentMethod: PaymentMethod;
  installments: number;
  installmentPrice: number;
}

interface Emits {
  (e: "update:modelValue", value: boolean): void;
}

const props = defineProps<Props>();
const emit = defineEmits<Emits>();

const acceptedTerms = computed({
  get: () => props.modelValue,
  set: (value) => emit("update:modelValue", value),
});

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
</script>

<template>
  <div class="rounded-xl bg-white p-6 shadow-sm">
    <h2 class="mb-4 text-lg font-semibold text-neutral-900">
      Termos e Condições de Compra e Uso
    </h2>

    <div
      class="max-h-96 space-y-4 overflow-y-auto rounded-lg bg-neutral-50 p-6 text-sm text-neutral-700"
    >
      <h3 class="font-semibold text-neutral-900">
        TERMOS E CONDIÇÕES DE COMPRA E USO
      </h3>

      <p>
        <strong>1. Condições Gerais</strong><br />
        Ao efetuar uma compra neste site, o cliente declara estar ciente e de
        acordo com todas as condições aqui descritas. O uso da plataforma
        implica aceitação integral destes Termos.
      </p>

      <p>
        <strong>2. Cadastro e Responsabilidade do Usuário</strong><br />
        O cliente se compromete a fornecer informações verdadeiras e
        atualizadas. A empresa não se responsabiliza por erros decorrentes de
        dados incorretos ou incompletos no momento da compra.
      </p>

      <p>
        <strong>3. Pagamento e Processamento do Pedido</strong><br />
        Os pedidos serão processados após a confirmação do pagamento. O não
        pagamento no prazo resultará no cancelamento automático. Todos os
        valores incluem tributos conforme a legislação vigente.
      </p>

      <div class="rounded-lg border border-emerald-200 bg-emerald-50 p-4">
        <p class="mb-2 font-semibold text-emerald-900">
          Detalhes da sua compra:
        </p>
        <ul class="list-inside list-disc space-y-1 pl-2 text-emerald-800">
          <li>Tipo de energia: {{ getTypeLabel(offer.energy_type) }}</li>
          <li>
            Quantidade: {{ quantityMwh ? formatQuantity(quantityMwh) : "—" }}
          </li>
          <li>Preço unitário: {{ formatPrice(pricePerMwh) }}/MWh</li>
          <li>Valor total: {{ formatPrice(totalPrice) }}</li>
          <li class="flex items-center gap-2">
            <span>Submercado:</span>
            <SubmarketBadge
              :offer-submarket="offer.submarket"
              size="small"
            />
          </li>
          <li>Período: {{ offer.period_start }} até {{ offer.period_end }}</li>
          <li>
            Forma de pagamento:
            <span v-if="paymentMethod === 'pix'"
              >PIX com aprovação instantânea</span
            >
            <span v-else-if="paymentMethod === 'card'">
              Cartão de crédito em {{ installments }}x de
              {{ formatPrice(installmentPrice) }}
            </span>
            <span v-else>Boleto bancário com vencimento em 3 dias úteis</span>
          </li>
        </ul>
      </div>

      <p>
        <strong>4. Fornecimento e Prazos</strong><br />
        O fornecimento de energia elétrica será realizado conforme o período
        contratado e disponibilidade na rede de distribuição. A empresa se
        compromete a informar previamente sobre qualquer alteração nos prazos ou
        condições de fornecimento.
      </p>

      <p>
        <strong>5. Cancelamento</strong><br />
        O cancelamento do contrato poderá ser solicitado conforme as condições
        estabelecidas e regulamentação vigente do setor elétrico brasileiro. O
        cliente deve estar ciente de que podem haver restrições de cancelamento
        dependendo do estágio da operação.
      </p>

      <p>
        <strong>6. Política de Privacidade</strong><br />
        Os dados coletados são utilizados exclusivamente para o processamento de
        pedidos e comunicações relacionadas. As informações são tratadas de
        forma confidencial, seguindo a legislação de proteção de dados (LGPD).
      </p>

      <p>
        <strong>7. Propriedade Intelectual</strong><br />
        Todos os direitos de propriedade intelectual relacionados à plataforma e
        seus serviços são reservados, incluindo marcas, patentes e direitos
        autorais aplicáveis.
      </p>

      <p>
        <strong>8. Alterações dos Termos</strong><br />
        A empresa reserva-se o direito de alterar estes Termos a qualquer
        momento, sendo as mudanças válidas a partir da publicação no site.
      </p>

      <p>
        <strong>9. Foro e Legislação Aplicável</strong><br />
        Em caso de controvérsia, aplica-se a legislação brasileira, elegendo-se
        o foro da comarca do domicílio do consumidor.
      </p>

      <p class="mt-6 text-xs text-neutral-500 italic">
        Este é um ambiente de demonstração. Nenhuma transação real será
        realizada.
      </p>
    </div>

    <div class="mt-4 flex items-start gap-3">
      <input
        id="accept-contract"
        v-model="acceptedTerms"
        type="checkbox"
        class="mt-1 h-5 w-5 rounded border-neutral-300 text-emerald-500 focus:ring-emerald-500"
      />
      <label for="accept-contract" class="text-sm text-neutral-700">
        Li e aceito os termos e condições de compra e uso
        <span class="text-red-500">*</span>
      </label>
    </div>
  </div>
</template>

<style scoped>
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
</style>
