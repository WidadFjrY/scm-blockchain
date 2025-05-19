<script setup>
import { useRoute } from 'vue-router'
import { web3 } from '@/assets/script/eth-transaction.js'
import { ref, watchEffect } from 'vue'


const route = useRoute()
const txHash = route.params.txHash

const transaction = ref({})
const transactionParse = ref({})
const network = ref({})
const loading = ref(true)

watchEffect(async () => {
    if (txHash) {
        loading.value = true;
        const tx = await web3.eth.getTransaction(txHash);
        const netId = await web3.eth.net.getId();
        network.value = {
            netId: netId,
            netName: getNetworkName(netId)
        }

        transaction.value = Object.fromEntries(
            Object.entries(tx).map(([key, val]) => [key, typeof val === 'bigint' ? val.toString() : val])
        );

        loading.value = false;
    }

    if (transaction.value) {
        transactionParse.value = convertTransaction(transaction.value)
    }
});

function convertTransaction(tx) {
    function weiToEth(wei) {
        return (Number(wei) / 1e18).toFixed(6);
    }

    function normalize(value) {
        if (typeof value === 'bigint') return value.toString();
        if (!value) return "-";
        return value.toString();
    }

    return {
        "Tipe Transaksi": normalize(tx.type),
        "Hash Transaksi": normalize(tx.hash),
        "ID Rantai": normalize(tx.chainId),
        "Nonce": normalize(tx.nonce),
        "Hash Blok": normalize(tx.blockHash),
        "Nomor Blok": normalize(tx.blockNumber),
        "Indeks Transaksi di Blok": normalize(tx.transactionIndex),
        "Alamat Pengirim": normalize(tx.from),
        "Alamat Penerima": normalize(tx.to),
        "Nilai (ETH)": weiToEth(tx.value),
        "Max Priority Fee Per Gas (Gwei)": (Number(tx.maxPriorityFeePerGas) / 1e9).toFixed(2),
        "Max Fee Per Gas (Gwei)": (Number(tx.maxFeePerGas) / 1e9).toFixed(2),
        "Gas Price (Gwei)": (Number(tx.gasPrice) / 1e9).toFixed(2),
        "Gas Limit": normalize(tx.gas),
        "Input Data": normalize(tx.input),
    }
}

function getNetworkName(chainId) {
    const networks = {
        1: "Ethereum Mainnet",
        3: "Ropsten",
        4: "Rinkeby",
        5: "Goerli",
        42: "Kovan",
        1337: "Ganache (Default)",
        5777: "Ganache",
    };

    return networks[chainId] || `Jaringan Tidak Dikenal (Chain ID: ${chainId})`;
}


</script>
<template>
    <div class="container">
        <h1>Blok {{ transactionParse["Nomor Blok"] }} - Ethereum Lokal ({{ network.netName }}, Network ID: {{
            network.netId }})</h1>
        <div v-if="loading">Loading...</div>
        <div v-else-if="transaction">
            <div v-for="(value, key) in transactionParse" :key="key" class="transaction-item">
                <h3>{{ key }}</h3>
                <p>{{ value }}</p>
            </div>
        </div>
        <div v-else>
            <p>Transaksi tidak ditemukan.</p>
        </div>
    </div>
</template>
<style scoped>
.container {
    background-color: white;
    width: 60%;
    padding: 2rem;
    margin: auto;
    margin-top: 2rem;
    margin-bottom: 2rem;
    border-radius: 1rem;
    box-shadow: rgba(0, 0, 0, 0.16) 0px 1px 4px;
    overflow: hidden;
}

.container h3 {
    margin-top: 1rem;
}

.container p {
    font-size: 1.2rem;
    word-wrap: break-word;
    white-space: normal;
    overflow-wrap: break-word;
}
</style>