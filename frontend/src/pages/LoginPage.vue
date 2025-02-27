<script setup>

import { ref } from 'vue'
import { useRouter } from 'vue-router'

import axios from 'axios'
import Cookies from 'js-cookie'

const BASE_URL_BACKEND = import.meta.env.VITE_BACKEND_BASE_URL

const router = useRouter()
const loginForm = ref({
    email: "",
    password: ""
});

const errorMsg = ref("")

async function loginHandle() {
    errorMsg.value = ""
    try {
        const response = await axios.post(`${BASE_URL_BACKEND}/user/login`, {
            email: loginForm.value.email,
            password: loginForm.value.password
        })

        const token = response.data.data.token

        Cookies.set('AUTH_TOKEN', token, { expires: 1, secure: true, sameSite: 'Strict' })
        window.location.href = '/'

    } catch (error) {
        errorMsg.value = error.response.data.message
    }
}
</script>

<template>
    <div class="login-container">
        <div class="login-card">
            <h1>Masuk</h1>
            <form action="">
                <label for="email">Email</label>
                <input v-model="loginForm.email" type="email" name="email" id="email" placeholder="Email">
                <label for="password">Kata Sandi</label>
                <input v-model="loginForm.password" type="password" name="password" id="password"
                    placeholder="Kata Sandi">
                <p class="error" v-if="errorMsg">{{ errorMsg }}</p>
                <button @click.prevent="loginHandle">Masuk</button>
            </form>
            <router-link to="/register"><button class="btn-reg">Daftar Sebagai Pelanggan</button></router-link>
        </div>
        <div class="info-card">
            <h2>Selamat Datang</h2>
            <p>Supply Chain Management Berbasis Blockchain</p>
            <img src="@/assets/img/icon-white.png">
        </div>
    </div>
</template>

<style scoped>
.login-container {
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

.login-card {
    width: 50%;
    height: 100%;
    padding: 5rem;
    display: flex;
    flex-direction: column;
    justify-content: center;

}

.login-card h1 {
    font-weight: 300;
    font-size: 3rem;
}

.login-card label {
    font-weight: 600;
    display: block;
    margin-top: 1rem;
}

.login-card input {
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

.login-card input::placeholder {
    color: #d5d5d5;
}

.login-card .error {
    color: red;
}

.login-card form button {
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

.login-card form button:hover {
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

.btn-reg {
    width: 100%;
    margin-top: 1rem;
    height: 3rem;
    border: 1px solid var(--background);
    border-radius: 1.5rem;
    background-color: white;
    color: var(--background);
    font-size: 1.5rem;
    cursor: pointer;
}
</style>