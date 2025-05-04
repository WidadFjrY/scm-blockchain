<script setup>
import { web3, SupplyChainContract } from '@/assets/script/eth-transaction.js'
import { ref } from 'vue';
import axios from 'axios';

import SideBar from '@/components/SideBar.vue';
import NavBarDash from '@/components/NavBarDash.vue';

import { useRouter, useRoute } from 'vue-router';

const route = useRoute()

const isModalOpen = ref(false)
const transactionsArray = ref([]);
const products = ref([]);
const groupedTransactions = ref([]);
const selectedTransaction = ref();
const selectedTransactionIndex = ref(0)
const status = ref("")
const userTrx = ref({})
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

        groupTransactions();

    } catch (error) {
        console.error("Gagal mengambil transaksi yang belum selesai:", error);
    }
}


async function getProduct(trxId, productId) {
    try {
        const response = await axios.get(`${BACKEND_BASE_URL}/product/${productId}`);
        products.value.push({ trxId: trxId, data: response.data.data });
    } catch (error) {
        console.log(error);
    }
}

const groupTransactions = () => {
    const result = Object.values(products.value.reduce((acc, curr) => {
        const trxId = curr.trxId.toString();

        if (!acc[trxId]) {
            acc[trxId] = { trxId: curr.trxId, products: [] };
        }

        const productExists = acc[trxId].products.some(
            (product) => product.data.id === curr.data.id
        );

        if (!productExists) {
            acc[trxId].products.push(curr);
        }

        return acc;
    }, {}));

    groupedTransactions.value = result;
    console.log(transactionsArray.value[0].quantities)
};

async function getDataUserByETHAddr(addr) {
    try {
        const response = await axios.get(`${BACKEND_BASE_URL}/user/eth_addr/${addr}`)
        userTrx.value = response.data.data
    } catch (error) {
        console.log(error)
    }
}

async function selectedTransactionHandle(transactions, index) {
    isModalOpen.value = true
    await getDataUserByETHAddr(transactionsArray.value[index].buyer)
    selectedTransaction.value = transactions
    selectedTransactionIndex.value = index
}

async function updateHandle(transactionId) {
    const accounts = await web3.eth.getAccounts();
    const userAddress = accounts[0];

    try {
        await SupplyChainContract.methods
            .updateTransactionStatus(transactionId, status.value)
            .send({ from: userAddress });
    } catch (error) {
        console.log(error)
    }
    isModalOpen.value = !isModalOpen.value
}

function getStatusClass(status) {
    return status.toLowerCase()
}

getDataUser()
getAllPendingTransactions();

</script>

