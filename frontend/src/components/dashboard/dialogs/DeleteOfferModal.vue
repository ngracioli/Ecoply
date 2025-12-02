<script setup lang="ts">
interface Props {
  visible: boolean;
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
      v-if="visible"
      class="fixed inset-0 z-50 flex items-center justify-center bg-black/60 p-4 backdrop-blur-sm sm:p-0"
      @click.self="handleClose"
    >
      <div
        class="relative mx-4 w-full max-w-md overflow-hidden rounded-2xl bg-white shadow-2xl sm:rounded-3xl"
      >
        <div
          class="flex flex-col items-center justify-center px-6 py-8 sm:p-12"
        >
          <div class="relative mb-5 sm:mb-6">
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
            class="mb-2 px-2 text-center text-xl font-bold text-red-600 sm:px-0 sm:text-2xl"
          >
            Não foi possível excluir a Oferta
          </h3>
          <p
            class="px-2 text-center text-sm text-neutral-600 sm:px-0 sm:text-base"
          >
            Você só pode excluir ofertas que ainda não tiveram nenhuma interação
          </p>
          <button
            @click="handleClose"
            class="mt-5 min-h-[44px] w-full touch-manipulation rounded-lg bg-neutral-100 px-6 py-2.5 text-sm font-medium text-neutral-700 transition-colors hover:bg-neutral-200 active:scale-[0.98] sm:mt-6 sm:w-auto"
          >
            Entendi
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

.animate-x-draw {
  stroke-dasharray: 100;
  animation: x-draw 0.6s ease forwards;
  animation-delay: 0.2s;
}
</style>
