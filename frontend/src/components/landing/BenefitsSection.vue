<script setup lang="ts">
import {
  Zap,
  ShoppingCart,
  TrendingUp,
  Shield,
  Users,
  Leaf,
} from "lucide-vue-next";
</script>

<template>
  <section
    id="benefits"
    class="bg-gray-50 px-4 py-12 sm:px-6 sm:py-16 lg:px-8 lg:py-24"
  >
    <div class="mx-auto max-w-7xl">
      <div class="mb-10 text-center sm:mb-14 lg:mb-20">
        <div
          class="mb-3 inline-flex items-center gap-2 rounded-full bg-emerald-50 px-3 py-1.5 sm:mb-4 sm:px-4 sm:py-2"
        >
          <div
            class="flex h-5 w-5 items-center justify-center rounded-full bg-emerald-500 sm:h-6 sm:w-6"
          >
            <i class="pi pi-heart text-[10px] text-white sm:text-xs"></i>
          </div>
          <span class="text-xs font-medium text-emerald-700 sm:text-sm">
            Vantagens Exclusivas
          </span>
        </div>
        <h2
          class="mb-3 text-3xl font-bold tracking-tight text-gray-900 sm:text-4xl lg:mb-4 lg:text-5xl"
        >
          Benefícios para Todos
        </h2>
        <p
          class="mx-auto max-w-2xl px-4 text-base text-gray-600 sm:px-0 sm:text-lg"
        >
          Uma plataforma que gera valor para produtores, compradores e toda a
          comunidade
        </p>
      </div>

      <div class="grid gap-5 sm:grid-cols-2 sm:gap-6 lg:grid-cols-3 lg:gap-8">
        <div
          class="group rounded-xl border border-gray-200 bg-white p-6 transition-all duration-300 hover:-translate-y-2 hover:border-emerald-200 hover:shadow-xl hover:shadow-emerald-500/10 sm:rounded-2xl sm:p-8"
        >
          <div
            class="mb-4 inline-flex h-12 w-12 items-center justify-center rounded-lg bg-gradient-to-br from-emerald-100 to-emerald-200 transition-all duration-300 group-hover:from-emerald-500 group-hover:to-emerald-600 sm:mb-6 sm:h-14 sm:w-14 sm:rounded-xl"
          >
            <Zap
              :size="24"
              class="text-emerald-600 transition-colors duration-300 group-hover:text-white sm:size-7"
              :stroke-width="2"
            />
          </div>
          <h3 class="mb-2 text-lg font-bold text-gray-900 sm:mb-3 sm:text-xl">
            Venda sua Energia
          </h3>
          <p class="text-sm leading-relaxed text-gray-600">
            Conecte-se com compradores locais e venda sua energia excedente de
            forma fácil e segura.
          </p>
        </div>

        <div
          class="group rounded-xl border border-gray-200 bg-white p-6 transition-all duration-300 hover:-translate-y-2 hover:border-emerald-200 hover:shadow-xl hover:shadow-emerald-500/10 sm:rounded-2xl sm:p-8"
        >
          <div
            class="mb-4 inline-flex h-12 w-12 items-center justify-center rounded-lg bg-gradient-to-br from-blue-100 to-blue-200 transition-all duration-300 group-hover:from-blue-500 group-hover:to-blue-600 sm:mb-6 sm:h-14 sm:w-14 sm:rounded-xl"
          >
            <ShoppingCart
              :size="24"
              class="text-blue-600 transition-colors duration-300 group-hover:text-white sm:size-7"
              :stroke-width="2"
            />
          </div>
          <h3 class="mb-2 text-lg font-bold text-gray-900 sm:mb-3 sm:text-xl">
            Compre Energia Limpa
          </h3>
          <p class="text-sm leading-relaxed text-gray-600">
            Encontre ofertas de energia limpa e sustentável, com preços justos e
            transparência total.
          </p>
        </div>

        <div
          class="group rounded-xl border border-gray-200 bg-white p-6 transition-all duration-300 hover:-translate-y-2 hover:border-emerald-200 hover:shadow-xl hover:shadow-emerald-500/10 sm:rounded-2xl sm:p-8"
        >
          <div
            class="mb-4 inline-flex h-12 w-12 items-center justify-center rounded-lg bg-gradient-to-br from-amber-100 to-amber-200 transition-all duration-300 group-hover:from-amber-500 group-hover:to-amber-600 sm:mb-6 sm:h-14 sm:w-14 sm:rounded-xl"
          >
            <TrendingUp
              :size="24"
              class="text-amber-600 transition-colors duration-300 group-hover:text-white sm:size-7"
              :stroke-width="2"
            />
          </div>
          <h3 class="mb-2 text-lg font-bold text-gray-900 sm:mb-3 sm:text-xl">
            Economia Garantida
          </h3>
          <p class="text-sm leading-relaxed text-gray-600">
            Reduza seus custos com energia através de transações diretas, sem
            intermediários desnecessários.
          </p>
        </div>

        <div
          class="group rounded-xl border border-gray-200 bg-white p-6 transition-all duration-300 hover:-translate-y-2 hover:border-emerald-200 hover:shadow-xl hover:shadow-emerald-500/10 sm:rounded-2xl sm:p-8"
        >
          <div
            class="mb-4 inline-flex h-12 w-12 items-center justify-center rounded-lg bg-gradient-to-br from-purple-100 to-purple-200 transition-all duration-300 group-hover:from-purple-500 group-hover:to-purple-600 sm:mb-6 sm:h-14 sm:w-14 sm:rounded-xl"
          >
            <Shield
              :size="24"
              class="text-purple-600 transition-colors duration-300 group-hover:text-white sm:size-7"
              :stroke-width="2"
            />
          </div>
          <h3 class="mb-2 text-lg font-bold text-gray-900 sm:mb-3 sm:text-xl">
            Transações Seguras
          </h3>
          <p class="text-sm leading-relaxed text-gray-600">
            Sistema robusto de verificação e proteção de dados para garantir a
            segurança de todas as negociações.
          </p>
        </div>

        <div
          class="group rounded-xl border border-gray-200 bg-white p-6 transition-all duration-300 hover:-translate-y-2 hover:border-emerald-200 hover:shadow-xl hover:shadow-emerald-500/10 sm:rounded-2xl sm:p-8"
        >
          <div
            class="mb-4 inline-flex h-12 w-12 items-center justify-center rounded-lg bg-gradient-to-br from-pink-100 to-pink-200 transition-all duration-300 group-hover:from-pink-500 group-hover:to-pink-600 sm:mb-6 sm:h-14 sm:w-14 sm:rounded-xl"
          >
            <Users
              :size="24"
              class="text-pink-600 transition-colors duration-300 group-hover:text-white sm:size-7"
              :stroke-width="2"
            />
          </div>
          <h3 class="mb-2 text-lg font-bold text-gray-900 sm:mb-3 sm:text-xl">
            Comunidade Ativa
          </h3>
          <p class="text-sm leading-relaxed text-gray-600">
            Faça parte de uma comunidade que valoriza sustentabilidade e
            prosperidade local.
          </p>
        </div>

        <div
          class="group rounded-xl border border-gray-200 bg-white p-6 transition-all duration-300 hover:-translate-y-2 hover:border-emerald-200 hover:shadow-xl hover:shadow-emerald-500/10 sm:rounded-2xl sm:p-8"
        >
          <div
            class="mb-4 inline-flex h-12 w-12 items-center justify-center rounded-lg bg-gradient-to-br from-green-100 to-green-200 transition-all duration-300 group-hover:from-green-500 group-hover:to-green-600 sm:mb-6 sm:h-14 sm:w-14 sm:rounded-xl"
          >
            <Leaf
              :size="24"
              class="text-green-600 transition-colors duration-300 group-hover:text-white sm:size-7"
              :stroke-width="2"
            />
          </div>
          <h3 class="mb-2 text-lg font-bold text-gray-900 sm:mb-3 sm:text-xl">
            Impacto Ambiental Positivo
          </h3>
          <p class="text-sm leading-relaxed text-gray-600">
            Acompanhe o impacto ambiental das suas transações e contribua para
            um futuro mais sustentável.
          </p>
        </div>
      </div>
    </div>
  </section>
</template>
