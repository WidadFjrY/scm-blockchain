<script setup>
import { ref, watch } from 'vue';

import axios from 'axios';

import SideBar from '@/components/SideBar.vue';
import NavBarDash from '@/components/NavBarDash.vue';
import { web3, SupplyChainContract } from '@/assets/script/eth-transaction.js'

import { useRouter, useRoute } from 'vue-router';

const route = useRoute()
const groupedTransaction = ref({})
const user = ref({
    name: "",
    ethAddr: "",
    role: "",
})
const isModalOpen = ref(false)
const userTrx = ref()
const selectedTransaction = ref()
const selectedDate = ref()
const selectedTransactionIndex = ref()
const products = ref([])
const BACKEND_BASE_URL = import.meta.env.VITE_BACKEND_BASE_URL
const dailyProfits = ref({})
const ethPrice = ref();
const receipt = ref()

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

function groupTransactionsByDate(transactions) {
    const grouped = {};

    transactions.forEach(tx => {
        const date = new Date(tx.timestamp * 1000).toISOString().split('T')[0];

        if (!grouped[date]) {
            grouped[date] = [];
        }

        grouped[date].push(tx);
    });

    return grouped;
}


async function getTransaction() {
    try {
        const accounts = await window.ethereum.request({ method: 'eth_requestAccounts' });
        const admin = accounts[0];

        const rawTxs = await SupplyChainContract.methods.getAllCompletedTransactions().call({ from: admin });

        const transactions = rawTxs.map(tx => ({
            ...tx,
            id: parseInt(tx.id),
            quantities: tx.quantities.map(q => parseInt(q)),
            totalPrice: web3.utils.fromWei(tx.totalPrice, 'ether'),
            timestamp: parseInt(tx.timestamp),
            blockNumber: Number(tx.blockNumber)
        }));

        groupedTransaction.value = groupTransactionsByDate(transactions);
        console.log(transactions);
    } catch (error) {
        console.error("Gagal mengambil transaksi selesai:", error);
    }
}

async function getTransactionReceipt(blockNumber) {
    const response = await axios.get(`${BACKEND_BASE_URL}/user/tx/${blockNumber}`)
    if (!response.data.data) {
        receipt.value = {}
        return
    }
    receipt.value = response.data.data
}


function maskAddress(address) {
    if (!address || address.length < 10) return address
    return `${address.slice(0, 5)}***${address.slice(-3)}`
}

async function getProduct(productId) {
    try {
        const response = await axios.get(`${BACKEND_BASE_URL}/product/${productId}`);
        products.value.push(response.data.data);
    } catch (error) {
        console.log(error);
    }
}


async function getDataUserByETHAddr(addr) {
    try {
        const response = await axios.get(`${BACKEND_BASE_URL}/user/eth_addr/${addr}`)
        userTrx.value = response.data.data
    } catch (error) {
        console.log(error)
    }
}

async function selectedTransactionHandle(transactions, date, index) {
    isModalOpen.value = true
    await getDataUserByETHAddr(transactions.buyer)
    await getTransactionReceipt(transactions.blockNumber)

    selectedTransaction.value = transactions
    selectedTransactionIndex.value = index
    selectedDate.value = date

    groupedTransaction.value[date][index].productIds.forEach(productId => {
        getProduct(productId)
    });
}

function clearProduct() {
    products.value.length = 0
}

function dailyTotals() {
    const totals = {};
    for (const [date, transactions] of Object.entries(groupedTransaction.value)) {
        if (Array.isArray(transactions)) {
            totals[date] = transactions.reduce((sum, t) => {
                // t.totalPrice bisa berupa string atau BigInt
                const price = typeof t.totalPrice === 'string'
                    ? parseFloat(t.totalPrice)
                    : fromWei(t.totalPrice);
                return sum + price;
            }, 0);
        } else {
            totals[date] = 0;
        }
    }
    return totals;
}

function fromWei(wei) {
    return Number(wei) / 1e18;
}

watch(dailyTotals, (val) => {
    dailyProfits.value = val
});


async function ETHPrice() {
    try {
        const response = await axios.get("https://api.coingecko.com/api/v3/simple/price?ids=ethereum&vs_currencies=idr");
        ethPrice.value = response.data.ethereum.idr;
    } catch (error) {
        ethPrice.value = "loading..."
        console.log(error)
    }
}

const convertToIDR = (ethAmount) => {
    return ethPrice.value ? Math.round(ethAmount * ethPrice.value) : "Loading...";
};



dailyTotals()
getTransaction()
getDataUser()
ETHPrice()

</script>

