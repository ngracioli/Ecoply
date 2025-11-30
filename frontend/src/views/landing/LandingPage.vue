<script setup lang="ts">
import { ref, onMounted } from "vue";
import api from "../../axios";
import { ANALYTICS_ENDPOINTS } from "../../api/endpoints";
import type { PlatformAnalyticsResponse } from "../../types/analytics";
import Header from "../../components/landing/Header.vue";
import Hero from "../../components/landing/Hero.vue";
import HowItWorks from "../../components/landing/HowItWorks.vue";
import BenefitsSection from "../../components/landing/BenefitsSection.vue";
import OfferDemonstrationSection from "../../components/landing/OfferDemonstrationSection.vue";
import JoinUsSection from "../../components/landing/JoinUsSection.vue";
import FooterLanding from "../../components/landing/FooterLanding.vue";

const platformData = ref({
  sucessful_purchases: 0,
  active_offers: 0,
  money_transacted: 0,
  energy_transacted: 0,
});

const loading = ref(true);

const fetchPlatformAnalytics = async () => {
  try {
    const response = await api.get<PlatformAnalyticsResponse>(
      ANALYTICS_ENDPOINTS.PLATFORM,
    );
    if (response.data?.data) {
      platformData.value = response.data.data;
    }
  } catch (error) {
    console.error("Erro ao buscar analytics:", error);
  } finally {
    loading.value = false;
  }
};

onMounted(() => {
  fetchPlatformAnalytics();
});
</script>

<template>
  <Header />
  <Hero :platform-data="platformData" :loading="loading" />
  <HowItWorks />
  <BenefitsSection />
  <OfferDemonstrationSection :platform-data="platformData" :loading="loading" />
  <JoinUsSection />
  <FooterLanding />
</template>

<style scoped></style>
