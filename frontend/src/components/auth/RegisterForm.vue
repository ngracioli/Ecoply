<script setup lang="ts">
import InputText from "primevue/inputtext";
import InputMask from "primevue/inputmask";
import Password from "primevue/password";
import Dropdown from "primevue/dropdown";
import Button from "primevue/button";
import Toast from "primevue/toast";
import Message from "primevue/message";
import { useToast } from "primevue/usetoast";

import { ref, reactive } from "vue";
import { z } from "zod";
import { useStore } from "vuex";
import { useRouter } from "vue-router";
import type { RegisterRequest } from "../../types/requests/auth";
import api from "../../axios";
import type { StateConsult } from "../../types/responses/brazilStates";
import type { CCEEAgentsResponse, CCEEAgent } from "../../types/responses/ccee";
import { CCEE_ENDPOINTS, BRASILAPI_ENDPOINTS } from "../../api/endpoints";

const toast = useToast();
const store = useStore();
const router = useRouter();

async function loadStates() {
  try {
    const resp = await api.get<{ data: StateConsult }>(
      BRASILAPI_ENDPOINTS.STATES,
    );

    brazilStates.value = resp.data.data.sort(
      (a: { sigla: string }, b: { sigla: string }) =>
        a.sigla.localeCompare(b.sigla),
    );
  } catch (err) {}
}

loadStates();

const step = ref<number>(1);
const loading = ref(false);
const brazilStates = ref<StateConsult>([]);
const cceeAgents = ref<CCEEAgent[]>([]);

const form = reactive<RegisterRequest>({
  name: "",
  email: "",
  password: "",
  confirm_password: "",
  user_type: "buyer" as "buyer" | "supplier",
  address: {
    cep: "",
    state_initials: "",
    state: "",
    city: "",
    neighborhood: "",
    street: "",
    number: "",
    complement: "",
  },
  agent: {
    cnpj: "",
    ccee_code: "",
    submarket_name: "",
    company_name: "",
  },
});

const errors = reactive<Record<string, string>>({});

const clearFieldError = (field: string) => {
  if (errors[field]) {
    delete errors[field];
  }
};

const step1Schema = z
  .object({
    name: z.string().min(1, "Nome completo é obrigatório"),
    email: z.string().email("E-mail inválido"),
    password: z
      .string()
      .regex(/[^a-zA-Z0-9]/, "A senha deve conter ao menos um símbolo")
      .regex(/[0-9]/, "A senha deve conter ao menos um número")
      .regex(/[A-Z]/, "A senha deve conter ao menos uma letra maiúscula")
      .regex(/[a-z]/, "A senha deve conter ao menos uma letra minúscula")
      .max(50, "A senha pode ter no máximo 50 caracteres")
      .min(8, "A senha precisa ter ao menos 8 caracteres"),
    confirm_password: z.string().min(1, "Confirmação de senha é obrigatória"),
  })
  .refine((data) => data.password === data.confirm_password, {
    message: "Senhas não conferem",
    path: ["confirm_password"],
  });

const showSuccess = (detail: string) =>
  toast.add({ severity: "success", summary: "Sucesso", detail, life: 3000 });

function clearErrors() {
  Object.keys(errors).forEach((k) => delete errors[k]);
}

