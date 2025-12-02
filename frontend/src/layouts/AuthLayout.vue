<script setup lang="ts">
import { useRouter } from "vue-router";
const router = useRouter();

const goToHome = () => {
  router.push({ name: "Home" });
};
</script>

<template>
  <main
    class="relative flex min-h-dvh items-center justify-center overflow-hidden p-3 sm:p-6 md:p-10 lg:p-16 xl:p-24"
  >
    <div class="absolute inset-0 z-0">
      <img
        src="/ecoply_login.jpg"
        alt="Background"
        class="h-full w-full object-cover"
      />
      <div class="absolute inset-0 bg-black/40 backdrop-blur-2xl"></div>
    </div>

    <div
      class="relative z-10 grid h-[calc(100dvh-1.5rem)] w-full grid-cols-1 gap-3 sm:h-[calc(100dvh-3rem)] sm:gap-4 md:h-[calc(100dvh-5rem)] md:gap-6 lg:h-[calc(100dvh-8rem)] lg:grid-cols-2 lg:gap-8 xl:h-[calc(100dvh-12rem)]"
    >
      <section
        class="relative hidden items-center justify-center overflow-hidden rounded-2xl bg-transparent shadow-2xl lg:flex lg:rounded-3xl"
      >
        <img
          src="/ecoply_login.jpg"
          alt="Ilustração Ecoply"
          class="absolute inset-0 h-full w-full rounded-2xl object-cover lg:rounded-3xl"
        />

        <div
          class="absolute inset-0 rounded-2xl bg-gradient-to-b from-black/40 via-black/20 to-black/50 lg:rounded-3xl"
        ></div>

        <div class="absolute top-4 left-4 z-10 lg:top-6 lg:left-8">
          <h1 class="text-2xl font-bold text-white drop-shadow-lg lg:text-3xl">
            ECOPLY
          </h1>
        </div>

        <button
          @click="goToHome"
          class="absolute top-4 right-4 z-10 rounded-lg bg-white/90 px-3 py-1.5 text-xs font-medium text-emerald-700 shadow-lg backdrop-blur-sm transition-all duration-300 hover:bg-white hover:shadow-xl lg:top-6 lg:right-8 lg:px-4 lg:py-2 lg:text-sm"
        >
          Voltar ao site
        </button>

        <div
          class="absolute right-0 bottom-4 left-0 z-10 px-4 text-center lg:bottom-8 lg:px-8"
        >
          <h2
            class="mb-2 text-3xl font-bold text-white drop-shadow-2xl lg:mb-3 lg:text-4xl xl:text-5xl"
          >
            Conectando Pessoas,
          </h2>
          <h2
            class="text-3xl font-bold text-emerald-300 drop-shadow-2xl lg:text-4xl xl:text-5xl"
          >
            Energizando o Futuro
          </h2>
        </div>
      </section>

      <section
        class="relative flex flex-col overflow-hidden rounded-2xl bg-white shadow-2xl lg:rounded-3xl"
      >
        <div class="absolute top-3 left-3 z-10 lg:hidden">
          <h1 class="text-xl font-bold text-emerald-600">ECOPLY</h1>
        </div>

        <button
          @click="goToHome"
          class="absolute top-3 right-3 z-10 rounded-lg bg-emerald-50 px-3 py-1.5 text-xs font-medium text-emerald-700 shadow-md transition-all duration-300 hover:bg-emerald-100 hover:shadow-lg lg:hidden"
        >
          Voltar
        </button>

        <div
          class="scrollbar-custom flex-1 overflow-y-auto p-4 pt-14 sm:p-6 sm:pt-6 md:p-8 lg:p-12"
        >
          <div class="mx-auto flex min-h-full w-full max-w-md items-center">
            <div class="w-full">
              <slot />
            </div>
          </div>
        </div>
      </section>
    </div>
  </main>
</template>

<style scoped>
section {
  box-shadow:
    0 10px 30px -5px rgba(0, 0, 0, 0.1),
    0 20px 60px -10px rgba(0, 0, 0, 0.05);
}

.scrollbar-custom::-webkit-scrollbar {
  width: 8px;
}

.scrollbar-custom::-webkit-scrollbar-track {
  background: rgba(0, 0, 0, 0.05);
  border-radius: 10px;
  margin: 8px 0;
}

.scrollbar-custom::-webkit-scrollbar-thumb {
  background: linear-gradient(180deg, #059669, #10b981);
  border-radius: 10px;
  transition: all 0.3s ease;
}

.scrollbar-custom::-webkit-scrollbar-thumb:hover {
  background: linear-gradient(180deg, #047857, #059669);
}

.scrollbar-custom {
  scrollbar-width: thin;
  scrollbar-color: #059669 rgba(0, 0, 0, 0.05);
}
</style>
