import { createRouter, createWebHistory } from 'vue-router'
import RegisterPage from '../views/RegisterPage.vue'

const routes = [
  {
    path: '/',
    name: 'Register',
    component: RegisterPage
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
