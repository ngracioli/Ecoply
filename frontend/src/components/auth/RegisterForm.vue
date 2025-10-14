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

const toast = useToast();
const store = useStore();
const router = useRouter();

async function loadEstados() {
  try {
    const resp = await fetch("https://brasilapi.com.br/api/ibge/uf/v1");
    if (!resp.ok) return;
    const data = await resp.json();
    estados.value = data
      .map((uf: any) => ({
        sigla: uf.sigla,
        nome: uf.nome,
      }))
      .sort((a: { sigla: string }, b: { sigla: string }) =>
        a.sigla.localeCompare(b.sigla),
      );
  } catch (err) {
    console.error("Erro ao carregar estados:", err);
  }
}

// Carrega os estados quando o componente é montado
loadEstados();

/* --- UI / flow --- */
const step = ref<number>(1); // 1..4
const loading = ref(false);
const estados = ref<Array<{ sigla: string; nome: string }>>([]);

/* --- Form model --- */
const form = reactive({
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

/* --- Errors por campo --- */
const errors = reactive<Record<string, string>>({});

/* --- Schemas zod por step --- */
const step1Schema = z
  .object({
    name: z.string().min(1, "Nome completo é obrigatório"),
    email: z.string().email("E-mail inválido"),
    password: z.string().min(6, "A senha precisa ter ao menos 6 caracteres"),
    confirm_password: z.string(),
  })
  .refine((data) => data.password === data.confirm_password, {
    message: "Senhas não conferem",
    path: ["confirm_password"],
  });

const step2Schema = z.object({
  user_type: z.enum(["buyer", "supplier"]),
  // cnpj opcional quando buyer, obrigatório quando supplier
  cnpj: z.string().optional(),
});

const step3Schema = z.object({
  // no step 3 validar se existe agent.company_name (preenchido pela API)
  agent_company_name: z.string().min(1, "CNPJ não confirmado"),
});

const step4Schema = z.object({
  "address.cep": z.string().min(8, "CEP inválido"),
  "address.city": z.string().min(1, "Cidade é obrigatória"),
  "address.state_initials": z.string().min(2, "UF é obrigatória"),
});

/* --- Helpers UI --- */
const steps = [
  { id: 1, label: "Dados pessoais" },
  { id: 2, label: "Tipo / CNPJ" },
  { id: 3, label: "Confirma CNPJ" },
  { id: 4, label: "Endereço" },
];

const showError = (detail: string) =>
  toast.add({ severity: "error", summary: "Erro", detail, life: 3500 });
const showSuccess = (detail: string) =>
  toast.add({ severity: "success", summary: "Sucesso", detail, life: 3000 });

function clearErrors() {
  Object.keys(errors).forEach((k) => delete errors[k]);
}

/* --- valida step atual --- */
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
    // Exige CNPJ numérico com 14 dígitos para ambos (comprador e fornecedor)
    const digits = form.agent.cnpj.replace(/\D/g, "");
    if (!/^\d{14}$/.test(digits)) {
      errors["agent.cnpj"] = "CNPJ deve ter 14 dígitos";
      return false;
    }
    return true;
  }

  if (n === 3) {
    if (!form.agent.company_name) {
      errors["agent.cnpj"] = "CNPJ não confirmado";
      return false;
    }
    // Valida o submercado energético
    if (!form.agent.submarket_name) {
      errors["agent.submarket_name"] = "Selecione o submercado energético";
      return false;
    }
    // Valida o código CCEE (1 a 6 dígitos numéricos)
    const cceeDigits = form.agent.ccee_code.replace(/\D/g, "");
    if (!cceeDigits || !/^\d{1,6}$/.test(cceeDigits)) {
      errors["agent.ccee_code"] = "Digite de 1 a 6 dígitos";
      return false;
    }
    return true;
  }

  if (n === 4) {
    const cepDigits = form.address.cep.replace(/\D/g, "");
    if (!/^\d{8}$/.test(cepDigits)) {
      errors["address.cep"] = "CEP inválido";
      return false;
    }
    if (!form.address.city) {
      errors["address.city"] = "Cidade obrigatória";
      return false;
    }
    if (!form.address.state_initials) {
      errors["address.state_initials"] = "UF obrigatória";
      return false;
    }
    return true;
  }

  return true;
}

