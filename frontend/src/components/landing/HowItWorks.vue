<script setup lang="ts">
import { UserPlus, Search, Handshake } from "lucide-vue-next";
</script>

<template>
  <section id="how-it-works" class="bg-white px-8 py-24">
    <div class="mx-auto max-w-7xl">
      <div class="mb-20 text-center">
        <div
          class="mb-4 inline-flex items-center gap-2 rounded-full bg-emerald-50 px-4 py-2"
        >
          <div
            class="flex h-6 w-6 items-center justify-center rounded-full bg-emerald-500"
          >
            <i class="pi pi-bolt text-xs text-white"></i>
          </div>
          <span class="text-sm font-medium text-emerald-700">
            Simples e Eficiente
          </span>
        </div>
        <h2 class="mb-4 text-5xl font-bold tracking-tight text-gray-900">
          Como funciona o Ecoply?
        </h2>
        <p class="mx-auto max-w-2xl text-lg text-gray-600">
          Em apenas três passos simples, você faz parte da revolução da energia
          descentralizada
        </p>
      </div>

      <div class="relative grid gap-12 md:grid-cols-3">
        <div
          class="absolute top-12 right-0 left-0 hidden h-0.5 bg-gradient-to-r from-emerald-200 via-emerald-300 to-emerald-200 md:block"
        ></div>

        <div class="group relative text-center">
          <div class="relative z-10 mx-auto mb-8 inline-flex">
            <div
              class="flex h-24 w-24 items-center justify-center rounded-2xl bg-gradient-to-br from-emerald-400 to-emerald-600 shadow-lg shadow-emerald-500/30 transition-all duration-300 group-hover:-translate-y-2 group-hover:shadow-xl group-hover:shadow-emerald-500/40"
            >
              <UserPlus :size="40" class="text-white" :stroke-width="2" />
            </div>
            <div
              class="absolute -top-2 -right-2 flex h-8 w-8 items-center justify-center rounded-full bg-gray-900 text-sm font-bold text-white"
            >
              1
            </div>
          </div>
          <h3 class="mb-3 text-xl font-bold text-gray-900">Cadastre-se</h3>
          <p class="text-sm leading-relaxed text-gray-600">
            Registre-se em nossa plataforma e defina seu perfil: você quer
            vender energia excedente ou comprar energia limpa local.
          </p>
        </div>

        <div class="group relative text-center">
          <div class="relative z-10 mx-auto mb-8 inline-flex">
            <div
              class="flex h-24 w-24 items-center justify-center rounded-2xl bg-gradient-to-br from-emerald-400 to-emerald-600 shadow-lg shadow-emerald-500/30 transition-all duration-300 group-hover:-translate-y-2 group-hover:shadow-xl group-hover:shadow-emerald-500/40"
            >
              <Search :size="40" class="text-white" :stroke-width="2" />
            </div>
            <div
              class="absolute -top-2 -right-2 flex h-8 w-8 items-center justify-center rounded-full bg-gray-900 text-sm font-bold text-white"
            >
              2
            </div>
          </div>
          <h3 class="mb-3 text-xl font-bold text-gray-900">
            Publique ou Encontre Ofertas
          </h3>
          <p class="text-sm leading-relaxed text-gray-600">
            Produtores publicam sua energia disponível. Compradores navegam
            pelas ofertas locais com preços transparentes.
          </p>
        </div>

        <div class="group relative text-center">
          <div class="relative z-10 mx-auto mb-8 inline-flex">
            <div
              class="flex h-24 w-24 items-center justify-center rounded-2xl bg-gradient-to-br from-emerald-400 to-emerald-600 shadow-lg shadow-emerald-500/30 transition-all duration-300 group-hover:-translate-y-2 group-hover:shadow-xl group-hover:shadow-emerald-500/40"
            >
              <Handshake :size="40" class="text-white" :stroke-width="2" />
            </div>
            <div
              class="absolute -top-2 -right-2 flex h-8 w-8 items-center justify-center rounded-full bg-gray-900 text-sm font-bold text-white"
            >
              3
            </div>
          </div>
          <h3 class="mb-3 text-xl font-bold text-gray-900">Negocie</h3>
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
