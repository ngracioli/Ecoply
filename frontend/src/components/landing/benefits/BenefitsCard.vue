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
    class="group hover:ring-primary-color/20 flex w-full max-w-md flex-col items-start rounded-xl px-4 py-6 ring-1 ring-neutral-200/50 transition-all duration-500 ease-out hover:-translate-y-2 hover:scale-102 hover:shadow-2xl sm:rounded-2xl sm:px-6 sm:py-8 lg:py-10"
  >
    <IconSquared :icon="icon" class="mb-3 text-2xl sm:mb-4 sm:text-3xl" />
    <h1 class="text-center text-xl font-bold text-neutral-900 sm:text-2xl">
      {{ title }}
    </h1>
    <h2
      class="text-primary-color mb-4 text-base font-semibold sm:mb-6 sm:text-lg"
    >
      Venda seu excedente com mais valor
    </h2>
    <h3
      class="text-base leading-relaxed text-neutral-700 sm:text-lg lg:text-xl"
    >
      {{ subtitle }}
    </h3>
  </article>
</template>
