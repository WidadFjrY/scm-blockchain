<script setup>
import { web3, SupplyChainContract } from '@/assets/script/eth-transaction.js'
import { ref } from 'vue';

import axios from 'axios';

import CardOverview from '@/components/CardOverview.vue';
import SideBar from '@/components/SideBar.vue';
import NavBarDash from '@/components/NavBarDash.vue';

import { useRouter, useRoute } from 'vue-router';

const route = useRoute()
const products = ref([])
const product = ref([])

function getStatusColor(status) {
    if (status === "Pengiriman") {
        return "yellow";
    } else if (status === "Diterima") {
        return "green";
    } else if (status === "Pending") {
        return "red";
    }
}

const totalUser = ref(0)
const totalProduct = ref(0)
const transactionsArray = ref([]);
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

async function getTotal() {
    try {
        const responseTotalUser = await axios.get(`${BACKEND_BASE_URL}/user/count`)
        totalUser.value = responseTotalUser.data.data.total_user

        const responseTotalProduct = await axios.get(`${BACKEND_BASE_URL}/product/count`)
        totalProduct.value = responseTotalProduct.data.data.total_product

    } catch (error) {
        console.log(error)
    }
}

async function getAllPrduct() {
    try {
        const response = await axios.get(`${BACKEND_BASE_URL}/products`)
        products.value = response.data.data.sort((a, b) => a.stock - b.stock);
    } catch (error) {
        console.log(error)
    }
}

async function getAllPendingTransactions() {
    try {
        const transactions = await SupplyChainContract.methods.getAllPendingTransactions().call();
        const productPromises = [];

        for (let i = 0; i < transactions.length; i++) {
            transactionsArray.value.push({
                id: transactions[i][0],
                buyer: transactions[i][1],
                productIds: transactions[i][2],
                quantities: transactions[i][3].map(qty => Number(qty).toLocaleString()),
                totalPrice: web3.utils.fromWei(transactions[i][4], 'ether'),
                shippingAddress: transactions[i][5],
                timestamp: new Date(Number(transactions[i][6]) * 1000).toLocaleString(),
                status: transactions[i][7]
            });

            transactionsArray.value[i].productIds.forEach(productId => {
                productPromises.push(getProduct(transactions[i][0], productId));
            });
        }

        await Promise.all(productPromises);

        console.log(transactionsArray.value[0])
    } catch (error) {
        console.error("Gagal mengambil transaksi yang belum selesai:", error);
    }
}

async function getProduct(trxId, productId) {
    try {
        const response = await axios.get(`${BACKEND_BASE_URL}/product/${productId}`);
        product.value.push({ trxId: trxId, data: response.data.data });

    } catch (error) {
        console.log(error);
    }
}

function getStatusClass(status) {
    return status.toLowerCase()
}

getAllPendingTransactions()
getAllPrduct()
getDataUser()
getTotal()

</script>

<template>
    <SideBar v-if="user.role" :role="user.role"></SideBar>

    <div class="container">
        <NavBarDash :user="user.name" :role="user.role" :title="route.name"></NavBarDash>
        <div class="overview">
            <h2>Ringkasan</h2>
            <div class="card">
                <CardOverview :title="'Total Produk'" :value="totalProduct" :icon="'Produk.png'"></CardOverview>
                <CardOverview :title="'Total Pengguna'" :value="totalUser" :icon="'Manajemen Distributor.png'">
                </CardOverview>
                <CardOverview :title="'Pesanan Aktif'" :value="transactionsArray.length"
                    :icon="'Manajemen Pesanan.png'"></CardOverview>
            </div>
        </div>
        <div class="list-table">
            <div class="list-product">
                <h2>Daftar Stok Barang</h2>
                <div class="product-content">
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
                            <tr v-if="!products">
                                <td colspan="4" style="text-align: center; color: gray;">Belum ada produk</td>
                            </tr>
                            <tr v-else v-for="(product, index) in products.slice(0, 5)" :key="product.id">
                                <td>{{ index + 1 }}</td>
                                <td>{{ product.id }}</td>
                                <td>{{ product.product_name }}</td>
                                <td :style="product.stock < 10 ? 'color: red;' : 'color: var(--dark-color)'">{{
                                    product.stock }}</td>
                            </tr>
                        </tbody>
                    </table>
                    <div style="margin-top: 1rem;">
                        <router-link to="/products" style="color: var(--background); font-size: 1.3rem; ">Lihat
                            semua></router-link>
                    </div>
                </div>
            </div>
            <div class="list-order">
                <h2>Daftar Pesanan Aktif</h2>
                <div class="order-content">
                    <table>
                        <thead>
                            <tr>
                                <th>ID</th>
                                <th>Barang</th>
                                <th>Status</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr v-if="transactionsArray.length === 0">
                                <td colspan="3" style="text-align: center; color: gray;">Belum ada pesanan</td>
                            </tr>
                            <tr v-else v-for="(transaction, index) in transactionsArray.slice(0, 5)"
                                :key="transaction.id">
                                <td>{{ transaction.id }}</td>
                                <td>
                                    {{ product[0].data.product_name }}
                                    {{ transaction.productIds.length - 1 === 0 ? "" : `+${transaction.productIds.length
                                        - 1}
                                    lainnya` }}
                                </td>
                                <td>
                                    <div :class="getStatusClass(transaction.status)">
                                        {{ transaction.status }}
                                    </div>
                                </td>
                            </tr>
                        </tbody>
                    </table>
                    <div style="margin-top: 1rem;">
                        <router-link to="/orders" style="color: var(--background); font-size: 1.3rem; ">Lihat
                            semua></router-link>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<style scoped>
.pending {
    color: #F59E0B;
}

.proses {
    color: #3B82F6;
}

.pengiriman {
    color: #8B5CF6;
}

.selesai {
    color: #10B981;
}

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
    margin-top: 2rem;
}

.list-product,
.list-order {
    width: 49%;
}

.order-content,
.product-content {
    margin-top: 0.5rem;
    width: 100%;
    background-color: white;
    padding: 1rem;
    border-radius: 1rem;
    box-shadow: rgba(0, 0, 0, 0.16) 0px 1px 4px;

}

table {
    width: 100%;
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
</style>