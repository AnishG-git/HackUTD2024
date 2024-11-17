<template>
  
  <div class="  h-full">
    <BarChart :chartData="chartData" :chartOptions="chartOptions" />
  </div>
    
  
</template>

<script>
import BarChart from "~/components/BarChart.vue";
import data from "../pages/home/data.json";

export default {
  components: {
    BarChart,
  },
  data() {
    return {
      chartData: null,
      chartOptions: {
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
      },
    };
  },
  methods: {
    processData() {
      const endpointStats = {};

      // Count successes and failures for each endpoint
      data.forEach((entry) => {
        const endpoint = entry.endpoint;
        const status = parseInt(entry.resp_states.split(" ")[0], 10); // Extract status code

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
    },
  },
  created() {
    this.chartData = this.processData();
  },
};
</script>
