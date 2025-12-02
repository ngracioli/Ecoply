<script setup lang="ts">
import { Zap } from "lucide-vue-next";

interface Props {
  status: "idle" | "processing" | "success" | "error";
  errorMessage?: string;
}

interface Emits {
  (e: "close"): void;
}

defineProps<Props>();
const emit = defineEmits<Emits>();

const handleClose = () => {
  emit("close");
};
</script>

<template>
  <Transition name="modal-fade">
    <div
      v-if="status !== 'idle'"
      class="fixed inset-0 z-50 flex items-center justify-center bg-black/60 px-4 py-6 backdrop-blur-sm sm:px-6"
      @click.self="status === 'error' ? handleClose() : null"
    >
      <div
        class="relative w-full max-w-md overflow-hidden rounded-2xl bg-white shadow-2xl sm:rounded-3xl"
      >
        <div
          v-if="status === 'processing'"
          class="flex flex-col items-center justify-center p-6 sm:p-10 md:p-12"
        >
          <div class="relative mb-4 sm:mb-6">
            <div
              class="absolute inset-0 animate-spin rounded-full border-4 border-emerald-200"
              style="animation-duration: 1.5s"
            ></div>
            <div
              class="flex h-20 w-20 animate-pulse items-center justify-center rounded-full bg-gradient-to-br from-emerald-400 to-green-500 sm:h-24 sm:w-24"
            >
              <Zap :size="36" class="text-white sm:hidden" />
              <Zap :size="40" class="hidden text-white sm:block" />
            </div>
          </div>
          <h3
            class="mb-2 px-2 text-center text-xl font-bold text-neutral-900 sm:text-2xl"
          >
            Processando Pagamento
          </h3>
          <p class="px-4 text-center text-sm text-neutral-600 sm:text-base">
            Aguarde enquanto confirmamos sua compra...
          </p>
          <div class="mt-4 flex gap-2 sm:mt-6">
            <div
              class="h-2 w-2 animate-bounce rounded-full bg-emerald-500"
              style="animation-delay: 0ms"
            ></div>
            <div
              class="h-2 w-2 animate-bounce rounded-full bg-emerald-500"
              style="animation-delay: 150ms"
            ></div>
            <div
              class="h-2 w-2 animate-bounce rounded-full bg-emerald-500"
              style="animation-delay: 300ms"
            ></div>
          </div>
        </div>

        <div
          v-else-if="status === 'success'"
          class="flex flex-col items-center justify-center p-6 sm:p-10 md:p-12"
        >
          <div class="relative mb-4 sm:mb-6">
            <div
              class="animate-scale-in flex h-20 w-20 items-center justify-center rounded-full bg-gradient-to-br from-emerald-400 to-green-500 sm:h-24 sm:w-24"
            >
              <svg
                class="animate-check-draw h-10 w-10 text-white sm:h-12 sm:w-12"
                viewBox="0 0 24 24"
                fill="none"
                stroke="currentColor"
                stroke-width="3"
                stroke-linecap="round"
                stroke-linejoin="round"
              >
                <polyline points="20 6 9 17 4 12" />
              </svg>
            </div>
            <div
              class="absolute -inset-2 animate-ping rounded-full border-4 border-emerald-200"
              style="animation-iteration-count: 1; animation-duration: 1s"
            ></div>
          </div>
          <h3
            class="mb-2 px-2 text-center text-xl font-bold text-emerald-600 sm:text-2xl"
          >
            Compra Concluída!
          </h3>
          <p class="px-4 text-center text-sm text-neutral-600 sm:text-base">
            Sua compra foi realizada com sucesso
          </p>
        </div>

        <div
          v-else-if="status === 'error'"
          class="flex flex-col items-center justify-center p-6 sm:p-10 md:p-12"
        >
          <div class="relative mb-4 sm:mb-6">
            <div
              class="animate-scale-in flex h-20 w-20 items-center justify-center rounded-full bg-gradient-to-br from-red-400 to-red-500 sm:h-24 sm:w-24"
            >
              <svg
                class="animate-x-draw h-10 w-10 text-white sm:h-12 sm:w-12"
                viewBox="0 0 24 24"
                fill="none"
                stroke="currentColor"
                stroke-width="3"
                stroke-linecap="round"
                stroke-linejoin="round"
              >
                <line x1="18" y1="6" x2="6" y2="18" />
                <line x1="6" y1="6" x2="18" y2="18" />
              </svg>
            </div>
            <div
              class="absolute -inset-2 animate-ping rounded-full border-4 border-red-200"
              style="animation-iteration-count: 1; animation-duration: 1s"
            ></div>
          </div>
          <h3
            class="mb-2 px-2 text-center text-xl font-bold text-red-600 sm:text-2xl"
          >
            Pagamento Recusado
          </h3>
          <p class="px-4 text-center text-sm text-neutral-600 sm:text-base">
            {{ errorMessage || "Não foi possível processar sua compra" }}
          </p>
          <button
            @click="handleClose"
            class="mt-4 min-h-[44px] rounded-lg bg-neutral-100 px-6 py-3 text-sm font-medium text-neutral-700 transition-colors hover:bg-neutral-200 active:bg-neutral-300 sm:mt-6 sm:min-h-0 sm:py-2.5"
          >
            Tentar Novamente
          </button>
        </div>
      </div>
    </div>
  </Transition>
</template>

<style scoped>
.modal-fade-enter-active,
.modal-fade-leave-active {
  transition: opacity 0.3s ease;
}

.modal-fade-enter-active > div,
.modal-fade-leave-active > div {
  transition: transform 0.3s ease;
}

.modal-fade-enter-from,
.modal-fade-leave-to {
  opacity: 0;
}

.modal-fade-enter-from > div {
  transform: scale(0.9) translateY(-20px);
}

.modal-fade-leave-to > div {
  transform: scale(0.9) translateY(20px);
}

@keyframes scale-in {
  0% {
    transform: scale(0);
    opacity: 0;
  }
  50% {
    transform: scale(1.1);
  }
  100% {
    transform: scale(1);
    opacity: 1;
  }
}

@keyframes check-draw {
  0% {
    stroke-dasharray: 0, 100;
    opacity: 0;
  }
  50% {
    opacity: 1;
  }
  100% {
    stroke-dasharray: 100, 0;
    opacity: 1;
  }
}

@keyframes x-draw {
  0% {
    stroke-dasharray: 0, 100;
    opacity: 0;
  }
  50% {
    opacity: 1;
  }
  100% {
    stroke-dasharray: 100, 0;
    opacity: 1;
  }
}

.animate-scale-in {
  animation: scale-in 0.5s cubic-bezier(0.34, 1.56, 0.64, 1);
}

.animate-check-draw {
  stroke-dasharray: 100;
  animation: check-draw 0.6s ease forwards;
  animation-delay: 0.2s;
}

.animate-x-draw {
  stroke-dasharray: 100;
  animation: x-draw 0.6s ease forwards;
  animation-delay: 0.2s;
}
</style>
