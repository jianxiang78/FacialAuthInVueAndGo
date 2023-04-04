import Vue from 'vue';
import Router from 'vue-router';
import dashboard from '../components/sysDashboard';
import login from '../components/userLogin';

Vue.use(Router);

let routerPush = Router.prototype.push;
Router.prototype.push = function push(location) {
    return routerPush.call(this, location).catch(err => err)
}

export default new Router({
    routes: [
    { path: '/', redirect: '/login' },
    {
      path: '/dashboard',
      name: 'dashboard',
      component: dashboard,
      meta: {
          keepAlive: true
      }
    },
    {
      path: '/login',
      name: 'login',
      component: login
    }
  ]
})