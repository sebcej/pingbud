<template>
    <div class="q-py-md">
        <q-table
            class="bg-blue-grey-1"
            title="Registered errors"
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
            { name: 'time', label: 'Time', field: 'time', sortable: true, format: val => new Date(val*1000).toLocaleString() }
        ]
    },
    computed: {
      latestData () {
        return this.$store.state['master'].errors
      }
    }
  });
</script>