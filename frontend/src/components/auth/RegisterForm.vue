<script setup lang="ts">
import EcoForm from "../ui/EcoForm.vue";
// import EcoFormField from "../ui/EcoFormField.vue"; // legacy
// import EcoInput from "../ui/EcoInput.vue"; // legacy
import EcoFieldInput from "../ui/EcoFieldInput.vue";
import EcoFormButton from "../ui/EcoFormButton.vue";

import { ref } from "vue";
import type {
  RegisterFirstStepProps,
  RegisterSecondStepProps,
} from "./intefaces";
import { transformFormData } from "./transformers";

const step = ref<number>(1);

const firstStepForm = ref<RegisterFirstStepProps>({
  name: "",
  email: "",
  password: "",
  confirmPassword: "",
});
const secondStepForm = ref<RegisterSecondStepProps>({
  address: "",
  city: "",
  zip: "",
});

function nextStep() {
  if (step.value < 3) step.value++;
}

function prevStep() {
  if (step.value > 1) step.value--;
}

function submitForm() {
  console.log(
    "Form enviado:",
    transformFormData(firstStepForm.value, secondStepForm.value),
  );
  alert("Registro concluído!");
}
</script>

<template>
  <div class="rounded-2xl bg-white p-8 shadow-xl">
    <!-- Cabeçalho / Stepper -->
    <div class="mb-8 text-center">
      <div
        class="bg-primary-color/10 mb-4 inline-flex h-16 w-16 items-center justify-center rounded-full"
      >
        <i class="pi pi-users text-primary-color text-2xl"></i>
      </div>
      <h1 class="mb-2 text-2xl font-bold text-gray-900">
        Registrar-se no Ecoply
      </h1>
      <div class="mb-4 flex items-center justify-center space-x-2">
        <div
          v-for="i in 3"
          :key="i"
          class="h-2 w-8 rounded-full transition-colors duration-200"
          :class="i <= step ? 'bg-primary-color' : 'bg-gray-200'"
        ></div>
      </div>
      <p class="text-primary-color font-medium">Passo {{ step }} de 3</p>
    </div>

    <!-- Form principal -->
    <EcoForm :as="'div'" class="space-y-10" @submit="submitForm">
      <!-- Conteúdo dos Steps -->
      <div class="min-h-[280px]">
        <!-- Step 1 -->
        <div v-if="step === 1" class="space-y-6">
          <h2 class="mb-6 text-center text-lg font-semibold text-gray-900">
            Informações de acesso
          </h2>
          <EcoFieldInput
            id="name"
            label="Nome Completo"
            v-model="firstStepForm.name"
            required
            placeholder="Digite seu nome completo"
          />
          <EcoFieldInput
            id="email"
            label="E-mail"
            v-model="firstStepForm.email"
            required
            placeholder="Digite seu e-mail"
          />
          <!-- Duplicate email field removed (was duplicated) -->
          <EcoFieldInput
            id="password"
            label="Senha"
            password
            v-model="firstStepForm.password"
            required
            placeholder="Crie uma senha segura"
          />
          <EcoFieldInput
            id="confirmPassword"
            label="Confirmar Senha"
            password
            v-model="firstStepForm.confirmPassword"
            required
            placeholder="Repita a senha"
          />
        </div>

        <!-- Step 2 -->
        <div v-if="step === 2" class="space-y-6">
          <h2 class="mb-6 text-center text-lg font-semibold text-gray-900">
            Informações pessoais
          </h2>
          <EcoFieldInput
            id="name2"
            label="Nome completo"
            v-model="firstStepForm.name"
            required
            placeholder="Digite seu nome completo"
          />
          <EcoFieldInput
            id="address"
            label="Endereço"
            v-model="secondStepForm.address"
            placeholder="Digite seu endereço"
          />
        </div>

        <!-- Step 3 -->
        <div v-if="step === 3" class="space-y-6">
          <h2 class="mb-6 text-center text-lg font-semibold text-gray-900">
            Localização
          </h2>
          <EcoFieldInput
            id="city"
            label="Cidade"
            v-model="secondStepForm.city"
            required
            placeholder="Digite sua cidade"
          />
          <EcoFieldInput
            id="zip"
            label="CEP"
            v-model="secondStepForm.zip"
            required
            placeholder="Digite seu CEP"
          />
        </div>
      </div>

      <!-- Botões de navegação -->
      <div class="mt-2 flex gap-4">
        <EcoFormButton
          v-if="step > 1"
          variant="outline"
          size="md"
          full
          type="button"
          @click="prevStep"
        >
          Voltar
        </EcoFormButton>
        <EcoFormButton
          v-if="step < 3"
          variant="primary"
          size="md"
          full
          type="button"
          @click="nextStep"
        >
          Próximo
        </EcoFormButton>
        <EcoFormButton
          v-if="step === 3"
          variant="primary"
          size="md"
          full
          type="submit"
        >
          Finalizar cadastro
        </EcoFormButton>
      </div>
    </EcoForm>
  </div>
</template>
