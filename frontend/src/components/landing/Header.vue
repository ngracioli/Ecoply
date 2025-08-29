<script setup lang="ts"></script>

<template>
  <header class="bg-primary-bg-color fixed z-100 w-full py-2">
    <nav class="flex items-center justify-evenly p-4">
      <div
        class="logo-container group flex cursor-default flex-row items-center"
      >
        <i
          class="pi pi-bolt text-secondary-color text-4xl font-extrabold transition-transform duration-300 ease-out group-hover:scale-110 group-hover:rotate-12"
        ></i>
        <h1
          class="text-secondary-color text-4xl font-bold transition-transform duration-300 ease-out group-hover:scale-105"
        >
          Ecoply
        </h1>
      </div>
      <ul class="flex space-x-12 text-xl">
        <li>
          <router-link
            to="/#home"
            class="nav-link text-secondary-color hover:text-tertiary-color relative inline-block py-2 transition-all duration-300 ease-out hover:-translate-y-0.5"
            >Início</router-link
          >
        </li>
        <li>
          <router-link
            to="/#how-it-works"
            class="nav-link text-secondary-color hover:text-tertiary-color relative inline-block py-2 transition-all duration-300 ease-out hover:-translate-y-0.5"
            >Como funciona</router-link
          >
        </li>
        <li>
          <router-link
            to="/#benefits"
            class="nav-link text-secondary-color hover:text-tertiary-color relative inline-block py-2 transition-all duration-300 ease-out hover:-translate-y-0.5"
            >Benefícios</router-link
          >
        </li>
        <li>
          <router-link
            to="/#offer-demonstration"
            class="nav-link text-secondary-color hover:text-tertiary-color relative inline-block py-2 transition-all duration-300 ease-out hover:-translate-y-0.5"
            >Ofertas</router-link
          >
        </li>
      </ul>
      <button
        class="cta-button bg-secondary-color text-primary-color text-md relative cursor-pointer overflow-hidden rounded border-2 border-transparent px-4 py-2 font-bold transition-all duration-300 ease-out hover:-translate-y-0.5 hover:scale-105 hover:border-white/20 hover:shadow-xl active:-translate-y-0.5 active:scale-[1.02]"
      >
        Acessar Dashboard
      </button>
    </nav>
  </header>
</template>

<style scoped>
.nav-link::after {
  content: "";
  position: absolute;
  bottom: 0;
  left: 50%;
  width: 0;
  height: 2px;
  background: currentColor;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  transform: translateX(-50%);
}

.nav-link:hover::after {
  width: 100%;
}

.cta-button::before {
  content: "";
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(
    90deg,
    transparent,
    rgba(255, 255, 255, 0.2),
    transparent
  );
  transition: left 0.5s;
}

.cta-button:hover::before {
  left: 100%;
}
</style>