function validateStep(n: number): boolean {
  clearErrors();
  if (n === 1) {
    const res = step1Schema.safeParse({
      name: form.name,
      email: form.email,
      password: form.password,
      confirm_password: form.confirm_password,
    });
    if (!res.success) {
      res.error.issues.forEach(
        (i) => (errors[String(i.path[0] ?? "")] = i.message),
      );
      return false;
    }
    return true;
  }

  if (n === 2) {
    const digits = form.agent.cnpj.replace(/\D/g, "");
    if (!/^\d{14}$/.test(digits)) {
      errors["agent.cnpj"] = "CNPJ deve ter 14 dígitos.";
      return false;
    }
    return true;
  }

  if (n === 3) {
    if (!form.agent.company_name) {
      errors["agent.cnpj"] = "CNPJ não confirmado.";
      return false;
    }
    if (!form.agent.ccee_code) {
      errors["agent.ccee_code"] = "Selecione o código CCEE.";
      return false;
    }
    return true;
  }

  if (n === 4) {
    const cepDigits = form.address.cep.replace(/\D/g, "");
    if (!/^\d{8}$/.test(cepDigits)) {
      errors["address.cep"] = "CEP inválido.";
      return false;
    }
    if (!form.address.street || form.address.street.trim() === "") {
      errors["address.street"] = "Rua é obrigatória.";
      return false;
    }
    if (!form.address.neighborhood || form.address.neighborhood.trim() === "") {
      errors["address.neighborhood"] = "Bairro é obrigatório.";
      return false;
    }
    if (!form.address.number || form.address.number.trim() === "") {
      errors["address.number"] = "Número é obrigatório.";
      return false;
    }
    if (!form.address.city || form.address.city.trim() === "") {
      errors["address.city"] = "Cidade é obrigatória.";
      return false;
    }
    if (!form.address.state_initials) {
      errors["address.state_initials"] = "UF é obrigatória.";
      return false;
    }
    return true;
  }

  return true;
}

async function checkEmailAvailability(): Promise<boolean> {
  try {
    await api.get<void>("/api/v1/auth/available", {
      params: {
        type: "email",
        value: form.email,
      },
    });

    return true;
  } catch (err: any) {
    if (err.response?.status === 409) {
      errors["email"] = "Este e-mail já está cadastrado.";
      return false;
    }

    errors["email"] = "Erro ao verificar disponibilidade do e-mail.";
    return false;
  }
}

async function checkCNPJAvailability(): Promise<boolean> {
  try {
    await api.get<void>("/api/v1/auth/available", {
      params: {
        type: "cnpj",
        value: onlyDigits(form.agent.cnpj),
      },
    });

    return true;
  } catch (err: any) {
    if (err.response?.status === 409) {
      errors["agent.cnpj"] = "Este CNPJ já está cadastrado.";
      return false;
    }

    errors["agent.cnpj"] = "Erro ao verificar disponibilidade do CNPJ.";
    return false;
  }
}

async function next() {
  if (!validateStep(step.value)) return;

  if (step.value === 1) {
    const isEmailAvailable = await checkEmailAvailability();
    if (!isEmailAvailable) return;
  }

  step.value++;
}
function back() {
  if (step.value > 1) step.value--;
}

function onlyDigits(s: string) {
  return (s || "").replace(/\D/g, "");
}

async function confirmCnpj() {
  clearErrors();
  if (!validateStep(2)) return;

  const isCNPJAvailable = await checkCNPJAvailability();
  if (!isCNPJAvailable) return;

  loading.value = true;
  try {
    const cnpjDigits = onlyDigits(form.agent.cnpj);

    const respBrasil = await api.get<{
      data: { razao_social: string; nome_fantasia: string };
    }>(BRASILAPI_ENDPOINTS.CNPJ(cnpjDigits));

    const respCCEE = await api.get<CCEEAgentsResponse>(
      CCEE_ENDPOINTS.AGENTS(cnpjDigits),
    );

    if (
      !respCCEE.data ||
      !respCCEE.data.data ||
      respCCEE.data.data.length === 0
    ) {
      errors["agent.cnpj"] = "CNPJ não está vinculado a nenhum agente CCEE.";
      return;
    }

    form.agent.company_name =
      respBrasil.data.data.razao_social ||
      respBrasil.data.data.nome_fantasia ||
      "";

    cceeAgents.value = respCCEE.data.data;

    form.agent.ccee_code = "";
    form.agent.submarket_name = "";

    showSuccess(
      "Dados do CNPJ carregados. Verifique e selecione o agente CCEE.",
    );
    step.value = 3;
  } catch (err: any) {
    if (err.response?.status === 404) {
      if (err.config?.url?.includes("/brasilapi/cnpj/")) {
        errors["agent.cnpj"] = "CNPJ não encontrado.";
      } else {
        errors["agent.cnpj"] = "CNPJ não está vinculado a nenhum agente CCEE.";
      }
    } else {
      errors["agent.cnpj"] =
        err.response?.data?.message ||
        "Falha ao consultar CNPJ. Tente novamente.";
    }
  } finally {
    loading.value = false;
  }
}

