<template>
  <q-page>
    <TopBar/>
    <Graph/>
    <div class="row q-col-gutter-md q-mb-md">
      <div class="col-12 col-md-9">
        <Latest/>
      </div>
      <div class="col-12 col-md-3">
        <Errors/>
      </div>
    </div>
  </q-page>
</template>

<script>
  import { defineComponent } from "vue";
  import TopBar from 'components/Home/TopBar'
  import Graph from "components/Home/Graph.vue"
  import Latest from "components/Home/LatestTable.vue";
  import Errors from "components/Home/ErrorsTable.vue";

  export default defineComponent({
    name: "PageIndex",
    beforeMount() {
      this.$store.dispatch('master/getSettings')
      setInterval(() => {
        this.$store.dispatch('master/getStats')
      }, 60000)
      this.$store.dispatch('master/getStats')
    },
    components: {
      TopBar,
      Graph,
      Latest,
      Errors
    }
  });
</script>
