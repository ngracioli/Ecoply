<script setup lang="ts"></script>

<template>
  <main
    id="home"
    class="bg-primary-bg-color flex h-screen flex-col items-center justify-center text-center"
  >
    <h3
      class="text-secondary-color mb-2 flex items-center gap-2 text-2xl font-bold"
    >
      Energia Renovável
      <i class="pi pi-check-circle"></i>
    </h3>
    <h1 class="text-secondary-color mb-6 text-8xl font-black">
      TRANSFORME ENERGIA<br />
      EM NOVAS OPORTUNIDADES
    </h1>
    <h2 class="text-secondary-color mb-10 text-3xl font-normal">
      O Ecoply conecta quem gera energia limpa a quem<br />precisa dela, de
      forma transparente e direta.
    </h2>
    <section class="flex flex-row gap-8 text-2xl font-semibold">
      <button
        class="hero-btn-primary group hover:bg-accent-color bg-secondary-color text-primary-color flex cursor-pointer items-center gap-4 rounded-xl px-8 py-4 transition-all duration-400 ease-out hover:-translate-y-1 hover:scale-105 hover:text-white hover:shadow-xl"
      >
        Explorar as Ofertas
        <i
          class="pi pi-arrow-right text-primary-color transition-transform duration-300 group-hover:translate-x-1 group-hover:text-white"
        ></i>
      </button>
      <button
        class="hero-btn-secondary group hover:bg-accent-color cursor-pointer rounded-xl border-2 border-lime-200 bg-transparent px-4 py-2 text-lime-200 transition-all duration-400 ease-out hover:-translate-y-1 hover:scale-105 hover:border-transparent hover:text-white hover:shadow-xl"
      >
        Saiba Mais
      </button>
    </section>
  </main>
</template>

<style scoped></style>
