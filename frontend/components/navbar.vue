<script setup>
import { UserCircleIcon } from "@heroicons/vue/24/solid";
import { ArrowLeftStartOnRectangleIcon } from "@heroicons/vue/24/solid";
const loggedIn = ref(true);
const token = useCookie("token");
if (!token.value) {
  loggedIn.value = false;
}
const logout = async () => {
  navigateTo("http://localhost:8080/api/auth/google/logout", {
    external: true,
  });
};
</script>

<template>
  <div
    class="relative text-white w-screen p-2 bg-gray-800 border border-b-yellow-300 border-l-transparent border-r-transparent border-t-transparent border-b-transparent"
  >
    <div class="grid grid-cols-3 relative z-10">
      <div>
        <h1 class="text-2xl font-bold uppercase space-x-1 text-yellow-200">
          <span>P</span><span>I</span><span>N</span><span>E</span><span>A</span
          ><span>P</span><span>P</span><span>L</span><span>E</span>
        </h1>
      </div>
      <div></div>
      <div class="flex justify-end">
        <button
          v-if="!loggedIn"
          class="bg-yellow-500 text-white px-4 py-2 rounded"
        >
          <UserCircleIcon class="h-6 w-6" />
        </button>
        <button
          v-if="loggedIn"
          class="bg-yellow-500 hover:bg-yellow-600 text-black px-4 py-2 rounded"
          @click="logout"
        >
          <ArrowLeftStartOnRectangleIcon class="h-6 w-6" />
          <!-- <span>Logout</span> -->
        </button>
      </div>
    </div>
  </div>
</template>
