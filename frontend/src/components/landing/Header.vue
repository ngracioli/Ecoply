<script setup lang="ts"></script>

<template>
  <header class="bg-primary-bg-color fixed z-100 w-full py-2">
    <nav class="flex items-center justify-evenly p-4">
      <div
        class="logo-container group flex cursor-default flex-row items-center"
      >
        <i
          class="pi pi-bolt text-secondary-color group-hover:text-tertiary-color text-4xl font-extrabold transition-all duration-500 ease-out group-hover:scale-105 group-hover:rotate-6"
        ></i>
        <h1
          class="text-secondary-color group-hover:text-tertiary-color text-4xl font-bold transition-all duration-500 ease-out group-hover:scale-102"
        >
          Ecoply
        </h1>
      </div>
      <ul class="flex space-x-12 text-xl">
        <li>
          <router-link
            to="/#home"
            class="nav-link text-secondary-color hover:text-tertiary-color relative inline-block py-2 transition-all duration-400 ease-out hover:-translate-y-1 hover:scale-105"
            >Início</router-link
          >
        </li>
        <li>
          <router-link
            to="/#how-it-works"
            class="nav-link text-secondary-color hover:text-tertiary-color relative inline-block py-2 transition-all duration-400 ease-out hover:-translate-y-1 hover:scale-105"
            >Como funciona</router-link
          >
        </li>
        <li>
          <router-link
            to="/#benefits"
            class="nav-link text-secondary-color hover:text-tertiary-color relative inline-block py-2 transition-all duration-400 ease-out hover:-translate-y-1 hover:scale-105"
            >Benefícios</router-link
          >
        </li>
        <li>
          <router-link
            to="/#offer-demonstration"
            class="nav-link text-secondary-color hover:text-tertiary-color relative inline-block py-2 transition-all duration-400 ease-out hover:-translate-y-1 hover:scale-105"
            >Ofertas</router-link
          >
        </li>
      </ul>
      <button
        class="cta-button bg-secondary-color text-primary-color text-md hover:border-primary-color/20 hover:bg-tertiary-color relative cursor-pointer overflow-hidden rounded border-2 border-transparent px-4 py-2 font-bold transition-all duration-400 ease-out hover:-translate-y-1 hover:scale-105 hover:shadow-2xl active:-translate-y-0.5 active:scale-[1.02]"
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
  background: linear-gradient(90deg, transparent, currentColor, transparent);
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
  transform: translateX(-50%);
  border-radius: 1px;
}

.nav-link:hover::after {
  width: 120%;
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
    rgba(30, 65, 4, 0.1),
    transparent
  );
  transition: left 0.6s ease-out;
}

.cta-button:hover::before {
  left: 100%;
}
</style>
