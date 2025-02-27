<script setup>

import Cookies from 'js-cookie';
import axios from 'axios';

const token = Cookies.get("AUTH_TOKEN")
axios.defaults.headers.common["Authorization"] = `Bearer ${token}`

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