/* --- navegação --- */
function next() {
  if (!validateStep(step.value)) return;
  step.value++;
}
function back() {
  if (step.value > 1) step.value--;
}

/* --- Utilidades para API calls --- */
function onlyDigits(s: string) {
  return (s || "").replace(/\D/g, "");
}

/* --- Consulta CNPJ (Confirmar) --- */
/* Usando BrasilAPI para consultar CNPJ */
async function confirmCnpj() {
  clearErrors();
  if (!validateStep(2)) return;

  loading.value = true;
  try {
    const cnpjDigits = onlyDigits(form.agent.cnpj);
    // BrasilAPI endpoint: https://brasilapi.com.br/api/cnpj/v1/{cnpj}
    const resp = await fetch(
      `https://brasilapi.com.br/api/cnpj/v1/${cnpjDigits}`,
    );
    if (!resp.ok) {
      const errorData = await resp.json().catch(() => ({}));
      errors["agent.cnpj"] = errorData.message || "Erro ao consultar CNPJ";
      return;
    }
    const data = await resp.json();

    // BrasilAPI retorna campos como 'razao_social', 'nome_fantasia', 'cnae_fiscal', etc.
    form.agent.company_name = data.razao_social || data.nome_fantasia || "";
    // Formata o CNPJ com máscara para exibir no input
    const cnpjFormatted = (data.cnpj || cnpjDigits).replace(
      /^(\d{2})(\d{3})(\d{3})(\d{4})(\d{2})$/,
      "$1.$2.$3/$4-$5",
    );
    form.agent.cnpj = cnpjFormatted;
    // Não preenche mais ccee_code e submarket_name automaticamente

    showSuccess("Dados do CNPJ carregados. Verifique e prossiga.");
    step.value = 3;
  } catch (err: any) {
    console.error(err);
    errors["agent.cnpj"] = "Falha ao consultar CNPJ. Tente novamente.";
  } finally {
    loading.value = false;
  }
}

/* --- Consulta CEP (BrasilAPI) --- */
async function fetchCep() {
  clearErrors();
  const cepDigits = onlyDigits(form.address.cep);
  if (!/^\d{8}$/.test(cepDigits)) {
    errors["address.cep"] = "CEP inválido";
    return;
  }

  loading.value = true;
  try {
    const resp = await fetch(
      `https://brasilapi.com.br/api/cep/v1/${cepDigits}`,
    );
    if (!resp.ok) {
      const errorData = await resp.json().catch(() => ({}));
      errors["address.cep"] = errorData.message || "CEP não encontrado";
      return;
    }
    const data = await resp.json();

    // popula o address (padrões BrasilAPI -> street, neighborhood, city, state)
    form.address.street = data.street || "";
    form.address.neighborhood = data.neighborhood || "";
    form.address.city = data.city || "";
    form.address.state_initials = data.state || "";
    // Também preenche o nome do estado quando busca pelo CEP
    const estadoEncontrado = estados.value.find((e) => e.sigla === data.state);
    if (estadoEncontrado) {
      form.address.state = estadoEncontrado.nome;
    }
    showSuccess("Endereço preenchido a partir do CEP.");
  } catch (err: any) {
    console.error(err);
    errors["address.cep"] = "Erro ao consultar CEP.";
  } finally {
    loading.value = false;
  }
}

/* --- Atualiza o nome do estado quando seleciona a UF --- */
function onStateChange() {
  const estadoSelecionado = estados.value.find(
    (e) => e.sigla === form.address.state_initials,
  );
  if (estadoSelecionado) {
    form.address.state = estadoSelecionado.nome;
  } else {
    form.address.state = "";
  }
}

