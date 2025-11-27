<script setup lang="ts">
import { computed, ref, watch, type Component } from "vue";
import { useStore } from "vuex";
import {
  ChevronDown,
  FileText,
  Home,
  LogOut,
  User,
  Zap,
} from "lucide-vue-next";

interface SubMenuItem {
  label: string;
  key: string;
}

interface MenuItem {
  icon: Component;
  label: string;
  key: string;
  hasDropdown?: boolean;
  subItems?: SubMenuItem[];
}

type MenuItemKey =
  | "overview"
  | "offers"
  | "history"
  | "profile"
  | "manage-offers";

const emit = defineEmits<{
  navigate: [key: string];
  logout: [];
}>();

const store = useStore();
const activeItem = ref<MenuItemKey>("overview");
const isProfileDropdownOpen = ref(false);

const isSupplier = computed(() => store.getters["user/isSupplier"]);

const baseMenuItems: MenuItem[] = [
  { icon: Home, label: "Resumo Geral", key: "overview" },
  { icon: Zap, label: "Ofertas de Energia", key: "offers" },
  { icon: FileText, label: "Histórico de Negociações", key: "history" },
];

const profileMenuItem = computed<MenuItem>(() => ({
  icon: User,
  label: "Perfil",
  key: "profile",
  hasDropdown: true,
  subItems: isSupplier.value
    ? [
        { label: "Informações Pessoais", key: "profile" },
        { label: "Gerenciar Ofertas", key: "manage-offers" },
      ]
    : [{ label: "Informações Pessoais", key: "profile" }],
}));

const menuItems = computed(() => [...baseMenuItems, profileMenuItem.value]);

const isProfileActive = computed(
  () =>
    activeItem.value === "profile" ||
    activeItem.value === "manage-offers" ||
    isProfileDropdownOpen.value,
);

const handleItemClick = (key: string) => {
  activeItem.value = key as MenuItemKey;
  isProfileDropdownOpen.value = false;
  emit("navigate", key);
};

const toggleProfileDropdown = () => {
  const wasOpen = isProfileDropdownOpen.value;
  isProfileDropdownOpen.value = !isProfileDropdownOpen.value;

  if (
    !wasOpen &&
    activeItem.value !== "profile" &&
    activeItem.value !== "manage-offers"
  ) {
    activeItem.value = "profile";
    emit("navigate", "profile");
  }
};

const handleSubItemClick = (key: string) => {
  activeItem.value = key as MenuItemKey;
  emit("navigate", key);
};

const handleLogout = () => {
  emit("logout");
};
</script>

<template>
  <nav
    class="flex h-screen w-64 flex-col justify-between border-r border-neutral-200 bg-white"
  >
    <div class="p-6">
      <div class="mb-8 flex items-center gap-3">
        <div
          class="flex h-10 w-10 items-center justify-center rounded-lg bg-gradient-to-br from-emerald-500 to-teal-600 text-white shadow-md"
        >
          <Zap :size="24" :stroke-width="2.5" />
        </div>
        <span class="text-xl font-bold text-neutral-800">Ecoply</span>
      </div>

      <ul class="flex flex-col gap-1">
        <li v-for="item in menuItems" :key="item.key">
          <button
            v-if="!item.hasDropdown"
            @click="handleItemClick(item.key)"
            :class="[
              'group flex w-full items-center gap-3 rounded-lg px-4 py-3 text-left transition-all duration-200',
              activeItem === item.key
                ? 'bg-emerald-50 text-emerald-700 shadow-sm'
                : 'text-neutral-600 hover:bg-neutral-50 hover:text-neutral-900',
            ]"
          >
            <component
              :is="item.icon"
              :size="20"
              :stroke-width="2"
              :class="[
                'transition-colors duration-200',
                activeItem === item.key
                  ? 'text-emerald-600'
                  : 'text-neutral-400 group-hover:text-neutral-600',
              ]"
            />
            <span class="text-sm font-medium">{{ item.label }}</span>
          </button>

          <div v-else-if="item.hasDropdown">
            <button
              @click="toggleProfileDropdown"
              :class="[
                'group flex w-full items-center justify-between rounded-lg px-4 py-3 text-left transition-all duration-200',
                isProfileActive
                  ? 'bg-emerald-50 text-emerald-700 shadow-sm'
                  : 'text-neutral-600 hover:bg-neutral-50 hover:text-neutral-900',
              ]"
            >
              <div class="flex items-center gap-3">
                <component
                  :is="item.icon"
                  :size="20"
                  :stroke-width="2"
                  :class="[
                    'transition-colors duration-200',
                    isProfileActive
                      ? 'text-emerald-600'
                      : 'text-neutral-400 group-hover:text-neutral-600',
                  ]"
                />
                <span class="text-sm font-medium">{{ item.label }}</span>
              </div>
              <ChevronDown
                :size="16"
                :class="[
                  'transition-transform duration-200',
                  isProfileDropdownOpen ? 'rotate-180' : '',
                  isProfileActive
                    ? 'text-emerald-600'
                    : 'text-neutral-400 group-hover:text-neutral-600',
                ]"
              />
            </button>

            <div
              v-show="isProfileDropdownOpen"
              class="mt-1 space-y-1 overflow-hidden transition-all duration-200"
            >
              <button
                v-for="subItem in item.subItems"
                :key="subItem.key"
                @click="handleSubItemClick(subItem.key)"
                :class="[
                  'group flex w-full items-center gap-3 rounded-lg py-2 pr-4 pl-12 text-left transition-all duration-200',
                  activeItem === subItem.key
                    ? 'bg-emerald-100 text-emerald-700'
                    : 'text-neutral-600 hover:bg-neutral-50 hover:text-neutral-900',
                ]"
              >
                <span class="text-sm">{{ subItem.label }}</span>
              </button>
            </div>
          </div>
        </li>
      </ul>
    </div>

    <div class="border-t border-neutral-200 p-6">
      <button
        @click="handleLogout"
        class="group flex w-full items-center gap-3 rounded-lg px-4 py-3 text-left text-neutral-600 transition-all duration-200 hover:bg-red-50 hover:text-red-600"
      >
        <LogOut
          :size="20"
          :stroke-width="2"
          class="text-neutral-400 transition-colors duration-200 group-hover:text-red-500"
        />
        <span class="text-sm font-medium">Logout</span>
      </button>
    </div>
  </nav>
</template>

<style scoped></style>
