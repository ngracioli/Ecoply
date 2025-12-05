<script setup lang="ts">
import { reactive, ref, watch, computed } from "vue";
import Dialog from "primevue/dialog";
import { useToast } from "primevue/usetoast";
import { Calendar, DollarSign, Package, Zap, Plus } from "lucide-vue-next";
import { useIconSize } from "../../../../composables/useIconSize";
import api from "../../../../axios";
import { OFFER_ENDPOINTS } from "../../../../api/endpoints";
import type { EnergyType } from "../../../../types/offer";
import type { CreateOfferRequest } from "../../../../types/requests/offer";
import type { CreateOfferResponse } from "../../../../types/responses/offer";
import type { OfferListItem } from "../../../../types/responses/offers";
import { offerSchema } from "./validation";

interface OfferForm {
  price_per_mwh: string;
  quantity_mwh: string;
  period_start: string;
  period_end: string;
  description: string;
  energy_type: EnergyType;
}

interface OfferFormErrors {
  price_per_mwh?: string;
  quantity_mwh?: string;
  period_start?: string;
  period_end?: string;
  description?: string;
  energy_type?: string;
}

interface CreateOfferDialogProps {
  visible: boolean;
  offer?: OfferListItem | null;
}

interface CreateOfferDialogEmits {
  (e: "update:visible", value: boolean): void;
  (e: "offer-created"): void;
  (e: "offer-updated"): void;
}

interface EnergyTypeConfig {
  label: string;
  gradient: string;
  borderColor: string;
  bgColor: string;
  hoverBorder: string;
  hoverBg: string;
  ringColor: string;
}

const INITIAL_FORM_STATE: OfferForm = {
  price_per_mwh: "",
  quantity_mwh: "",
  period_start: "",
  period_end: "",
  description: "",
  energy_type: "solar",
};

const ENERGY_TYPES_CONFIG: Record<EnergyType, EnergyTypeConfig> = {
  solar: {
    label: "Solar",
    gradient: "from-amber-400 to-orange-500",
    borderColor: "border-amber-500",
    bgColor: "bg-amber-50",
    hoverBorder: "hover:border-amber-300",
    hoverBg: "hover:bg-amber-50/50",
    ringColor: "ring-amber-500/10",
  },
  eolic: {
    label: "Eólica",
    gradient: "from-cyan-400 to-blue-500",
    borderColor: "border-cyan-500",
    bgColor: "bg-cyan-50",
    hoverBorder: "hover:border-cyan-300",
    hoverBg: "hover:bg-cyan-50/50",
    ringColor: "ring-cyan-500/10",
  },
  hydroelectric: {
    label: "Hidrelétrica",
    gradient: "from-blue-500 to-indigo-600",
    borderColor: "border-blue-600",
    bgColor: "bg-blue-50",
    hoverBorder: "hover:border-blue-400",
    hoverBg: "hover:bg-blue-50/50",
    ringColor: "ring-blue-600/10",
  },
  geothermal: {
    label: "Geotérmica",
    gradient: "from-red-500 to-orange-600",
    borderColor: "border-red-600",
    bgColor: "bg-red-50",
    hoverBorder: "hover:border-red-400",
    hoverBg: "hover:bg-red-50/50",
    ringColor: "ring-red-600/10",
  },
};

const props = defineProps<CreateOfferDialogProps>();
const emit = defineEmits<CreateOfferDialogEmits>();
const toast = useToast();
const { iconSize } = useIconSize();

const newOffer = reactive<OfferForm>({ ...INITIAL_FORM_STATE });
const originalOffer = ref<OfferForm | null>(null);
const isSubmitting = ref(false);
const errors = ref<OfferFormErrors>({});

const isEditMode = computed(() => !!props.offer);

const dialogTitle = computed(() =>
  isEditMode.value ? "Editar Oferta" : "Criar Nova Oferta",
);

const dialogSubtitle = computed(() =>
  isEditMode.value
    ? "Atualize os dados da sua oferta de energia"
    : "Preencha os dados da sua oferta de energia",
);

const submitButtonText = computed(() => {
  if (isSubmitting.value) {
    return isEditMode.value ? "Atualizando..." : "Criando...";
  }
  return isEditMode.value ? "Atualizar Oferta" : "Criar Oferta";
});

