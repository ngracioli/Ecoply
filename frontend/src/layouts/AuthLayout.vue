<script setup lang="ts"></script>

<template>
  <main class="grid min-h-screen grid-cols-1 lg:grid-cols-2">
    <section
      class="from-primary-color to-primary-dark-color flex flex-col justify-center bg-gradient-to-br p-8 text-white lg:p-16"
    >
      <div class="max-w-lg">
        <h1 class="mb-4 text-4xl font-bold lg:text-5xl">Bem-vindo ao Ecoply</h1>
        <p class="mb-8 text-lg opacity-90 lg:text-xl">
          O marketplace de energia renovável que conecta você a um futuro
          sustentável
        </p>
        <ul class="space-y-3 text-base lg:text-lg">
          <li class="flex items-center">
            <span class="mr-3">✓</span>
            Leilões transparentes de energia limpa
          </li>
          <li class="flex items-center">
            <span class="mr-3">✓</span>
            Certificações verificadas
          </li>
          <li class="flex items-center">
            <span class="mr-3">✓</span>
            Impacto ambiental positivo
          </li>
        </ul>
      </div>
    </section>

    <section class="flex items-center justify-center bg-gray-50 p-6 lg:p-12">
      <div class="w-full max-w-md">
        <slot />
      </div>
    </section>
  </main>
</template>
