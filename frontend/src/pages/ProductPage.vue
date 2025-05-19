<script setup>
import { reactive, ref, watch } from 'vue';

import { formatedDate } from '../assets/script/formated-date';
import axios from 'axios'
import { useRoute } from 'vue-router';

import SideBar from '@/components/SideBar.vue';
import NavBarDash from '@/components/NavBarDash.vue';

const route = useRoute()
const ethPrice = ref(null);
const newStock = ref(0)
const state = reactive({
    search: null,
    products: null,
    product: null,
    filteredProducts: null,
})

const isModalOpen = ref(false)

const BACKEND_BASE_URL = import.meta.env.VITE_BACKEND_BASE_URL

const convertToETH = (idrPrice) => {
    return ethPrice.value ? (idrPrice / ethPrice.value).toFixed(6) : "Loading...";
};

async function getProductById(id) {
    try {
        const response = await axios.get(`${BACKEND_BASE_URL}/product/${id}`)
        state.product = response.data.data
    } catch (error) {
        console.log(error)
    }

    isModalOpen.value = !isModalOpen.value
}

async function getAllPrduct() {
    try {
        const response = await axios.get(`${BACKEND_BASE_URL}/products`)
        state.products = response.data.data;
        state.filteredProducts = response.data.data;
    } catch (error) {
        console.log(error)
    }
}

const user = ref({
    name: "",
    ethAddr: "",
    role: "",
})

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

async function ETHPrice() {
    try {
        const response = await axios.get("https://api.coingecko.com/api/v3/simple/price?ids=ethereum&vs_currencies=idr");
        ethPrice.value = response.data.ethereum.idr;
    } catch (error) {
        ethPrice.value = "loading..."
        console.log(error)
    }
}

async function stockUpdateHandle(productId) {
    try {
        await axios.put(`${BACKEND_BASE_URL}/product/stock`,
            {
                product_id: productId,
                stock_in: newStock.value,
                stock_out: 0
            }
        )
        state.product.stock = state.product.stock + newStock.value
        newStock.value = 0
        await getAllPrduct()
    } catch (error) {
        console.log(error)
    }
}

ETHPrice()
getDataUser()
getAllPrduct()

watch(
    () => state.search,
    (newValue) => {
        if (newValue) {
            state.filteredProducts = state.products.filter((item) =>
                newValue
                    .toLowerCase()
                    .split(" ")
                    .every(
                        (term) =>
                            item.product_name.toLowerCase().includes(term))
            );
        } else {
            state.filteredProducts = state.products;
        }
    },

    { immediate: true }
);

</script>

<template>
    <div v-if="isModalOpen" class="modal-overlay" @click.self="isModalOpen = !isModalOpen">
        <div class="modal" id="modal">
            <div style="display: flex; justify-content: space-between; align-items: center;">
                <h2>Detail Barang</h2>
                <button @click="isModalOpen = !isModalOpen" class="btn-close">X</button>
            </div>
            <div class="content">
                <div>
                    <h3>ID Produk</h3>
                    <p>{{ state.product.id }}</p>
                    <h3>Nama Produk</h3>
                    <p>{{ state.product.product_name }}</p>
                    <h3>Brand/Merek</h3>
                    <p> {{ state.product.brand }} </p>
                    <h3>Harga</h3>
                    <p>Rp. {{ state.product.price.toLocaleString("id-ID") }}</p>
                    <p>{{ convertToETH(state.product.price) }} ETH</p>
                    <h3>Stok</h3>
                    <p>{{ state.product.stock.toLocaleString("id-ID") }}</p>
                    <h3>Satuan</h3>
                    <p>{{ state.product.unit }}</p>
                    <h3>Ditambahkan Pada</h3>
                    <p>{{ formatedDate(state.product.created_at) }}</p>
                </div>
                <div>
                    <h3>Gambar Produk</h3>
                    <div class="card">
                        <img :src="`${BACKEND_BASE_URL}/${state.product.filepath}`" alt="">
                    </div>
                    <template v-if="user.role === 'Warehouse_Staff'">
                        <h3>Tambah Stok</h3>
                        <div style="display: flex; align-items: center; gap: 1rem; margin-top: 0.5rem;">
                            <input type="number" v-model="newStock">
                            <button class="btn" @click.prevent="stockUpdateHandle(state.product.id)"
                                :disabled="newStock < 1">
                                <h1 class="add-stock">+</h1>
                            </button>
                        </div>
                    </template>
                </div>
            </div>
        </div>
    </div>

    <SideBar v-if="user.role" :role="user.role"></SideBar>

    <div class="container">
        <NavBarDash :user="user.name" :role="user.role" :title="route.name"></NavBarDash>
        <div v-if="state.products !== null">
            <h2>Daftar Produk</h2>
            <input type="text" placeholder="Cari Product" v-model="state.search">
        </div>
        <div v-if="state.products !== null" class="card-container">
            <table>
                <thead>
                    <tr>
                        <th>No</th>
                        <th>Nama Produk</th>
                        <th>Merek</th>
                        <th>Stok</th>
                        <th>Harga</th>
                        <th>Aksi</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="(product, index) in state.filteredProducts" :key="product.id">
                        <td>{{ index + 1 }}</td>
                        <td>{{ product.product_name }}</td>
                        <td>{{ product.brand }}</td>
                        <td>{{ product.stock.toLocaleString("id-ID") }}</td>
                        <td>Rp. {{ product.price.toLocaleString("id-ID") }}</td>
                        <td style="display: flex; justify-content: center;"><button
                                @click.prevent="getProductById(product.id)">Lihat
                                Detail</button></td>
                    </tr>
                </tbody>
            </table>
        </div>
        <div class="card-container" v-else>
            <h2 style="text-align: center; padding: 3rem;">Belum ada produk</h2>
        </div>
        <router-link v-if="user.role === 'Warehouse_Staff'" to="/add/product">
            Tambah Produk
        </router-link>
    </div>
