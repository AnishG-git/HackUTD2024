<script setup>
import { ref } from "vue";
import { useAuth, setLoggedIn } from "~/composables/useAuth";

const { loggedIn } = useAuth();

const login = async () => {
  setLoggedIn(true);
  navigateTo("http://localhost:8080/api/auth/google/sso", {
    external: true,
  });
};
</script>

<template>
  <div class="flex flex-col w-screen h-screen bg-gray-900 text-white">
    <navbar />
    <div class="flex justify-center items-center h-full">
      <div class="w-1/3 h-1/2 rounded p-3">
        <div class="flex flex-col justify-center items-center h-full">
          <div
            class="border-2 border-yellow-200 mb-10 rounded p-3 w-1/3 h-2/3 m-3 flex justify-center items-center"
          >
            <img
              src="/public/logo.svg"
              alt="logo"
              class="w-2/3"
              draggable="false"
            />
          </div>
          <button
            class="bg-yellow-500 hover:bg-yellow-600 text-black px-4 py-2 rounded"
            @click="login"
          >
            Login
          </button>
        </div>
      </div>
    </div>
  </div>
</template>