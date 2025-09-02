<script setup lang="ts">
import IconCircle from "../shared/IconCircle.vue";
</script>

<template>
  <main class="bg-secondary-bg-color py-16">
    <section class="flex flex-col items-center justify-center">
      <div class="flex w-full max-w-2/3 flex-col items-center justify-center">
        <IconCircle
          :icon="'pi-user'"
          class="mb-4 transition-transform duration-300 group-hover:scale-110"
        />
        <h1 class="mb-8 text-center text-6xl font-extrabold text-neutral-900">
          Participe do futuro da energia descentralizada
        </h1>
        <h2 class="mb-10 text-center text-2xl font-bold text-neutral-700">
          Junte-se à revolução da energia renovável. Seja parte de uma
          comunidade que valoriza sustentabilidade, transparência e prosperidade
          local.
        </h2>
        <button
          class="bg-primary-color mb-16 flex flex-row items-center gap-4 rounded-2xl px-8 py-4 text-xl text-white transition-colors duration-300 hover:bg-neutral-600"
        >
          Testar Agora
          <i class="pi pi-arrow-right text-sm"></i>
        </button>
        <article>
          <ul
            class="text-md flex w-full flex-row items-center justify-between gap-32 whitespace-nowrap text-neutral-700"
          >
            <li class="flex min-w-0 flex-1 flex-col items-center text-center">
              <span class="text-primary-color text-3xl font-bold">234</span>
              Ofertas Abertas
            </li>
            <li class="flex min-w-0 flex-1 flex-col items-center text-center">
              <span class="text-primary-color text-3xl font-bold">301</span>
              Transações Concluídas
            </li>
            <li class="flex min-w-0 flex-1 flex-col items-center text-center">
              <span class="text-primary-color text-3xl font-bold">24/7</span>
              Suporte ao Cliente
            </li>
          </ul>
        </article>
      </div>
    </section>
  </main>
</template>

<style scoped></style>
