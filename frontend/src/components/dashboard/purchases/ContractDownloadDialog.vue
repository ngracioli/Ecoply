<script setup lang="ts">
import { computed } from "vue";
import ConfirmDialog from "../../shared/ConfirmDialog.vue";

interface Props {
  visible: boolean;
}

interface Emits {
  (e: "update:visible", value: boolean): void;
  (e: "download"): void;
  (e: "skip"): void;
}

const props = defineProps<Props>();
const emit = defineEmits<Emits>();

const isVisible = computed({
  get: () => props.visible,
  set: (value) => emit("update:visible", value),
});

const handleDownload = () => {
  emit("download");
};

const handleSkip = () => {
  emit("skip");
};
</script>

<template>
  <ConfirmDialog
    v-model:visible="isVisible"
    title="Baixar Contrato de Compra"
    message="Sua compra foi realizada com sucesso! Deseja baixar o contrato da compra com todas as informações da transação?"
    confirm-text="Baixar Contrato"
    cancel-text="Agora Não"
    variant="info"
    @confirm="handleDownload"
    @cancel="handleSkip"
  />
</template>
