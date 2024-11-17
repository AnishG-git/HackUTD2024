<script setup>
const login = async () => {
  navigateTo("http://localhost:8080/api/auth/google/sso", {
    external: true,
  });
};
</script>

<template>
  <div class="flex flex-col w-screen h-screen bg-gray-900 text-white">
    <navbar />
    <div class="flex h-full justify-center items-center">
      <button
        class="bg-yellow-500 hover:bg-yellow-600 text-black px-4 py-2 rounded"
        @click="login"
      >
        Login
      </button>
    </div>
  </div>
</template>
