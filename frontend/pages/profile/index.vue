<script setup>
import {
  ArrowUturnLeftIcon,
  ArrowRightStartOnRectangleIcon,
} from "@heroicons/vue/24/solid";

const show_key = ref("false");

function parseJwt(token) {
  const arrayToken = token.split(".");
  const tokenPayload = JSON.parse(atob(arrayToken[1]));
  return tokenPayload;
}

const token = useCookie("jwt");
const user = ref("");
if (token.value) {
  user.value = parseJwt(token.value).sdk_key;
  console.log(user);
}

const show_key_toggle = () => {
  navigator.clipboard.writeText(user.value);
  show_key.value = !show_key.value;
};

const logout = async () => {
//   setLoggedIn(false);
  // console.log(loggedIn);
  await navigateTo("http://localhost:8080/api/auth/google/logout", {
    external: true,
  });
  // await navigateTo("/"); // Redirect to home page after logout
};

const nav = (path) => {
  navigateTo(path);
};
</script>

<template>
  <div class="flex flex-col w-screen h-screen bg-gray-900 text-white">
    <navbar />
    <div class="flex p-3">
      <span class="w-1/2">
        <button
          class="rounded p-2 bg-yellow-200 text-black"
          @click="nav('/home')"
        >
          <ArrowUturnLeftIcon class="h-6 w-6" />
        </button>
      </span>
      <span class="w-1/2 flex justify-end">
        <button class="rounded p-2 bg-yellow-200 text-black" @click="logout">
          <ArrowRightStartOnRectangleIcon class="h-6 w-6" />
        </button>
      </span>
    </div>
    <div class="flex justify-center items-center h-full">
      <button
        v-if="show_key"
        @click="show_key_toggle"
        class="rounded bg-yellow-200 text-black p-3 border border-gray-100 hover:bg-yellow-300"
      >
        Show SDK Key
      </button>
      <button
        v-if="!show_key"
        @click="show_key_toggle"
        class="rounded w-3/4 p-3 overflow-auto bg-green-400 hover:bg-green-500 text-wrap text-black"
      >
        {{ user }}
      </button>
    </div>
  </div>
</template>