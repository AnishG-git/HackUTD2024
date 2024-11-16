<script setup>
const res = await useFetch("http://localhost:8080/api/ping");
console.log("res", res);
</script>

<template>
  <div class="flex justify-center items-center w-screen h-screen">
    <div class="underline font-mono text-4xl text-red-900">
      {{ res.data.value }}
    </div>
    <nuxt-link to="/login" class="bg-red-400 p-5 rounded-xl text-2xl shadow-lg text-white hover:bg-red-600">
      <button>register</button>
    </nuxt-link>
  </div>
</template>
