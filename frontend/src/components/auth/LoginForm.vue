<script setup lang="ts">
import InputText from "../../components/shared/forms/InputText.vue";
import Password from "../../components/shared/forms/Password.vue";

import { ref } from "vue";
import { useStore } from "vuex";
const store = useStore();

const email = ref<string>("");
const password = ref<string>("");

const login = async () => {
  try {
    await store.dispatch("auth/login", {
      email: email.value,
      password: password.value,
    });
  } catch (error) {
    console.error("Erro ao fazer login:", error);
  }
};
</script>

<template>
  <div class="rounded-2xl bg-white p-8 shadow-xl">
    <div class="mb-8 text-center">
      <div
        class="bg-primary-color/10 mb-4 inline-flex h-16 w-16 items-center justify-center rounded-full"
      >
        <i class="pi pi-users text-primary-color text-2xl"></i>
      </div>
      <h1 class="mb-2 text-2xl font-bold text-gray-900">Entrar no Ecoply</h1>
      <p class="text-gray-600">Acesse sua conta para continuar</p>
    </div>

    <form class="space-y-6" @submit.prevent="login">
      <div>
        <label for="email" class="mb-2 block text-sm font-medium text-gray-700">
          E-mail
        </label>
        <InputText
          id="email"
          v-model="email"
          placeholder="Digite seu e-mail"
          class="w-full"
        />
      </div>

      <div>
        <label
          for="password"
          class="mb-2 block text-sm font-medium text-gray-700"
        >
          Senha
        </label>
        <Password
          id="password"
          v-model="password"
          placeholder="Digite sua senha"
          class="w-full"
        />
      </div>

      <button
        type="submit"
        class="bg-primary-color hover:bg-primary-dark-color w-full transform rounded-xl px-6 py-3 font-semibold text-white transition duration-200 ease-in-out hover:scale-[1.02]"
      >
        Entrar
      </button>
    </form>

    <div class="mt-6 text-center">
      <p class="text-gray-600">
        Não tem uma conta?
        <router-link
          to="/register"
          class="text-primary-color hover:text-primary-dark-color ml-1 font-medium transition-colors duration-200"
        >
          Cadastre-se
        </router-link>
      </p>
    </div>
  </div>
</template>
