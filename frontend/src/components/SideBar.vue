<script setup>

import { ref } from 'vue';
import { useRoute } from 'vue-router';

import axios from 'axios'
import Cookies from 'js-cookie'

const props = defineProps({
    role: String,
});
const route = useRoute()
const appName = import.meta.env.VITE_APP_NAME
const links = ref([
    { name: "Dashboard", link: "/dashboard", icon: '/assets/icon/Dashboard.png', isVisible: props.role === "Admin" || props.role === "Warehouse_Staff" },
    { name: "Manajemen Pesanan", link: "/orders", icon: '/assets/icon/Manajemen Pesanan.png', isVisible: props.role === "Admin" || props.role === "Warehouse_Staff" },
    { name: "Produk", link: "/products", icon: '/assets/icon/Produk.png', isVisible: props.role === "Admin" || props.role === "Warehouse_Staff" },
    { name: "Laporan", link: "/reports", icon: '/assets/icon/Laporan.png', isVisible: props.role === "Admin" },
    { name: "Pengguna", link: "/users", icon: '/assets/icon/Pengguna.png', isVisible: props.role === "Admin" || props.role === "Warehouse_Staff" },
])

console.log(props.role)

const BACKEND_BASE_URL = import.meta.env.VITE_BACKEND_BASE_URL
const token = Cookies.get("AUTH_TOKEN")


async function logoutHandle() {
    try {
        axios.defaults.headers.common["Authorization"] = `Bearer ${token}`;
        await axios.post(`${BACKEND_BASE_URL}/user/logout`)

        Cookies.remove("AUTH_TOKEN")
        window.location.href = '/'
    } catch (error) {
        console.log(error)
    }
}

</script>
<template>
    <div class="side-bar">
        <div class="head">
            <img src="@/assets/img/icon.png" width="100" alt="">
            <h1>{{ appName }}</h1>
            <p>CV. Sindang Agung Jaya</p>
        </div>
        <ul class="item">
            <template v-for="(item, index) in links" :key="index">
                <router-link :to="item.link" :class="{ active: item.link === route.path }" v-if="item.isVisible">
                    <img :src="item.icon" width="40">
                    {{ item.name }}
                </router-link>
            </template>
        </ul>
        <button @click.prevent="logoutHandle">
            <img src="/assets/icon/logout.png" alt="">
            Logout</button>
    </div>

</template>
<style scoped>
.side-bar {
    padding: 2rem;
    width: 25%;
    height: 100vh;
    position: fixed;
    top: 0;
    left: 0;
    background-color: white;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    box-shadow: rgba(0, 0, 0, 0.16) 0px 1px 4px;

}

.head {
    display: flex;
    flex-direction: column;
    align-items: center;
}

.head h1 {
    margin-top: 1rem;
    color: var(--dark-color);
    text-align: center;
    font-weight: 500;
    font-size: 1.5rem;
}

.head p {
    color: var(--dark-color);
}

.item a {
    display: block;
    margin-top: 1rem;
    padding: 0.5rem;
    border-radius: 1rem;
    color: rgb(172, 172, 172);
    font-weight: 600;
    font-size: 1.2rem;
    display: flex;
    align-items: center;
    gap: 1rem;
    transition: 0.5s ease;
}

.item a img {
    pointer-events: none;
    filter: grayscale(100);
}


.active {
    padding-left: 2rem !important;
    background-color: rgb(241, 241, 241);
    color: var(--dark-color) !important;
}

.active img {
    filter: grayscale(0) !important;

}

.side-bar button {
    background-color: white;
    border: none;
    display: flex;
    align-items: center;
    color: var(--dark-color);
    font-weight: 600;
    font-size: 1.2rem;
    margin-left: 2rem;
    cursor: pointer;
}
</style>