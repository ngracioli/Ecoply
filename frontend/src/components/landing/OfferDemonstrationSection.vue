<script setup lang="ts">
import OfferCard from "../shared/OfferCard.vue";
</script>

<template>
  <main class="bg-primary-gradient mb-10 py-24" id="offer-demonstration">
    <section class="mb-12 flex flex-col items-center justify-center">
      <h1 class="mb-6 text-6xl font-extrabold text-neutral-900">
        Marketplace de Energia
      </h1>
      <h2 class="mb-10 text-2xl font-bold text-neutral-700">
        Uma plataforma que gera valor para produtores, compradores e toda a
        comunidade
      </h2>
    </section>
    <section class="flex flex-row items-center justify-center px-8">
      <div
        class="transform transition-all duration-500 hover:-translate-y-2 hover:scale-105"
      >
        <div class="scale-90 opacity-85 hover:opacity-100">
          <OfferCard />
        </div>
      </div>

      <div
        class="relative z-10 mx-8 transform transition-all duration-500 hover:-translate-y-3 hover:scale-105"
      >
        <div class="relative">
          <div
            class="bg-accent-color/10 absolute inset-0 scale-110 rounded-3xl blur-2xl"
          ></div>

          <div
            class="shadow-accent-color/20 relative scale-110 rounded-2xl shadow-2xl"
          >
            <OfferCard />
          </div>
        </div>
      </div>

      <div
        class="transform transition-all duration-500 hover:-translate-y-2 hover:scale-105"
      >
        <div class="scale-90 opacity-85 hover:opacity-100">
          <OfferCard />
        </div>
      </div>
    </section>
  </main>
</template>

<style scoped></style>
