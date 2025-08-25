<script setup lang="ts">
import NumericCard from "../shared/NumericCard.vue";
</script>

<template>
  <main class="bg-primary-gradient py-16">
    <section class="flex flex-col items-center justify-center py-12">
      <h1 class="mb-6 text-6xl font-extrabold text-neutral-700">
        Como funciona o Ecoply?
      </h1>
      <h2 class="mb-10 text-2xl font-bold text-neutral-600">
        Em apenas três passos simples, você faz parte da revolução da energia
        descentralizada
      </h2>
    </section>
    <section class="flex flex-row justify-center gap-32">
      <NumericCard
        :step="1"
        icon="pi-user"
        title="Cadastre-se"
        subtitle="Registre-se em nossa plataforma e defina seu perfil: você quer vender energia excedente ou comprar energia limpa local."
      />
      <NumericCard
        :step="2"
        icon="pi-cog"
        title="Publique ou Encontre Ofertas"
        subtitle="Produtores publicam sua energia disponível. Compradores navegam pelas ofertas locais com preços transparentes."
      />
      <NumericCard
        :step="3"
        icon="pi-check"
        title="Negocie"
        subtitle="Complete transações seguras e acompanhe o impacto ambiental e financeiro em tempo real."
      />
    </section>
  </main>
</template>

<style scoped></style>
