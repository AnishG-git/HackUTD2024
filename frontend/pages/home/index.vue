<script setup>
import LineChartWrapper from "~/components/LineChartWrapper.vue";
import MapChartWrapper from "~/components/MapChartWrapper.vue";
import PieChartWrapper from "~/components/PieChartWrapper.vue";
import BarChartWrapper from "~/components/BarChartWrapper.vue";
import { PinataSDK } from "pinata";

const pinata = new PinataSDK({
    pinataJwt: "ZXlKaGJHY2lPaUpJVXpJMU5pSXNJblI1Y0NJNklrcFhWQ0o5LmV5SjFjMlZ5U1c1bWIzSnRZWFJwYjI0aU9uc2lhV1FpT2lKaE1qbGxPVEl3WVMxak1EbGpMVFEzWlRJdE9USmtZeTA1WmpZeVptUTBaVFEzT1RBaUxDSmxiV0ZwYkNJNkluSmxZV05vZEhWemFHRnlkMkZ1YVVCbmJXRnBiQzVqYjIwaUxDSmxiV0ZwYkY5MlpYSnBabWxsWkNJNmRISjFaU3dpY0dsdVgzQnZiR2xqZVNJNmV5SnlaV2RwYjI1eklqcGJleUprWlhOcGNtVmtVbVZ3YkdsallYUnBiMjVEYjNWdWRDSTZNU3dpYVdRaU9pSkdVa0V4SW4wc2V5SmtaWE5wY21Wa1VtVndiR2xqWVhScGIyNURiM1Z1ZENJNk1Td2lhV1FpT2lKT1dVTXhJbjFkTENKMlpYSnphVzl1SWpveGZTd2liV1poWDJWdVlXSnNaV1FpT21aaGJITmxMQ0p6ZEdGMGRYTWlPaUpCUTFSSlZrVWlmU3dpWVhWMGFHVnVkR2xqWVhScGIyNVVlWEJsSWpvaWMyTnZjR1ZrUzJWNUlpd2ljMk52Y0dWa1MyVjVTMlY1SWpvaVpEaGhPREk1TVdRME9HTmpZelV5TldZME9ERWlMQ0p6WTI5d1pXUkxaWGxUWldOeVpYUWlPaUkyWVdVeVpHRmhOR0kwWVdVMFlqSXlZVFkyT1dNMVpqWmpNMk16T0RBd01UYzRNbUZrT0dVMFl6SXlNRGt6TjJZeFpHVmxZMkprT0dKaE1qSTVaamcxSWl3aVpYaHdJam94TnpZek16UTNOalV5ZlEubTFVa3J5TmFGWjhRNS1UMzR0bXlrSmFPd0c4dnRhR1NUdmZWWWJvOEFOOA==",
    pinataGateway: "moccasin-improved-smelt-48.mypinata.cloud",
});

const logout = async () => {
  navigateTo("http://localhost:8080/api/auth/google/logout", {
    external: true,
  });
};

const wap = async () => {
  console.log("Wap");
  console.log(pinata);
  const data = await pinata.testAuthentication();
  // gateways.get("bafkreicd3p3hcsu7633nqcmnrlxohyvsfxtwvdgzxwp7uornwf66j6debu")
    console.log(data.String());


};

const currentIndex = ref(0);
const charts = [
  LineChartWrapper,
  MapChartWrapper,
  PieChartWrapper,
  BarChartWrapper,
  // Add more components if needed
];

const next = () => {
  currentIndex.value = (currentIndex.value + 1) % charts.length;
};

const prev = () => {
  currentIndex.value = (currentIndex.value - 1 + charts.length) % charts.length;
};


</script>

<template>
  <div class="flex flex-col w-screen h-screen bg-gray-900 text-white">
    <navbar />
    <h1 class="p-6 pt-10 text-4xl text-yellow-400 font-semibold">API Analysis</h1>
    <div class="w-full h-2/5 flex flex-wrap">
      <div class="w-2/3 min-h-full p-10 pr-5 pt-0">
        <LineChartWrapper/>
      </div>
      <div class="w-1/3 min-h-full p-10 pl-5 pt-0">
        <PieChartWrapper/>
      </div>
    </div>

    <div class="w-full h-2/5 flex flex-wrap pt-0">
      <div class="w-2/3 min-h-full p-10 pt-0 pr-5">
        <BarChartWrapper/>
      </div>
      <div class="w-1/3 min-h-full p-10 pt-0 pl-5">
        <MapChartWrapper/>
      </div>
    </div>
    
    
    
    
    
    
  </div>
</template>

