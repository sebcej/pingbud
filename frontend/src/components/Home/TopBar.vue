<template>
    <div class="row q-col-gutter-md q-mb-md">
      <div class="col-md-4 col-6">
        <q-card class="bg-positive">
          <q-card-section>
            <div class="text-h6">Latence</div>
          </q-card-section>

          <q-card-section class="text-center text-h4">
            {{round(stats.avgPing)}}ms
          </q-card-section>
        </q-card>
      </div>
      <div class="col-md-4 col-6">
          <q-card class="bg-info">
            <q-card-section>
              <div class="text-h6">Jitter</div>
            </q-card-section>

            <q-card-section class="text-center text-h4">
              {{round(stats.avgJitter)}}ms
            </q-card-section>
          </q-card>
        </div>
        <div class="col-md-1 col-6">
          <q-card class="bg-negative">
            <q-card-section>
              <div class="text-h6">Errors</div>
            </q-card-section>

            <q-card-section class="text-center text-h4">
              {{stats.errors || '-'}}
            </q-card-section>
          </q-card>
        </div>
        <div class="col-md-3 col-6">
          <SettingsPopup ref="settings"/>
          <q-card>
            <q-card-section>
              <div class="text-h6">Control panel</div>
            </q-card-section>

            <q-card-section class="row">
              <div class="col-4">
                <q-icon size="2.4rem" :name="enabledStatus.icon" :color="enabledStatus.color" class="cursor-pointer" @click="toggleEnabled"/>
                <q-icon @click="$refs.settings.show()" size="2.4rem" name="settings" class="q-ml-lg cursor-pointer"/>
              </div>
              <div class="col-8">
                <q-input v-model="filter" readonly dense placeholder="Last 24h">
                  <template v-slot:append>
                    <q-icon name="event" class="cursor-pointer">
                      <q-popup-proxy ref="qDateProxy" cover @hide="filterEv" transition-show="scale" transition-hide="scale">
                        <q-date v-model="filter" mask="YYYY-MM-DD">
                          <div class="row items-center justify-end">
                            <q-btn v-close-popup label="Close" color="negative" flat />
                            <q-btn v-close-popup label="Last 24h" color="primary" flat @click="resetFilter"/>
                          </div>
                        </q-date>
                      </q-popup-proxy>
                    </q-icon>
                  </template>
                </q-input>
              </div>
            </q-card-section>
          </q-card>
        </div>
    </div>
</template>

<script>
  import { defineComponent } from "vue";
  import SettingsPopup from "components/SettingsPopup.vue";

  export default defineComponent({
    name: "TopBar",
    methods: {
      round (val) {
        return val ? val.toFixed(3) : ' - '
      },
      async toggleEnabled () {
        await this.$store.dispatch('master/toggleEnabled')
        await this.$store.dispatch('master/getSettings')
      },
      async resetFilter() {
        this.$store.commit('master/setFilter', "")
        this.filterEv()
      },
      async filterEv() {
        this.$store.dispatch('master/getStats')
      }
    },
    computed: {
      stats () {
        return this.$store.state['master'].stats
      },
      filter: {
        get() {
          return this.$store.state['master'].dateFilter
        },
        set(data) {
          this.$store.commit('master/setFilter', data)
        }
      },
      enabledStatus() {
        const settings = this.$store.state['master'].settings
        return {
          icon: settings.enabled ? 'stop' : 'play_arrow',
          color: settings.enabled ? 'red': 'green'
        }
      }
    },
    components: {
      SettingsPopup
    }
  });
</script>