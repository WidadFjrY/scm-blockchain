<script setup>
import NavBar from '@/components/NavBar.vue';
import ProductCard from '@/components/ProductCard.vue';

import { ref, reactive } from 'vue';

import axios from 'axios';

const user = ref({
    name: "",
    ethAddr: "",
    role: "",
})

const state = reactive({
    products: null,
})

const ethPrice = ref();
const count = ref(0)
const BACKEND_BASE_URL = import.meta.env.VITE_BACKEND_BASE_URL

const convertToETH = (idrPrice) => {
    return ethPrice.value ? (idrPrice / ethPrice.value).toFixed(6) : "Loading...";
};

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

async function getAllPrduct() {
    try {
        const response = await axios.get(`${BACKEND_BASE_URL}/products`)
        state.products = response.data.data;
    } catch (error) {
        console.log(error)
    }
}

getAllPrduct()
ETHPrice()
getDataUser()

</script>
<template>
    <NavBar :name="user.name" :ethAddr="user.ethAddr" :role="user.role"></NavBar>
    <div class="container">
        <div class="head">
            <h1>Daftar Barang</h1>
            <div style="display: flex; align-items: center;">
                <img src="@/assets/img/eth-logo.png" width="50">
                <h2>1 ETH = Rp. {{ ethPrice.toLocaleString("id-ID") }}</h2>
            </div>
        </div>
        <div class="card-container">
            <div v-for="(product, index) in state.products" :key="index">
                <ProductCard :img="`${BACKEND_BASE_URL}/${product.filepath}`" :productId="product.id"
                    :productName="product.product_name" :stock=product.stock :unit=product.unit :brand=product.brand
                    :price="convertToETH(product.price)" />
            </div>
        </div>
    </div>
</template>
<style scoped>
.container {
    background-color: white;
    margin: 2rem;
    padding: 2rem;
    box-shadow: rgba(60, 64, 67, 0.3) 0px 1px 2px 0px, rgba(60, 64, 67, 0.15) 0px 1px 3px 1px;
    border-radius: 1rem;
}


.container .card-container {
    display: flex;
    flex-wrap: wrap;
    gap: 1rem;
}

.head {
    display: flex;
    align-items: center;
    justify-content: space-between;
}
</style>