<script setup lang="ts">
import { UserPlus, Search, Handshake } from "lucide-vue-next";
</script>

<template>
  <section
    id="how-it-works"
    class="bg-white px-4 py-12 sm:px-6 sm:py-16 md:px-8 lg:py-24"
  >
    <div class="mx-auto max-w-7xl">
      <div class="mb-12 text-center sm:mb-16 lg:mb-20">
        <div
          class="mb-3 inline-flex items-center gap-2 rounded-full bg-emerald-50 px-3 py-1.5 sm:mb-4 sm:px-4 sm:py-2"
        >
          <div
            class="flex h-5 w-5 items-center justify-center rounded-full bg-emerald-500 sm:h-6 sm:w-6"
          >
            <i class="pi pi-bolt text-[10px] text-white sm:text-xs"></i>
          </div>
          <span class="text-xs font-medium text-emerald-700 sm:text-sm">
            Simples e Eficiente
          </span>
        </div>
        <h2
          class="mb-3 px-2 text-3xl font-bold tracking-tight text-gray-900 sm:mb-4 sm:px-0 sm:text-4xl md:text-5xl"
        >
          Como funciona o Ecoply?
        </h2>
        <p
          class="mx-auto max-w-2xl px-4 text-base text-gray-600 sm:px-0 sm:text-lg"
        >
          Em apenas três passos simples, você faz parte da revolução da energia
          descentralizada
        </p>
      </div>

      <div class="relative grid gap-8 sm:gap-10 md:grid-cols-3 md:gap-12">
        <div
          class="absolute top-12 right-0 left-0 hidden h-0.5 bg-gradient-to-r from-emerald-200 via-emerald-300 to-emerald-200 md:block"
        ></div>

        <div class="group relative px-2 text-center sm:px-0">
          <div class="relative z-10 mx-auto mb-6 inline-flex sm:mb-8">
            <div
              class="flex h-20 w-20 items-center justify-center rounded-xl bg-gradient-to-br from-emerald-400 to-emerald-600 shadow-lg shadow-emerald-500/30 transition-all duration-300 group-hover:-translate-y-2 group-hover:shadow-xl group-hover:shadow-emerald-500/40 sm:h-24 sm:w-24 sm:rounded-2xl"
            >
              <UserPlus
                :size="32"
                class="text-white sm:hidden"
                :stroke-width="2"
              />
              <UserPlus
                :size="40"
                class="hidden text-white sm:block"
                :stroke-width="2"
              />
            </div>
            <div
              class="absolute -top-1.5 -right-1.5 flex h-7 w-7 items-center justify-center rounded-full bg-gray-900 text-xs font-bold text-white sm:-top-2 sm:-right-2 sm:h-8 sm:w-8 sm:text-sm"
            >
              1
            </div>
          </div>
          <h3 class="mb-2 text-lg font-bold text-gray-900 sm:mb-3 sm:text-xl">
            Cadastre-se
          </h3>
          <p class="text-sm leading-relaxed text-gray-600">
            Registre-se em nossa plataforma e defina seu perfil: você quer
            vender energia excedente ou comprar energia limpa local.
          </p>
        </div>

        <div class="group relative px-2 text-center sm:px-0">
          <div class="relative z-10 mx-auto mb-6 inline-flex sm:mb-8">
            <div
              class="flex h-20 w-20 items-center justify-center rounded-xl bg-gradient-to-br from-emerald-400 to-emerald-600 shadow-lg shadow-emerald-500/30 transition-all duration-300 group-hover:-translate-y-2 group-hover:shadow-xl group-hover:shadow-emerald-500/40 sm:h-24 sm:w-24 sm:rounded-2xl"
            >
              <Search
                :size="32"
                class="text-white sm:hidden"
                :stroke-width="2"
              />
              <Search
                :size="40"
                class="hidden text-white sm:block"
                :stroke-width="2"
              />
            </div>
            <div
              class="absolute -top-1.5 -right-1.5 flex h-7 w-7 items-center justify-center rounded-full bg-gray-900 text-xs font-bold text-white sm:-top-2 sm:-right-2 sm:h-8 sm:w-8 sm:text-sm"
            >
              2
            </div>
          </div>
          <h3 class="mb-2 text-lg font-bold text-gray-900 sm:mb-3 sm:text-xl">
            Publique ou Encontre Ofertas
          </h3>
          <p class="text-sm leading-relaxed text-gray-600">
            Produtores publicam sua energia disponível. Compradores navegam
            pelas ofertas locais com preços transparentes.
          </p>
        </div>

        <div class="group relative px-2 text-center sm:px-0">
          <div class="relative z-10 mx-auto mb-6 inline-flex sm:mb-8">
            <div
              class="flex h-20 w-20 items-center justify-center rounded-xl bg-gradient-to-br from-emerald-400 to-emerald-600 shadow-lg shadow-emerald-500/30 transition-all duration-300 group-hover:-translate-y-2 group-hover:shadow-xl group-hover:shadow-emerald-500/40 sm:h-24 sm:w-24 sm:rounded-2xl"
            >
              <Handshake
                :size="32"
                class="text-white sm:hidden"
                :stroke-width="2"
              />
              <Handshake
                :size="40"
                class="hidden text-white sm:block"
                :stroke-width="2"
              />
            </div>
            <div
              class="absolute -top-1.5 -right-1.5 flex h-7 w-7 items-center justify-center rounded-full bg-gray-900 text-xs font-bold text-white sm:-top-2 sm:-right-2 sm:h-8 sm:w-8 sm:text-sm"
            >
              3
            </div>
          </div>
          <h3 class="mb-2 text-lg font-bold text-gray-900 sm:mb-3 sm:text-xl">
            Negocie
          </h3>
          <p class="text-sm leading-relaxed text-gray-600">
            Complete transações seguras e acompanhe o impacto ambiental e
            financeiro em tempo real.
          </p>
        </div>
      </div>
    </div>
  </section>
</template>

<style scoped></style>
