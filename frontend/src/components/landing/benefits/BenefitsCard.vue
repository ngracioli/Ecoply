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
    class="group hover:ring-primary-color/20 flex w-full max-w-md flex-col items-start rounded-2xl px-6 py-10 ring-1 ring-neutral-200/50 transition-all duration-500 ease-out hover:-translate-y-2 hover:scale-102 hover:shadow-2xl"
  >
    <IconSquared :icon="icon" class="mb-4 text-3xl" />
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
