<script setup>
import { ref } from "vue";
const res = await useFetch("http://localhost:8080/api/ping");
console.log("res", res);

const register = async () => {
  window.location.href = "http://localhost:8080/api/auth/google/sso";
  console.log(document.cookie);
};

const logout = async () => {
  window.location.href = "http://localhost:8080/api/auth/google/logout";
  console.log(document.cookie);
};
</script>

<template>
  <div class="flex flex-col justify-center items-center h-screen">
    <div class="text-black text-2xl mb-10">Make an account, uggo</div>
    <button
      @click="register"
      class="bg-blue-500 hover:bg-blue-700 text-white font-bold p-3 rounded-full w-400vh"
    >
      Register
    </button>
    <button
      @click="logout"
      class="bg-blue-500 hover:bg-blue-700 text-white font-bold p-3 rounded-full w-400vh"
    >
      Logout
    </button>
  </div>
</template>
