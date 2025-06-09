import { Dashboard } from '@/pages/dashboard';
import { Home } from '@/pages/home';
import { Login } from '@/pages/login';
import { Register } from '@/pages/register';
import AppLayout from '@/layouts/AppLayout.vue';
import { createRouter, createWebHistory } from 'vue-router';
import { Charges, ChargesAdd } from '@/pages/charges';

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      component: AppLayout,
      children: [
        {
          path: '/',
          name: 'home',
          component: Home,
        },
        {
          path: '/register',
          name: 'register',
          component: Register,
        },
        {
          path: '/login',
          name: 'login',
          component: Login,
        },
        {
          path: '/dashboard',
          name: 'dashboard',
          component: Dashboard,
        },
        {
          path: '/charges',
          name: 'charges',
          component: Charges,
        },
        {
          path: '/charges/add',
          name: 'charges add',
          component: ChargesAdd,
        },
      ],
    },
  ],
});

export default router;
