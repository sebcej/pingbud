<template>
  <q-page>
    <TopBar/>
    <Graph/>
    <div class="row">
      <div class="col">
        <Latest/>
      </div>
      <div class="col"></div>
    </div>
  </q-page>
</template>

<script>
  import { defineComponent } from "vue";
  import TopBar from 'components/Home/TopBar'
  import Graph from "components/Home/Graph.vue"
  import Latest from "components/Home/LatestTable.vue";

  export default defineComponent({
    name: "PageIndex",
    beforeMount() {
      this.$store.dispatch('master/getSettings')
      setInterval(() => {
        this.$store.dispatch('master/getStats')
      }, 5000)
      this.$store.dispatch('master/getStats')
    },
    components: {
      TopBar,
      Graph,
      Latest
    }
  });
</script>
