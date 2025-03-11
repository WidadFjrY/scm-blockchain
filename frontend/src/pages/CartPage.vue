<script setup>
import NavBar from '@/components/NavBar.vue';

import { web3, SupplyChainContract } from '@/assets/script/eth-transaction.js'
import { ref, reactive } from 'vue';

import axios from 'axios';

const user = ref({
    name: "",
    ethAddr: "",
    role: "",
})
const ethPrice = ref();
const carts = ref([]);

const BACKEND_BASE_URL = import.meta.env.VITE_BACKEND_BASE_URL
const convertToETH = (idrPrice) => {
    if (!ethPrice.value || ethPrice.value <= 0) {
        return "Loading...";
    }
    return parseFloat((idrPrice / ethPrice.value).toFixed(6));
};

const prepareTransactionData = (carts) => {
    const productIds = carts.map(cart => cart.product_id)
    const quantities = carts.map(cart => cart.quantity)

    return { productIds, quantities }
}

const totalPrice = () => {
    let price = 0;

    if (carts.value) {
        carts.value.forEach(cart => {
            price += cart.product.price * cart.quantity;
        });
    }

    return price;
};

async function getData() {
    try {
        const userResponse = await axios.get(`${BACKEND_BASE_URL}/user`)
        const ETHResponse = await axios.get("https://api.coingecko.com/api/v3/simple/price?ids=ethereum&vs_currencies=idr");
        const cartResponse = await axios.get(`${BACKEND_BASE_URL}/cart`)


        user.value.name = userResponse.data.data.name
        user.value.ethAddr = userResponse.data.data.eth_addr
        user.value.role = userResponse.data.data.role

        ethPrice.value = ETHResponse.data.ethereum.idr

        carts.value = cartResponse.data.data

    } catch (error) {
        console.log(error)
    }

}

async function updateQty(productId, quantity) {
    try {
        await axios.put(`${BACKEND_BASE_URL}/cart?product_id=${productId}&qty=${quantity}`)
    } catch (error) {
        console.log(error)
    }
}

async function checkOutHandle() {
    const { productIds, quantities } = prepareTransactionData(carts.value);
    const account = await web3.eth.getAccounts();
    const userAddress = account[0];

    try {
        const valueInWei = web3.utils.toWei(convertToETH(totalPrice()).toString(), 'ether');
        console.log(valueInWei)
        const tx = await SupplyChainContract.methods
            .createTransaction(productIds, quantities)
            .send({ from: userAddress, value: valueInWei });

        console.log("Transaction Hash:", tx.transactionHash);

        window.location.href = "/";

    } catch (error) {
        console.error("Transaksi gagal:", error);
    }
}


getData()
</script>
<template>
    <NavBar :name="user.name" :ethAddr="user.ethAddr" :role="user.role"></NavBar>
    <div class="container">
        <div class="head">
            <h1>Keranjang</h1>
            <div style="display: flex; align-items: center;">
                <img src="@/assets/img/eth-logo.png" width="50">
                <h2>1 ETH = Rp. {{ ethPrice.toLocaleString("id-ID") }}</h2>
            </div>
        </div>
        <p v-if="!carts" style="text-align: center;">Tidak ada produk</p>
        <div v-else class="card-container">
            <div v-for="(cart, index) in carts" :key="index">
                <div class="card">
                    <div style="display: flex;gap:1rem ;width: 20rem;">
                        <div class="card-img">
                            <img :src="`${BACKEND_BASE_URL}/${cart.product.filepath}`" alt="">
                        </div>
                        <div>
                            <h2>{{ cart.product.product_name }}</h2>
                            <p>{{ cart.product.brand }}</p>
                            <p>Satuan {{ cart.product.unit }}</p>
                        </div>
                    </div>
                    <div class="card-btn">
                        <button @click.prevent="cart.quantity--, updateQty(cart.product_id, cart.quantity)">-</button>
                        <p>{{ cart.quantity }}</p>
                        <button @click.prevent="cart.quantity++, updateQty(cart.product_id, cart.quantity)">+</button>
                    </div>
                    <div class="card-price">
                        <p>Rp. {{ (cart.product.price * cart.quantity).toLocaleString("id-ID") }}</p>
                        <h3>{{ convertToETH(cart.product.price * cart.quantity) }} ETH</h3>
                    </div>
                </div>
            </div>
            <h2 style="text-align: end;">Total: {{ convertToETH(totalPrice()) }} ETH</h2>
            <div class="card-btn">
                <button @click.prevent="checkOutHandle">Check Out</button>
            </div>
        </div>
    </div>
</template>
<style scoped>
p {
    font-size: 1.2rem;
}

.container {
    background-color: white;
    margin: 2rem;
    padding: 2rem;
    box-shadow: rgba(60, 64, 67, 0.3) 0px 1px 2px 0px, rgba(60, 64, 67, 0.15) 0px 1px 3px 1px;
    border-radius: 1rem;
}

.card-container {
    margin-top: 2rem;
    padding: 0 4rem;
}

.head {
    display: flex;
    align-items: center;
    justify-content: space-between;
}

.card {
    display: flex;
    align-items: center;
    justify-content: space-between;
}

.card-img {
    width: 100px;
    height: 100px;
    margin-bottom: 1rem;
}

.card-img img {
    width: 100%;
    height: 100%;
    object-fit: cover;
    border-radius: 0.5rem;
}

.card-btn {
    display: flex;
    align-items: center;
}

.card-btn p {
    width: 2rem;
    text-align: center;
}

.card-btn button {
    background-color: white;
    width: 40px;
    height: 40px;
    font-size: 2rem;
    border-radius: 100%;
    border: var(--dark-color) 1px solid;
    cursor: pointer;
    color: gray;
}

.card-price {
    width: 10rem;
    display: flex;
    flex-direction: column;
    align-items: end;
}

.card-container>.card-btn {
    display: flex;
    justify-content: end;
    margin-top: 1rem;
}

.card-container>.card-btn button {
    width: 10rem;
    margin-top: 1rem;
    padding: 0.5rem;
    border: none;
    border-radius: 0.5rem;
    background: var(--background);
    background: var(--gradient);
    color: white;
    font-size: 1.4rem;
    cursor: pointer;
}
</style>