const getInputClass = (error?: string): string => {
  const baseClasses =
    "transition-all duration-200 placeholder:text-neutral-400 focus:ring-4 focus:outline-none";
  const errorClasses = error
    ? "border-red-500 focus:border-red-500 focus:ring-red-500/10"
    : "border-neutral-200 focus:border-emerald-500 focus:ring-emerald-500/10";
  return `${baseClasses} ${errorClasses}`;
};

const validateField = (fieldName: keyof OfferFormErrors) => {
  const fieldValue = newOffer[fieldName];

  if (fieldName === "price_per_mwh" || fieldName === "quantity_mwh") {
    if (!fieldValue || fieldValue === "") {
      errors.value[fieldName] = undefined;
      return;
    }

    const schema =
      fieldName === "price_per_mwh"
        ? offerSchema.shape.price_per_mwh
        : offerSchema.shape.quantity_mwh;

    const result = schema.safeParse(fieldValue);

    if (!result.success) {
      errors.value[fieldName] = result.error.issues[0]?.message;
    } else {
      errors.value[fieldName] = undefined;
    }
  }
};

const resetForm = () => {
  Object.assign(newOffer, INITIAL_FORM_STATE);
  errors.value = {};
};

const loadOfferData = () => {
  if (props.offer) {
    newOffer.price_per_mwh = props.offer.price_per_mwh.toString();
    newOffer.quantity_mwh = props.offer.remaining_quantity_mwh.toString();
    newOffer.period_start = props.offer.period_start;
    newOffer.period_end = props.offer.period_end;
    newOffer.description = props.offer.description;
    newOffer.energy_type = props.offer.energy_type;

    originalOffer.value = {
      price_per_mwh: props.offer.price_per_mwh.toString(),
      quantity_mwh: props.offer.remaining_quantity_mwh.toString(),
      period_start: props.offer.period_start,
      period_end: props.offer.period_end,
      description: props.offer.description,
      energy_type: props.offer.energy_type,
    };
  } else {
    resetForm();
    originalOffer.value = null;
  }
};

const closeDialog = () => {
  emit("update:visible", false);
  resetForm();
};

const showToast = (severity: "success" | "error", message: string) => {
  toast.add({
    severity,
    summary: severity === "success" ? "Sucesso" : "Erro",
    detail: message,
    life: 3000,
  });
};

const handleSubmit = async () => {
  errors.value = {};

  if (isEditMode.value && originalOffer.value) {
    const fieldsToValidate: Partial<OfferForm> = {};

    if (newOffer.price_per_mwh !== originalOffer.value.price_per_mwh) {
      fieldsToValidate.price_per_mwh = newOffer.price_per_mwh;
    }
    if (newOffer.quantity_mwh !== originalOffer.value.quantity_mwh) {
      fieldsToValidate.quantity_mwh = newOffer.quantity_mwh;
    }
    if (newOffer.period_start !== originalOffer.value.period_start) {
      fieldsToValidate.period_start = newOffer.period_start;
    }
    if (newOffer.period_end !== originalOffer.value.period_end) {
      fieldsToValidate.period_end = newOffer.period_end;
    }
    if (newOffer.description !== originalOffer.value.description) {
      fieldsToValidate.description = newOffer.description;
    }
    if (newOffer.energy_type !== originalOffer.value.energy_type) {
      fieldsToValidate.energy_type = newOffer.energy_type;
    }

    if (Object.keys(fieldsToValidate).length === 0) {
      showToast("error", "Nenhuma alteração foi feita.");
      return;
    }

    const result = offerSchema.partial().safeParse(fieldsToValidate);

    if (!result.success) {
      result.error.issues.forEach((err) => {
        const field = err.path[0] as keyof OfferFormErrors;
        if (field) {
          errors.value[field] = err.message;
        }
      });
      showToast("error", "Por favor, corrija os erros no formulário.");
      return;
    }
  } else {
    const result = offerSchema.safeParse(newOffer);

    if (!result.success) {
      result.error.issues.forEach((err) => {
        const field = err.path[0] as keyof OfferFormErrors;
        if (field) {
          errors.value[field] = err.message;
        }
      });
      showToast("error", "Por favor, corrija os erros no formulário.");
      return;
    }
  }

  try {
    isSubmitting.value = true;

    if (isEditMode.value && props.offer) {
      const payload: CreateOfferRequest = {
        price_per_mwh: parseFloat(newOffer.price_per_mwh),
        quantity_mwh: parseFloat(newOffer.quantity_mwh),
        period_start: newOffer.period_start,
        period_end: newOffer.period_end,
        description: newOffer.description,
        energy_type: newOffer.energy_type,
      };

      await api.put<CreateOfferResponse>(
        OFFER_ENDPOINTS.UPDATE(props.offer.uuid),
        payload,
      );
      showToast("success", "Oferta atualizada com sucesso!");
      emit("offer-updated");
    } else {
      const payload: CreateOfferRequest = {
        price_per_mwh: parseFloat(newOffer.price_per_mwh),
        quantity_mwh: parseFloat(newOffer.quantity_mwh),
        period_start: newOffer.period_start,
        period_end: newOffer.period_end,
        description: newOffer.description,
        energy_type: newOffer.energy_type,
      };

      await api.post<CreateOfferResponse>(OFFER_ENDPOINTS.CREATE, payload);
      showToast("success", "Oferta criada com sucesso!");
      emit("offer-created");
    }

    closeDialog();
  } catch (error) {
    const errorMessage = isEditMode.value
      ? "Erro ao atualizar oferta. Por favor, tente novamente."
      : "Erro ao criar oferta. Por favor, tente novamente.";
    showToast("error", errorMessage);
  } finally {
    isSubmitting.value = false;
  }
};

