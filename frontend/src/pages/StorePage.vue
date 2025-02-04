<script setup>
import { reactive, ref, watch } from 'vue';

import axios from 'axios'
import { formatedDate } from '@/assets/script/formated-date';

const state = reactive({
    search: null,
    stores: null,
    store: null,
    filteredStore: null,
})

const storeManager = ref([])
const isFormModalOpen = ref(false)
const isModalOpen = ref(false)
const storeForm = ref({
    user_id: "",
    name: "",
    address: "",
    telp: "",
})

const BACKEND_BASE_URL = import.meta.env.VITE_BACKEND_BASE_URL

async function getAllManager() {
    try {
        const response = await axios.get(`${BACKEND_BASE_URL}/users/manager`)
        storeManager.value = response.data.data
    } catch (error) {

    }
}

async function getAllStore() {
    try {
        const response = await axios.get(`${BACKEND_BASE_URL}/store`)
        state.stores = response.data.data
        state.filteredStore = response.data.data
    } catch (error) {
        console.log(error)
    }
}

async function addStoreHandle() {
    try {
        await axios.post(`${BACKEND_BASE_URL}/store`, {
            user_id: storeForm.value.user_id,
            name: storeForm.value.name,
            address: storeForm.value.address,
            telp: storeForm.value.telp.toString(),
        })
        isFormModalOpen.value = false
        getAllStore()
    } catch (error) {
        console.log(error)
    }
}

async function getStoreById(id) {
    try {
        const response = await axios.get(`${BACKEND_BASE_URL}/store/${id}`)
        state.store = response.data.data
        console.log(response.data)
    } catch (error) {
        console.log(error)
    }
    isModalOpen.value = !isModalOpen.value
}

getAllManager()
getAllStore()

watch(
    () => state.search,
    (newValue) => {
        if (newValue) {
            state.filteredStore = state.stores.filter((item) =>
                newValue
                    .toLowerCase()
                    .split(" ")
                    .every(
                        (term) =>
                            item.name.toLowerCase().includes(term)
                    )
            );
        } else {
            state.filteredStore = state.stores;
        }
    },
    { immediate: true }
);
</script>

<template>

    <div v-if="isFormModalOpen" class="modal-overlay" @click.self="isFormModalOpen = !isFormModalOpen">
        <div class="modal" id="modal">
            <div style="display: flex; justify-content: space-between; align-items: center;">
                <h2>Tambah Toko</h2>
                <button @click="isFormModalOpen = !isFormModalOpen" class="btn-close">X</button>
            </div>
            <form action="">
                <label for="name">Nama Toko</label>
                <input type="text" name="name" id="neme" v-model="storeForm.name" required>

                <label for="user_id">Manajer Toko</label>
                <select name="user_id" id="user_id" v-model="storeForm.user_id">
                    <option v-for="(manager, index) in storeManager" :key="index" :value="manager.id">{{
                        manager.name }}</option>
                </select>

                <label for="address">Alamat Toko</label>
                <input type="text" name="address" id="address" v-model="storeForm.address" required>

                <label for="telp">No. Telp</label>
                <input type="number" name="telp" id="telp" v-model="storeForm.telp" required>

                <button @click.prevent="addStoreHandle">Tambah</button>
            </form>
        </div>
    </div>
    <div v-if="isModalOpen" class="modal-overlay" @click.self="isModalOpen = !isModalOpen">
        <div class="modal" id="modal">
            <div style="display: flex; justify-content: space-between; align-items: center;">
                <h2>Detail Distributor</h2>
                <button @click="isModalOpen = !isModalOpen" class="btn-close">X</button>
            </div>
            <div class="content">
                <h3>ID Toko</h3>
                <p>{{ state.store.id }}</p>
                <h3>Manajer</h3>
                <p>{{ state.store.manager_name }}</p>
                <h3>Nama Toko</h3>
                <p>{{ state.store.name }}</p>
                <h3>Alamat</h3>
                <p>{{ state.store.address }}</p>
                <h3>No. Telp</h3>
                <p>{{ state.store.telp }}</p>
                <h3>Ditambahkan Pada</h3>
                <p>{{ formatedDate(state.store.created_at) }}</p>
                <h3>Diperbaharui Pada</h3>
                <p>{{ formatedDate(state.store.updated_at) }}</p>
            </div>
        </div>
    </div>

    <div class="container">
        <h2>Daftar Toko</h2>
        <input type="text" placeholder="Cari Toko" v-model="state.search">
        <table v-if="state.stores != null">
            <thead>
                <tr>
                    <th style="width: 5rem;">No</th>
                    <th style="width: 20rem;">Toko</th>
                    <th style="width: 35rem;">Alamat</th>
                    <th style="text-align: center;">Aksi</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="(store, index) in state.filteredStore" :key="store.id">
                    <td>{{ index + 1 }}</td>
                    <td>{{ store.name }}</td>
                    <td>{{ store.address }}</td>
                    <td style="display: flex; justify-content: center;"><button
                            @click.prevent="getStoreById(store.id)">Lihat Detail</button></td>
                </tr>
            </tbody>
        </table>
        <button @click="isFormModalOpen = !isFormModalOpen">Tambah Toko</button>
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

.modal .btn-close {
    cursor: pointer;
    background-color: white;
    border: none;
    font-size: 2rem;
    font-weight: 900;
    color: #686868;
}
</style>