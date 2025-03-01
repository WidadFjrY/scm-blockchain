<script setup>

import axios from 'axios';
import Cookies from 'js-cookie';

const props = defineProps({
    name: "",
    ethAddr: "",
    role: ""
});

const BACKEND_BASE_URL = import.meta.env.VITE_BACKEND_BASE_URL

async function logoutHandle() {
    try {
        const response = await axios.post(`${BACKEND_BASE_URL}/user/logout`)
        if (response.status === 200) {
            Cookies.remove("AUTH_TOKEN")
            window.location.href = "/login"
        }
    } catch (error) {
        console.log(error)
    }
}

</script>
<template>
    <nav>
        <div class="logo">
            <img src="@/assets/img/saj.jpg" width="60" alt="">
            <h1>CV. Sindang Agung Jaya</h1>
        </div>
        <ul>
            <li><router-link to="/transaction-hisoty">Riwayat Transasksi</router-link></li>
            <li><router-link to="/tracking">Lacak Pesanan</router-link></li>
            <li><router-link to="/cart">Keranjang</router-link></li>
            <li v-if="props.role === 'Admin' || props.role === 'Warehouse_Staff'"><router-link
                    to="/dashboard">Dashboard</router-link></li>
        </ul>
        <ul>
            <li>{{ props.name }}</li>
            <li>{{ props.ethAddr }}</li>
            <li @click.prevent="logoutHandle" style="cursor: pointer !important;"><img src="/assets/icon/logout.png"
                    width="40" alt="">
            </li>
        </ul>
    </nav>
</template>

<style scoped>
nav {
    display: flex;
    padding: 1rem 0;
    align-items: center;
    justify-content: space-around;
    background-color: white;
    box-shadow: rgba(60, 64, 67, 0.3) 0px 1px 2px 0px, rgba(60, 64, 67, 0.15) 0px 1px 3px 1px;
}

.logo {
    display: flex;
    align-items: center;
    gap: 1rem;
}

ul {
    display: flex;
    align-items: center;
    gap: 2rem;
}

ul li {
    font-size: 1.3rem;
    font-weight: 500;
    color: var(--background);
}

a:hover {
    opacity: 0.5;
}
</style>