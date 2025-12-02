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
} from "../../../types/user";

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
  <div v-if="currentUser" class="flex flex-col gap-4 sm:gap-6">
    <div>
      <h2 class="text-xl font-bold text-neutral-900 sm:text-2xl">
        Informações Pessoais
      </h2>
      <p class="mt-1 text-xs text-neutral-500 sm:text-sm">
        Gerencie seus dados pessoais e preferências
      </p>
    </div>

    <div class="grid grid-cols-1 gap-4 sm:gap-6 lg:grid-cols-3">
      <div
        class="flex items-center justify-center rounded-lg border border-neutral-200 bg-white p-6 shadow-sm transition-all duration-300 hover:shadow-md sm:rounded-xl sm:p-6 lg:col-span-1"
      >
        <div class="flex w-full flex-col items-center text-center">
          <div
            class="mb-3 flex h-20 w-20 items-center justify-center rounded-full bg-gradient-to-br from-emerald-500 to-teal-600 text-white shadow-lg sm:mb-4 sm:h-24 sm:w-24"
          >
            <User :size="40" :stroke-width="2" class="sm:hidden" />
            <User :size="48" :stroke-width="2" class="hidden sm:block" />
          </div>
          <h3 class="text-lg font-bold text-neutral-900 sm:text-xl">
            {{ currentUser.name }}
          </h3>
          <p class="mt-1 text-xs text-neutral-500 sm:text-sm">
            {{ userTypeLabel }}
          </p>
          <div class="mt-2 rounded-full bg-emerald-100 px-3 py-1">
            <p
              class="max-w-[200px] truncate text-xs font-medium text-emerald-700 sm:max-w-none"
            >
              {{ currentUser.agent.company_name }}
            </p>
          </div>
        </div>
      </div>

      <div
        class="rounded-lg border border-neutral-200 bg-white p-4 shadow-sm transition-all duration-300 hover:shadow-md sm:rounded-xl sm:p-6 lg:col-span-2"
      >
        <h3
          class="mb-4 text-base font-semibold text-neutral-900 sm:mb-6 sm:text-lg"
        >
          Detalhes da Conta
        </h3>
        <div class="grid grid-cols-1 gap-4 sm:gap-6 md:grid-cols-2">
          <div
            v-for="field in profileFields"
            :key="field.label"
            class="group flex items-start gap-3 rounded-lg p-3 transition-all duration-200 hover:bg-neutral-50 sm:gap-4 sm:p-4"
          >
            <div
              class="flex h-9 w-9 flex-shrink-0 items-center justify-center rounded-lg bg-emerald-50 text-emerald-600 transition-all duration-200 group-hover:bg-emerald-100 sm:h-10 sm:w-10"
            >
              <component
                :is="field.icon"
                :size="18"
                :stroke-width="2"
                class="sm:hidden"
              />
              <component
                :is="field.icon"
                :size="20"
                :stroke-width="2"
                class="hidden sm:block"
              />
            </div>
            <div class="min-w-0 flex-1">
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

    <div
      class="rounded-lg border border-neutral-200 bg-white p-4 shadow-sm transition-all duration-300 hover:shadow-md sm:rounded-xl sm:p-6"
    >
      <h3
        class="mb-4 text-base font-semibold text-neutral-900 sm:mb-6 sm:text-lg"
      >
        Endereço
      </h3>
      <div class="mb-4 rounded-lg bg-emerald-50 p-3 sm:p-4">
        <p class="text-xs text-neutral-700 sm:text-sm">
          {{ fullAddress }}
        </p>
      </div>
      <div
        class="grid grid-cols-1 gap-3 sm:grid-cols-2 sm:gap-4 md:grid-cols-3"
      >
        <div
          v-for="field in addressFields"
          :key="field.label"
          class="rounded-lg border border-neutral-200 p-3 transition-all duration-200 hover:border-emerald-300 sm:p-4"
        >
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

  <div v-else class="flex h-64 items-center justify-center sm:h-96">
    <div class="text-center">
      <div class="mb-4 flex justify-center">
        <div
          class="h-10 w-10 animate-spin rounded-full border-4 border-emerald-500 border-t-transparent sm:h-12 sm:w-12"
        ></div>
      </div>
      <p class="text-xs text-neutral-500 sm:text-sm">
        Carregando informações do perfil...
      </p>
    </div>
  </div>
</template>

<style scoped></style>
