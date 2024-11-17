<template>
    
      <PieChart :chartData="chartData" :chartOptions="chartOptions" />
    
  </template>
  
  <script>
  import PieChart from "~/components/PieChart.vue";
  import data from "../pages/home/data.json";
  
  export default {
    components: {
      PieChart,
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
        },
      };
    },
    methods: {
      processData() {
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
              backgroundColor: this.generateColors(labels.length),
              hoverOffset: 4,
            },
          ],
        };
      },
      generateColors(count) {
        const colors = [];
        for (let i = 0; i < count; i++) {
          colors.push(`hsl(${(360 / count) * i}, 70%, 60%)`);
        }
        return colors;
      },
    },
    created() {
      this.chartData = this.processData();
    },
  };
  </script>
  