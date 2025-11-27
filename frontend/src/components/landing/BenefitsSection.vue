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
  <section id="benefits" class="bg-gray-50 px-8 py-24">
    <div class="mx-auto max-w-7xl">
      <div class="mb-20 text-center">
        <div
          class="mb-4 inline-flex items-center gap-2 rounded-full bg-emerald-50 px-4 py-2"
        >
          <div
            class="flex h-6 w-6 items-center justify-center rounded-full bg-emerald-500"
          >
            <i class="pi pi-heart text-xs text-white"></i>
          </div>
          <span class="text-sm font-medium text-emerald-700">
            Vantagens Exclusivas
          </span>
        </div>
        <h2 class="mb-4 text-5xl font-bold tracking-tight text-gray-900">
          Benefícios para Todos
        </h2>
        <p class="mx-auto max-w-2xl text-lg text-gray-600">
          Uma plataforma que gera valor para produtores, compradores e toda a
          comunidade
        </p>
      </div>

      <div class="grid gap-8 md:grid-cols-2 lg:grid-cols-3">
        <div
          class="group rounded-2xl border border-gray-200 bg-white p-8 transition-all duration-300 hover:-translate-y-2 hover:border-emerald-200 hover:shadow-xl hover:shadow-emerald-500/10"
        >
          <div
            class="mb-6 inline-flex h-14 w-14 items-center justify-center rounded-xl bg-gradient-to-br from-emerald-100 to-emerald-200 transition-all duration-300 group-hover:from-emerald-500 group-hover:to-emerald-600"
          >
            <Zap
              :size="28"
              class="text-emerald-600 transition-colors duration-300 group-hover:text-white"
              :stroke-width="2"
            />
          </div>
          <h3 class="mb-3 text-xl font-bold text-gray-900">
            Venda sua Energia
          </h3>
          <p class="text-sm leading-relaxed text-gray-600">
            Conecte-se com compradores locais e venda sua energia excedente de
            forma fácil e segura.
          </p>
        </div>

        <div
          class="group rounded-2xl border border-gray-200 bg-white p-8 transition-all duration-300 hover:-translate-y-2 hover:border-emerald-200 hover:shadow-xl hover:shadow-emerald-500/10"
        >
          <div
            class="mb-6 inline-flex h-14 w-14 items-center justify-center rounded-xl bg-gradient-to-br from-blue-100 to-blue-200 transition-all duration-300 group-hover:from-blue-500 group-hover:to-blue-600"
          >
            <ShoppingCart
              :size="28"
              class="text-blue-600 transition-colors duration-300 group-hover:text-white"
              :stroke-width="2"
            />
          </div>
          <h3 class="mb-3 text-xl font-bold text-gray-900">
            Compre Energia Limpa
          </h3>
          <p class="text-sm leading-relaxed text-gray-600">
            Encontre ofertas de energia limpa e sustentável, com preços justos e
            transparência total.
          </p>
        </div>

        <div
          class="group rounded-2xl border border-gray-200 bg-white p-8 transition-all duration-300 hover:-translate-y-2 hover:border-emerald-200 hover:shadow-xl hover:shadow-emerald-500/10"
        >
          <div
            class="mb-6 inline-flex h-14 w-14 items-center justify-center rounded-xl bg-gradient-to-br from-amber-100 to-amber-200 transition-all duration-300 group-hover:from-amber-500 group-hover:to-amber-600"
          >
            <TrendingUp
              :size="28"
              class="text-amber-600 transition-colors duration-300 group-hover:text-white"
              :stroke-width="2"
            />
          </div>
          <h3 class="mb-3 text-xl font-bold text-gray-900">
            Economia Garantida
          </h3>
          <p class="text-sm leading-relaxed text-gray-600">
            Reduza seus custos com energia através de transações diretas, sem
            intermediários desnecessários.
          </p>
        </div>

        <div
          class="group rounded-2xl border border-gray-200 bg-white p-8 transition-all duration-300 hover:-translate-y-2 hover:border-emerald-200 hover:shadow-xl hover:shadow-emerald-500/10"
        >
          <div
            class="mb-6 inline-flex h-14 w-14 items-center justify-center rounded-xl bg-gradient-to-br from-purple-100 to-purple-200 transition-all duration-300 group-hover:from-purple-500 group-hover:to-purple-600"
          >
            <Shield
              :size="28"
              class="text-purple-600 transition-colors duration-300 group-hover:text-white"
              :stroke-width="2"
            />
          </div>
          <h3 class="mb-3 text-xl font-bold text-gray-900">
            Transações Seguras
          </h3>
          <p class="text-sm leading-relaxed text-gray-600">
            Sistema robusto de verificação e proteção de dados para garantir a
            segurança de todas as negociações.
          </p>
        </div>

        <div
          class="group rounded-2xl border border-gray-200 bg-white p-8 transition-all duration-300 hover:-translate-y-2 hover:border-emerald-200 hover:shadow-xl hover:shadow-emerald-500/10"
        >
          <div
            class="mb-6 inline-flex h-14 w-14 items-center justify-center rounded-xl bg-gradient-to-br from-pink-100 to-pink-200 transition-all duration-300 group-hover:from-pink-500 group-hover:to-pink-600"
          >
            <Users
              :size="28"
              class="text-pink-600 transition-colors duration-300 group-hover:text-white"
              :stroke-width="2"
            />
          </div>
          <h3 class="mb-3 text-xl font-bold text-gray-900">Comunidade Ativa</h3>
          <p class="text-sm leading-relaxed text-gray-600">
            Faça parte de uma comunidade que valoriza sustentabilidade e
            prosperidade local.
          </p>
        </div>

        <div
          class="group rounded-2xl border border-gray-200 bg-white p-8 transition-all duration-300 hover:-translate-y-2 hover:border-emerald-200 hover:shadow-xl hover:shadow-emerald-500/10"
        >
          <div
            class="mb-6 inline-flex h-14 w-14 items-center justify-center rounded-xl bg-gradient-to-br from-green-100 to-green-200 transition-all duration-300 group-hover:from-green-500 group-hover:to-green-600"
          >
            <Leaf
              :size="28"
              class="text-green-600 transition-colors duration-300 group-hover:text-white"
              :stroke-width="2"
            />
          </div>
          <h3 class="mb-3 text-xl font-bold text-gray-900">
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
