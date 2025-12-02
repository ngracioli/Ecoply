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
    class="group hover:shadow-card relative flex w-full max-w-sm flex-col items-center gap-5 rounded-xl px-4 py-6 transition-all duration-500 ease-out hover:-translate-y-2 hover:scale-102 sm:gap-6 sm:rounded-2xl sm:py-8 md:gap-8"
  >
    <span
      class="bg-accent-color text-secondary-color absolute -top-2 -right-2 flex h-7 w-7 items-center justify-center rounded-full text-sm font-bold shadow transition-all duration-400 ease-out group-hover:scale-110 group-hover:rotate-12 group-hover:shadow-lg sm:-top-3 sm:-right-3 sm:h-8 sm:w-8 sm:text-base"
    >
      {{ step }}
    </span>
    <IconCircle
      :icon="icon"
      class="transition-all duration-400 ease-out group-hover:scale-110 group-hover:rotate-3"
    />
    <div class="flex flex-col items-center justify-center">
      <h1
        class="mb-3 text-center text-lg font-bold text-neutral-700 sm:mb-4 sm:text-xl md:text-2xl"
      >
        {{ title }}
      </h1>
      <h3
        class="text-center text-base leading-relaxed text-neutral-600 sm:text-lg md:text-xl"
      >
        {{ subtitle }}
      </h3>
    </div>
  </article>
</template>
