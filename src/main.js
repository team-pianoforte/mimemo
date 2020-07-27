import Vue from 'vue'
import VueTouch from 'vue-touch'
import { firestorePlugin } from 'vuefire'
import App from './App.vue'
import './registerServiceWorker'
import router from './router'
import store from './store'
import vuetify from './plugins/vuetify'

Vue.use(VueTouch, { name: 'v-touch' })
Vue.use(firestorePlugin)

Vue.config.productionTip = false

new Vue({
  router,
  store,
  vuetify,
  render: (h) => h(App),
}).$mount('#app')
