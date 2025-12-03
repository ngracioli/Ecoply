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
  <div class="flex flex-col gap-4 px-4 sm:gap-6 sm:px-0">
    <div class="mb-2 text-center sm:mb-4">
      <div
        class="mb-3 inline-flex h-12 w-12 items-center justify-center rounded-full bg-gradient-to-br from-emerald-400 to-emerald-600 shadow-lg sm:mb-4 sm:h-16 sm:w-16"
      >
        <i class="pi pi-lock !text-2xl text-white sm:!text-3xl"></i>
      </div>
      <h1
        class="mb-1 px-4 text-2xl font-bold text-gray-900 sm:mb-2 sm:px-0 sm:text-3xl"
      >
        Bem-vindo de volta
      </h1>
      <p class="px-2 text-sm text-gray-600 sm:px-0 sm:text-base">
        Entre com suas credenciais para continuar
      </p>
    </div>

    <form
      @submit.prevent="login"
      class="mobile-scrollable flex flex-col gap-4 sm:gap-5"
    >
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
          class="w-full touch-manipulation !rounded-xl !border-2 !border-gray-300 !bg-white !py-2.5 !text-base transition-colors hover:!border-gray-400 focus:!border-emerald-500 sm:!py-3"
          size="large"
        />
        <small
          v-if="errors.email"
          class="text-xs font-medium text-red-500 sm:text-sm"
        >
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
          inputClass="w-full !rounded-xl !border-2 !border-gray-300 focus:!border-emerald-500 !py-2.5 sm:!py-3 !bg-white hover:!border-gray-400 transition-colors !text-base touch-manipulation"
          size="large"
        />
        <small
          v-if="errors.password"
          class="text-xs font-medium text-red-500 sm:text-sm"
        >
          {{ errors.password }}
        </small>
      </div>

      <Button
        type="submit"
        label="Entrar"
        icon="pi pi-sign-in"
        :loading="loading"
        :disabled="loading"
        class="mt-2 min-h-[44px] w-full touch-manipulation !rounded-xl !border-none !bg-gradient-to-r !from-emerald-600 !to-emerald-700 !py-3 !text-base !font-semibold shadow-lg transition-all duration-300 hover:!from-emerald-700 hover:!to-emerald-800 hover:shadow-xl active:scale-[0.98] sm:mt-4 sm:!py-3"
        size="large"
      />
    </form>

    <div class="mt-2 text-center sm:mt-4">
      <p class="px-2 text-sm text-gray-700 sm:px-0 sm:text-base">
        Não tem uma conta?
        <router-link
          to="/register"
          class="-my-3 ml-1 inline-block min-h-[44px] touch-manipulation leading-[44px] font-semibold text-emerald-600 transition-colors duration-200 hover:text-emerald-700 active:text-emerald-800"
        >
          Cadastre-se gratuitamente
        </router-link>
      </p>
    </div>
  </div>
</template>

<style scoped></style>
