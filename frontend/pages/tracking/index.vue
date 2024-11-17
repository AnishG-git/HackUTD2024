<template>
    <div class="h-full w-64 bg-gray-100 p-4 overflow-y-auto shadow-lg">
      <h2 class="text-lg font-bold mb-4">API List</h2>
      <ul class="space-y-4">
        <li
          v-for="(api, index) in apiData"
          :key="index"
          class="p-4 bg-white rounded-lg shadow-md"
        >
          <p><span class="font-bold">Endpoint:</span> {{ api.endpoint }}</p>
          <p><span class="font-bold">Request In:</span> {{ formatTimestamp(api.req_in) }}</p>
          <p><span class="font-bold">Method:</span> {{ api.method }}</p>
          <button
            class="mt-2 px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600"
            @click="handleButtonClick(api)"
          >
            View Details
          </button>
        </li>
      </ul>
    </div>
  </template>
  
  <script>
  import { defineComponent } from "vue";
  import data from "../home/data.json";
  
  export default defineComponent({
    name: "ApiList",
    data() {
      return {
        apiData: [],
      };
    },
    methods: {
      formatTimestamp(timestamp) {
        return new Date(timestamp).toLocaleString();
      },
      handleButtonClick(api) {
        console.log("Button clicked for API:", api);
        // Perform additional actions, such as navigating to another page or showing details.
      },
    },
    created() {
      // Retrieve data from data.json
      this.apiData = data;
    },
  });
  </script>