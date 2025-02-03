<script setup>
import axios from 'axios';

import { ref } from 'vue'
import { useRouter } from 'vue-router';

const router = useRouter()

const distributor = ref([]);
const fileName = ref("Pilih Gambar");

const formProduct = ref({
    product_name: "",
    distributor_id: "",
    brand: "",
    price: 0,
    unit: "",
    stock: 0,
    status: "",
    image: null
})

const BACKEND_BASE_URL = import.meta.env.VITE_BACKEND_BASE_URL

async function getAllDistributor() {
    try {
        const response = await axios.get(`${BACKEND_BASE_URL}/distributor`)
        response.data.data.forEach(data => {
            distributor.value.push(data)
        });
    } catch (error) {
        console.log(error)
    }
}

getAllDistributor()

function handleFileChange(event) {
    const selectedFile = event.target.files[0];
    if (selectedFile) {
        fileName.value = selectedFile.name;
        formProduct.value.image = selectedFile
    }
}

async function addProductHandle() {
    const formData = new FormData()

    formData.append("product_name", formProduct.value.product_name)
    formData.append("distributor_id", formProduct.value.distributor_id)
    formData.append("brand", formProduct.value.brand)
    formData.append("price", formProduct.value.price)
    formData.append("unit", formProduct.value.unit)
    formData.append("stock", formProduct.value.stock)
    formData.append("status", formProduct.value.status)
    formData.append("image", formProduct.value.image)

    try {
        const response = await axios.post(`${BACKEND_BASE_URL}/product`, formData, {
            headers: {
                "Content-Type": "multipart/form-data"
            }
        })
        console.log(response)
        router.push('/products')

    } catch (error) {
        console.log(error)
    }
}

</script>
<template>
    <div class="container">
        <h2>Informasi Produk</h2>
        <form action="">
            <div class="card">
                <div>
                    <label for="product_name">Nama Produk</label>
                    <input required type="text" name="product_name" id="product_name"
                        v-model="formProduct.product_name">
                </div>
                <div>
                    <label for="brand">Merek/Brand</label>
                    <input required type="text" name="brand" id="brand" v-model="formProduct.brand">
                </div>
            </div>
            <div class="card">
                <div>
                    <label for="distributor_id">Distributor</label>
                    <select name="distribtor_id" id="distributor_id" v-model="formProduct.distributor_id">
                        <option v-for="(distributor, index) in distributor" :key="index" :value="distributor.id">{{
                            distributor.name }}</option>
                    </select>
                </div>
                <div>
                    <label for="price">Harga</label>
                    <input required type="number" name="price" id="price" v-model="formProduct.price">
                </div>
            </div>

            <div class="card">
                <div>
                    <label for="stock">Stok</label>
                    <input required type="number" name="stock" id="stock" v-model="formProduct.stock">
                </div>
                <div>
                    <label for="unit">Satuan</label>
                    <input required type="text" name="unit" id="unit" v-model="formProduct.unit">
                </div>
            </div>
            <div class="card">
                <div>
                    <label for="status">Status</label>
                    <select name="status" id="status" v-model="formProduct.status">
                        <option value="Tersedia">Tersedia</option>
                        <option value="Hampir Habis">Hampir Habis</option>
                        <option value="Habis">Habis</option>
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
    background-color: white;
    margin-top: 2rem;
    padding: 2rem;
    box-shadow: rgba(0, 0, 0, 0.16) 0px 1px 4px;
    border-radius: 1rem;

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