async function fetchCep() {
  clearErrors();
  const cepDigits = onlyDigits(form.address.cep);
  if (!/^\d{8}$/.test(cepDigits)) {
    errors["address.cep"] = "CEP inválido.";
    return;
  }

  loading.value = true;
  try {
    const resp = await api.get<{
      data: {
        street: string;
        neighborhood: string;
        city: string;
        state: string;
      };
    }>(BRASILAPI_ENDPOINTS.CEP(cepDigits));

    const data = resp.data.data;

    form.address.street = data.street || "";
    form.address.neighborhood = data.neighborhood || "";
    form.address.city = data.city || "";
    form.address.state_initials = data.state || "";
    const estadoEncontrado = brazilStates.value.find(
      (e) => e.sigla === data.state,
    );
    if (estadoEncontrado) {
      form.address.state = estadoEncontrado.nome;
    }
    showSuccess("Endereço preenchido a partir do CEP.");
  } catch (err: any) {
    if (err.response?.status === 404) {
      errors["address.cep"] = "CEP não encontrado.";
    } else {
      errors["address.cep"] =
        err.response?.data?.message || "Erro ao consultar CEP.";
    }
  } finally {
    loading.value = false;
  }
}

function onStateChange() {
  const estadoSelecionado = brazilStates.value.find(
    (e) => e.sigla === form.address.state_initials,
  );
  if (estadoSelecionado) {
    form.address.state = estadoSelecionado.nome;
  } else {
    form.address.state = "";
  }
}

function onCCEEAgentChange() {
  const agenteEscolhido = cceeAgents.value.find(
    (a) => a.value === form.agent.ccee_code,
  );
  if (agenteEscolhido) {
    form.agent.submarket_name = agenteEscolhido.submarket;
  } else {
    form.agent.submarket_name = "";
  }
}

async function submitFinal() {
  if (!validateStep(4)) return;

  loading.value = true;
  try {
    const payload = {
      name: form.name,
      email: form.email,
      password: form.password,
      confirm_password: form.confirm_password,
      user_type: form.user_type,
      address: {
        ...form.address,
        cep: onlyDigits(form.address.cep),
      },
      agent: {
        ...form.agent,
        cnpj: onlyDigits(form.agent.cnpj),
        submarket_name: form.agent.submarket_name || null,
      },
    };

    await store.dispatch("auth/register", payload);

    showSuccess("Cadastro realizado com sucesso!");
    router.push({ name: "Dashboard" });
  } catch (err: any) {
    errors["submit"] =
      err.response?.data?.message || "Erro ao finalizar cadastro.";
  } finally {
    loading.value = false;
  }
}
</script>

