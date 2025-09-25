<script setup lang="ts">
import EcoForm from "../ui/EcoForm.vue";
import EcoFieldInput from "../ui/EcoFieldInput.vue";
import EcoFormButton from "../ui/EcoFormButton.vue";

import { ref } from "vue";
import { z } from "zod";
import type { ZodIssue } from "zod";
import { useStore } from "vuex";
const store = useStore();

const email = ref<string>("");
const password = ref<string>("");
const loading = ref(false);
const errors = ref<{ email?: string; password?: string }>({});

const loginSchema = z.object({
  email: z.string().email({ message: "E-mail inválido" }),
  password: z
    .string()
    .min(6, { message: "A senha precisa ter ao menos 6 caracteres" }),
});

const login = async () => {
  errors.value = {};

  const result = loginSchema.safeParse({
    email: email.value,
    password: password.value,
  });

  if (!result.success) {
    result.error.issues.forEach((err: ZodIssue) => {
      if (err.path[0] === "email") {
        errors.value.email = err.message;
      } else if (err.path[0] === "password") {
        errors.value.password = err.message;
      }
    });
    return;
  }

  try {
    loading.value = true;
    await store.dispatch("auth/login", {
      email: email.value,
      password: password.value,
    });
  } catch (error) {
    console.error("Erro ao fazer login:", error);
  } finally {
    loading.value = false;
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

    <EcoForm @submit="login">
      <EcoFieldInput
        id="email"
        label="E-mail"
        v-model="email"
        :required="true"
        placeholder="Digite seu e-mail"
        :error-message="errors.email"
      />

      <EcoFieldInput
        id="password"
        label="Senha"
        password
        v-model="password"
        :required="true"
        placeholder="Digite sua senha"
        :error-message="errors.password"
      />

      <EcoFormButton
        type="submit"
        variant="primary"
        :full="true"
        :loading="loading"
      >
        Entrar
      </EcoFormButton>
    </EcoForm>

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
