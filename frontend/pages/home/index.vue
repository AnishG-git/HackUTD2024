<script setup>
const logout = async () => {
  navigateTo("http://localhost:8080/api/auth/google/logout", {
    external: true,
  });
};
</script>

<template>
  <div class="flex flex-col w-screen h-screen bg-gray-900 text-white">
    <navbar />
    <div class="flex h-full justify-center items-center"></div>
  </div>
</template>
