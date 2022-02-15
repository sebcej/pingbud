<template>
    <div class="q-py-md">
        <q-table
            class="bg-blue-grey-1"
            title="Latest data"
            :rows="latestData"
            :columns="cols"
        />
    </div>
</template>

<script>
  import { defineComponent } from "vue";

  export default defineComponent({
    name: "TopBar",
    created () {
        this.cols = [
            { name: 'time', label: 'Time', field: 'time', sortable: true, format: val => new Date(val*1000).toLocaleString() },
            { name: 'avg', label: 'Average', field: 'avg', sortable: true },
            { name: 'max', label: 'Max', field: 'max', sortable: true },
            { name: 'min', label: 'Min', field: 'min', sortable: true },
            { name: 'isOnline', label: 'Is Online', field: 'isOnline', sortable: true }
        ]
    },
    computed: {
      latestData () {
        return this.$store.state['master'].latest
      }
    }
  });
</script>