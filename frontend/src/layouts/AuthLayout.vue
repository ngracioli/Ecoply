<script setup lang="ts"></script>

<template>
  <main class="grid min-h-screen grid-cols-1 lg:grid-cols-2">
    <section
      class="flex flex-col justify-center bg-[#059669] bg-gradient-to-br p-8 text-white lg:p-16"
    >
      <div class="flex flex-col items-center">
        <img src="/ecoply_login.svg" alt="" class="w-4/5" />
        <div class="mt-8 text-center">
          <h1 class="mb-2 text-4xl font-bold">ECOPLY</h1>
        </div>
      </div>
    </section>

    <section class="flex items-center justify-center bg-gray-50 p-6 lg:p-12">
      <div class="w-full max-w-md">
        <slot />
      </div>
    </section>
  </main>
</template>
