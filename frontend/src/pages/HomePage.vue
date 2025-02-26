<script setup>
import NavBar from '@/components/NavBar.vue';
import ProductCard from '@/components/ProductCard.vue';
import semenImg from '@/assets/img/semen.png';

import { CheckAuth } from '@/assets/script/check-auth';
import { ref } from 'vue';

import axios from 'axios';
import Cookies from 'js-cookie'

CheckAuth()
const user = ref({
    name: "",
    ethAddr: "",
})

const ethPrice = ref();
const count = ref(0)
const token = Cookies.get("AUTH_TOKEN")
const BASE_URL_BACKEND = import.meta.env.VITE_BACKEND_BASE_URL

const convertToETH = (idrPrice) => {
    return ethPrice.value ? (idrPrice / ethPrice.value).toFixed(6) : "Loading...";
};

async function getDataUser() {
    if (token) {
        try {
            axios.defaults.headers.common["Authorization"] = `Bearer ${token}`
            const response = await axios.get(`${BASE_URL_BACKEND}/user`)

            user.value.name = response.data.data.name
            user.value.ethAddr = response.data.data.eth_addr

        } catch (error) {
            console.log(error)
        }
    }
}

async function ETHPrice() {
    const response = await axios.get("https://api.coingecko.com/api/v3/simple/price?ids=ethereum&vs_currencies=idr");
    ethPrice.value = response.data.ethereum.idr;
}

ETHPrice()
getDataUser()

</script>
<template>
    <NavBar :name="user.name" :ethAddr="user.ethAddr"></NavBar>
    <div class="container">
        <div class="head">
            <h1>Daftar Barang</h1>
            <div style="display: flex; align-items: center;">
                <img src="@/assets/img/eth-logo.png" width="50">
                <h2>1 ETH = Rp. {{ ethPrice.toLocaleString("id-ID") }}</h2>
            </div>
        </div>
        <ProductCard :img="semenImg" :productName="'Semen'" :stock="420" :unit="'Sak'" :brand="'Tiga Roda'"
            :price="convertToETH(67000)" />
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

.head {
    display: flex;
    align-items: center;
    justify-content: space-between;
}
</style>