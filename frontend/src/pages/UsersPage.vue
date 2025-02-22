<script setup>
import axios from 'axios';
import { reactive, watch, ref } from 'vue';

const state = reactive({
    search: null,
    users: null,
    user: null,
    filteredusers: null,
})

const isFormModalOpen = ref(false)
const isModalOpen = ref(false)
const errMsg = ref("")
const userForm = ref({
    email: "",
    role: "",
    password: "",
    verify_password: ""
})

const BACKEND_BASE_URL = import.meta.env.VITE_BACKEND_BASE_URL

async function addUser() {
    try {
        console.log(userForm)
        await axios.post(`${BACKEND_BASE_URL}/user/register`, {
            email: userForm.value.email,
            role: userForm.value.role,
            password: userForm.value.password,
            verify_password: userForm.value.verify_password,
        })
        getAllUser()
        isFormModalOpen.value = false
    } catch (error) {
        errMsg.value = error.response.data.message
    }
}

async function getUserById(id) {

}

async function getAllUser() {
    try {
        const response = await axios.get(`${BACKEND_BASE_URL}/users`)
        state.users = response.data.data
        state.filteredusers = response.data.data
    } catch (error) {
        console.log(error)
    }
    console.log(state.filteredusers)
}

getAllUser()

watch(
    () => state.search,
    (newValue) => {
        if (newValue) {
            state.filteredusers = state.users.filter((item) =>
                newValue
                    .toLowerCase()
                    .split(" ")
                    .every(
                        (term) =>
                            item.name.toLowerCase().includes(term) ||
                            item.email.toLowerCase().includes(term)
                    )
            );
        } else {
            state.filteredusers = state.users;
        }
    },
    { immediate: true }
);

</script>

<template>
    <div v-if="isFormModalOpen" class="modal-overlay" @click.self="isFormModalOpen = !isFormModalOpen">
        <div class="modal" id="modal">
            <div style="display: flex; justify-content: space-between; align-items: center;">
                <h2>Tambah Pengguna</h2>
                <button @click="isFormModalOpen = !isFormModalOpen" class="btn-close">X</button>
            </div>
            <p style="color: red;" v-if="errMsg">{{ errMsg }}</p>

            <form action="">
                <label for="email">Email</label>
                <input type="email" name="email" id="email" v-model="userForm.email" required>

                <label for="role">Peran</label>

                <select name="role" id="role" v-model="userForm.role">
                    <option value="Admin">Admin</option>
                    <option value="Distributor">Distributor</option>
                    <option value="Store Manager">Manajer Toko</option>
                </select>

                <label for="password">Kata Sandi</label>
                <input type="password" name="password" id="password" v-model="userForm.password" required>

                <label for="verify_password">Ulangi Kata Sandi</label>
                <input type="password" name="verify_password" id="verify_password" v-model="userForm.verify_password"
                    required>

                <button @click.prevent="addUser">Tambah</button>
            </form>
        </div>
    </div>
    <div v-if="isModalOpen" class="modal-overlay" @click.self="isModalOpen = !isModalOpen">
        <div class="modal" id="modal">
            <div style="display: flex; justify-content: space-between; align-items: center;">
                <h2>Detail Pengguna</h2>
                <button @click="isModalOpen = !isModalOpen" class="btn-close">X</button>
            </div>
            <div class="content">

            </div>
        </div>
    </div>
    <div class="container">
        <h2>Daftar Pengguna</h2>
        <input type="text" placeholder="Cari Distributor" v-model="state.search" @input="filteredusers">
        <table v-if="state.users != null">
            <thead>
                <tr>
                    <th style="width: 5rem;">No</th>
                    <th style="width: 15rem;">Nama</th>
                    <th style="width: 25rem;">Email</th>
                    <th style="width: 15rem;">Role</th>
                    <th style="text-align: center;">Aksi</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="(user, index) in state.filteredusers" :key="user.id">
                    <td>{{ index + 1 }}</td>
                    <td>{{ user.name }}</td>
                    <td>{{ user.email }}</td>
                    <td>{{ user.role }}</td>
                    <td style="display: flex; justify-content: center;"><button
                            @click.prevent="getUserById(user.id)">Lihat Detail</button></td>
                </tr>
            </tbody>
        </table>
        <button @click="isFormModalOpen = !isFormModalOpen">Tambah Pengguna</button>
    </div>
</template>

<style scoped>
.container {
    padding: 2rem;
    margin-top: 1rem;
}

.container input {
    padding: 1rem;
    margin: 1rem 0;
    color: var(--dark-color);
    border-radius: 1rem;
    background-color: white;
    border: none;
    outline: none;
    width: 100%;
    font-size: 1.5rem;
    box-shadow: rgba(0, 0, 0, 0.16) 0px 1px 4px;
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
    text-align: start;
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
    border-radius: 1rem;
    font-size: 1.3rem;
    background-color: rgb(32, 193, 243);
    color: white;
    border: none;
    cursor: pointer;
    transition: 0.5s ease;
}


.container>button {
    margin-top: 2rem;
    padding: 1rem;
    border-radius: 1rem;
    font-size: 1.3rem;
    background-color: var(--background);
    color: white;
    border: none;
    cursor: pointer;
    transition: 0.5s ease;
}

button:hover {
    background-color: rgb(55, 208, 255);
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