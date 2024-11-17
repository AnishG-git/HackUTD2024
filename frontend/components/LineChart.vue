<template>
    
    <canvas ref="chartCanvas"></canvas>
    
  </template>
  
  <script>
  import { defineComponent, ref, onMounted } from "vue";
  import { Chart, registerables } from "chart.js";
  
  // Register Chart.js components
  Chart.register(...registerables);
  
  export default defineComponent({
    props: {
      chartData: {
        type: Object,
        required: true,
      },
      chartOptions: {
        type: Object,
        required: false,
        default: () => ({}),
      },
    },
    setup(props) {
      const chartCanvas = ref(null);
  
      onMounted(() => {
        new Chart(chartCanvas.value, {
          type: "line",
          data: props.chartData,
          options: {
            ...props.chartOptions,
            responsive: true,
            maintainAspectRatio: false, // Allows the chart to stretch to fill the container
          },
        });
      });
  
      return {
        chartCanvas,
      };
    },
  });
  </script>
  