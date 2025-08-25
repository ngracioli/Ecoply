<script setup lang="ts">
import BenefitsCard from "./benefits/BenefitsCard.vue";
</script>

<template>
  <main class="bg-primary-gradient py-16">
    <section class="flex flex-col items-center justify-center">
      <h1 class="mb-6 text-6xl font-extrabold text-neutral-700">
        Benefícios para Todos
      </h1>
      <h2 class="mb-10 text-2xl font-bold text-neutral-600">
        Uma plataforma que gera valor para produtores, compradores e toda a
        comunidade
      </h2>
    </section>
    <section class="flex flex-row justify-evenly gap-32 py-12 text-4xl">
      <BenefitsCard
        icon="pi-user"
        title="Produtores"
        subtitle="Conecte-se com compradores locais e venda sua energia excedente de forma fácil e segura."
      />
      <BenefitsCard
        icon="pi-user"
        title="Compradores"
        subtitle="Encontre ofertas de energia limpa e sustentável, com preços justos e transparência total."
      />
    </section>
  </main>
</template>
