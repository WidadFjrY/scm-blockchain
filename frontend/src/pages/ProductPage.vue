<script setup>
import { reactive, ref, watch } from 'vue';

import axios from 'axios'
import { formatedDate } from '../assets/script/formated-date';

const state = reactive({
    search: null,
    products: null,
    product: null,
    filteredProducts: null,
})

const isModalOpen = ref(false)

const BACKEND_BASE_URL = import.meta.env.VITE_BACKEND_BASE_URL

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
        const response = await axios.get(`${BACKEND_BASE_URL}/product`)
        state.products = response.data.data;
        state.filteredProducts = response.data.data;
    } catch (error) {
        console.log(error)
    }
}

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
                <h2>Detail Product</h2>
                <button @click="isModalOpen = !isModalOpen" class="btn-close">X</button>
            </div>
            <div class="content">
                <div>
                    <h3>ID Produk</h3>
                    <p>{{ state.product.product_id }}</p>
                    <h3>Nama Produk</h3>
                    <p>{{ state.product.product_name }}</p>
                    <h3>Brand/Merek</h3>
                    <p> {{ state.product.brand }} </p>
                    <h3>Distributor</h3>
                    <p>{{ state.product.distributor_name }}</p>
                    <h3>Harga</h3>
                    <p>{{ state.product.price }}</p>
                    <h3>Stok</h3>
                    <p>{{ state.product.stock }}</p>
                    <h3>Satuan</h3>
                    <p>{{ state.product.unit }}</p>
                    <h3>Status</h3>
                    <p>{{ state.product.status }}</p>
                    <h3>Ditambahkan Pada</h3>
                    <p> {{ formatedDate(state.product.created_at) }} </p>
                    <h3>Diperbaharui Pada</h3>
                    <p> {{ formatedDate(state.product.updated_at) }} </p>
                </div>
                <div>
                    <h3>Gambar Produk</h3>
                    <div class="card">
                        <img :src="`${BACKEND_BASE_URL}/${state.product.url_picture}`" width="400" alt="">
                    </div>
                </div>
            </div>
        </div>
    </div>

    <div class="container">
        <h2>Daftar Produk</h2>
        <input type="text" placeholder="Cari Product" v-model="state.search">
        <table v-if="state.products != null">
            <thead>
                <tr>
                    <th style="width: 7rem;">No</th>
                    <th style="width: 20rem;">Nama Produk</th>
                    <th style="width: 10rem;">Stok</th>
                    <th style="width: 20rem;">Harga</th>
                    <th style="text-align: center;">Aksi</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="(product, index) in state.filteredProducts" :key="product.product_id">
                    <td>{{ index + 1 }}</td>
                    <td>{{ product.product_name }}</td>
                    <td>{{ product.stock }}</td>
                    <td>{{ product.price }}</td>
                    <td style="display: flex; justify-content: center;"><button
                            @click.prevent="getProductById(product.product_id)">Lihat
                            Detail</button></td>
                </tr>
            </tbody>
        </table>

        <router-link to="/add/product">
            Tambah Product
        </router-link>
    </div>
</template>

<style scoped>
.container {
    padding: 2rem;
    margin-top: 1rem;
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

table button {
    padding: 1rem;
    border-radius: 1rem;
    font-size: 1.3rem;
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
}

.modal .content .card {
    box-shadow: rgba(0, 0, 0, 0.16) 0px 1px 4px;
    border-radius: 1rem;
}
</style>