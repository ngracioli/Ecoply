<script setup lang="ts">
import IconCircle from "../shared/IconCircle.vue";

interface NumericCardProps {
  step: number;
  icon: string;
  title: string;
  subtitle: string;
}

defineProps<NumericCardProps>();
</script>

<template>
  <article
    class="group hover:shadow-card relative flex w-full max-w-sm flex-col items-center gap-8 rounded-2xl px-4 py-8 transition-all duration-500 ease-out hover:-translate-y-2 hover:scale-102"
  >
    <span
      class="bg-accent-color text-secondary-color absolute -top-3 -right-3 flex h-8 w-8 items-center justify-center rounded-full font-bold shadow transition-all duration-400 ease-out group-hover:scale-110 group-hover:rotate-12 group-hover:shadow-lg"
    >
      {{ step }}
    </span>
    <IconCircle
      :icon="icon"
      class="transition-all duration-400 ease-out group-hover:scale-110 group-hover:rotate-3"
    />
    <div class="flex flex-col items-center justify-center">
      <h1 class="mb-4 text-center text-2xl font-bold text-neutral-700">
        {{ title }}
      </h1>
      <h3 class="text-center text-xl leading-relaxed text-neutral-600">
        {{ subtitle }}
      </h3>
    </div>
  </article>
</template>