<template>
    <div v-if="isModalOpen" class="modal-overlay" @click.self="isModalOpen = !isModalOpen">
        <div class="modal" id="modal">
            <div style="display: flex; justify-content: space-between; align-items: center;">
                <h2>Detail Product</h2>
                <button @click="isModalOpen = !isModalOpen" class="btn-close">X</button>
            </div>
            <div class="content">
                <div style="width: 100%;">
                    <h3>ID Transaksi</h3>
                    <p>{{ transactionsArray[selectedTransactionIndex].id }}</p>
                    <h3>Nama Pembeli</h3>
                    <p>{{ userTrx.name }}</p>
                    <h3>Alamat ETH</h3>
                    <p>{{ transactionsArray[selectedTransactionIndex].buyer }}</p>
                    <h3>Alamat Pengiriman</h3>
                    <p>{{ transactionsArray[selectedTransactionIndex].shippingAddress }}</p>
                    <h3>No. Telepon</h3>
                    <p>{{ userTrx.telp }}</p>
                    <h3>Produk:</h3>
                    <div style="max-height: 350px; overflow-y: auto;">
                        <div v-for="(product, index) in selectedTransaction.products" :key="index">
                            <div style="display: flex; justify-content: space-between; align-items: center;">
                                <div>
                                    <h3>Nama Produk</h3>
                                    <p>{{ product.data.product_name }}</p>
                                    <h3>Banyaknya</h3>
                                    <p>{{ `${transactionsArray[selectedTransactionIndex].quantities[index]}
                                        ${product.data.unit}` }}
                                    </p>
                                </div>
                                <div class="card">
                                    <img :src="`${BACKEND_BASE_URL}/${product.data.filepath}`" width="100" alt="">
                                </div>
                            </div>
                            <div style="border-top: 1px solid grey; width: 100%; margin-top: 1rem"></div>

                        </div>
                    </div>

                </div>
                <div style="width: 100%;">
                    <div>
                        <h3>Total Harga</h3>
                        <p>{{ transactionsArray[selectedTransactionIndex].totalPrice }} ETH</p>
                    </div>
                    <div>
                        <h3>Ubah Status</h3>
                        <select name="status" id="status" v-model="status">
                            <option disabled value="">-- Pilih status --</option>
                            <option value="Proses">Proses</option>
                            <option value="Pengiriman">Pengiriman</option>
                            <option value="Selesai">Selesai</option>
                        </select>
                    </div>
                    <button class="btn-update"
                        @click.prevent="updateHandle(transactionsArray[selectedTransactionIndex].id)">Update</button>
                </div>
            </div>
        </div>
    </div>

    <SideBar v-if="user.role" :role="user.role"></SideBar>

    <div class="container">
        <NavBarDash :user="user.name" :role="user.role" :title="route.name"></NavBarDash>
        <div class="card-container" v-if="groupedTransactions.length != 0">
            <div v-for="(transactions, index) in groupedTransactions" :key="index" class="card-product"
                @click.prevent="selectedTransactionHandle(transactions, index)">
                <div style="display: flex; align-items: center; gap:2rem">
                    <div class="card-img">
                        <img :src="`${BACKEND_BASE_URL}/${transactions.products[0].data.filepath}`" alt="">
                    </div>
                    <div>
                        <h2>{{ transactions.products[0].data.product_name }} {{ transactions.products.length - 1 === 0
                            ? "" : `+${transactions.products.length - 1} lainnya` }}
                        </h2>
                        <p> {{ transactionsArray[index].buyer }}</p>
                        <h3>{{ transactionsArray[index].totalPrice }} ETH</h3>
                        <p>{{ transactionsArray[index].timestamp }}</p>
                        <p>Terverifikasi di Blockchain ✔️</p>
                    </div>
                </div>
                <h1 :class="getStatusClass(transactionsArray[index].status)">{{ transactionsArray[index].status }}</h1>
            </div>
        </div>
        <div class="card-container" v-else>
            <h2 style="text-align: center; padding: 3rem;">Belum ada pesanan</h2>
        </div>
    </div>

</template>

<style scoped>
.container {
    padding: 2rem;
    margin-left: 25%;
}

.card-container {
    background-color: white;
    border-radius: 1rem;
    box-shadow: rgba(0, 0, 0, 0.16) 0px 1px 4px;
    overflow: hidden;
}

.card-product {
    padding: 2rem;
    display: flex;
    align-items: center;
    justify-content: space-between;
    width: 100%;
    cursor: pointer;
    transition: 0.2s;
}

.card-product:hover {
    background-color: var(--background);
    color: white
}



.card-img {
    width: 100px;
    height: 100px;
}

.card-img img {
    width: 100%;
    height: 100%;
    object-fit: cover;
    border-radius: 0.5rem;
}

p {
    font-size: 1.2rem;
}

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
    width: 70%;
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
    justify-content: space-between;
    gap: 4rem;
}


.modal select {
    width: 100%;
    height: 3rem;
    border: none;
    outline: none;
    background-color: #f2f2f2;
    font-size: 1.2rem;
    padding: 0 1rem;
    border-radius: 1.5rem;
}

.btn-update {
    width: 100%;
    margin-top: 2rem;
    height: 3rem;
    border: none;
    border-radius: 1.5rem;
    background: var(--background);
    background: var(--gradient);
    color: white;
    font-size: 1.5rem;
    cursor: pointer;
}
</style>