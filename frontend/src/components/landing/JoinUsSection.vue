<script setup lang="ts">
import { useRouter } from "vue-router";
import { Zap } from "lucide-vue-next";

const router = useRouter();

const goToDashboard = () => {
  router.push({ name: "Dashboard" });
};
</script>

<template>
  <section
    class="relative overflow-hidden bg-gradient-to-br from-emerald-900 via-emerald-800 to-emerald-900 px-4 py-12 sm:px-6 sm:py-16 md:px-8 lg:py-24"
  >
    <div
      class="absolute inset-0 bg-[radial-gradient(circle_at_30%_50%,rgba(16,185,129,0.15),transparent_50%)]"
    ></div>
    <div
      class="absolute inset-0 bg-[radial-gradient(circle_at_70%_50%,rgba(5,150,105,0.15),transparent_50%)]"
    ></div>

    <div class="relative z-10 mx-auto max-w-5xl text-center">
      <div class="mb-6 flex justify-center sm:mb-8">
        <div
          class="flex h-16 w-16 items-center justify-center rounded-xl bg-gradient-to-br from-emerald-400 to-emerald-600 shadow-lg shadow-emerald-500/30 sm:h-20 sm:w-20 sm:rounded-2xl"
        >
          <Zap :size="32" class="text-white sm:hidden" :stroke-width="2.5" />
          <Zap
            :size="40"
            class="hidden text-white sm:block"
            :stroke-width="2.5"
          />
        </div>
      </div>

      <h2
        class="mb-4 text-3xl leading-tight font-bold tracking-tight text-white sm:mb-6 sm:text-4xl md:text-5xl"
      >
        Participe do Futuro da<br class="hidden sm:block" /><span
          class="sm:hidden"
        >
        </span
        >Energia Descentralizada
      </h2>

      <p
        class="mx-auto mb-8 max-w-2xl px-2 text-base text-emerald-100/90 sm:mb-12 sm:px-0 sm:text-lg"
      >
        Junte-se à revolução da energia renovável. Seja parte de uma comunidade
        que valoriza sustentabilidade, transparência e prosperidade local.
      </p>

      <button
        @click="goToDashboard"
        class="group inline-flex items-center gap-2 rounded-full bg-white px-6 py-3 text-sm font-semibold text-emerald-900 shadow-xl shadow-emerald-500/30 transition-all duration-300 hover:-translate-y-1 hover:shadow-2xl hover:shadow-emerald-500/40 active:scale-95 sm:gap-3 sm:px-10 sm:py-5 sm:text-lg"
      >
        Começar Agora
        <i
          class="pi pi-arrow-right text-xs transition-transform duration-300 group-hover:translate-x-1 sm:text-base"
        ></i>
      </button>
    </div>
  </section>
</template>

<style scoped></style>
