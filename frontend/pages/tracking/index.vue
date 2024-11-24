<template>
  <div class="flex flex-col w-screen h-screen bg-gray-900 overflow-hidden">
    <navbar />

    <!-- Main Container -->
    <div class="flex h-full">
      <!-- Sidebar -->
      <div class="h-full w-1/3 p-4 overflow-y-auto shadow-lg">
        <h2 class="text-lg font-bold mb-4 text-white">API List</h2>

        <!-- Filter Section -->
        <div class="mb-4">
          <label for="filter-time" class="text-white font-bold"
            >Filter by Request Time:</label
          >
          <input
            id="filter-time"
            type="datetime-local"
            v-model="filterTime"
            class="w-full p-2 rounded-xl py-4 border mb-2"
          />

          <label for="filter-endpoint" class="text-white font-bold block"
            >Filter by Endpoint:</label
          >
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
            <div class="w-1/2 overflow-hidden">
              <p class="font-bold">{{ api.endpoint }}</p>
              <p>{{ formatTimestamp(api.req_in) }}</p>
            </div>

            <div
              class="bg-red-200 border-red-600 border-2 text-red-900 rounded-full px-2 flex items-center justify-center text-md"
            >
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
          <p>
            <span class="font-bold">Endpoint:</span> {{ selectedApi.endpoint }}
          </p>
          <p>
            <span class="font-bold">Request In:</span>
            {{ formatTimestamp(selectedApi.req_in) }}
          </p>
          <p>
            <span class="font-bold">Request Out:</span>
            {{ formatTimestamp(selectedApi.req_out) }}
          </p>
          <p><span class="font-bold">Method:</span> {{ selectedApi.method }}</p>
          <p>
            <span class="font-bold">Raw Request ID:</span>
            {{ selectedApi.raw_req_id }}
          </p>
          <p>
            <span class="font-bold">Raw Response ID:</span>
            {{ selectedApi.raw_resp_id }}
          </p>
          <p>
            <span class="font-bold">Response States:</span>
            {{ selectedApi.status_code }}
          </p>
          <p>
            <span class="font-bold">Latency:</span>
            {{ selectedApi.latency }}
          </p>
        </div>
        <div v-else>
          <p class="text-gray-500">Select an API to view its details.</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
//import { useCookie } from "vue-composable"; // Ensure you are using a compatible library
//import data from "../home/data.json"; // Placeholder for additional data if needed

// Reactive variables
const apiData = ref([]); // API data storage
const filteredApiData = ref([]); // Filtered data based on user input
const filterTime = ref(""); // User-inputted time
const filterEndpoint = ref(""); // User-inputted endpoint
const selectedApi = ref(null); // Currently selected API
const user = ref("");

// Helper function to parse JWT token
function parseJwt(token) {
  const arrayToken = token.split(".");
  const tokenPayload = JSON.parse(atob(arrayToken[1]));
  return tokenPayload;
}

// Fetch API data
async function fetchData() {
  const token = useCookie("jwt");
  if (token.value) {
    user.value = parseJwt(token.value).sdk_key;

    try {
      const response = await fetch(
        `http://localhost:8080/api/fetch-data/${user.value}`,
        {
          method: "GET",
          headers: {
            "Content-Type": "application/json",
          },
        }
      );

      if (!response.ok) {
        throw new Error(`HTTP error! Status: ${response.status}`);
      }

      const data = await response.json();
      console.log("Fetched Data:", data);

      // Assign the fetched data to apiData
      apiData.value = data;

      // Optionally, initialize filteredApiData
      filteredApiData.value = apiData.value;
    } catch (error) {
      console.error("Error fetching data:", error);
    }
  }
}

// Filter and utility functions
function formatTimestamp(timestamp) {
  return new Date(timestamp).toLocaleString();
}

async function handleButtonClick(api) {
  selectedApi.value = api; // Set the selected API
  const token = useCookie("jwt");
  console.log(token.value);
  console.log(api.raw_req_id);
  const options = {
    method: "POST",
    headers: {
      "Authorization": `Bearer ${token.value}`,
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      url: `https://moccasin-improved-smelt-48.mypinata.cloud/files/${api.raw_req_id}`,
      date: 1731847517,
      expires: 600,
      method: "GET",
    }),
  };

  try {
    const signResponse = await fetch(
      "https://api.pinata.cloud/v3/files/sign",
      options
    );
    const signData = await signResponse.json();
    console.log(signData);

    const fileResponse = await fetch(signData.data, options);
    const fileData = await fileResponse.json();
    console.log(fileData);
  } catch (err) {
    console.error(err);
  }
}

function applyFilter() {
  const filterTimestamp = filterTime.value
    ? new Date(filterTime.value).getTime()
    : null;
  const endpointFilter = filterEndpoint.value.toLowerCase();

  filteredApiData.value = apiData.value.filter((api) => {
    const reqInTimestamp = new Date(api.req_in).getTime();
    const matchesTime = filterTimestamp
      ? reqInTimestamp > filterTimestamp
      : true;
    const matchesEndpoint = endpointFilter
      ? api.endpoint.toLowerCase().includes(endpointFilter)
      : true;

    return matchesTime && matchesEndpoint;
  });
}

// Fetch data when component is mounted
onMounted(() => {
  fetchData();
});
</script>
