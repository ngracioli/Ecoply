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
    class="group hover:shadow-card relative flex w-full max-w-sm flex-col items-center gap-8 rounded-2xl px-4 py-8 transition-shadow duration-300"
  >
    <span
      class="bg-secondary-color absolute -top-3 -right-3 flex h-8 w-8 items-center justify-center rounded-full font-bold text-white shadow transition-transform duration-300 group-hover:scale-90"
    >
      {{ step }}
    </span>
    <IconCircle
      :icon="icon"
      class="transition-transform duration-300 group-hover:scale-110"
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