</template>

<style scoped>
.container {
    padding: 2rem;
    margin-left: 25%;
}

.container input {
    padding: 1rem;
    margin: 1rem 0;
    color: var(--dark-color);
    border-radius: 1rem;
    background-color: white;
    border: none;
    outline: none;
    width: 100%;
    font-size: 1.5rem;
    box-shadow: rgba(0, 0, 0, 0.16) 0px 1px 4px;

}


.card-container {
    background-color: white;
    border-radius: 1rem;
    box-shadow: rgba(0, 0, 0, 0.16) 0px 1px 4px;
    overflow: hidden;
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
    text-align: center;
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

table button {
    padding: 1rem;
    border-radius: 0.5rem;
    font-size: 1rem;
    background-color: rgb(32, 193, 243);
    color: white;
    border: none;
    cursor: pointer;
    transition: 0.5s ease;
}


.container>a {
    display: inline-block;
    margin-top: 2rem;
    padding: 1rem;
    border-radius: 1rem;
    font-size: 1.3rem;
    background-color: var(--background);
    color: white;
    border: none;
    cursor: pointer;
    transition: 0.5s ease;
}

button:hover,
.container>a:hover {
    background-color: rgb(55, 208, 255);
}

.modal-overlay {
    position: fixed;
    background-color: rgba(63, 63, 63, 0.267);
    width: 100vw;
    height: 100vh;
    top: 0;
    left: 0;
    z-index: 999;
    display: flex;
    justify-content: center;
    align-items: center;
}

.modal {
    background-color: white;
    width: 50%;
    padding: 3rem;
    border-radius: 1rem;
    transition: 0.5;
}

.modal .btn-close {
    cursor: pointer;
    background-color: white;
    border: none;
    font-size: 2rem;
    font-weight: 900;
    color: #686868;
}

.modal h2 {
    font-size: 2rem;
}

.modal h3 {
    margin-top: 1rem;
    font-size: 1.3rem;
}

.modal p {
    font-size: 1.3rem;
}

.modal .content {
    margin-top: 1.5rem;
    display: flex;
    justify-content: center;
    gap: 4rem;
}

.modal .content .card {
    box-shadow: rgba(0, 0, 0, 0.16) 0px 1px 4px;
    border-radius: 1rem;
    width: 350px;
    height: 350px;
    margin-top: 1rem;
    overflow: hidden;
}

.modal .content .card img {
    width: 100%;
    height: 100%;
    object-fit: cover;
}

.modal input {
    display: flex;
    align-items: center;
    padding-left: 1rem;
    width: 100% !important;
    height: 3rem;
    border-radius: 0.5rem;
    border: none;
    outline: none;
    color: var(--dark-color);
    background-color: #f2f2f2;
    font-size: 1.3rem;
}

.modal .btn {
    background-color: var(--background);
    display: flex;
    justify-content: center;
    align-items: center;
    height: 40px;
    width: 50px;
    border-radius: 0.5rem;
    border: none;
    cursor: pointer;
}

.modal button:disabled {
    cursor: not-allowed;
    background-color: rgb(154, 235, 255);
}

.modal .add-stock {
    font-size: 2.5rem;
    color: white;
    font-weight: 400;
}
</style>