<template>
  <div class="flex flex-col gap-4 px-4 sm:gap-6 sm:px-0">
    <div class="mb-2 text-center sm:mb-4">
      <div
        class="mb-3 inline-flex h-12 w-12 items-center justify-center rounded-full bg-gradient-to-br from-emerald-400 to-emerald-600 shadow-lg sm:mb-4 sm:h-16 sm:w-16"
      >
        <i class="pi pi-user-plus !text-2xl text-white sm:!text-3xl"></i>
      </div>
      <h1
        class="mb-1 px-4 text-2xl font-bold text-gray-900 sm:mb-2 sm:px-0 sm:text-3xl"
      >
        Crie sua conta
      </h1>
      <p class="px-2 text-sm text-gray-600 sm:px-0 sm:text-base">
        Preencha os passos para começar
      </p>
    </div>

    <div class="mb-2 flex items-center justify-center gap-2 sm:mb-4">
      <div
        v-for="i in 4"
        :key="i"
        :class="[
          'h-2 rounded-full transition-all duration-300',
          i === step
            ? 'w-10 bg-emerald-600 shadow-md sm:w-12'
            : 'w-2 bg-gray-400',
        ]"
      ></div>
    </div>

    <form
      @submit.prevent="step < 4 ? next() : submitFinal()"
      class="mobile-scrollable flex flex-col gap-4 sm:gap-5"
    >
      <Toast />

      <div v-show="step === 1" class="flex flex-col gap-4 sm:gap-5">
        <div class="flex flex-col gap-2">
          <label for="name" class="text-sm font-semibold text-gray-800"
            >Nome completo</label
          >
          <InputText
            id="name"
            v-model="form.name"
            @input="clearFieldError('name')"
            :disabled="loading"
            :invalid="!!errors.name"
            placeholder="Digite seu nome completo"
            class="w-full touch-manipulation !rounded-xl !border-2 !border-gray-300 !bg-white !py-2.5 !text-base transition-colors hover:!border-gray-400 focus:!border-emerald-500 sm:!py-3"
            size="large"
          />
          <small
            v-if="errors.name"
            class="text-xs font-medium text-red-500 sm:text-sm"
            >{{ errors.name }}</small
          >
        </div>

        <div class="flex flex-col gap-2">
          <label for="email" class="text-sm font-semibold text-gray-800"
            >E-mail</label
          >
          <InputText
            id="email"
            v-model="form.email"
            @input="clearFieldError('email')"
            :disabled="loading"
            :invalid="!!errors.email"
            placeholder="seu@email.com"
            class="w-full touch-manipulation !rounded-xl !border-2 !border-gray-300 !bg-white !py-2.5 !text-base transition-colors hover:!border-gray-400 focus:!border-emerald-500 sm:!py-3"
            size="large"
          />
          <small
            v-if="errors.email"
            class="text-xs font-medium text-red-500 sm:text-sm"
            >{{ errors.email }}</small
          >
        </div>

        <div class="flex flex-col gap-2">
          <label for="password" class="text-sm font-semibold text-gray-800"
            >Senha</label
          >
          <Password
            id="password"
            v-model="form.password"
            @input="clearFieldError('password')"
            :disabled="loading"
            :invalid="!!errors.password"
            :feedback="false"
            toggleMask
            placeholder="••••••••"
            class="w-full"
            inputClass="w-full !rounded-xl !border-2 !border-gray-300 focus:!border-emerald-500 !py-2.5 sm:!py-3 !bg-white hover:!border-gray-400 transition-colors !text-base touch-manipulation"
            size="large"
          />
          <small
            v-if="errors.password"
            class="text-xs font-medium text-red-500 sm:text-sm"
            >{{ errors.password }}</small
          >
          <small v-else class="text-xs text-gray-600">
            Mín. 8 caracteres, com letras maiúsculas, minúsculas, números e
            símbolos.
          </small>
        </div>

        <div class="flex flex-col gap-2">
          <label
            for="confirm_password"
            class="text-sm font-semibold text-gray-800"
            >Confirmar senha</label
          >
          <Password
            id="confirm_password"
            v-model="form.confirm_password"
            @input="clearFieldError('confirm_password')"
            :disabled="loading"
            :invalid="!!errors.confirm_password"
            :feedback="false"
            toggleMask
            placeholder="••••••••"
            class="w-full"
            inputClass="w-full !rounded-xl !border-2 !border-gray-300 focus:!border-emerald-500 !py-2.5 sm:!py-3 !bg-white hover:!border-gray-400 transition-colors !text-base touch-manipulation"
            size="large"
          />
          <small
            v-if="errors.confirm_password"
            class="text-xs font-medium text-red-500 sm:text-sm"
            >{{ errors.confirm_password }}</small
          >
        </div>

        <Button
          type="submit"
          label="Próximo"
          icon="pi pi-arrow-right"
          iconPos="right"
          :loading="loading"
          :disabled="loading"
          class="mt-2 min-h-[44px] w-full touch-manipulation !rounded-xl !border-none !bg-gradient-to-r !from-emerald-600 !to-emerald-700 !py-3 !text-base !font-semibold shadow-lg transition-all duration-300 hover:!from-emerald-700 hover:!to-emerald-800 hover:shadow-xl active:scale-[0.98] sm:mt-4"
          size="large"
        />
      </div>

      <div v-show="step === 2" class="flex flex-col gap-4 sm:gap-5">
        <Message
          severity="info"
          :closable="false"
          icon="pi pi-exclamation-circle"
          class="!rounded-xl text-sm"
        >
          O CNPJ do agente é obrigatório para todos os usuários.
        </Message>

        <div class="flex flex-col gap-2">
          <label for="user_type" class="text-sm font-semibold text-gray-700"
            >Tipo de usuário</label
          >
          <Dropdown
            id="user_type"
            :options="[
              { label: 'Comprador', value: 'buyer' },
              { label: 'Fornecedor', value: 'supplier' },
            ]"
            optionLabel="label"
            optionValue="value"
            v-model="form.user_type"
            :disabled="loading"
            placeholder="Selecione o tipo"
            class="w-full touch-manipulation !rounded-xl"
            size="large"
          />
        </div>

        <div class="flex flex-col gap-2">
          <label for="cnpj" class="text-sm font-semibold text-gray-700"
            >CNPJ do Agente CCEE</label
          >
          <InputMask
            id="cnpj"
            v-model="form.agent.cnpj"
            @input="clearFieldError('agent.cnpj')"
            mask="99.999.999/9999-99"
            :disabled="loading"
            :invalid="!!errors['agent.cnpj']"
            placeholder="00.000.000/0000-00"
            class="w-full touch-manipulation !rounded-xl !border-2 !border-gray-200 !py-2.5 !text-base focus:!border-emerald-500 sm:!py-3"
            size="large"
          />
          <small
            v-if="errors['agent.cnpj']"
            class="text-xs font-medium text-red-500 sm:text-sm"
            >{{ errors["agent.cnpj"] }}</small
          >
        </div>

        <div class="mt-2 flex flex-col gap-3 sm:flex-row">
          <Button
            label="Voltar"
            icon="pi pi-arrow-left"
            severity="secondary"
            size="large"
            class="min-h-[44px] flex-1 touch-manipulation !rounded-xl !py-3"
            @click.prevent="back"
          />
          <Button
            label="Confirmar CNPJ"
            icon="pi pi-search"
            iconPos="right"
            :loading="loading"
            :disabled="loading"
            class="min-h-[44px] flex-1 touch-manipulation !rounded-xl !border-none !bg-emerald-600 !py-3 !font-semibold transition-all duration-300 hover:!bg-emerald-700 active:scale-[0.98]"
            size="large"
            @click.prevent="confirmCnpj"
          />
        </div>
      </div>

      <div v-show="step === 3" class="flex flex-col gap-4 sm:gap-5">
        <Message
          severity="info"
          :closable="false"
          icon="pi pi-exclamation-circle"
          class="!rounded-xl text-sm"
        >
          Confirme os dados retornados pelo CNPJ e selecione o código CCEE.
        </Message>

        <div class="flex flex-col gap-2">
          <label for="company_name" class="text-sm font-semibold text-gray-700"
            >Razão Social / Nome</label
          >
          <InputText
            id="company_name"
            v-model="form.agent.company_name"
            :disabled="true"
            class="w-full !rounded-xl !bg-gray-50 !py-2.5 !text-base sm:!py-3"
            size="large"
          />
        </div>

        <div class="flex flex-col gap-2">
          <label for="cnpj_confirm" class="text-sm font-semibold text-gray-700"
            >CNPJ</label
          >
          <InputText
            id="cnpj_confirm"
            v-model="form.agent.cnpj"
            :disabled="true"
            class="w-full !rounded-xl !bg-gray-50 !py-2.5 !text-base sm:!py-3"
            size="large"
          />
        </div>

        <div class="flex flex-col gap-2">
          <label for="ccee_code" class="text-sm font-semibold text-gray-700"
            >Código de Perfil do Agente CCEE</label
          >
          <Dropdown
            id="ccee_code"
            v-model="form.agent.ccee_code"
            :options="cceeAgents"
            optionLabel="label"
            optionValue="value"
            :disabled="loading"
            :invalid="!!errors['agent.ccee_code']"
            placeholder="Selecione o código CCEE"
            class="w-full touch-manipulation !rounded-xl"
            size="large"
            @change="
              () => {
                onCCEEAgentChange();
                clearFieldError('agent.ccee_code');
              }
            "
          />
          <small
            v-if="errors['agent.ccee_code']"
            class="text-xs font-medium text-red-500 sm:text-sm"
            >{{ errors["agent.ccee_code"] }}</small
          >
        </div>

        <div class="flex flex-col gap-2">
          <label for="submarket" class="text-sm font-semibold text-gray-700"
            >Submercado Energético</label
          >
          <InputText
            id="submarket"
            v-model="form.agent.submarket_name"
            :disabled="true"
            placeholder="Agente não selecionado"
            class="w-full !rounded-xl !bg-gray-50 !py-2.5 !text-base sm:!py-3"
            size="large"
          />
          <small class="text-xs text-gray-500"
            >Será preenchido automaticamente ao selecionar o código CCEE.</small
          >
        </div>

        <div class="mt-2 flex flex-col gap-3 sm:flex-row">
          <Button
            label="Voltar"
            icon="pi pi-arrow-left"
            severity="secondary"
            size="large"
            class="min-h-[44px] flex-1 touch-manipulation !rounded-xl !py-3"
            @click.prevent="back"
          />
          <Button
            label="Confirmar"
            icon="pi pi-check"
            iconPos="right"
            class="min-h-[44px] flex-1 touch-manipulation !rounded-xl !border-none !bg-emerald-600 !py-3 !font-semibold transition-all duration-300 hover:!bg-emerald-700 active:scale-[0.98]"
            size="large"
            @click.prevent="next"
          />
        </div>
      </div>

      <div v-show="step === 4" class="flex flex-col gap-4 sm:gap-5">
        <div class="flex flex-col gap-2">
          <label for="cep" class="text-sm font-semibold text-gray-700"
            >CEP</label
          >
          <div class="flex gap-2">
            <InputMask
              id="cep"
              v-model="form.address.cep"
              @input="clearFieldError('address.cep')"
              mask="99999-999"
              :disabled="loading"
              :invalid="!!errors['address.cep']"
              placeholder="00000-000"
              class="flex-1 touch-manipulation !rounded-xl !border-2 !border-gray-200 !py-2.5 !text-base focus:!border-emerald-500 sm:!py-3"
              size="large"
            />
            <Button
              label="Buscar"
              icon="pi pi-search"
              :loading="loading"
              :disabled="loading"
              size="large"
              class="min-h-[44px] touch-manipulation !rounded-xl !border-none !bg-emerald-600 !px-4 hover:!bg-emerald-700 active:scale-[0.98] sm:!px-6"
              @click.prevent="fetchCep"
            />
          </div>
          <small
            v-if="errors['address.cep']"
            class="text-xs font-medium text-red-500 sm:text-sm"
            >{{ errors["address.cep"] }}</small
          >
        </div>

        <div class="flex flex-col gap-2">
          <label for="street" class="text-sm font-semibold text-gray-700"
            >Rua</label
          >
          <InputText
            id="street"
            v-model="form.address.street"
            @input="clearFieldError('address.street')"
            :disabled="loading"
            :invalid="!!errors['address.street']"
            placeholder="Nome da rua"
            class="w-full touch-manipulation !rounded-xl !border-2 !border-gray-200 !py-2.5 !text-base focus:!border-emerald-500 sm:!py-3"
            size="large"
          />
          <small
            v-if="errors['address.street']"
            class="text-xs font-medium text-red-500 sm:text-sm"
            >{{ errors["address.street"] }}</small
          >
        </div>

        <div class="flex flex-col gap-3 sm:flex-row">
          <div class="flex flex-1 flex-col gap-2">
            <label
              for="neighborhood"
              class="text-sm font-semibold text-gray-700"
              >Bairro</label
            >
            <InputText
              id="neighborhood"
              v-model="form.address.neighborhood"
              @input="clearFieldError('address.neighborhood')"
              :disabled="loading"
              :invalid="!!errors['address.neighborhood']"
              placeholder="Bairro"
              class="touch-manipulation !rounded-xl !border-2 !border-gray-200 !py-2.5 !text-base focus:!border-emerald-500 sm:!py-3"
              size="large"
            />
            <small
              v-if="errors['address.neighborhood']"
              class="text-xs font-medium text-red-500 sm:text-sm"
              >{{ errors["address.neighborhood"] }}</small
            >
          </div>
          <div class="flex flex-1 flex-col gap-2 sm:w-32">
            <label for="number" class="text-sm font-semibold text-gray-700"
              >Número</label
            >
            <InputText
              id="number"
              v-model="form.address.number"
              @input="clearFieldError('address.number')"
              :disabled="loading"
              :invalid="!!errors['address.number']"
              placeholder="Nº"
              class="touch-manipulation !rounded-xl !border-2 !border-gray-200 !py-2.5 !text-base focus:!border-emerald-500 sm:!py-3"
              size="large"
            />
            <small
              v-if="errors['address.number']"
              class="text-xs font-medium text-red-500 sm:text-sm"
              >{{ errors["address.number"] }}</small
            >
          </div>
        </div>

        <div class="flex flex-col gap-3 sm:flex-row">
          <div class="flex flex-1 flex-col gap-2">
            <label for="city" class="text-sm font-semibold text-gray-700"
              >Cidade</label
            >
            <InputText
              id="city"
              v-model="form.address.city"
              @input="clearFieldError('address.city')"
              :disabled="loading"
              :invalid="!!errors['address.city']"
              placeholder="Cidade"
              class="touch-manipulation !rounded-xl !border-2 !border-gray-200 !py-2.5 !text-base focus:!border-emerald-500 sm:!py-3"
              size="large"
            />
            <small
              v-if="errors['address.city']"
              class="text-xs font-medium text-red-500 sm:text-sm"
              >{{ errors["address.city"] }}</small
            >
          </div>
          <div class="flex flex-1 flex-col gap-2 sm:w-32">
            <label for="state" class="text-sm font-semibold text-gray-700"
              >Estado</label
            >
            <Dropdown
              id="state"
              v-model="form.address.state_initials"
              @change="
                () => {
                  onStateChange();
                  clearFieldError('address.state_initials');
                }
              "
              :options="brazilStates"
              optionLabel="sigla"
              optionValue="sigla"
              :disabled="loading"
              :invalid="!!errors['address.state_initials']"
              placeholder="UF"
              class="touch-manipulation !rounded-xl"
              size="large"
              :showClear="true"
            />
            <small
              v-if="errors['address.state_initials']"
              class="text-xs font-medium text-red-500 sm:text-sm"
              >{{ errors["address.state_initials"] }}</small
            >
          </div>
        </div>

        <div class="flex flex-col gap-2">
          <label for="complement" class="text-sm font-semibold text-gray-700"
            >Complemento (opcional)</label
          >
          <InputText
            id="complement"
            v-model="form.address.complement"
            :disabled="loading"
            placeholder="Apto, bloco, etc"
            class="w-full touch-manipulation !rounded-xl !border-2 !border-gray-200 !py-2.5 !text-base focus:!border-emerald-500 sm:!py-3"
            size="large"
          />
        </div>

        <div class="mt-2 flex flex-col gap-3 sm:flex-row">
          <Button
            label="Voltar"
            icon="pi pi-arrow-left"
            severity="secondary"
            size="large"
            class="min-h-[44px] flex-1 touch-manipulation !rounded-xl !py-3"
            @click.prevent="back"
          />
          <Button
            type="submit"
            label="Finalizar Cadastro"
            icon="pi pi-check"
            iconPos="right"
            :loading="loading"
            :disabled="loading"
            class="min-h-[44px] flex-1 touch-manipulation !rounded-xl !border-none !bg-emerald-600 !py-3 !font-semibold shadow-lg transition-all duration-300 hover:!bg-emerald-700 hover:shadow-xl active:scale-[0.98]"
            size="large"
          />
        </div>
      </div>
    </form>

    <div class="mt-2 text-center sm:mt-4">
      <p class="px-2 text-sm text-gray-600 sm:px-0 sm:text-base">
        Já tem uma conta?
        <router-link
          to="/login"
          class="-my-3 ml-1 inline-block min-h-[44px] touch-manipulation leading-[44px] font-semibold text-emerald-600 transition-colors duration-200 hover:text-emerald-700 active:text-emerald-800"
        >
          Entrar
        </router-link>
      </p>
    </div>
  </div>
</template>

<style scoped>
:deep(.p-message-text) {
  padding-left: 0.5rem;
}
</style>
