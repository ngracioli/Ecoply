<script setup lang="ts">
import InputText from "primevue/inputtext";
import Password from "primevue/password";
import Button from "primevue/button";
import { Toast } from "primevue";
import { useToast } from "primevue/usetoast";

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
const toast = useToast();

// Limpa o erro de um campo específico quando o usuário digita
const clearFieldError = (field: "email" | "password") => {
  if (errors.value[field]) {
    errors.value[field] = undefined;
  }
  if (hasLoginError.value) {
    hasLoginError.value = false;
  }
};

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
  password: z.string().min(1, { message: "Senha é obrigatória" }),
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
  <div class="flex flex-col gap-6">
    <div class="mb-4 text-center">
      <div
        class="mb-4 inline-flex h-16 w-16 items-center justify-center rounded-full bg-gradient-to-br from-emerald-400 to-emerald-600 shadow-lg"
      >
        <i class="pi pi-lock !text-3xl text-white"></i>
      </div>
      <h1 class="mb-2 text-3xl font-bold text-gray-900">Bem-vindo de volta</h1>
      <p class="text-gray-600">Entre com suas credenciais para continuar</p>
    </div>

    <form @submit.prevent="login" class="flex flex-col gap-5">
      <Toast />
      <div class="flex flex-col gap-2">
        <label for="email" class="text-sm font-semibold text-gray-800"
          >E-mail</label
        >
        <InputText
          id="email"
          v-model="email"
          @input="clearFieldError('email')"
          placeholder="seu@email.com"
          :invalid="!!errors.email || hasLoginError"
          :disabled="loading"
          class="w-full !rounded-xl !border-2 !border-gray-300 !bg-white !py-3 transition-colors hover:!border-gray-400 focus:!border-emerald-500"
          size="large"
        />
        <small v-if="errors.email" class="font-medium text-red-500">
          {{ errors.email }}
        </small>
      </div>

      <div class="flex flex-col gap-2">
        <label for="password" class="text-sm font-semibold text-gray-800"
          >Senha</label
        >
        <Password
          id="password"
          v-model="password"
          @input="clearFieldError('password')"
          placeholder="••••••••"
          :invalid="!!errors.password || hasLoginError"
          :disabled="loading"
          :feedback="false"
          toggleMask
          class="w-full"
          inputClass="w-full !rounded-xl !border-2 !border-gray-300 focus:!border-emerald-500 !py-3 !bg-white hover:!border-gray-400 transition-colors"
          size="large"
        />
        <small v-if="errors.password" class="font-medium text-red-500">
          {{ errors.password }}
        </small>
      </div>

      <Button
        type="submit"
        label="Entrar"
        icon="pi pi-sign-in"
        :loading="loading"
        :disabled="loading"
        class="mt-4 w-full !rounded-xl !border-none !bg-gradient-to-r !from-emerald-600 !to-emerald-700 !py-3 !text-base !font-semibold shadow-lg transition-all duration-300 hover:!from-emerald-700 hover:!to-emerald-800 hover:shadow-xl"
        size="large"
      />
    </form>

    <div class="mt-4 text-center">
      <p class="text-gray-700">
        Não tem uma conta?
        <router-link
          to="/register"
          class="ml-1 font-semibold text-emerald-600 transition-colors duration-200 hover:text-emerald-700"
        >
          Cadastre-se gratuitamente
        </router-link>
      </p>
    </div>
  </div>
</template>

<style scoped></style>
