<script setup>
import { ref } from 'vue';

import axios from 'axios';

import CardOverview from '@/components/CardOverview.vue';
import SideBar from '@/components/SideBar.vue';
import NavBarDash from '@/components/NavBarDash.vue';

import { useRouter, useRoute } from 'vue-router';

const route = useRoute()

const products = ref([
    { name: "Nippon Paint 5 kg", id: 'ID381729', stock: 23 },
    { name: "Pipa 5M", id: 'ID381729', stock: 10 },
    { name: "Keran Onda", id: 'ID381729', stock: 3 },
    { name: "Paku Beton", id: 'ID381729', stock: '10 Kg' },
    { name: "Semen 3 Roda 50 Kg", id: 'ID381729', stock: 32 },
])

const orders = ref([
    { name: "PT. Besi", status: 'Dikirim' },
    { name: "Batu", status: 'Pending' },
    { name: "PVC", status: 'Dikirim' },
    { name: "Nippon", status: 'Diterima' },
    { name: "Tiga Roda", status: 'Diterima' },
])

function getStatusColor(status) {
    if (status === "Dikirim") {
        return "yellow";
    } else if (status === "Diterima") {
        return "green";
    } else if (status === "Pending") {
        return "red";
    }
}

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
        <div class="overview">
            <h2>Ringkasan</h2>
            <div class="card">
                <CardOverview :title="'Total Produk'" :value="10" :icon="'Produk.png'"></CardOverview>
                <CardOverview :title="'Total Distributor'" :value="10" :icon="'Manajemen Distributor.png'">
                </CardOverview>
                <CardOverview :title="'Pesanan Aktif'" :value="10" :icon="'Manajemen Pesanan.png'"></CardOverview>
            </div>
        </div>
        <div class="list-table">
            <div class="list-product">
                <h2>Daftar Stok Barang</h2>
                <table>
                    <thead>
                        <tr>
                            <th>No</th>
                            <th>ID Barang</th>
                            <th>Nama</th>
                            <th>Stok</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr v-for="(product, index) in products" :key="product.id">
                            <td>{{ index + 1 }}</td>
                            <td>{{ product.id }}</td>
                            <td>{{ product.name }}</td>
                            <td style="color: red;">{{ product.stock }}</td>
                        </tr>
                    </tbody>
                </table>
            </div>
            <div class="list-order">
                <h2>Daftar Pesanan Aktif</h2>
                <table>
                    <thead>
                        <tr>
                            <th>No</th>
                            <th>Distributor</th>
                            <th>Status</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr v-for="(order, index) in orders" :key="order.id">
                            <td>{{ index + 1 }}</td>
                            <td>{{ order.name }}</td>
                            <td>
                                <div class="status"
                                    :style="{ backgroundColor: getStatusColor(order.status), color: order.status === 'Dikirim' ? 'var(--dark-color)' : 'white' }">
                                    {{
                                        order.status }}</div>
                            </td>
                        </tr>
                    </tbody>
                </table>
            </div>
        </div>
    </div>
</template>

<style scoped>
.container {
    padding: 2rem;
    margin-left: 25%;
}

.overview {
    margin-top: 1rem;
}

.overview .card {
    margin-top: 0.5rem;
    display: flex;
    justify-content: center;
    gap: 2rem
}

.list-table {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-top: 2rem;
}

.list-product,
.list-order {
    width: 49%;
}


table {
    margin-top: 0.5rem;
    width: 100%;
    background-color: white;
    padding: 1rem;
    border-radius: 1rem;
    box-shadow: rgba(0, 0, 0, 0.16) 0px 1px 4px;

}

table th,
table td {
    padding: 1rem;
    text-align: start;
    padding-top: 1rem;
    border-bottom: 1px solid rgb(207, 207, 207);

}

table th {
    font-size: 1.3rem;
    font-weight: bold;
}

table td {
    font-size: 1.3rem;
}

.status {
    width: 90%;
    border-radius: 1rem;
    color: white;
    padding: 0 1rem;
    text-align: center;
}
</style>