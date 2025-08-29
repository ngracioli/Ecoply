<script setup lang="ts">
import IconSquared from "../../shared/IconSquared.vue";

interface BenefitsCardProps {
  icon: string;
  title: string;
  subtitle: string;
}

defineProps<BenefitsCardProps>();
</script>

<template>
  <article
    class="group flex w-full max-w-sm flex-col items-start rounded-2xl px-4 py-8 ring-1 ring-neutral-200/50 transition-shadow duration-300 hover:shadow-2xl"
  >
    <IconSquared :icon="icon" class="mb-4" />
    <h1 class="text-center text-2xl font-bold text-neutral-900">
      {{ title }}
    </h1>
    <h2 class="text-primary-color mb-6 text-lg font-semibold">
      Venda seu excedente com mais valor
    </h2>
    <h3 class="text-xl leading-relaxed text-neutral-700">
      {{ subtitle }}
    </h3>
  </article>
</template>
