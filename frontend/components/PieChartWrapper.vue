<script setup>
import { ref, reactive, onMounted } from "vue";
import PieChart from "~/components/PieChart.vue";
//import { useCookie } from "vue-composable"; // Ensure you have this dependency

const chartData = ref(null);
const chartOptions = reactive({
  responsive: true,
  plugins: {
    legend: {
      display: true,
      position: "top",
    },
  },
});

const generateColors = (count) => {
  const colors = [];
  for (let i = 0; i < count; i++) {
    colors.push(`hsl(${(360 / count) * i}, 70%, 60%)`);
  }
  return colors;
};

const fetchData = async () => {
  const token = useCookie("jwt");
  const user = ref("");

  if (token.value) {
    user.value = JSON.parse(atob(token.value.split(".")[1])).sdk_key;

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
      return data;
    } catch (error) {
      console.error("Error fetching data:", error);
      return [];
    }
  }
  return [];
};

const processData = (data) => {
  const endpointCounts = {};

  // Count occurrences of each endpoint
  data.forEach((entry) => {
    const endpoint = entry.endpoint;
    endpointCounts[endpoint] = (endpointCounts[endpoint] || 0) + 1;
  });

  // Prepare data for Chart.js
  const labels = Object.keys(endpointCounts);
  const values = Object.values(endpointCounts);

  return {
    labels,
    datasets: [
      {
        data: values,
        backgroundColor: generateColors(labels.length),
        hoverOffset: 4,
      },
    ],
  };
};

onMounted(async () => {
  const data = await fetchData();
  chartData.value = processData(data);

  // For debugging
  console.log("Chart Data:", chartData.value);
});
</script>

<template>
  <div
    class="min-w-full min-h-full p-4 bg-white rounded-lg shadow-md border-blue-600 border-4"
  >
    <PieChart
      v-if="chartData"
      :chartData="chartData"
      :chartOptions="chartOptions"
    />
  </div>
</template>
