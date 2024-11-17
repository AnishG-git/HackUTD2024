<template>
  <div class="flex flex-col w-screen h-screen bg-gray-900">
    <navbar />

    <!-- Main Container -->
    <div class="flex h-full">
      <!-- Sidebar -->
      <div class="h-full w-1/3 p-4 overflow-y-auto shadow-lg ">
        <h2 class="text-lg font-bold mb-4 text-white">API List</h2>

        <!-- Filter Section -->
        <div class="mb-4">
          <label for="filter-time" class=" text-white font-bold">Filter by Request Time:</label>
          <input
            id="filter-time"
            type="datetime-local"
            v-model="filterTime"
            class="w-full p-2 rounded-xl py-4 border mb-2"
          />

          <label for="filter-endpoint" class=" text-white font-bold block">Filter by Endpoint:</label>
          <input
            id="filter-endpoint"
            type="text"
            v-model="filterEndpoint"
            placeholder="Enter endpoint..."
            class="w-full p-2 rounded-xl py-4 border"
          />

          <button
            @click="applyFilter"
            class="mt-2 px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600"
          >
            Apply Filter
          </button>
        </div>

        <!-- API List -->
        <ul class="space-y-4">
          <li
            v-for="(api, index) in filteredApiData"
            :key="index"
            class="p-4 bg-white rounded-lg shadow-md flex flex-wrap justify-between"
          >
            <div>
              <p class="font-bold">{{ api.endpoint }}</p>
              <p>{{ formatTimestamp(api.req_in) }}</p>
            </div>

            <div class="bg-red-200 border-red-600 border-2 text-red-900 rounded-full px-2 flex items-center justify-center text-md">
              <p>{{ api.method }}</p>
            </div>

            <button
              class="mt-2 px-4 py-2 bg-blue-500 text-white rounded-full hover:bg-blue-600"
              @click="handleButtonClick(api)"
            >
              View
            </button>
          </li>
        </ul>
      </div>

      <!-- Details Section -->
      <div class="h-full w-2/3 p-8 text-white">
        <h2 class="text-2xl font-bold mb-4">API Details</h2>
        <div v-if="selectedApi" class="space-y-4">
          <p><span class="font-bold">Endpoint:</span> {{ selectedApi.endpoint }}</p>
          <p><span class="font-bold">Request In:</span> {{ formatTimestamp(selectedApi.req_in) }}</p>
          <p><span class="font-bold">Request Out:</span> {{ formatTimestamp(selectedApi.req_out) }}</p>
          <p><span class="font-bold">Method:</span> {{ selectedApi.method }}</p>
          <p><span class="font-bold">Raw Request ID:</span> {{ selectedApi.raw_req_id }}</p>
          <p><span class="font-bold">Raw Response ID:</span> {{ selectedApi.raw_resp_id }}</p>
          <p><span class="font-bold">Response States:</span> {{ selectedApi.resp_states }}</p>
        </div>
        <div v-else>
          <p class="text-gray-500">Select an API to view its details.</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { defineComponent } from "vue";
import data from "../home/data.json";

export default defineComponent({
  name: "ApiList",
  data() {
    return {
      apiData: [], // Original data from JSON
      filteredApiData: [], // Filtered data based on user input
      filterTime: "", // User-inputted time
      filterEndpoint: "", // User-inputted endpoint
      selectedApi: null, // Currently selected API
    };
  },
  methods: {
    formatTimestamp(timestamp) {
      return new Date(timestamp).toLocaleString();
    },
    handleButtonClick(api) {
      this.selectedApi = api; // Set the selected API
    },
    applyFilter() {
      const filterTimestamp = this.filterTime ? new Date(this.filterTime).getTime() : null;
      const endpointFilter = this.filterEndpoint.toLowerCase();

      this.filteredApiData = this.apiData.filter((api) => {
        const reqInTimestamp = new Date(api.req_in).getTime();
        const matchesTime = filterTimestamp ? reqInTimestamp > filterTimestamp : true;
        const matchesEndpoint = endpointFilter
          ? api.endpoint.toLowerCase().includes(endpointFilter)
          : true;

        return matchesTime && matchesEndpoint;
      });
    },
  },
  created() {
    // Retrieve data from data.json
    this.apiData = data;
    this.filteredApiData = this.apiData; // Initialize with all data
  },
});
</script>
