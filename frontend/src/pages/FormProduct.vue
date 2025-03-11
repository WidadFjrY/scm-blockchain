<script setup>
import axios from 'axios';

import { ref } from 'vue'
import { useRoute, useRouter } from 'vue-router';
import { web3, SupplyChainContract } from '../assets/script/eth-transaction';

import SideBar from '@/components/SideBar.vue';
import NavBarDash from '@/components/NavBarDash.vue';

const route = useRoute()
const router = useRouter()

const brands = ref([]);
const units = ref([]);
const fileName = ref("Pilih Gambar");

const formProduct = ref({
    product_name: "",
    brand_id: "",
    price: 0,
    unit_id: "",
    stock: 0,
    description: "",
    image: null
})

const BACKEND_BASE_URL = import.meta.env.VITE_BACKEND_BASE_URL

const user = ref({
    name: "",
    ethAddr: "",
    role: "",
})

async function getAllData() {
    try {
        const brandResponse = await axios.get(`${BACKEND_BASE_URL}/brands`)
        brandResponse.data.data.forEach(data => {
            brands.value.push(data)
        });

        const unitResponse = await axios.get(`${BACKEND_BASE_URL}/units`)
        unitResponse.data.data.forEach(data => {
            units.value.push(data)
        });

        const userResponse = await axios.get(`${BACKEND_BASE_URL}/user`)
        user.value.name = userResponse.data.data.name
        user.value.ethAddr = userResponse.data.data.eth_addr
        user.value.role = userResponse.data.data.role
    } catch (error) {
        console.log(error)
    }

}

getAllData()

function handleFileChange(event) {
    const selectedFile = event.target.files[0];
    if (selectedFile) {
        fileName.value = selectedFile.name;
        formProduct.value.image = selectedFile
    }
}

async function addProductHandle() {
    const formData = new FormData()
    const account = await web3.eth.getAccounts()
    const userAddress = account[0]

    formData.append("product_name", formProduct.value.product_name)
    formData.append("brand_id", formProduct.value.brand_id)
    formData.append("price", formProduct.value.price)
    formData.append("unit_id", formProduct.value.unit_id)
    formData.append("stock", formProduct.value.stock)
    formData.append("description", formProduct.value.description)
    formData.append("image", formProduct.value.image)

    try {
        const response = await axios.post(`${BACKEND_BASE_URL}/product/add`, formData, {
            headers: {
                "Content-Type": "multipart/form-data"
            }
        })

        if (response.status === 201) {
            const txReceipt = await SupplyChainContract.methods
                .addProduct(response.data.data.product_id, formProduct.value.product_name, web3.utils.keccak256(`${formProduct.value.product_name, formProduct.value.brand_id, formProduct.value.unit_id}`))
                .send({ from: userAddress })

            if (!txReceipt.status) {
                throw new Error("Transaksi blockchain gagal")
            }
        }

        router.push('/products')
    } catch (error) {
        console.log(error)
    }
}

</script>
<template>
    <SideBar></SideBar>
    <div class="container">
        <NavBarDash :user="user.name" :role="user.role" :title="route.name"></NavBarDash>
        <h2>Informasi Produk</h2>
        <form action="">
            <div class="card">
                <div>
                    <label for="product_name">Nama Produk</label>
                    <input required type="text" name="product_name" id="product_name"
                        v-model="formProduct.product_name">
                </div>
                <div>
                    <label for="brand_id">Merek</label>
                    <select name="brand_id" id="brand_id" v-model="formProduct.brand_id">
                        <option v-for="(brand, index) in brands" :key="index" :value="brand.id">{{
                            brand.brand_name }}</option>
                    </select>
                </div>
            </div>
            <div class="card">
                <div>
                    <label for="price">Harga</label>
                    <input required type="number" name="price" id="price" v-model="formProduct.price">
                </div>
                <div>
                    <label for="stock">Stok</label>
                    <input required type="number" name="stock" id="stock" v-model="formProduct.stock">
                </div>
            </div>

            <div class="card">
                <div>
                    <label for="brand_id">Satuan</label>
                    <select name="brand_id" id="brand_id" v-model="formProduct.unit_id">
                        <option v-for="(unit, index) in units" :key="index" :value="unit.id">{{
                            unit.unit_name }}</option>
                    </select>
                </div>
                <div>
                    <label>Gambar</label>
                    <div class="input" for="image">
                        <label for="image">{{ fileName }}</label>
                    </div>
                    <input required type="file" name="image" id="image" accept="image/*" @change="handleFileChange">
                </div>
            </div>
            <div class="card-textarea">
                <div>
                    <label for="description">Deskripsi</label>
                    <textarea type="text" name="description" id="description" v-model="formProduct.description"
                        required></textarea>
                </div>

            </div>
            <button @click.prevent="addProductHandle">Tambah Produk</button>
        </form>
    </div>
</template>
<style scoped>
#image {
    display: none;
}

.input {
    display: flex;
    align-items: center;
    padding-left: 1rem;
    width: 100% !important;
    height: 3rem;
    border-radius: 1.5rem;
    border: none;
    outline: none;
    color: var(--dark-color);
    background-color: #f2f2f2;
    margin-top: 0.5rem;
}

.input label {
    font-size: 1.2rem !important;
    font-weight: 400 !important;
    width: 100%;
}

.container {
    padding: 2rem;
    margin-left: 25%;
}

.container form {
    margin-top: 1rem;
    background-color: white;
    padding: 2rem;
    border-radius: 1rem;
    box-shadow: rgba(0, 0, 0, 0.16) 0px 1px 4px;

}

.container .card:nth-child(1) {
    margin-top: 0 !important;
}

.container .card {
    display: flex;
    gap: 1rem;
    margin-top: 2rem;
}

.container .card div {
    width: 50%;
}

.container label:nth-child(1) {
    margin-top: 0 !important;
}

.container label {
    display: block;
    font-size: 1.3rem;
    font-weight: 600;
}

.container input {
    width: 100%;
    height: 3rem;
    border-radius: 1.5rem;
    border: none;
    outline: none;
    font-size: 1.2rem;
    padding: 1rem;
    color: var(--dark-color);
    background-color: #f2f2f2;
    margin-top: 0.5rem;
}

.container .card-textarea {
    margin-top: 2rem;
    width: 100% !important;
}

textarea {
    width: 100%;
    font-family: 'Poppins', sans-serif !important;
    resize: none;
    border-radius: 1.5rem;
    border: none;
    outline: none;
    font-size: 1.2rem;
    padding: 1rem;
    color: var(--dark-color);
    background-color: #f2f2f2;
    margin-top: 0.5rem;
}

.container input::placeholder {
    color: #d5d5d5;
}

.container select {
    width: 100%;
    height: 3rem;
    border: none;
    outline: none;
    background-color: #f2f2f2;
    font-size: 1.2rem;
    padding: 0 1rem;
    border-radius: 1.5rem;
    margin-top: 0.5rem;

}

.container button {
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

.container button:hover {
    background: rgb(32, 193, 243);
    background: linear-gradient(90deg, rgba(32, 193, 243, 1) 23%, rgba(32, 125, 243, 1) 72%);

}
</style>