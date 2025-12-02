<script setup lang="ts">
import { computed } from "vue";
import type { PlatformAnalytics } from "../../types/analytics";

interface Props {
  platformData: PlatformAnalytics;
  loading: boolean;
}

const props = defineProps<Props>();

const scrollToOffers = () => {
  const element = document.getElementById("offer-demonstration");
  element?.scrollIntoView({ behavior: "smooth" });
};

const scrollToHowItWorks = () => {
  const element = document.getElementById("how-it-works");
  element?.scrollIntoView({ behavior: "smooth" });
};

const formatNumber = (num: number): string => {
  if (num >= 1000) {
    return (num / 1000).toFixed(1) + "k";
  }
  return num.toString();
};

const formatMoney = (num: number): string => {
  if (num >= 1000000) {
    return "R$ " + (num / 1000000).toFixed(1) + "M+";
  }
  if (num >= 1000) {
    return "R$ " + (num / 1000).toFixed(1) + "k+";
  }
  return "R$ " + num.toFixed(0) + "+";
};

const formatEnergy = (num: number): string => {
  if (num >= 1000) {
    return (num / 1000).toFixed(1) + "k+";
  }
  return num.toFixed(1) + "+";
};

const activeOffersText = computed(() => {
  return props.loading
    ? "..."
    : formatNumber(props.platformData.active_offers) + "+";
});

const successfulPurchasesText = computed(() => {
  return props.loading
    ? "..."
    : formatNumber(props.platformData.sucessful_purchases) + "+";
});

const moneyTransactedText = computed(() => {
  return props.loading
    ? "..."
    : formatMoney(props.platformData.money_transacted);
});

const energyTransactedText = computed(() => {
  return props.loading
    ? "..."
    : formatEnergy(props.platformData.energy_transacted);
});
</script>

<template>
  <main
    id="home"
    class="relative flex min-h-dvh items-center justify-center overflow-hidden bg-gradient-to-br from-emerald-950 via-emerald-900 to-emerald-800 px-4 pt-16 sm:px-6 sm:pt-20 md:px-8"
  >
    <div
      class="absolute inset-0 bg-[radial-gradient(circle_at_30%_50%,rgba(16,185,129,0.1),transparent_50%)]"
    ></div>
    <div
      class="absolute inset-0 bg-[radial-gradient(circle_at_70%_50%,rgba(5,150,105,0.1),transparent_50%)]"
    ></div>

    <div
      class="relative z-10 mx-auto w-full max-w-6xl px-2 text-center sm:px-0"
    >
      <div
        class="mb-6 inline-flex items-center gap-2 rounded-full border border-emerald-400/30 bg-emerald-900/30 px-3 py-1.5 backdrop-blur-sm sm:mb-8 sm:px-4 sm:py-2"
      >
        <div
          class="flex h-5 w-5 items-center justify-center rounded-full bg-emerald-500 sm:h-6 sm:w-6"
        >
          <i class="pi pi-bolt text-[10px] text-white sm:text-xs"></i>
        </div>
        <span class="text-xs font-medium text-emerald-100 sm:text-sm">
          Energia Renovável
        </span>
      </div>

      <h1
        class="mb-4 bg-gradient-to-r from-white via-emerald-50 to-white bg-clip-text px-2 text-4xl leading-tight font-black tracking-tight text-transparent sm:mb-6 sm:px-0 sm:text-5xl md:text-6xl lg:text-7xl"
      >
        Transforme Energia<br />
        em Novas Oportunidades
      </h1>

      <p
        class="mx-auto mb-8 max-w-2xl px-4 text-base leading-relaxed text-emerald-100/90 sm:mb-12 sm:px-0 sm:text-lg md:text-xl"
      >
        O Ecoply conecta quem gera energia limpa a quem precisa dela, de forma
        transparente e direta.
      </p>

      <div
        class="flex flex-col items-center justify-center gap-3 sm:flex-row sm:gap-4"
      >
        <button
          @click="scrollToOffers"
          class="group flex w-full items-center justify-center gap-2 rounded-full bg-white px-6 py-3 text-sm font-semibold text-emerald-900 shadow-lg shadow-emerald-500/20 transition-all duration-300 hover:-translate-y-1 hover:shadow-xl hover:shadow-emerald-500/30 active:scale-95 sm:w-auto sm:gap-3 sm:px-8 sm:py-4 sm:text-base"
        >
          Explorar Ofertas
          <i
            class="pi pi-arrow-right text-xs transition-transform duration-300 group-hover:translate-x-1 sm:text-sm"
          ></i>
        </button>
        <button
          @click="scrollToHowItWorks"
          class="group w-full rounded-full border-2 border-white/30 bg-white/10 px-6 py-3 text-sm font-semibold text-white backdrop-blur-sm transition-all duration-300 hover:-translate-y-1 hover:border-white/50 hover:bg-white/20 active:scale-95 sm:w-auto sm:px-8 sm:py-4 sm:text-base"
        >
          Saiba Mais
        </button>
      </div>

      <div
        class="mt-12 grid grid-cols-2 gap-4 sm:mt-16 sm:gap-6 md:grid-cols-4 md:gap-8 lg:mt-20"
      >
        <div class="group cursor-default">
          <div
            class="text-2xl font-bold text-white transition-transform duration-300 group-hover:scale-110 sm:text-3xl md:text-4xl"
          >
            {{ activeOffersText }}
          </div>
          <div class="mt-1 text-xs text-emerald-200/80 sm:mt-2 sm:text-sm">
            Ofertas Ativas
          </div>
        </div>
        <div class="group cursor-default">
          <div
            class="text-2xl font-bold text-white transition-transform duration-300 group-hover:scale-110 sm:text-3xl md:text-4xl"
          >
            {{ energyTransactedText }}
          </div>
          <div class="mt-1 text-xs text-emerald-200/80 sm:mt-2 sm:text-sm">
            Energia Movimentada (MWh)
          </div>
        </div>
        <div class="group cursor-default">
          <div
            class="text-2xl font-bold text-white transition-transform duration-300 group-hover:scale-110 sm:text-3xl md:text-4xl"
          >
            {{ moneyTransactedText }}
          </div>
          <div class="mt-1 text-xs text-emerald-200/80 sm:mt-2 sm:text-sm">
            Volume Transacionado
          </div>
        </div>
        <div class="group cursor-default">
          <div
            class="text-2xl font-bold text-white transition-transform duration-300 group-hover:scale-110 sm:text-3xl md:text-4xl"
          >
            {{ successfulPurchasesText }}
          </div>
          <div class="mt-1 text-xs text-emerald-200/80 sm:mt-2 sm:text-sm">
            Transações Concluídas
          </div>
        </div>
      </div>
    </div>

    <div
      class="absolute bottom-6 left-1/2 -translate-x-1/2 animate-bounce sm:bottom-8"
    >
      <i class="pi pi-chevron-down text-xl text-white/50 sm:text-2xl"></i>
    </div>
  </main>
</template>

<style scoped>
@keyframes bounce {
  0%,
  100% {
    transform: translateY(0) translateX(-50%);
  }
  50% {
    transform: translateY(-10px) translateX(-50%);
  }
}

.animate-bounce {
  animation: bounce 2s infinite;
}
</style>
