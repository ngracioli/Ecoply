<script setup lang="ts">
import { computed } from "vue";
import Dialog from "primevue/dialog";
import { AlertTriangle } from "lucide-vue-next";

interface ConfirmDialogProps {
  visible: boolean;
  title?: string;
  message?: string;
  confirmText?: string;
  cancelText?: string;
  variant?: "danger" | "warning" | "info";
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
    },
    warning: {
      gradient: "from-amber-500 to-orange-600",
      iconBg: "bg-amber-100",
      iconColor: "text-amber-600",
      confirmBg: "bg-amber-600 hover:bg-amber-700",
      confirmBgDisabled: "bg-amber-400",
    },
    info: {
      gradient: "from-blue-500 to-indigo-600",
      iconBg: "bg-blue-100",
      iconColor: "text-blue-600",
      confirmBg: "bg-blue-600 hover:bg-blue-700",
      confirmBgDisabled: "bg-blue-400",
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
    class="w-full max-w-md"
  >
    <template #header>
      <div class="flex items-center gap-3">
        <div :class="['rounded-full p-2', variantConfig.iconBg]">
          <AlertTriangle :size="24" :class="variantConfig.iconColor" />
        </div>
        <h3 class="text-xl font-bold text-neutral-900">
          {{ title }}
        </h3>
      </div>
    </template>

    <div class="py-4">
      <p class="text-sm leading-relaxed text-neutral-600">
        {{ message }}
      </p>
    </div>

    <template #footer>
      <div class="flex gap-3">
        <button
          @click="handleCancel"
          :disabled="loading"
          class="flex-1 rounded-lg border-2 border-neutral-300 bg-white px-4 py-2.5 text-sm font-medium text-neutral-700 transition-all duration-200 hover:border-neutral-400 hover:bg-neutral-50 disabled:cursor-not-allowed disabled:opacity-50"
        >
          {{ cancelText }}
        </button>
        <button
          @click="handleConfirm"
          :disabled="loading"
          :class="[
            'flex-1 rounded-lg px-4 py-2.5 text-sm font-medium text-white shadow-md transition-all duration-200 hover:shadow-lg disabled:cursor-not-allowed disabled:shadow-none',
            loading ? variantConfig.confirmBgDisabled : variantConfig.confirmBg,
          ]"
        >
          <span v-if="loading" class="flex items-center justify-center gap-2">
            <div
              class="h-4 w-4 animate-spin rounded-full border-2 border-white border-t-transparent"
            ></div>
            Processando...
          </span>
          <span v-else>{{ confirmText }}</span>
        </button>
      </div>
    </template>
  </Dialog>
</template>
