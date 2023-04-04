import Vue from 'vue'
import App from './App.vue'
import router from './router/index.js'

import {AjaxPost} from './assets/build/js/scripts.js'

import './assets/vendors/bootstrap/dist/css/bootstrap.min.css'
import './assets/vendors/font-awesome/css/font-awesome.min.css'

import 'bootstrap/dist/js/bootstrap.min.js'

Vue.prototype.AjaxPost = AjaxPost

Vue.config.productionTip = false

new Vue({
    el: '#app',
    router,
    render: h => h(App)
})
