<template>
    <div class="bg-blue-grey-1 q-pa-md shadow-1 rounded-borders">
        <apexchart width="100%" height="400px" type="line" :options="completeStats.options" :series="completeStats.series"></apexchart>
    </div>
</template>

<script>
  import { defineComponent } from "vue";

  export default defineComponent({
    name: "PageGraph",
    computed: {
      completeStats() {
        const chart = this.$store.getters['master/parsedStats']

        return {
          options: {
            chart: {
              id: 'ping-stats',
              animations: {
                enabled: false
              }
            },
            stroke: {
              width: 2,
              curve: "smooth"
            },
            xaxis: {
              categories: chart.categories
            }
          },
          series: chart.series
        }
      }
    }
  });
</script>