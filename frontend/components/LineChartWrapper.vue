<template>
   
        <div class="min-w-full min-h-full p-4 bg-white rounded-lg shadow-md border-blue-600 border-4">
        <LineChart :chartData="chartData" :chartOptions="chartOptions" />
      </div>
    
  </template>
  
  <script>
  import LineChart from "~/components/LineChart.vue";
  import data from "../pages/home/data.json";
  
  export default {
    components: {
      LineChart,
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
        },
      };
    },
    methods: {
      processData() {
        const intervals = {};
        const intervalDuration = 60 * 1000; // 1 minute
  
        data.forEach((api) => {
          const reqIn = new Date(api.req_in).getTime();
          const reqOut = new Date(api.req_out).getTime();
  
          for (let t = reqIn; t <= reqOut; t += intervalDuration) {
            const interval = new Date(t).toISOString().slice(0, 16); // Format as 'YYYY-MM-DDTHH:mm'
            intervals[interval] = (intervals[interval] || 0) + 1;
          }
        });
  
        const labels = Object.keys(intervals).sort();
        const values = labels.map((label) => intervals[label]);
  
        return { labels, values };
      },
    },
    created() {
      const { labels, values } = this.processData();
  
      this.chartData = {
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
    },
  };
  </script>
  