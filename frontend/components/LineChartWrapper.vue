<script setup>
import { ref, onMounted } from "vue";
import LineChart from "~/components/LineChart.vue";

const apiData = ref([]); // API data storage
const filteredApiData = ref([]); // Filtered data based on user input
const filterTime = ref(""); // User-inputted time
const filterEndpoint = ref(""); // User-inputted endpoint
const selectedApi = ref(null); // Currently selected API
const user = ref("");

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
        text: "Time Intervals",
      },
    },
    y: {
      title: {
        display: true,
        text: "Active APIs",
      },
    },
  },
};

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

      // Initialize filteredApiData with the fetched data
      filteredApiData.value = apiData.value;

      // Process the data for the chart
      processData();
    } catch (error) {
      console.error("Error fetching data:", error);
    }
  }
}

function processData() {
  const intervals = {};
  const intervalDuration = 60 * 1000; // 1 minute

  apiData.value.forEach((api) => {
    const reqIn = new Date(api.req_in).getTime();
    const reqOut = new Date(api.req_out).getTime();

    for (let t = reqIn; t <= reqOut; t += intervalDuration) {
      const interval = new Date(t).toISOString().slice(0, 16); // Format as 'YYYY-MM-DDTHH:mm'
      intervals[interval] = (intervals[interval] || 0) + 1;
    }
  });

  const labels = Object.keys(intervals).sort();
  const values = labels.map((label) => intervals[label]);

  console.log(labels, values); // Debugging
  chartData.value = {
    labels,
    datasets: [
      {
        label: "Active APIs",
        data: values,
        borderColor: "#4A90E2",
        backgroundColor: "rgba(74, 144, 226, 0.2)",
        tension: 0.4, // Smooth line
      },
    ],
  };

  console.log(chartData.value); // Debugging
}

onMounted(() => {
  fetchData(); // Fetch the data when the component is mounted
});
</script>

<template>
  <div
    class="min-w-full min-h-full p-4 bg-white rounded-lg shadow-md border-blue-600 border-4"
  >
    <LineChart
      v-if="chartData"
      :chartData="chartData"
      :chartOptions="chartOptions"
    />
  </div>
</template>
