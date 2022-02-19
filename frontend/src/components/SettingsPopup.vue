<template>
    <q-dialog v-model="visible">
      <q-card>
        <q-card-section class="text-h4">
          Settings
        </q-card-section>
        <q-card-section style="min-width: 500px;">
          <q-input type="number" label="N. of pings" v-model="form.pingcount"/>
          <q-input label="Frequency" v-model="form.pingcron"/>
          <q-input label="Ip to call" v-model="form.pingroute"/>
          <q-input type="number" label="Data retention (days)" v-model="form.retention"/>
          <q-input type="number" label="Timeout (seconds)" v-model="form.timeout"/>
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
        visible: false,
        form: {}
      }
    },
    methods: {
      async getData () {
        await this.$store.dispatch('master/getSettings')
      },
      async show() {
        await this.getData()
        this.form = this.settings
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