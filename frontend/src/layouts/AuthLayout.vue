<script setup lang="ts">
import { useRouter } from "vue-router";
const router = useRouter();

const goToHome = () => {
  router.push({ name: "Home" });
};
</script>

<template>
  <main
    class="relative flex min-h-screen items-center justify-center overflow-hidden p-16 lg:p-24"
  >
    <!-- Imagem de fundo com blur e escurecida -->
    <div class="absolute inset-0 z-0">
      <img
        src="/ecoply_login.jpg"
        alt="Background"
        class="h-full w-full object-cover"
      />
      <!-- Overlay escuro + blur -->
      <div class="absolute inset-0 bg-black/40 backdrop-blur-2xl"></div>
    </div>

    <div
      class="relative z-10 grid h-[calc(100vh-8rem)] w-full grid-cols-1 gap-8 lg:h-[calc(100vh-12rem)] lg:grid-cols-2"
    >
      <!-- Painel Ilustrativo - Lado Esquerdo -->
      <section
        class="relative flex items-center justify-center overflow-hidden rounded-3xl bg-transparent shadow-2xl"
      >
        <!-- Imagem de fundo -->
        <img
          src="/ecoply_login.jpg"
          alt="Ilustração Ecoply"
          class="absolute inset-0 h-full w-full rounded-3xl object-cover"
        />

        <!-- Overlay escuro para melhor legibilidade -->
        <div
          class="absolute inset-0 rounded-3xl bg-gradient-to-b from-black/40 via-black/20 to-black/50"
        ></div>

        <!-- Logo no canto superior esquerdo -->
        <div class="absolute top-6 left-8 z-10">
          <h1 class="text-3xl font-bold text-white drop-shadow-lg">ECOPLY</h1>
        </div>

        <!-- Botão "Voltar ao site" no canto superior direito -->
        <button
          @click="goToHome"
          class="absolute top-6 right-8 z-10 rounded-lg bg-white/90 px-4 py-2 text-sm font-medium text-emerald-700 shadow-lg backdrop-blur-sm transition-all duration-300 hover:bg-white hover:shadow-xl"
        >
          Voltar ao site
        </button>

        <!-- Texto central sobreposto - BOTTOM CENTER -->
        <div class="absolute right-0 bottom-8 left-0 z-10 px-8 text-center">
          <h2
            class="mb-3 text-4xl font-bold text-white drop-shadow-2xl lg:text-5xl"
          >
            Conectando Pessoas,
          </h2>
          <h2
            class="text-4xl font-bold text-emerald-300 drop-shadow-2xl lg:text-5xl"
          >
            Energizando o Futuro
          </h2>
        </div>
      </section>

      <!-- Formulário - Lado Direito -->
      <section
        class="relative flex flex-col overflow-hidden rounded-3xl bg-white shadow-2xl"
      >
        <!-- Container com scroll -->
        <div class="scrollbar-custom flex-1 overflow-y-auto p-8 lg:p-12">
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
/* Efeito de sombra suave e moderna */
section {
  box-shadow:
    0 10px 30px -5px rgba(0, 0, 0, 0.1),
    0 20px 60px -10px rgba(0, 0, 0, 0.05);
}

/* Scrollbar customizada */
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

/* Para Firefox */
.scrollbar-custom {
  scrollbar-width: thin;
  scrollbar-color: #059669 rgba(0, 0, 0, 0.05);
}
</style>
