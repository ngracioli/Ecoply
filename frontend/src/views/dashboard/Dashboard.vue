<script setup lang="ts">
import { onMounted, ref } from "vue";
import type { Component } from "vue";
import { useRouter } from "vue-router";
import { useStore } from "vuex";
import HistorySection from "../../components/dashboard/sections/HistorySection.vue";
import ManageOffersSection from "../../components/dashboard/sections/ManageOffersSection.vue";
import OffersSection from "../../components/dashboard/sections/OffersSection.vue";
import OverviewSection from "../../components/dashboard/sections/OverviewSection.vue";
import ProfileSection from "../../components/dashboard/sections/ProfileSection.vue";
import SideBarDashboard from "../../components/dashboard/layout/SideBarDashboard.vue";
import TopBar from "../../components/dashboard/layout/TopBar.vue";

type DashboardSection =
  | "overview"
  | "offers"
  | "history"
  | "profile"
  | "manage-offers";

const store = useStore();
const router = useRouter();

const activeSection = ref<DashboardSection>("overview");

const sectionComponents: Record<DashboardSection, Component> = {
  overview: OverviewSection,
  offers: OffersSection,
  history: HistorySection,
  profile: ProfileSection,
  "manage-offers": ManageOffersSection,
};

onMounted(async () => {
  try {
    await store.dispatch("user/fetchCurrentUser");
  } catch (error) {
    router.push({ name: "Login" });
  }
});

const handleNavigate = (key: string) => {
  activeSection.value = key as DashboardSection;
};

const logout = async () => {
  await store.dispatch("auth/logout");
  router.push({ name: "Login" });
};
</script>

<template>
  <main class="flex h-dvh w-full overflow-hidden bg-neutral-50">
    <SideBarDashboard @navigate="handleNavigate" @logout="logout" />

    <div class="flex flex-1 flex-col overflow-hidden">
      <TopBar />
      <div class="flex-1 overflow-y-auto">
        <div class="p-3 sm:p-5 md:p-6 lg:p-8">
          <Transition
            name="fade"
            mode="out-in"
            enter-active-class="transition-all duration-300 ease-out"
            leave-active-class="transition-all duration-200 ease-in"
            enter-from-class="opacity-0 translate-y-4"
            enter-to-class="opacity-100 translate-y-0"
            leave-from-class="opacity-100 translate-y-0"
            leave-to-class="opacity-0 -translate-y-4"
          >
            <component
              :is="sectionComponents[activeSection]"
              :key="activeSection"
            />
          </Transition>
        </div>
      </div>
    </div>
  </main>
</template>
