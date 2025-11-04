<script setup lang="ts">
import { computed } from "vue";
import { useStore } from "vuex";
import {
  Briefcase,
  Building,
  FileText,
  Hash,
  Mail,
  MapPin,
  User,
} from "lucide-vue-next";
import type { Component } from "vue";
import type {
  User as UserType,
  UserType as UserTypeEnum,
} from "../../types/user";

interface ProfileField {
  icon: Component;
  label: string;
  value: string;
}

interface AddressField {
  label: string;
  value: string;
}

const USER_TYPE_LABELS: Record<UserTypeEnum, string> = {
  supplier: "Fornecedor",
  buyer: "Comprador",
};

const formatCEP = (cep: string): string => {
  return cep.replace(/(\d{5})(\d{3})/, "$1-$2");
};

const formatCNPJ = (cnpj: string): string => {
  return cnpj.replace(/(\d{2})(\d{3})(\d{3})(\d{4})(\d{2})/, "$1.$2.$3/$4-$5");
};

const store = useStore();

const currentUser = computed<UserType | null>(
  () => store.state.user.currentUser,
);

const userTypeLabel = computed<string>(() => {
  if (!currentUser.value) return "";
  return USER_TYPE_LABELS[currentUser.value.user_type as UserTypeEnum];
});

const fullAddress = computed<string>(() => {
  if (!currentUser.value?.address) return "";

  const addr = currentUser.value.address;
  const complement = addr.complement ? ` - ${addr.complement}` : "";

  return `${addr.street}, ${addr.number}${complement}, ${addr.neighborhood}, ${addr.city} - ${addr.state_initial}`;
});

const cepFormatted = computed<string>(() => {
  if (!currentUser.value?.address?.cep) return "";
  return formatCEP(currentUser.value.address.cep);
});

const cnpjFormatted = computed<string>(() => {
  if (!currentUser.value?.agent?.cnpj) return "";
  return formatCNPJ(currentUser.value.agent.cnpj);
});

const profileFields = computed<ProfileField[]>(() => {
  if (!currentUser.value) return [];

  return [
    {
      icon: Mail,
      label: "E-mail",
      value: currentUser.value.email,
    },
    {
      icon: Building,
      label: "Empresa",
      value: currentUser.value.agent.company_name,
    },
    {
      icon: FileText,
      label: "CNPJ",
      value: cnpjFormatted.value,
    },
    {
      icon: Hash,
      label: "Código CCEE",
      value: currentUser.value.agent.ccee_code,
    },
    {
      icon: Briefcase,
      label: "Submercado",
      value: currentUser.value.agent.submarket_name,
    },
    {
      icon: MapPin,
      label: "CEP",
      value: cepFormatted.value,
    },
  ];
});

const addressFields = computed<AddressField[]>(() => {
  if (!currentUser.value?.address) return [];

  const addr = currentUser.value.address;

  return [
    { label: "Rua", value: addr.street },
    { label: "Número", value: addr.number },
    { label: "Complemento", value: addr.complement || "—" },
    { label: "Bairro", value: addr.neighborhood },
    { label: "Cidade", value: addr.city },
    { label: "Estado", value: `${addr.state} (${addr.state_initial})` },
  ];
});
</script>

