import Vue from 'vue'
import App from './App.vue'
import VueRouter from 'vue-router'

Vue.config.productionTip = false
Vue.use(VueRouter)

const Bar = {template: `<div class="bar"><h1>bar</h1></div>`}
import HelloWorld from './components/HelloWorld'
import Foo from './components/Foo'

const routes = [
    {path: '/', component: HelloWorld},
    {path: '/foo', component: Foo},
    {path: '/bar', component: Bar}
]

const router = new VueRouter({
    routes // short for `routes: routes`
})

new Vue({
    router,
    render: h => h(App)
}).$mount('#app')