watch(
  () => newOffer.price_per_mwh,
  () => {
    validateField("price_per_mwh");
  },
);

watch(
  () => newOffer.quantity_mwh,
  () => {
    validateField("quantity_mwh");
  },
);

watch(
  () => props.visible,
  (isVisible) => {
    if (isVisible) {
      loadOfferData();
    } else {
      resetForm();
    }
  },
);

watch(
  () => props.offer,
  () => {
    if (props.visible) {
      loadOfferData();
    }
  },
);
</script>

<template>
  <Dialog
    v-if="visible"
    :visible="visible"
    @update:visible="emit('update:visible', $event)"
    modal
    :closable="true"
    :draggable="false"
    :style="{ width: '50rem' }"
    :breakpoints="{ '1199px': '75vw', '575px': '95vw' }"
    class="p-0"
  >
    <template #header>
      <div class="flex items-center gap-2.5 px-2 sm:gap-3 sm:px-0">
        <div
          class="flex h-10 w-10 shrink-0 items-center justify-center rounded-lg bg-emerald-100 p-1.5 sm:h-auto sm:w-auto sm:p-2"
        >
          <Plus :size="iconSize(20, 24)" class="text-emerald-600" />
        </div>
        <div class="min-w-0 flex-1">
          <h3 class="truncate text-lg font-bold text-neutral-900 sm:text-xl">
            {{ dialogTitle }}
          </h3>
          <p class="truncate text-xs text-neutral-500 sm:text-sm">
            {{ dialogSubtitle }}
          </p>
        </div>
      </div>
    </template>

    <div class="space-y-5 px-4 py-3 sm:space-y-6 sm:px-6 sm:py-4">
      <div class="group">
        <label
          for="offer-type"
          class="mb-2 flex items-center gap-1.5 text-xs font-semibold text-neutral-700 sm:gap-2 sm:text-sm"
        >
          <Zap :size="iconSize(14, 16)" class="text-emerald-600" />
          Tipo de Energia
          <span class="text-red-500">*</span>
        </label>
        <div class="grid grid-cols-2 gap-2.5 sm:grid-cols-4 sm:gap-3">
          <label
            v-for="(config, type) in ENERGY_TYPES_CONFIG"
            :key="type"
            :class="[
              'flex cursor-pointer touch-manipulation flex-col items-center gap-1.5 rounded-lg border-2 p-3 transition-all duration-200 sm:gap-2 sm:p-4',
              newOffer.energy_type === type
                ? `${config.borderColor} ${config.bgColor} ring-4 ${config.ringColor}`
                : `border-neutral-200 ${config.hoverBorder} ${config.hoverBg}`,
            ]"
          >
            <input
              v-model="newOffer.energy_type"
              type="radio"
              :value="type"
              class="hidden"
            />
            <div
              :class="[
                'flex h-10 w-10 items-center justify-center rounded-full bg-gradient-to-br sm:h-12 sm:w-12',
                config.gradient,
              ]"
            >
              <Zap :size="iconSize(18, 20)" class="text-white" />
            </div>
            <span
              class="text-center text-xs font-medium text-neutral-900 sm:text-sm"
            >
              {{ config.label }}
            </span>
          </label>
        </div>
        <small
          v-if="errors.energy_type"
          class="mt-1 block text-xs text-red-600 sm:text-sm"
        >
          {{ errors.energy_type }}
        </small>
      </div>

      <div class="grid grid-cols-1 gap-4 sm:grid-cols-2">
        <div class="group">
          <label
            for="offer-price"
            class="mb-2 flex items-center gap-1.5 text-xs font-semibold text-neutral-700 sm:gap-2 sm:text-sm"
          >
            <DollarSign :size="iconSize(14, 16)" class="text-emerald-600" />
            Preço por MWh
            <span class="text-red-500">*</span>
          </label>
          <div class="relative">
            <span
              class="absolute top-1/2 left-3 -translate-y-1/2 text-sm text-neutral-500 sm:left-4 sm:text-base"
            >
              R$
            </span>
            <input
              id="offer-price"
              v-model="newOffer.price_per_mwh"
              type="number"
              step="0.01"
              min="0.01"
              max="99999999.99"
              placeholder="10.35"
              :class="getInputClass(errors.price_per_mwh)"
              class="w-full min-w-0 touch-manipulation rounded-lg border-2 py-2.5 pr-3 pl-10 text-base sm:py-3 sm:pr-4 sm:pl-12"
            />
          </div>
          <small
            v-if="errors.price_per_mwh"
            class="mt-1 block text-xs text-red-600 sm:text-sm"
          >
            {{ errors.price_per_mwh }}
          </small>
        </div>

        <div class="group">
          <label
            for="offer-quantity"
            class="mb-2 flex items-center gap-1.5 text-xs font-semibold text-neutral-700 sm:gap-2 sm:text-sm"
          >
            <Package :size="iconSize(14, 16)" class="text-emerald-600" />
            Quantidade
            <span class="text-red-500">*</span>
          </label>
          <div class="relative">
            <input
              id="offer-quantity"
              v-model="newOffer.quantity_mwh"
              type="number"
              step="0.001"
              min="0.001"
              max="9999999.999"
              placeholder="100.35"
              :class="getInputClass(errors.quantity_mwh)"
              class="w-full min-w-0 touch-manipulation rounded-lg border-2 px-3 py-2.5 pr-14 text-base sm:px-4 sm:py-3 sm:pr-16"
            />
            <span
              class="absolute top-1/2 right-3 -translate-y-1/2 text-xs text-neutral-500 sm:right-4 sm:text-sm"
            >
              MWh
            </span>
          </div>
          <small
            v-if="errors.quantity_mwh"
            class="mt-1 block text-xs text-red-600 sm:text-sm"
          >
            {{ errors.quantity_mwh }}
          </small>
        </div>
      </div>

      <div class="grid grid-cols-1 gap-4 sm:grid-cols-2">
        <div class="group min-w-0">
          <label
            for="period-start"
            class="mb-2 flex items-center gap-1.5 text-xs font-semibold text-neutral-700 sm:gap-2 sm:text-sm"
          >
            <Calendar :size="iconSize(14, 16)" class="text-emerald-600" />
            Data de Início
            <span class="text-red-500">*</span>
          </label>
          <div class="relative overflow-hidden">
            <input
              id="period-start"
              v-model="newOffer.period_start"
              type="date"
              lang="pt-BR"
              :class="getInputClass(errors.period_start)"
              class="box-border w-full max-w-full min-w-0 touch-manipulation rounded-lg border-2 px-2 py-2.5 pr-8 text-base sm:px-4 sm:py-3"
              style="font-size: 16px"
            />
          </div>
          <small
            v-if="errors.period_start"
            class="mt-1 block text-xs text-red-600 sm:text-sm"
          >
            {{ errors.period_start }}
          </small>
        </div>

        <div class="group min-w-0">
          <label
            for="period-end"
            class="mb-2 flex items-center gap-1.5 text-xs font-semibold text-neutral-700 sm:gap-2 sm:text-sm"
          >
            <Calendar :size="iconSize(14, 16)" class="text-emerald-600" />
            Data de Término
            <span class="text-red-500">*</span>
          </label>
          <div class="relative overflow-hidden">
            <input
              id="period-end"
              v-model="newOffer.period_end"
              type="date"
              lang="pt-BR"
              :class="getInputClass(errors.period_end)"
              class="box-border w-full max-w-full min-w-0 touch-manipulation rounded-lg border-2 px-2 py-2.5 pr-8 text-base sm:px-4 sm:py-3"
              style="font-size: 16px"
            />
          </div>
          <small
            v-if="errors.period_end"
            class="mt-1 block text-xs text-red-600 sm:text-sm"
          >
            {{ errors.period_end }}
          </small>
        </div>
      </div>

      <div class="group">
        <label
          for="offer-description"
          class="mb-2 flex items-center gap-1.5 text-xs font-semibold text-neutral-700 sm:gap-2 sm:text-sm"
        >
          Descrição
          <span class="text-red-500">*</span>
        </label>
        <textarea
          id="offer-description"
          v-model="newOffer.description"
          rows="3"
          placeholder="Descreva sua oferta de energia renovável..."
          :class="getInputClass(errors.description)"
          class="w-full touch-manipulation resize-none rounded-lg border-2 px-3 py-2.5 text-base sm:px-4 sm:py-3"
        ></textarea>
        <div class="mt-1 flex items-center justify-between gap-2">
          <small
            v-if="errors.description"
            class="min-w-0 flex-1 truncate text-xs text-red-600 sm:text-sm"
          >
            {{ errors.description }}
          </small>
          <small
            class="shrink-0 text-xs text-neutral-500 sm:text-sm"
            :class="{ 'ml-auto': !errors.description }"
          >
            {{ newOffer.description.length }}/500
          </small>
        </div>
      </div>
    </div>

    <template #footer>
      <div
        class="flex flex-col items-stretch justify-between gap-3 border-t border-neutral-100 px-4 py-3 sm:flex-row sm:items-center sm:gap-4 sm:px-6 sm:py-4"
      >
        <p
          class="order-2 text-center text-xs text-neutral-500 sm:order-1 sm:text-left"
        >
          <span class="text-red-500">*</span> Campos obrigatórios
        </p>
        <div
          class="order-1 flex w-full shrink-0 gap-2.5 sm:order-2 sm:w-auto sm:gap-3"
        >
          <button
            @click="closeDialog"
            :disabled="isSubmitting"
            class="min-h-[44px] flex-1 touch-manipulation rounded-lg border-2 border-neutral-300 px-4 py-2.5 text-sm font-medium text-neutral-700 transition-all duration-200 hover:bg-neutral-50 active:scale-[0.98] disabled:cursor-not-allowed disabled:opacity-50 sm:flex-initial sm:px-5"
          >
            Cancelar
          </button>
          <button
            @click="handleSubmit"
            :disabled="isSubmitting"
            class="min-h-[44px] flex-1 touch-manipulation rounded-lg bg-gradient-to-r from-emerald-500 to-teal-600 px-4 py-2.5 text-sm font-medium whitespace-nowrap text-white shadow-md transition-all duration-200 hover:shadow-lg hover:brightness-110 active:scale-[0.98] disabled:cursor-not-allowed disabled:opacity-50 sm:flex-initial sm:px-5"
          >
            {{ submitButtonText }}
          </button>
        </div>
      </div>
    </template>
  </Dialog>
</template>

<style scoped>
input[type="date"] {
  appearance: none;
  -webkit-appearance: none;
  -moz-appearance: none;
}

input[type="date"]::-webkit-calendar-picker-indicator {
  position: absolute;
  right: 8px;
  width: 20px;
  height: 20px;
  cursor: pointer;
  flex-shrink: 0;
}

input[type="date"]::-webkit-date-and-time-value {
  text-align: left;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  margin-right: 28px;
}

@media (max-width: 639px) {
  input[type="date"] {
    font-size: 16px !important;
    padding-right: 32px !important;
  }
}

@media (max-width: 768px) {
  :deep(.p-dialog-content) {
    max-height: 60vh;
    overflow-y: auto !important;
    -webkit-overflow-scrolling: touch;
  }
}
</style>
