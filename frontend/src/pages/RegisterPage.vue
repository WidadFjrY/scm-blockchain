<script setup>

import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { web3, userContract } from '@/assets/script/eth-transaction.js'
import { v4 as uuidv4 } from 'uuid'

import axios from 'axios'
import Cookies from 'js-cookie'

const BASE_URL_BACKEND = import.meta.env.VITE_BACKEND_BASE_URL

const router = useRouter()
const registerForm = ref({
    name: "",
    email: "",
    telp: "",
    password: "",
    verify_password: "",
});

const errorMsg = ref("")

const token = Cookies.get('AUTH_TOKEN')

if (token) {
    checkToken()
}

async function checkToken() {
    try {
        axios.defaults.headers.common['Authorization'] = `Bearer ${token}`
        const response = await axios.get(`${BASE_URL_BACKEND}/checkToken`)
        if (response.status == 200) {
            router.push('/')
        }
    } catch (error) {
        router.push('/login')
    }
}

async function registerHandle() {
    errorMsg.value = ""
    try {

        const response = await axios.post(`${BASE_URL_BACKEND}/user/validate`, {
            name: registerForm.value.name,
            email: registerForm.value.email,
            eth_address: "userAddress",
            telp: registerForm.value.telp,
            password: registerForm.value.password,
            verify_password: registerForm.value.verify_password
        })

        if (response.data.data) {
            const account = await web3.eth.getAccounts()
            const userAddress = account[0]

            const txReceipt = await userContract.methods
                .registerUser(registerForm.value.name, "Customer")
                .send({ from: userAddress })

            if (!txReceipt.status) {
                throw new Error("Transaksi blockchain gagal")
            }

            await axios.post(`${BASE_URL_BACKEND}/user/register`, {
                name: registerForm.value.name,
                email: registerForm.value.email,
                eth_address: userAddress,
                telp: registerForm.value.telp,
                password: registerForm.value.password,
                verify_password: registerForm.value.verify_password
            })
            router.push('/login')
        }
    } catch (error) {
        errorMsg.value = error.response.data.message
    }
}

</script>

<template>
    <div class="register-container">
        <div class="register-card">
            <h1>Daftar</h1>
            <form action="">
                <label for="name">Nama</label>
                <input v-model="registerForm.name" type="text" name="name" id="name" placeholder="Nama" required>
                <label for="email">Email</label>
                <input v-model="registerForm.email" type="email" name="email" id="email" placeholder="Email" required>
                <label for="email">No. Telp</label>
                <input v-model="registerForm.telp" type="text" name="telp" id="telp" placeholder="No. Telp" required>
                <label for="password">Kata Sandi</label>
                <input v-model="registerForm.password" type="password" name="password" id="password"
                    placeholder="Kata Sandi" required>
                <label for="password">Ulangi Kata Sandi</label>
                <input v-model="registerForm.verify_password" type="password" name="verify_password"
                    id="verify Password" placeholder="Ulangi Kata Sandi" :disabled="registerForm.password === ''"
                    required>
                <p class="error" v-if="errorMsg">{{ errorMsg }}</p>
                <button @click.prevent="registerHandle">Daftar</button>
            </form>
        </div>
        <div class="info-card">
            <h2>Selamat Datang</h2>
            <p>Supply Chain Management Berbasis Blockchain</p>
            <img src="@/assets/img/icon-white.png">
        </div>
    </div>
</template>

<style scoped>
.register-container {
    margin: auto;
    margin-top: 7rem;
    background-color: white;
    width: 80vw;
    height: 70vh;
    display: flex;
    justify-content: space-between;
    align-items: center;
    border-radius: 2rem;
    overflow: hidden;
    box-shadow: rgba(0, 0, 0, 0.1) 0px 4px 12px;
}

.register-card {
    width: 50%;
    height: 100%;
    padding: 5rem;
    display: flex;
    flex-direction: column;
    justify-content: center;

}

.register-card h1 {
    font-weight: 300;
    font-size: 3rem;
}

.register-card label {
    font-weight: 600;
    display: block;
    margin-top: 1rem;
}

.register-card input {
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

.register-card input::placeholder {
    color: #d5d5d5;
}

.register-card .error {
    color: red;
}

.register-card button {
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

.register-card button:hover {
    background: rgb(32, 193, 243);
    background: linear-gradient(90deg, rgba(32, 193, 243, 1) 23%, rgba(32, 125, 243, 1) 72%);

}

.info-card {
    background: var(--background);
    background: var(--gradient);
    height: 100%;
    width: 50%;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
}

.info-card h2 {
    font-size: 3rem;
    color: white;
}

.info-card p {
    font-size: 1.3rem;
    color: white;
    text-align: center;
}


.info-card img {
    margin-top: 3rem;
    width: 200px;
}
</style>