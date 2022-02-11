<template>
    <div class="row q-col-gutter-md q-mb-md">
      <div class="col-4">
        <q-card
            class="bg-positive"
        >
          <q-card-section>
            <div class="text-h6">Latence</div>
          </q-card-section>

          <q-card-section class="text-center text-h4">
            {{round(stats.avgPing)}}ms
          </q-card-section>
        </q-card>
      </div>
      <div class="col-4">
          <q-card
              class="bg-info"
          >
            <q-card-section>
              <div class="text-h6">Jitter</div>
            </q-card-section>

            <q-card-section class="text-center text-h4">
              {{round(stats.avgJitter)}}ms
            </q-card-section>
          </q-card>
        </div>
        <div class="col-1">
          <q-card
              class="bg-negative"
          >
            <q-card-section>
              <div class="text-h6">Errors</div>
            </q-card-section>

            <q-card-section class="text-center text-h4">
              {{stats.errors}}
            </q-card-section>
          </q-card>
        </div>
        <div class="col-3">
          <SettingsPopup ref="settings"/>
          <q-card
              class="bg-negative"
          >
            <q-card-section>
              <a href="#" @click="$refs.settings.show()">Settings</a>
              <a href="#" >Start</a>
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
      }
    },
    computed: {
      stats () {
        return this.$store.state['master'].stats
      }
    },
    components: {
      SettingsPopup
    }
  });
</script>