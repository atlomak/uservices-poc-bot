import { createRouter, createWebHistory } from 'vue-router';
import Home from '../views/HomeView.vue';
import Login from '../views/LoginView.vue';

const routes = [
  {
    path: '/home',
    name: 'Home',
    component: Home
  },
  {
    path: '/login',
    name: 'Login',
    component: Login
  },
  {
    path: "/",
    redirect: "/login"
  }
];

const router = createRouter({
  history: createWebHistory(),
  routes
});

export default router;
