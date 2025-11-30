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
    class="relative flex min-h-screen items-center justify-center overflow-hidden bg-gradient-to-br from-emerald-950 via-emerald-900 to-emerald-800 px-8 pt-20"
  >
    <div
      class="absolute inset-0 bg-[radial-gradient(circle_at_30%_50%,rgba(16,185,129,0.1),transparent_50%)]"
    ></div>
    <div
      class="absolute inset-0 bg-[radial-gradient(circle_at_70%_50%,rgba(5,150,105,0.1),transparent_50%)]"
    ></div>

    <div class="relative z-10 mx-auto max-w-6xl text-center">
      <div
        class="mb-8 inline-flex items-center gap-2 rounded-full border border-emerald-400/30 bg-emerald-900/30 px-4 py-2 backdrop-blur-sm"
      >
        <div
          class="flex h-6 w-6 items-center justify-center rounded-full bg-emerald-500"
        >
          <i class="pi pi-bolt text-xs text-white"></i>
        </div>
        <span class="text-sm font-medium text-emerald-100">
          Energia Renovável
        </span>
      </div>

      <h1
        class="mb-6 bg-gradient-to-r from-white via-emerald-50 to-white bg-clip-text text-7xl leading-tight font-black tracking-tight text-transparent"
      >
        Transforme Energia<br />
        em Novas Oportunidades
      </h1>

      <p
        class="mx-auto mb-12 max-w-2xl text-xl leading-relaxed text-emerald-100/90"
      >
        O Ecoply conecta quem gera energia limpa a quem precisa dela, de forma
        transparente e direta.
      </p>

      <div class="flex items-center justify-center gap-4">
        <button
          @click="scrollToOffers"
          class="group flex items-center gap-3 rounded-full bg-white px-8 py-4 text-base font-semibold text-emerald-900 shadow-lg shadow-emerald-500/20 transition-all duration-300 hover:-translate-y-1 hover:shadow-xl hover:shadow-emerald-500/30"
        >
          Explorar Ofertas
          <i
            class="pi pi-arrow-right text-sm transition-transform duration-300 group-hover:translate-x-1"
          ></i>
        </button>
        <button
          @click="scrollToHowItWorks"
          class="group rounded-full border-2 border-white/30 bg-white/10 px-8 py-4 text-base font-semibold text-white backdrop-blur-sm transition-all duration-300 hover:-translate-y-1 hover:border-white/50 hover:bg-white/20"
        >
          Saiba Mais
        </button>
      </div>

      <div class="mt-20 grid grid-cols-4 gap-8">
        <div class="group cursor-default">
          <div
            class="text-4xl font-bold text-white transition-transform duration-300 group-hover:scale-110"
          >
            {{ activeOffersText }}
          </div>
          <div class="mt-2 text-sm text-emerald-200/80">Ofertas Ativas</div>
        </div>
        <div class="group cursor-default">
          <div
            class="text-4xl font-bold text-white transition-transform duration-300 group-hover:scale-110"
          >
            {{ energyTransactedText }}
          </div>
          <div class="mt-2 text-sm text-emerald-200/80">
            Energia Movimentada (MWh)
          </div>
        </div>
        <div class="group cursor-default">
          <div
            class="text-4xl font-bold text-white transition-transform duration-300 group-hover:scale-110"
          >
            {{ moneyTransactedText }}
          </div>
          <div class="mt-2 text-sm text-emerald-200/80">
            Volume Transacionado
          </div>
        </div>
        <div class="group cursor-default">
          <div
            class="text-4xl font-bold text-white transition-transform duration-300 group-hover:scale-110"
          >
            {{ successfulPurchasesText }}
          </div>
          <div class="mt-2 text-sm text-emerald-200/80">
            Transações Concluídas
          </div>
        </div>
      </div>
    </div>

    <div class="absolute bottom-8 left-1/2 -translate-x-1/2 animate-bounce">
      <i class="pi pi-chevron-down text-2xl text-white/50"></i>
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