<template>
  <div v-if="currentUser" class="flex flex-col gap-6">
    <div>
      <h2 class="text-2xl font-bold text-neutral-900">Informações Pessoais</h2>
      <p class="mt-1 text-sm text-neutral-500">
        Gerencie seus dados pessoais e preferências
      </p>
    </div>

    <div class="grid grid-cols-1 gap-6 lg:grid-cols-3">
      <!-- Profile Card -->
      <div
        class="flex items-center justify-center rounded-xl border border-neutral-200 bg-white p-6 shadow-sm transition-all duration-300 hover:shadow-md lg:col-span-1"
      >
        <div class="flex w-full flex-col items-center text-center">
          <div
            class="mb-4 flex h-24 w-24 items-center justify-center rounded-full bg-gradient-to-br from-emerald-500 to-teal-600 text-white shadow-lg"
          >
            <User :size="48" :stroke-width="2" />
          </div>
          <h3 class="text-xl font-bold text-neutral-900">
            {{ currentUser.name }}
          </h3>
          <p class="mt-1 text-sm text-neutral-500">{{ userTypeLabel }}</p>
          <div class="mt-2 rounded-full bg-emerald-100 px-3 py-1">
            <p class="text-xs font-medium text-emerald-700">
              {{ currentUser.agent.company_name }}
            </p>
          </div>
          <!-- <button
            class="mt-6 w-full rounded-lg border-2 border-emerald-500 bg-transparent py-2.5 text-sm font-medium text-emerald-600 transition-all duration-200 hover:bg-emerald-500 hover:text-white"
          >
            Editar Perfil
          </button> -->
        </div>
      </div>

      <!-- Information Details -->
      <div
        class="rounded-xl border border-neutral-200 bg-white p-6 shadow-sm transition-all duration-300 hover:shadow-md lg:col-span-2"
      >
        <h3 class="mb-6 text-lg font-semibold text-neutral-900">
          Detalhes da Conta
        </h3>
        <div class="grid grid-cols-1 gap-6 md:grid-cols-2">
          <div
            v-for="field in profileFields"
            :key="field.label"
            class="group flex items-start gap-4 rounded-lg p-4 transition-all duration-200 hover:bg-neutral-50"
          >
            <div
              class="flex h-10 w-10 flex-shrink-0 items-center justify-center rounded-lg bg-emerald-50 text-emerald-600 transition-all duration-200 group-hover:bg-emerald-100"
            >
              <component :is="field.icon" :size="20" :stroke-width="2" />
            </div>
            <div class="flex-1">
              <p class="text-xs font-medium text-neutral-500 uppercase">
                {{ field.label }}
              </p>
              <p class="mt-1 text-sm font-medium break-words text-neutral-900">
                {{ field.value }}
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Address Section -->
    <div
      class="rounded-xl border border-neutral-200 bg-white p-6 shadow-sm transition-all duration-300 hover:shadow-md"
    >
      <h3 class="mb-6 text-lg font-semibold text-neutral-900">Endereço</h3>
      <div class="mb-4 rounded-lg bg-emerald-50 p-4">
        <p class="text-sm text-neutral-700">
          {{ fullAddress }}
        </p>
      </div>
      <div class="grid grid-cols-1 gap-4 md:grid-cols-3">
        <div
          v-for="field in addressFields"
          :key="field.label"
          class="rounded-lg border border-neutral-200 p-4 transition-all duration-200 hover:border-emerald-300"
        >
          <p class="text-xs font-medium text-neutral-500 uppercase">
            {{ field.label }}
          </p>
          <p class="mt-1 text-sm font-medium text-neutral-900">
            {{ field.value }}
          </p>
        </div>
      </div>
    </div>

    <!-- Additional Settings -->
    <div
      class="rounded-xl border border-neutral-200 bg-white p-6 shadow-sm transition-all duration-300 hover:shadow-md"
    >
      <h3 class="mb-6 text-lg font-semibold text-neutral-900">Preferências</h3>
      <div class="grid grid-cols-1 gap-6 md:grid-cols-2">
        <div
          class="flex items-center justify-between rounded-lg border border-neutral-200 p-4 transition-all duration-200 hover:border-emerald-300"
        >
          <div>
            <h4 class="font-medium text-neutral-900">
              Notificações por E-mail
            </h4>
            <p class="mt-1 text-sm text-neutral-500">
              Receba atualizações importantes
            </p>
          </div>
          <label class="relative inline-flex cursor-pointer items-center">
            <input type="checkbox" class="peer sr-only" checked />
            <div
              class="peer h-6 w-11 rounded-full bg-neutral-300 peer-checked:bg-emerald-500 peer-focus:ring-2 peer-focus:ring-emerald-500/20 peer-focus:outline-none after:absolute after:top-[2px] after:left-[2px] after:h-5 after:w-5 after:rounded-full after:bg-white after:transition-all after:content-[''] peer-checked:after:translate-x-full"
            ></div>
          </label>
        </div>

        <div
          class="flex items-center justify-between rounded-lg border border-neutral-200 p-4 transition-all duration-200 hover:border-emerald-300"
        >
          <div>
            <h4 class="font-medium text-neutral-900">Alertas de Ofertas</h4>
            <p class="mt-1 text-sm text-neutral-500">
              Novidades no marketplace
            </p>
          </div>
          <label class="relative inline-flex cursor-pointer items-center">
            <input type="checkbox" class="peer sr-only" checked />
            <div
              class="peer h-6 w-11 rounded-full bg-neutral-300 peer-checked:bg-emerald-500 peer-focus:ring-2 peer-focus:ring-emerald-500/20 peer-focus:outline-none after:absolute after:top-[2px] after:left-[2px] after:h-5 after:w-5 after:rounded-full after:bg-white after:transition-all after:content-[''] peer-checked:after:translate-x-full"
            ></div>
          </label>
        </div>
      </div>
    </div>
  </div>

  <!-- Loading or No User State -->
  <div v-else class="flex h-96 items-center justify-center">
    <div class="text-center">
      <div class="mb-4 flex justify-center">
        <div
          class="h-12 w-12 animate-spin rounded-full border-4 border-emerald-500 border-t-transparent"
        ></div>
      </div>
      <p class="text-neutral-500">Carregando informações do perfil...</p>
    </div>
  </div>
</template>

<style scoped></style>
