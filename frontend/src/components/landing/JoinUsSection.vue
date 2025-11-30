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
    class="relative overflow-hidden bg-gradient-to-br from-emerald-900 via-emerald-800 to-emerald-900 px-8 py-24"
  >
    <div
      class="absolute inset-0 bg-[radial-gradient(circle_at_30%_50%,rgba(16,185,129,0.15),transparent_50%)]"
    ></div>
    <div
      class="absolute inset-0 bg-[radial-gradient(circle_at_70%_50%,rgba(5,150,105,0.15),transparent_50%)]"
    ></div>

    <div class="relative z-10 mx-auto max-w-5xl text-center">
      <div class="mb-8 flex justify-center">
        <div
          class="flex h-20 w-20 items-center justify-center rounded-2xl bg-gradient-to-br from-emerald-400 to-emerald-600 shadow-lg shadow-emerald-500/30"
        >
          <Zap :size="40" class="text-white" :stroke-width="2.5" />
        </div>
      </div>

      <h2
        class="mb-6 text-5xl leading-tight font-bold tracking-tight text-white"
      >
        Participe do Futuro da<br />Energia Descentralizada
      </h2>

      <p class="mx-auto mb-12 max-w-2xl text-lg text-emerald-100/90">
        Junte-se à revolução da energia renovável. Seja parte de uma comunidade
        que valoriza sustentabilidade, transparência e prosperidade local.
      </p>

      <button
        @click="goToDashboard"
        class="group inline-flex items-center gap-3 rounded-full bg-white px-10 py-5 text-lg font-semibold text-emerald-900 shadow-xl shadow-emerald-500/30 transition-all duration-300 hover:-translate-y-1 hover:shadow-2xl hover:shadow-emerald-500/40"
      >
        Começar Agora
        <i
          class="pi pi-arrow-right text-base transition-transform duration-300 group-hover:translate-x-1"
        ></i>
      </button>
    </div>
  </section>
</template>

<style scoped></style>
