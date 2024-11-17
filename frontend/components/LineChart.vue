<template>
    <div class="p-4 bg-white rounded-lg shadow-md">
      <canvas ref="chartCanvas"></canvas>
    </div>
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
          options: props.chartOptions,
        });
      });
  
      return {
        chartCanvas,
      };
    },
  });
  </script>
  