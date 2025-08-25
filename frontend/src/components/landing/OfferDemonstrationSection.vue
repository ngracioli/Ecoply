<script setup lang="ts">
import OfferCard from "../shared/OfferCard.vue";
</script>

<template>
  <main class="bg-primary-gradient py-24">
    <section class="mb-12 flex flex-col items-center justify-center">
      <h1 class="mb-6 text-6xl font-extrabold text-neutral-700">
        Marketplace de Energia
      </h1>
      <h2 class="mb-10 text-2xl font-bold text-neutral-600">
        Uma plataforma que gera valor para produtores, compradores e toda a
        comunidade
      </h2>
    </section>
    <section class="flex flex-row justify-center gap-32">
      <OfferCard />
      <OfferCard />
      <OfferCard />
    </section>
  </main>
</template>

<style scoped></style>
