<script setup lang="ts">
import InputText from "../../components/shared/forms/InputText.vue";
import Password from "../../components/shared/forms/Password.vue";
</script>

<template>
  <main class="grid grid-cols-2">
    <section
      class="flex h-screen w-full flex-col justify-center bg-gradient-to-br from-green-600 to-green-400 p-24 text-left text-white"
    >
      <h1 class="mb-2 text-5xl font-bold">Bem-vindo ao Ecoply</h1>
      <h2 class="mb-8 text-xl">
        O marketplace de energia renovável que conecta você a um futuro
        sustentável
      </h2>
      <ul class="list-disc pl-5 text-lg">
        <li>Leilões transparentes de energia limpa</li>
        <li>Certificações verificadas</li>
        <li>Impacto ambiental positivo</li>
      </ul>
    </section>
    <section
      class="flex h-screen w-full flex-col items-center justify-center bg-gradient-to-tl from-neutral-400 to-neutral-100 text-left"
    >
      <div class="w-full max-w-8/12 rounded-xl bg-white px-12 py-12">
        <section class="text-center">
          <i class="pi pi-users text-5xl text-green-600"></i>
          <h1 class="text-2xl font-bold">Entrar no Ecoply</h1>
          <h2 class="mb-8 text-xl">Acesse sua conta para continuar</h2>
        </section>
        <section>
          <form action="" class="flex flex-col gap-5">
            <div>
              <label for="email">Teste</label>
              <InputText id="email" placeholder="Email" />
            </div>
            <div>
              <label for="password">Senha</label>
              <Password id="password" placeholder="Senha" />
            </div>
            <button
              type="submit"
              class="mt-4 w-full cursor-pointer items-center self-center rounded-xl bg-green-600 px-2 py-3 text-xl font-bold text-white hover:bg-green-700"
            >
              Entrar
            </button>
            <p class="text-md self-center text-gray-700">
              Não tem uma conta?
              <a href="#" class="ml-2 text-green-600 hover:text-green-800"
                >Cadastre-se</a
              >
            </p>
          </form>
        </section>
      </div>
    </section>
  </main>
</template>

<style scoped></style>