<template>
    <div v-if="isModalOpen" class="modal-overlay" @click.self="isModalOpen = !isModalOpen, clearProduct()">
        <div class="modal" id="modal">
            <div style="display: flex; justify-content: space-between; align-items: center;">
                <h2>Detail Product</h2>
                <button @click="isModalOpen = !isModalOpen, clearProduct()" class="btn-close">X</button>
            </div>
            <div class="content">
                <div style="width: 50%;">
                    <h3>ID Transaksi</h3>
                    <p>{{ groupedTransaction[selectedDate][selectedTransactionIndex].id }}</p>
                    <h3>Nama Pembeli</h3>
                    <p>{{ userTrx.name }}</p>
                    <h3>Alamat ETH</h3>
                    <p>{{ groupedTransaction[selectedDate][selectedTransactionIndex].buyer }}</p>
                    <h3>Alamat Pengiriman</h3>
                    <p>{{ groupedTransaction[selectedDate][selectedTransactionIndex].shippingAddress }}</p>
                    <h3>Produk:</h3>
                    <div style="max-height: 350px; overflow-y: auto;">
                        <div v-for="(product, index) in products" :key="index">
                            <div style="display: flex; justify-content: space-between; align-items: center;">
                                <div>
                                    <h3>Nama Produk</h3>
                                    <p>{{ product.product_name }}</p>
                                    <h3>Banyaknya</h3>
                                    <p>{{
                                        `${groupedTransaction[selectedDate][selectedTransactionIndex].quantities[index]}
                                        ${product.unit}` }}
                                    </p>
                                </div>
                                <div class="card">
                                    <img :src="`${BACKEND_BASE_URL}/${product.filepath}`" width="100" alt="">
                                </div>
                            </div>
                            <div style="border-top: 1px solid grey; width: 100%; margin-top: 1rem"></div>

                        </div>
                    </div>
                </div>
                <div style="width: 50%;">
                    <div>
                        <h3>Total Harga</h3>
                        <p>{{ groupedTransaction[selectedDate][selectedTransactionIndex].totalPrice }} ETH</p>
                        <h3>Status</h3>
                        <p>{{ groupedTransaction[selectedDate][selectedTransactionIndex].status }}</p>
                        <h3>Tanggal Pembelian</h3>
                        <p>{{ selectedDate }}</p>
                    </div>
                    <h2 style="margin-top: 2rem;">Informasi Blockchain</h2>
                    <div>
                        <h3>Transaction Hash</h3>
                        <p class="tx-hash">{{ receipt.tx_hash }}</p>
                    </div>
                    <div>
                        <h3>Nomor Blok</h3>
                        <p>{{ receipt.block_number }}</p>
                    </div>
                    <div style="margin-top: 1.5rem;">
                        <a :href="/verification/ + receipt.tx_hash" target="_blank" class="btn-verify">Lihat Detail
                            Blok</a>
                    </div>
                </div>
            </div>
        </div>
    </div>
    <SideBar v-if="user.role" :role="user.role"></SideBar>

    <div class="container">
        <NavBarDash :user="user.name" :role="user.role" :title="route.name"></NavBarDash>

        <div v-if="groupedTransaction" class="card-container">
            <table>
                <thead>
                    <tr>
                        <th>Tanggal</th>
                        <th>ID Transaksi</th>
                        <th>Pembeli</th>
                        <th>Total Harga (ETH)</th>
                        <th>Status</th>
                        <th>Aksi</th>
                    </tr>
                </thead>
                <template v-for="(transactions, date) in groupedTransaction" :key="date">
                    <tr v-for="(transaction, index) in transactions" :key="transaction.id">
                        <td v-if="index === 0" :rowspan="transactions.length">{{ date }}</td>
                        <td>{{ transaction.id }}</td>
                        <td>{{ maskAddress(transaction.buyer) }}</td>
                        <td>{{ transaction.totalPrice }}</td>
                        <td>{{ transaction.status }}</td>
                        <td style="display: flex; justify-content: center;"><button
                                @click.prevent="selectedTransactionHandle(transaction, date, index)">Lihat
                                Detail</button>
                        </td>
                    </tr>
                    <tr>
                        <td style="text-align: right; font-weight: bold;">Total Uang Masuk :</td>
                        <td style="font-weight: bold;">{{ dailyProfits[date] }} ETH / Rp.
                            {{ convertToIDR(dailyProfits[date]).toLocaleString("ID-id") }}</td>
                    </tr>
                </template>
            </table>
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

.btn-verify {
    padding: 0.5rem 1rem;
    border-radius: 0.5rem;
    margin-top: 2rem !important;
    background-color: #10B981;
    color: white;
    font-size: 1.2rem;
    cursor: pointer;
}

.tx-hash {
    word-wrap: break-word;
    white-space: normal;
    overflow-wrap: break-word;
}
</style>