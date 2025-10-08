<script setup lang="ts">
import InputText from "primevue/inputtext";
import Password from "primevue/password";
import Button from "primevue/button";

import { ref } from "vue";
import { z } from "zod";

import { useStore } from "vuex";
const store = useStore();

import { useRouter } from "vue-router";
const router = useRouter();

const email = ref<string>("");
const password = ref<string>("");
const loading = ref<boolean>(false);
const errors = ref<{ email?: string; password?: string }>({});
const hasLoginError = ref<boolean>(false);

import { Toast } from "primevue";
import { useToast } from "primevue/usetoast";
const toast = useToast();

const showLoginErrorToast = () => {
  toast.add({
    severity: "error",
    summary: "Erro",
    detail: "Credenciais inválidas. Tente novamente.",
    life: 3000,
  });

  hasLoginError.value = true;
};

const loginSchema = z.object({
  email: z.email({ message: "E-mail inválido" }),
  password: z
    .string()
    .min(6, { message: "A senha precisa ter ao menos 6 caracteres" }),
});

const login = async () => {
  hasLoginError.value = false;
  errors.value = {};

  const result = loginSchema.safeParse({
    email: email.value,
    password: password.value,
  });

  if (!result.success) {
    result.error.issues.forEach((err: z.core.$ZodIssue) => {
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
    router.push({ name: "Dashboard" });
  } catch (error) {
    showLoginErrorToast();
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
        <i class="pi pi-users text-primary-color !text-3xl"></i>
      </div>
      <h1 class="mb-2 text-2xl font-bold text-gray-900">Entrar no Ecoply</h1>
      <p class="text-gray-600">Acesse sua conta para continuar</p>
    </div>

    <form @submit.prevent="login" class="flex flex-col gap-5">
      <Toast />
      <div class="flex flex-col gap-2">
        <label for="email" class="font-medium text-gray-700">E-mail</label>
        <InputText
          id="email"
          v-model="email"
          placeholder="Digite seu e-mail"
          :invalid="!!errors.email || hasLoginError"
          :disabled="loading"
          class="w-full"
          size="large"
        />
        <small v-if="errors.email" class="text-red-600">
          {{ errors.email }}
        </small>
      </div>

      <div class="flex flex-col gap-2">
        <label for="password" class="font-medium text-gray-700">Senha</label>
        <Password
          id="password"
          v-model="password"
          placeholder="Digite sua senha"
          :invalid="!!errors.password || hasLoginError"
          :disabled="loading"
          :feedback="false"
          toggleMask
          class="w-full"
          inputClass="w-full"
          size="large"
        />
        <small v-if="errors.password" class="text-red-600">
          {{ errors.password }}
        </small>
      </div>

      <Button
        type="submit"
        label="Entrar"
        icon="pi pi-sign-in"
        :loading="loading"
        :disabled="loading"
        severity="primary"
        size="large"
        class="mt-2 w-full"
      />
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

<style scoped></style>
