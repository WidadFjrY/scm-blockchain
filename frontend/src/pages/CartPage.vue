<script setup>
import NavBar from '@/components/NavBar.vue';

import { web3, SupplyChainContract } from '@/assets/script/eth-transaction.js'
import { ref } from 'vue';

import axios from 'axios';

const user = ref({
    name: "",
    ethAddr: "",
    role: "",
})
const ethPrice = ref();
const paymentMethod = ref("eth");
const carts = ref([]);
const shippingAddress = ref("");

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

    if (paymentMethod.value == "eth") {
        try {
            const valueInWei = web3.utils.toWei(convertToETH(totalPrice()).toString(), 'ether');
            const tx = await SupplyChainContract.methods
                .createTransaction(productIds, quantities, shippingAddress.value, "eth")
                .send({ from: userAddress, value: valueInWei });

            for (const cart of carts.value) {
                await deleteCartHandle(cart.product_id);
            }

            for (let i = 0; i < productIds.length; i++) {
                await axios.put(`${BACKEND_BASE_URL}/product/stock`, {
                    product_id: productIds[i],
                    stock_in: 0,
                    stock_out: quantities[i]
                });
            }

            const receipt = await web3.eth.getTransactionReceipt(tx.transactionHash);

            try {
                await axios.post(`${BACKEND_BASE_URL}/user/tx`, {
                    tx_hash: tx.transactionHash,
                    block_address: receipt.blockHash,
                    block_number: Number(receipt.blockNumber),
                })
            } catch (errorAPI) {
                console.log(errorAPI)
            }
            window.location.href = "/"
        } catch (error) {
            console.error("Transaksi gagal:", error);
        }
    } else {
        try {
            const tx = await SupplyChainContract.methods
                .createTransactionWithOtherPayments(productIds, quantities, shippingAddress.value, totalPrice(), "cash")
                .send({ from: userAddress });

            for (const cart of carts.value) {
                await deleteCartHandle(cart.product_id);
            }

            for (let i = 0; i < productIds.length; i++) {
                await axios.put(`${BACKEND_BASE_URL}/product/stock`, {
                    product_id: productIds[i],
                    stock_in: 0,
                    stock_out: quantities[i]
                });
            }

            const receipt = await web3.eth.getTransactionReceipt(tx.transactionHash);

            try {
                await axios.post(`${BACKEND_BASE_URL}/user/tx`, {
                    tx_hash: tx.transactionHash,
                    block_address: receipt.blockHash,
                    block_number: Number(receipt.blockNumber),
                })
            } catch (errorAPI) {
                console.log(errorAPI)
            }
            window.location.href = "/"
        } catch (error) {
            console.error("Transaksi gagal:", error);
        }
    }
}

