<script setup lang="ts">
import { computed } from "vue";
import Dialog from "primevue/dialog";
import { AlertTriangle, CheckCircle, Info } from "lucide-vue-next";

interface ConfirmDialogProps {
  visible: boolean;
  title?: string;
  message?: string;
  confirmText?: string;
  cancelText?: string;
  variant?: "danger" | "warning" | "info" | "success";
  loading?: boolean;
}

interface ConfirmDialogEmits {
  (e: "update:visible", value: boolean): void;
  (e: "confirm"): void;
  (e: "cancel"): void;
}

const props = withDefaults(defineProps<ConfirmDialogProps>(), {
  title: "Confirmar ação",
  message: "Tem certeza que deseja realizar esta ação?",
  confirmText: "Confirmar",
  cancelText: "Cancelar",
  variant: "warning",
  loading: false,
});

const emit = defineEmits<ConfirmDialogEmits>();

const isVisible = computed({
  get: () => props.visible,
  set: (value) => emit("update:visible", value),
});

const variantConfig = computed(() => {
  const configs = {
    danger: {
      gradient: "from-red-500 to-rose-600",
      iconBg: "bg-red-100",
      iconColor: "text-red-600",
      confirmBg: "bg-red-600 hover:bg-red-700",
      confirmBgDisabled: "bg-red-400",
      icon: AlertTriangle,
    },
    warning: {
      gradient: "from-amber-500 to-orange-600",
      iconBg: "bg-amber-100",
      iconColor: "text-amber-600",
      confirmBg: "bg-amber-600 hover:bg-amber-700",
      confirmBgDisabled: "bg-amber-400",
      icon: AlertTriangle,
    },
    info: {
      gradient: "from-blue-500 to-indigo-600",
      iconBg: "bg-blue-100",
      iconColor: "text-blue-600",
      confirmBg: "bg-blue-600 hover:bg-blue-700",
      confirmBgDisabled: "bg-blue-400",
      icon: Info,
    },
    success: {
      gradient: "from-emerald-500 to-green-600",
      iconBg: "bg-emerald-100",
      iconColor: "text-emerald-600",
      confirmBg: "bg-emerald-600 hover:bg-emerald-700",
      confirmBgDisabled: "bg-emerald-400",
      icon: CheckCircle,
    },
  };

  return configs[props.variant];
});

const handleConfirm = () => {
  emit("confirm");
};

const handleCancel = () => {
  isVisible.value = false;
  emit("cancel");
};
</script>

<template>
  <Dialog
    v-model:visible="isVisible"
    modal
    :draggable="false"
    :closable="!loading"
    :closeOnEscape="!loading"
    :style="{ width: '90vw', maxWidth: '28rem' }"
    :breakpoints="{ '640px': '90vw' }"
    class="mx-4 sm:mx-0"
  >
    <template #header>
      <div class="flex items-center gap-2 px-2 sm:gap-3 sm:px-0">
        <div
          :class="['shrink-0 rounded-full p-1.5 sm:p-2', variantConfig.iconBg]"
        >
          <component
            :is="variantConfig.icon"
            :size="20"
            :class="['sm:hidden', variantConfig.iconColor]"
          />
          <component
            :is="variantConfig.icon"
            :size="24"
            :class="['hidden sm:block', variantConfig.iconColor]"
          />
        </div>
        <h3
          class="min-w-0 flex-1 text-base font-bold text-neutral-900 sm:text-lg md:text-xl"
        >
          {{ title }}
        </h3>
      </div>
    </template>

    <div class="px-2 py-3 sm:px-0 sm:py-4 md:py-6">
      <p class="text-sm leading-relaxed text-neutral-600 sm:text-base">
        {{ message }}
      </p>
    </div>

    <template #footer>
      <div class="flex flex-col gap-2.5 px-2 sm:flex-row sm:gap-3 sm:px-0">
        <button
          v-if="cancelText"
          @click="handleCancel"
          :disabled="loading"
          class="order-2 min-h-[44px] flex-1 touch-manipulation rounded-lg border-2 border-neutral-300 bg-white px-4 py-2.5 text-sm font-medium text-neutral-700 transition-all duration-200 hover:border-neutral-400 hover:bg-neutral-50 active:scale-[0.98] disabled:cursor-not-allowed disabled:opacity-50 sm:order-1 sm:px-5 sm:py-3 md:px-6"
        >
          {{ cancelText }}
        </button>
        <button
          @click="handleConfirm"
          :disabled="loading"
          :class="[
            cancelText ? 'order-1 flex-1 sm:order-2' : 'w-full',
            'min-h-[44px] touch-manipulation rounded-lg px-4 py-2.5 text-sm font-medium text-white shadow-md transition-all duration-200 hover:shadow-lg active:scale-[0.98] disabled:cursor-not-allowed disabled:shadow-none sm:px-5 sm:py-3 md:px-6',
            loading ? variantConfig.confirmBgDisabled : variantConfig.confirmBg,
          ]"
        >
          <span
            v-if="loading"
            class="flex items-center justify-center gap-1.5 sm:gap-2"
          >
            <div
              class="h-4 w-4 animate-spin rounded-full border-2 border-white border-t-transparent"
            ></div>
            <span>Processando...</span>
          </span>
          <span v-else>{{ confirmText }}</span>
        </button>
      </div>
    </template>
  </Dialog>
</template>
