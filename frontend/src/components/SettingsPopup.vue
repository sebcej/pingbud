<template>
    <q-dialog v-model="visible">
      <q-card>
        <q-card-section class="text-h4">
          Settings
        </q-card-section>
        <q-card-section style="min-width: 500px;">
          <q-form @submit="submit">
            <q-input type="number" label="N. of pings" hint="Number of pings performed. The final value will be an average" :rules="[requiredField]" v-model.number="form.pingcount"/>
            <q-input label="Frequency" hint="Crontab notation with seconds" :rules="[validCron]" v-model="form.pingcron"/>
            <q-input label="Ip to call" :rules="[requiredField, validIp]" v-model="form.pingroute"/>
            <q-input type="number" label="Data retention (days)" :rules="[requiredField]" v-model.number="form.retention"/>
            <q-input type="number" label="Timeout (seconds)" hint="Time before the connection is considered faulty" :rules="[requiredField]" v-model.number="form.timeout"/>
           <div class="q-pt-sm">
              <q-checkbox label="Privilegied mode" v-model="form.privilegedmode" dense/>
              <div class="q-field__bottom q-pl-none">
                In some systems privileged mode (root) is necessary. See <a href='https://github.com/go-ping/ping#supported-operating-systems'>here</a> why
              </div>
           </div>
            <div class="q-pt-md">
              <q-btn type="submit" color="primary" size="md">Save</q-btn>
            </div>
          </q-form>
        </q-card-section>
      </q-card>
    </q-dialog>
</template>

<script>
  import { defineComponent } from "vue";
  import { validCron, requiredField, validIp } from "src/common/validators"

  export default defineComponent({
    name: "SettingsPopup",
    data () {
      return {
        visible: false,
        form: {}
      }
    },
    methods: {
      validCron,
      requiredField,
      validIp,
      async getData () {
        await this.$store.dispatch('master/getSettings')
      },
      async show() {
        await this.getData()
        this.form = {...this.settings}
        this.visible = true
      },
      hide() {
        this.visible = false
      },
      async submit() {
        await this.$store.dispatch('master/saveSettings', this.form)
        await this.getData()
      }
    },
    computed: {
      settings () {
        return this.$store.state['master'].settings
      }
    }
  });
</script>