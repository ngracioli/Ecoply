<script setup lang="ts">
import { useRouter } from "vue-router";
import { ref, onMounted, onUnmounted } from "vue";
import { Zap } from "lucide-vue-next";

const router = useRouter();
const scrolled = ref(false);

const goToDashboard = () => {
  router.push({ name: "Dashboard" });
};

const handleScroll = () => {
  scrolled.value = window.scrollY > 20;
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
    <nav class="mx-auto flex max-w-7xl items-center justify-between px-8 py-4">
      <div class="logo-container group flex cursor-pointer items-center gap-2">
        <div
          :class="[
            'flex h-10 w-10 items-center justify-center rounded-xl transition-all duration-300',
            scrolled
              ? 'bg-gradient-to-br from-emerald-500 to-emerald-600'
              : 'bg-gradient-to-br from-emerald-400/90 to-emerald-600/90',
          ]"
        >
          <Zap :size="24" :stroke-width="2.5" class="text-white" />
        </div>
        <h1
          :class="[
            'text-2xl font-bold transition-colors duration-300',
            scrolled ? 'text-gray-900' : 'text-white',
          ]"
        >
          Ecoply
        </h1>
      </div>

      <ul class="flex items-center gap-8">
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

      <button
        @click="goToDashboard"
        class="group relative overflow-hidden rounded-full bg-emerald-600 px-6 py-2.5 text-sm font-medium text-white transition-all duration-300 hover:bg-emerald-700 hover:shadow-lg hover:shadow-emerald-500/25"
      >
        <span class="relative z-10">Acessar Dashboard</span>
        <div
          class="absolute inset-0 translate-y-full bg-emerald-700 transition-transform duration-300 group-hover:translate-y-0"
        ></div>
      </button>
    </nav>
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
