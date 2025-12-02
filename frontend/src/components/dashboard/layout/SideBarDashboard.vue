<script setup lang="ts">
import { computed, ref, type Component } from "vue";
import { useStore } from "vuex";
import {
  ChevronDown,
  FileText,
  Home,
  LogOut,
  User,
  Zap,
  Menu,
  X,
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

const isMobileMenuOpen = ref(false);

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
  isMobileMenuOpen.value = false;
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
  isMobileMenuOpen.value = false;
  emit("navigate", key);
};

const handleLogout = () => {
  isMobileMenuOpen.value = false;
  emit("logout");
};

const toggleMobileMenu = () => {
  isMobileMenuOpen.value = !isMobileMenuOpen.value;
};
</script>

<template>
  <button
    @click="toggleMobileMenu"
    class="fixed top-4 left-4 z-50 flex touch-manipulation items-center justify-center rounded-lg border border-neutral-200 bg-white p-2.5 shadow-lg md:hidden"
  >
    <Menu v-if="!isMobileMenuOpen" :size="24" class="text-neutral-700" />
    <X v-else :size="24" class="text-neutral-700" />
  </button>

  <Transition
    name="mobile-menu"
    enter-active-class="transition-opacity duration-300"
    leave-active-class="transition-opacity duration-200"
    enter-from-class="opacity-0"
    enter-to-class="opacity-100"
    leave-from-class="opacity-100"
    leave-to-class="opacity-0"
  >
    <div
      v-if="isMobileMenuOpen"
      @click="isMobileMenuOpen = false"
      class="fixed inset-0 z-40 bg-black/50 backdrop-blur-sm md:hidden"
    ></div>
  </Transition>

  <nav
    :class="[
      'flex flex-col justify-between border-r border-neutral-200 bg-white transition-transform duration-300 ease-in-out',
      'fixed inset-y-0 left-0 z-40 h-screen w-64 md:static md:h-screen',
      isMobileMenuOpen ? 'translate-x-0' : '-translate-x-full md:translate-x-0',
    ]"
  >
    <div class="p-4 sm:p-6">
      <div class="mb-6 flex items-center gap-2.5 sm:mb-8 sm:gap-3">
        <div
          class="flex h-9 w-9 shrink-0 items-center justify-center rounded-lg bg-gradient-to-br from-emerald-500 to-teal-600 text-white shadow-md sm:h-10 sm:w-10"
        >
          <Zap :size="22" class="sm:hidden" :stroke-width="2.5" />
          <Zap :size="24" class="hidden sm:block" :stroke-width="2.5" />
        </div>
        <span class="text-lg font-bold text-neutral-800 sm:text-xl"
          >Ecoply</span
        >
      </div>

      <ul class="flex flex-col gap-1">
        <li v-for="item in menuItems" :key="item.key">
          <button
            v-if="!item.hasDropdown"
            @click="handleItemClick(item.key)"
            :class="[
              'group flex w-full touch-manipulation items-center gap-2.5 rounded-lg px-3 py-2.5 text-left transition-all duration-200 sm:gap-3 sm:px-4 sm:py-3',
              activeItem === item.key
                ? 'bg-emerald-50 text-emerald-700 shadow-sm'
                : 'text-neutral-600 hover:bg-neutral-50 hover:text-neutral-900 active:scale-[0.98]',
            ]"
          >
            <component
              :is="item.icon"
              :size="18"
              :stroke-width="2"
              :class="[
                'shrink-0 transition-colors duration-200 sm:hidden',
                activeItem === item.key
                  ? 'text-emerald-600'
                  : 'text-neutral-400 group-hover:text-neutral-600',
              ]"
            />
            <component
              :is="item.icon"
              :size="20"
              :stroke-width="2"
              :class="[
                'hidden shrink-0 transition-colors duration-200 sm:block',
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
                'group flex w-full touch-manipulation items-center justify-between rounded-lg px-3 py-2.5 text-left transition-all duration-200 sm:px-4 sm:py-3',
                isProfileActive
                  ? 'bg-emerald-50 text-emerald-700 shadow-sm'
                  : 'text-neutral-600 hover:bg-neutral-50 hover:text-neutral-900 active:scale-[0.98]',
              ]"
            >
              <div class="flex items-center gap-2.5 sm:gap-3">
                <component
                  :is="item.icon"
                  :size="18"
                  :stroke-width="2"
                  :class="[
                    'shrink-0 transition-colors duration-200 sm:hidden',
                    isProfileActive
                      ? 'text-emerald-600'
                      : 'text-neutral-400 group-hover:text-neutral-600',
                  ]"
                />
                <component
                  :is="item.icon"
                  :size="20"
                  :stroke-width="2"
                  :class="[
                    'hidden shrink-0 transition-colors duration-200 sm:block',
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
                  'shrink-0 transition-transform duration-200',
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
                  'group flex w-full touch-manipulation items-center gap-2.5 rounded-lg py-2 pr-3 pl-10 text-left transition-all duration-200 sm:gap-3 sm:pr-4 sm:pl-12',
                  activeItem === subItem.key
                    ? 'bg-emerald-100 text-emerald-700'
                    : 'text-neutral-600 hover:bg-neutral-50 hover:text-neutral-900 active:scale-[0.98]',
                ]"
              >
                <span class="text-sm">{{ subItem.label }}</span>
              </button>
            </div>
          </div>
        </li>
      </ul>
    </div>

    <div class="border-t border-neutral-200 p-4 sm:p-6">
      <button
        @click="handleLogout"
        class="group flex w-full touch-manipulation items-center gap-2.5 rounded-lg px-3 py-2.5 text-left text-neutral-600 transition-all duration-200 hover:bg-red-50 hover:text-red-600 active:scale-[0.98] sm:gap-3 sm:px-4 sm:py-3"
      >
        <LogOut
          :size="18"
          :stroke-width="2"
          class="shrink-0 text-neutral-400 transition-colors duration-200 group-hover:text-red-500 sm:hidden"
        />
        <LogOut
          :size="20"
          :stroke-width="2"
          class="hidden shrink-0 text-neutral-400 transition-colors duration-200 group-hover:text-red-500 sm:block"
        />
        <span class="text-sm font-medium">Logout</span>
      </button>
    </div>
  </nav>
</template>

<style scoped></style>