async function deleteCartHandle(productId) {
    try {
        await axios.delete(`${BACKEND_BASE_URL}/cart/${productId}`)
    } catch (error) {
        console.log(error)
    }
    carts.value = carts.value.filter(cart => cart.product_id !== productId);
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
        <div class="card-container" v-if="carts">
            <div v-for="(cart, index) in carts" :key="index">
                <div class="card">
                    <div style="display: flex;gap:1rem ;width: 20rem;">
                        <div class="card-img">
                            <img :src="`${BACKEND_BASE_URL}/${cart.product.filepath}`" alt="">
                        </div>
                        <div>
                            <h2>{{ cart.product.product_name }}</h2>
                            <p>{{ cart.product.brand }}</p>
                            <p>Stok: {{ cart.product.stock }} {{ cart.product.unit }}</p>
                        </div>
                    </div>
                    <div style="display: flex; align-items: center; gap:1rem">
                        <div class="card-btn-qty">
                            <button @click.prevent="cart.quantity--, updateQty(cart.product_id, cart.quantity)"
                                :disabled="cart.quantity <= 1">-</button>
                            <p>{{ cart.quantity }}</p>
                            <button @click.prevent="cart.quantity++, updateQty(cart.product_id, cart.quantity)"
                                :disabled="cart.quantity >= cart.product.stock">+</button>
                        </div>
                    </div>
                    <div style="display: flex; align-items: center; gap: 1rem;">
                        <div class="card-price">
                            <p>Rp. {{ (cart.product.price * cart.quantity).toLocaleString("id-ID") }}</p>
                            <h3>{{ convertToETH(cart.product.price * cart.quantity) }} ETH</h3>
                        </div>
                        <img src="/assets/icon/Trash.png" alt="" width="40"
                            style="cursor: pointer !important; pointer-events: all;"
                            @click.prevent="deleteCartHandle(cart.product_id)">
                    </div>
                </div>
            </div>
            <label for="shippingAddress" class="shipping">Alamat Pengiriman</label>
            <input type="text" id="shippingAddress" v-model="shippingAddress" class="shippingInput">
            <div class="radio-button">
                <h3>Pilih Metode Pembayaran</h3>
                <label :class="paymentMethod == 'eth' ? 'active' : ''">
                    <input type="radio" value="eth" v-model="paymentMethod" />
                    <div style="display: flex; align-items: center; justify-content: space-between;">
                        <div style="display: flex; align-items: center; gap: 1rem;">
                            <img src="@/assets/img/eth-logo.png" width="40" alt="">
                            <span>ETH</span>
                        </div>
                        <span style="text-align: end;">{{ convertToETH(totalPrice()) }} ETH</span>
                    </div>
                </label>
                <label :class="paymentMethod == 'cash' ? 'active' : ''">
                    <input type="radio" value="cash" v-model="paymentMethod" />
                    <div style="display: flex; align-items: center; justify-content: space-between;">
                        <div style="display: flex; align-items: center; gap: 1rem;">
                            <img src="@/assets/img/money.png" alt="">
                            <span>Cash (Bayar di Toko)</span>
                        </div>
                        <span style="text-align: end;">Rp. {{ totalPrice().toLocaleString("ID-id") }}</span>

                    </div>
                </label>
            </div>
            <h2 style="text-align: end; margin-top: 2rem;" v-if="paymentMethod == 'eth'">Total (ETH): {{
                convertToETH(totalPrice()) }} ETH</h2>
            <h2 style="text-align: end; margin-top: 2rem;" v-else>Total (IDR): Rp. {{
                totalPrice().toLocaleString("ID-id") }}</h2>
            <div class="card-btn">
                <button @click.prevent="checkOutHandle" :disabled="!shippingAddress">Check Out</button>
            </div>
        </div>
        <div v-else class="card-container">
            <h2 style="text-align: center;">Tidak ada produk</h2>
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

.card-btn-qty {
    display: flex;
    align-items: center;
}

.card-btn-qty p {
    padding: 1rem;
    text-align: center;
}

.card-btn-qty button {
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

button:disabled {
    cursor: not-allowed;
    opacity: 0.4;
}

.shipping {
    display: block;
    font-size: 1.3rem;
    font-weight: 600;
}

.shippingInput {
    width: 100%;
    height: 3rem;
    border-radius: 0.5rem;
    border: none;
    outline: none;
    font-size: 1.2rem;
    padding: 1rem;
    color: var(--dark-color);
    background-color: #f2f2f2;
    margin-top: 0.5rem;
}

.radio-button h3 {
    margin-top: 2rem;
    display: block;
    font-size: 1.3rem;
    font-weight: 600;
}

.radio-button label {

    margin-top: 1rem;
    font-size: 1.3rem;
    padding: 1rem;
    display: block;
    background-color: white;
    border: 1px solid rgb(209, 209, 209);
    border-radius: 0.5rem;
    cursor: pointer;
}

.active {
    background-color: rgb(234, 246, 255) !important;
    border: 1px solid var(--background) !important;
    cursor: default !important;
}

.radio-button input {
    display: none;
}
</style>