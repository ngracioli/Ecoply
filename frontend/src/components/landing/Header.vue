<script setup lang="ts">
import { useRouter } from "vue-router";
import { ref, onMounted, onUnmounted } from "vue";
import { Zap, Menu, X } from "lucide-vue-next";

const router = useRouter();
const scrolled = ref(false);
const mobileMenuOpen = ref(false);

const goToDashboard = () => {
  router.push({ name: "Dashboard" });
  mobileMenuOpen.value = false;
};

const handleScroll = () => {
  scrolled.value = window.scrollY > 20;
};

const closeMobileMenu = () => {
  mobileMenuOpen.value = false;
};

const handleNavClick = () => {
  closeMobileMenu();
};

onMounted(() => {
  window.addEventListener("scroll", handleScroll);
});

onUnmounted(() => {
  window.removeEventListener("scroll", handleScroll);
});
</script>

<template>
  <header
    :class="[
      'fixed top-0 right-0 left-0 z-50 transition-all duration-300',
      scrolled ? 'bg-white/80 shadow-sm backdrop-blur-xl' : 'bg-transparent',
    ]"
  >
    <nav
      class="mx-auto flex max-w-7xl items-center justify-between px-4 py-3 sm:px-6 lg:px-8 lg:py-4"
    >
      <div class="logo-container group flex cursor-pointer items-center gap-2">
        <div
          :class="[
            'flex h-8 w-8 items-center justify-center rounded-lg transition-all duration-300 sm:h-10 sm:w-10 sm:rounded-xl',
            scrolled
              ? 'bg-gradient-to-br from-emerald-500 to-emerald-600'
              : 'bg-gradient-to-br from-emerald-400/90 to-emerald-600/90',
          ]"
        >
          <Zap :size="20" :stroke-width="2.5" class="text-white sm:size-6" />
        </div>
        <h1
          :class="[
            'text-xl font-bold transition-colors duration-300 sm:text-2xl',
            scrolled ? 'text-gray-900' : 'text-white',
          ]"
        >
          Ecoply
        </h1>
      </div>

      <ul class="hidden items-center gap-8 lg:flex">
        <li>
          <a
            href="#home"
            :class="[
              'nav-link text-sm font-medium transition-colors duration-300',
              scrolled
                ? 'text-gray-700 hover:text-emerald-600'
                : 'text-white/90 hover:text-white',
            ]"
          >
            Início
          </a>
        </li>
        <li>
          <a
            href="#how-it-works"
            :class="[
              'nav-link text-sm font-medium transition-colors duration-300',
              scrolled
                ? 'text-gray-700 hover:text-emerald-600'
                : 'text-white/90 hover:text-white',
            ]"
          >
            Como funciona
          </a>
        </li>
        <li>
          <a
            href="#benefits"
            :class="[
              'nav-link text-sm font-medium transition-colors duration-300',
              scrolled
                ? 'text-gray-700 hover:text-emerald-600'
                : 'text-white/90 hover:text-white',
            ]"
          >
            Benefícios
          </a>
        </li>
        <li>
          <a
            href="#offer-demonstration"
            :class="[
              'nav-link text-sm font-medium transition-colors duration-300',
              scrolled
                ? 'text-gray-700 hover:text-emerald-600'
                : 'text-white/90 hover:text-white',
            ]"
          >
            Ofertas
          </a>
        </li>
      </ul>

      <div class="flex items-center gap-3">
        <button
          @click="goToDashboard"
          class="group relative hidden overflow-hidden rounded-full bg-emerald-600 px-6 py-2.5 text-sm font-medium text-white transition-all duration-300 hover:bg-emerald-700 hover:shadow-lg hover:shadow-emerald-500/25 sm:block"
        >
          <span class="relative z-10">Acessar Dashboard</span>
          <div
            class="absolute inset-0 translate-y-full bg-emerald-700 transition-transform duration-300 group-hover:translate-y-0"
          ></div>
        </button>

        <button
          @click="mobileMenuOpen = !mobileMenuOpen"
          :class="[
            'flex h-10 w-10 items-center justify-center rounded-lg transition-all duration-300 lg:hidden',
            scrolled
              ? 'bg-gray-100 hover:bg-gray-200'
              : 'bg-white/10 backdrop-blur-sm hover:bg-white/20',
          ]"
          :aria-label="mobileMenuOpen ? 'Fechar menu' : 'Abrir menu'"
          :aria-expanded="mobileMenuOpen"
        >
          <X
            v-if="mobileMenuOpen"
            :size="24"
            :class="[
              'transition-colors duration-300',
              scrolled ? 'text-gray-700' : 'text-white',
            ]"
          />
          <Menu
            v-else
            :size="24"
            :class="[
              'transition-colors duration-300',
              scrolled ? 'text-gray-700' : 'text-white',
            ]"
          />
        </button>
      </div>
    </nav>

    <Transition
      enter-active-class="transition-all duration-300 ease-out"
      enter-from-class="opacity-0 -translate-y-2"
      enter-to-class="opacity-100 translate-y-0"
      leave-active-class="transition-all duration-200 ease-in"
      leave-from-class="opacity-100 translate-y-0"
      leave-to-class="opacity-0 -translate-y-2"
    >
      <div
        v-if="mobileMenuOpen"
        class="border-t border-gray-200/20 bg-white/95 backdrop-blur-xl lg:hidden"
      >
        <ul class="flex flex-col px-4 py-3">
          <li>
            <a
              href="#home"
              @click="handleNavClick"
              class="block rounded-lg px-4 py-3 text-base font-medium text-gray-700 transition-colors duration-200 hover:bg-emerald-50 hover:text-emerald-600"
            >
              Início
            </a>
          </li>
          <li>
            <a
              href="#how-it-works"
              @click="handleNavClick"
              class="block rounded-lg px-4 py-3 text-base font-medium text-gray-700 transition-colors duration-200 hover:bg-emerald-50 hover:text-emerald-600"
            >
              Como funciona
            </a>
          </li>
          <li>
            <a
              href="#benefits"
              @click="handleNavClick"
              class="block rounded-lg px-4 py-3 text-base font-medium text-gray-700 transition-colors duration-200 hover:bg-emerald-50 hover:text-emerald-600"
            >
              Benefícios
            </a>
          </li>
          <li>
            <a
              href="#offer-demonstration"
              @click="handleNavClick"
              class="block rounded-lg px-4 py-3 text-base font-medium text-gray-700 transition-colors duration-200 hover:bg-emerald-50 hover:text-emerald-600"
            >
              Ofertas
            </a>
          </li>
          <li class="mt-2 border-t border-gray-200 pt-3">
            <button
              @click="goToDashboard"
              class="group relative w-full overflow-hidden rounded-full bg-emerald-600 px-6 py-3 text-base font-medium text-white transition-all duration-300 hover:bg-emerald-700 hover:shadow-lg hover:shadow-emerald-500/25"
            >
              <span class="relative z-10">Acessar Dashboard</span>
              <div
                class="absolute inset-0 translate-y-full bg-emerald-700 transition-transform duration-300 group-hover:translate-y-0"
              ></div>
            </button>
          </li>
        </ul>
      </div>
    </Transition>
  </header>
</template>

<style scoped>
.nav-link {
  position: relative;
}

.nav-link::after {
  content: "";
  position: absolute;
  bottom: -4px;
  left: 0;
  width: 0;
  height: 2px;
  background: currentColor;
  transition: width 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.nav-link:hover::after {
  width: 100%;
}
</style>
