<script setup>
const props = defineProps({
    img: String,
    productName: String,
    unit: String,
    stock: Number,
    brand: String,
    price: String,
})
</script>
<template>
    <div class="card">
        <div class="card-img">
            <img :src="props.img" alt="" srcset="">
        </div>
        <div style="display: flex; align-items: center; justify-content: space-between;">
            <h2>{{ props.productName }}</h2>
            <p style="color: gray;">{{ props.stock }} {{ props.unit }}</p>

        </div>
        <p>{{ props.brand }}</p>
        <div style="display: flex; align-items: center; gap:0.3rem">
            <p class="price">{{ props.price }} ETH </p>
            <p>/{{ props.unit }}</p>
        </div>
        <button>+ Keranjang</button>
    </div>
</template>
<style scoped>
.card {
    display: inline-block;
    box-shadow: rgba(0, 0, 0, 0.1) 0px 0px 5px 0px, rgba(0, 0, 0, 0.1) 0px 0px 1px 0px;
    border-radius: 0.5rem;
    padding: 1rem;
    margin-top: 2rem;
}

.card-img {
    width: 300px;
    height: 400px;
    margin-bottom: 1rem;
}

.card .card-img img {
    width: 100%;
    height: 100%;
    object-fit: cover;
    border-radius: 0.5rem;
}


.price {
    font-size: 1.3rem;
    font-weight: 700;
}

.card>button {
    width: 100%;
    margin-top: 1rem;
    padding: 0.5rem;
    border: none;
    border-radius: 0.5rem;
    background: var(--background);
    background: var(--gradient);
    color: white;
    font-size: 1.1rem;
    cursor: pointer;
}
</style>