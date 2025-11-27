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
    <div class="mx-auto max-w-7xl px-8 py-12">
      <div class="grid gap-12 md:grid-cols-4">
        <div class="md:col-span-2">
          <div class="mb-4 flex items-center gap-2">
            <div
              class="flex h-10 w-10 items-center justify-center rounded-xl bg-gradient-to-br from-emerald-500 to-emerald-600"
            >
              <i class="pi pi-bolt text-xl font-bold text-white"></i>
            </div>
            <h2 class="text-2xl font-bold text-gray-900">Ecoply</h2>
          </div>
          <p class="mb-6 max-w-sm text-sm leading-relaxed text-gray-600">
            Conectando produtores e consumidores de energia renovável para um
            futuro mais sustentável e descentralizado.
          </p>
          <div class="flex gap-3">
            <a
              href="#"
              class="flex h-10 w-10 items-center justify-center rounded-lg border border-gray-200 text-gray-600 transition-all duration-300 hover:-translate-y-1 hover:border-emerald-500 hover:bg-emerald-50 hover:text-emerald-600"
            >
              <Facebook :size="18" :stroke-width="2" />
            </a>
            <a
              href="#"
              class="flex h-10 w-10 items-center justify-center rounded-lg border border-gray-200 text-gray-600 transition-all duration-300 hover:-translate-y-1 hover:border-emerald-500 hover:bg-emerald-50 hover:text-emerald-600"
            >
              <Twitter :size="18" :stroke-width="2" />
            </a>
            <a
              href="#"
              class="flex h-10 w-10 items-center justify-center rounded-lg border border-gray-200 text-gray-600 transition-all duration-300 hover:-translate-y-1 hover:border-emerald-500 hover:bg-emerald-50 hover:text-emerald-600"
            >
              <Instagram :size="18" :stroke-width="2" />
            </a>
            <a
              href="#"
              class="flex h-10 w-10 items-center justify-center rounded-lg border border-gray-200 text-gray-600 transition-all duration-300 hover:-translate-y-1 hover:border-emerald-500 hover:bg-emerald-50 hover:text-emerald-600"
            >
              <Linkedin :size="18" :stroke-width="2" />
            </a>
          </div>
        </div>

        <div>
          <h3 class="mb-4 text-sm font-semibold text-gray-900">Navegação</h3>
          <ul class="space-y-3">
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
          <h3 class="mb-4 text-sm font-semibold text-gray-900">Contato</h3>
          <ul class="space-y-3">
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
        class="mt-12 flex flex-col items-center justify-between gap-4 border-t border-gray-200 pt-8 md:flex-row"
      >
        <p class="text-sm text-gray-600">
          © 2025
          <span class="font-semibold text-emerald-600">Ecoply</span>. Todos os
          direitos reservados.
        </p>
        <div class="flex gap-6">
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
