<script setup>
import { ref } from 'vue';

import axios from 'axios';

import SideBar from '@/components/SideBar.vue';
import NavBarDash from '@/components/NavBarDash.vue';

import { useRouter, useRoute } from 'vue-router';

const route = useRoute()

const user = ref({
    name: "",
    ethAddr: "",
    role: "",
})

const BACKEND_BASE_URL = import.meta.env.VITE_BACKEND_BASE_URL

async function getDataUser() {
    try {
        const response = await axios.get(`${BACKEND_BASE_URL}/user`)
        user.value.name = response.data.data.name
        user.value.ethAddr = response.data.data.eth_addr
        user.value.role = response.data.data.role
    } catch (error) {
        console.log(error)
    }

}

getDataUser()

</script>

<template>
    <SideBar></SideBar>

    <div class="container">
        <NavBarDash :user="user.name" :role="user.role" :title="route.name"></NavBarDash>
    </div>

</template>

<style scoped>
.container {
    padding: 2rem;
    margin-left: 25%;
}
</style>