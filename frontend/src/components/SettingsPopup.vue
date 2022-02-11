<template>
    <q-dialog v-model="visible">
      <q-card>
        <q-card-section class="text-h4">
          Settings
        </q-card-section>
        <q-card-section>
          {{settings}}
        </q-card-section>
        <q-card-section>
          <q-btn color="primary" size="md">Save</q-btn>
        </q-card-section>
      </q-card>
    </q-dialog>
</template>

<script>
  import { defineComponent } from "vue";

  export default defineComponent({
    name: "SettingsPopup",
    data () {
      return {
        visible: false
      }
    },
    methods: {
      async getData () {
        await this.$store.dispatch('master/getSettings')
      },
      async show() {
        await this.getData()
        this.visible = true
      },
      hide() {
        this.visible = false
      }
    },
    computed: {
      settings () {
        return this.$store.state['master'].settings
      }
    }
  });
</script>