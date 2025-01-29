<script setup>

const props = defineProps({
    title: String,
    value: Number,
    icon: String
});


</script>
<template>
    <div class="card">
        <div class="card-img">
            <img :src="`/assets/icon/${props.icon}`" width="80">
        </div>
        <div class="info">
            <h3>{{ props.title }}</h3>
            <p>{{ props.value }}</p>
        </div>
    </div>
</template>
<style scoped>
.card {
    width: 25rem;
    display: inline-flex;
    align-items: center;
    background-color: white;
    border-radius: 1rem;
    padding: 2rem;
    box-shadow: rgba(0, 0, 0, 0.16) 0px 1px 4px;
}

.card h3 {
    font-weight: 400;
    font-size: 1.3rem;
}

.card p {
    font-size: 2rem;
    font-weight: 600;
}
</style>