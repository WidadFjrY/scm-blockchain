<script setup>
import NavBar from '@/components/NavBar.vue';

import { web3, SupplyChainContract } from '@/assets/script/eth-transaction.js'
import { ref } from 'vue';

import axios from 'axios';

const user = ref({})
const transactionsArray = ref([])
const products = ref([])
const groupedTransactions = ref([]);

const ethPrice = ref();
const BACKEND_BASE_URL = import.meta.env.VITE_BACKEND_BASE_URL

async function getDataUser() {
    try {
        const response = await axios.get(`${BACKEND_BASE_URL}/user`)
        user.value = response.data.data
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

async function getUserTransactions() {
    const account = await web3.eth.getAccounts();
    const userAddress = account[0];

    try {
        const transactions = await SupplyChainContract.methods
            .getTransactionsByBuyer(userAddress)
            .call();

        const productPromises = [];

        for (let i = 0; i < transactions.length; i++) {
            transactionsArray.value.push({
                id: transactions[i][0],
                buyer: transactions[i][1],
                productIds: transactions[i][2],
                quantities: transactions[i][3][0].toLocaleString(),
                totalPrice: transactions[i][5] == 'eth' ? `${web3.utils.fromWei(transactions[i][4], 'ether')} ETH` : `Rp. ${transactions[i][4].toLocaleString("id-ID")}`,
                timestamp: new Date(Number(transactions[i][7]) * 1000).toLocaleString(),
                status: transactions[i][8],
                txHash: await getTransactionReceipt(transactions[i][9])
            });

            transactions[i][2].forEach(productId => {
                productPromises.push(getProduct(transactions[i][0], productId));
            });
        }

        await Promise.all(productPromises);
        groupTransactions();
    } catch (error) {
        console.error("Gagal mengambil transaksi:", error);
    }
}


async function getTransactionReceipt(blockNumber) {
    const response = await axios.get(`${BACKEND_BASE_URL}/user/tx/${blockNumber}`)
    if (!response.data.data) {
        receipt.value = {}
        return
    }
    return response.data.data.tx_hash
}

async function getProduct(trxId, productId) {
    try {
        const response = await axios.get(`${BACKEND_BASE_URL}/product/${productId}`)
        products.value.push({ trxId: trxId, data: response.data.data })
    } catch (error) {
        console.log(error)
        return {}
    }
}

function getStatusClass(status) {
    return status.toLowerCase()
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
}

getUserTransactions()
ETHPrice()
getDataUser()

</script>
<template>
    <NavBar :name="user.name" :ethAddr="user.eth_addr" :role="user.role"></NavBar>
    <div class="container">
        <div class="head">
            <h1>Lacak Pesanan</h1>
            <div style="display: flex; align-items: center;">
                <img src="@/assets/img/eth-logo.png" width="50">
                <h2>1 ETH = Rp. {{ ethPrice.toLocaleString("id-ID") }}</h2>
            </div>
        </div>
        <div class="card-container" v-if="groupedTransactions.length > 0">
            <div v-for="(transactions, index) in groupedTransactions" :key="index" class="card-product">
                <div style="display: flex; align-items: center; gap:2rem; margin-top: 1rem;">
                    <div class="card-img">
                        <img :src="`${BACKEND_BASE_URL}/${transactions.products[0].data.filepath}`" alt="">
                    </div>
                    <div>
                        <h2>{{ transactions.products[0].data.product_name }} {{ transactions.products.length - 1 === 0
                            ? "" : `+${transactions.products.length - 1} lainnya` }}
                        </h2>
                        <p> {{ transactionsArray[index].buyer }}</p>
                        <h3>{{ transactionsArray[index].totalPrice }}</h3>
                        <p>{{ transactionsArray[index].timestamp }}</p>
                        <p>Terverifikasi di Blockchain ✔️</p>
                        <div style="margin-top: 1rem;">
                            <a :href="/verification/ + transactionsArray[index].txHash" target="_blank"
                                class="btn-verify">Lihat Detail
                                Blok</a>
                        </div>
                    </div>
                </div>
                <h1 :class="getStatusClass(transactionsArray[index].status)">{{ transactionsArray[index].status }}</h1>
            </div>
        </div>
        <div class="card-container" v-else>
            <h2 style="text-align: center; width: 100%;">Tidak ada transaksi</h2>
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
    margin-top: 2rem;
    display: flex;
    flex-wrap: wrap;
    gap: 1rem;
}

.head {
    display: flex;
    align-items: center;
    justify-content: space-between;
}

.card-product {
    display: flex;
    align-items: center;
    margin: 0 5rem;
    justify-content: space-between;
    width: 100%;
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