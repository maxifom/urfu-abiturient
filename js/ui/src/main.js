import Vue from 'vue'
import App from './App.vue'
import vuetify from './plugins/vuetify'
import VueTimeago from 'vue-timeago'

Vue.use(VueTimeago, {
    name: 'Timeago',
    locale: 'en',
})

Vue.config.productionTip = false

new Vue({
    vuetify,
    render: h => h(App)
}).$mount('#app')
