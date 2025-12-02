<script setup lang="ts">
import {
  Facebook,
  Twitter,
  Instagram,
  Linkedin,
  Mail,
  Phone,
  MapPin,
} from "lucide-vue-next";
</script>

<template>
  <footer class="border-t border-gray-200 bg-white">
    <div class="mx-auto max-w-7xl px-4 py-8 sm:px-6 sm:py-10 lg:px-8 lg:py-12">
      <div class="grid gap-8 sm:grid-cols-2 sm:gap-10 lg:grid-cols-4 lg:gap-12">
        <div class="sm:col-span-2">
          <div class="mb-3 flex items-center gap-2 sm:mb-4">
            <div
              class="flex h-9 w-9 items-center justify-center rounded-lg bg-gradient-to-br from-emerald-500 to-emerald-600 sm:h-10 sm:w-10 sm:rounded-xl"
            >
              <i class="pi pi-bolt text-lg font-bold text-white sm:text-xl"></i>
            </div>
            <h2 class="text-xl font-bold text-gray-900 sm:text-2xl">Ecoply</h2>
          </div>
          <p
            class="mb-4 max-w-sm text-sm leading-relaxed text-gray-600 sm:mb-6"
          >
            Conectando produtores e consumidores de energia renovável para um
            futuro mais sustentável e descentralizado.
          </p>
          <div class="flex flex-wrap gap-2.5 sm:gap-3">
            <a
              href="#"
              class="flex h-9 w-9 items-center justify-center rounded-lg border border-gray-200 text-gray-600 transition-all duration-300 hover:-translate-y-1 hover:border-emerald-500 hover:bg-emerald-50 hover:text-emerald-600 sm:h-10 sm:w-10"
            >
              <Facebook :size="17" :stroke-width="2" class="sm:size-[18px]" />
            </a>
            <a
              href="#"
              class="flex h-9 w-9 items-center justify-center rounded-lg border border-gray-200 text-gray-600 transition-all duration-300 hover:-translate-y-1 hover:border-emerald-500 hover:bg-emerald-50 hover:text-emerald-600 sm:h-10 sm:w-10"
            >
              <Twitter :size="17" :stroke-width="2" class="sm:size-[18px]" />
            </a>
            <a
              href="#"
              class="flex h-9 w-9 items-center justify-center rounded-lg border border-gray-200 text-gray-600 transition-all duration-300 hover:-translate-y-1 hover:border-emerald-500 hover:bg-emerald-50 hover:text-emerald-600 sm:h-10 sm:w-10"
            >
              <Instagram :size="17" :stroke-width="2" class="sm:size-[18px]" />
            </a>
            <a
              href="#"
              class="flex h-9 w-9 items-center justify-center rounded-lg border border-gray-200 text-gray-600 transition-all duration-300 hover:-translate-y-1 hover:border-emerald-500 hover:bg-emerald-50 hover:text-emerald-600 sm:h-10 sm:w-10"
            >
              <Linkedin :size="17" :stroke-width="2" class="sm:size-[18px]" />
            </a>
          </div>
        </div>

        <div>
          <h3 class="mb-3 text-sm font-semibold text-gray-900 sm:mb-4">
            Navegação
          </h3>
          <ul class="space-y-2.5 sm:space-y-3">
            <li>
              <a
                href="#home"
                class="text-sm text-gray-600 transition-colors hover:text-emerald-600"
              >
                Início
              </a>
            </li>
            <li>
              <a
                href="#how-it-works"
                class="text-sm text-gray-600 transition-colors hover:text-emerald-600"
              >
                Como Funciona
              </a>
            </li>
            <li>
              <a
                href="#benefits"
                class="text-sm text-gray-600 transition-colors hover:text-emerald-600"
              >
                Benefícios
              </a>
            </li>
            <li>
              <a
                href="#offer-demonstration"
                class="text-sm text-gray-600 transition-colors hover:text-emerald-600"
              >
                Ofertas
              </a>
            </li>
          </ul>
        </div>

        <div>
          <h3 class="mb-3 text-sm font-semibold text-gray-900 sm:mb-4">
            Contato
          </h3>
          <ul class="space-y-2.5 sm:space-y-3">
            <li>
              <a
                href="mailto:contato@ecoply.com.br"
                class="flex items-center gap-2 text-sm text-gray-600 transition-colors hover:text-emerald-600"
              >
                <Mail :size="16" :stroke-width="2" />
                <span>contato@ecoply.com.br</span>
              </a>
            </li>
            <li>
              <a
                href="tel:+551999999999"
                class="flex items-center gap-2 text-sm text-gray-600 transition-colors hover:text-emerald-600"
              >
                <Phone :size="16" :stroke-width="2" />
                <span>(19) 9999-9999</span>
              </a>
            </li>
            <li class="flex items-center gap-2 text-sm text-gray-600">
              <MapPin :size="16" :stroke-width="2" />
              <span>Rio Claro, SP</span>
            </li>
          </ul>
        </div>
      </div>

      <div
        class="mt-8 flex flex-col items-center justify-between gap-3 border-t border-gray-200 pt-6 sm:mt-10 sm:gap-4 sm:pt-7 md:flex-row lg:mt-12 lg:pt-8"
      >
        <p class="text-center text-sm text-gray-600 md:text-left">
          © 2025
          <span class="font-semibold text-emerald-600">Ecoply</span>. Todos os
          direitos reservados.
        </p>
        <div class="flex gap-4 sm:gap-6">
          <a
            href="#"
            class="text-sm text-gray-600 transition-colors hover:text-emerald-600"
          >
            Privacidade
          </a>
          <a
            href="#"
            class="text-sm text-gray-600 transition-colors hover:text-emerald-600"
          >
            Termos de Uso
          </a>
        </div>
      </div>
    </div>
  </footer>
</template>

<style scoped></style>
