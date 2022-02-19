<template>
    <div class="q-py-md">
        <q-table
            class="bg-blue-grey-1"
            title="Latest data"
            :rows="latestData"
            :columns="cols"
        >
        <template #body-cell-isOnline="{value}">
          <q-td class="text-right">
            <q-icon size="1.5rem" :color="value ? 'green' : 'red'" :name="value ? 'check' : 'error'"/>
          </q-td>
        </template>
      </q-table>
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
            { name: 'isOnline', label: 'Connection', field: 'isOnline' }
        ]
    },
    computed: {
      latestData () {
        return this.$store.state['master'].latest
      }
    }
  });
</script>