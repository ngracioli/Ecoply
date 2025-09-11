<script setup lang="ts">
import BenefitsCard from "./benefits/BenefitsCard.vue";
</script>

<template>
  <main class="bg-primary-gradient py-16" id="benefits">
    <section class="flex flex-col items-center justify-center">
      <h1 class="mb-6 text-6xl font-extrabold text-neutral-900">
        Benefícios para Todos
      </h1>
      <h2 class="mb-10 text-2xl font-bold text-neutral-700">
        Uma plataforma que gera valor para produtores, compradores e toda a
        comunidade
      </h2>
    </section>
    <section class="container mx-auto px-6 py-12">
      <div class="flex items-center justify-center gap-32">
        <BenefitsCard
          icon="pi-user"
          title="Produtores"
          subtitle="Conecte-se com compradores locais e venda sua energia excedente de forma fácil e segura."
        />

        <div class="flex items-center gap-4">
          <div
            class="via-primary-color h-px w-12 bg-gradient-to-r from-transparent to-transparent"
          ></div>
          <div class="bg-primary-color h-3 w-3 rounded-full"></div>
          <div
            class="via-primary-color h-px w-12 bg-gradient-to-r from-transparent to-transparent"
          ></div>
        </div>

        <BenefitsCard
          icon="pi-user"
          title="Compradores"
          subtitle="Encontre ofertas de energia limpa e sustentável, com preços justos e transparência total."
        />
      </div>
    </section>
  </main>
</template>
