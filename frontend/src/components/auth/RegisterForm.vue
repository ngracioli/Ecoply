<script setup lang="ts">
import InputText from "../../components/shared/forms/InputText.vue";
import Password from "../../components/shared/forms/Password.vue";
import { ref } from "vue";

const step = ref(1);

const form = ref({
  nome: "",
  email: "",
  senha: "",
  endereco: "",
  cidade: "",
  cep: "",
});

function nextStep() {
  if (step.value < 3) step.value++;
}

function prevStep() {
  if (step.value > 1) step.value--;
}

function submitForm() {
  console.log("Form enviado:", form.value);
  alert("Registro concluído!");
}
</script>

<template>
  <div class="rounded-2xl bg-white p-8 shadow-xl">
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

    <!-- Conteúdo dos Steps -->
    <div class="min-h-[280px]">
      <!-- Step 1 -->
      <div v-if="step === 1" class="space-y-6">
        <h2 class="mb-6 text-center text-lg font-semibold text-gray-900">
          Informações de acesso
        </h2>
        <div>
          <label
            for="email"
            class="mb-2 block text-sm font-medium text-gray-700"
          >
            E-mail
          </label>
          <InputText
            id="email"
            v-model="form.email"
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
            v-model="form.senha"
            placeholder="Crie uma senha segura"
            class="w-full"
          />
        </div>
      </div>

      <!-- Step 2 -->
      <div v-if="step === 2" class="space-y-6">
        <h2 class="mb-6 text-center text-lg font-semibold text-gray-900">
          Informações pessoais
        </h2>
        <div>
          <label
            for="nome"
            class="mb-2 block text-sm font-medium text-gray-700"
          >
            Nome completo
          </label>
          <InputText
            id="nome"
            v-model="form.nome"
            placeholder="Digite seu nome completo"
            class="w-full"
          />
        </div>
        <div>
          <label
            for="endereco"
            class="mb-2 block text-sm font-medium text-gray-700"
          >
            Endereço
          </label>
          <InputText
            id="endereco"
            v-model="form.endereco"
            placeholder="Digite seu endereço"
            class="w-full"
          />
        </div>
      </div>

      <!-- Step 3 -->
      <div v-if="step === 3" class="space-y-6">
        <h2 class="mb-6 text-center text-lg font-semibold text-gray-900">
          Localização
        </h2>
        <div>
          <label
            for="cidade"
            class="mb-2 block text-sm font-medium text-gray-700"
          >
            Cidade
          </label>
          <InputText
            id="cidade"
            v-model="form.cidade"
            placeholder="Digite sua cidade"
            class="w-full"
          />
        </div>
        <div>
          <label for="cep" class="mb-2 block text-sm font-medium text-gray-700">
            CEP
          </label>
          <InputText
            id="cep"
            v-model="form.cep"
            placeholder="Digite seu CEP"
            class="w-full"
          />
        </div>
      </div>
    </div>

    <!-- Botões de navegação -->
    <div class="mt-8 flex gap-4">
      <button
        v-if="step > 1"
        @click="prevStep"
        class="flex-1 rounded-xl bg-gray-100 px-6 py-3 font-semibold text-gray-700 transition duration-200 hover:bg-gray-200"
      >
        Voltar
      </button>
      <button
        v-if="step < 3"
        @click="nextStep"
        class="bg-primary-color hover:bg-primary-dark-color flex-1 transform rounded-xl px-6 py-3 font-semibold text-white transition duration-200 ease-in-out hover:scale-[1.02]"
      >
        Próximo
      </button>
      <button
        v-if="step === 3"
        @click="submitForm"
        class="bg-primary-color hover:bg-primary-dark-color flex-1 transform rounded-xl px-6 py-3 font-semibold text-white transition duration-200 ease-in-out hover:scale-[1.02]"
      >
        Finalizar cadastro
      </button>
    </div>
  </div>
</template>
