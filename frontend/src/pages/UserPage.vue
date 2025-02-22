<script setup>
import { ref } from 'vue'
import axios from 'axios'

import userPict from '@/assets/img/user_default.png'

const user = ref({
    id: "",
    email: "",
    name: "",
    role: "",
    telp: "",
    url_profile: ""
})

const formUser = ref({
    email: "",
    name: "",
    telp: "",
})

const isFormModalOpen = ref(false)

const BASE_URL_BACKEND = import.meta.env.VITE_BACKEND_BASE_URL

async function getDataUser() {
    try {
        const response = await axios.get(`${BASE_URL_BACKEND}/user`)

        user.value = response.data.data

        formUser.value.email = user.value.email
        formUser.value.name = user.value.name
        formUser.value.telp = user.value.telp
    } catch (error) {
        console.log(error)
    }
}

async function updateUserHandle() {
    try {
        await axios.put(`${BASE_URL_BACKEND}/user`, {
            email: formUser.value.email,
            name: formUser.value.name,
            telp: formUser.value.telp
        })
        isFormModalOpen.value = false
        getDataUser()
    } catch (error) {
        console.log(error)
    }

    console.log(formUser.value)
}

getDataUser()
</script>

<template>
    <div v-if="isFormModalOpen" class="modal-overlay" @click.self="isFormModalOpen = !isFormModalOpen">
        <div class="modal" id="modal">
            <div style="display: flex; justify-content: space-between; align-items: center;">
                <h2>Perbaharui Profil</h2>
                <button @click="isFormModalOpen = !isFormModalOpen" class="btn-close">X</button>
            </div>
            <form action="">
                <label for="name">ID</label>
                <input type="text" name="id" id="id" disabled :value="user.id" style="color: rgb(203, 203, 203);">

                <label for="name">Nama</label>
                <input type="text" name="name" id="name" required v-model="formUser.name">

                <label for="name">Role</label>
                <input type="text" name="role" id="role" disabled :value="user.role" style="color: rgb(203, 203, 203);">

                <label for=" email">Email</label>
                <input type="email" name="email" id="email" required v-model="formUser.email">

                <label for="telp">Telp</label>
                <input type="text" name="telp" id="telp" required v-model="formUser.telp">

                <button @click.prevent="updateUserHandle">Tambah</button>
            </form>
        </div>
    </div>
    <div class="container">
        <div class="card-info">
            <div style="display: flex; gap:2rem">
                <div class="card-img">
                    <img :src="user.url_profile ? user.url_profile : userPict" alt="" srcset="">
                </div>
                <div class="info">
                    <p>ID</p>
                    <h1>{{ user.id }}</h1>
                    <p>Nama</p>
                    <h1>{{ user.name }}</h1>
                    <p>Role</p>
                    <h1>{{ user.role }}</h1>
                    <p>Email</p>
                    <h1>{{ user.email }}</h1>
                    <p>Telp</p>
                    <h1>{{ user.telp }}</h1>
                </div>
            </div>
            <div>
                <button @click="isFormModalOpen = !isFormModalOpen">Perbaharui Profil</button>
            </div>
        </div>
        <div class="card-form">
            <h1>Riwayat Pengguna</h1>
        </div>
    </div>
</template>

<style scoped>
.container {
    padding: 2rem;
}

.container .card-info {
    display: flex;
    justify-content: space-between;
}

.info p:nth-child(1) {
    margin-top: 0 !important;
}

.info p {
    margin-top: 0.5rem;
    font-size: 1.2rem;
}

.card-img {
    width: 20rem;
    background-color: white;
    aspect-ratio: 3 / 4;
    overflow: hidden;
    border-radius: 8px;
    box-shadow: rgba(0, 0, 0, 0.16) 0px 1px 4px;

}

.card-img img {
    width: 100%;
    height: 100%;
    object-fit: cover;
    object-position: center;
}

.card-form {
    margin-top: 2rem;
}

.container button {
    padding: 1rem;
    border-radius: 1rem;
    font-size: 1.3rem;
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
    width: 50%;
    padding: 3rem;
    border-radius: 1rem;
    transition: 0.5;
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
}

.modal label {
    font-weight: 600;
    display: block;
    margin-top: 1rem;
    font-size: 1.3rem;
}

.modal input {
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

.modal input::placeholder {
    color: #d5d5d5;
}

.modal form button {
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

.modal form button:hover {
    background: rgb(32, 193, 243);
    background: linear-gradient(90deg, rgba(32, 193, 243, 1) 23%, rgba(32, 125, 243, 1) 72%);

}

.modal .btn-close {
    cursor: pointer;
    background-color: white;
    border: none;
    font-size: 2rem;
    font-weight: 900;
    color: #686868;
}
</style>