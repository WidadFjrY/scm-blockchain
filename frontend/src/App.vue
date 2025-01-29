<script setup>
import { useRoute } from 'vue-router';
import { CheckAuth } from '@/assets/script/check-auth';
import { ref, computed } from 'vue';

import axios from 'axios';
import Cookies from 'js-cookie'

import SideBar from './components/SideBar.vue';
import NavBar from './components/NavBar.vue';

CheckAuth()
const route = useRoute()

const user = ref({
  email: "",
  name: "",
  role: "",
})

const token = Cookies.get("AUTH_TOKEN")
const BASE_URL_BACKEND = import.meta.env.VITE_BACKEND_BASE_URL

async function getDataUser() {
  try {
    axios.defaults.headers.common["Authorization"] = `Bearer ${token}`
    const response = await axios.get(`${BASE_URL_BACKEND}/user`)

    user.value.email = response.data.data.email
    user.value.name = response.data.data.name
    user.value.role = response.data.data.role
  } catch (error) {
    console.log(error)
  }
}

getDataUser()

</script>

<template>
  <div v-if="route.path !== '/login'">
    <SiddeBar :role="user.role"></SiddeBar>
    <div class="view">
      <NavBar :title="route.name" :user="user.name" :role="user.role"></NavBar>
      <router-view></router-view>
    </div>
  </div>
  <router-view v-else></router-view>
</template>

<style scoped>
.view {
  margin-left: 25%;
  padding: 3rem;
}
</style>
