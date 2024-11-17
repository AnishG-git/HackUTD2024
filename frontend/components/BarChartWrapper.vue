<template>
  <div
    class="min-w-full min-h-full p-4 bg-white rounded-lg shadow-md border-blue-600 border-4"
  >
    <BarChart
      v-if="chartData"
      :chartData="chartData"
      :chartOptions="chartOptions"
    />
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import BarChart from "~/components/BarChart.vue";
import data from "../pages/home/data.json";
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
      console.log("processed Data:", processData(data));
      chartData.value = processData(data);
      // console.log("hehe:", processData(data));

      // Assign the fetched data to apiData

      // Optionally, initialize filteredApiData
      filteredApiData.value = apiData.value;
    } catch (error) {
      console.error("Error fetching data:", error);
    }
  }
}

const chartData = ref(null);
const chartOptions = {
  responsive: true,
  plugins: {
    legend: {
      display: true,
      position: "top",
    },
  },
  scales: {
    x: {
      title: {
        display: true,
        text: "Endpoints",
      },
    },
    y: {
      title: {
        display: true,
        text: "Count",
      },
    },
  },
};

function processData(data) {
  const endpointStats = {};

  // Count successes and failures for each endpoint
  data.forEach((entry) => {
    const endpoint = entry.endpoint;
    const status = parseInt(entry.status_code); // Extract status code

    if (!endpointStats[endpoint]) {
      endpointStats[endpoint] = { success: 0, failure: 0 };
    }

    if (status === 200 || status === 201) {
      endpointStats[endpoint].success += 1;
    } else {
      endpointStats[endpoint].failure += 1;
    }
  });

  // Prepare data for Chart.js
  const labels = Object.keys(endpointStats);
  const successData = labels.map((label) => endpointStats[label].success);
  const failureData = labels.map((label) => endpointStats[label].failure);

  return {
    labels,
    datasets: [
      {
        label: "Successful Responses",
        data: successData,
        backgroundColor: "rgba(54, 162, 235, 0.6)", // Blue
        borderColor: "rgba(54, 162, 235, 1)",
        borderWidth: 1,
      },
      {
        label: "Unsuccessful Responses",
        data: failureData,
        backgroundColor: "rgba(255, 99, 132, 0.6)", // Red
        borderColor: "rgba(255, 99, 132, 1)",
        borderWidth: 1,
      },
    ],
  };
}

onMounted(() => {
  console.log("haha:", data);

  fetchData();
});
</script>
