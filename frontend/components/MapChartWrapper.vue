<template>
    
      <MapChart :chartData="chartData" :chartOptions="chartOptions" />
    
  </template>
  
  <script>
  import MapChart from "~/components/MapChart.vue";
  
  export default {
    components: {
      MapChart,
    },
    data() {
      return {
        chartData: {
          datasets: [
            {
              label: "Random Points on US Map",
              data: this.generateRandomPoints(100), // Generate 100 random points
              backgroundColor: "#FF6384",
            },
          ],
        },
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
                text: "Longitude",
              },
              min: -125, // Westernmost point of the US
              max: -66,  // Easternmost point of the US
            },
            y: {
              title: {
                display: true,
                text: "Latitude",
              },
              min: 25, // Southernmost point of the US
              max: 50, // Northernmost point of the US
            },
          },
        },
      };
    },
    methods: {
      generateRandomPoints(count) {
        const points = [];
        for (let i = 0; i < count; i++) {
          const longitude = (Math.random() * (-66 - -125)) + -125; // Longitude range
          const latitude = (Math.random() * (50 - 25)) + 25;       // Latitude range
          points.push({ x: longitude, y: latitude });
        }
        return points;
      },
    },
  };
  </script>
  