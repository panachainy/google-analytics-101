import Vue from 'vue'
import App from './App.vue'
import VueGtag from "vue-gtag";
import VueGtm from "vue-gtm";

Vue.config.productionTip = false

Vue.use(VueGtag, {
  config: { id: process.env.GOOGLE_ANALYTICS_ID }
});

Vue.use(VueGtm, {
  id: process.env.GOOGLE_TAG_MANAGER,
  defer: false,
  compatibility: false,
  enabled: true,
  debug: true,
  loadScript: true,
  trackOnNextTick: false,
});

new Vue({
  render: h => h(App),
}).$mount('#app')
