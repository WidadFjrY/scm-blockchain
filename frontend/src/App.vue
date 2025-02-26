<script setup>
import { useRoute, useRouter } from 'vue-router';
import { CheckAuth } from '@/assets/script/check-auth';
import { ref } from 'vue';

import axios from 'axios';
import Cookies from 'js-cookie'

CheckAuth()
const route = useRoute()
const router = useRouter()
const user = ref({
  id: "",
  email: "",
  name: "",
  role: "",
})

const token = Cookies.get("AUTH_TOKEN")
const BASE_URL_BACKEND = import.meta.env.VITE_BACKEND_BASE_URL

async function getDataUser() {
  if (token) {
    try {
      axios.defaults.headers.common["Authorization"] = `Bearer ${token}`
      const response = await axios.get(`${BASE_URL_BACKEND}/user`)

      user.value.email = response.data.data.email
      user.value.name = response.data.data.name
      user.value.role = response.data.data.role
      user.value.id = response.data.data.id

    } catch (error) {
      console.log(error)
    }
  }
}

getDataUser()

</script>

<template>
  <router-view></router-view>
</template>

<style scoped>
.view {
  margin-left: 25%;
  padding: 3rem;
}
</style>