/* --- submit final --- */
async function submitFinal() {
  if (!validateStep(4)) return;

  loading.value = true;
  try {
    // prepare payload conforme seu backend
    const payload = {
      name: form.name,
      email: form.email,
      password: form.password,
      confirm_password: form.confirm_password,
      user_type: form.user_type === "supplier" ? "supplier" : "buyer",
      address: {
        ...form.address,
        cep: onlyDigits(form.address.cep),
      },
      agent: {
        ...form.agent,
        // Remove a máscara do CNPJ (envia só números para o backend)
        cnpj: onlyDigits(form.agent.cnpj),
      },
    };

    // ajuste a action conforme sua store
    await store.dispatch("auth/register", payload);

    showSuccess("Cadastro realizado com sucesso!");
    router.push({ name: "Dashboard" });
  } catch (err: any) {
    console.error(err);
    errors["submit"] =
      err.response?.data?.message || "Erro ao finalizar cadastro.";
  } finally {
    loading.value = false;
  }
}
</script>

<template>
  <div class="rounded-2xl bg-white p-8 shadow-xl">
    <div class="mb-8 text-center">
      <div
        class="bg-primary-color/10 mb-4 inline-flex h-16 w-16 items-center justify-center rounded-full"
      >
        <i class="pi pi-user-plus text-primary-color !text-3xl"></i>
      </div>
      <h1 class="mb-2 text-2xl font-bold text-gray-900">
        Criar conta no Ecoply
      </h1>
      <p class="text-gray-600">Preencha os passos para criar sua conta</p>
    </div>

    <form
      @submit.prevent="step < 4 ? next() : submitFinal()"
      class="flex flex-col gap-5"
    >
      <Toast />

      <!-- STEP 1: Dados Pessoais -->
      <div v-show="step === 1" class="flex flex-col gap-5">
        <div class="flex flex-col gap-2">
          <label for="name" class="font-medium text-gray-700"
            >Nome completo</label
          >
          <InputText
            id="name"
            v-model="form.name"
            :disabled="loading"
            :invalid="!!errors.name"
            placeholder="Digite seu nome completo"
            class="w-full"
            size="large"
          />
          <small v-if="errors.name" class="text-red-600">{{
            errors.name
          }}</small>
        </div>

        <div class="flex flex-col gap-2">
          <label for="email" class="font-medium text-gray-700">E-mail</label>
          <InputText
            id="email"
            v-model="form.email"
            :disabled="loading"
            :invalid="!!errors.email"
            placeholder="exemplo@dominio.com"
            class="w-full"
            size="large"
          />
          <small v-if="errors.email" class="text-red-600">{{
            errors.email
          }}</small>
        </div>

        <div class="flex flex-col gap-2">
          <label for="password" class="font-medium text-gray-700">Senha</label>
          <Password
            id="password"
            v-model="form.password"
            :disabled="loading"
            :invalid="!!errors.password"
            :feedback="false"
            toggleMask
            placeholder="Digite sua senha"
            class="w-full"
            inputClass="w-full"
            size="large"
          />
          <small v-if="errors.password" class="text-red-600">{{
            errors.password
          }}</small>
        </div>

        <div class="flex flex-col gap-2">
          <label for="confirm_password" class="font-medium text-gray-700"
            >Confirmar senha</label
          >
          <Password
            id="confirm_password"
            v-model="form.confirm_password"
            :disabled="loading"
            :invalid="!!errors.confirm_password"
            :feedback="false"
            toggleMask
            placeholder="Confirme sua senha"
            class="w-full"
            inputClass="w-full"
            size="large"
          />
          <small v-if="errors.confirm_password" class="text-red-600">{{
            errors.confirm_password
          }}</small>
        </div>

        <Button
          type="submit"
          label="Próximo"
          icon="pi pi-arrow-right"
          iconPos="right"
          :loading="loading"
          :disabled="loading"
          severity="primary"
          size="large"
          class="mt-2 w-full"
        />
      </div>

      <!-- STEP 2: Tipo e CNPJ -->
      <div v-show="step === 2" class="flex flex-col gap-5">
        <Message
          severity="info"
          :closable="false"
          icon="pi pi-exclamation-circle"
        >
          O CNPJ é obrigatório para todos os usuários.
        </Message>

        <div class="flex flex-col gap-2">
          <label for="user_type" class="font-medium text-gray-700"
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
            class="w-full"
            size="large"
          />
        </div>

        <div class="flex flex-col gap-2">
          <label for="cnpj" class="font-medium text-gray-700">CNPJ</label>
          <InputMask
            id="cnpj"
            v-model="form.agent.cnpj"
            mask="99.999.999/9999-99"
            :disabled="loading"
            :invalid="!!errors['agent.cnpj']"
            placeholder="00.000.000/0000-00"
            class="w-full"
            size="large"
          />
          <small v-if="errors['agent.cnpj']" class="text-red-600">{{
            errors["agent.cnpj"]
          }}</small>
        </div>

        <div class="mt-2 flex gap-3">
          <Button
            label="Voltar"
            icon="pi pi-arrow-left"
            severity="secondary"
            size="large"
            class="flex-1"
            @click.prevent="back"
          />
          <Button
            label="Confirmar CNPJ"
            icon="pi pi-search"
            iconPos="right"
            :loading="loading"
            :disabled="loading"
            severity="primary"
            size="large"
            class="flex-1"
            @click.prevent="confirmCnpj"
          />
        </div>
      </div>

      <!-- STEP 3: Confirma CNPJ (readonly fields) -->
      <div v-show="step === 3" class="flex flex-col gap-5">
        <Message
          severity="info"
          :closable="false"
          icon="pi pi-exclamation-circle"
        >
          Confirme os dados retornados pelo CNPJ.
        </Message>

        <div class="flex flex-col gap-2">
          <label for="company_name" class="font-medium text-gray-700"
            >Razão Social / Nome</label
          >
          <InputText
            id="company_name"
            v-model="form.agent.company_name"
            :disabled="true"
            class="w-full"
            size="large"
          />
        </div>

        <div class="flex flex-col gap-2">
          <label for="cnpj_confirm" class="font-medium text-gray-700"
            >CNPJ</label
          >
          <InputText
            id="cnpj_confirm"
            v-model="form.agent.cnpj"
            :disabled="true"
            class="w-full"
            size="large"
          />
        </div>

        <div class="flex flex-col gap-2">
          <label for="submarket" class="font-medium text-gray-700"
            >Submercado Energético</label
          >
          <Dropdown
            id="submarket"
            v-model="form.agent.submarket_name"
            :options="[
              { label: 'N - Norte', value: 'N' },
              { label: 'NE - Nordeste', value: 'NE' },
              { label: 'S - Sul', value: 'S' },
              { label: 'SE/CO - Sudeste/Centro-Oeste', value: 'SE_CO' },
            ]"
            optionLabel="label"
            optionValue="value"
            :disabled="loading"
            :invalid="!!errors['agent.submarket_name']"
            placeholder="Selecione o submercado"
            class="w-full"
            size="large"
          />
          <small v-if="errors['agent.submarket_name']" class="text-red-600">{{
            errors["agent.submarket_name"]
          }}</small>
        </div>

        <div class="flex flex-col gap-2">
          <label for="ccee_code" class="font-medium text-gray-700"
            >Código de Perfil do Agente CCEE</label
          >
          <InputText
            id="ccee_code"
            v-model="form.agent.ccee_code"
            :disabled="loading"
            :invalid="!!errors['agent.ccee_code']"
            placeholder="Ex: 12345"
            class="w-full"
            size="large"
          />
          <small
            :class="
              errors['agent.ccee_code'] ? 'text-red-600' : 'text-gray-600'
            "
            >Digite apenas números inteiros, de 1 a 6 dígitos.</small
          >
        </div>

        <div class="mt-2 flex gap-3">
          <Button
            label="Voltar"
            icon="pi pi-arrow-left"
            severity="secondary"
            size="large"
            class="flex-1"
            @click.prevent="back"
          />
          <Button
            label="Confirmar"
            icon="pi pi-check"
            iconPos="right"
            severity="primary"
            size="large"
            class="flex-1"
            @click.prevent="next"
          />
        </div>
      </div>

      <!-- STEP 4: Endereço via CEP -->
      <div v-show="step === 4" class="flex flex-col gap-5">
        <div class="flex flex-col gap-2">
          <label for="cep" class="font-medium text-gray-700">CEP</label>
          <div class="flex gap-2">
            <InputMask
              id="cep"
              v-model="form.address.cep"
              mask="99999-999"
              :disabled="loading"
              :invalid="!!errors['address.cep']"
              placeholder="00000-000"
              class="flex-1"
              size="large"
            />
            <Button
              label="Buscar"
              icon="pi pi-search"
              :loading="loading"
              :disabled="loading"
              size="large"
              @click.prevent="fetchCep"
            />
          </div>
          <small v-if="errors['address.cep']" class="text-red-600">{{
            errors["address.cep"]
          }}</small>
        </div>

        <div class="flex flex-col gap-2">
          <label for="street" class="font-medium text-gray-700">Rua</label>
          <InputText
            id="street"
            v-model="form.address.street"
            :disabled="loading"
            placeholder="Nome da rua"
            class="w-full"
            size="large"
          />
        </div>

        <div class="flex gap-3">
          <div class="flex flex-1 flex-col gap-2">
            <label for="neighborhood" class="font-medium text-gray-700"
              >Bairro</label
            >
            <InputText
              id="neighborhood"
              v-model="form.address.neighborhood"
              :disabled="loading"
              placeholder="Bairro"
              size="large"
            />
          </div>
          <div class="flex w-32 flex-col gap-2">
            <label for="number" class="font-medium text-gray-700">Número</label>
            <InputText
              id="number"
              v-model="form.address.number"
              :disabled="loading"
              placeholder="Nº"
              size="large"
            />
          </div>
        </div>

        <div class="flex gap-3">
          <div class="flex flex-1 flex-col gap-2">
            <label for="city" class="font-medium text-gray-700">Cidade</label>
            <InputText
              id="city"
              v-model="form.address.city"
              :disabled="loading"
              :invalid="!!errors['address.city']"
              placeholder="Cidade"
              size="large"
            />
            <small v-if="errors['address.city']" class="text-red-600">{{
              errors["address.city"]
            }}</small>
          </div>
          <div class="flex w-32 flex-col gap-2">
            <label for="state" class="font-medium text-gray-700">Estado</label>
            <Dropdown
              id="state"
              v-model="form.address.state_initials"
              :options="estados"
              optionLabel="sigla"
              optionValue="sigla"
              :disabled="loading"
              :invalid="!!errors['address.state_initials']"
              placeholder="UF"
              size="large"
              :showClear="true"
              @change="onStateChange"
            />
            <small
              v-if="errors['address.state_initials']"
              class="text-red-600"
              >{{ errors["address.state_initials"] }}</small
            >
          </div>
        </div>

        <div class="flex flex-col gap-2">
          <label for="complement" class="font-medium text-gray-700"
            >Complemento (opcional)</label
          >
          <InputText
            id="complement"
            v-model="form.address.complement"
            :disabled="loading"
            placeholder="Apto, bloco, etc"
            class="w-full"
            size="large"
          />
        </div>

        <div class="mt-2 flex gap-3">
          <Button
            label="Voltar"
            icon="pi pi-arrow-left"
            severity="secondary"
            size="large"
            class="flex-1"
            @click.prevent="back"
          />
          <Button
            type="submit"
            label="Finalizar Cadastro"
            icon="pi pi-check"
            iconPos="right"
            :loading="loading"
            :disabled="loading"
            severity="primary"
            size="large"
            class="flex-1"
          />
        </div>
      </div>
    </form>

    <div class="mt-6 text-center">
      <p class="text-gray-600">
        Já tem uma conta?
        <router-link
          to="/login"
          class="text-primary-color hover:text-primary-dark-color ml-1 font-medium transition-colors duration-200"
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
